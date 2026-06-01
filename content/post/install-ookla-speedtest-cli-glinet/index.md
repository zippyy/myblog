---
title: "Install Ookla Speedtest CLI on GL.iNet Routers in Seconds"
slug: "install-ookla-speedtest-cli-glinet"
date: 2026-06-01T09:00:00-06:00
lastmod: 2026-06-01T09:00:00-06:00
draft: false
description: "Install the official Ookla Speedtest CLI on GL.iNet routers using one command and run accurate speed tests directly from the router."
summary: "A quick guide for installing the official Ookla Speedtest CLI directly on ARM64 GL.iNet routers like the Flint 3, Slate 7, Slate 7 Pro, Mudi 7, and Spitz AX."
categories:
  - Networking
  - Homelab
tags:
  - GL.iNet
  - OpenWrt
  - Speedtest
  - Ookla
  - Flint 3
  - Slate 7
  - Slate 7 Pro
  - Travel Router
  - Networking
toc: true
usePageBundles: true
---

> **Affiliate disclosure:** This post may contain affiliate links. If you purchase through those links, Tech Relay may earn a commission at no additional cost to you.

## Overview

If you are troubleshooting internet speed, testing a VPN tunnel, comparing WAN connections, or validating a 5G modem, running a speed test from your laptop is not always the cleanest test.

Wi-Fi performance, browser overhead, client hardware, background apps, and local network congestion can all skew the results.

The better option is to run the test directly from the router.

This guide shows how to install the official **Ookla Speedtest CLI** on ARM64 GL.iNet routers using one command.

## Supported GL.iNet Routers

This method is for **aarch64 / ARM64** GL.iNet routers.

Good targets include:

- GL.iNet Flint 3 / GL-BE9300
- GL.iNet Slate 7 / GL-BE3600
- GL.iNet Slate 7 Pro / GL-BE10000
- GL.iNet Mudi 7
- GL.iNet Spitz AX
- Other ARM64 OpenWrt-based GL.iNet devices

Before installing, you can confirm the CPU architecture with:

```bash
uname -m
```

Expected output:

```text
aarch64
```

If your router does not show `aarch64`, do not use the ARM64 package shown in this guide.

## Hardware Used

### GL.iNet Flint 3 / GL-BE9300

The Flint 3 is a high-performance Wi-Fi 7 home router and a great place to run speed tests directly from the router.

**Affiliate link:**  
`https://amzn.to/4fhCJaf`

### GL.iNet Slate 7 / GL-BE3600

The Slate 7 is a compact Wi-Fi 7 travel router that can use Speedtest CLI to validate hotel internet, hotspot connections, tethering, or VPN performance.

**Affiliate link:**  
`https://amzn.to/4o55c5w`

### GL.iNet Slate 7 Pro / GL-BE10000

The Slate 7 Pro is the higher-end travel router option. If you are using it for OpenClash, WireGuard, Tailscale, AstroWarp, or travel routing, Speedtest CLI is an easy way to test real throughput from the router itself.

**Non-affiliate link:**  
`https://store.gl-inet.com/products/slate-7-pro-gl-be10000-tri-band-wi-fi-7-travel-router`

## Install Ookla Speedtest CLI

SSH into your GL.iNet router.

Most GL.iNet routers use:

```bash
ssh root@192.168.8.1
```

If you changed your LAN subnet, use your router's actual management IP.

Now run the install command:

```bash
curl -L "https://install.speedtest.net/app/cli/ookla-speedtest-1.2.0-linux-aarch64.tgz" | tar -xzvC /usr/bin/ speedtest
```

This command does a few things at once:

1. Downloads the official ARM64 Ookla Speedtest CLI archive.
2. Streams it directly into `tar`.
3. Extracts only the `speedtest` binary.
4. Places the binary directly into `/usr/bin/`.

That means there is no leftover archive to clean up afterward.

## Make Speedtest Executable

After installing the binary, make it executable:

```bash
chmod +x /usr/bin/speedtest
```

## Run Your First Speed Test

Now run:

```bash
speedtest
```

The first time you run it, Ookla will ask you to accept the license agreement and privacy policy.

After accepting, the test will start.

Example output will look similar to this:

```text
Speedtest by Ookla

Server: Example ISP - Denver, CO
ISP: Example Fiber

Idle Latency: 6.21 ms

Download: 941.82 Mbps
Upload:   928.44 Mbps
Packet Loss: 0.0%
```

## Useful Speedtest CLI Commands

### Show Available Servers

```bash
speedtest -L
```

This lists nearby Speedtest servers and their server IDs.

### Test Against a Specific Server

```bash
speedtest -s SERVER_ID
```

Example:

```bash
speedtest -s 12345
```

### Output JSON

JSON output is useful for scripts, dashboards, and logging.

```bash
speedtest -f json
```

### Output CSV

```bash
speedtest -f csv
```

### Check the Installed Version

```bash
speedtest --version
```

### Confirm the Binary Location

```bash
which speedtest
```

Expected output:

```text
/usr/bin/speedtest
```

## Automate Hourly Speed Tests

You can log periodic speed tests with cron.

Open the cron editor:

```bash
crontab -e
```

Add this line to run a JSON speed test every hour:

```bash
0 * * * * /usr/bin/speedtest -f json >> /root/speedtest.log
```

Restart cron if needed:

```bash
/etc/init.d/cron restart
```

You can view the log with:

```bash
tail -f /root/speedtest.log
```

## Why This Is Useful on GL.iNet Routers

Installing Speedtest CLI directly on the router is useful for testing:

- ISP speed from the WAN edge
- WireGuard throughput
- OpenVPN throughput
- Tailscale throughput
- AstroWarp performance
- OpenClash routing performance
- 5G modem speed
- Failover WAN speed
- Hotel or captive portal internet performance
- Travel router performance before connecting all your devices

This is especially helpful when comparing results before and after enabling VPN routing.

## Troubleshooting

### `speedtest: not found`

Confirm the binary exists:

```bash
ls -lah /usr/bin/speedtest
```

If it exists but will not run, reapply executable permissions:

```bash
chmod +x /usr/bin/speedtest
```

### Architecture Error

If the binary will not execute, verify the router architecture:

```bash
uname -m
```

This guide uses the `linux-aarch64` build. If your router is not ARM64, you need a different Speedtest CLI build.

### Missing `curl`

If your router does not have `curl`, install it with:

```bash
opkg update
opkg install curl
```

Then rerun the install command.

### Storage Space Issues

Check available space with:

```bash
df -h
```

The one-line install method keeps things cleaner because it does not leave the `.tgz` archive in `/tmp`.

## Full Install Block

If you just want the copy-and-paste version, use this:

```bash
curl -L "https://install.speedtest.net/app/cli/ookla-speedtest-1.2.0-linux-aarch64.tgz" | tar -xzvC /usr/bin/ speedtest
chmod +x /usr/bin/speedtest
speedtest
```

## Final Thoughts

The official Ookla Speedtest CLI is one of the easiest and most useful tools to add to a GL.iNet router.

Whether you are validating your ISP connection, benchmarking a VPN tunnel, testing a 5G gateway, or troubleshooting a travel router setup, running Speedtest directly from the router gives you a cleaner and more repeatable result than testing from a random client device.

<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "FAQPage",
  "mainEntity": [
    {
      "@type": "Question",
      "name": "Can I install Ookla Speedtest CLI on a GL.iNet router?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Yes. On ARM64 GL.iNet routers, you can install the official Ookla Speedtest CLI by downloading the linux-aarch64 package and extracting the speedtest binary to /usr/bin."
      }
    },
    {
      "@type": "Question",
      "name": "Which GL.iNet routers does this work on?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "This method is intended for aarch64 ARM64 GL.iNet routers such as the Flint 3, Slate 7, Slate 7 Pro, Mudi 7, Spitz AX, and other compatible OpenWrt-based ARM64 devices."
      }
    },
    {
      "@type": "Question",
      "name": "Why run Speedtest from the router instead of a laptop?",
      "acceptedAnswer": {
        "@type": "Answer",
        "text": "Running Speedtest directly from the router removes Wi-Fi, browser, and client-device limitations from the test, giving a cleaner view of WAN, VPN, or cellular modem performance."
      }
    }
  ]
}
</script>
