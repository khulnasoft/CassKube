#!/bin/bash
all_ids=(`docker container ls | egrep kindest | cut -d ' ' -f 1`)
for id in ${all_ids[*]}
do
   echo "Deleting old operator Docker image from Docker container: $id"
   docker exec $id crictl rmi docker.io/khulnasoft/casskube:latest
done
echo "Loading new operator Docker image into KIND cluster"
kind load docker-image khulnasoft/casskube:latest
echo "Done."
