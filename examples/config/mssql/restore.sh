#!/bin/sh

sqlpackage /Action:Publish \
  /TargetServerName:127.0.0.1,1433 \
  /TargetDatabaseName:example \
  /TargetUser:sa \
  /TargetPassword:Password1 \
  /SourceFile:./tmp/backup.dacpac \
  /TargetTrustServerCertificate:True
