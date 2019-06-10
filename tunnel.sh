#!/bin/bash

# USAGE
# $ ./tunnel.sh start|stop

# starts an ssh tunnel to $HOST
# stops the tunnel on SIGINT if being run by a daemon

if test $# -eq 0; then
  echo missing arg
  exit 2
fi

CMD=$1
SOCKET="pets-mongo-socket"
HOST="pets-mongo"

function start_sock() {
  if test -S $SOCKET; then
    echo "already running"
  else
    echo "starting socket"
    ssh -f -N -M -S $SOCKET -L 4321:localhost:27017 $HOST
  fi
}

function stop_sock {
  if test -S $SOCKET; then
    echo "stopping socket"
    ssh -S $SOCKET -O exit $HOST > /dev/null 2>&1
  else
    echo "no socket to stop"
  fi
}

if test $CMD = "start"; then
  start_sock
elif test $CMD = "stop"; then
  stop_sock
fi

if ! test -t 1; then
  trap 'stop_sock; exit 0' SIGINT
  while true; do
    :
  done
fi
