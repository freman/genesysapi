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

// PatchKnowledgeKnowledgebaseLanguageCategoryReader is a Reader for the PatchKnowledgeKnowledgebaseLanguageCategory structure.
type PatchKnowledgeKnowledgebaseLanguageCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryOK creates a PatchKnowledgeKnowledgebaseLanguageCategoryOK with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryOK() *PatchKnowledgeKnowledgebaseLanguageCategoryOK {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryOK{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryOK handles this case with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryOK struct {
	Payload *models.KnowledgeExtendedCategory
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) GetPayload() *models.KnowledgeExtendedCategory {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeExtendedCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryBadRequest creates a PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryBadRequest() *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized creates a PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized() *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryForbidden creates a PatchKnowledgeKnowledgebaseLanguageCategoryForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryForbidden() *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryForbidden{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryNotFound creates a PatchKnowledgeKnowledgebaseLanguageCategoryNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryNotFound() *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryNotFound{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType() *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests creates a PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests() *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError creates a PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError() *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable creates a PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable() *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout creates a PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout() *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout{}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/categories/{categoryId}][%d] patchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageCategoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
