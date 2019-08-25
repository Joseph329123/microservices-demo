#!/bin/sh

# script to copy the csv from testservice container into local machine

test_service_name=$(kubectl get pods | grep testservice | awk '{print $1}')

while [ ! $(kubectl exec $test_service_name -- ls | grep 'results_unformatted.csv') ]
do
    sleep 10s
done

# copy to local machine
kubectl exec $test_service_name -- cat 'results_unformatted.csv' >> output/results_unformatted.csv
