package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"kafka_producer/pkg/kafkadriver"
)

const (
	host  = "localhost"
	port  = "8080"
	topic = "test"
)

var (
	readTimeout = 5 * time.Second
)

func main() {
	producer, notif, producerErr := kafkadriver.CreateProducer(topic)

	go func() {
		for msg := range notif {
			log.Printf("Got msg event: %v\n", msg)
		}
	}()

	go func() {
		select {
		case err := <-producerErr:
			log.Fatalf("got producer error: %v", err)
		}
	}()

	r := chi.NewRouter()
	r.Post("/kafka", func(w http.ResponseWriter, r *http.Request) {
		raw, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var data kafkadriver.KafkaMsg
		if err := json.Unmarshal(raw, &data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		producer <- data
		w.WriteHeader(http.StatusCreated)
	})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
