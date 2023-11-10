#!/bin/bash

if [ "$(docker ps -q -f name=task-scheduler-mysql-db)" ]; then
  docker rm -f task-scheduler-mysql-db
fi