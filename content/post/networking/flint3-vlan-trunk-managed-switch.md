---
title: "How to Configure a VLAN Trunk Between a GL.iNet Flint 3 and Any Managed Switch"
slug: "flint3-vlan-trunk-managed-switch"
date: 2026-06-02T10:00:00-06:00
lastmod: 2026-06-02T10:00:00-06:00
draft: false
author: "Nicholas Bennett"
description: "Step-by-step guide for configuring a VLAN trunk port between a GL.iNet Flint 3 and any managed switch using LuCI, OpenWrt bridge VLAN filtering, and tagged VLANs."
summary: "Build a clean multi-VLAN trunk between a GL.iNet Flint 3 and any managed switch. This guide covers Flint 3 LuCI VLAN filtering, tagged trunk ports, switch-side setup, DHCP, firewall zones, and troubleshooting."
categories:
  - Networking
  - GL.iNet
  - OpenWrt
tags:
  - GL.iNet Flint 3
  - VLAN
  - Trunk Port
  - Managed Switch
  - OpenWrt
  - LuCI
  - Network Segmentation
  - Small Business Networking
keywords:
  - Flint 3 VLAN trunk
  - GL.iNet Flint 3 VLAN setup
  - OpenWrt bridge VLAN filtering
  - managed switch VLAN trunk
  - LuCI VLAN filtering
  - tagged VLAN trunk
  - GL-BE9300 VLAN
featured: true
weight: 20
toc: true
featured_image: "/images/posts/flint3-vlan-trunk-managed-switch/flint3-vlan-trunk-featured.svg"
images:
  - "/images/posts/flint3-vlan-trunk-managed-switch/flint3-vlan-trunk-featured.svg"
seo:
  title: "GL.iNet Flint 3 VLAN Trunk Setup for Any Managed Switch"
  description: "Configure a VLAN trunk between a GL.iNet Flint 3 and any managed switch with this practical OpenWrt/LuCI guide."
  canonical: "https://techrelay.xyz/posts/flint3-vlan-trunk-managed-switch/"
  noindex: false
openGraph:
  type: "article"
  title: "How to Configure a VLAN Trunk Between a GL.iNet Flint 3 and Any Managed Switch"
  description: "A practical Flint 3 VLAN trunk setup guide for OpenWrt, LuCI, and managed switches."
  image: "/images/posts/flint3-vlan-trunk-managed-switch/flint3-vlan-trunk-featured.svg"
twitter:
  card: "summary_large_image"
  title: "GL.iNet Flint 3 VLAN Trunk Setup Guide"
  description: "Build a clean VLAN trunk between a Flint 3 and any managed switch."
  image: "/images/posts/flint3-vlan-trunk-managed-switch/flint3-vlan-trunk-featured.svg"
schema:
  type: "TechArticle"
  faq: true
  howto: true
---

# How to Configure a VLAN Trunk Between a GL.iNet Flint 3 and Any Managed Switch

A VLAN trunk lets you carry multiple isolated networks over one Ethernet cable. If you are using a **GL.iNet Flint 3** as your router and a managed switch for your wired network, trunking is the clean way to move VLANs from the router to the switch without burning one physical port per network.

This guide walks through a generic setup that works with almost any managed switch: HP, Aruba, Cisco, UniFi, TP-Link, Netgear, MikroTik, and others.

The names change between vendors, but the concept is the same:

> The Flint 3 tags the VLANs. The switch port connected to the Flint 3 accepts those tagged VLANs. Access ports on the switch place client devices into the correct VLAN.

---

## Example Network Design

For this guide, we will use four VLANs:

| VLAN ID | Name | Example Subnet | Purpose |
|---:|---|---|---|
| 10 | Main | 192.168.10.0/24 | Trusted computers and phones |
| 20 | IoT | 192.168.20.0/24 | Smart TVs, plugs, cameras, random cloud gadgets |
| 30 | Guest | 192.168.30.0/24 | Guest Wi-Fi or temporary devices |
| 40 | Servers | 192.168.40.0/24 | NAS, Proxmox, lab servers, services |

You can replace these with your own VLAN IDs. The process is the same.

---

## Example Topology

```text
Internet
   |
WAN
   |
GL.iNet Flint 3
   |
LAN trunk port carrying VLANs 10, 20, 30, 40
   |
Managed switch
   |
Access ports for clients, APs, servers, IoT, or guest devices
```

The key idea is that the Ethernet cable between the Flint 3 and the managed switch carries multiple VLANs at the same time.

---

## Before You Start

You should have:

- A GL.iNet Flint 3
- A managed switch with VLAN support
- Admin access to LuCI on the Flint 3
- Admin access to your switch
- A list of VLAN IDs you want to use
- A backup or export of your current router configuration

Do not make VLAN changes remotely unless you have a fallback path. A bad VLAN setting can lock you out of the router.

---

## Step 1: Open LuCI on the Flint 3

The regular GL.iNet interface is great for normal router management, but VLAN trunking is handled through LuCI.

Open:

```text
http://192.168.8.1/cgi-bin/luci
```

Then go to:

```text
Network → Devices
```

Find:

```text
br-lan
```

Click **Configure**.

---

## Step 2: Enable Bridge VLAN Filtering

Inside the `br-lan` device settings, enable:

```text
Bridge VLAN Filtering
```

This tells OpenWrt to treat the LAN bridge as a VLAN-aware bridge.

After enabling VLAN filtering, you will be able to define which VLANs are tagged, untagged, or excluded on each bridge port.

---

## Step 3: Pick the Flint 3 Trunk Port

Choose one physical LAN port on the Flint 3 to connect to your managed switch.

For example:

```text
Flint 3 LAN port → Managed switch uplink port
```

In LuCI, the port names may appear as `lan1`, `lan2`, `lan3`, `lan4`, or sometimes as numbered ports depending on firmware and switch driver behavior.

If you are unsure which LuCI port matches the physical jack, plug a laptop into the port and check link status, or verify from the command line later with:

```bash
bridge link
bridge vlan show
```

---

## Step 4: Create the VLAN Table on the Flint 3

In:

```text
Network → Devices → br-lan → Configure → Bridge VLAN Filtering
```

Add each VLAN.

For a pure tagged trunk carrying VLANs 10, 20, 30, and 40, the table should conceptually look like this:

| VLAN | CPU / br-lan self | Flint trunk port |
|---:|---|---|
| 10 | Tagged | Tagged |
| 20 | Tagged | Tagged |
| 30 | Tagged | Tagged |
| 40 | Tagged | Tagged |

The important part:

```text
CPU/self = Tagged
Trunk port = Tagged
```

The CPU/self side must be tagged so OpenWrt can create and use interfaces such as `br-lan.10`, `br-lan.20`, `br-lan.30`, and `br-lan.40`.

The trunk port must be tagged so the managed switch receives the VLAN tags.

---

## Step 5: Decide Whether You Need a Native/Untagged VLAN

Some networks use a native VLAN on the trunk. That means one VLAN travels untagged across the trunk.

For example:

| VLAN | CPU / br-lan self | Flint trunk port | PVID |
|---:|---|---|---|
| 10 | Tagged | Untagged | Yes |
| 20 | Tagged | Tagged | No |
| 30 | Tagged | Tagged | No |
| 40 | Tagged | Tagged | No |

This can be useful if your switch expects an untagged management VLAN.

However, for cleaner setups, especially with managed switches, I usually prefer tagging every VLAN on the trunk when possible.

Recommended simple approach:

```text
Tag all VLANs on the trunk.
Avoid native VLANs unless you specifically need one.
```

---

## Step 6: Create Flint 3 Interfaces for Routed VLANs

Now go to:

```text
Network → Interfaces
```

Create a new interface for each VLAN the Flint 3 will route.

Example:

| Interface Name | Device | Protocol | Example IP |
|---|---|---|---|
| VLAN10_MAIN | br-lan.10 | Static address | 192.168.10.1/24 |
| VLAN20_IOT | br-lan.20 | Static address | 192.168.20.1/24 |
| VLAN30_GUEST | br-lan.30 | Static address | 192.168.30.1/24 |
| VLAN40_SERVERS | br-lan.40 | Static address | 192.168.40.1/24 |

Use whatever naming style makes sense for your network.

You only need an interface if the Flint 3 should participate in that VLAN. If the Flint 3 is routing, providing DHCP, or firewalling that VLAN, create the interface.

If the Flint 3 is only passing a VLAN through at Layer 2, you may not need a routed interface for that VLAN.

---

## Step 7: Enable DHCP Where Needed

For each VLAN where the Flint 3 should hand out IP addresses:

Go to:

```text
Network → Interfaces → Edit VLAN interface → DHCP Server
```

Enable DHCP.

Example DHCP ranges:

| VLAN | Router IP | DHCP Range |
|---:|---|---|
| 10 | 192.168.10.1 | 192.168.10.100 - 192.168.10.249 |
| 20 | 192.168.20.1 | 192.168.20.100 - 192.168.20.249 |
| 30 | 192.168.30.1 | 192.168.30.100 - 192.168.30.249 |
| 40 | 192.168.40.1 | 192.168.40.100 - 192.168.40.249 |

For server VLANs, you may choose a smaller DHCP range or use static addresses.

---

## Step 8: Create Firewall Zones

VLANs are not very useful if everything can freely talk to everything else. The firewall is where segmentation actually becomes security.

A common setup:

| VLAN | Zone | Input | Output | Forward |
|---:|---|---|---|---|
| 10 | main | Accept | Accept | Accept |
| 20 | iot | Reject | Accept | Reject |
| 30 | guest | Reject | Accept | Reject |
| 40 | servers | Reject or Accept | Accept | Reject |

Then allow forwarding from each limited zone to WAN:

```text
iot → wan
guest → wan
servers → wan
```

Avoid allowing unrestricted forwarding between VLANs unless you actually need it.

Better approach:

- Let IoT reach the internet.
- Block IoT from reaching your main LAN.
- Let guest devices reach only the internet.
- Allow main LAN to reach servers.
- Add specific firewall rules for services you intentionally expose.

---

## Step 9: Configure the Managed Switch Uplink Port

On the managed switch, the port connected to the Flint 3 must carry the same VLANs.

Different vendors use different words:

| Vendor / Platform | Typical Term |
|---|---|
| Cisco | Trunk port |
| HP / Aruba | Trunk or tagged VLAN membership |
| UniFi | Port profile with tagged networks |
| TP-Link | Tagged VLANs |
| Netgear | Tagged VLAN membership |
| MikroTik | Bridge VLAN table / tagged port |
| Generic web UI | Tagged member |

The switch uplink should be:

```text
Tagged VLANs: 10, 20, 30, 40
```

If you are using a native VLAN, set the matching untagged/native VLAN on both the Flint 3 and the switch. Do not make one side tagged and the other side untagged for the same VLAN unless you are intentionally translating behavior at the edge.

---

## Step 10: Configure Switch Access Ports

Client-facing switch ports are usually access ports.

Examples:

| Switch Port | Mode | VLAN |
|---|---|---:|
| Port 1 | Access / Untagged | 10 |
| Port 2 | Access / Untagged | 20 |
| Port 3 | Access / Untagged | 30 |
| Port 4 | Access / Untagged | 40 |

A normal laptop, printer, TV, or desktop usually expects untagged traffic. The switch places that untagged client traffic into the assigned VLAN.

For access points that broadcast multiple SSIDs mapped to different VLANs, the AP port usually needs to be a trunk too.

---

## Example: Generic Switch Uplink Settings

Your switch uplink to the Flint 3 should look conceptually like this:

```text
Port connected to Flint 3:
Mode: Trunk / Tagged
Tagged VLANs: 10, 20, 30, 40
Native/Untagged VLAN: none, unless required
```

A regular client port should look like this:

```text
Port connected to IoT device:
Mode: Access / Untagged
Untagged VLAN: 20
PVID: 20
```

The exact menu names differ, but these settings exist in some form on nearly every managed switch.

---

## Step 11: Verify from the Flint 3

SSH into the Flint 3 and run:

```bash
bridge vlan show
```

You should see the trunk port listed with the VLANs you configured.

You can also check interfaces:

```bash
ip addr show
```

You should see devices like:

```text
br-lan.10
br-lan.20
br-lan.30
br-lan.40
```

If you enabled DHCP, connect a client to a switch access port and verify it receives the correct subnet.

---

## Step 12: Test Each VLAN

For each VLAN, test:

1. Does the client receive the correct IP range?
2. Can the client reach the internet?
3. Can the client reach the Flint 3 gateway IP for that VLAN?
4. Is inter-VLAN access blocked or allowed as intended?
5. Does DNS resolve correctly?

Example:

```bash
ping 192.168.20.1
nslookup techrelay.xyz
ping 1.1.1.1
```

If DNS fails but pinging an IP works, the VLAN path is probably fine and the issue is DNS or firewall input rules.

---

## Common Problems

### Client Gets No IP Address

Check:

- VLAN exists on the Flint 3 bridge VLAN table.
- Trunk port is tagged for that VLAN on both sides.
- Switch access port has the correct untagged VLAN and PVID.
- DHCP is enabled on the Flint 3 VLAN interface.
- Firewall input is not blocking DHCP.

### Client Gets an IP but No Internet

Check:

- Firewall zone allows forwarding to WAN.
- The VLAN interface has the correct gateway IP.
- DNS is configured.
- Masquerading is enabled on the WAN zone.

### You Lose Access to the Flint 3

This usually happens when the management VLAN gets moved or untagged traffic changes unexpectedly.

Recovery options:

- Try another LAN port that still has the default LAN.
- Set a static IP on your laptop in the expected subnet.
- Connect over Wi-Fi if Wi-Fi is still bridged to management.
- Factory reset only as a last resort.

### VLAN Works on One Switch But Not Another

Vendor terminology is usually the problem.

Look for these concepts:

- Tagged VLAN membership
- Untagged VLAN membership
- PVID
- Native VLAN
- Trunk port
- Access port
- Allowed VLANs

Do not rely only on the word “trunk.” Some switches use “trunk” to mean link aggregation instead of VLAN trunking.

---

## Security Notes

VLANs are segmentation, not magic.

Good VLAN design should be paired with:

- Firewall rules
- Strong Wi-Fi security
- Separate guest networks
- Limited management access
- Regular firmware updates
- Backups of router and switch configs
- Monitoring for unknown devices

The point is not to make the network complicated. The point is to make it controlled.

---

## Final Thoughts

The GL.iNet Flint 3 is a strong little router for advanced home labs and small-business networks because it exposes OpenWrt under the hood. With LuCI bridge VLAN filtering and a managed switch, you can build a clean multi-VLAN network using a single trunk cable.

Once the trunk is working, scaling is simple. Add the VLAN on the Flint 3, allow it on the switch uplink, assign switch ports or SSIDs, and build firewall rules that match how the network should actually behave.

For most networks, this is the point where the setup starts feeling less like a pile of devices and more like an intentional infrastructure design.

---

## FAQ

### Do I need a managed switch for VLANs?

Yes. If you want to carry multiple VLANs over one cable and assign ports to different VLANs, you need a managed switch or smart switch with VLAN support.

### Can an unmanaged switch pass VLAN tags?

Sometimes unmanaged switches will pass tagged frames, but they cannot properly assign access ports, manage PVIDs, or enforce VLAN membership. Use a managed switch.

### Should my trunk port be tagged or untagged?

For a clean router-to-switch trunk, tag every VLAN unless you specifically need a native VLAN.

### What is the CPU or self port in LuCI?

The CPU/self entry represents the router side of the bridge. It must be tagged for VLANs that OpenWrt needs to route, firewall, or provide DHCP for.

### Do I need DHCP on every VLAN?

Only if the Flint 3 should hand out addresses on that VLAN. Some VLANs may use static addressing or a different DHCP server.

### Can I use the same trunk for access points?

Yes. If your access point supports VLAN-tagged SSIDs, connect it to a switch port configured as a trunk and tag the SSID VLANs on that port.

### Why does my switch use the word trunk for link aggregation?

Some vendors use “trunk” to mean LAG or port-channel. In VLAN terms, you are looking for tagged VLAN membership or 802.1Q VLAN settings.

### Can I route between VLANs?

Yes, but only allow what you actually need. The Flint 3 firewall can permit specific traffic between VLANs while blocking everything else.
