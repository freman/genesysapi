// Code generated by go-swagger; DO NOT EDIT.

package voicemail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetVoicemailQueueMessagesReader is a Reader for the GetVoicemailQueueMessages structure.
type GetVoicemailQueueMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailQueueMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailQueueMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailQueueMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailQueueMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailQueueMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailQueueMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailQueueMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailQueueMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailQueueMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailQueueMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailQueueMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailQueueMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailQueueMessagesOK creates a GetVoicemailQueueMessagesOK with default headers values
func NewGetVoicemailQueueMessagesOK() *GetVoicemailQueueMessagesOK {
	return &GetVoicemailQueueMessagesOK{}
}

/*GetVoicemailQueueMessagesOK handles this case with default header values.

successful operation
*/
type GetVoicemailQueueMessagesOK struct {
	Payload *models.VoicemailMessageEntityListing
}

func (o *GetVoicemailQueueMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailQueueMessagesOK) GetPayload() *models.VoicemailMessageEntityListing {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMessageEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesBadRequest creates a GetVoicemailQueueMessagesBadRequest with default headers values
func NewGetVoicemailQueueMessagesBadRequest() *GetVoicemailQueueMessagesBadRequest {
	return &GetVoicemailQueueMessagesBadRequest{}
}

/*GetVoicemailQueueMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailQueueMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailQueueMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesUnauthorized creates a GetVoicemailQueueMessagesUnauthorized with default headers values
func NewGetVoicemailQueueMessagesUnauthorized() *GetVoicemailQueueMessagesUnauthorized {
	return &GetVoicemailQueueMessagesUnauthorized{}
}

/*GetVoicemailQueueMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailQueueMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailQueueMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesForbidden creates a GetVoicemailQueueMessagesForbidden with default headers values
func NewGetVoicemailQueueMessagesForbidden() *GetVoicemailQueueMessagesForbidden {
	return &GetVoicemailQueueMessagesForbidden{}
}

/*GetVoicemailQueueMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailQueueMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailQueueMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesNotFound creates a GetVoicemailQueueMessagesNotFound with default headers values
func NewGetVoicemailQueueMessagesNotFound() *GetVoicemailQueueMessagesNotFound {
	return &GetVoicemailQueueMessagesNotFound{}
}

/*GetVoicemailQueueMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetVoicemailQueueMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailQueueMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesRequestEntityTooLarge creates a GetVoicemailQueueMessagesRequestEntityTooLarge with default headers values
func NewGetVoicemailQueueMessagesRequestEntityTooLarge() *GetVoicemailQueueMessagesRequestEntityTooLarge {
	return &GetVoicemailQueueMessagesRequestEntityTooLarge{}
}

/*GetVoicemailQueueMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetVoicemailQueueMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesUnsupportedMediaType creates a GetVoicemailQueueMessagesUnsupportedMediaType with default headers values
func NewGetVoicemailQueueMessagesUnsupportedMediaType() *GetVoicemailQueueMessagesUnsupportedMediaType {
	return &GetVoicemailQueueMessagesUnsupportedMediaType{}
}

/*GetVoicemailQueueMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailQueueMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesTooManyRequests creates a GetVoicemailQueueMessagesTooManyRequests with default headers values
func NewGetVoicemailQueueMessagesTooManyRequests() *GetVoicemailQueueMessagesTooManyRequests {
	return &GetVoicemailQueueMessagesTooManyRequests{}
}

/*GetVoicemailQueueMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetVoicemailQueueMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailQueueMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesInternalServerError creates a GetVoicemailQueueMessagesInternalServerError with default headers values
func NewGetVoicemailQueueMessagesInternalServerError() *GetVoicemailQueueMessagesInternalServerError {
	return &GetVoicemailQueueMessagesInternalServerError{}
}

/*GetVoicemailQueueMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailQueueMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailQueueMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesServiceUnavailable creates a GetVoicemailQueueMessagesServiceUnavailable with default headers values
func NewGetVoicemailQueueMessagesServiceUnavailable() *GetVoicemailQueueMessagesServiceUnavailable {
	return &GetVoicemailQueueMessagesServiceUnavailable{}
}

/*GetVoicemailQueueMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailQueueMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesGatewayTimeout creates a GetVoicemailQueueMessagesGatewayTimeout with default headers values
func NewGetVoicemailQueueMessagesGatewayTimeout() *GetVoicemailQueueMessagesGatewayTimeout {
	return &GetVoicemailQueueMessagesGatewayTimeout{}
}

/*GetVoicemailQueueMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetVoicemailQueueMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
