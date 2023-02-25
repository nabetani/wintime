#!/bin/bash

set -eu

GOOS=windows GOARCH=amd64 go build -o bin/wintime.exe
ssh nabetani@miniregu.local 'if not exist wintime mkdir wintime'
scp bin/wintime.exe nabetani@miniregu.local:"./wintime"
(ssh nabetani@miniregu.local 'wintime\\wintime.exe') 2>&1 | ruby conv.rb
