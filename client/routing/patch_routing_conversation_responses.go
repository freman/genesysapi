// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchRoutingConversationReader is a Reader for the PatchRoutingConversation structure.
type PatchRoutingConversationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingConversationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoutingConversationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingConversationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingConversationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingConversationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingConversationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchRoutingConversationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingConversationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingConversationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingConversationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingConversationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingConversationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingConversationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingConversationOK creates a PatchRoutingConversationOK with default headers values
func NewPatchRoutingConversationOK() *PatchRoutingConversationOK {
	return &PatchRoutingConversationOK{}
}

/*
PatchRoutingConversationOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchRoutingConversationOK struct {
	Payload *models.RoutingConversationAttributesResponse
}

// IsSuccess returns true when this patch routing conversation o k response has a 2xx status code
func (o *PatchRoutingConversationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch routing conversation o k response has a 3xx status code
func (o *PatchRoutingConversationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation o k response has a 4xx status code
func (o *PatchRoutingConversationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch routing conversation o k response has a 5xx status code
func (o *PatchRoutingConversationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation o k response a status code equal to that given
func (o *PatchRoutingConversationOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchRoutingConversationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingConversationOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingConversationOK) GetPayload() *models.RoutingConversationAttributesResponse {
	return o.Payload
}

func (o *PatchRoutingConversationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingConversationAttributesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationBadRequest creates a PatchRoutingConversationBadRequest with default headers values
func NewPatchRoutingConversationBadRequest() *PatchRoutingConversationBadRequest {
	return &PatchRoutingConversationBadRequest{}
}

/*
PatchRoutingConversationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingConversationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation bad request response has a 2xx status code
func (o *PatchRoutingConversationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation bad request response has a 3xx status code
func (o *PatchRoutingConversationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation bad request response has a 4xx status code
func (o *PatchRoutingConversationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation bad request response has a 5xx status code
func (o *PatchRoutingConversationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation bad request response a status code equal to that given
func (o *PatchRoutingConversationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchRoutingConversationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingConversationBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingConversationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationUnauthorized creates a PatchRoutingConversationUnauthorized with default headers values
func NewPatchRoutingConversationUnauthorized() *PatchRoutingConversationUnauthorized {
	return &PatchRoutingConversationUnauthorized{}
}

/*
PatchRoutingConversationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingConversationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation unauthorized response has a 2xx status code
func (o *PatchRoutingConversationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation unauthorized response has a 3xx status code
func (o *PatchRoutingConversationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation unauthorized response has a 4xx status code
func (o *PatchRoutingConversationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation unauthorized response has a 5xx status code
func (o *PatchRoutingConversationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation unauthorized response a status code equal to that given
func (o *PatchRoutingConversationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchRoutingConversationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingConversationUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingConversationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationForbidden creates a PatchRoutingConversationForbidden with default headers values
func NewPatchRoutingConversationForbidden() *PatchRoutingConversationForbidden {
	return &PatchRoutingConversationForbidden{}
}

/*
PatchRoutingConversationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingConversationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation forbidden response has a 2xx status code
func (o *PatchRoutingConversationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation forbidden response has a 3xx status code
func (o *PatchRoutingConversationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation forbidden response has a 4xx status code
func (o *PatchRoutingConversationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation forbidden response has a 5xx status code
func (o *PatchRoutingConversationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation forbidden response a status code equal to that given
func (o *PatchRoutingConversationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchRoutingConversationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingConversationForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingConversationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationNotFound creates a PatchRoutingConversationNotFound with default headers values
func NewPatchRoutingConversationNotFound() *PatchRoutingConversationNotFound {
	return &PatchRoutingConversationNotFound{}
}

/*
PatchRoutingConversationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchRoutingConversationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation not found response has a 2xx status code
func (o *PatchRoutingConversationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation not found response has a 3xx status code
func (o *PatchRoutingConversationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation not found response has a 4xx status code
func (o *PatchRoutingConversationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation not found response has a 5xx status code
func (o *PatchRoutingConversationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation not found response a status code equal to that given
func (o *PatchRoutingConversationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchRoutingConversationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingConversationNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingConversationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationRequestTimeout creates a PatchRoutingConversationRequestTimeout with default headers values
func NewPatchRoutingConversationRequestTimeout() *PatchRoutingConversationRequestTimeout {
	return &PatchRoutingConversationRequestTimeout{}
}

/*
PatchRoutingConversationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchRoutingConversationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation request timeout response has a 2xx status code
func (o *PatchRoutingConversationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation request timeout response has a 3xx status code
func (o *PatchRoutingConversationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation request timeout response has a 4xx status code
func (o *PatchRoutingConversationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation request timeout response has a 5xx status code
func (o *PatchRoutingConversationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation request timeout response a status code equal to that given
func (o *PatchRoutingConversationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchRoutingConversationRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchRoutingConversationRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchRoutingConversationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationRequestEntityTooLarge creates a PatchRoutingConversationRequestEntityTooLarge with default headers values
func NewPatchRoutingConversationRequestEntityTooLarge() *PatchRoutingConversationRequestEntityTooLarge {
	return &PatchRoutingConversationRequestEntityTooLarge{}
}

/*
PatchRoutingConversationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchRoutingConversationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation request entity too large response has a 2xx status code
func (o *PatchRoutingConversationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation request entity too large response has a 3xx status code
func (o *PatchRoutingConversationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation request entity too large response has a 4xx status code
func (o *PatchRoutingConversationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation request entity too large response has a 5xx status code
func (o *PatchRoutingConversationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation request entity too large response a status code equal to that given
func (o *PatchRoutingConversationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchRoutingConversationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingConversationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingConversationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationUnsupportedMediaType creates a PatchRoutingConversationUnsupportedMediaType with default headers values
func NewPatchRoutingConversationUnsupportedMediaType() *PatchRoutingConversationUnsupportedMediaType {
	return &PatchRoutingConversationUnsupportedMediaType{}
}

/*
PatchRoutingConversationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingConversationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation unsupported media type response has a 2xx status code
func (o *PatchRoutingConversationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation unsupported media type response has a 3xx status code
func (o *PatchRoutingConversationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation unsupported media type response has a 4xx status code
func (o *PatchRoutingConversationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation unsupported media type response has a 5xx status code
func (o *PatchRoutingConversationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation unsupported media type response a status code equal to that given
func (o *PatchRoutingConversationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchRoutingConversationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingConversationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingConversationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationTooManyRequests creates a PatchRoutingConversationTooManyRequests with default headers values
func NewPatchRoutingConversationTooManyRequests() *PatchRoutingConversationTooManyRequests {
	return &PatchRoutingConversationTooManyRequests{}
}

/*
PatchRoutingConversationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchRoutingConversationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation too many requests response has a 2xx status code
func (o *PatchRoutingConversationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation too many requests response has a 3xx status code
func (o *PatchRoutingConversationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation too many requests response has a 4xx status code
func (o *PatchRoutingConversationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch routing conversation too many requests response has a 5xx status code
func (o *PatchRoutingConversationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch routing conversation too many requests response a status code equal to that given
func (o *PatchRoutingConversationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchRoutingConversationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingConversationTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingConversationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationInternalServerError creates a PatchRoutingConversationInternalServerError with default headers values
func NewPatchRoutingConversationInternalServerError() *PatchRoutingConversationInternalServerError {
	return &PatchRoutingConversationInternalServerError{}
}

/*
PatchRoutingConversationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingConversationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation internal server error response has a 2xx status code
func (o *PatchRoutingConversationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation internal server error response has a 3xx status code
func (o *PatchRoutingConversationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation internal server error response has a 4xx status code
func (o *PatchRoutingConversationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch routing conversation internal server error response has a 5xx status code
func (o *PatchRoutingConversationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch routing conversation internal server error response a status code equal to that given
func (o *PatchRoutingConversationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchRoutingConversationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingConversationInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingConversationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationServiceUnavailable creates a PatchRoutingConversationServiceUnavailable with default headers values
func NewPatchRoutingConversationServiceUnavailable() *PatchRoutingConversationServiceUnavailable {
	return &PatchRoutingConversationServiceUnavailable{}
}

/*
PatchRoutingConversationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingConversationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation service unavailable response has a 2xx status code
func (o *PatchRoutingConversationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation service unavailable response has a 3xx status code
func (o *PatchRoutingConversationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation service unavailable response has a 4xx status code
func (o *PatchRoutingConversationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch routing conversation service unavailable response has a 5xx status code
func (o *PatchRoutingConversationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch routing conversation service unavailable response a status code equal to that given
func (o *PatchRoutingConversationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchRoutingConversationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingConversationServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingConversationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingConversationGatewayTimeout creates a PatchRoutingConversationGatewayTimeout with default headers values
func NewPatchRoutingConversationGatewayTimeout() *PatchRoutingConversationGatewayTimeout {
	return &PatchRoutingConversationGatewayTimeout{}
}

/*
PatchRoutingConversationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchRoutingConversationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch routing conversation gateway timeout response has a 2xx status code
func (o *PatchRoutingConversationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch routing conversation gateway timeout response has a 3xx status code
func (o *PatchRoutingConversationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch routing conversation gateway timeout response has a 4xx status code
func (o *PatchRoutingConversationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch routing conversation gateway timeout response has a 5xx status code
func (o *PatchRoutingConversationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch routing conversation gateway timeout response a status code equal to that given
func (o *PatchRoutingConversationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchRoutingConversationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingConversationGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/conversations/{conversationId}][%d] patchRoutingConversationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingConversationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingConversationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
