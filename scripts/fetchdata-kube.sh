#!/bin/sh

# sleep before fetching so that if we're in some sort of fast crashloop/reschedule mode, 
# we don't ping testgrid everytime we come back up
echo "Doing initial sleep before fetching testgrid data"
sleep 60 # 1 minutes
while [ true ]; do
  echo "Fetching new testgrid data"
  rm -rf /data/*
  /go/src/sippy/sippy --fetch-data /data --dashboard=kube-master=sig-release-master-blocking,sig-release-master-informing= -v 4
  echo "Done fetching data, refreshing server"
  curl localhost:8080/refresh
  echo "Done refreshing data, sleeping"
  sleep 3600  # 1 hour
done
