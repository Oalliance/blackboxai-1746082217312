#!/bin/bash
# Script to backup database and upload to AWS S3

DB_HOST="your-db-host"
DB_USER="your-db-user"
DB_NAME="your-db-name"
BACKUP_DIR="/tmp/db_backups"
S3_BUCKET="s3://your-backup-bucket"
TIMESTAMP=$(date +"%Y%m%d%H%M%S")
BACKUP_FILE="$BACKUP_DIR/db_backup_$TIMESTAMP.sql.gz"

mkdir -p $BACKUP_DIR

# Dump and compress database
pg_dump -h $DB_HOST -U $DB_USER $DB_NAME | gzip > $BACKUP_FILE

# Upload to S3
aws s3 cp $BACKUP_FILE $S3_BUCKET/

# Remove local backup
rm $BACKUP_FILE
