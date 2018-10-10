# Golan go

From basic HTTP server to gateway for kafka.

## Work

### DONE

- [x] launch the http server
- [x] explore request headers
- [x] launch kafka
- [x] create topic, produce/consume messages
- [x] make the server produces message for each incoming request

### IN PROGRESS

- [ ] package in docker + docker-compose (?)
- [ ] explore request content
- [ ] define a data model to handle the request data

### TODO

- [ ] test on kafka cluster
- [ ] optimize kafka connectio creation
- [ ] check error handling

## Launch kafka

### From source

https://kafka.apache.org/quickstart

### With docker-compose in cluster mode

https://github.com/simplesteph/kafka-stack-docker-compose

```shell
docker-compose -f zk-single-kafka-single.yml up
```

## Run the server

### Installing Confluent's Golang Client for Apache KafkaTM

From https://github.com/confluentinc/confluent-kafka-go

1) Install librdkafka:

``` shell
git clone https://github.com/edenhill/librdkafka.git
cd librdkafka
./configure --prefix /usr
make
sudo make install
```
2) Install go client

```shell
go get -u github.com/confluentinc/confluent-kafka-go/kafka
```

### Run baby run!

```shell
cd golang_helloooo
go build
./golang_helloooo
```

### Launch the consumer

```shell
cd golang_helloooo/cons
go build
./cons
```

## Query the HTTP server

```shell
curl --header "Content-Type: application/json" --request POST --data '{"message":"everything is awesome"}' 'http://localhost:8080/'
```