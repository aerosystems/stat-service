package elastic

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	numberConnectionAttempts = 10
	caCertPath               = "/app/certs/es01.crt"
)

func NewClient() *elasticsearch.Client {
	host := os.Getenv("ELASTIC_HOST")
	if host == "" {
		log.Panic("Elasticsearch host wasn't set")
	}
	password := os.Getenv("ELASTIC_PASSWORD")
	if password == "" {
		log.Panic("Elasticsearch password wasn't set")
	}

	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		log.Panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := elasticsearch.Config{
		Addresses: []string{
			host,
		},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	auth := base64.StdEncoding.EncodeToString([]byte("elastic:" + password))
	cfg.Header = http.Header{}
	cfg.Header.Set("Authorization", "Basic "+auth)

	count := 0
	for {
		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			log.Panic(err)
		}
		_, err = es.Info()
		if err != nil {
			log.Println("Elasticsearch not ready...")
			count++
		} else {
			log.Println("Connected to Elasticsearch")
			return es
		}
		log.Println("Backing off for two seconds...")
		time.Sleep(time.Second * 2)
		if count == numberConnectionAttempts {
			log.Panic(err)
			return nil
		}
	}
}
