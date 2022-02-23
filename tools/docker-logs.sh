#!/bin/bash
containerName=anekdot-service
dockerID=$(docker ps -a --filter "ancestor=${containerName}" | grep ${containerName} | awk '{print $1}')
docker logs --tail $1 $dockerID
