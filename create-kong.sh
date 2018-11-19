#!/bin/bash

KONG_ADDR=kong:8001
SERVICE_ADDR=auth-service:8070
API=auth

curl $KONG_ADDR

# curl -i -X POST http://kong:8001/services/ --header 'Content-Type: application/json' --data '{"name": "auth-service", "host": "auth-service", "port": 8070, "protocol": "http",
# "path": "/auth"}'
# Create Service
curl -i -X POST "http://$KONG_ADDR/services/" --header 'Content-Type: application/json' \
     --data '{ "name": "'$API'-service", "host": "'$API'-service", "port": 80, "protocol": "http", "path": "/'${API}'" }'
# curl -i -X POST http://kong:8001/services/auth-service/routes --header 'Content-Type: application/json' --data '{"paths": ["/auth"], "strip_path": true}'
# Create Route to service
curl -i -X POST "http://$KONG_ADDR/services/$API-service/routes" --header 'Content-Type: application/json' \
     --data '{ "paths": ["/'${API}'"], "strip_path": true }'
# curl -i -X POST http://kong:8001/upstreams --header 'Content-Type: application/json' --data '{"name": "auth-service"}'
# Add UPSTREAM
curl -i -X POST "http://$KONG_ADDR/upstreams/" --header 'Content-Type: application/json'  \
     --data '{ "name": "'${API}'-service" }'
# curl -i -X POST http://kong:8001/upstreams/auth-service/targets --header 'Content-Type: application/json' --data '{"target": "auth-service:8070"}'
# Add target to upstream
curl -i -X POST "http://$KONG_ADDR/upstreams/${API}-service/targets" --header 'Content-Type: application/json'  \
     --data '{ "target": "'$SERVICE_ADDR'" }'
