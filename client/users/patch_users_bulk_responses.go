// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchUsersBulkReader is a Reader for the PatchUsersBulk structure.
type PatchUsersBulkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUsersBulkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUsersBulkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUsersBulkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUsersBulkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUsersBulkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUsersBulkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchUsersBulkRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchUsersBulkRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchUsersBulkUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchUsersBulkTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUsersBulkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchUsersBulkServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchUsersBulkGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUsersBulkOK creates a PatchUsersBulkOK with default headers values
func NewPatchUsersBulkOK() *PatchUsersBulkOK {
	return &PatchUsersBulkOK{}
}

/*
PatchUsersBulkOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchUsersBulkOK struct {
	Payload *models.UserEntityListing
}

// IsSuccess returns true when this patch users bulk o k response has a 2xx status code
func (o *PatchUsersBulkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch users bulk o k response has a 3xx status code
func (o *PatchUsersBulkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk o k response has a 4xx status code
func (o *PatchUsersBulkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch users bulk o k response has a 5xx status code
func (o *PatchUsersBulkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk o k response a status code equal to that given
func (o *PatchUsersBulkOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchUsersBulkOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkOK  %+v", 200, o.Payload)
}

func (o *PatchUsersBulkOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkOK  %+v", 200, o.Payload)
}

func (o *PatchUsersBulkOK) GetPayload() *models.UserEntityListing {
	return o.Payload
}

func (o *PatchUsersBulkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkBadRequest creates a PatchUsersBulkBadRequest with default headers values
func NewPatchUsersBulkBadRequest() *PatchUsersBulkBadRequest {
	return &PatchUsersBulkBadRequest{}
}

/*
PatchUsersBulkBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchUsersBulkBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk bad request response has a 2xx status code
func (o *PatchUsersBulkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk bad request response has a 3xx status code
func (o *PatchUsersBulkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk bad request response has a 4xx status code
func (o *PatchUsersBulkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk bad request response has a 5xx status code
func (o *PatchUsersBulkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk bad request response a status code equal to that given
func (o *PatchUsersBulkBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchUsersBulkBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUsersBulkBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUsersBulkBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkUnauthorized creates a PatchUsersBulkUnauthorized with default headers values
func NewPatchUsersBulkUnauthorized() *PatchUsersBulkUnauthorized {
	return &PatchUsersBulkUnauthorized{}
}

/*
PatchUsersBulkUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchUsersBulkUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk unauthorized response has a 2xx status code
func (o *PatchUsersBulkUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk unauthorized response has a 3xx status code
func (o *PatchUsersBulkUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk unauthorized response has a 4xx status code
func (o *PatchUsersBulkUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk unauthorized response has a 5xx status code
func (o *PatchUsersBulkUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk unauthorized response a status code equal to that given
func (o *PatchUsersBulkUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchUsersBulkUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUsersBulkUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUsersBulkUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkForbidden creates a PatchUsersBulkForbidden with default headers values
func NewPatchUsersBulkForbidden() *PatchUsersBulkForbidden {
	return &PatchUsersBulkForbidden{}
}

/*
PatchUsersBulkForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchUsersBulkForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk forbidden response has a 2xx status code
func (o *PatchUsersBulkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk forbidden response has a 3xx status code
func (o *PatchUsersBulkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk forbidden response has a 4xx status code
func (o *PatchUsersBulkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk forbidden response has a 5xx status code
func (o *PatchUsersBulkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk forbidden response a status code equal to that given
func (o *PatchUsersBulkForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchUsersBulkForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkForbidden  %+v", 403, o.Payload)
}

func (o *PatchUsersBulkForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkForbidden  %+v", 403, o.Payload)
}

func (o *PatchUsersBulkForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkNotFound creates a PatchUsersBulkNotFound with default headers values
func NewPatchUsersBulkNotFound() *PatchUsersBulkNotFound {
	return &PatchUsersBulkNotFound{}
}

/*
PatchUsersBulkNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchUsersBulkNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk not found response has a 2xx status code
func (o *PatchUsersBulkNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk not found response has a 3xx status code
func (o *PatchUsersBulkNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk not found response has a 4xx status code
func (o *PatchUsersBulkNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk not found response has a 5xx status code
func (o *PatchUsersBulkNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk not found response a status code equal to that given
func (o *PatchUsersBulkNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchUsersBulkNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkNotFound  %+v", 404, o.Payload)
}

func (o *PatchUsersBulkNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkNotFound  %+v", 404, o.Payload)
}

func (o *PatchUsersBulkNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkRequestTimeout creates a PatchUsersBulkRequestTimeout with default headers values
func NewPatchUsersBulkRequestTimeout() *PatchUsersBulkRequestTimeout {
	return &PatchUsersBulkRequestTimeout{}
}

/*
PatchUsersBulkRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchUsersBulkRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk request timeout response has a 2xx status code
func (o *PatchUsersBulkRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk request timeout response has a 3xx status code
func (o *PatchUsersBulkRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk request timeout response has a 4xx status code
func (o *PatchUsersBulkRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk request timeout response has a 5xx status code
func (o *PatchUsersBulkRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk request timeout response a status code equal to that given
func (o *PatchUsersBulkRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchUsersBulkRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchUsersBulkRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchUsersBulkRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkRequestEntityTooLarge creates a PatchUsersBulkRequestEntityTooLarge with default headers values
func NewPatchUsersBulkRequestEntityTooLarge() *PatchUsersBulkRequestEntityTooLarge {
	return &PatchUsersBulkRequestEntityTooLarge{}
}

/*
PatchUsersBulkRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchUsersBulkRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk request entity too large response has a 2xx status code
func (o *PatchUsersBulkRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk request entity too large response has a 3xx status code
func (o *PatchUsersBulkRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk request entity too large response has a 4xx status code
func (o *PatchUsersBulkRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk request entity too large response has a 5xx status code
func (o *PatchUsersBulkRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk request entity too large response a status code equal to that given
func (o *PatchUsersBulkRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchUsersBulkRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUsersBulkRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUsersBulkRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkUnsupportedMediaType creates a PatchUsersBulkUnsupportedMediaType with default headers values
func NewPatchUsersBulkUnsupportedMediaType() *PatchUsersBulkUnsupportedMediaType {
	return &PatchUsersBulkUnsupportedMediaType{}
}

/*
PatchUsersBulkUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchUsersBulkUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk unsupported media type response has a 2xx status code
func (o *PatchUsersBulkUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk unsupported media type response has a 3xx status code
func (o *PatchUsersBulkUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk unsupported media type response has a 4xx status code
func (o *PatchUsersBulkUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk unsupported media type response has a 5xx status code
func (o *PatchUsersBulkUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk unsupported media type response a status code equal to that given
func (o *PatchUsersBulkUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchUsersBulkUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUsersBulkUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUsersBulkUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkTooManyRequests creates a PatchUsersBulkTooManyRequests with default headers values
func NewPatchUsersBulkTooManyRequests() *PatchUsersBulkTooManyRequests {
	return &PatchUsersBulkTooManyRequests{}
}

/*
PatchUsersBulkTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchUsersBulkTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk too many requests response has a 2xx status code
func (o *PatchUsersBulkTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk too many requests response has a 3xx status code
func (o *PatchUsersBulkTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk too many requests response has a 4xx status code
func (o *PatchUsersBulkTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch users bulk too many requests response has a 5xx status code
func (o *PatchUsersBulkTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch users bulk too many requests response a status code equal to that given
func (o *PatchUsersBulkTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchUsersBulkTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUsersBulkTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUsersBulkTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkInternalServerError creates a PatchUsersBulkInternalServerError with default headers values
func NewPatchUsersBulkInternalServerError() *PatchUsersBulkInternalServerError {
	return &PatchUsersBulkInternalServerError{}
}

/*
PatchUsersBulkInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchUsersBulkInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk internal server error response has a 2xx status code
func (o *PatchUsersBulkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk internal server error response has a 3xx status code
func (o *PatchUsersBulkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk internal server error response has a 4xx status code
func (o *PatchUsersBulkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch users bulk internal server error response has a 5xx status code
func (o *PatchUsersBulkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch users bulk internal server error response a status code equal to that given
func (o *PatchUsersBulkInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchUsersBulkInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUsersBulkInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUsersBulkInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkServiceUnavailable creates a PatchUsersBulkServiceUnavailable with default headers values
func NewPatchUsersBulkServiceUnavailable() *PatchUsersBulkServiceUnavailable {
	return &PatchUsersBulkServiceUnavailable{}
}

/*
PatchUsersBulkServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchUsersBulkServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk service unavailable response has a 2xx status code
func (o *PatchUsersBulkServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk service unavailable response has a 3xx status code
func (o *PatchUsersBulkServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk service unavailable response has a 4xx status code
func (o *PatchUsersBulkServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch users bulk service unavailable response has a 5xx status code
func (o *PatchUsersBulkServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch users bulk service unavailable response a status code equal to that given
func (o *PatchUsersBulkServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchUsersBulkServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUsersBulkServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUsersBulkServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsersBulkGatewayTimeout creates a PatchUsersBulkGatewayTimeout with default headers values
func NewPatchUsersBulkGatewayTimeout() *PatchUsersBulkGatewayTimeout {
	return &PatchUsersBulkGatewayTimeout{}
}

/*
PatchUsersBulkGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchUsersBulkGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch users bulk gateway timeout response has a 2xx status code
func (o *PatchUsersBulkGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch users bulk gateway timeout response has a 3xx status code
func (o *PatchUsersBulkGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch users bulk gateway timeout response has a 4xx status code
func (o *PatchUsersBulkGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch users bulk gateway timeout response has a 5xx status code
func (o *PatchUsersBulkGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch users bulk gateway timeout response a status code equal to that given
func (o *PatchUsersBulkGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchUsersBulkGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUsersBulkGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/bulk][%d] patchUsersBulkGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUsersBulkGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUsersBulkGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
