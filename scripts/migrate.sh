#!/bin/bash

if [[ $1 == "force" ]]; then
    if [[ $# != 2 ]]; then
        echo "force コマンドにはバージョン番号の指定が必要です"
        echo "使用方法: $0 force <version>"
        exit 1
    fi
    docker compose exec web migrate -path migrations -database "mysql://yatter:yatter@tcp(mysql:3306)/yatter" force "$2"
elif [[ $1 == "up" ]] || [[ $1 == "down" ]] || [[ $1 == "version" ]]; then
    if [[ $# != 1 ]]; then
        echo "不正な引数です"
        echo "使用方法: $0 (up|down|version)"
        exit 1
    fi
    docker compose exec web migrate -path migrations -database "mysql://yatter:yatter@tcp(mysql:3306)/yatter" "$1"
else
    echo "不正なコマンドです"
    echo "使用方法: $0 (up|down|version|force <version>)"
    exit 1
fi