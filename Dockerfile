FROM ubuntu:18.04

RUN apt-get update && apt-get install -y git make g++

# Installing Confluent's Golang Client for Apache KafkaTM
RUN \
git clone https://github.com/edenhill/librdkafka.git && \
cd librdkafka && \
./configure --prefix /usr && \
make && \
make install

# Install go client
RUN apt-get install -y golang-go
RUN go get -u github.com/confluentinc/confluent-kafka-go/kafka

# Install golang_helloooo
RUN \
git clone https://github.com/ggdupont/golang_helloooo.git && \
cd golang_helloooo && \
go build && \
./golang_helloooo
