#!/bin/sh

for ((i=0;i<10;++i))
do
    printf "\n\n"
    echo "--------------------------------------------------------"
    echo "Iteration: $i"
    ./orchestrate.sh
    echo "--------------------------------------------------------"
done

printf "\n"
