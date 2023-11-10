#!/bin/bash

if [ "$(docker ps -q -f name=task-scheduler-mongo-db)" ]; then
  docker rm -f task-scheduler-mongo-db
fi