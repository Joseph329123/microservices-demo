#!/bin/sh


deploy() {
    cd ..

    # https://github.com/GoogleCloudPlatform/microservices-demo
    # skaffold run --default-repo=gcr.io/[PROJECT_ID]
    skaffold run
    
    # IF pods aren't ready in three min (i.e. excessive CrashLoopBackoff), restart deployment)
    end=$((SECONDS+300))
    while [ $flag -eq 1 ] && [ $SECONDS -lt $end ]
    do
        # e.g. READY 1/1 1/1 0/1 .....
        pods_ready=($(kubectl get pods | awk '{print $2}'))
        pods_ready=(${pods_ready[@]/$'READY'})

        # e.g. STATUS Running Running InitContainer Waiting ....
        pods_status=($(kubectl get pods | awk '{print $3}'))
        pods_status=(${pods_status[@]/$'STATUS'})

        for ((i=0;i<${#pods_ready[@]};++i))
        do
            # if some pod is not READY or status is not RUNNING
            if [ "${pods_ready[i]}" != "1/1" ] || [ "${pods_status[i]}" != "Running" ]
            then
                sleep 10s
                flag=1
                break
            fi
            flag=0
        done
    done

    cd experiment_scripts
}

main() {
    local flag=1
    deploy

    # IF pods aren't ready in three min (i.e. excessive CrashLoopBackoff), restart deployment)
    while [ $flag -eq 1 ]
    do
        echo destroying pods and restarting deployment
        ./destroy_pods.sh
        deploy
    done
}

main