#!/bin/sh
set -e

mkdir -p /opt/xray
cd /opt/xray

wget -O xray.zip https://github.com/XTLS/Xray-core/releases/latest/download/Xray-linux-arm64-v8a.zip
unzip -o xray.zip
chmod +x xray

cat >/etc/init.d/xray <<'EOF'
#!/bin/sh /etc/rc.common

START=99
USE_PROCD=1

start_service() {
    procd_open_instance
    procd_set_param command /opt/xray/xray run -config /opt/xray/config.json
    procd_set_param env XRAY_LOCATION_ASSET=/opt/xray
    procd_set_param respawn
    procd_set_param stdout 1
    procd_set_param stderr 1
    procd_close_instance
}
EOF

chmod +x /etc/init.d/xray
/etc/init.d/xray enable

echo "Now create /opt/xray/config.json, then run: /etc/init.d/xray start"
