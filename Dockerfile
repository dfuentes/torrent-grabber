FROM ubuntu
# comment
COPY ./torrent-grabber /bin/torrent-grabber

ENTRYPOINT ["/bin/torrent-grabber", "--config", "/mnt/config.yml"]
