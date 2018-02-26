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

// SubscriptionsForCustomerSubscriptionReader is a Reader for the SubscriptionsForCustomerSubscription structure.
type SubscriptionsForCustomerSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubscriptionsForCustomerSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubscriptionsForCustomerSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubscriptionsForCustomerSubscriptionOK creates a SubscriptionsForCustomerSubscriptionOK with default headers values
func NewSubscriptionsForCustomerSubscriptionOK() *SubscriptionsForCustomerSubscriptionOK {
	return &SubscriptionsForCustomerSubscriptionOK{}
}

/*SubscriptionsForCustomerSubscriptionOK handles this case with default header values.

subscriptionsForCustomerSubscription response
*/
type SubscriptionsForCustomerSubscriptionOK struct {
	Payload SubscriptionsForCustomerSubscriptionOKBody
}

func (o *SubscriptionsForCustomerSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /customers/{id}/subscriptions][%d] subscriptionsForCustomerSubscriptionOK  %+v", 200, o.Payload)
}

func (o *SubscriptionsForCustomerSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SubscriptionsForCustomerSubscriptionOKBody subscriptions for customer subscription o k body
swagger:model SubscriptionsForCustomerSubscriptionOKBody
*/

type SubscriptionsForCustomerSubscriptionOKBody struct {

	// list
	// Required: true
	List []*models.SubscriptionResponse `json:"list"`

	// next offset
	// Required: true
	NextOffset *string `json:"next_offset"`
}

/* polymorph SubscriptionsForCustomerSubscriptionOKBody list false */

/* polymorph SubscriptionsForCustomerSubscriptionOKBody next_offset false */

// Validate validates this subscriptions for customer subscription o k body
func (o *SubscriptionsForCustomerSubscriptionOKBody) Validate(formats strfmt.Registry) error {
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

func (o *SubscriptionsForCustomerSubscriptionOKBody) validateList(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionsForCustomerSubscriptionOK"+"."+"list", "body", o.List); err != nil {
		return err
	}

	for i := 0; i < len(o.List); i++ {

		if swag.IsZero(o.List[i]) { // not required
			continue
		}

		if o.List[i] != nil {

			if err := o.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptionsForCustomerSubscriptionOK" + "." + "list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *SubscriptionsForCustomerSubscriptionOKBody) validateNextOffset(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionsForCustomerSubscriptionOK"+"."+"next_offset", "body", o.NextOffset); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SubscriptionsForCustomerSubscriptionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubscriptionsForCustomerSubscriptionOKBody) UnmarshalBinary(b []byte) error {
	var res SubscriptionsForCustomerSubscriptionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
