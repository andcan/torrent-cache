Torrent Cache
=============
[![Build Status](https://travis-ci.org/k4rr1m/torrent-cache.png)](https://travis-ci.org/k4rr1m/torrent-cache)

Simple application to cache torrents written in Go.

Requirements
------------
`archivemount` command is not really required but `start.sh` and `stop.sh` won't work without it. They also expect a file named `cached.tar.gz`. Mount point is `$HOME/.torrent_cache`.


Compiling
---------
Run
```
    ./setup.sh
```
or
```
    go build torrent_cache.go
```
Api
---
Current version is v1.
Api URL is http://{hostname}/api/{version}/{torrent_infohash}.
* Cache torrent submitting a __POST__ request with a form containing the file with key __torrent__. Returns torrent's infohash.
* To retrieve cached torrents submit a __GET__ request to api URL
* Deletion is implemented but disabled.
