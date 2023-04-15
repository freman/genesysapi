// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchKnowledgeGuestSessionDocumentsSearchSearchIDReader is a Reader for the PatchKnowledgeGuestSessionDocumentsSearchSearchID structure.
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent describes a response with status code 204, with default header values.

Search updated successfully.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent struct {
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id no content response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id no content response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id no content response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge guest session documents search search Id no content response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id no content response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdNoContent ", 204)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdNoContent ", 204)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id bad request response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id bad request response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id bad request response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id bad request response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id bad request response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id unauthorized response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id unauthorized response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id unauthorized response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id unauthorized response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id unauthorized response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id forbidden response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id forbidden response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id forbidden response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id forbidden response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id forbidden response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id not found response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id not found response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id not found response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id not found response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id not found response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id request timeout response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id request timeout response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id request timeout response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id request timeout response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id request timeout response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id request entity too large response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id request entity too large response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id request entity too large response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id request entity too large response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id request entity too large response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id unsupported media type response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id unsupported media type response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id unsupported media type response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id unsupported media type response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id unsupported media type response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id too many requests response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id too many requests response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id too many requests response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge guest session documents search search Id too many requests response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge guest session documents search search Id too many requests response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id internal server error response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id internal server error response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id internal server error response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge guest session documents search search Id internal server error response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge guest session documents search search Id internal server error response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id service unavailable response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id service unavailable response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id service unavailable response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge guest session documents search search Id service unavailable response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge guest session documents search search Id service unavailable response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout creates a PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout with default headers values
func NewPatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout() *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout {
	return &PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout{}
}

/*
PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge guest session documents search search Id gateway timeout response has a 2xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge guest session documents search search Id gateway timeout response has a 3xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge guest session documents search search Id gateway timeout response has a 4xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge guest session documents search search Id gateway timeout response has a 5xx status code
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge guest session documents search search Id gateway timeout response a status code equal to that given
func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/guest/sessions/{sessionId}/documents/search/{searchId}][%d] patchKnowledgeGuestSessionDocumentsSearchSearchIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeGuestSessionDocumentsSearchSearchIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
