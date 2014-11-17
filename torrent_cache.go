package main

import (
	"github.com/marksamman/bencode"
	
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"time"
)
var TORRENT_DIR string

func apiV1Handler(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	id := strings.Replace(uri, "/api/v1/", "", -1)
	switch r.Method {
		case "GET":
			path := strings.Join([]string{TORRENT_DIR, string(os.PathSeparator),
				id, ".torrent"}, "")
			_, err := os.Stat(path)
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				return
			} else if nil != err {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			content, err := ioutil.ReadFile(path)
			if nil != err {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			header := w.Header()
			header.Add("Content-Disposition", "inline; filename=" + id  +
				".torrent")
			header.Add("Content-Type", "application/x-torrent")
			w.WriteHeader(http.StatusOK)
			w.Write(content)
			break
		case "DELETE":
			w.WriteHeader(http.StatusMethodNotAllowed)
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
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			var buf bytes.Buffer
			buf.ReadFrom(file)
			bs := buf.Bytes()
			decoded, err := bencode.Decode(bytes.NewReader(bs))
			if (nil != err) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			hash := sha1.New()
			_, err = hash.Write(bencode.Encode(decoded["info"]))
			if nil != err {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))				
				return
			}
			id := hex.EncodeToString(hash.Sum(nil))
			path := strings.Join([]string{TORRENT_DIR, string(os.PathSeparator),
				id, ".torrent"}, "")
			
			err = ioutil.WriteFile(path, bs, 0644)
			if nil != err {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(id))
			break
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			break
	}
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
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	http.HandleFunc("/api/v1/", apiV1Handler)
	log.Fatal(server.ListenAndServe())
}
