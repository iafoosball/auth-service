#!/bin/bash
# This file should be used only outside the running container (on host).

API=auth
SERVICE_ADDR=$API-service:8070

ADMIN_HOSTNAME=localhost
ADMIN_PORT=8050
ADMIN_ADDR=$ADMIN_HOSTNAME:$ADMIN_PORT

# Get route id on Service
ROUTE_ID1=`curl -s "http://$ADMIN_ADDR/services/$API-service/routes" | jq ".data[0].id" | tr -d \" `
ROUTE_ID2=`curl -s "http://$ADMIN_ADDR/services/$API-service/routes" | jq ".data[1].id" | tr -d \" `
ROUTE_ID2=`curl -s "http://$ADMIN_ADDR/services/$API-service/routes" | jq ".data[2].id" | tr -d \" `
# Delete target on service
if [ "$ROUTE_ID1" != "" ] ; then
    curl -i -X DELETE "http://$ADMIN_ADDR/routes/$ROUTE_ID1"
else
    echo "Route to $API-service was not found"
fi

if [ "$ROUTE_ID2" != "" ] ; then
    curl -i -X DELETE "http://$ADMIN_ADDR/routes/$ROUTE_ID2"
else
    echo "Route to $API-service was not found"
fi

if [ "$ROUTE_ID3" != "" ] ; then
    curl -i -X DELETE "http://$ADMIN_ADDR/routes/$ROUTE_ID3"
else
    echo "Route to $API-service was not found"
fi
# Delete service
curl -i -X DELETE "http://$ADMIN_ADDR/services/$API-service"
# Delete UPSTREAM
curl -i -X DELETE "http://$ADMIN_ADDR/upstreams/$API-service"

