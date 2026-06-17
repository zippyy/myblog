---
title: "Running AstroWarp, Tailscale, ZeroTier, and WireGuard Together on a GL.iNet Flint 3"
slug: "astrowarp-tailscale-zerotier-wireguard-flint3"
date: 2026-06-17T12:00:00-06:00
lastmod: 2026-06-17T12:00:00-06:00
draft: false
description: "AstroWarp, Tailscale, ZeroTier, and WireGuard can run together on a GL.iNet Flint 3. The catch is that GL.iNet's management layer may disable Tailscale and ZeroTier when AstroWarp is enabled. Here is the fix."
summary: "A real-world Flint 3 Mesh Alpha firmware test showing how AstroWarp, Tailscale, ZeroTier, and WireGuard can coexist, why Tailscale and ZeroTier get disabled, and the exact recovery script to restore them."
categories:
  - Networking
  - Homelab
  - VPN
tags:
  - GL.iNet
  - Flint 3
  - AstroWarp
  - Tailscale
  - ZeroTier
  - WireGuard
  - OpenWrt
  - Mesh Alpha
  - Remote Access
toc: true
codeLineNumbers: false
codeMaxLines: 25
usePageBundles: true
featureImage: "/images/posts/astrowarp-tailscale-zerotier-wireguard-flint3/hero.svg"
featureImageAlt: "GL.iNet Flint 3 running AstroWarp, Tailscale, ZeroTier, and WireGuard together"
featureImageCap: "AstroWarp, Tailscale, ZeroTier, and WireGuard running together on a GL.iNet Flint 3."
shareImage: "/images/posts/astrowarp-tailscale-zerotier-wireguard-flint3/og-image.svg"
thumbnail: "/images/posts/astrowarp-tailscale-zerotier-wireguard-flint3/hero.svg"
year: "2026"
month: "2026-06"
---

> **Affiliate disclosure:** This post may contain affiliate links. If you purchase through those links, Tech Relay may earn a commission at no additional cost to you.

## Overview

The GL.iNet Flint 3 is one of those routers that starts out looking like a normal Wi-Fi box and then slowly turns into a homelab appliance the deeper you go.

On my Flint 3, I already had:

- **WireGuard Server**
- **Tailscale**
- **ZeroTier**
- VLANs
- Mesh Alpha firmware
- Multiple LANs
- Remote subnet access

Then I started testing **AstroWarp**.

The GL.iNet UI makes it look like AstroWarp does not want to coexist with other overlay/VPN services. When AstroWarp comes up, Tailscale and ZeroTier may suddenly disappear. Tailscale reports that `tailscaled` is not running. ZeroTier reports that `zerotier-one.port` is missing.

At first glance, that looks like a hard compatibility problem.

After digging into it from SSH, that is not what was happening.

The Linux networking stack could run everything together just fine.

The actual problem was this:

**The GL.iNet management layer was flipping the Tailscale and ZeroTier UCI settings to disabled.**

Once those settings were restored and the services were restarted, AstroWarp, Tailscale, ZeroTier, and WireGuard all ran together.

This post documents the testing, the commands, the symptoms, and the recovery script I now keep on the router.

---

## Test Environment

This was tested on:

```text
Router:   GL.iNet Flint 3 / GL-BE9300
Firmware: Mesh Alpha firmware
Base:     OpenWrt 23.05 snapshot-style firmware
Shell:    root SSH access
```

The relevant services were:

| Service | Interface / Role |
|---|---|
| AstroWarp | `mptun0` |
| Tailscale | `tailscale0` |
| ZeroTier | `zt...` interface |
| WireGuard Server | `wgserver` |
| Main LAN | `br-lan`, `192.168.80.0/24` |

The test networks used in this write-up were:

```text
Local LAN:                 192.168.80.0/24
Remote Tailscale subnet:   192.168.42.0/24
ZeroTier network:          10.121.15.0/24
AstroWarp tunnel:          10.3.128.0/24 via mptun0
```

Your networks will almost certainly be different. Adjust the IPs accordingly.

---

## The Breakage

After enabling AstroWarp, Tailscale stopped responding:

```sh
tailscale status
```

Output:

```text
failed to connect to local tailscaled; it doesn't appear to be running
```

ZeroTier also stopped responding:

```sh
zerotier-cli info
```

Output:

```text
zerotier-cli: missing port and zerotier-one.port not found in /var/lib/zerotier-one
```

That usually means the daemons are not running.

The next step was to check UCI.

---

## Checking Tailscale UCI

```sh
uci show tailscale
```

Broken state:

```text
tailscale.settings=settings
tailscale.settings.log_stderr='1'
tailscale.settings.log_stdout='1'
tailscale.settings.port='41641'
tailscale.settings.state_file='/etc/tailscale/tailscaled.state'
tailscale.settings.enabled='0'
tailscale.settings.lan_enabled='0'
tailscale.settings.wan_enabled='0'
tailscale.settings.ssh_enabled='1'
tailscale.settings.masq='0'
tailscale.settings.run_exit_node='0'
```

The important lines:

```text
tailscale.settings.enabled='0'
tailscale.settings.lan_enabled='0'
tailscale.settings.wan_enabled='0'
tailscale.settings.masq='0'
```

So Tailscale was not broken at the binary level. It had been disabled in GL.iNet's UCI config.

---

## Checking ZeroTier UCI

```sh
uci show zerotier
```

Broken state:

```text
zerotier.gl=zerotier
zerotier.gl.enabled='0'
zerotier.gl.nat='1'
zerotier.gl.wan_nat='0'
zerotier.gl.join='69f5545babee4bbf'
```

The important line:

```text
zerotier.gl.enabled='0'
```

Again, not a mysterious tunnel failure. ZeroTier was simply disabled.

---

## Verifying AstroWarp

AstroWarp created an `mptun0` interface:

```sh
ip -br addr | grep mptun
```

Example output:

```text
mptun0           UNKNOWN        10.3.128.2/24
```

Interface stats:

```sh
ip -s link show mptun0
```

Example output:

```text
126: mptun0: <POINTOPOINT,NOARP,UP,LOWER_UP> mtu 1420 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/none
    RX:  bytes packets errors dropped  missed   mcast
          1196      13     72       0       0       0
    TX:  bytes packets errors dropped carrier collsns
         16286     204    690       0       0       0
```

AstroWarp's routing table:

```sh
ip route show table mptcp_mptun0
```

Output:

```text
default dev mptun0 scope link
```

Policy rules:

```sh
ip rule show
```

Relevant rule:

```text
30:     from all fwmark 0x2/0x2 lookup mptcp_mptun0
```

That tells us AstroWarp is policy-routed. Traffic marked with `0x2` uses table `mptcp_mptun0`, which defaults through `mptun0`.

AstroWarp was not installing simple remote subnet routes like Tailscale or ZeroTier. It was using policy routing.

---

## Verifying Tailscale Routing

Once Tailscale was re-enabled, the router itself could reach the remote Tailscale subnet:

```sh
ping 192.168.42.1
```

Example output:

```text
PING 192.168.42.1 (192.168.42.1): 56 data bytes
64 bytes from 192.168.42.1: seq=0 ttl=64 time=19.552 ms
64 bytes from 192.168.42.1: seq=1 ttl=64 time=22.230 ms
```

Route lookup:

```sh
ip route get 192.168.42.1
```

Output:

```text
192.168.42.1 dev tailscale0 table 52 src 100.97.167.103 uid 0
    cache
```

Tailscale route table:

```sh
ip route show table 52
```

Relevant output:

```text
192.168.42.0/24 dev tailscale0
```

That proved Tailscale routing itself was good.

---

## The LAN Client Symptom

The router could ping the remote Tailscale subnet, but a Windows PC behind the Flint 3 could not.

From Windows:

```powershell
ping 192.168.42.1
```

Broken output:

```text
Pinging 192.168.42.1 with 32 bytes of data:
Reply from 192.168.80.1: Destination port unreachable.
Reply from 192.168.80.1: Destination port unreachable.
Reply from 192.168.80.1: Destination port unreachable.
```

The reply came from:

```text
192.168.80.1
```

That is the Flint 3 LAN gateway.

So the PC was sending packets to the router, and the router was rejecting them.

The missing piece was:

```text
tailscale.settings.masq='0'
```

Once Tailscale masquerading was restored and the service restarted, LAN clients could reach the remote subnet.

Working output from Windows:

```text
Pinging 192.168.42.1 with 32 bytes of data:
Reply from 192.168.42.1: bytes=32 time=19ms TTL=63
Reply from 192.168.42.1: bytes=32 time=23ms TTL=63
Reply from 192.168.42.1: bytes=32 time=20ms TTL=63
```

---

## Verifying ZeroTier

After restoring ZeroTier:

```sh
zerotier-cli info
```

Output:

```text
200 info 73da18a38c 1.14.1 ONLINE
```

Router test:

```sh
ping 10.121.15.226
```

Example output:

```text
64 bytes from 10.121.15.226: seq=0 ttl=64 time=134.733 ms
64 bytes from 10.121.15.226: seq=1 ttl=64 time=17.723 ms
```

LAN client test from Windows:

```powershell
ping 10.121.15.226
```

Working output:

```text
Pinging 10.121.15.226 with 32 bytes of data:
Reply from 10.121.15.226: bytes=32 time=23ms TTL=64
Reply from 10.121.15.226: bytes=32 time=28ms TTL=64
Reply from 10.121.15.226: bytes=32 time=24ms TTL=64
Reply from 10.121.15.226: bytes=32 time=21ms TTL=64
```

ZeroTier did not need much help. Once `zerotier.gl.enabled` was put back to `1`, it came back normally.

---

## The Minimal Fix

This is the final command set that restored both Tailscale and ZeroTier:

```sh
uci set tailscale.settings.enabled='1'
uci set tailscale.settings.lan_enabled='1'
uci set tailscale.settings.wan_enabled='1'
uci set tailscale.settings.masq='1'
uci set zerotier.gl.enabled='1'
uci commit tailscale
uci commit zerotier
/etc/init.d/tailscale restart
/etc/init.d/zerotier restart
```

That alone was enough to restore:

- Tailscale daemon
- Tailscale LAN subnet access
- ZeroTier daemon
- ZeroTier LAN client access

---

## Recovery Script

I keep this script on the Flint 3 as `/root/fix-astrowarp-overlays.sh`.

```sh
cat >/root/fix-astrowarp-overlays.sh <<'EOF'
#!/bin/sh

echo "[+] Restoring Tailscale settings..."
uci set tailscale.settings.enabled='1'
uci set tailscale.settings.lan_enabled='1'
uci set tailscale.settings.wan_enabled='1'
uci set tailscale.settings.masq='1'

echo "[+] Restoring ZeroTier settings..."
uci set zerotier.gl.enabled='1'

echo "[+] Committing UCI changes..."
uci commit tailscale
uci commit zerotier

echo "[+] Restarting Tailscale..."
/etc/init.d/tailscale restart

echo "[+] Restarting ZeroTier..."
/etc/init.d/zerotier restart

sleep 5

echo
echo "=== Tailscale ==="
tailscale status 2>/dev/null || echo "Tailscale is not ready yet."

echo
echo "=== ZeroTier ==="
zerotier-cli info 2>/dev/null || echo "ZeroTier is not ready yet."

echo
echo "=== Interfaces ==="
ip -br addr | grep -E 'tailscale0|zt|mptun0|wgserver' || true

echo
echo "[+] Done."
EOF

chmod +x /root/fix-astrowarp-overlays.sh
```

Run it whenever AstroWarp disables the overlays:

```sh
/root/fix-astrowarp-overlays.sh
```

---

## Optional Watchdog

If AstroWarp repeatedly disables Tailscale and ZeroTier, you can run a simple watchdog.

Create:

```sh
cat >/root/watch-astrowarp-overlays.sh <<'EOF'
#!/bin/sh

while true; do
    fixed=0

    if [ "$(uci -q get tailscale.settings.enabled)" != "1" ]; then
        logger -t astrowarp-overlay-watch "Tailscale disabled; restoring"
        uci set tailscale.settings.enabled='1'
        uci set tailscale.settings.lan_enabled='1'
        uci set tailscale.settings.wan_enabled='1'
        uci set tailscale.settings.masq='1'
        fixed=1
    fi

    if [ "$(uci -q get zerotier.gl.enabled)" != "1" ]; then
        logger -t astrowarp-overlay-watch "ZeroTier disabled; restoring"
        uci set zerotier.gl.enabled='1'
        fixed=1
    fi

    if [ "$fixed" = "1" ]; then
        uci commit tailscale
        uci commit zerotier
        /etc/init.d/tailscale restart
        /etc/init.d/zerotier restart
    fi

    sleep 30
done
EOF

chmod +x /root/watch-astrowarp-overlays.sh
```

Run manually:

```sh
/root/watch-astrowarp-overlays.sh &
```

Or add it to `/etc/rc.local` before `exit 0`:

```sh
/root/watch-astrowarp-overlays.sh &
```

I prefer the manual recovery script unless AstroWarp keeps flipping the settings repeatedly.

---

## Diagnostic Command Reference

### Check AstroWarp

```sh
ip -br addr | grep mptun
ip -s link show mptun0
ip route show table mptcp_mptun0
ip rule show
```

### Check Tailscale

```sh
uci show tailscale
tailscale status
ip route show table 52
ip route get 192.168.42.1
```

### Check ZeroTier

```sh
uci show zerotier
zerotier-cli info
zerotier-cli listnetworks
ip -br addr | grep zt
```

### Check WireGuard Server

```sh
ip -br addr | grep wg
wg show
```

### Check Everything at Once

```sh
ip -br addr | grep -E 'mptun0|tailscale0|zt|wgserver'
ps w | grep -E 'mptun|tailscaled|zerotier|wireguard'
```

---

## What I Think Is Happening

Based on testing, AstroWarp itself does not appear to be technically incompatible with Tailscale or ZeroTier.

The `mptun0` interface works.

Tailscale works once re-enabled.

ZeroTier works once re-enabled.

WireGuard keeps working.

The thing that breaks Tailscale and ZeroTier is that GL.iNet's management layer changes their UCI settings when AstroWarp is enabled or refreshed.

That means this is less of a kernel/networking problem and more of a UI/service-management problem.

---

## Important Warning About Routes

Running multiple overlay networks is powerful, but you need to avoid overlapping routes.

Bad idea:

```text
192.168.1.0/24 via Tailscale
192.168.1.0/24 via ZeroTier
192.168.1.0/24 via WireGuard
```

That can cause real routing conflicts.

Better:

```text
192.168.42.0/24 via Tailscale
10.121.15.0/24 via ZeroTier
10.1.0.0/24 via WireGuard
AstroWarp via mptun0 policy routing
```

Keep the routed subnets distinct and predictable.

---

## Final Result

After restoring the UCI settings, the Flint 3 successfully ran:

| Service | Result |
|---|---|
| AstroWarp | Working |
| Tailscale | Working |
| ZeroTier | Working |
| WireGuard Server | Working |

From a LAN client, both Tailscale and ZeroTier targets were reachable:

```powershell
ping 192.168.42.1
ping 10.121.15.226
```

That is exactly what I wanted: AstroWarp enabled without sacrificing the existing remote-access stack.

---

## Final Thoughts

The Flint 3 is flexible enough to run AstroWarp, Tailscale, ZeroTier, and WireGuard at the same time.

The Mesh Alpha firmware may try to protect users from route conflicts by disabling Tailscale and ZeroTier when AstroWarp is enabled. But in this test, the underlying system handled all of them once the settings were restored.

For advanced users, the fix is simple:

```sh
/root/fix-astrowarp-overlays.sh
```

Keep that script on the router, and you can quickly recover when the GL.iNet UI decides to "help."

---

## FAQ

### Can AstroWarp and Tailscale run together on the Flint 3?

Yes. They worked together after Tailscale was re-enabled in UCI and restarted.

### Can AstroWarp and ZeroTier run together on the Flint 3?

Yes. ZeroTier worked after restoring `zerotier.gl.enabled='1'` and restarting the service.

### Does AstroWarp break WireGuard?

In this test, WireGuard server continued working.

### Why did Tailscale stop working?

The GL.iNet management layer changed `tailscale.settings.enabled` to `0`, along with LAN/WAN/masquerade options.

### Why did ZeroTier stop working?

The GL.iNet management layer changed `zerotier.gl.enabled` to `0`.

### What is the quickest fix?

Run the recovery command block or use `/root/fix-astrowarp-overlays.sh`.

### Is this officially supported?

No. This is an advanced, unsupported configuration tested on Mesh Alpha firmware.

### Should everyone do this?

No. If you are not comfortable with SSH, UCI, and routing, stick with the GL.iNet UI-supported combinations.

