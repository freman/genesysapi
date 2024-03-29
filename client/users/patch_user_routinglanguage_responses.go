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

// PatchUserRoutinglanguageReader is a Reader for the PatchUserRoutinglanguage structure.
type PatchUserRoutinglanguageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserRoutinglanguageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserRoutinglanguageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUserRoutinglanguageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUserRoutinglanguageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUserRoutinglanguageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUserRoutinglanguageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchUserRoutinglanguageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchUserRoutinglanguageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchUserRoutinglanguageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchUserRoutinglanguageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchUserRoutinglanguageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUserRoutinglanguageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchUserRoutinglanguageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchUserRoutinglanguageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUserRoutinglanguageOK creates a PatchUserRoutinglanguageOK with default headers values
func NewPatchUserRoutinglanguageOK() *PatchUserRoutinglanguageOK {
	return &PatchUserRoutinglanguageOK{}
}

/*
PatchUserRoutinglanguageOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchUserRoutinglanguageOK struct {
	Payload *models.UserRoutingLanguage
}

// IsSuccess returns true when this patch user routinglanguage o k response has a 2xx status code
func (o *PatchUserRoutinglanguageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch user routinglanguage o k response has a 3xx status code
func (o *PatchUserRoutinglanguageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage o k response has a 4xx status code
func (o *PatchUserRoutinglanguageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user routinglanguage o k response has a 5xx status code
func (o *PatchUserRoutinglanguageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage o k response a status code equal to that given
func (o *PatchUserRoutinglanguageOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchUserRoutinglanguageOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageOK  %+v", 200, o.Payload)
}

func (o *PatchUserRoutinglanguageOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageOK  %+v", 200, o.Payload)
}

func (o *PatchUserRoutinglanguageOK) GetPayload() *models.UserRoutingLanguage {
	return o.Payload
}

func (o *PatchUserRoutinglanguageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoutingLanguage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageBadRequest creates a PatchUserRoutinglanguageBadRequest with default headers values
func NewPatchUserRoutinglanguageBadRequest() *PatchUserRoutinglanguageBadRequest {
	return &PatchUserRoutinglanguageBadRequest{}
}

/*
PatchUserRoutinglanguageBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchUserRoutinglanguageBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage bad request response has a 2xx status code
func (o *PatchUserRoutinglanguageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage bad request response has a 3xx status code
func (o *PatchUserRoutinglanguageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage bad request response has a 4xx status code
func (o *PatchUserRoutinglanguageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage bad request response has a 5xx status code
func (o *PatchUserRoutinglanguageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage bad request response a status code equal to that given
func (o *PatchUserRoutinglanguageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchUserRoutinglanguageBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUserRoutinglanguageBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUserRoutinglanguageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageUnauthorized creates a PatchUserRoutinglanguageUnauthorized with default headers values
func NewPatchUserRoutinglanguageUnauthorized() *PatchUserRoutinglanguageUnauthorized {
	return &PatchUserRoutinglanguageUnauthorized{}
}

/*
PatchUserRoutinglanguageUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchUserRoutinglanguageUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage unauthorized response has a 2xx status code
func (o *PatchUserRoutinglanguageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage unauthorized response has a 3xx status code
func (o *PatchUserRoutinglanguageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage unauthorized response has a 4xx status code
func (o *PatchUserRoutinglanguageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage unauthorized response has a 5xx status code
func (o *PatchUserRoutinglanguageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage unauthorized response a status code equal to that given
func (o *PatchUserRoutinglanguageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchUserRoutinglanguageUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUserRoutinglanguageUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUserRoutinglanguageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageForbidden creates a PatchUserRoutinglanguageForbidden with default headers values
func NewPatchUserRoutinglanguageForbidden() *PatchUserRoutinglanguageForbidden {
	return &PatchUserRoutinglanguageForbidden{}
}

/*
PatchUserRoutinglanguageForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchUserRoutinglanguageForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage forbidden response has a 2xx status code
func (o *PatchUserRoutinglanguageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage forbidden response has a 3xx status code
func (o *PatchUserRoutinglanguageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage forbidden response has a 4xx status code
func (o *PatchUserRoutinglanguageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage forbidden response has a 5xx status code
func (o *PatchUserRoutinglanguageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage forbidden response a status code equal to that given
func (o *PatchUserRoutinglanguageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchUserRoutinglanguageForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserRoutinglanguageForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserRoutinglanguageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageNotFound creates a PatchUserRoutinglanguageNotFound with default headers values
func NewPatchUserRoutinglanguageNotFound() *PatchUserRoutinglanguageNotFound {
	return &PatchUserRoutinglanguageNotFound{}
}

/*
PatchUserRoutinglanguageNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchUserRoutinglanguageNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage not found response has a 2xx status code
func (o *PatchUserRoutinglanguageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage not found response has a 3xx status code
func (o *PatchUserRoutinglanguageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage not found response has a 4xx status code
func (o *PatchUserRoutinglanguageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage not found response has a 5xx status code
func (o *PatchUserRoutinglanguageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage not found response a status code equal to that given
func (o *PatchUserRoutinglanguageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchUserRoutinglanguageNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserRoutinglanguageNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserRoutinglanguageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageRequestTimeout creates a PatchUserRoutinglanguageRequestTimeout with default headers values
func NewPatchUserRoutinglanguageRequestTimeout() *PatchUserRoutinglanguageRequestTimeout {
	return &PatchUserRoutinglanguageRequestTimeout{}
}

/*
PatchUserRoutinglanguageRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchUserRoutinglanguageRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage request timeout response has a 2xx status code
func (o *PatchUserRoutinglanguageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage request timeout response has a 3xx status code
func (o *PatchUserRoutinglanguageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage request timeout response has a 4xx status code
func (o *PatchUserRoutinglanguageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage request timeout response has a 5xx status code
func (o *PatchUserRoutinglanguageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage request timeout response a status code equal to that given
func (o *PatchUserRoutinglanguageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchUserRoutinglanguageRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchUserRoutinglanguageRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchUserRoutinglanguageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageConflict creates a PatchUserRoutinglanguageConflict with default headers values
func NewPatchUserRoutinglanguageConflict() *PatchUserRoutinglanguageConflict {
	return &PatchUserRoutinglanguageConflict{}
}

/*
PatchUserRoutinglanguageConflict describes a response with status code 409, with default header values.

Resource conflict - Unexpected version was provided
*/
type PatchUserRoutinglanguageConflict struct {
}

// IsSuccess returns true when this patch user routinglanguage conflict response has a 2xx status code
func (o *PatchUserRoutinglanguageConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage conflict response has a 3xx status code
func (o *PatchUserRoutinglanguageConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage conflict response has a 4xx status code
func (o *PatchUserRoutinglanguageConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage conflict response has a 5xx status code
func (o *PatchUserRoutinglanguageConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage conflict response a status code equal to that given
func (o *PatchUserRoutinglanguageConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PatchUserRoutinglanguageConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageConflict ", 409)
}

func (o *PatchUserRoutinglanguageConflict) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageConflict ", 409)
}

func (o *PatchUserRoutinglanguageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchUserRoutinglanguageRequestEntityTooLarge creates a PatchUserRoutinglanguageRequestEntityTooLarge with default headers values
func NewPatchUserRoutinglanguageRequestEntityTooLarge() *PatchUserRoutinglanguageRequestEntityTooLarge {
	return &PatchUserRoutinglanguageRequestEntityTooLarge{}
}

/*
PatchUserRoutinglanguageRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchUserRoutinglanguageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage request entity too large response has a 2xx status code
func (o *PatchUserRoutinglanguageRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage request entity too large response has a 3xx status code
func (o *PatchUserRoutinglanguageRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage request entity too large response has a 4xx status code
func (o *PatchUserRoutinglanguageRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage request entity too large response has a 5xx status code
func (o *PatchUserRoutinglanguageRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage request entity too large response a status code equal to that given
func (o *PatchUserRoutinglanguageRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchUserRoutinglanguageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUserRoutinglanguageRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUserRoutinglanguageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageUnsupportedMediaType creates a PatchUserRoutinglanguageUnsupportedMediaType with default headers values
func NewPatchUserRoutinglanguageUnsupportedMediaType() *PatchUserRoutinglanguageUnsupportedMediaType {
	return &PatchUserRoutinglanguageUnsupportedMediaType{}
}

/*
PatchUserRoutinglanguageUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchUserRoutinglanguageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage unsupported media type response has a 2xx status code
func (o *PatchUserRoutinglanguageUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage unsupported media type response has a 3xx status code
func (o *PatchUserRoutinglanguageUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage unsupported media type response has a 4xx status code
func (o *PatchUserRoutinglanguageUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage unsupported media type response has a 5xx status code
func (o *PatchUserRoutinglanguageUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage unsupported media type response a status code equal to that given
func (o *PatchUserRoutinglanguageUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchUserRoutinglanguageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUserRoutinglanguageUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUserRoutinglanguageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageTooManyRequests creates a PatchUserRoutinglanguageTooManyRequests with default headers values
func NewPatchUserRoutinglanguageTooManyRequests() *PatchUserRoutinglanguageTooManyRequests {
	return &PatchUserRoutinglanguageTooManyRequests{}
}

/*
PatchUserRoutinglanguageTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchUserRoutinglanguageTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage too many requests response has a 2xx status code
func (o *PatchUserRoutinglanguageTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage too many requests response has a 3xx status code
func (o *PatchUserRoutinglanguageTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage too many requests response has a 4xx status code
func (o *PatchUserRoutinglanguageTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch user routinglanguage too many requests response has a 5xx status code
func (o *PatchUserRoutinglanguageTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch user routinglanguage too many requests response a status code equal to that given
func (o *PatchUserRoutinglanguageTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchUserRoutinglanguageTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUserRoutinglanguageTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUserRoutinglanguageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageInternalServerError creates a PatchUserRoutinglanguageInternalServerError with default headers values
func NewPatchUserRoutinglanguageInternalServerError() *PatchUserRoutinglanguageInternalServerError {
	return &PatchUserRoutinglanguageInternalServerError{}
}

/*
PatchUserRoutinglanguageInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchUserRoutinglanguageInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage internal server error response has a 2xx status code
func (o *PatchUserRoutinglanguageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage internal server error response has a 3xx status code
func (o *PatchUserRoutinglanguageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage internal server error response has a 4xx status code
func (o *PatchUserRoutinglanguageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user routinglanguage internal server error response has a 5xx status code
func (o *PatchUserRoutinglanguageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch user routinglanguage internal server error response a status code equal to that given
func (o *PatchUserRoutinglanguageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchUserRoutinglanguageInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUserRoutinglanguageInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUserRoutinglanguageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageServiceUnavailable creates a PatchUserRoutinglanguageServiceUnavailable with default headers values
func NewPatchUserRoutinglanguageServiceUnavailable() *PatchUserRoutinglanguageServiceUnavailable {
	return &PatchUserRoutinglanguageServiceUnavailable{}
}

/*
PatchUserRoutinglanguageServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchUserRoutinglanguageServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage service unavailable response has a 2xx status code
func (o *PatchUserRoutinglanguageServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage service unavailable response has a 3xx status code
func (o *PatchUserRoutinglanguageServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage service unavailable response has a 4xx status code
func (o *PatchUserRoutinglanguageServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user routinglanguage service unavailable response has a 5xx status code
func (o *PatchUserRoutinglanguageServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch user routinglanguage service unavailable response a status code equal to that given
func (o *PatchUserRoutinglanguageServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchUserRoutinglanguageServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUserRoutinglanguageServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUserRoutinglanguageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserRoutinglanguageGatewayTimeout creates a PatchUserRoutinglanguageGatewayTimeout with default headers values
func NewPatchUserRoutinglanguageGatewayTimeout() *PatchUserRoutinglanguageGatewayTimeout {
	return &PatchUserRoutinglanguageGatewayTimeout{}
}

/*
PatchUserRoutinglanguageGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchUserRoutinglanguageGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch user routinglanguage gateway timeout response has a 2xx status code
func (o *PatchUserRoutinglanguageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch user routinglanguage gateway timeout response has a 3xx status code
func (o *PatchUserRoutinglanguageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch user routinglanguage gateway timeout response has a 4xx status code
func (o *PatchUserRoutinglanguageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch user routinglanguage gateway timeout response has a 5xx status code
func (o *PatchUserRoutinglanguageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch user routinglanguage gateway timeout response a status code equal to that given
func (o *PatchUserRoutinglanguageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchUserRoutinglanguageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUserRoutinglanguageGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/routinglanguages/{languageId}][%d] patchUserRoutinglanguageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUserRoutinglanguageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserRoutinglanguageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
