#!/bin/bash
# dockerインストール
yum update -y
yum install -y docker
service docker start
usermod -a -G docker ec2-user

# docker composeインストール
mkdir -p /usr/local/lib/docker/cli-plugins
curl -SL https://github.com/docker/compose/releases/download/v2.18.1/docker-compose-linux-x86_64 -o /usr/local/lib/docker/cli-plugins/docker-compose
chmod +x /usr/local/lib/docker/cli-plugins/docker-compose

# gitインストール
yum install -y git

dd if=/dev/zero of=/swapfile bs=1M count=2048
chmod 600 /swapfile
mkswap /swapfile
swapon /swapfile
echo "/swapfile swap swap defaults 0 0" > /etc/fstab

# ssh接続して以下やる
# ssh keyをgithubに登録
# git clone ~
# sudo docker compose up -d