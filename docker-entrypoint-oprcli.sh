#!/bin/sh
set -e

# this script will start relayproxy

PROGNAME=$(basename $0)

USER="eoracle"
GROUP="eoracle"
WORKDIR="/app/eoracle"
STARTUP="oprcli $@"

echo "$PROGNAME: Starting $STARTUP"
if [[ "$(id -u)" = '0' ]]; then
    cd ${WORKDIR}
    exec su-exec ${USER} ${STARTUP}
else
    # allow the container to be started with `--user`, in this case we cannot use su-exec
    cd ${WORKDIR}
    exec ${STARTUP}
fi