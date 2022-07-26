#!/bin/bash

if [[ $1 == "UP" ]]; then
  echo "Starting localstack..."
  TMPDIR=/private$TMPDIR docker-compose -f local/localstack.yaml up -d
  echo "Localstack started"
elif [[ $1 == "DOWN" ]]; then
  echo "Stopping localstack..."
  docker-compose -f local/localstack.yaml down
  echo "Localstack stopped"
else
  echo "Usage: ./localstack.sh [UP|DOWN]"
fi
