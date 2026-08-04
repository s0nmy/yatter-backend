#!/bin/bash

if [ $# -lt 1 ]; then
  echo 出力ファイルは必須です.
elif [ $# -gt 1 ]; then
  echo 不正な引数です.
else
  docker compose exec web migrate create -ext sql -dir migrations -seq $1
fi
