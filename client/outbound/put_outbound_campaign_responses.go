// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOutboundCampaignReader is a Reader for the PutOutboundCampaign structure.
type PutOutboundCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundCampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundCampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundCampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundCampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundCampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundCampaignConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundCampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundCampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundCampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundCampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundCampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundCampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundCampaignOK creates a PutOutboundCampaignOK with default headers values
func NewPutOutboundCampaignOK() *PutOutboundCampaignOK {
	return &PutOutboundCampaignOK{}
}

/*PutOutboundCampaignOK handles this case with default header values.

successful operation
*/
type PutOutboundCampaignOK struct {
	Payload *models.Campaign
}

func (o *PutOutboundCampaignOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignOK  %+v", 200, o.Payload)
}

func (o *PutOutboundCampaignOK) GetPayload() *models.Campaign {
	return o.Payload
}

func (o *PutOutboundCampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Campaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignBadRequest creates a PutOutboundCampaignBadRequest with default headers values
func NewPutOutboundCampaignBadRequest() *PutOutboundCampaignBadRequest {
	return &PutOutboundCampaignBadRequest{}
}

/*PutOutboundCampaignBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundCampaignBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundCampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignUnauthorized creates a PutOutboundCampaignUnauthorized with default headers values
func NewPutOutboundCampaignUnauthorized() *PutOutboundCampaignUnauthorized {
	return &PutOutboundCampaignUnauthorized{}
}

/*PutOutboundCampaignUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundCampaignUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundCampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignForbidden creates a PutOutboundCampaignForbidden with default headers values
func NewPutOutboundCampaignForbidden() *PutOutboundCampaignForbidden {
	return &PutOutboundCampaignForbidden{}
}

/*PutOutboundCampaignForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundCampaignForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundCampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignNotFound creates a PutOutboundCampaignNotFound with default headers values
func NewPutOutboundCampaignNotFound() *PutOutboundCampaignNotFound {
	return &PutOutboundCampaignNotFound{}
}

/*PutOutboundCampaignNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOutboundCampaignNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundCampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignConflict creates a PutOutboundCampaignConflict with default headers values
func NewPutOutboundCampaignConflict() *PutOutboundCampaignConflict {
	return &PutOutboundCampaignConflict{}
}

/*PutOutboundCampaignConflict handles this case with default header values.

Conflict.
*/
type PutOutboundCampaignConflict struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundCampaignConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignRequestEntityTooLarge creates a PutOutboundCampaignRequestEntityTooLarge with default headers values
func NewPutOutboundCampaignRequestEntityTooLarge() *PutOutboundCampaignRequestEntityTooLarge {
	return &PutOutboundCampaignRequestEntityTooLarge{}
}

/*PutOutboundCampaignRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOutboundCampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundCampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignUnsupportedMediaType creates a PutOutboundCampaignUnsupportedMediaType with default headers values
func NewPutOutboundCampaignUnsupportedMediaType() *PutOutboundCampaignUnsupportedMediaType {
	return &PutOutboundCampaignUnsupportedMediaType{}
}

/*PutOutboundCampaignUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundCampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundCampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignTooManyRequests creates a PutOutboundCampaignTooManyRequests with default headers values
func NewPutOutboundCampaignTooManyRequests() *PutOutboundCampaignTooManyRequests {
	return &PutOutboundCampaignTooManyRequests{}
}

/*PutOutboundCampaignTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutOutboundCampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundCampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignInternalServerError creates a PutOutboundCampaignInternalServerError with default headers values
func NewPutOutboundCampaignInternalServerError() *PutOutboundCampaignInternalServerError {
	return &PutOutboundCampaignInternalServerError{}
}

/*PutOutboundCampaignInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundCampaignInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundCampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignServiceUnavailable creates a PutOutboundCampaignServiceUnavailable with default headers values
func NewPutOutboundCampaignServiceUnavailable() *PutOutboundCampaignServiceUnavailable {
	return &PutOutboundCampaignServiceUnavailable{}
}

/*PutOutboundCampaignServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundCampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundCampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCampaignGatewayTimeout creates a PutOutboundCampaignGatewayTimeout with default headers values
func NewPutOutboundCampaignGatewayTimeout() *PutOutboundCampaignGatewayTimeout {
	return &PutOutboundCampaignGatewayTimeout{}
}

/*PutOutboundCampaignGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOutboundCampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/campaigns/{campaignId}][%d] putOutboundCampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundCampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
