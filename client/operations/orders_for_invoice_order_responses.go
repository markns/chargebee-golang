// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/markns/chargebee-golang/models"
)

// OrdersForInvoiceOrderReader is a Reader for the OrdersForInvoiceOrder structure.
type OrdersForInvoiceOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrdersForInvoiceOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrdersForInvoiceOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrdersForInvoiceOrderOK creates a OrdersForInvoiceOrderOK with default headers values
func NewOrdersForInvoiceOrderOK() *OrdersForInvoiceOrderOK {
	return &OrdersForInvoiceOrderOK{}
}

/*OrdersForInvoiceOrderOK handles this case with default header values.

ordersForInvoiceOrder response
*/
type OrdersForInvoiceOrderOK struct {
	Payload *models.OrdersForInvoiceOrderOKBody
}

func (o *OrdersForInvoiceOrderOK) Error() string {
	return fmt.Sprintf("[GET /invoices/{id}/orders][%d] ordersForInvoiceOrderOK  %+v", 200, o.Payload)
}

func (o *OrdersForInvoiceOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrdersForInvoiceOrderOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
