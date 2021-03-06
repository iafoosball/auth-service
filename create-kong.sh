#!/bin/bash
# This file should be used only outside the running container (on host).

API=auth
API_PORT=8010
SERVICE_ADDR=$API-service:$API_PORT

ADMIN_HOSTNAME=localhost
ADMIN_PORT=8051
ADMIN_ADDR=$ADMIN_HOSTNAME:$ADMIN_PORT

# Create Service
curl -i -X POST "http://$ADMIN_ADDR/services/" --header 'Content-Type: application/json' \
     --data '{ "name": "'$API'-service", "host": "'$API'-service", "port": '$API_PORT', "protocol": "http", "path": "/" }'

# Create Route to login
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "paths": ["/oauth/login"], "strip_path": false }'

# Get ID of created login route and register basic-auth on it
ROUTE_ID=`curl -s "http://$ADMIN_ADDR/services/$API-service/routes" | jq ".data[].id" | tr -d \" `
curl -i -X POST "http://$ADMIN_ADDR/routes/$ROUTE_ID/plugins" \
     --data 'name=basic-auth'

# Create Route to token verification, and strip path to hide it from client
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "paths": ["/oauth/verify"], "strip_path": true }'

# Create Route for other endpoints
curl -i -X POST "http://$ADMIN_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "paths": ["/oauth"], "strip_path": false  }'

# Add UPSTREAM
curl -i -X POST "http://$ADMIN_ADDR/upstreams/" --header 'Content-Type: application/json'  \
     --data '{ "name": "'$API'-service" }'

# Add target to upstream
curl -i -X POST "http://$ADMIN_ADDR/upstreams/${API}-service/targets" --header 'Content-Type: application/json'  \
     --data '{ "target": "'$SERVICE_ADDR'" }'
