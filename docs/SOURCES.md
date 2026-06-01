# Sources Used

This production package is based on the source material provided by the user and rewritten in Tech Relay's voice with expanded implementation notes, troubleshooting, SEO support, and publishing assets.

## Source article
Colton Idle, "How to install tailscale on your Unifi router (UDM)", DEV Community.
https://dev.to/coltonidle/how-to-install-tailscale-on-your-unifi-router-udm-5a35

Important source points used:
- Goal: replace UniFi Teleport-style access with Tailscale.
- Goals: direct IP access, custom DNS access, and exit-node behavior.
- Install flow: enable SSH, run installer, tailscale up.
- Bullseye backports fix.
- dnsmasq tailscale0 fix.

## Project documentation
SierraSoftworks tailscale-unifi:
https://github.com/SierraSoftworks/tailscale-unifi

## Official Tailscale docs
Subnet routers:
https://tailscale.com/docs/features/subnet-routers

Exit nodes:
https://tailscale.com/docs/features/exit-nodes

DNS:
https://tailscale.com/docs/reference/dns-in-tailscale
