#!/bin/sh

# script to copy the csv (one line) from testservice container into local machine

test_service_name=$(kubectl get pods | grep testservice | awk '{print $1}')

while [ ! $(kubectl exec $test_service_name -- ls | grep 'test_service_measurements.csv') ]
do
    echo "cannot find file: ${pod_with_csv}"
done

# copy to local machine
kubectl exec $test_service_name -- cat 'test_service_measurements.csv' >> output/test_service_measurements.csv
printf "\n" >> output/test_service_measurements.csv
