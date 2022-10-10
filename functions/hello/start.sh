#!/bin/bash
MESSAGE="👋 hello world 🌍" capsule \
  -wasm=hello.wasm \
  -mode=http \
  -httpPort=7071

# gp preview $(gp url 7071)