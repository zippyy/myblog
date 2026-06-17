#!/bin/sh
# fix-astrowarp-overlays.sh
# Restores Tailscale and ZeroTier after GL.iNet AstroWarp disables them.
# Tested on GL.iNet Flint 3 / GL-BE9300 Mesh Alpha firmware.

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
