
Source code
---

- https://github.com/sceneryback/kafka-best-practices/
- refs: https://medium.com/swlh/how-to-consume-kafka-efficiently-in-golang-264f7fe2155b


Test
---

- Run `docker-compose up` to start kafka (9092), zoo, kafdrop / web-ui (9000)

- Run test `./pkg/comsumer/consumer_test.go`
  + `TestProducer`: push msg to topic
  + `TestMultiBatchConsumer`: read msg from topic