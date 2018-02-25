// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateOrder create order API
*/
func (a *Client) CreateOrder(params *CreateOrderParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrder",
		Method:             "POST",
		PathPattern:        "/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOrderOK), nil

}

/*
ListOrder list order API
*/
func (a *Client) ListOrder(params *ListOrderParams, authInfo runtime.ClientAuthInfoWriter) (*ListOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listOrder",
		Method:             "GET",
		PathPattern:        "/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListOrderOK), nil

}

/*
OrdersForInvoiceOrder orders for invoice order API
*/
func (a *Client) OrdersForInvoiceOrder(params *OrdersForInvoiceOrderParams, authInfo runtime.ClientAuthInfoWriter) (*OrdersForInvoiceOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrdersForInvoiceOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ordersForInvoiceOrder",
		Method:             "GET",
		PathPattern:        "/invoices/{id}/orders",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OrdersForInvoiceOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OrdersForInvoiceOrderOK), nil

}

/*
RetrieveOrder retrieve order API
*/
func (a *Client) RetrieveOrder(params *RetrieveOrderParams, authInfo runtime.ClientAuthInfoWriter) (*RetrieveOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveOrder",
		Method:             "GET",
		PathPattern:        "/orders/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetrieveOrderOK), nil

}

/*
UpdateOrder update order API
*/
func (a *Client) UpdateOrder(params *UpdateOrderParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrder",
		Method:             "POST",
		PathPattern:        "/orders/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateOrderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateOrderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}