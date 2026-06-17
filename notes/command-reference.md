# AstroWarp + Tailscale + ZeroTier + WireGuard Flint 3 Command Reference

## Minimal recovery

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

## Install recovery script

```sh
cat >/root/fix-astrowarp-overlays.sh <<'EOF'
#!/bin/sh
uci set tailscale.settings.enabled='1'
uci set tailscale.settings.lan_enabled='1'
uci set tailscale.settings.wan_enabled='1'
uci set tailscale.settings.masq='1'
uci set zerotier.gl.enabled='1'
uci commit tailscale
uci commit zerotier
/etc/init.d/tailscale restart
/etc/init.d/zerotier restart
sleep 5
tailscale status
zerotier-cli info
ip -br addr | grep -E 'tailscale0|zt|mptun0|wgserver' || true
EOF
chmod +x /root/fix-astrowarp-overlays.sh
```

## Run recovery

```sh
/root/fix-astrowarp-overlays.sh
```

## Diagnostics

```sh
uci show tailscale
uci show zerotier
tailscale status
zerotier-cli info
zerotier-cli listnetworks
ip -br addr | grep -E 'mptun0|tailscale0|zt|wgserver'
ip route show table 52
ip route show table mptcp_mptun0
ip rule show
```

## Test from router

```sh
ping 192.168.42.1
ping 10.121.15.226
```

## Test from LAN client

```powershell
ping 192.168.42.1
ping 10.121.15.226
```
