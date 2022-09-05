#!/bin/bash

docker buildx build --push --platform=linux/amd64,linux/arm64 . -t srcio/ash -t registry.cn-hangzhou.aliyuncs.com/srcio/ash -f Dockerfile
