#!/bin/sh
# watch-astrowarp-overlays.sh
# Optional watchdog for GL.iNet AstroWarp repeatedly disabling Tailscale/ZeroTier.

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
