echo "Stopping server"
kill $(cat torrent_cache.pid)
rm torrent_cache.pid
echo "Unmounting cached.tar.gz @ $HOME/.torrent_cache"
./sh/umount.sh
