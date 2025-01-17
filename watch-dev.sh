#!/bin/bash
watchexec -r -e go,templ,css,json,js,ts,html -i _templ.go -d 1s -E DEV=true -- ./dev.sh
