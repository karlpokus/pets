#!/bin/bash

# USAGE
# $ ./tunnel.sh start|stop

if test $# -eq 0; then
  echo missing arg
  exit 2
fi

CMD=$1
SOCKET="pets-mongo-socket"
HOST="pets-mongo"

if test $CMD = "start"; then
  if test -S $SOCKET; then
    echo "already running"
  else
    echo "starting socket"
    ssh -f -N -M -S $SOCKET -L 4321:localhost:27017 $HOST
  fi
elif test $CMD = "stop"; then
  if test -S $SOCKET; then
    echo "stopping socket"
    ssh -S $SOCKET -O exit $HOST
  else
    echo "no socket to stop"
  fi
fi
