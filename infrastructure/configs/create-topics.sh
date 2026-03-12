#!/bin/bash
# Создание топиков при старте Kafka (Apache официальный образ)


BOOTSTRAP="kafka:9092"
TOPICS=("match-events")

while ! /opt/kafka/bin/kafka-topics.sh --bootstrap-server $BOOTSTRAP --list >/dev/null 2>&1
do
  echo "Kafka not ready yet..."
  sleep 3
done

for TOPIC in "${TOPICS[@]}"; do
  /opt/kafka/bin/kafka-topics.sh --create --topic $TOPIC --bootstrap-server $BOOTSTRAP --partitions 3 --replication-factor 1
done

echo "Kafka topics created: ${TOPICS[*]}"