#!/bin/sh

# deploy pods
./deploy_pods


# LOOP 

# delete loadgenerator service (so it will re-pull latest version from GCR)
loadgen_service_name=$(kubectl get pods | grep loadgen | awk '{print $1}')
kubectl delete pod $loadgen_service_name

# wait for loadgen pod to be ready READY 1/1 so testservice
# is taking measurements with simulated users
loadgen_ready=$(kubectl get pods | grep loadgen | awk '{print $2}')
loadgen_status=$(kubectl get pods | grep loadgen | awk '{print $3}')
while [ "${loadgen_ready}" != "1/1" ] || [ "${loadgen_status}" != "Running" ]
do
	sleep 10s
	loadgen_ready=$(kubectl get pods | grep loadgen | awk '{print $2}')
	loadgen_status=$(kubectl get pods | grep loadgen | awk '{print $3}')
done

# delete testservice (so it will re-pull latest version from GCR and take new measurements)
testservice_ready=$(kubectl get pods | grep testservice | awk '{print $2}')
testservice_status=$(kubectl get pods | grep testservice | awk '{print $3}')
while [ "${testservice_ready}" != "1/1" ] || [ "${testservice_status}" != "Running" ]
do
	sleep 10s
	testservice_ready=$(kubectl get pods | grep testservice | awk '{print $2}')
	testservice_status=$(kubectl get pods | grep testservice | awk '{print $3}')
done

# copy csv from testservice container to local machine
./copy_csv_testservice.sh

# copy avg req/ sec

