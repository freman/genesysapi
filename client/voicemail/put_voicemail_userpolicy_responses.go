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

/*
PutVoicemailUserpolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type PutVoicemailUserpolicyOK struct {
	Payload *models.VoicemailUserPolicy
}

// IsSuccess returns true when this put voicemail userpolicy o k response has a 2xx status code
func (o *PutVoicemailUserpolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put voicemail userpolicy o k response has a 3xx status code
func (o *PutVoicemailUserpolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy o k response has a 4xx status code
func (o *PutVoicemailUserpolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail userpolicy o k response has a 5xx status code
func (o *PutVoicemailUserpolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy o k response a status code equal to that given
func (o *PutVoicemailUserpolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutVoicemailUserpolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyOK  %+v", 200, o.Payload)
}

func (o *PutVoicemailUserpolicyOK) String() string {
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

/*
PutVoicemailUserpolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutVoicemailUserpolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy bad request response has a 2xx status code
func (o *PutVoicemailUserpolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy bad request response has a 3xx status code
func (o *PutVoicemailUserpolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy bad request response has a 4xx status code
func (o *PutVoicemailUserpolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy bad request response has a 5xx status code
func (o *PutVoicemailUserpolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy bad request response a status code equal to that given
func (o *PutVoicemailUserpolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutVoicemailUserpolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutVoicemailUserpolicyBadRequest) String() string {
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

/*
PutVoicemailUserpolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutVoicemailUserpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy unauthorized response has a 2xx status code
func (o *PutVoicemailUserpolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy unauthorized response has a 3xx status code
func (o *PutVoicemailUserpolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy unauthorized response has a 4xx status code
func (o *PutVoicemailUserpolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy unauthorized response has a 5xx status code
func (o *PutVoicemailUserpolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy unauthorized response a status code equal to that given
func (o *PutVoicemailUserpolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutVoicemailUserpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutVoicemailUserpolicyUnauthorized) String() string {
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

/*
PutVoicemailUserpolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutVoicemailUserpolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy forbidden response has a 2xx status code
func (o *PutVoicemailUserpolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy forbidden response has a 3xx status code
func (o *PutVoicemailUserpolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy forbidden response has a 4xx status code
func (o *PutVoicemailUserpolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy forbidden response has a 5xx status code
func (o *PutVoicemailUserpolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy forbidden response a status code equal to that given
func (o *PutVoicemailUserpolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutVoicemailUserpolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutVoicemailUserpolicyForbidden) String() string {
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

/*
PutVoicemailUserpolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutVoicemailUserpolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy not found response has a 2xx status code
func (o *PutVoicemailUserpolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy not found response has a 3xx status code
func (o *PutVoicemailUserpolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy not found response has a 4xx status code
func (o *PutVoicemailUserpolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy not found response has a 5xx status code
func (o *PutVoicemailUserpolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy not found response a status code equal to that given
func (o *PutVoicemailUserpolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutVoicemailUserpolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutVoicemailUserpolicyNotFound) String() string {
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

/*
PutVoicemailUserpolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutVoicemailUserpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy request timeout response has a 2xx status code
func (o *PutVoicemailUserpolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy request timeout response has a 3xx status code
func (o *PutVoicemailUserpolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy request timeout response has a 4xx status code
func (o *PutVoicemailUserpolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy request timeout response has a 5xx status code
func (o *PutVoicemailUserpolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy request timeout response a status code equal to that given
func (o *PutVoicemailUserpolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutVoicemailUserpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutVoicemailUserpolicyRequestTimeout) String() string {
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

/*
PutVoicemailUserpolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutVoicemailUserpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy request entity too large response has a 2xx status code
func (o *PutVoicemailUserpolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy request entity too large response has a 3xx status code
func (o *PutVoicemailUserpolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy request entity too large response has a 4xx status code
func (o *PutVoicemailUserpolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy request entity too large response has a 5xx status code
func (o *PutVoicemailUserpolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy request entity too large response a status code equal to that given
func (o *PutVoicemailUserpolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutVoicemailUserpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutVoicemailUserpolicyRequestEntityTooLarge) String() string {
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

/*
PutVoicemailUserpolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutVoicemailUserpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy unsupported media type response has a 2xx status code
func (o *PutVoicemailUserpolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy unsupported media type response has a 3xx status code
func (o *PutVoicemailUserpolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy unsupported media type response has a 4xx status code
func (o *PutVoicemailUserpolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy unsupported media type response has a 5xx status code
func (o *PutVoicemailUserpolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy unsupported media type response a status code equal to that given
func (o *PutVoicemailUserpolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutVoicemailUserpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutVoicemailUserpolicyUnsupportedMediaType) String() string {
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

/*
PutVoicemailUserpolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutVoicemailUserpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy too many requests response has a 2xx status code
func (o *PutVoicemailUserpolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy too many requests response has a 3xx status code
func (o *PutVoicemailUserpolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy too many requests response has a 4xx status code
func (o *PutVoicemailUserpolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail userpolicy too many requests response has a 5xx status code
func (o *PutVoicemailUserpolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail userpolicy too many requests response a status code equal to that given
func (o *PutVoicemailUserpolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutVoicemailUserpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutVoicemailUserpolicyTooManyRequests) String() string {
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

/*
PutVoicemailUserpolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutVoicemailUserpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy internal server error response has a 2xx status code
func (o *PutVoicemailUserpolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy internal server error response has a 3xx status code
func (o *PutVoicemailUserpolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy internal server error response has a 4xx status code
func (o *PutVoicemailUserpolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail userpolicy internal server error response has a 5xx status code
func (o *PutVoicemailUserpolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put voicemail userpolicy internal server error response a status code equal to that given
func (o *PutVoicemailUserpolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutVoicemailUserpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutVoicemailUserpolicyInternalServerError) String() string {
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

/*
PutVoicemailUserpolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutVoicemailUserpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy service unavailable response has a 2xx status code
func (o *PutVoicemailUserpolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy service unavailable response has a 3xx status code
func (o *PutVoicemailUserpolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy service unavailable response has a 4xx status code
func (o *PutVoicemailUserpolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail userpolicy service unavailable response has a 5xx status code
func (o *PutVoicemailUserpolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put voicemail userpolicy service unavailable response a status code equal to that given
func (o *PutVoicemailUserpolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutVoicemailUserpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutVoicemailUserpolicyServiceUnavailable) String() string {
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

/*
PutVoicemailUserpolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutVoicemailUserpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail userpolicy gateway timeout response has a 2xx status code
func (o *PutVoicemailUserpolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail userpolicy gateway timeout response has a 3xx status code
func (o *PutVoicemailUserpolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail userpolicy gateway timeout response has a 4xx status code
func (o *PutVoicemailUserpolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail userpolicy gateway timeout response has a 5xx status code
func (o *PutVoicemailUserpolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put voicemail userpolicy gateway timeout response a status code equal to that given
func (o *PutVoicemailUserpolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutVoicemailUserpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/userpolicies/{userId}][%d] putVoicemailUserpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutVoicemailUserpolicyGatewayTimeout) String() string {
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
