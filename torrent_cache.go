// torrent_cache project main.go
package main

import (
	//bencode "github.com/jackpal/bencode-go"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

var TORRENT_DIR string

func apiV1Handler(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	id := strings.Replace(uri, "/api/v1/", "", -1)
	switch r.Method {
		case "GET":
			path := TORRENT_DIR + string(os.PathSeparator) + id + ".torrent"
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				w.WriteHeader(404)
				return
			} else if nil != err {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
			content, err := ioutil.ReadFile(path)
			if nil != err {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
			header := w.Header()
			header.Add("Content-Disposition", "inline; filename=" + id  + ".torrent")
			header.Add("Content-Type", "application/x-torrent")
			w.Write(content)
			break
		case "DELETE":
			/*path := TORRENT_DIR + string(os.PathSeparator) + id + ".torrent"
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				w.WriteHeader(404)
				return
			} else if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return			
			}
			err = os.Remove(path)
			if nil != err {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
			*/
			break
		case "POST":
			file, _, err := r.FormFile("torrent")
			if nil != err {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
			var buf bytes.Buffer
			buf.ReadFrom(file)
			bs := buf.Bytes()
			hash := sha256.New()
			_, err = hash.Write(bs)
			if nil != err {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
			id := hex.EncodeToString(hash.Sum(nil))
			path := TORRENT_DIR + string(os.PathSeparator) + id + ".torrent"
			
			err = ioutil.WriteFile(path, bs, 0644)
			if nil != err {
				w.Write([]byte(err.Error()))
				w.WriteHeader(500)
				return
			}
			w.Write([]byte(id))
			break
	}
	w.WriteHeader(200)
}

func main() {
	u, err := user.Current()	
	if nil != err {
		panic(err)
	}
	TORRENT_DIR = strings.Join([]string{u.HomeDir, "/.torrent_cache"}, "")
	
	if _, err := os.Stat(TORRENT_DIR); os.IsNotExist(err) || nil != err {
		panic(err)
	}
	http.HandleFunc("/api/v1/", apiV1Handler)
	http.ListenAndServe(":8080", nil)
}
