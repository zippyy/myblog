#!/bin/sh
# Tech Relay - UU Game Booster health check for GL.iNet/OpenWrt

echo "=== OpenWrt release ==="
cat /etc/openwrt_release 2>/dev/null | sed -n '1,8p'
echo

echo "=== Architecture ==="
uname -m
echo

echo "=== TUN device ==="
ls -l /dev/net/tun 2>/dev/null || echo "Missing /dev/net/tun"
echo

echo "=== LAN bridge ==="
ip addr show br-lan 2>/dev/null || echo "Missing br-lan"
echo

echo "=== UU processes ==="
ps | grep -i uu | grep -v grep || echo "No UU process found"
echo

echo "=== UU init scripts ==="
ls /etc/init.d | grep -i uu || echo "No UU init script found"
echo

echo "=== DNS test ==="
nslookup uu.163.com 2>/dev/null || nslookup openwrt.org 2>/dev/null || echo "DNS lookup failed"
