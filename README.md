# Golan go

From basic HTTP server to gateway for kafka.

## Run the server

TODO:

- install go: https://golang.org/doc/install
- compile
- run

## Launch kafka

### From source

https://kafka.apache.org/quickstart

### With docker-compose

https://github.com/simplesteph/kafka-stack-docker-compose

## Testing Confluent's Golang Client for Apache KafkaTM

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
## Query the HTTP server

```shell
curl --header "Content-Type: application/json" --request POST --data '{"message":"everything is awesome"}' 'http://localhost:8080/'
```