#!/bin/bash

kafka-topics --create --topic test --partitions 3 --replication-factor 3 --if-not-exists --zookeeper 192.168.99.100:2181
kafka-topics --create --topic test_dlx --partitions 3 --replication-factor 3 --if-not-exists --zookeeper 192.168.99.100:2181
kafka-topics --describe --topic test --zookeeper 192.168.99.100:22181

seq 1 | kafka-console-producer --broker-list 192.168.99.100:32770 --topic test && echo 'Produced 42 messages.'

kafka-console-consumer --bootstrap-server 192.168.99.100:32770 --topic test --from-beginning --max-messages 42
kafka-console-consumer --bootstrap-server 192.168.99.100:32770 --topic test_dlx --from-beginning --max-messages 42



eval $(docker-machine env confluent)