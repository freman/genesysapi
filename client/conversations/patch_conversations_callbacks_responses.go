// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchConversationsCallbacksReader is a Reader for the PatchConversationsCallbacks structure.
type PatchConversationsCallbacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsCallbacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsCallbacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsCallbacksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsCallbacksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsCallbacksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsCallbacksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationsCallbacksRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPatchConversationsCallbacksPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsCallbacksRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsCallbacksUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsCallbacksTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsCallbacksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsCallbacksServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsCallbacksGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsCallbacksOK creates a PatchConversationsCallbacksOK with default headers values
func NewPatchConversationsCallbacksOK() *PatchConversationsCallbacksOK {
	return &PatchConversationsCallbacksOK{}
}

/*
PatchConversationsCallbacksOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchConversationsCallbacksOK struct {
	Payload *models.PatchCallbackResponse
}

// IsSuccess returns true when this patch conversations callbacks o k response has a 2xx status code
func (o *PatchConversationsCallbacksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversations callbacks o k response has a 3xx status code
func (o *PatchConversationsCallbacksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks o k response has a 4xx status code
func (o *PatchConversationsCallbacksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations callbacks o k response has a 5xx status code
func (o *PatchConversationsCallbacksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks o k response a status code equal to that given
func (o *PatchConversationsCallbacksOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchConversationsCallbacksOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallbacksOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsCallbacksOK) GetPayload() *models.PatchCallbackResponse {
	return o.Payload
}

func (o *PatchConversationsCallbacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchCallbackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksBadRequest creates a PatchConversationsCallbacksBadRequest with default headers values
func NewPatchConversationsCallbacksBadRequest() *PatchConversationsCallbacksBadRequest {
	return &PatchConversationsCallbacksBadRequest{}
}

/*
PatchConversationsCallbacksBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsCallbacksBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks bad request response has a 2xx status code
func (o *PatchConversationsCallbacksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks bad request response has a 3xx status code
func (o *PatchConversationsCallbacksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks bad request response has a 4xx status code
func (o *PatchConversationsCallbacksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks bad request response has a 5xx status code
func (o *PatchConversationsCallbacksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks bad request response a status code equal to that given
func (o *PatchConversationsCallbacksBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationsCallbacksBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallbacksBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsCallbacksBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksUnauthorized creates a PatchConversationsCallbacksUnauthorized with default headers values
func NewPatchConversationsCallbacksUnauthorized() *PatchConversationsCallbacksUnauthorized {
	return &PatchConversationsCallbacksUnauthorized{}
}

/*
PatchConversationsCallbacksUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsCallbacksUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks unauthorized response has a 2xx status code
func (o *PatchConversationsCallbacksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks unauthorized response has a 3xx status code
func (o *PatchConversationsCallbacksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks unauthorized response has a 4xx status code
func (o *PatchConversationsCallbacksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks unauthorized response has a 5xx status code
func (o *PatchConversationsCallbacksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks unauthorized response a status code equal to that given
func (o *PatchConversationsCallbacksUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationsCallbacksUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallbacksUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsCallbacksUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksForbidden creates a PatchConversationsCallbacksForbidden with default headers values
func NewPatchConversationsCallbacksForbidden() *PatchConversationsCallbacksForbidden {
	return &PatchConversationsCallbacksForbidden{}
}

/*
PatchConversationsCallbacksForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsCallbacksForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks forbidden response has a 2xx status code
func (o *PatchConversationsCallbacksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks forbidden response has a 3xx status code
func (o *PatchConversationsCallbacksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks forbidden response has a 4xx status code
func (o *PatchConversationsCallbacksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks forbidden response has a 5xx status code
func (o *PatchConversationsCallbacksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks forbidden response a status code equal to that given
func (o *PatchConversationsCallbacksForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationsCallbacksForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallbacksForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsCallbacksForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksNotFound creates a PatchConversationsCallbacksNotFound with default headers values
func NewPatchConversationsCallbacksNotFound() *PatchConversationsCallbacksNotFound {
	return &PatchConversationsCallbacksNotFound{}
}

/*
PatchConversationsCallbacksNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationsCallbacksNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks not found response has a 2xx status code
func (o *PatchConversationsCallbacksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks not found response has a 3xx status code
func (o *PatchConversationsCallbacksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks not found response has a 4xx status code
func (o *PatchConversationsCallbacksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks not found response has a 5xx status code
func (o *PatchConversationsCallbacksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks not found response a status code equal to that given
func (o *PatchConversationsCallbacksNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationsCallbacksNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallbacksNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsCallbacksNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksRequestTimeout creates a PatchConversationsCallbacksRequestTimeout with default headers values
func NewPatchConversationsCallbacksRequestTimeout() *PatchConversationsCallbacksRequestTimeout {
	return &PatchConversationsCallbacksRequestTimeout{}
}

/*
PatchConversationsCallbacksRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationsCallbacksRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks request timeout response has a 2xx status code
func (o *PatchConversationsCallbacksRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks request timeout response has a 3xx status code
func (o *PatchConversationsCallbacksRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks request timeout response has a 4xx status code
func (o *PatchConversationsCallbacksRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks request timeout response has a 5xx status code
func (o *PatchConversationsCallbacksRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks request timeout response a status code equal to that given
func (o *PatchConversationsCallbacksRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationsCallbacksRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallbacksRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationsCallbacksRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksPreconditionFailed creates a PatchConversationsCallbacksPreconditionFailed with default headers values
func NewPatchConversationsCallbacksPreconditionFailed() *PatchConversationsCallbacksPreconditionFailed {
	return &PatchConversationsCallbacksPreconditionFailed{}
}

/*
PatchConversationsCallbacksPreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type PatchConversationsCallbacksPreconditionFailed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks precondition failed response has a 2xx status code
func (o *PatchConversationsCallbacksPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks precondition failed response has a 3xx status code
func (o *PatchConversationsCallbacksPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks precondition failed response has a 4xx status code
func (o *PatchConversationsCallbacksPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks precondition failed response has a 5xx status code
func (o *PatchConversationsCallbacksPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks precondition failed response a status code equal to that given
func (o *PatchConversationsCallbacksPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

func (o *PatchConversationsCallbacksPreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PatchConversationsCallbacksPreconditionFailed) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PatchConversationsCallbacksPreconditionFailed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksRequestEntityTooLarge creates a PatchConversationsCallbacksRequestEntityTooLarge with default headers values
func NewPatchConversationsCallbacksRequestEntityTooLarge() *PatchConversationsCallbacksRequestEntityTooLarge {
	return &PatchConversationsCallbacksRequestEntityTooLarge{}
}

/*
PatchConversationsCallbacksRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchConversationsCallbacksRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks request entity too large response has a 2xx status code
func (o *PatchConversationsCallbacksRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks request entity too large response has a 3xx status code
func (o *PatchConversationsCallbacksRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks request entity too large response has a 4xx status code
func (o *PatchConversationsCallbacksRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks request entity too large response has a 5xx status code
func (o *PatchConversationsCallbacksRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks request entity too large response a status code equal to that given
func (o *PatchConversationsCallbacksRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationsCallbacksRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallbacksRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsCallbacksRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksUnsupportedMediaType creates a PatchConversationsCallbacksUnsupportedMediaType with default headers values
func NewPatchConversationsCallbacksUnsupportedMediaType() *PatchConversationsCallbacksUnsupportedMediaType {
	return &PatchConversationsCallbacksUnsupportedMediaType{}
}

/*
PatchConversationsCallbacksUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsCallbacksUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks unsupported media type response has a 2xx status code
func (o *PatchConversationsCallbacksUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks unsupported media type response has a 3xx status code
func (o *PatchConversationsCallbacksUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks unsupported media type response has a 4xx status code
func (o *PatchConversationsCallbacksUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks unsupported media type response has a 5xx status code
func (o *PatchConversationsCallbacksUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks unsupported media type response a status code equal to that given
func (o *PatchConversationsCallbacksUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationsCallbacksUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallbacksUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsCallbacksUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksTooManyRequests creates a PatchConversationsCallbacksTooManyRequests with default headers values
func NewPatchConversationsCallbacksTooManyRequests() *PatchConversationsCallbacksTooManyRequests {
	return &PatchConversationsCallbacksTooManyRequests{}
}

/*
PatchConversationsCallbacksTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationsCallbacksTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks too many requests response has a 2xx status code
func (o *PatchConversationsCallbacksTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks too many requests response has a 3xx status code
func (o *PatchConversationsCallbacksTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks too many requests response has a 4xx status code
func (o *PatchConversationsCallbacksTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversations callbacks too many requests response has a 5xx status code
func (o *PatchConversationsCallbacksTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversations callbacks too many requests response a status code equal to that given
func (o *PatchConversationsCallbacksTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationsCallbacksTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallbacksTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsCallbacksTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksInternalServerError creates a PatchConversationsCallbacksInternalServerError with default headers values
func NewPatchConversationsCallbacksInternalServerError() *PatchConversationsCallbacksInternalServerError {
	return &PatchConversationsCallbacksInternalServerError{}
}

/*
PatchConversationsCallbacksInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsCallbacksInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks internal server error response has a 2xx status code
func (o *PatchConversationsCallbacksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks internal server error response has a 3xx status code
func (o *PatchConversationsCallbacksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks internal server error response has a 4xx status code
func (o *PatchConversationsCallbacksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations callbacks internal server error response has a 5xx status code
func (o *PatchConversationsCallbacksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations callbacks internal server error response a status code equal to that given
func (o *PatchConversationsCallbacksInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationsCallbacksInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallbacksInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsCallbacksInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksServiceUnavailable creates a PatchConversationsCallbacksServiceUnavailable with default headers values
func NewPatchConversationsCallbacksServiceUnavailable() *PatchConversationsCallbacksServiceUnavailable {
	return &PatchConversationsCallbacksServiceUnavailable{}
}

/*
PatchConversationsCallbacksServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsCallbacksServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks service unavailable response has a 2xx status code
func (o *PatchConversationsCallbacksServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks service unavailable response has a 3xx status code
func (o *PatchConversationsCallbacksServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks service unavailable response has a 4xx status code
func (o *PatchConversationsCallbacksServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations callbacks service unavailable response has a 5xx status code
func (o *PatchConversationsCallbacksServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations callbacks service unavailable response a status code equal to that given
func (o *PatchConversationsCallbacksServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationsCallbacksServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallbacksServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsCallbacksServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsCallbacksGatewayTimeout creates a PatchConversationsCallbacksGatewayTimeout with default headers values
func NewPatchConversationsCallbacksGatewayTimeout() *PatchConversationsCallbacksGatewayTimeout {
	return &PatchConversationsCallbacksGatewayTimeout{}
}

/*
PatchConversationsCallbacksGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationsCallbacksGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversations callbacks gateway timeout response has a 2xx status code
func (o *PatchConversationsCallbacksGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversations callbacks gateway timeout response has a 3xx status code
func (o *PatchConversationsCallbacksGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversations callbacks gateway timeout response has a 4xx status code
func (o *PatchConversationsCallbacksGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversations callbacks gateway timeout response has a 5xx status code
func (o *PatchConversationsCallbacksGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversations callbacks gateway timeout response a status code equal to that given
func (o *PatchConversationsCallbacksGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationsCallbacksGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallbacksGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/callbacks][%d] patchConversationsCallbacksGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsCallbacksGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsCallbacksGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
