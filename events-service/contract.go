package events-service

import (
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

type events-servicecontract interface {
	HelloWorld() (string, error)
}

type events-service struct{}

// AddServices sets up and starts the service.
func AddServices() {
	ctx := context.Background()
	svc := events-service{}

	events-serviceHandler := httptransport.NewServer(
		ctx,
		makeHelloWorldEndpoint(svc),
		decodeHelloWorldRequest,
		encodeResponse,
	)

	http.Handle("/", events-serviceHandler)
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
