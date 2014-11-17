import os
from subprocess import call

if not os.path.isfile("cached.tar.gz") :
    call(["tar", "cfzT", "cached.tar.gz", "/dev/null"])
