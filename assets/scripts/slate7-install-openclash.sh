#!/bin/sh
set -e

opkg update

opkg install bash iptables dnsmasq-full curl ca-bundle ipset ip-full \
iptables-mod-tproxy iptables-mod-extra ruby ruby-yaml kmod-tun \
kmod-inet-diag unzip luci-compat luci luci-base

cd /tmp

curl -L --retry 2 https://api.github.com/repos/vernesong/OpenClash/releases/latest \
  -o /tmp/openclash_version

download_url=$(cat /tmp/openclash_version | jsonfilter -e '@.assets[*].browser_download_url' | grep '\.ipk$')

curl -L --retry 2 "$download_url" -o /tmp/openclash.ipk

opkg install /tmp/openclash.ipk || true

echo "OpenClash install attempted. Check LuCI Services > OpenClash and verify Meta core."
