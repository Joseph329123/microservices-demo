#!/bin/sh

echo "starting..."
cd ..
skaffold run --default-repo=gcr.io/liquid-champion-244004
sleep 5s
echo "printing names"
for p in $(kubectl get pods | awk '{print $1}')
do
    echo $p
done
# sleep 60s
# kubectl get pods
# sleep 5s
# skaffold delete
# sleep 60s
# #for p in $(kubectl get pods | grep Terminating | awk '{print $1}'); do kubectl delete pod $p --grace-period=0 --force;done
# kubectl get pods
# echo "done"
# kubectl exec frontend-5d9b6fcd8-h8w2m
# kubectl delete pod <pod>
