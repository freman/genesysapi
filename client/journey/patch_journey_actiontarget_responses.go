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

// PatchJourneyActiontargetReader is a Reader for the PatchJourneyActiontarget structure.
type PatchJourneyActiontargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchJourneyActiontargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchJourneyActiontargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchJourneyActiontargetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchJourneyActiontargetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchJourneyActiontargetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchJourneyActiontargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchJourneyActiontargetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchJourneyActiontargetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchJourneyActiontargetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchJourneyActiontargetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchJourneyActiontargetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchJourneyActiontargetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchJourneyActiontargetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchJourneyActiontargetOK creates a PatchJourneyActiontargetOK with default headers values
func NewPatchJourneyActiontargetOK() *PatchJourneyActiontargetOK {
	return &PatchJourneyActiontargetOK{}
}

/*
PatchJourneyActiontargetOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchJourneyActiontargetOK struct {
	Payload *models.ActionTarget
}

// IsSuccess returns true when this patch journey actiontarget o k response has a 2xx status code
func (o *PatchJourneyActiontargetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch journey actiontarget o k response has a 3xx status code
func (o *PatchJourneyActiontargetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget o k response has a 4xx status code
func (o *PatchJourneyActiontargetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch journey actiontarget o k response has a 5xx status code
func (o *PatchJourneyActiontargetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget o k response a status code equal to that given
func (o *PatchJourneyActiontargetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchJourneyActiontargetOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetOK  %+v", 200, o.Payload)
}

func (o *PatchJourneyActiontargetOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetOK  %+v", 200, o.Payload)
}

func (o *PatchJourneyActiontargetOK) GetPayload() *models.ActionTarget {
	return o.Payload
}

func (o *PatchJourneyActiontargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetBadRequest creates a PatchJourneyActiontargetBadRequest with default headers values
func NewPatchJourneyActiontargetBadRequest() *PatchJourneyActiontargetBadRequest {
	return &PatchJourneyActiontargetBadRequest{}
}

/*
PatchJourneyActiontargetBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchJourneyActiontargetBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget bad request response has a 2xx status code
func (o *PatchJourneyActiontargetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget bad request response has a 3xx status code
func (o *PatchJourneyActiontargetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget bad request response has a 4xx status code
func (o *PatchJourneyActiontargetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget bad request response has a 5xx status code
func (o *PatchJourneyActiontargetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget bad request response a status code equal to that given
func (o *PatchJourneyActiontargetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchJourneyActiontargetBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetBadRequest  %+v", 400, o.Payload)
}

func (o *PatchJourneyActiontargetBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetBadRequest  %+v", 400, o.Payload)
}

func (o *PatchJourneyActiontargetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetUnauthorized creates a PatchJourneyActiontargetUnauthorized with default headers values
func NewPatchJourneyActiontargetUnauthorized() *PatchJourneyActiontargetUnauthorized {
	return &PatchJourneyActiontargetUnauthorized{}
}

/*
PatchJourneyActiontargetUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchJourneyActiontargetUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget unauthorized response has a 2xx status code
func (o *PatchJourneyActiontargetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget unauthorized response has a 3xx status code
func (o *PatchJourneyActiontargetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget unauthorized response has a 4xx status code
func (o *PatchJourneyActiontargetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget unauthorized response has a 5xx status code
func (o *PatchJourneyActiontargetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget unauthorized response a status code equal to that given
func (o *PatchJourneyActiontargetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchJourneyActiontargetUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchJourneyActiontargetUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchJourneyActiontargetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetForbidden creates a PatchJourneyActiontargetForbidden with default headers values
func NewPatchJourneyActiontargetForbidden() *PatchJourneyActiontargetForbidden {
	return &PatchJourneyActiontargetForbidden{}
}

/*
PatchJourneyActiontargetForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchJourneyActiontargetForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget forbidden response has a 2xx status code
func (o *PatchJourneyActiontargetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget forbidden response has a 3xx status code
func (o *PatchJourneyActiontargetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget forbidden response has a 4xx status code
func (o *PatchJourneyActiontargetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget forbidden response has a 5xx status code
func (o *PatchJourneyActiontargetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget forbidden response a status code equal to that given
func (o *PatchJourneyActiontargetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchJourneyActiontargetForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetForbidden  %+v", 403, o.Payload)
}

func (o *PatchJourneyActiontargetForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetForbidden  %+v", 403, o.Payload)
}

func (o *PatchJourneyActiontargetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetNotFound creates a PatchJourneyActiontargetNotFound with default headers values
func NewPatchJourneyActiontargetNotFound() *PatchJourneyActiontargetNotFound {
	return &PatchJourneyActiontargetNotFound{}
}

/*
PatchJourneyActiontargetNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchJourneyActiontargetNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget not found response has a 2xx status code
func (o *PatchJourneyActiontargetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget not found response has a 3xx status code
func (o *PatchJourneyActiontargetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget not found response has a 4xx status code
func (o *PatchJourneyActiontargetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget not found response has a 5xx status code
func (o *PatchJourneyActiontargetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget not found response a status code equal to that given
func (o *PatchJourneyActiontargetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchJourneyActiontargetNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetNotFound  %+v", 404, o.Payload)
}

func (o *PatchJourneyActiontargetNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetNotFound  %+v", 404, o.Payload)
}

func (o *PatchJourneyActiontargetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetRequestTimeout creates a PatchJourneyActiontargetRequestTimeout with default headers values
func NewPatchJourneyActiontargetRequestTimeout() *PatchJourneyActiontargetRequestTimeout {
	return &PatchJourneyActiontargetRequestTimeout{}
}

/*
PatchJourneyActiontargetRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchJourneyActiontargetRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget request timeout response has a 2xx status code
func (o *PatchJourneyActiontargetRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget request timeout response has a 3xx status code
func (o *PatchJourneyActiontargetRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget request timeout response has a 4xx status code
func (o *PatchJourneyActiontargetRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget request timeout response has a 5xx status code
func (o *PatchJourneyActiontargetRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget request timeout response a status code equal to that given
func (o *PatchJourneyActiontargetRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchJourneyActiontargetRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchJourneyActiontargetRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchJourneyActiontargetRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetRequestEntityTooLarge creates a PatchJourneyActiontargetRequestEntityTooLarge with default headers values
func NewPatchJourneyActiontargetRequestEntityTooLarge() *PatchJourneyActiontargetRequestEntityTooLarge {
	return &PatchJourneyActiontargetRequestEntityTooLarge{}
}

/*
PatchJourneyActiontargetRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchJourneyActiontargetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget request entity too large response has a 2xx status code
func (o *PatchJourneyActiontargetRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget request entity too large response has a 3xx status code
func (o *PatchJourneyActiontargetRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget request entity too large response has a 4xx status code
func (o *PatchJourneyActiontargetRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget request entity too large response has a 5xx status code
func (o *PatchJourneyActiontargetRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget request entity too large response a status code equal to that given
func (o *PatchJourneyActiontargetRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchJourneyActiontargetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchJourneyActiontargetRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchJourneyActiontargetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetUnsupportedMediaType creates a PatchJourneyActiontargetUnsupportedMediaType with default headers values
func NewPatchJourneyActiontargetUnsupportedMediaType() *PatchJourneyActiontargetUnsupportedMediaType {
	return &PatchJourneyActiontargetUnsupportedMediaType{}
}

/*
PatchJourneyActiontargetUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchJourneyActiontargetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget unsupported media type response has a 2xx status code
func (o *PatchJourneyActiontargetUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget unsupported media type response has a 3xx status code
func (o *PatchJourneyActiontargetUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget unsupported media type response has a 4xx status code
func (o *PatchJourneyActiontargetUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget unsupported media type response has a 5xx status code
func (o *PatchJourneyActiontargetUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget unsupported media type response a status code equal to that given
func (o *PatchJourneyActiontargetUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchJourneyActiontargetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchJourneyActiontargetUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchJourneyActiontargetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetTooManyRequests creates a PatchJourneyActiontargetTooManyRequests with default headers values
func NewPatchJourneyActiontargetTooManyRequests() *PatchJourneyActiontargetTooManyRequests {
	return &PatchJourneyActiontargetTooManyRequests{}
}

/*
PatchJourneyActiontargetTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchJourneyActiontargetTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget too many requests response has a 2xx status code
func (o *PatchJourneyActiontargetTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget too many requests response has a 3xx status code
func (o *PatchJourneyActiontargetTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget too many requests response has a 4xx status code
func (o *PatchJourneyActiontargetTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch journey actiontarget too many requests response has a 5xx status code
func (o *PatchJourneyActiontargetTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch journey actiontarget too many requests response a status code equal to that given
func (o *PatchJourneyActiontargetTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchJourneyActiontargetTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchJourneyActiontargetTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchJourneyActiontargetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetInternalServerError creates a PatchJourneyActiontargetInternalServerError with default headers values
func NewPatchJourneyActiontargetInternalServerError() *PatchJourneyActiontargetInternalServerError {
	return &PatchJourneyActiontargetInternalServerError{}
}

/*
PatchJourneyActiontargetInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchJourneyActiontargetInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget internal server error response has a 2xx status code
func (o *PatchJourneyActiontargetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget internal server error response has a 3xx status code
func (o *PatchJourneyActiontargetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget internal server error response has a 4xx status code
func (o *PatchJourneyActiontargetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch journey actiontarget internal server error response has a 5xx status code
func (o *PatchJourneyActiontargetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch journey actiontarget internal server error response a status code equal to that given
func (o *PatchJourneyActiontargetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchJourneyActiontargetInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchJourneyActiontargetInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchJourneyActiontargetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetServiceUnavailable creates a PatchJourneyActiontargetServiceUnavailable with default headers values
func NewPatchJourneyActiontargetServiceUnavailable() *PatchJourneyActiontargetServiceUnavailable {
	return &PatchJourneyActiontargetServiceUnavailable{}
}

/*
PatchJourneyActiontargetServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchJourneyActiontargetServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget service unavailable response has a 2xx status code
func (o *PatchJourneyActiontargetServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget service unavailable response has a 3xx status code
func (o *PatchJourneyActiontargetServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget service unavailable response has a 4xx status code
func (o *PatchJourneyActiontargetServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch journey actiontarget service unavailable response has a 5xx status code
func (o *PatchJourneyActiontargetServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch journey actiontarget service unavailable response a status code equal to that given
func (o *PatchJourneyActiontargetServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchJourneyActiontargetServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchJourneyActiontargetServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchJourneyActiontargetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActiontargetGatewayTimeout creates a PatchJourneyActiontargetGatewayTimeout with default headers values
func NewPatchJourneyActiontargetGatewayTimeout() *PatchJourneyActiontargetGatewayTimeout {
	return &PatchJourneyActiontargetGatewayTimeout{}
}

/*
PatchJourneyActiontargetGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchJourneyActiontargetGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch journey actiontarget gateway timeout response has a 2xx status code
func (o *PatchJourneyActiontargetGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch journey actiontarget gateway timeout response has a 3xx status code
func (o *PatchJourneyActiontargetGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch journey actiontarget gateway timeout response has a 4xx status code
func (o *PatchJourneyActiontargetGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch journey actiontarget gateway timeout response has a 5xx status code
func (o *PatchJourneyActiontargetGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch journey actiontarget gateway timeout response a status code equal to that given
func (o *PatchJourneyActiontargetGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchJourneyActiontargetGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchJourneyActiontargetGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actiontargets/{actionTargetId}][%d] patchJourneyActiontargetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchJourneyActiontargetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActiontargetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
