+++
categories = ['Technology', 'Networking', 'Gaming', 'OpenWrt']
codeLineNumbers = true
codeMaxLines = 30
date = "2026-06-02T13:15:00-06:00"
year = "2026"
month = "2026-06"
description = 'Install NetEase UU Game Booster on a GL.iNet router running OpenWrt so consoles and gaming devices can use router-level game acceleration.'
draft = false
featureImage = 'uu-game-booster-glinet-router.svg'
featureImageAlt = 'GL.iNet router providing NetEase UU Game Booster acceleration for a gaming console and handheld device'
featureImageCap = 'Router-level game acceleration lets consoles and handhelds use UU without installing a client on each device.'
figurePositionShow = true
shareImage = 'uu-game-booster-glinet-router.svg'
tags = ['GL.iNet', 'OpenWrt', 'Gaming', 'UU Game Booster', 'NetEase UU', 'Nintendo Switch', 'PlayStation', 'Xbox']
thumbnail = 'uu-game-booster-glinet-router.svg'
title = 'How to Install NetEase UU Game Booster on a GL.iNet Router'
toc = true
usePageBundles = true
series = 'GL.iNet Router Guides'
keywords = ['UU Game Booster', 'NetEase UU', 'GL.iNet router', 'OpenWrt UU plugin', 'game acceleration router', 'Nintendo Switch router acceleration', 'PlayStation router acceleration', 'Xbox router acceleration']
+++

## Overview

NetEase UU Game Booster is a gaming acceleration service that can optimize the route between your local network and supported game services. A lot of players know it as a PC or mobile app, but the router plugin is the more interesting option for home networks because it can help devices that cannot run the desktop client directly.

That matters for consoles and handhelds. A PlayStation, Xbox, Nintendo Switch, Steam Deck, gaming handheld, or living-room PC can sit behind the router and use the acceleration path selected by UU. Instead of installing software on every device, you install the router plugin once, bind the router in the UU app, and then manage acceleration from your phone.

GL.iNet routers are a good fit for this because they run OpenWrt underneath the normal GL.iNet interface. The GL.iNet web UI is friendly, but SSH gives you the same package manager and Linux tools you would expect on OpenWrt. This guide walks through installing the UU router plugin on a GL.iNet router, verifying TUN support, binding the router, and troubleshooting the common failure points.

## What this setup does

The finished setup looks like this:

```text
Gaming device -> GL.iNet router -> UU acceleration service -> game server
```

Your console or gaming device stays configured like a normal client on your LAN. The GL.iNet router handles the UU plugin, tunneling, and acceleration logic. The UU mobile app is used to bind and manage the router.

This guide covers:

- Installing the required TUN kernel module
- Downloading the UU OpenWrt installer
- Running the installer with the correct architecture argument
- Verifying the service is installed
- Binding the router in the UU app
- Adding basic troubleshooting commands
- Creating an optional reinstall script for repeat deployments

## Before you start

You need the following:

- A GL.iNet router with internet access
- SSH access to the router
- The router admin password
- A phone connected to the router Wi-Fi
- The UU Game Booster mobile app
- A UU account

You should also update the router firmware before you begin. Do not start by installing random packages on a router that is already unstable, low on storage, or halfway through another network change.

## Step 1: Find the router IP address

Most GL.iNet routers use this LAN address by default:

```bash
192.168.8.1
```

If you changed your LAN subnet, use the address of your router instead. On a client connected to the router, you can check the default gateway.

On Windows:

```powershell
ipconfig
```

Look for **Default Gateway** under the active Ethernet or Wi-Fi adapter.

On macOS or Linux:

```bash
ip route | grep default
```

The gateway address is the router address you will SSH into.

## Step 2: SSH into the GL.iNet router

From Windows Terminal, PowerShell, macOS Terminal, or Linux Terminal, connect as root:

```bash
ssh root@192.168.8.1
```

If your router uses a different IP, replace `192.168.8.1` with the router address.

Enter your router password when prompted. On GL.iNet firmware, the SSH password is usually the same password used for the GL.iNet admin interface unless you changed it.

## Step 3: Confirm the router architecture

The UU installer expects an architecture value. The common install command passes the output of `uname -m` automatically, but it is still useful to see what your router reports.

Run:

```bash
uname -m
```

Common outputs include:

```text
aarch64
armv7l
x86_64
mips
mipsel
```

Modern higher-end GL.iNet routers are often ARM64/aarch64, but do not guess. Let the router report it.

## Step 4: Update OpenWrt package lists

Update the package lists before installing anything:

```bash
opkg update
```

If this fails, fix basic internet and DNS first. The router must be able to resolve package repository hostnames and reach the internet.

A quick DNS test:

```bash
nslookup openwrt.org
```

A quick connectivity test:

```bash
ping -c 4 1.1.1.1
```

If IP ping works but DNS lookup fails, the router has a DNS problem. Fix that before continuing.

## Step 5: Install TUN support

The UU plugin needs TUN support because it creates tunnel interfaces for traffic handling. Install `kmod-tun`:

```bash
opkg install kmod-tun
```

Then verify that the TUN device exists:

```bash
ls -l /dev/net/tun
```

Expected result:

```text
crw-rw-rw-    1 root     root       10, 200 /dev/net/tun
```

If the file exists, continue. If it does not exist, reboot and check again:

```bash
reboot
```

After the router comes back online:

```bash
ssh root@192.168.8.1
ls -l /dev/net/tun
```

Do not skip this step. If TUN is missing, the plugin may install but fail to function correctly.

## Step 6: Download the UU OpenWrt installer

Change into `/tmp` so the installer is not stored permanently on the overlay:

```bash
cd /tmp
```

Download the installer:

```bash
wget http://uu.gdl.netease.com/uuplugin-script/20231117102400/install.sh -O install.sh
```

Make it executable:

```bash
chmod +x install.sh
```

The public OpenWrt install method commonly uses NetEase's `uuplugin-script` installer and runs it with `openwrt` plus the router architecture. Several OpenWrt guides and community notes document the same general pattern: download `install.sh`, run it with `openwrt $(uname -m)`, and make sure `kmod-tun` is present.

## Step 7: Run the installer

Run:

```bash
/bin/sh install.sh openwrt $(uname -m)
```

The installer should download and place the plugin files for your router architecture. Watch the output. If the installer prints an unsupported architecture message, the router may not have a matching plugin build.

When the installer completes, check for UU-related files and processes.

```bash
ps | grep -i uu
```

Also check init scripts:

```bash
ls /etc/init.d | grep -i uu
```

Depending on the plugin build, service names can vary, so use the broad `grep -i uu` checks first.

## Step 8: Start and enable the service

If the installer created an init script, enable it at boot and start it now.

First list the service name:

```bash
ls /etc/init.d | grep -i uu
```

If the service is named `uuplugin`, run:

```bash
/etc/init.d/uuplugin enable
/etc/init.d/uuplugin start
```

If the service is named `uu`, run:

```bash
/etc/init.d/uu enable
/etc/init.d/uu start
```

Then check processes again:

```bash
ps | grep -i uu
```

If you are not sure which init script exists, this command will show the matching files:

```bash
for f in /etc/init.d/*; do grep -qi uu "$f" 2>/dev/null && echo "$f"; done
```

## Step 9: Confirm the LAN bridge exists

UU router binding expects the router to behave like a normal bridge/NAT router. On GL.iNet firmware, the LAN bridge is usually `br-lan`.

Check it:

```bash
ip link show br-lan
```

You should see output showing the bridge interface. If `br-lan` does not exist, your router may be in an unusual mode or your LAN configuration may be customized.

Also check LAN addressing:

```bash
ip addr show br-lan
```

You should see the router LAN IP on `br-lan`.

## Step 10: Bind the router in the UU mobile app

Now move to your phone.

1. Connect the phone to the GL.iNet router Wi-Fi.
2. Open the UU Game Booster app.
3. Sign in to your UU account.
4. Look for the router/plugin acceleration section.
5. Add or bind a router.
6. Follow the in-app discovery and pairing steps.

The phone and router should be on the same LAN during binding. Do not try to bind while your phone is on cellular data.

After binding, select the device or game platform you want to accelerate. Depending on the app version and region, the labels may differ, but the basic workflow is the same: bind router, choose device/platform, enable acceleration.

## Step 11: Test from a gaming device

Connect your console or gaming device to the GL.iNet router. Wired Ethernet is best for testing because it removes Wi-Fi from the troubleshooting path.

Recommended test order:

1. Confirm the device gets an IP address from the GL.iNet router.
2. Confirm the device can reach the internet normally.
3. Enable acceleration in the UU app.
4. Launch the game.
5. Compare latency, matchmaking behavior, and disconnects.

Do not judge the setup from a single speed test. Gaming acceleration is about route quality, packet loss, latency, and stability to specific game servers. It is not the same thing as maximizing download speed.

## Optional: one-command install script

For repeat deployments, create this helper script on the router.

```bash
cat > /root/install-uu-game-booster.sh <<'EOF'
#!/bin/sh
set -e

echo "[1/6] Updating package lists..."
opkg update

echo "[2/6] Installing kmod-tun..."
opkg install kmod-tun || true

echo "[3/6] Checking TUN..."
if [ ! -e /dev/net/tun ]; then
  echo "ERROR: /dev/net/tun is missing. Reboot and try again."
  exit 1
fi

echo "[4/6] Downloading UU installer..."
cd /tmp
wget http://uu.gdl.netease.com/uuplugin-script/20231117102400/install.sh -O install.sh
chmod +x install.sh

echo "[5/6] Running UU installer..."
/bin/sh install.sh openwrt $(uname -m)

echo "[6/6] Checking installation..."
ps | grep -i uu | grep -v grep || true
ls /etc/init.d | grep -i uu || true

echo "Done. Connect your phone to this router and bind it in the UU app."
EOF

chmod +x /root/install-uu-game-booster.sh
/root/install-uu-game-booster.sh
```

This does not replace understanding the steps, but it makes the process easier when you are testing multiple GL.iNet routers.

## Optional: basic health check script

After installation, you can create a quick diagnostic script:

```bash
cat > /root/check-uu-game-booster.sh <<'EOF'
#!/bin/sh

echo "=== Router ==="
cat /etc/openwrt_release 2>/dev/null | sed -n '1,8p'
echo

echo "=== Architecture ==="
uname -m
echo

echo "=== TUN ==="
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
EOF

chmod +x /root/check-uu-game-booster.sh
/root/check-uu-game-booster.sh
```

This gives you a simple checklist: firmware info, architecture, TUN, LAN bridge, processes, init script, and DNS.

## Troubleshooting

### The installer fails to download

Check internet and DNS from the router:

```bash
ping -c 4 1.1.1.1
nslookup uu.163.com
```

If DNS fails, configure a working DNS server on the router and retry.

### `/dev/net/tun` is missing

Install the kernel module:

```bash
opkg update
opkg install kmod-tun
```

Then reboot:

```bash
reboot
```

After reboot:

```bash
ls -l /dev/net/tun
```

If the package cannot be installed because of a kernel mismatch, you may be running a firmware build whose package repositories no longer match the installed kernel. In that case, update or reinstall the router firmware so package versions line up again.

### The phone cannot find the router

Make sure your phone is connected to the GL.iNet router's LAN or Wi-Fi. Do not bind over cellular data.

Then check that the LAN bridge exists:

```bash
ip link show br-lan
ip addr show br-lan
```

If the router is in repeater, extender, AP-only, or another non-standard mode, temporarily test in normal router mode.

### The service is installed but not running

Search for the service:

```bash
ls /etc/init.d | grep -i uu
```

Start the matching service name:

```bash
/etc/init.d/uuplugin start
```

or:

```bash
/etc/init.d/uu start
```

Then check logs:

```bash
logread | grep -i uu
```

### Games still have bad latency

UU can improve routing, but it cannot fix every upstream issue. Check the basics:

- Use Ethernet for the console if possible.
- Test with the router close to the modem or primary gateway.
- Avoid double Wi-Fi hops.
- Confirm your WAN connection is stable.
- Reboot the router after installation.
- Test more than one game/server region.

### Download speed looks slower

Gaming acceleration is not the same as raw throughput optimization. Some acceleration paths may reduce peak download speed while improving latency or packet consistency for a game. Judge it by in-game ping, packet loss, disconnects, and matchmaking behavior instead of only using a speed test.

## Best practices for GL.iNet routers

### Use wired backhaul when possible

For consoles, Ethernet is still the cleanest option. If you are using a portable GL.iNet router, connect the console or gaming PC by Ethernet for your first test.

### Keep the router simple

Install UU on a clean router first. Once you confirm it works, then decide whether you want to combine it with VLANs, policy routing, VPN clients, or other advanced features.

### Watch storage space

Check available space before and after installation:

```bash
df -h
```

If overlay storage is nearly full, remove unused packages or use a router with more storage.

### Keep a backup

Before changing a router you rely on, back up the GL.iNet configuration from the web UI. If something goes sideways, restoring a backup is faster than rebuilding everything from memory.

## Uninstall notes

The exact uninstall process can vary depending on the plugin version and package format. Start by checking whether it was installed as a package:

```bash
opkg list-installed | grep -i uu
```

If an installed package appears, remove it with:

```bash
opkg remove PACKAGE_NAME
```

Replace `PACKAGE_NAME` with the actual package name shown by `opkg list-installed`.

Also check init scripts:

```bash
ls /etc/init.d | grep -i uu
```

Disable any matching service before removing files:

```bash
/etc/init.d/uuplugin stop
/etc/init.d/uuplugin disable
```

or:

```bash
/etc/init.d/uu stop
/etc/init.d/uu disable
```

If you manually remove files, be careful. Do not delete random system files just because they have similar names.

## FAQ

### Does UU Game Booster work on GL.iNet routers?

Yes, many GL.iNet routers can run the UU OpenWrt router plugin because GL.iNet firmware is OpenWrt-based. The router still needs compatible architecture support, working internet access, and TUN support.

### Do I need `kmod-tun`?

Yes. TUN support is one of the most important requirements. Install `kmod-tun` and verify `/dev/net/tun` before assuming the plugin is broken.

### Can this accelerate Nintendo Switch, PlayStation, and Xbox?

Yes. That is one of the main reasons to use the router plugin instead of a PC-only client. The console does not need to run UU directly; it connects through the router.

### Is this the same thing as a VPN?

No. A VPN generally tunnels traffic through a selected endpoint for privacy, access, or routing. UU Game Booster is focused on supported game acceleration and route optimization.

### Should I install this on my main router?

If this is your first test, use a spare GL.iNet router or a simple test network. Once you know it works with your account, devices, and games, you can decide whether to move it to the main router.

### What if my router architecture is not supported?

The installer may fail or report an unsupported platform. In that case, try a different GL.iNet model or an OpenWrt device with a supported architecture.

### Why does the app not discover my router?

The most common causes are the phone being on cellular instead of Wi-Fi, the router not being in normal router mode, missing LAN bridge behavior, or the UU service not running.

### Will this improve speed tests?

Not necessarily. The goal is better game routing, lower latency, fewer disconnects, and more stable gameplay. Raw download speed is not the main measurement.

## Final thoughts

A GL.iNet router plus the UU Game Booster OpenWrt plugin is a clean way to provide gaming acceleration to devices that cannot run a desktop client. The process is straightforward: install TUN support, run the UU OpenWrt installer, confirm the service is active, and bind the router from the mobile app.
