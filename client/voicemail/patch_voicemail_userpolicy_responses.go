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

// PatchVoicemailUserpolicyReader is a Reader for the PatchVoicemailUserpolicy structure.
type PatchVoicemailUserpolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVoicemailUserpolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchVoicemailUserpolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchVoicemailUserpolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchVoicemailUserpolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchVoicemailUserpolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchVoicemailUserpolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchVoicemailUserpolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchVoicemailUserpolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchVoicemailUserpolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchVoicemailUserpolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchVoicemailUserpolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchVoicemailUserpolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchVoicemailUserpolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchVoicemailUserpolicyOK creates a PatchVoicemailUserpolicyOK with default headers values
func NewPatchVoicemailUserpolicyOK() *PatchVoicemailUserpolicyOK {
	return &PatchVoicemailUserpolicyOK{}
}

/*
PatchVoicemailUserpolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchVoicemailUserpolicyOK struct {
	Payload *models.VoicemailUserPolicy
}

// IsSuccess returns true when this patch voicemail userpolicy o k response has a 2xx status code
func (o *PatchVoicemailUserpolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch voicemail userpolicy o k response has a 3xx status code
func (o *PatchVoicemailUserpolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy o k response has a 4xx status code
func (o *PatchVoicemailUserpolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail userpolicy o k response has a 5xx status code
func (o *PatchVoicemailUserpolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy o k response a status code equal to that given
func (o *PatchVoicemailUserpolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchVoicemailUserpolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyOK  %+v", 200, o.Payload)
}

func (o *PatchVoicemailUserpolicyOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyOK  %+v", 200, o.Payload)
}

func (o *PatchVoicemailUserpolicyOK) GetPayload() *models.VoicemailUserPolicy {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailUserPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyBadRequest creates a PatchVoicemailUserpolicyBadRequest with default headers values
func NewPatchVoicemailUserpolicyBadRequest() *PatchVoicemailUserpolicyBadRequest {
	return &PatchVoicemailUserpolicyBadRequest{}
}

/*
PatchVoicemailUserpolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchVoicemailUserpolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy bad request response has a 2xx status code
func (o *PatchVoicemailUserpolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy bad request response has a 3xx status code
func (o *PatchVoicemailUserpolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy bad request response has a 4xx status code
func (o *PatchVoicemailUserpolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy bad request response has a 5xx status code
func (o *PatchVoicemailUserpolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy bad request response a status code equal to that given
func (o *PatchVoicemailUserpolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchVoicemailUserpolicyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVoicemailUserpolicyBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVoicemailUserpolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyUnauthorized creates a PatchVoicemailUserpolicyUnauthorized with default headers values
func NewPatchVoicemailUserpolicyUnauthorized() *PatchVoicemailUserpolicyUnauthorized {
	return &PatchVoicemailUserpolicyUnauthorized{}
}

/*
PatchVoicemailUserpolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchVoicemailUserpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy unauthorized response has a 2xx status code
func (o *PatchVoicemailUserpolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy unauthorized response has a 3xx status code
func (o *PatchVoicemailUserpolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy unauthorized response has a 4xx status code
func (o *PatchVoicemailUserpolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy unauthorized response has a 5xx status code
func (o *PatchVoicemailUserpolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy unauthorized response a status code equal to that given
func (o *PatchVoicemailUserpolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchVoicemailUserpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchVoicemailUserpolicyUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchVoicemailUserpolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyForbidden creates a PatchVoicemailUserpolicyForbidden with default headers values
func NewPatchVoicemailUserpolicyForbidden() *PatchVoicemailUserpolicyForbidden {
	return &PatchVoicemailUserpolicyForbidden{}
}

/*
PatchVoicemailUserpolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchVoicemailUserpolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy forbidden response has a 2xx status code
func (o *PatchVoicemailUserpolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy forbidden response has a 3xx status code
func (o *PatchVoicemailUserpolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy forbidden response has a 4xx status code
func (o *PatchVoicemailUserpolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy forbidden response has a 5xx status code
func (o *PatchVoicemailUserpolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy forbidden response a status code equal to that given
func (o *PatchVoicemailUserpolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchVoicemailUserpolicyForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PatchVoicemailUserpolicyForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PatchVoicemailUserpolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyNotFound creates a PatchVoicemailUserpolicyNotFound with default headers values
func NewPatchVoicemailUserpolicyNotFound() *PatchVoicemailUserpolicyNotFound {
	return &PatchVoicemailUserpolicyNotFound{}
}

/*
PatchVoicemailUserpolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchVoicemailUserpolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy not found response has a 2xx status code
func (o *PatchVoicemailUserpolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy not found response has a 3xx status code
func (o *PatchVoicemailUserpolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy not found response has a 4xx status code
func (o *PatchVoicemailUserpolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy not found response has a 5xx status code
func (o *PatchVoicemailUserpolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy not found response a status code equal to that given
func (o *PatchVoicemailUserpolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchVoicemailUserpolicyNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PatchVoicemailUserpolicyNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PatchVoicemailUserpolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyRequestTimeout creates a PatchVoicemailUserpolicyRequestTimeout with default headers values
func NewPatchVoicemailUserpolicyRequestTimeout() *PatchVoicemailUserpolicyRequestTimeout {
	return &PatchVoicemailUserpolicyRequestTimeout{}
}

/*
PatchVoicemailUserpolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchVoicemailUserpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy request timeout response has a 2xx status code
func (o *PatchVoicemailUserpolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy request timeout response has a 3xx status code
func (o *PatchVoicemailUserpolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy request timeout response has a 4xx status code
func (o *PatchVoicemailUserpolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy request timeout response has a 5xx status code
func (o *PatchVoicemailUserpolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy request timeout response a status code equal to that given
func (o *PatchVoicemailUserpolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchVoicemailUserpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchVoicemailUserpolicyRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchVoicemailUserpolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyRequestEntityTooLarge creates a PatchVoicemailUserpolicyRequestEntityTooLarge with default headers values
func NewPatchVoicemailUserpolicyRequestEntityTooLarge() *PatchVoicemailUserpolicyRequestEntityTooLarge {
	return &PatchVoicemailUserpolicyRequestEntityTooLarge{}
}

/*
PatchVoicemailUserpolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchVoicemailUserpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy request entity too large response has a 2xx status code
func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy request entity too large response has a 3xx status code
func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy request entity too large response has a 4xx status code
func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy request entity too large response has a 5xx status code
func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy request entity too large response a status code equal to that given
func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyUnsupportedMediaType creates a PatchVoicemailUserpolicyUnsupportedMediaType with default headers values
func NewPatchVoicemailUserpolicyUnsupportedMediaType() *PatchVoicemailUserpolicyUnsupportedMediaType {
	return &PatchVoicemailUserpolicyUnsupportedMediaType{}
}

/*
PatchVoicemailUserpolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchVoicemailUserpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy unsupported media type response has a 2xx status code
func (o *PatchVoicemailUserpolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy unsupported media type response has a 3xx status code
func (o *PatchVoicemailUserpolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy unsupported media type response has a 4xx status code
func (o *PatchVoicemailUserpolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy unsupported media type response has a 5xx status code
func (o *PatchVoicemailUserpolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy unsupported media type response a status code equal to that given
func (o *PatchVoicemailUserpolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchVoicemailUserpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchVoicemailUserpolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchVoicemailUserpolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyTooManyRequests creates a PatchVoicemailUserpolicyTooManyRequests with default headers values
func NewPatchVoicemailUserpolicyTooManyRequests() *PatchVoicemailUserpolicyTooManyRequests {
	return &PatchVoicemailUserpolicyTooManyRequests{}
}

/*
PatchVoicemailUserpolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchVoicemailUserpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy too many requests response has a 2xx status code
func (o *PatchVoicemailUserpolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy too many requests response has a 3xx status code
func (o *PatchVoicemailUserpolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy too many requests response has a 4xx status code
func (o *PatchVoicemailUserpolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail userpolicy too many requests response has a 5xx status code
func (o *PatchVoicemailUserpolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail userpolicy too many requests response a status code equal to that given
func (o *PatchVoicemailUserpolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchVoicemailUserpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchVoicemailUserpolicyTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchVoicemailUserpolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyInternalServerError creates a PatchVoicemailUserpolicyInternalServerError with default headers values
func NewPatchVoicemailUserpolicyInternalServerError() *PatchVoicemailUserpolicyInternalServerError {
	return &PatchVoicemailUserpolicyInternalServerError{}
}

/*
PatchVoicemailUserpolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchVoicemailUserpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy internal server error response has a 2xx status code
func (o *PatchVoicemailUserpolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy internal server error response has a 3xx status code
func (o *PatchVoicemailUserpolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy internal server error response has a 4xx status code
func (o *PatchVoicemailUserpolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail userpolicy internal server error response has a 5xx status code
func (o *PatchVoicemailUserpolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch voicemail userpolicy internal server error response a status code equal to that given
func (o *PatchVoicemailUserpolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchVoicemailUserpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchVoicemailUserpolicyInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchVoicemailUserpolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyServiceUnavailable creates a PatchVoicemailUserpolicyServiceUnavailable with default headers values
func NewPatchVoicemailUserpolicyServiceUnavailable() *PatchVoicemailUserpolicyServiceUnavailable {
	return &PatchVoicemailUserpolicyServiceUnavailable{}
}

/*
PatchVoicemailUserpolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchVoicemailUserpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy service unavailable response has a 2xx status code
func (o *PatchVoicemailUserpolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy service unavailable response has a 3xx status code
func (o *PatchVoicemailUserpolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy service unavailable response has a 4xx status code
func (o *PatchVoicemailUserpolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail userpolicy service unavailable response has a 5xx status code
func (o *PatchVoicemailUserpolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch voicemail userpolicy service unavailable response a status code equal to that given
func (o *PatchVoicemailUserpolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchVoicemailUserpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchVoicemailUserpolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchVoicemailUserpolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailUserpolicyGatewayTimeout creates a PatchVoicemailUserpolicyGatewayTimeout with default headers values
func NewPatchVoicemailUserpolicyGatewayTimeout() *PatchVoicemailUserpolicyGatewayTimeout {
	return &PatchVoicemailUserpolicyGatewayTimeout{}
}

/*
PatchVoicemailUserpolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchVoicemailUserpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail userpolicy gateway timeout response has a 2xx status code
func (o *PatchVoicemailUserpolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail userpolicy gateway timeout response has a 3xx status code
func (o *PatchVoicemailUserpolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail userpolicy gateway timeout response has a 4xx status code
func (o *PatchVoicemailUserpolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail userpolicy gateway timeout response has a 5xx status code
func (o *PatchVoicemailUserpolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch voicemail userpolicy gateway timeout response a status code equal to that given
func (o *PatchVoicemailUserpolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchVoicemailUserpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchVoicemailUserpolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/userpolicies/{userId}][%d] patchVoicemailUserpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchVoicemailUserpolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailUserpolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
