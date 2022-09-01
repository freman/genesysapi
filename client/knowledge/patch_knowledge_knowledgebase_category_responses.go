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

/*PatchKnowledgeKnowledgebaseCategoryOK handles this case with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseCategoryOK struct {
	Payload *models.CategoryResponse
}

func (o *PatchKnowledgeKnowledgebaseCategoryOK) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseCategoryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryBadRequest) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseCategoryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnauthorized) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseCategoryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryForbidden) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseCategoryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryNotFound) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseCategoryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestTimeout) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryConflict handles this case with default header values.

Conflict
*/
type PatchKnowledgeKnowledgebaseCategoryConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryConflict) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryUnsupportedMediaType) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseCategoryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryTooManyRequests) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseCategoryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryInternalServerError) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseCategoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryServiceUnavailable) Error() string {
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

/*PatchKnowledgeKnowledgebaseCategoryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseCategoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseCategoryGatewayTimeout) Error() string {
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
