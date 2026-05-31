---
title: "GL.iNet Flint 3 + Slate 7 Pro: Self-Hosted VLESS REALITY with OpenClash"
slug: "glinet-flint3-slate7-vless-reality-openclash"
date: 2026-05-30T12:00:00-06:00
lastmod: 2026-05-30T12:00:00-06:00
draft: false
description: "A complete guide to installing OpenClash on a GL.iNet Slate 7 Pro, running a self-hosted VLESS REALITY server on a GL.iNet Flint 3, and tunneling travel-router clients through home."
summary: "Install OpenClash on a GL.iNet Slate 7 Pro and connect it to a self-hosted Xray VLESS REALITY server on a GL.iNet Flint 3."
categories:
  - Networking
  - Homelab
  - VPN
tags:
  - GL.iNet
  - Flint 3
  - Slate 7 Pro
  - OpenClash
  - Mihomo
  - Xray
  - VLESS
  - REALITY
  - Travel Router
toc: true
---

> **Affiliate disclosure:** This post may contain affiliate links. If you purchase through those links, Tech Relay may earn a commission at no additional cost to you.

## Overview

This guide walks through a complete end-to-end setup:

1. Install OpenClash on a **GL.iNet Slate 7 Pro**.
2. Install and configure **Xray** on a **GL.iNet Flint 3**.
3. Build a **VLESS REALITY** tunnel using **XTLS Vision**.
4. Configure OpenClash/Mihomo on the Slate 7 Pro.
5. Route all Slate-connected client devices through the Flint 3 at home.

The goal is to turn the Slate 7 Pro into a travel router that tunnels connected devices back through the Flint 3 without using a traditional WireGuard/OpenVPN-style tunnel.

The finished path looks like this:

```text
Phone / Laptop / Tablet
        ↓
GL.iNet Slate 7 Pro
        ↓
OpenClash / Mihomo Meta
        ↓
VLESS + REALITY + XTLS Vision
        ↓
GL.iNet Flint 3
        ↓
Home Internet Connection
```

{{< figure src="/images/posts/glinet-flint3-slate7-vless-reality-openclash/architecture.png" alt="Network architecture for Flint 3 and Slate 7 Pro VLESS REALITY setup" caption="Client devices connect to the Slate 7 Pro, which tunnels traffic through OpenClash to the Flint 3 running Xray." >}}

## Hardware Used

### GL.iNet Flint 3 / GL-BE9300

The Flint 3 is the home-side endpoint. It runs the Xray server and accepts inbound VLESS REALITY connections.

**Affiliate link placeholder:**  
`https://example.com/flint3-affiliate-link`

### GL.iNet Slate 7 Pro / GL-BE10000

The Slate 7 Pro is the travel-side router. It runs OpenClash and transparently tunnels connected client devices back through the Flint 3.

**Affiliate link placeholder:**  
`https://example.com/slate7pro-affiliate-link`

### Optional Accessories

- Router UPS: `https://example.com/router-ups-affiliate-link`
- USB-C PD battery: `https://example.com/usb-c-pd-power-bank-affiliate-link`
- Portable 5G modem/hotspot: `https://example.com/5g-modem-affiliate-link`

## VLESS REALITY vs WireGuard vs Tailscale

| Feature | VLESS REALITY | WireGuard | Tailscale |
|---|---:|---:|---:|
| Self-hostable | Yes | Yes | Partially |
| Requires TLS certs | No | No | No |
| Looks like HTTPS | Yes | No | No |
| Easy client config | Medium | Easy | Easy |
| Router-level travel setup | Yes | Yes | Yes |
| Built-in mesh | No | No | Yes |

Use WireGuard when you want simple and fast VPN tunneling.

Use Tailscale when you want easy mesh networking.

Use VLESS REALITY when you want a self-hosted transport that more closely resembles normal HTTPS traffic and does not require TLS certificate management.

## Requirements

You need:

- Flint 3 with SSH access
- Slate 7 Pro with SSH access
- DDNS hostname pointing to your home connection
- TCP port forward to the Flint 3
- Xray Core for ARM64
- OpenClash installed on the Slate 7 Pro
- Mihomo Meta core for OpenClash

This guide uses sanitized placeholders:

```text
YOUR_DDNS_HOSTNAME
YOUR_UUID
YOUR_PRIVATE_KEY
YOUR_PUBLIC_KEY
YOUR_SHORT_ID
YOUR_PROXY_PASSWORD
YOUR_OPENCLASH_SECRET
```

Do not publish real keys, UUIDs, short IDs, hostnames, IPs, or passwords.

---

# Part 1: Install OpenClash on the Slate 7 Pro

The Slate 7 Pro does not ship with OpenClash installed by default. Install it first.

## Check the Slate 7 Pro Architecture

SSH into the Slate 7 Pro:

```bash
ssh root@192.168.8.1
```

Check the platform:

```bash
cat /etc/openwrt_release
```

A Slate 7 Pro may report something like:

```text
DISTRIB_TARGET='mediatek/mt7987'
DISTRIB_ARCH='aarch64_cortex-a53'
```

That means you want the **IPK** OpenClash package and the **ARM64/AArch64** Mihomo core.

## Update Package Lists

```bash
opkg update
```

## Install OpenClash Dependencies

For the Slate 7 Pro firmware used in this setup, the working OpenClash dependency path was the **iptables/IPK** install path, not the nftables one.

Use:

```bash
opkg install bash iptables dnsmasq-full curl ca-bundle ipset ip-full \
iptables-mod-tproxy iptables-mod-extra ruby ruby-yaml kmod-tun \
kmod-inet-diag unzip luci-compat luci luci-base
```

### Why not kmod-nft-tproxy?

On this GL.iNet OpenWrt 21.02-based firmware, `kmod-nft-tproxy` was not available in the GL.iNet feeds. The router already had iptables TProxy support:

```bash
opkg list-installed | grep tproxy
```

Expected:

```text
iptables-mod-tproxy
kmod-ipt-tproxy
```

So the iptables OpenClash install path is the safer path on this firmware.

## Download the Latest OpenClash IPK

```bash
cd /tmp

curl -L --retry 2 https://api.github.com/repos/vernesong/OpenClash/releases/latest \
  -o /tmp/openclash_version

download_url=$(cat /tmp/openclash_version | jsonfilter -e '@.assets[*].browser_download_url' | grep '\.ipk$')

curl -L --retry 2 "$download_url" -o /tmp/openclash.ipk
```

Verify it is not a tiny failed download:

```bash
ls -lh /tmp/openclash.ipk
```

If the file is only a few bytes, the URL failed.

Install it:

```bash
opkg install /tmp/openclash.ipk
```

## Verify OpenClash Files

```bash
ls -lah /etc/openclash/
```

Check the core directory:

```bash
ls -lah /etc/openclash/core/
```

You should eventually have:

```text
clash_meta
```

If the Meta core is missing, install or download it through the OpenClash LuCI interface:

```text
Services → OpenClash → Version Update
```

Select or update the **Mihomo Meta** core.

## Verify the Mihomo Meta Core

```bash
/etc/openclash/core/clash_meta -v
```

Expected output should identify Mihomo Meta for Linux ARM64.

Example:

```text
Mihomo Meta ... linux arm64
```

## Basic OpenClash Settings

In LuCI:

```text
Services → OpenClash
```

Recommended initial settings:

```text
Mode: Rule
DNS Mode: redir-host
Core: Meta / Mihomo
```

Avoid Fake-IP while first testing. Fake-IP may return `198.18.x.x` addresses, which is normal for that mode but makes troubleshooting harder.

From CLI, force redir-host:

```bash
uci set openclash.config.en_mode='redir-host'
uci set openclash.config.operation_mode='redir-host'
uci commit openclash
```

---

# Part 2: Install Xray on the Flint 3

SSH into the Flint 3.

```bash
mkdir -p /opt/xray
cd /opt/xray
```

Download Xray:

```bash
wget -O xray.zip https://github.com/XTLS/Xray-core/releases/latest/download/Xray-linux-arm64-v8a.zip
unzip -o xray.zip
chmod +x xray
```

Generate REALITY keys:

```bash
./xray x25519
```

Example placeholder output:

```text
PrivateKey: YOUR_PRIVATE_KEY
PublicKey:  YOUR_PUBLIC_KEY
```

Generate a UUID:

```bash
./xray uuid
```

Example:

```text
YOUR_UUID
```

Generate a short ID:

```bash
head -c 8 /dev/urandom | xxd -p
```

Example:

```text
YOUR_SHORT_ID
```

---

# Part 3: Create the Xray Server Config on the Flint 3

Create `/opt/xray/config.json`:

```bash
cat >/opt/xray/config.json <<'EOF'
{
  "log": {
    "loglevel": "warning"
  },
  "inbounds": [
    {
      "listen": "0.0.0.0",
      "port": 9443,
      "protocol": "vless",
      "settings": {
        "clients": [
          {
            "id": "YOUR_UUID",
            "flow": "xtls-rprx-vision"
          }
        ],
        "decryption": "none"
      },
      "streamSettings": {
        "network": "tcp",
        "security": "reality",
        "realitySettings": {
          "show": false,
          "dest": "www.cloudflare.com:443",
          "serverNames": [
            "www.cloudflare.com"
          ],
          "privateKey": "YOUR_PRIVATE_KEY",
          "shortIds": [
            "YOUR_SHORT_ID"
          ]
        }
      }
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom"
    }
  ]
}
EOF
```

Test the config:

```bash
/opt/xray/xray run -test -config /opt/xray/config.json
```

Expected:

```text
Configuration OK.
```

---

# Part 4: Open the Flint 3 Firewall

```bash
uci -q delete firewall.xray

uci set firewall.xray='rule'
uci set firewall.xray.name='Xray-9443'
uci set firewall.xray.src='wan'
uci set firewall.xray.proto='tcp'
uci set firewall.xray.dest_port='9443'
uci set firewall.xray.target='ACCEPT'

uci commit firewall
/etc/init.d/firewall restart
```

If the Flint is behind an ISP gateway, forward:

```text
TCP 9443 → Flint 3 LAN IP → TCP 9443
```

---

# Part 5: Create an Xray Startup Service

Create `/etc/init.d/xray`:

```bash
cat >/etc/init.d/xray <<'EOF'
#!/bin/sh /etc/rc.common

START=99
USE_PROCD=1

start_service() {
    procd_open_instance
    procd_set_param command /opt/xray/xray run -config /opt/xray/config.json
    procd_set_param env XRAY_LOCATION_ASSET=/opt/xray
    procd_set_param respawn
    procd_set_param stdout 1
    procd_set_param stderr 1
    procd_close_instance
}
EOF
```

Enable and start:

```bash
chmod +x /etc/init.d/xray
/etc/init.d/xray enable
/etc/init.d/xray start
```

Verify:

```bash
netstat -lntp | grep 9443
```

Expected:

```text
:::9443 LISTEN
```

---

# Part 6: Create the OpenClash Config on the Slate 7 Pro

On the Slate 7 Pro, create:

```text
/etc/openclash/config/flint3-reality.yaml
```

Use:

```yaml
mixed-port: 7890
allow-lan: true
mode: rule
log-level: info
ipv6: false
tcp-concurrent: true

dns:
  enable: true
  listen: 0.0.0.0:7874
  ipv6: false
  enhanced-mode: redir-host
  nameserver:
    - https://1.1.1.1/dns-query
    - https://dns.google/dns-query
  fallback:
    - https://dns.cloudflare.com/dns-query
    - https://dns.google/dns-query

proxies:
  - name: Flint3-Reality
    type: vless
    server: YOUR_DDNS_HOSTNAME
    port: 9443
    uuid: YOUR_UUID
    network: tcp
    tls: true
    udp: true
    flow: xtls-rprx-vision
    servername: www.cloudflare.com
    client-fingerprint: chrome
    reality-opts:
      public-key: YOUR_PUBLIC_KEY
      short-id: YOUR_SHORT_ID

proxy-groups:
  - name: Proxy
    type: select
    proxies:
      - Flint3-Reality

rules:
  - MATCH,Proxy
```

In OpenClash LuCI:

```text
Config Manage → Select flint3-reality.yaml
```

Start OpenClash.

---

# Part 7: Fix the OpenClash SAFE_PATHS Issue

On this Slate 7 Pro setup, OpenClash generated a runtime config at:

```text
/etc/openclash/flint3-reality.yaml
```

It included:

```yaml
external-ui: "/usr/share/openclash/ui"
external-ui-name: metacubexd
external-ui-url: https://codeload.github.com/MetaCubeX/metacubexd/zip/refs/heads/gh-pages
```

Mihomo refused to start with:

```text
Parse config error: path is not subpath of home directory or SAFE_PATHS: /usr/share/openclash/ui
allowed paths: [/etc/openclash]
```

Remove the bad UI lines:

```bash
sed -i '/external-ui:/d' /etc/openclash/flint3-reality.yaml
sed -i '/external-ui-name:/d' /etc/openclash/flint3-reality.yaml
sed -i '/external-ui-url:/d' /etc/openclash/flint3-reality.yaml
```

Restart OpenClash:

```bash
/etc/init.d/openclash restart
```

Verify Mihomo is running:

```bash
ps w | grep clash
```

Expected:

```text
/etc/openclash/clash -d /etc/openclash -f /etc/openclash/flint3-reality.yaml
```

---

# Part 8: Verify OpenClash Ports

```bash
netstat -lnt | grep 789
```

Expected listeners include:

```text
7890
7891
7892
7893
7895
```

Check DNS:

```bash
netstat -lnt | grep 7874
```

---

# Part 9: Test the Tunnel

If HTTP proxy authentication is enabled, test with:

```bash
curl -x http://Clash:YOUR_PROXY_PASSWORD@127.0.0.1:7890 https://ipv4.icanhazip.com
```

Expected:

```text
YOUR_FLINT_PUBLIC_IP
```

This should be the Flint 3 public IP, not the Slate 7 Pro WAN IP.

Check the OpenClash API:

```bash
curl -s http://127.0.0.1:9090/proxies \
  -H "Authorization: Bearer YOUR_OPENCLASH_SECRET"
```

Confirm:

```json
"Proxy": {
  "now": "Flint3-Reality"
}
```

Check active connections:

```bash
curl -s http://127.0.0.1:9090/connections \
  -H "Authorization: Bearer YOUR_OPENCLASH_SECRET"
```

Expected:

```json
"chains": ["Flint3-Reality", "Proxy"]
```

---

# Part 10: Verify Phone Clients

Connect a phone or laptop to the Slate 7 Pro Wi-Fi.

Visit:

```text
https://ipv4.icanhazip.com
```

Expected:

```text
YOUR_FLINT_PUBLIC_IP
```

If it shows the Slate 7 Pro WAN IP instead, OpenClash is running but client traffic is not being transparently redirected.

Check:

```bash
tail -100 /tmp/openclash.log
curl -s http://127.0.0.1:9090/connections \
  -H "Authorization: Bearer YOUR_OPENCLASH_SECRET"
```

---

# Troubleshooting

## OpenClash IPK Download Is Only a Few Bytes

If the IPK is only a few bytes, the GitHub URL was wrong or truncated.

Use the GitHub API method from this guide instead of manually copying a release URL.

## kmod-nft-tproxy Cannot Be Installed

Use the iptables OpenClash dependency path instead.

Check:

```bash
opkg list-installed | grep tproxy
```

If `iptables-mod-tproxy` and `kmod-ipt-tproxy` exist, the iptables path is appropriate.

## DNS Returns 198.18.x.x

That means Fake-IP mode is active.

Use redir-host:

```bash
uci set openclash.config.en_mode='redir-host'
uci set openclash.config.operation_mode='redir-host'
uci commit openclash
/etc/init.d/openclash restart
```

## OpenClash Says Already Start but Nothing Works

Check if Mihomo is actually running:

```bash
pidof clash
pidof clash_meta
ps w | grep clash
```

Check logs:

```bash
tail -100 /tmp/openclash.log
```

## SAFE_PATHS Error

Remove:

```yaml
external-ui:
external-ui-name:
external-ui-url:
```

from:

```text
/etc/openclash/flint3-reality.yaml
```

## 407 Proxy Error

If this fails:

```bash
curl -x http://127.0.0.1:7890 https://ipv4.icanhazip.com
```

with:

```text
CONNECT tunnel failed, response 407
```

use credentials:

```bash
curl -x http://Clash:YOUR_PROXY_PASSWORD@127.0.0.1:7890 https://ipv4.icanhazip.com
```

## Xray Works in Foreground but Not as a Service

Make sure `/etc/init.d/xray` includes:

```bash
procd_set_param env XRAY_LOCATION_ASSET=/opt/xray
```

## Port Already in Use

Check:

```bash
netstat -lntp | grep 9443
```

If another service owns the port, change both the Xray server and OpenClash client to a different port.

---

# Final Validation Checklist

A complete working deployment should show:

- OpenClash installed on the Slate 7 Pro
- Mihomo Meta core present and executable
- Xray running on the Flint 3
- TCP `9443` listening on the Flint 3
- TCP `9443` reachable from the Slate 7 Pro
- OpenClash proxy group set to `Flint3-Reality`
- `curl` through the local proxy returning the Flint public IP
- Phone clients connected to the Slate 7 Pro showing the Flint public IP
- Xray logs showing accepted TCP sessions

## Conclusion

This setup turns a GL.iNet Flint 3 into a self-hosted VLESS REALITY endpoint and a GL.iNet Slate 7 Pro into a travel router that tunnels connected clients back through home.

The complete working chain is:

```text
Client Device → Slate 7 Pro → OpenClash/Mihomo → VLESS REALITY → Flint 3/Xray → Internet
```

The most important troubleshooting lesson was that the tunnel itself may be working even when OpenClash is not. In this case, the key failure was OpenClash generating a Mihomo config with an invalid `external-ui` path that violated Mihomo's SAFE_PATHS restriction.
