#!/bin/bash

### BUILD
go get ..
go build ..
sudo cp hero-backend /usr/bin/

### CONFIGURE SYSLOG
cat <<EOF >/etc/rsyslog.d/hero_backend.conf
# shellcheck disable=SC2154
if $programname == 'hero_backend' then /var/log/hero_backend.log
& stop
EOF
touch /var/log/hero_backend.log
chown syslog /var/log/hero_backend.log
systemctl daemon-reload

### SETUP SYSTEMD SERVICE
systemctl stop hero_backend.service
cat <<EOF >/lib/systemd/system/hero_backend.service
[Unit]
Description=Hero RESTful APIs

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/usr/bin/hero-backend
StandardOutput=append:/var/log/hero_backend.log
StandardError=append:/var/log/hero_backend.log

[Install]
WantedBy=multi-user.target
EOF
systemctl daemon-reload
systemctl start hero_backend.service
