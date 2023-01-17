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

// PatchConversationSecureattributesReader is a Reader for the PatchConversationSecureattributes structure.
type PatchConversationSecureattributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationSecureattributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchConversationSecureattributesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationSecureattributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationSecureattributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationSecureattributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationSecureattributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchConversationSecureattributesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchConversationSecureattributesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationSecureattributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationSecureattributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationSecureattributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationSecureattributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationSecureattributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationSecureattributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationSecureattributesAccepted creates a PatchConversationSecureattributesAccepted with default headers values
func NewPatchConversationSecureattributesAccepted() *PatchConversationSecureattributesAccepted {
	return &PatchConversationSecureattributesAccepted{}
}

/*
PatchConversationSecureattributesAccepted describes a response with status code 202, with default header values.

The secure attributes update request was accepted.
*/
type PatchConversationSecureattributesAccepted struct {
	Payload string
}

// IsSuccess returns true when this patch conversation secureattributes accepted response has a 2xx status code
func (o *PatchConversationSecureattributesAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch conversation secureattributes accepted response has a 3xx status code
func (o *PatchConversationSecureattributesAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes accepted response has a 4xx status code
func (o *PatchConversationSecureattributesAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation secureattributes accepted response has a 5xx status code
func (o *PatchConversationSecureattributesAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes accepted response a status code equal to that given
func (o *PatchConversationSecureattributesAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PatchConversationSecureattributesAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesAccepted  %+v", 202, o.Payload)
}

func (o *PatchConversationSecureattributesAccepted) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesAccepted  %+v", 202, o.Payload)
}

func (o *PatchConversationSecureattributesAccepted) GetPayload() string {
	return o.Payload
}

func (o *PatchConversationSecureattributesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesBadRequest creates a PatchConversationSecureattributesBadRequest with default headers values
func NewPatchConversationSecureattributesBadRequest() *PatchConversationSecureattributesBadRequest {
	return &PatchConversationSecureattributesBadRequest{}
}

/*
PatchConversationSecureattributesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationSecureattributesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes bad request response has a 2xx status code
func (o *PatchConversationSecureattributesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes bad request response has a 3xx status code
func (o *PatchConversationSecureattributesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes bad request response has a 4xx status code
func (o *PatchConversationSecureattributesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes bad request response has a 5xx status code
func (o *PatchConversationSecureattributesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes bad request response a status code equal to that given
func (o *PatchConversationSecureattributesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchConversationSecureattributesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationSecureattributesBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationSecureattributesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesUnauthorized creates a PatchConversationSecureattributesUnauthorized with default headers values
func NewPatchConversationSecureattributesUnauthorized() *PatchConversationSecureattributesUnauthorized {
	return &PatchConversationSecureattributesUnauthorized{}
}

/*
PatchConversationSecureattributesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationSecureattributesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes unauthorized response has a 2xx status code
func (o *PatchConversationSecureattributesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes unauthorized response has a 3xx status code
func (o *PatchConversationSecureattributesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes unauthorized response has a 4xx status code
func (o *PatchConversationSecureattributesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes unauthorized response has a 5xx status code
func (o *PatchConversationSecureattributesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes unauthorized response a status code equal to that given
func (o *PatchConversationSecureattributesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchConversationSecureattributesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationSecureattributesUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationSecureattributesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesForbidden creates a PatchConversationSecureattributesForbidden with default headers values
func NewPatchConversationSecureattributesForbidden() *PatchConversationSecureattributesForbidden {
	return &PatchConversationSecureattributesForbidden{}
}

/*
PatchConversationSecureattributesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationSecureattributesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes forbidden response has a 2xx status code
func (o *PatchConversationSecureattributesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes forbidden response has a 3xx status code
func (o *PatchConversationSecureattributesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes forbidden response has a 4xx status code
func (o *PatchConversationSecureattributesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes forbidden response has a 5xx status code
func (o *PatchConversationSecureattributesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes forbidden response a status code equal to that given
func (o *PatchConversationSecureattributesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchConversationSecureattributesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationSecureattributesForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationSecureattributesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesNotFound creates a PatchConversationSecureattributesNotFound with default headers values
func NewPatchConversationSecureattributesNotFound() *PatchConversationSecureattributesNotFound {
	return &PatchConversationSecureattributesNotFound{}
}

/*
PatchConversationSecureattributesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchConversationSecureattributesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes not found response has a 2xx status code
func (o *PatchConversationSecureattributesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes not found response has a 3xx status code
func (o *PatchConversationSecureattributesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes not found response has a 4xx status code
func (o *PatchConversationSecureattributesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes not found response has a 5xx status code
func (o *PatchConversationSecureattributesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes not found response a status code equal to that given
func (o *PatchConversationSecureattributesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchConversationSecureattributesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationSecureattributesNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationSecureattributesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesRequestTimeout creates a PatchConversationSecureattributesRequestTimeout with default headers values
func NewPatchConversationSecureattributesRequestTimeout() *PatchConversationSecureattributesRequestTimeout {
	return &PatchConversationSecureattributesRequestTimeout{}
}

/*
PatchConversationSecureattributesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchConversationSecureattributesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes request timeout response has a 2xx status code
func (o *PatchConversationSecureattributesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes request timeout response has a 3xx status code
func (o *PatchConversationSecureattributesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes request timeout response has a 4xx status code
func (o *PatchConversationSecureattributesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes request timeout response has a 5xx status code
func (o *PatchConversationSecureattributesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes request timeout response a status code equal to that given
func (o *PatchConversationSecureattributesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchConversationSecureattributesRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationSecureattributesRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchConversationSecureattributesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesConflict creates a PatchConversationSecureattributesConflict with default headers values
func NewPatchConversationSecureattributesConflict() *PatchConversationSecureattributesConflict {
	return &PatchConversationSecureattributesConflict{}
}

/*
PatchConversationSecureattributesConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchConversationSecureattributesConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes conflict response has a 2xx status code
func (o *PatchConversationSecureattributesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes conflict response has a 3xx status code
func (o *PatchConversationSecureattributesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes conflict response has a 4xx status code
func (o *PatchConversationSecureattributesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes conflict response has a 5xx status code
func (o *PatchConversationSecureattributesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes conflict response a status code equal to that given
func (o *PatchConversationSecureattributesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PatchConversationSecureattributesConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesConflict  %+v", 409, o.Payload)
}

func (o *PatchConversationSecureattributesConflict) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesConflict  %+v", 409, o.Payload)
}

func (o *PatchConversationSecureattributesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesRequestEntityTooLarge creates a PatchConversationSecureattributesRequestEntityTooLarge with default headers values
func NewPatchConversationSecureattributesRequestEntityTooLarge() *PatchConversationSecureattributesRequestEntityTooLarge {
	return &PatchConversationSecureattributesRequestEntityTooLarge{}
}

/*
PatchConversationSecureattributesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchConversationSecureattributesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes request entity too large response has a 2xx status code
func (o *PatchConversationSecureattributesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes request entity too large response has a 3xx status code
func (o *PatchConversationSecureattributesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes request entity too large response has a 4xx status code
func (o *PatchConversationSecureattributesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes request entity too large response has a 5xx status code
func (o *PatchConversationSecureattributesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes request entity too large response a status code equal to that given
func (o *PatchConversationSecureattributesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchConversationSecureattributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationSecureattributesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationSecureattributesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesUnsupportedMediaType creates a PatchConversationSecureattributesUnsupportedMediaType with default headers values
func NewPatchConversationSecureattributesUnsupportedMediaType() *PatchConversationSecureattributesUnsupportedMediaType {
	return &PatchConversationSecureattributesUnsupportedMediaType{}
}

/*
PatchConversationSecureattributesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationSecureattributesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes unsupported media type response has a 2xx status code
func (o *PatchConversationSecureattributesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes unsupported media type response has a 3xx status code
func (o *PatchConversationSecureattributesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes unsupported media type response has a 4xx status code
func (o *PatchConversationSecureattributesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes unsupported media type response has a 5xx status code
func (o *PatchConversationSecureattributesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes unsupported media type response a status code equal to that given
func (o *PatchConversationSecureattributesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchConversationSecureattributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationSecureattributesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationSecureattributesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesTooManyRequests creates a PatchConversationSecureattributesTooManyRequests with default headers values
func NewPatchConversationSecureattributesTooManyRequests() *PatchConversationSecureattributesTooManyRequests {
	return &PatchConversationSecureattributesTooManyRequests{}
}

/*
PatchConversationSecureattributesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchConversationSecureattributesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes too many requests response has a 2xx status code
func (o *PatchConversationSecureattributesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes too many requests response has a 3xx status code
func (o *PatchConversationSecureattributesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes too many requests response has a 4xx status code
func (o *PatchConversationSecureattributesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch conversation secureattributes too many requests response has a 5xx status code
func (o *PatchConversationSecureattributesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch conversation secureattributes too many requests response a status code equal to that given
func (o *PatchConversationSecureattributesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchConversationSecureattributesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationSecureattributesTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationSecureattributesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesInternalServerError creates a PatchConversationSecureattributesInternalServerError with default headers values
func NewPatchConversationSecureattributesInternalServerError() *PatchConversationSecureattributesInternalServerError {
	return &PatchConversationSecureattributesInternalServerError{}
}

/*
PatchConversationSecureattributesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationSecureattributesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes internal server error response has a 2xx status code
func (o *PatchConversationSecureattributesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes internal server error response has a 3xx status code
func (o *PatchConversationSecureattributesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes internal server error response has a 4xx status code
func (o *PatchConversationSecureattributesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation secureattributes internal server error response has a 5xx status code
func (o *PatchConversationSecureattributesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversation secureattributes internal server error response a status code equal to that given
func (o *PatchConversationSecureattributesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchConversationSecureattributesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationSecureattributesInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationSecureattributesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesServiceUnavailable creates a PatchConversationSecureattributesServiceUnavailable with default headers values
func NewPatchConversationSecureattributesServiceUnavailable() *PatchConversationSecureattributesServiceUnavailable {
	return &PatchConversationSecureattributesServiceUnavailable{}
}

/*
PatchConversationSecureattributesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationSecureattributesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes service unavailable response has a 2xx status code
func (o *PatchConversationSecureattributesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes service unavailable response has a 3xx status code
func (o *PatchConversationSecureattributesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes service unavailable response has a 4xx status code
func (o *PatchConversationSecureattributesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation secureattributes service unavailable response has a 5xx status code
func (o *PatchConversationSecureattributesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversation secureattributes service unavailable response a status code equal to that given
func (o *PatchConversationSecureattributesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchConversationSecureattributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationSecureattributesServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationSecureattributesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationSecureattributesGatewayTimeout creates a PatchConversationSecureattributesGatewayTimeout with default headers values
func NewPatchConversationSecureattributesGatewayTimeout() *PatchConversationSecureattributesGatewayTimeout {
	return &PatchConversationSecureattributesGatewayTimeout{}
}

/*
PatchConversationSecureattributesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchConversationSecureattributesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch conversation secureattributes gateway timeout response has a 2xx status code
func (o *PatchConversationSecureattributesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch conversation secureattributes gateway timeout response has a 3xx status code
func (o *PatchConversationSecureattributesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch conversation secureattributes gateway timeout response has a 4xx status code
func (o *PatchConversationSecureattributesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch conversation secureattributes gateway timeout response has a 5xx status code
func (o *PatchConversationSecureattributesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch conversation secureattributes gateway timeout response a status code equal to that given
func (o *PatchConversationSecureattributesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchConversationSecureattributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationSecureattributesGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/{conversationId}/secureattributes][%d] patchConversationSecureattributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationSecureattributesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationSecureattributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
