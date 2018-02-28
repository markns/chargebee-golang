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

// ListSiteMigrationDetailReader is a Reader for the ListSiteMigrationDetail structure.
type ListSiteMigrationDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSiteMigrationDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSiteMigrationDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSiteMigrationDetailOK creates a ListSiteMigrationDetailOK with default headers values
func NewListSiteMigrationDetailOK() *ListSiteMigrationDetailOK {
	return &ListSiteMigrationDetailOK{}
}

/*ListSiteMigrationDetailOK handles this case with default header values.

listSiteMigrationDetail response
*/
type ListSiteMigrationDetailOK struct {
	Payload *models.ListSiteMigrationDetailOKBody
}

func (o *ListSiteMigrationDetailOK) Error() string {
	return fmt.Sprintf("[GET /site_migration_details][%d] listSiteMigrationDetailOK  %+v", 200, o.Payload)
}

func (o *ListSiteMigrationDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSiteMigrationDetailOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
