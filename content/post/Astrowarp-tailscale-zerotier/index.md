---
title: "Running AstroWarp, Tailscale, ZeroTier, and WireGuard Together on a GL.iNet Flint 3"
slug: "astrowarp-tailscale-zerotier-wireguard-flint3"
date: 2026-06-17T12:00:00-06:00
lastmod: 2026-06-17T12:00:00-06:00
draft: false
author: "Nicholas Bennett"
description: "AstroWarp, Tailscale, ZeroTier, and WireGuard can run together on a GL.iNet Flint 3. The catch is that GL.iNet’s management layer may disable Tailscale and ZeroTier when AstroWarp is enabled. Here is the fix."
summary: "A real-world Flint 3 Mesh Alpha firmware test showing how AstroWarp, Tailscale, ZeroTier, and WireGuard can coexist, why Tailscale and ZeroTier get disabled, and the exact recovery script to restore them."
keywords:
  - GL.iNet Flint 3
  - GL-BE9300
  - AstroWarp
  - Tailscale
  - ZeroTier
  - WireGuard
  - OpenWrt
  - Mesh Alpha Firmware
  - mptun0
  - tailscale0
  - zerotier-one
categories:
  - Networking
  - GL.iNet
tags:
  - Flint 3
  - GL.iNet
  - AstroWarp
  - Tailscale
  - ZeroTier
  - WireGuard
  - OpenWrt
  - VPN
  - Homelab
image: "/images/astrowarp-flint3-overlays/astrowarp-flint3-overlays.svg"
featured_image: "/images/astrowarp-flint3-overlays/astrowarp-flint3-overlays.svg"
toc: true
---

# Running AstroWarp, Tailscale, ZeroTier, and WireGuard Together on a GL.iNet Flint 3

GL.iNet’s Flint 3 is already a ridiculous little networking box. Between Wi-Fi 7, OpenWrt under the hood, WireGuard, Tailscale, ZeroTier, VLANs, and the newer AstroWarp feature, it can do a lot more than the average consumer router.

But there is one annoying behavior:

**When AstroWarp is enabled, GL.iNet may disable Tailscale and ZeroTier.**

At first glance, that makes it look like AstroWarp cannot coexist with Tailscale, ZeroTier, or WireGuard.

After testing it directly on a Flint 3 running the Mesh Alpha firmware, that is not what I found.

The underlying Linux networking stack can run all of them together:

- AstroWarp
- Tailscale
- ZeroTier
- WireGuard Server

The actual issue is that GL.iNet’s management layer flips the Tailscale and ZeroTier UCI settings to disabled when AstroWarp is enabled or refreshed.

This post documents the investigation, the commands used to prove what was happening, and the recovery script I now use to bring everything back when AstroWarp stomps on the config.

---

## The Setup

This was tested on:

```text
Device: GL.iNet Flint 3 / GL-BE9300
Firmware: Mesh Alpha firmware
Base: OpenWrt 23.05 snapshot-style build
Kernel observed: 5.4.213
```

The router was already running multiple overlay and VPN services:

| Service | Interface / Role |
|---|---|
| AstroWarp | `mptun0` |
| Tailscale | `tailscale0` |
| ZeroTier | `zt...` interface |
| WireGuard Server | `wgserver` |
| LAN | `br-lan`, using `192.168.80.0/24` |

The remote networks used during testing included:

```text
Tailscale remote subnet: 192.168.42.0/24
ZeroTier network:        10.121.15.0/24
Local LAN:               192.168.80.0/24
```

The exact IPs do not matter. Replace them with your own routes.

---

## The Symptom

After AstroWarp was enabled, Tailscale stopped responding:

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

At first, this looks like a service conflict.

But the real clue was in UCI.

---

## What AstroWarp Changed

Checking Tailscale:

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

Checking ZeroTier:

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

That explains why both daemons disappeared.

AstroWarp, or the GL.iNet management layer around AstroWarp, was not merely causing a route conflict. It was flipping both services off in UCI.

---

## Proving AstroWarp Itself Was Up

AstroWarp created the `mptun0` interface:

```sh
ip -br addr | grep mptun
```

Example output:

```text
mptun0           UNKNOWN        10.3.128.2/24
```

Interface statistics:

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

AstroWarp routing table:

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

That tells us AstroWarp is policy-routed. Traffic marked with `0x2` uses the `mptcp_mptun0` table, which defaults out `mptun0`.

This is very different from Tailscale or ZeroTier.

---

## How Tailscale Was Behaving

Tailscale uses policy routing and table `52`.

When Tailscale was running, the router could reach the remote Tailscale subnet:

```sh
ping 192.168.42.1
```

Example:

```text
PING 192.168.42.1 (192.168.42.1): 56 data bytes
64 bytes from 192.168.42.1: seq=0 ttl=64 time=19.552 ms
64 bytes from 192.168.42.1: seq=1 ttl=64 time=22.230 ms
```

Route lookup from the router:

```sh
ip route get 192.168.42.1
```

Output:

```text
192.168.42.1 dev tailscale0 table 52 src 100.97.167.103 uid 0
    cache
```

And table 52 contained the expected subnet route:

```sh
ip route show table 52
```

Relevant output:

```text
192.168.42.0/24 dev tailscale0
```

That proved the Tailscale side was technically fine once the daemon was running.

---

## The LAN Client Problem

The router itself could ping a remote Tailscale subnet, but a LAN client behind the router could not.

From a Windows PC:

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

That response came from the Flint 3 LAN gateway:

```text
192.168.80.1
```

So the PC was sending packets to the router, but the router was rejecting them.

The key UCI value was:

```text
tailscale.settings.masq='0'
```

Re-enabling Tailscale with masquerading fixed LAN client access.

---

## How ZeroTier Was Behaving

ZeroTier was simpler.

When disabled, it showed:

```sh
zerotier-cli info
```

Output:

```text
zerotier-cli: missing port and zerotier-one.port not found in /var/lib/zerotier-one
```

That means `zerotier-one` was not running.

After re-enabling it:

```sh
zerotier-cli info
```

Output:

```text
200 info 73da18a38c 1.14.1 ONLINE
```

Router-to-ZeroTier peer test:

```sh
ping 10.121.15.226
```

Example:

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
Reply from 10.121.15.226: bytes=32 time=23ms TTL=64
Reply from 10.121.15.226: bytes=32 time=28ms TTL=64
Reply from 10.121.15.226: bytes=32 time=24ms TTL=64
Reply from 10.121.15.226: bytes=32 time=21ms TTL=64
```

ZeroTier worked immediately once the daemon and UCI setting were restored.

---

## The Minimal Fix

This was the final command block that restored both Tailscale and ZeroTier:

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

That alone was enough to bring both services back.

---

## One-Command Recovery Script

Create this script on the router:

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

Run it whenever AstroWarp disables Tailscale or ZeroTier:

```sh
/root/fix-astrowarp-overlays.sh
```

---

## Optional Watchdog Script

If AstroWarp keeps disabling Tailscale and ZeroTier, create a watchdog.

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

Use the watchdog only if AstroWarp repeatedly disables the services. For most people, the manual fix script is cleaner.

---

## Useful Diagnostic Commands

### Check AstroWarp Interface

```sh
ip -br addr | grep mptun
```

### Check AstroWarp Routing Table

```sh
ip route show table mptcp_mptun0
```

### Check AstroWarp Policy Rules

```sh
ip rule show
```

### Check Tailscale Status

```sh
tailscale status
```

### Check Tailscale UCI

```sh
uci show tailscale
```

### Check Tailscale Routes

```sh
ip route show table 52
```

### Check ZeroTier Status

```sh
zerotier-cli info
zerotier-cli listnetworks
```

### Check ZeroTier UCI

```sh
uci show zerotier
```

### Check Interfaces

```sh
ip -br addr | grep -E 'mptun0|tailscale0|zt|wgserver'
```

### Check Running Processes

```sh
ps w | grep -E 'mptun|tailscaled|zerotier|wireguard'
```

---

## What This Proves

This testing showed that all of these can run together on the Flint 3:

| Service | Result |
|---|---|
| AstroWarp | Works |
| Tailscale | Works after UCI restore |
| ZeroTier | Works after UCI restore |
| WireGuard Server | Continues working |

The problem is not that Linux cannot run these overlays together.

The problem is that GL.iNet’s management layer disables some of them when AstroWarp is enabled.

---

## Important Caveats

This is not an officially supported configuration.

Future firmware updates may change:

- UCI paths
- init scripts
- firewall behavior
- AstroWarp behavior
- GL.iNet UI enforcement logic

Also, running multiple overlay networks at once can cause routing problems if you advertise overlapping subnets.

Avoid overlapping routes like:

```text
192.168.1.0/24 via Tailscale
192.168.1.0/24 via ZeroTier
192.168.1.0/24 via WireGuard
```

That is where real conflicts happen.

The working setup documented here used distinct routes.

---

## Final Thoughts

GL.iNet’s Flint 3 is flexible enough to run AstroWarp, Tailscale, ZeroTier, and WireGuard together.

The Mesh Alpha firmware tries to prevent that by disabling Tailscale and ZeroTier when AstroWarp is enabled. But once those settings are restored, all services can coexist.

For homelab users, consultants, and anyone building remote access networks, this is useful because it means AstroWarp does not have to replace your existing overlay network.

It can run beside it.

Just keep the recovery script handy.

---

## FAQ

### Can AstroWarp and Tailscale run at the same time?

Yes. In this test, AstroWarp and Tailscale ran at the same time once Tailscale was re-enabled through UCI.

### Can AstroWarp and ZeroTier run at the same time?

Yes. ZeroTier also worked after restoring `zerotier.gl.enabled='1'`.

### Does AstroWarp break WireGuard?

WireGuard server continued running in this setup.

### Why does Tailscale stop working after enabling AstroWarp?

The GL.iNet management layer changes Tailscale’s UCI settings to disabled. The daemon then stops.

### Why does ZeroTier stop working after enabling AstroWarp?

The GL.iNet management layer changes `zerotier.gl.enabled` to `0`, which prevents `zerotier-one` from running.

### What is the quickest fix?

Run:

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

### Is this safe?

It is safe enough for advanced users who understand routing and VPN overlays. It is not vendor-supported, and future firmware updates may change behavior.

## So Long and Thanks for All the FISH!!