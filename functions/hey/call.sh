#!/bin/bash
curl -v -X POST \
  http://localhost:7070 \
  -H 'content-type: application/json; charset=utf-8' \
  -H 'my-token: I love Pandas' \
  -d '{"message": "Golang 💜 wasm", "author": "k33g"}'
  