Mariano Barragán

The solution for this challengue is composed of various applications, which are described in the attached PDF.

To be able to meet the deadline, I wasn't able to complete the docker-compose.yml. This is the WIP version:

```
services:
  kafka:
    container_name: kafka-1
    image: docker.io/bitnami/kafka:3.8
    ports:
      - "9092:9092"
    volumes:
      - "kafka_data:/bitnami"
    environment:
      # KRaft settings
      - KAFKA_CFG_NODE_ID=0
      - KAFKA_CFG_PROCESS_ROLES=controller,broker
      - KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka:9093
      # Listeners
      - KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093
      - KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
      - KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT
      - KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER
      - KAFKA_CFG_INTER_BROKER_LISTENER_NAME=PLAINTEXT
  timeline-service:
    build:
      context: ./timeline-service
    container_name: timeline-service
    ports:
      - "8081:8081"
  user-service:
    build:
      context: ./users-service
    container_name: users-service
    ports:
      - "8083:8083"
  tweets-service:
    depends_on:
      - kafka
    build:
      context: ./tweets-service
    container_name: tweets-service
    ports:
      - "8082:8082"
    restart: unless-stopped
  timeline-subscriber:
    depends_on:
      - kafka
    build:
      context: ./timeline-subscriber
    container_name: timeline-subscriber
    restart: unless-stopped


volumes:
  kafka_data:
    driver: local
```
