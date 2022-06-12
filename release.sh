#!/bin/bash
set -e

docker build --tag vgchl/terminal-ui:latest .
docker push vgchl/terminal-ui:latest
