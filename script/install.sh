#!/bin/bash

go get ..
go build ..

cat <<EOF >>hero-backend.service
[Unit]
Description=Hero RESTful APIs

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/usr/bin/hero-backend

[Install]
WantedBy=multi-user.target
EOF

systemctl stop hero-backend.service
sudo cp hero-backend /usr/bin/
sudo cp hero-backend.service /lib/systemd/system/
systemctl start hero-backend.service
systemctl status hero-backend.service