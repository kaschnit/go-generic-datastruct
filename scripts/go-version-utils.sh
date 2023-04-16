#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROJ_DIR=$(dirname ${SCRIPT_DIR})

function comparableVersion {
    echo "$@" | tr -d " " | tr -d "go" | awk -F. '{ printf("%d%03d%03d%03d\n", $1,$2,$3,$4); }'
}

function goModVersion() {
    grep '^go \d\+[.]\d\+$' ${PROJ_DIR}/go.mod | tr -d " "
}

$*
