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

// PutOutboundMessagingcampaignReader is a Reader for the PutOutboundMessagingcampaign structure.
type PutOutboundMessagingcampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundMessagingcampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundMessagingcampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundMessagingcampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundMessagingcampaignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundMessagingcampaignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundMessagingcampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundMessagingcampaignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundMessagingcampaignConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundMessagingcampaignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundMessagingcampaignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundMessagingcampaignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundMessagingcampaignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundMessagingcampaignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundMessagingcampaignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundMessagingcampaignOK creates a PutOutboundMessagingcampaignOK with default headers values
func NewPutOutboundMessagingcampaignOK() *PutOutboundMessagingcampaignOK {
	return &PutOutboundMessagingcampaignOK{}
}

/*PutOutboundMessagingcampaignOK handles this case with default header values.

successful operation
*/
type PutOutboundMessagingcampaignOK struct {
	Payload *models.MessagingCampaign
}

func (o *PutOutboundMessagingcampaignOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignOK  %+v", 200, o.Payload)
}

func (o *PutOutboundMessagingcampaignOK) GetPayload() *models.MessagingCampaign {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessagingCampaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignBadRequest creates a PutOutboundMessagingcampaignBadRequest with default headers values
func NewPutOutboundMessagingcampaignBadRequest() *PutOutboundMessagingcampaignBadRequest {
	return &PutOutboundMessagingcampaignBadRequest{}
}

/*PutOutboundMessagingcampaignBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundMessagingcampaignBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundMessagingcampaignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignUnauthorized creates a PutOutboundMessagingcampaignUnauthorized with default headers values
func NewPutOutboundMessagingcampaignUnauthorized() *PutOutboundMessagingcampaignUnauthorized {
	return &PutOutboundMessagingcampaignUnauthorized{}
}

/*PutOutboundMessagingcampaignUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundMessagingcampaignUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundMessagingcampaignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignForbidden creates a PutOutboundMessagingcampaignForbidden with default headers values
func NewPutOutboundMessagingcampaignForbidden() *PutOutboundMessagingcampaignForbidden {
	return &PutOutboundMessagingcampaignForbidden{}
}

/*PutOutboundMessagingcampaignForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundMessagingcampaignForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundMessagingcampaignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignNotFound creates a PutOutboundMessagingcampaignNotFound with default headers values
func NewPutOutboundMessagingcampaignNotFound() *PutOutboundMessagingcampaignNotFound {
	return &PutOutboundMessagingcampaignNotFound{}
}

/*PutOutboundMessagingcampaignNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOutboundMessagingcampaignNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundMessagingcampaignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignRequestTimeout creates a PutOutboundMessagingcampaignRequestTimeout with default headers values
func NewPutOutboundMessagingcampaignRequestTimeout() *PutOutboundMessagingcampaignRequestTimeout {
	return &PutOutboundMessagingcampaignRequestTimeout{}
}

/*PutOutboundMessagingcampaignRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundMessagingcampaignRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundMessagingcampaignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignConflict creates a PutOutboundMessagingcampaignConflict with default headers values
func NewPutOutboundMessagingcampaignConflict() *PutOutboundMessagingcampaignConflict {
	return &PutOutboundMessagingcampaignConflict{}
}

/*PutOutboundMessagingcampaignConflict handles this case with default header values.

Conflict.
*/
type PutOutboundMessagingcampaignConflict struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundMessagingcampaignConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignRequestEntityTooLarge creates a PutOutboundMessagingcampaignRequestEntityTooLarge with default headers values
func NewPutOutboundMessagingcampaignRequestEntityTooLarge() *PutOutboundMessagingcampaignRequestEntityTooLarge {
	return &PutOutboundMessagingcampaignRequestEntityTooLarge{}
}

/*PutOutboundMessagingcampaignRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutOutboundMessagingcampaignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundMessagingcampaignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignUnsupportedMediaType creates a PutOutboundMessagingcampaignUnsupportedMediaType with default headers values
func NewPutOutboundMessagingcampaignUnsupportedMediaType() *PutOutboundMessagingcampaignUnsupportedMediaType {
	return &PutOutboundMessagingcampaignUnsupportedMediaType{}
}

/*PutOutboundMessagingcampaignUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundMessagingcampaignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundMessagingcampaignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignTooManyRequests creates a PutOutboundMessagingcampaignTooManyRequests with default headers values
func NewPutOutboundMessagingcampaignTooManyRequests() *PutOutboundMessagingcampaignTooManyRequests {
	return &PutOutboundMessagingcampaignTooManyRequests{}
}

/*PutOutboundMessagingcampaignTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundMessagingcampaignTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundMessagingcampaignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignInternalServerError creates a PutOutboundMessagingcampaignInternalServerError with default headers values
func NewPutOutboundMessagingcampaignInternalServerError() *PutOutboundMessagingcampaignInternalServerError {
	return &PutOutboundMessagingcampaignInternalServerError{}
}

/*PutOutboundMessagingcampaignInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundMessagingcampaignInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundMessagingcampaignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignServiceUnavailable creates a PutOutboundMessagingcampaignServiceUnavailable with default headers values
func NewPutOutboundMessagingcampaignServiceUnavailable() *PutOutboundMessagingcampaignServiceUnavailable {
	return &PutOutboundMessagingcampaignServiceUnavailable{}
}

/*PutOutboundMessagingcampaignServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundMessagingcampaignServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundMessagingcampaignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundMessagingcampaignGatewayTimeout creates a PutOutboundMessagingcampaignGatewayTimeout with default headers values
func NewPutOutboundMessagingcampaignGatewayTimeout() *PutOutboundMessagingcampaignGatewayTimeout {
	return &PutOutboundMessagingcampaignGatewayTimeout{}
}

/*PutOutboundMessagingcampaignGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOutboundMessagingcampaignGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundMessagingcampaignGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/messagingcampaigns/{messagingCampaignId}][%d] putOutboundMessagingcampaignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundMessagingcampaignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundMessagingcampaignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
