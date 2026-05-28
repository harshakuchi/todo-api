#!/bin/bash

set -e

# Load .env file
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

COMMAND=$1
NAME=$2

case "$COMMAND" in
  up)
    migrate -path migrations -database "$DATABASE_URL" up
    ;;

  down)
    COUNT=${NAME:-1}
    echo "Rolling back $COUNT migration(s). Continue? [y/N]"
    read CONFIRM

    if [ "$CONFIRM" = "y" ]; then
      migrate -path migrations -database "$DATABASE_URL" down "$COUNT"
    fi
    ;;

  create)
    migrate create -ext sql -dir migrations -seq "$NAME"
    ;;

  force)
    migrate -path migrations -database "$DATABASE_URL" force "$NAME"
    ;;

  *)
    echo "Usage:"
    echo "./scripts/migrate.sh up"
    echo "./scripts/migrate.sh down [count]"
    echo "./scripts/migrate.sh create [name]"
    echo "./scripts/migrate.sh force [version]"
    ;;
esac