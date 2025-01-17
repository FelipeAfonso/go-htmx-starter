#!/bin/bash
watchexec -r -e go,templ,css,json,js,ts,html -d 1500ms -E DEV=true -- curl -X POST -s -o /dev/null http://localhost:3000/reload/new &
air
