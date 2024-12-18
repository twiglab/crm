#!/usr/bin/env bash

cd ../

find . -maxdepth 3 -mindepth 3 -type  d -wholename '*/cmd/*' | awk 'BEGIN{FS="/"} {print $2 "-" $4}'
