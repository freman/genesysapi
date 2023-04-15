// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyActiontargetsReader is a Reader for the GetJourneyActiontargets structure.
type GetJourneyActiontargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyActiontargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyActiontargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyActiontargetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyActiontargetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyActiontargetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyActiontargetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyActiontargetsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyActiontargetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyActiontargetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyActiontargetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyActiontargetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyActiontargetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyActiontargetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyActiontargetsOK creates a GetJourneyActiontargetsOK with default headers values
func NewGetJourneyActiontargetsOK() *GetJourneyActiontargetsOK {
	return &GetJourneyActiontargetsOK{}
}

/*
GetJourneyActiontargetsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJourneyActiontargetsOK struct {
	Payload *models.ActionTargetListing
}

// IsSuccess returns true when this get journey actiontargets o k response has a 2xx status code
func (o *GetJourneyActiontargetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey actiontargets o k response has a 3xx status code
func (o *GetJourneyActiontargetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets o k response has a 4xx status code
func (o *GetJourneyActiontargetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontargets o k response has a 5xx status code
func (o *GetJourneyActiontargetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets o k response a status code equal to that given
func (o *GetJourneyActiontargetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJourneyActiontargetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActiontargetsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActiontargetsOK) GetPayload() *models.ActionTargetListing {
	return o.Payload
}

func (o *GetJourneyActiontargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTargetListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsBadRequest creates a GetJourneyActiontargetsBadRequest with default headers values
func NewGetJourneyActiontargetsBadRequest() *GetJourneyActiontargetsBadRequest {
	return &GetJourneyActiontargetsBadRequest{}
}

/*
GetJourneyActiontargetsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyActiontargetsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets bad request response has a 2xx status code
func (o *GetJourneyActiontargetsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets bad request response has a 3xx status code
func (o *GetJourneyActiontargetsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets bad request response has a 4xx status code
func (o *GetJourneyActiontargetsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets bad request response has a 5xx status code
func (o *GetJourneyActiontargetsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets bad request response a status code equal to that given
func (o *GetJourneyActiontargetsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJourneyActiontargetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActiontargetsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActiontargetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsUnauthorized creates a GetJourneyActiontargetsUnauthorized with default headers values
func NewGetJourneyActiontargetsUnauthorized() *GetJourneyActiontargetsUnauthorized {
	return &GetJourneyActiontargetsUnauthorized{}
}

/*
GetJourneyActiontargetsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyActiontargetsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets unauthorized response has a 2xx status code
func (o *GetJourneyActiontargetsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets unauthorized response has a 3xx status code
func (o *GetJourneyActiontargetsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets unauthorized response has a 4xx status code
func (o *GetJourneyActiontargetsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets unauthorized response has a 5xx status code
func (o *GetJourneyActiontargetsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets unauthorized response a status code equal to that given
func (o *GetJourneyActiontargetsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetJourneyActiontargetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActiontargetsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActiontargetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsForbidden creates a GetJourneyActiontargetsForbidden with default headers values
func NewGetJourneyActiontargetsForbidden() *GetJourneyActiontargetsForbidden {
	return &GetJourneyActiontargetsForbidden{}
}

/*
GetJourneyActiontargetsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyActiontargetsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets forbidden response has a 2xx status code
func (o *GetJourneyActiontargetsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets forbidden response has a 3xx status code
func (o *GetJourneyActiontargetsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets forbidden response has a 4xx status code
func (o *GetJourneyActiontargetsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets forbidden response has a 5xx status code
func (o *GetJourneyActiontargetsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets forbidden response a status code equal to that given
func (o *GetJourneyActiontargetsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetJourneyActiontargetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActiontargetsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActiontargetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsNotFound creates a GetJourneyActiontargetsNotFound with default headers values
func NewGetJourneyActiontargetsNotFound() *GetJourneyActiontargetsNotFound {
	return &GetJourneyActiontargetsNotFound{}
}

/*
GetJourneyActiontargetsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetJourneyActiontargetsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets not found response has a 2xx status code
func (o *GetJourneyActiontargetsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets not found response has a 3xx status code
func (o *GetJourneyActiontargetsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets not found response has a 4xx status code
func (o *GetJourneyActiontargetsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets not found response has a 5xx status code
func (o *GetJourneyActiontargetsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets not found response a status code equal to that given
func (o *GetJourneyActiontargetsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJourneyActiontargetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActiontargetsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActiontargetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsRequestTimeout creates a GetJourneyActiontargetsRequestTimeout with default headers values
func NewGetJourneyActiontargetsRequestTimeout() *GetJourneyActiontargetsRequestTimeout {
	return &GetJourneyActiontargetsRequestTimeout{}
}

/*
GetJourneyActiontargetsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyActiontargetsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets request timeout response has a 2xx status code
func (o *GetJourneyActiontargetsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets request timeout response has a 3xx status code
func (o *GetJourneyActiontargetsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets request timeout response has a 4xx status code
func (o *GetJourneyActiontargetsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets request timeout response has a 5xx status code
func (o *GetJourneyActiontargetsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets request timeout response a status code equal to that given
func (o *GetJourneyActiontargetsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetJourneyActiontargetsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActiontargetsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActiontargetsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsRequestEntityTooLarge creates a GetJourneyActiontargetsRequestEntityTooLarge with default headers values
func NewGetJourneyActiontargetsRequestEntityTooLarge() *GetJourneyActiontargetsRequestEntityTooLarge {
	return &GetJourneyActiontargetsRequestEntityTooLarge{}
}

/*
GetJourneyActiontargetsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetJourneyActiontargetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets request entity too large response has a 2xx status code
func (o *GetJourneyActiontargetsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets request entity too large response has a 3xx status code
func (o *GetJourneyActiontargetsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets request entity too large response has a 4xx status code
func (o *GetJourneyActiontargetsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets request entity too large response has a 5xx status code
func (o *GetJourneyActiontargetsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets request entity too large response a status code equal to that given
func (o *GetJourneyActiontargetsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetJourneyActiontargetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActiontargetsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActiontargetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsUnsupportedMediaType creates a GetJourneyActiontargetsUnsupportedMediaType with default headers values
func NewGetJourneyActiontargetsUnsupportedMediaType() *GetJourneyActiontargetsUnsupportedMediaType {
	return &GetJourneyActiontargetsUnsupportedMediaType{}
}

/*
GetJourneyActiontargetsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyActiontargetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets unsupported media type response has a 2xx status code
func (o *GetJourneyActiontargetsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets unsupported media type response has a 3xx status code
func (o *GetJourneyActiontargetsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets unsupported media type response has a 4xx status code
func (o *GetJourneyActiontargetsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets unsupported media type response has a 5xx status code
func (o *GetJourneyActiontargetsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets unsupported media type response a status code equal to that given
func (o *GetJourneyActiontargetsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetJourneyActiontargetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActiontargetsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActiontargetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsTooManyRequests creates a GetJourneyActiontargetsTooManyRequests with default headers values
func NewGetJourneyActiontargetsTooManyRequests() *GetJourneyActiontargetsTooManyRequests {
	return &GetJourneyActiontargetsTooManyRequests{}
}

/*
GetJourneyActiontargetsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyActiontargetsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets too many requests response has a 2xx status code
func (o *GetJourneyActiontargetsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets too many requests response has a 3xx status code
func (o *GetJourneyActiontargetsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets too many requests response has a 4xx status code
func (o *GetJourneyActiontargetsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actiontargets too many requests response has a 5xx status code
func (o *GetJourneyActiontargetsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actiontargets too many requests response a status code equal to that given
func (o *GetJourneyActiontargetsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetJourneyActiontargetsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActiontargetsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActiontargetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsInternalServerError creates a GetJourneyActiontargetsInternalServerError with default headers values
func NewGetJourneyActiontargetsInternalServerError() *GetJourneyActiontargetsInternalServerError {
	return &GetJourneyActiontargetsInternalServerError{}
}

/*
GetJourneyActiontargetsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyActiontargetsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets internal server error response has a 2xx status code
func (o *GetJourneyActiontargetsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets internal server error response has a 3xx status code
func (o *GetJourneyActiontargetsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets internal server error response has a 4xx status code
func (o *GetJourneyActiontargetsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontargets internal server error response has a 5xx status code
func (o *GetJourneyActiontargetsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actiontargets internal server error response a status code equal to that given
func (o *GetJourneyActiontargetsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJourneyActiontargetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActiontargetsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActiontargetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsServiceUnavailable creates a GetJourneyActiontargetsServiceUnavailable with default headers values
func NewGetJourneyActiontargetsServiceUnavailable() *GetJourneyActiontargetsServiceUnavailable {
	return &GetJourneyActiontargetsServiceUnavailable{}
}

/*
GetJourneyActiontargetsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyActiontargetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets service unavailable response has a 2xx status code
func (o *GetJourneyActiontargetsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets service unavailable response has a 3xx status code
func (o *GetJourneyActiontargetsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets service unavailable response has a 4xx status code
func (o *GetJourneyActiontargetsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontargets service unavailable response has a 5xx status code
func (o *GetJourneyActiontargetsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actiontargets service unavailable response a status code equal to that given
func (o *GetJourneyActiontargetsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetJourneyActiontargetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActiontargetsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActiontargetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetsGatewayTimeout creates a GetJourneyActiontargetsGatewayTimeout with default headers values
func NewGetJourneyActiontargetsGatewayTimeout() *GetJourneyActiontargetsGatewayTimeout {
	return &GetJourneyActiontargetsGatewayTimeout{}
}

/*
GetJourneyActiontargetsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetJourneyActiontargetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actiontargets gateway timeout response has a 2xx status code
func (o *GetJourneyActiontargetsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actiontargets gateway timeout response has a 3xx status code
func (o *GetJourneyActiontargetsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actiontargets gateway timeout response has a 4xx status code
func (o *GetJourneyActiontargetsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actiontargets gateway timeout response has a 5xx status code
func (o *GetJourneyActiontargetsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actiontargets gateway timeout response a status code equal to that given
func (o *GetJourneyActiontargetsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetJourneyActiontargetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActiontargetsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets][%d] getJourneyActiontargetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActiontargetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
