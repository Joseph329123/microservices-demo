#!/bin/sh
echo "destroy_pods.sh"

cd ..
skaffold delete
sleep 120s

# if there are pods still terminating after 1 min we manually remove them
terminating_pods=($(kubectl get pods | grep Terminating | awk '{print $1}'))

for ((i=0;i<${#terminating_pods[@]};++i))
do
    echo "manually deleting "${terminating_pods[i]}""
    kubectl delete pod ${terminating_pods[i]} --grace-period=0 --force
done

cd experiment_scripts
