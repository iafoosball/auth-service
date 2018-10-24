#!/bin/bash

docker run --rm \
  --network kong_iafoosball \
  --name ory-hydra-iafoosball--postgres \
  -e POSTGRES_USER=hydra \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=hydra \
  -d postgres:10

# export SYSTEM_SECRET=$(export LC_CTYPE=C; cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
export SYSTEM_SECRET=this_needs_to_be_the_same_always_and_also_very_$3cuR3-._

export DATABASE_URL=postgres://hydra:secret@ory-hydra-iafoosball--postgres:5432/hydra?sslmode=disable

docker pull oryd/hydra:latest-alpine

# docker run -it --rm --entrypoint hydra oryd/hydra:latest-alpine help serve

docker run -it --rm \
  --network kong_iafoosball \
  oryd/hydra:latest-alpine \
  migrate sql $DATABASE_URL

docker run -d \
  --name ory-hydra-iafoosball--hydra \
  --network kong_iafoosball \
  -p 8070:4444 \
  -p 8071:4445 \
  -e SYSTEM_SECRET=$SYSTEM_SECRET \
  -e DATABASE_URL=$DATABASE_URL \
  -e OAUTH2_ISSUER_URL=https://localhost:8070/ \
  -e OAUTH2_CONSENT_URL=http://localhost:8001/consent \
  -e OAUTH2_LOGIN_URL=http://localhost:8001/login \
  oryd/hydra:latest-alpine serve all
