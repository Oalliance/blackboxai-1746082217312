#!/bin/bash
# Script to setup log rotation for application logs

LOG_FILE="/var/log/logistics-marketplace/app.log"
ROTATE_CONF="/etc/logrotate.d/logistics-marketplace"

sudo bash -c "cat > $ROTATE_CONF" <<EOL
$LOG_FILE {
    daily
    missingok
    rotate 14
    compress
    delaycompress
    notifempty
    create 0640 root root
    sharedscripts
    postrotate
        systemctl restart logistics-marketplace.service > /dev/null 2>&1 || true
    endscript
}
EOL

echo "Log rotation configured for $LOG_FILE"
