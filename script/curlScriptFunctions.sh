#!/bin/bash
function  newLink {
  curl -X POST -H "Content-Type: application/json" -d \
  '{
   "resourcePath":"myPath",
   "redirectionPath":"www.google.com"
  }
  ' \
  http://localhost:8080/api/newlink
}

function  redirect {
  curl http://localhost:8080/myPath
}
