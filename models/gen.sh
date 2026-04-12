#/bin/bash

# 获取脚本所在的目录
SCRIPT_DIR="$(dirname "$0")"

# 切换到脚本所在的目录
cd "$SCRIPT_DIR"

coos.sh gen -m ./ -o ../ -r ../repository
