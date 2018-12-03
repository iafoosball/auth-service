#!/bin/bash
# This file should be used only outside the running container (on host).

API=auth
SERVICE_ADDR=$API-service:8070

ADMIN_HOSTNAME=localhost
ADMIN_PORT=8050
ADMIN_ADDR=$ADMIN_HOSTNAME:$ADMIN_PORT

# Create Service
curl -i -X POST "http://$ADMIN_ADDR/services/" --header 'Content-Type: application/json' \
     --data '{ "name": "'$API'-service", "host": "'$API'-service", "port": '$ADMIN_PORT', "protocol": "http", "path": "/" }'

# Create Routes to service
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "paths": ["/oauth/login"] }'

# Get ID of created login route
ROUTE_ID=`curl -s "http://$ADMIN_ADDR/services/$API-service/routes" | jq ".data[].id" | tr -d \" `
# Register basic-auth plugin
curl -i -X POST "http://$ADMIN_ADDR/routes/$ROUTE_ID/plugins" \
     --data 'name=basic-auth'
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \

# Add UPSTREAM
curl -i -X POST "http://$ADMIN_ADDR/upstreams/" --header 'Content-Type: application/json'  \
     --data '{ "name": "'$API'-service" }'

# Add target to upstream
curl -i -X POST "http://$ADMIN_ADDR/upstreams/${API}-service/targets" --header 'Content-Type: application/json'  \
     --data '{ "target": "'$SERVICE_ADDR'" }'
