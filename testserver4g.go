package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// start HTTP server
	log.Print("Starting test server for G")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type request struct {
	Language string `json:"language"`
}

type response struct {
	Ok bool `json:"ok"`
}

type click struct {
	referer  string
	URL      *url.URL
	datetime int64
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request request
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, `{"error_message": "internal error - json request expected"}`, http.StatusBadRequest)
		return
	}

	// TODO process request and push to topic
	log.Printf("host: %v, referer: %v", r.Host, r.Referer())
	log.Printf("language: %v", request.Language)
	log.Printf("headers:")
	for i := range r.Header {
		log.Printf("\t%v => %v", i, r.Header.Get(i))
	}

	log.Printf("body:")

	// initiate kafka connection
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// push to topic
	topic := "incoming_req"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(`"referer": "` + r.Host + `"`),
	}, nil)

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Wait for message deliveries
	p.Flush(15 * 1000)

	// write http response
	payload, err := json.Marshal(response{
		Ok: true,
	})
	if err != nil {
		http.Error(w, `{"error_message": "internal error"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
