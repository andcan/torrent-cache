import os
import sys
from subprocess import Popen, PIPE

def call(args) :

    p = Popen(args, stdin=PIPE, stdout=PIPE, stderr=PIPE)
    for line in p.stdout:
        print(line)
    return

if not os.path.isfile("cached.tar.gz") :
    call(["tar", "cfzT", "cached.tar.gz", "/dev/null"])

dir = os.getenv("HOME") + "/.torrent_cache"
if not os.path.exists(dir) :
    os.makedirs(dir)

call(["godep", "get"])
call(["go", "build"])
call(["go", "install"])

