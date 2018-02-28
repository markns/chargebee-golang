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

// CollectNowHostedPageReader is a Reader for the CollectNowHostedPage structure.
type CollectNowHostedPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectNowHostedPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCollectNowHostedPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCollectNowHostedPageOK creates a CollectNowHostedPageOK with default headers values
func NewCollectNowHostedPageOK() *CollectNowHostedPageOK {
	return &CollectNowHostedPageOK{}
}

/*CollectNowHostedPageOK handles this case with default header values.

collectNowHostedPage response
*/
type CollectNowHostedPageOK struct {
	Payload *models.HostedPageResponse
}

func (o *CollectNowHostedPageOK) Error() string {
	return fmt.Sprintf("[POST /hosted_pages/collect_now][%d] collectNowHostedPageOK  %+v", 200, o.Payload)
}

func (o *CollectNowHostedPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostedPageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
