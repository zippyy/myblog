# Command Cheat Sheet

## Enable Tailscale on UniFi OS

```bash
curl -sSLq https://raw.githubusercontent.com/SierraSoftworks/tailscale-unifi/main/install.sh | sh
tailscale up
tailscale status
```

## Check UniFi OS Version

```bash
/usr/bin/ubnt-device-info firmware_detail
```

## Bullseye Backports Fix

```bash
vim /etc/apt/sources.list
```

Use:

```text
deb https://archive.debian.org/debian/ bullseye-backports main
```

## DNS Fix

```bash
touch /run/dnsmasq.dhcp.conf.d/tailscale0.conf
echo "interface=tailscale0" > /run/dnsmasq.dhcp.conf.d/tailscale0.conf
pkill dnsmasq
```

## Advertise Routes

```bash
tailscale up --advertise-routes=192.168.1.0/24
```

## Advertise Multiple Routes and Exit Node

```bash
tailscale up \
  --advertise-routes=192.168.1.0/24,192.168.10.0/24,192.168.42.0/24 \
  --advertise-exit-node
```

## Restart Tailscale

```bash
systemctl restart tailscaled
```

## Upgrade Tailscale

```bash
/data/tailscale/manage.sh update
```

## Remove Tailscale

```bash
/data/tailscale/manage.sh uninstall
```
