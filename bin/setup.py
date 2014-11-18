import os
from subprocess import call

if not os.path.isfile("cached.tar.gz") :
    call(["tar", "cfzT", "cached.tar.gz", "/dev/null"])

dir = os.getenv("HOME") + "/.torrent_cache"
if not os.path.isfile(dir) :
    os.makedirs(dir)

call(["godep", "get"])
call(["go", "build", "torrent_cache.go"])

