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

// PutVoicemailUserpolicyReader is a Reader for the PutVoicemailUserpolicy structure.
type PutVoicemailUserpolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVoicemailUserpolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutVoicemailUserpolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutVoicemailUserpolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutVoicemailUserpolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutVoicemailUserpolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutVoicemailUserpolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutVoicemailUserpolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutVoicemailUserpolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutVoicemailUserpolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutVoicemailUserpolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutVoicemailUserpolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutVoicemailUserpolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutVoicemailUserpolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutVoicemailUserpolicyOK creates a PutVoicemailUserpolicyOK with default headers values
func NewPutVoicemailUserpolicyOK() *PutVoicemailUserpolicyOK {
	return &PutVoicemailUserpolicyOK{}
}

/*PutVoicemailUserpolicyOK handles this case with default header values.

successful operation
*/
type PutVoicemailUserpolicyOK struct {
	Payload *models.VoicemailUserPolicy
}

func (o *PutVoicemailUserpolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyOK  %+v", 200, o.Payload)
}

func (o *PutVoicemailUserpolicyOK) GetPayload() *models.VoicemailUserPolicy {
	return o.Payload
}

func (o *PutVoicemailUserpolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailUserPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyBadRequest creates a PutVoicemailUserpolicyBadRequest with default headers values
func NewPutVoicemailUserpolicyBadRequest() *PutVoicemailUserpolicyBadRequest {
	return &PutVoicemailUserpolicyBadRequest{}
}

/*PutVoicemailUserpolicyBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutVoicemailUserpolicyBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutVoicemailUserpolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyUnauthorized creates a PutVoicemailUserpolicyUnauthorized with default headers values
func NewPutVoicemailUserpolicyUnauthorized() *PutVoicemailUserpolicyUnauthorized {
	return &PutVoicemailUserpolicyUnauthorized{}
}

/*PutVoicemailUserpolicyUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutVoicemailUserpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutVoicemailUserpolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyForbidden creates a PutVoicemailUserpolicyForbidden with default headers values
func NewPutVoicemailUserpolicyForbidden() *PutVoicemailUserpolicyForbidden {
	return &PutVoicemailUserpolicyForbidden{}
}

/*PutVoicemailUserpolicyForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutVoicemailUserpolicyForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutVoicemailUserpolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyNotFound creates a PutVoicemailUserpolicyNotFound with default headers values
func NewPutVoicemailUserpolicyNotFound() *PutVoicemailUserpolicyNotFound {
	return &PutVoicemailUserpolicyNotFound{}
}

/*PutVoicemailUserpolicyNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutVoicemailUserpolicyNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutVoicemailUserpolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyRequestTimeout creates a PutVoicemailUserpolicyRequestTimeout with default headers values
func NewPutVoicemailUserpolicyRequestTimeout() *PutVoicemailUserpolicyRequestTimeout {
	return &PutVoicemailUserpolicyRequestTimeout{}
}

/*PutVoicemailUserpolicyRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutVoicemailUserpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutVoicemailUserpolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyRequestEntityTooLarge creates a PutVoicemailUserpolicyRequestEntityTooLarge with default headers values
func NewPutVoicemailUserpolicyRequestEntityTooLarge() *PutVoicemailUserpolicyRequestEntityTooLarge {
	return &PutVoicemailUserpolicyRequestEntityTooLarge{}
}

/*PutVoicemailUserpolicyRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutVoicemailUserpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutVoicemailUserpolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyUnsupportedMediaType creates a PutVoicemailUserpolicyUnsupportedMediaType with default headers values
func NewPutVoicemailUserpolicyUnsupportedMediaType() *PutVoicemailUserpolicyUnsupportedMediaType {
	return &PutVoicemailUserpolicyUnsupportedMediaType{}
}

/*PutVoicemailUserpolicyUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutVoicemailUserpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutVoicemailUserpolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyTooManyRequests creates a PutVoicemailUserpolicyTooManyRequests with default headers values
func NewPutVoicemailUserpolicyTooManyRequests() *PutVoicemailUserpolicyTooManyRequests {
	return &PutVoicemailUserpolicyTooManyRequests{}
}

/*PutVoicemailUserpolicyTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutVoicemailUserpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutVoicemailUserpolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyInternalServerError creates a PutVoicemailUserpolicyInternalServerError with default headers values
func NewPutVoicemailUserpolicyInternalServerError() *PutVoicemailUserpolicyInternalServerError {
	return &PutVoicemailUserpolicyInternalServerError{}
}

/*PutVoicemailUserpolicyInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutVoicemailUserpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutVoicemailUserpolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyServiceUnavailable creates a PutVoicemailUserpolicyServiceUnavailable with default headers values
func NewPutVoicemailUserpolicyServiceUnavailable() *PutVoicemailUserpolicyServiceUnavailable {
	return &PutVoicemailUserpolicyServiceUnavailable{}
}

/*PutVoicemailUserpolicyServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutVoicemailUserpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutVoicemailUserpolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailUserpolicyGatewayTimeout creates a PutVoicemailUserpolicyGatewayTimeout with default headers values
func NewPutVoicemailUserpolicyGatewayTimeout() *PutVoicemailUserpolicyGatewayTimeout {
	return &PutVoicemailUserpolicyGatewayTimeout{}
}

/*PutVoicemailUserpolicyGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutVoicemailUserpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutVoicemailUserpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutVoicemailUserpolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailUserpolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}