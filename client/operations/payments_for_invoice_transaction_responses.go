// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/markns/chargebee-golang/models"
)

// PaymentsForInvoiceTransactionReader is a Reader for the PaymentsForInvoiceTransaction structure.
type PaymentsForInvoiceTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentsForInvoiceTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPaymentsForInvoiceTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPaymentsForInvoiceTransactionOK creates a PaymentsForInvoiceTransactionOK with default headers values
func NewPaymentsForInvoiceTransactionOK() *PaymentsForInvoiceTransactionOK {
	return &PaymentsForInvoiceTransactionOK{}
}

/*PaymentsForInvoiceTransactionOK handles this case with default header values.

paymentsForInvoiceTransaction response
*/
type PaymentsForInvoiceTransactionOK struct {
	Payload PaymentsForInvoiceTransactionOKBody
}

func (o *PaymentsForInvoiceTransactionOK) Error() string {
	return fmt.Sprintf("[GET /invoices/{id}/payments][%d] paymentsForInvoiceTransactionOK  %+v", 200, o.Payload)
}

func (o *PaymentsForInvoiceTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PaymentsForInvoiceTransactionOKBody payments for invoice transaction o k body
swagger:model PaymentsForInvoiceTransactionOKBody
*/

type PaymentsForInvoiceTransactionOKBody struct {

	// list
	// Required: true
	List []*models.TransactionResponse `json:"list"`

	// next offset
	// Required: true
	NextOffset *string `json:"next_offset"`
}

/* polymorph PaymentsForInvoiceTransactionOKBody list false */

/* polymorph PaymentsForInvoiceTransactionOKBody next_offset false */

// Validate validates this payments for invoice transaction o k body
func (o *PaymentsForInvoiceTransactionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateList(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateNextOffset(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PaymentsForInvoiceTransactionOKBody) validateList(formats strfmt.Registry) error {

	if err := validate.Required("paymentsForInvoiceTransactionOK"+"."+"list", "body", o.List); err != nil {
		return err
	}

	for i := 0; i < len(o.List); i++ {

		if swag.IsZero(o.List[i]) { // not required
			continue
		}

		if o.List[i] != nil {

			if err := o.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paymentsForInvoiceTransactionOK" + "." + "list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PaymentsForInvoiceTransactionOKBody) validateNextOffset(formats strfmt.Registry) error {

	if err := validate.Required("paymentsForInvoiceTransactionOK"+"."+"next_offset", "body", o.NextOffset); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PaymentsForInvoiceTransactionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PaymentsForInvoiceTransactionOKBody) UnmarshalBinary(b []byte) error {
	var res PaymentsForInvoiceTransactionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
