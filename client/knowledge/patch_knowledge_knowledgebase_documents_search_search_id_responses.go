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

// PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDReader is a Reader for the PatchKnowledgeKnowledgebaseDocumentsSearchSearchID structure.
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent handles this case with default header values.

Search updated successfully.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent struct {
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdNoContent ", 204)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout creates a PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout() *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout{}
}

/*PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/search/{searchId}][%d] patchKnowledgeKnowledgebaseDocumentsSearchSearchIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseDocumentsSearchSearchIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}