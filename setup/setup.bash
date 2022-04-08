#!/usr/bin/env bash

: '
    Author: Aliasgar Khimani (NovusEdge)
    Project: github.com/ARaChn3/web-byte
    Copyright: MPL 2.0
    See the LICENSE file for more info.
    All Rights Reserved
'

PROJECT_DIR=$(dirname $(dirname $0))

if [ ! -d "$PROJECT_DIR/logs" ]; then
    printf "\033[33m[W]: Could not find \"logs/\" in $PROJECT_DIR\033[0m\n";
    printf "\033[1;36m[I]: Making the logging directory\033[0m\n\tLogging directory: $PROJECT_DIR/logs/\n\n";
    mkdir -p $PROJECT_DIR/logs;
fi

if ! type "go" > /dev/null; then
    printf "\033[31m[E]: Golang not detected!\n";
    printf "     Please install it from: https://go.dev/ \033[0m\n\n";
    exit 1;
fi

printf "\033[1;36m[I]: Getting golang dependencies...\033[0m\n";
cd $PROJECT_DIR; go get
printf "\033[1;32m[+]: Done!\033[0m\n";
