#!/bin/sh
# Tech Relay - Install NetEase UU Game Booster on OpenWrt/GL.iNet
set -e

echo "[1/6] Updating package lists..."
opkg update

echo "[2/6] Installing kmod-tun..."
opkg install kmod-tun || true

echo "[3/6] Checking /dev/net/tun..."
if [ ! -e /dev/net/tun ]; then
  echo "ERROR: /dev/net/tun is missing. Reboot after installing kmod-tun, then run this script again."
  exit 1
fi

echo "[4/6] Downloading UU installer..."
cd /tmp
wget http://uu.gdl.netease.com/uuplugin-script/20231117102400/install.sh -O install.sh
chmod +x install.sh

echo "[5/6] Running installer for architecture: $(uname -m)"
/bin/sh install.sh openwrt $(uname -m)

echo "[6/6] Checking for UU files and processes..."
ps | grep -i uu | grep -v grep || true
ls /etc/init.d | grep -i uu || true

echo "Done. Connect your phone to this router and bind it in the UU Game Booster app."
