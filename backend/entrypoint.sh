#!/bin/sh

if [ -f "$POSTGRES_PASSWORD_FILE" ]; then
	export POSTGRES_PASSWORD=$(cat "$POSTGRES_PASSWORD_FILE")
fi

if [ -f "$POSTGRES_USER_FILE" ]; then
	export POSTGRES_USER=$(cat "$POSTGRES_USER_FILE")
fi

if [ -f "$POSTGRES_DB_FILE" ]; then
	export POSTGRES_DB=$(cat "$POSTGRES_DB_FILE")
fi

exec "$@"
