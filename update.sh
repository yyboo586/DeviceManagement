#!/bin/bash

set -e

log() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [INFO] $1"
}

error() {
  echo "$(date '+%Y-%m-%d %H:%M:%S') [ERROR] $1" >&2
}

log "[1] 停止服务"
systemctl stop device_management

log "[2] 检测服务是否停止"
# 使用 systemctl is-active 检查服务是否真的停止（仅关注状态，不输出详细日志）
if systemctl is-active --quiet device_management; then
  error "服务未成功停止，无法继续操作"
  exit 1
else
  log "服务已成功停止"
fi

log "[3] 拷贝可执行文件到/usr/share/device_management"
/bin/cp -f ./bin/device_management /usr/share/device_management/
/bin/cp -f ./config.yaml /usr/share/device_management/

log "[4] 重启device_management服务"
systemctl restart device_management

log "[5] 查看device_management服务状态"
systemctl status device_management
