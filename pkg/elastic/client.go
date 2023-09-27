package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"os"
	"time"
)

func NewClient() *elasticsearch.Client {
	host := os.Getenv("ELASTIC_HOST")
	if host == "" {
		panic("Elasticsearch host wasn't set")
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			host,
		},
	}

	count := 0
	for {
		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Println("Elasticsearch not ready...")
			count++
		} else {
			log.Println("Connected to Elasticsearch")
			return es
		}

		if count > 10 {
			log.Panic(err)
			return nil
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(time.Second * 2)
	}
}
