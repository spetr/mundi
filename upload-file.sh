#!/bin/sh

curl -F "file=@$1" http://127.0.0.1:25794/api/v1/auto-convert
