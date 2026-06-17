# AstroWarp + Tailscale + ZeroTier + WireGuard Command Reference

## Minimal Recovery

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
