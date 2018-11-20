#!/usr/bin/env bash

set -o nounset

run_cmd="go test -v -tags=unit -race -coverprofile=profile.out -covermode=atomic"

coveragemode="mode: atomic"
echo $coveragemode > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    # go test -v -tags=unit -race -coverprofile=profile.out -covermode=atomic $d
    $run_cmd $d
    if [ -f profile.out ]; then
        grep -h -v "^${coveragemode}" profile.out >> coverage.txt
        rm profile.out
    fi
done

go tool cover -html=coverage.txt -o=coverage.html
