#!/bin/bash

./cf-whitelist

old=`sha256sum traefik.toml`
new=`sha256sum traefik.toml.new`

if [ "$old" != "$new" ]; then
	echo "Files are diff"
	cp traefik.toml.new traefik.toml
        ./start_traefik.sh
fi

rm traefik.toml.new
