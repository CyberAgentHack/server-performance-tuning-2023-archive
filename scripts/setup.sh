#!/bin/bash
# pprofのwebui用
sudo yum install graphviz -y

# redis-cli
sudo amazon-linux-extras install epel -y
    sudo yum install gcc jemalloc-devel openssl-devel tcl tcl-devel -y
    sudo wget http://download.redis.io/redis-stable.tar.gz
    sudo tar xvzf redis-stable.tar.gz
    cd redis-stable
    sudo make BUILD_TLS=yes
    sudo install -m 755 src/redis-cli /usr/local/bin/