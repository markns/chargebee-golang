package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/markns/chargebee-golang/client"
	"github.com/markns/chargebee-golang/client/operations"
)

// This function is a little bit mental unfortunately. The Chargebee API accepts form urlencoded
// params eg. email=foo%40bar.com&id=test (and produces JSON... but whatever). We can't simply
// add our request structs url.Values{}, call encode and be done with it, because the field names
// should be in snake_case. So, the solution is to serialize to JSON, deserialize, to a map[string]string,
// and finally encode to the form. Joy.
func URLEncodedProducer() runtime.Producer {
	return runtime.ProducerFunc(func(writer io.Writer, data interface{}) error {
		v := reflect.Indirect(reflect.ValueOf(data))
		if t := v.Type(); t.Kind() == reflect.Struct || t.Kind() == reflect.Slice {
			b, err := swag.WriteJSON(data)
			if err != nil {
				return err
			}

			var jsonMap map[string]interface{}
			err = json.Unmarshal(b, &jsonMap)
			if err != nil {
				return err
			}

			form := url.Values{}
			for k, v := range jsonMap {
				form.Add(k, fmt.Sprintf("%v", v))
			}

			_, err = writer.Write([]byte(form.Encode()))
			return err
		}
		return nil
	})
}


func main() {
	ctx := context.Background()

	config := client.DefaultTransportConfig().WithHost("tenant-test.chargebee.com")
	transport := openapi.New(config.Host, config.BasePath, config.Schemes)
	transport.Producers["application/x-www-form-urlencoded"] = URLEncodedProducer()
	chargebee := client.New(transport, nil)
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