#!/usr/bin/env bash

# Exit immediately if any command exits with a non-zero exit code.
set -e

# Generate Go files with all possible combinations of struct field types exported.
go run gen.go

max=63
for i in $(seq 0 $max)
do
    x=$(printf '%02d' "$i")

    go run ../../cmd/jay/main.go -p -d -v -o jay_"$x".go pkg_"$x".go
    echo "$x" "$(go test jay_"$x".go jay_"$x"_test.go pkg_"$x".go)"
done
