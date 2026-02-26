#!/bin/sh
set -e

mkdir -p /storage/logs

LOGFILE="/storage/logs/docker_$(date +%d_%m_%Y__%H_%M).log"
touch "$LOGFILE"

exec /compay >>"$LOGFILE" 2>&1
