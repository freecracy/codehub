#!/bin/bash

set -u # 遇到不存在的变量时报错
set -e # 遇到错误时终止脚本执行
# set -x 脚本调试,输出命令和执行结果

{
    cd /tmp # 切换到缓存目录

    OS=""
    VERSION=$(curl -sSf https://api.github.com/repos/freecracy/codehub/tags | grep name | cut -c14- | tr -d '",' | sort -nr | head -1)

    if [ $(uname) == "Darwin" ]; then
        OS="Darwin"
    fi

    if [ $(uname) == "Linux" ]; then
        OS="Linux"
    fi

    if [ $OS == "" ]; then
        echo "not suppurted OS ..."
    fi

    DOWNLOAD_FILE=ip_${VERSION:1}_${OS}_"x86_64".tar.gz
    DOWNLOAD_URL="https://github.com/freecracy/codehub/releases/download/$VERSION/$DOWNLOAD_FILE"

    if command -v curl; then
        curl -SLO --progress $DOWNLOAD_URL # -L:解决 Unrecognized archive format
    elif command -v wget; then
        echo "wget"
    else
        echo "wget or curl not found ..."
    fi

    if [[ ! -e $DOWNLOAD_FILE ]]; then
        echo "download failed ..."
    fi

    tar -xvzf $DOWNLOAD_FILE
    BIN_FILE="ip"
    BIN_FILE2="ip2" # linux下有ip命令
    if [[ -e $BIN_FILE ]]; then
        mv $BIN_FILE /usr/local/bin/$BIN_FILE2
    fi

    if [[ -e /usr/local/bin/$BIN_FILE2 ]]; then

        cat <<-'EOF'

    ip2 --help
    
EOF
    fi
}
