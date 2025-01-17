#!/bin/bash
tailwindcss -i ./pages/tailwind.css -o ./static/styles.css 
templ generate 
(sleep 2 && curl -X POST -s -o /dev/null http://localhost:3000/reload/new)&
go run .
