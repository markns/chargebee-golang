// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/markns/chargebee-golang/models"
)

// TransactionsForCustomerTransactionReader is a Reader for the TransactionsForCustomerTransaction structure.
type TransactionsForCustomerTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransactionsForCustomerTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTransactionsForCustomerTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTransactionsForCustomerTransactionOK creates a TransactionsForCustomerTransactionOK with default headers values
func NewTransactionsForCustomerTransactionOK() *TransactionsForCustomerTransactionOK {
	return &TransactionsForCustomerTransactionOK{}
}

/*TransactionsForCustomerTransactionOK handles this case with default header values.

transactionsForCustomerTransaction response
*/
type TransactionsForCustomerTransactionOK struct {
	Payload *models.TransactionResponse
}

func (o *TransactionsForCustomerTransactionOK) Error() string {
	return fmt.Sprintf("[GET /customers/{id}/transactions][%d] transactionsForCustomerTransactionOK  %+v", 200, o.Payload)
}

func (o *TransactionsForCustomerTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
