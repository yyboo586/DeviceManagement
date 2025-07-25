#!/bin/bash

set -e

error() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') [ERROR] $1" >&2
}

log() {
    echo "$(date '+%Y-%m-%d %H:%M:%S') [INFO] $1"
}

log "[1] 更新代码"
git fetch origin && git reset --hard origin/master

log "[2] 清理旧可执行文件..."
rm -rf ./bin/device_management

log "[3] 编译可执行文件..."
if ! go build -o ./bin/device_management main.go; then
  error "编译失败，请检查代码"
  exit 1
fi

log "[4] 检查编译结果..."
if [ ! -f "./bin/device_management" ]; then
  error "可执行文件不存在，编译可能失败"
  exit 1
fi