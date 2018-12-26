#!/bin/sh
set -x
set -e

# Create user for Gin
addgroup -S ginmusic
adduser -G ginmusic -H -D -g 'ginmusic User' ginmusic -h /data/ginmusic -s /bin/bash && usermod -p '*' ginmusic && passwd -u ginmusic
echo "export GINMUSIC_CUSTOM=${GINMUSIC_CUSTOM}" >> /etc/profile

# Final cleaning
rm /app/ginmusic/docker/finalize.sh
rm /app/ginmusic/docker/nsswitch.conf
