#!/bin/bash

docker stop ory-hydra-iafoosball--hydra

docker rm ory-hydra-iafoosball--hydra

docker stop ory-hydra-iafoosball--postgres

docker rm ory-hydra-iafoosball--postgres