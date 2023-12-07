#!/bin/bash

# commands
CMD=$1
FORCE=$2

# default commands
if [ "$CMD" = "" ]; then
  CMD="0"
fi
if [ "$FORCE" = "" ]; then
  FORCE="1"
fi

echo "------------------------------------------------"
echo "  Sintaxe this script:" 
echo "  > sh scripts/migration.sh CMD FORCE_VERSION" 
echo "  CMD is 0 = up, 1 = down, 2 = reverse, and 3 = force" 
echo "  to use force set FORCE_VERSION for version"
echo "------------------------------------------------"

go run cmd/migrate/migrate.go -c configs/app.yaml -cmd $CMD -force $FORCE