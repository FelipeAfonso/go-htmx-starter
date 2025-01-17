#!/bin/bash
watchexec -r -e go,templ,css,json,js,ts,html -d 2500ms -E DEV=true -- curl -X POST -s -o /dev/null http://localhost:3000/reload/new &
watchexec -r -e templ,html -- tailwindcss -i ./pages/tailwind.css -o ./static/styles.css &
air
