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

// GetVoicemailMessagesReader is a Reader for the GetVoicemailMessages structure.
type GetVoicemailMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetVoicemailMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailMessagesOK creates a GetVoicemailMessagesOK with default headers values
func NewGetVoicemailMessagesOK() *GetVoicemailMessagesOK {
	return &GetVoicemailMessagesOK{}
}

/*GetVoicemailMessagesOK handles this case with default header values.

successful operation
*/
type GetVoicemailMessagesOK struct {
	Payload *models.VoicemailMessageEntityListing
}

func (o *GetVoicemailMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailMessagesOK) GetPayload() *models.VoicemailMessageEntityListing {
	return o.Payload
}

func (o *GetVoicemailMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMessageEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesBadRequest creates a GetVoicemailMessagesBadRequest with default headers values
func NewGetVoicemailMessagesBadRequest() *GetVoicemailMessagesBadRequest {
	return &GetVoicemailMessagesBadRequest{}
}

/*GetVoicemailMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesUnauthorized creates a GetVoicemailMessagesUnauthorized with default headers values
func NewGetVoicemailMessagesUnauthorized() *GetVoicemailMessagesUnauthorized {
	return &GetVoicemailMessagesUnauthorized{}
}

/*GetVoicemailMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesForbidden creates a GetVoicemailMessagesForbidden with default headers values
func NewGetVoicemailMessagesForbidden() *GetVoicemailMessagesForbidden {
	return &GetVoicemailMessagesForbidden{}
}

/*GetVoicemailMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesNotFound creates a GetVoicemailMessagesNotFound with default headers values
func NewGetVoicemailMessagesNotFound() *GetVoicemailMessagesNotFound {
	return &GetVoicemailMessagesNotFound{}
}

/*GetVoicemailMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetVoicemailMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesRequestTimeout creates a GetVoicemailMessagesRequestTimeout with default headers values
func NewGetVoicemailMessagesRequestTimeout() *GetVoicemailMessagesRequestTimeout {
	return &GetVoicemailMessagesRequestTimeout{}
}

/*GetVoicemailMessagesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetVoicemailMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesRequestEntityTooLarge creates a GetVoicemailMessagesRequestEntityTooLarge with default headers values
func NewGetVoicemailMessagesRequestEntityTooLarge() *GetVoicemailMessagesRequestEntityTooLarge {
	return &GetVoicemailMessagesRequestEntityTooLarge{}
}

/*GetVoicemailMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetVoicemailMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesUnsupportedMediaType creates a GetVoicemailMessagesUnsupportedMediaType with default headers values
func NewGetVoicemailMessagesUnsupportedMediaType() *GetVoicemailMessagesUnsupportedMediaType {
	return &GetVoicemailMessagesUnsupportedMediaType{}
}

/*GetVoicemailMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesTooManyRequests creates a GetVoicemailMessagesTooManyRequests with default headers values
func NewGetVoicemailMessagesTooManyRequests() *GetVoicemailMessagesTooManyRequests {
	return &GetVoicemailMessagesTooManyRequests{}
}

/*GetVoicemailMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetVoicemailMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesInternalServerError creates a GetVoicemailMessagesInternalServerError with default headers values
func NewGetVoicemailMessagesInternalServerError() *GetVoicemailMessagesInternalServerError {
	return &GetVoicemailMessagesInternalServerError{}
}

/*GetVoicemailMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesServiceUnavailable creates a GetVoicemailMessagesServiceUnavailable with default headers values
func NewGetVoicemailMessagesServiceUnavailable() *GetVoicemailMessagesServiceUnavailable {
	return &GetVoicemailMessagesServiceUnavailable{}
}

/*GetVoicemailMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMessagesGatewayTimeout creates a GetVoicemailMessagesGatewayTimeout with default headers values
func NewGetVoicemailMessagesGatewayTimeout() *GetVoicemailMessagesGatewayTimeout {
	return &GetVoicemailMessagesGatewayTimeout{}
}

/*GetVoicemailMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetVoicemailMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetVoicemailMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/messages][%d] getVoicemailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
