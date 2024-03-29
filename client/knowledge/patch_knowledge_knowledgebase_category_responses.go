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

// PatchKnowledgeKnowledgebaseCategoryReader is a Reader for the PatchKnowledgeKnowledgebaseCategory structure.
type PatchKnowledgeKnowledgebaseCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseCategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseCategoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchKnowledgeKnowledgebaseCategoryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseCategoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseCategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseCategoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseCategoryOK creates a PatchKnowledgeKnowledgebaseCategoryOK with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryOK() *PatchKnowledgeKnowledgebaseCategoryOK {
	return &PatchKnowledgeKnowledgebaseCategoryOK{}
}

/*
PatchKnowledgeKnowledgebaseCategoryOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseCategoryOK struct {
	Payload *models.CategoryResponse
}

// IsSuccess returns true when this patch knowledge knowledgebase category o k response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch knowledge knowledgebase category o k response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category o k response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase category o k response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category o k response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchKnowledgeKnowledgebaseCategoryOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryOK) GetPayload() *models.CategoryResponse {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryBadRequest creates a PatchKnowledgeKnowledgebaseCategoryBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryBadRequest() *PatchKnowledgeKnowledgebaseCategoryBadRequest {
	return &PatchKnowledgeKnowledgebaseCategoryBadRequest{}
}

/*
PatchKnowledgeKnowledgebaseCategoryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseCategoryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category bad request response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category bad request response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category bad request response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category bad request response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category bad request response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryUnauthorized creates a PatchKnowledgeKnowledgebaseCategoryUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryUnauthorized() *PatchKnowledgeKnowledgebaseCategoryUnauthorized {
	return &PatchKnowledgeKnowledgebaseCategoryUnauthorized{}
}

/*
PatchKnowledgeKnowledgebaseCategoryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseCategoryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category unauthorized response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category unauthorized response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category unauthorized response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category unauthorized response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category unauthorized response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryForbidden creates a PatchKnowledgeKnowledgebaseCategoryForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryForbidden() *PatchKnowledgeKnowledgebaseCategoryForbidden {
	return &PatchKnowledgeKnowledgebaseCategoryForbidden{}
}

/*
PatchKnowledgeKnowledgebaseCategoryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseCategoryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category forbidden response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category forbidden response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category forbidden response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category forbidden response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category forbidden response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryNotFound creates a PatchKnowledgeKnowledgebaseCategoryNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryNotFound() *PatchKnowledgeKnowledgebaseCategoryNotFound {
	return &PatchKnowledgeKnowledgebaseCategoryNotFound{}
}

/*
PatchKnowledgeKnowledgebaseCategoryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseCategoryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category not found response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category not found response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category not found response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category not found response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category not found response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryRequestTimeout creates a PatchKnowledgeKnowledgebaseCategoryRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryRequestTimeout() *PatchKnowledgeKnowledgebaseCategoryRequestTimeout {
	return &PatchKnowledgeKnowledgebaseCategoryRequestTimeout{}
}

/*
PatchKnowledgeKnowledgebaseCategoryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseCategoryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category request timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category request timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category request timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category request timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category request timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryConflict creates a PatchKnowledgeKnowledgebaseCategoryConflict with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryConflict() *PatchKnowledgeKnowledgebaseCategoryConflict {
	return &PatchKnowledgeKnowledgebaseCategoryConflict{}
}

/*
PatchKnowledgeKnowledgebaseCategoryConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchKnowledgeKnowledgebaseCategoryConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category conflict response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category conflict response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category conflict response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category conflict response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category conflict response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PatchKnowledgeKnowledgebaseCategoryConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryConflict  %+v", 409, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryConflict) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryConflict  %+v", 409, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge{}
}

/*
PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category request entity too large response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category request entity too large response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category request entity too large response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category request entity too large response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category request entity too large response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType() *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType{}
}

/*
PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category unsupported media type response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category unsupported media type response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category unsupported media type response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category unsupported media type response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category unsupported media type response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryTooManyRequests creates a PatchKnowledgeKnowledgebaseCategoryTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryTooManyRequests() *PatchKnowledgeKnowledgebaseCategoryTooManyRequests {
	return &PatchKnowledgeKnowledgebaseCategoryTooManyRequests{}
}

/*
PatchKnowledgeKnowledgebaseCategoryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseCategoryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category too many requests response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category too many requests response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category too many requests response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch knowledge knowledgebase category too many requests response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch knowledge knowledgebase category too many requests response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryInternalServerError creates a PatchKnowledgeKnowledgebaseCategoryInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryInternalServerError() *PatchKnowledgeKnowledgebaseCategoryInternalServerError {
	return &PatchKnowledgeKnowledgebaseCategoryInternalServerError{}
}

/*
PatchKnowledgeKnowledgebaseCategoryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseCategoryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category internal server error response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category internal server error response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category internal server error response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase category internal server error response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase category internal server error response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryServiceUnavailable creates a PatchKnowledgeKnowledgebaseCategoryServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryServiceUnavailable() *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseCategoryServiceUnavailable{}
}

/*
PatchKnowledgeKnowledgebaseCategoryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseCategoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category service unavailable response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category service unavailable response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category service unavailable response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase category service unavailable response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase category service unavailable response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseCategoryGatewayTimeout creates a PatchKnowledgeKnowledgebaseCategoryGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseCategoryGatewayTimeout() *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseCategoryGatewayTimeout{}
}

/*
PatchKnowledgeKnowledgebaseCategoryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseCategoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch knowledge knowledgebase category gateway timeout response has a 2xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch knowledge knowledgebase category gateway timeout response has a 3xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch knowledge knowledgebase category gateway timeout response has a 4xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch knowledge knowledgebase category gateway timeout response has a 5xx status code
func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch knowledge knowledgebase category gateway timeout response a status code equal to that given
func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
