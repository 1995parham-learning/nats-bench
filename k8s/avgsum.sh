#!/bin/bash

total_pub=0
total_sub=0

for pod in $(kubectl get pods -l app=nats-bench -o jsonpath="{.items[*].metadata.name}"); do
        echo "Log for $pod"
        log=$(kubectl logs $pod | grep "avg")

        pub=$(echo $log | cut -d '|' -f 2 - | cut -d ' ' -f 3 -)
        sub=$(echo $log | cut -d '|' -f 5 - | cut -d ' ' -f 3 -)

        pub=${pub/,/}
        sub=${sub/,/}

        total_pub=$((total_pub + pub))
        total_sub=$((total_sub + sub))

        echo pub=$total_pub
        echo sub=$total_sub
done

echo "---"
echo pub=$total_pub
echo sub=$total_sub
