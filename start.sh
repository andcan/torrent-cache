#!/bin/bash
echo "Mounting cached.tar.gz @ $HOME/.torrent_cache"
./sh/mount.sh
echo "Running server"
./torrent_cache & echo $! > torrent_cache.pid
