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

// PostVoicemailMessagesReader is a Reader for the PostVoicemailMessages structure.
type PostVoicemailMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostVoicemailMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostVoicemailMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostVoicemailMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostVoicemailMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostVoicemailMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostVoicemailMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostVoicemailMessagesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostVoicemailMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostVoicemailMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostVoicemailMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostVoicemailMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostVoicemailMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostVoicemailMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostVoicemailMessagesOK creates a PostVoicemailMessagesOK with default headers values
func NewPostVoicemailMessagesOK() *PostVoicemailMessagesOK {
	return &PostVoicemailMessagesOK{}
}

/*PostVoicemailMessagesOK handles this case with default header values.

successful operation
*/
type PostVoicemailMessagesOK struct {
	Payload *models.VoicemailMessage
}

func (o *PostVoicemailMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesOK  %+v", 200, o.Payload)
}

func (o *PostVoicemailMessagesOK) GetPayload() *models.VoicemailMessage {
	return o.Payload
}

func (o *PostVoicemailMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesBadRequest creates a PostVoicemailMessagesBadRequest with default headers values
func NewPostVoicemailMessagesBadRequest() *PostVoicemailMessagesBadRequest {
	return &PostVoicemailMessagesBadRequest{}
}

/*PostVoicemailMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostVoicemailMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostVoicemailMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesUnauthorized creates a PostVoicemailMessagesUnauthorized with default headers values
func NewPostVoicemailMessagesUnauthorized() *PostVoicemailMessagesUnauthorized {
	return &PostVoicemailMessagesUnauthorized{}
}

/*PostVoicemailMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostVoicemailMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostVoicemailMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesForbidden creates a PostVoicemailMessagesForbidden with default headers values
func NewPostVoicemailMessagesForbidden() *PostVoicemailMessagesForbidden {
	return &PostVoicemailMessagesForbidden{}
}

/*PostVoicemailMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostVoicemailMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostVoicemailMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesNotFound creates a PostVoicemailMessagesNotFound with default headers values
func NewPostVoicemailMessagesNotFound() *PostVoicemailMessagesNotFound {
	return &PostVoicemailMessagesNotFound{}
}

/*PostVoicemailMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostVoicemailMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostVoicemailMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesConflict creates a PostVoicemailMessagesConflict with default headers values
func NewPostVoicemailMessagesConflict() *PostVoicemailMessagesConflict {
	return &PostVoicemailMessagesConflict{}
}

/*PostVoicemailMessagesConflict handles this case with default header values.

Conflict
*/
type PostVoicemailMessagesConflict struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesConflict  %+v", 409, o.Payload)
}

func (o *PostVoicemailMessagesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesRequestEntityTooLarge creates a PostVoicemailMessagesRequestEntityTooLarge with default headers values
func NewPostVoicemailMessagesRequestEntityTooLarge() *PostVoicemailMessagesRequestEntityTooLarge {
	return &PostVoicemailMessagesRequestEntityTooLarge{}
}

/*PostVoicemailMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostVoicemailMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostVoicemailMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesUnsupportedMediaType creates a PostVoicemailMessagesUnsupportedMediaType with default headers values
func NewPostVoicemailMessagesUnsupportedMediaType() *PostVoicemailMessagesUnsupportedMediaType {
	return &PostVoicemailMessagesUnsupportedMediaType{}
}

/*PostVoicemailMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostVoicemailMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostVoicemailMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesTooManyRequests creates a PostVoicemailMessagesTooManyRequests with default headers values
func NewPostVoicemailMessagesTooManyRequests() *PostVoicemailMessagesTooManyRequests {
	return &PostVoicemailMessagesTooManyRequests{}
}

/*PostVoicemailMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostVoicemailMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostVoicemailMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesInternalServerError creates a PostVoicemailMessagesInternalServerError with default headers values
func NewPostVoicemailMessagesInternalServerError() *PostVoicemailMessagesInternalServerError {
	return &PostVoicemailMessagesInternalServerError{}
}

/*PostVoicemailMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostVoicemailMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostVoicemailMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesServiceUnavailable creates a PostVoicemailMessagesServiceUnavailable with default headers values
func NewPostVoicemailMessagesServiceUnavailable() *PostVoicemailMessagesServiceUnavailable {
	return &PostVoicemailMessagesServiceUnavailable{}
}

/*PostVoicemailMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostVoicemailMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostVoicemailMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostVoicemailMessagesGatewayTimeout creates a PostVoicemailMessagesGatewayTimeout with default headers values
func NewPostVoicemailMessagesGatewayTimeout() *PostVoicemailMessagesGatewayTimeout {
	return &PostVoicemailMessagesGatewayTimeout{}
}

/*PostVoicemailMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostVoicemailMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostVoicemailMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/voicemail/messages][%d] postVoicemailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostVoicemailMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostVoicemailMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
