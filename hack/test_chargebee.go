package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/markns/chargebee-golang/client"
	"github.com/markns/chargebee-golang/client/operations"
)

func main() {
	ctx := context.Background()

	config := client.DefaultTransportConfig().
		WithHost("tenant-test.chargebee.com")
	chargebee := client.NewHTTPClientWithConfig(nil, config)

	basicAuth := httptransport.BasicAuth("test_xxxx", "")

	limit := int32(3)
	customerOK, err := chargebee.Operations.ListCustomer(&operations.ListCustomerParams{
		Context:ctx,
		Limit: &limit,
	}, basicAuth)
	//customerOK, err := chargebee.Operations.RetrieveCustomer(&operations.RetrieveCustomerParams{
	//	Context:ctx,
	//	ID: "aufsdfdfsdh7",
	//}, basicAuth)

	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(customerOK.Payload.List); i++ {
		log.Info(customerOK.Payload.List[i].Customer)
	}
}