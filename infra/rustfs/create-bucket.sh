#!/bin/sh

ACCESS_KEY=$(cat /run/secrets/rustfs_access_key) &&
SECRET_KEY=$(cat /run/secrets/rustfs_secret_key) &&
sleep 5 &&

mc alias set local http://s3:9001 $ACCESS_KEY $SECRET_KEY &&

mc mb local/profile-photos || echo 'bucket exists' &&
mc ls local/profile-photos

echo 'Buckets created'
