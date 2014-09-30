#!/bin/bash

function NewLink {


  curl -X POST -H "Content-Type: application/json" -d \
  '{
   "resourcePath":"myPath",
   "redirectionPath":"http://www.google.com"
  }
  ' \
  http://localhost:8080/api/newlink
}

function Redirect {
  curl http://localhost:8080/myPath
}
