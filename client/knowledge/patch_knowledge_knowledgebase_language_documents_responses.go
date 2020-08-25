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

// PatchKnowledgeKnowledgebaseLanguageDocumentsReader is a Reader for the PatchKnowledgeKnowledgebaseLanguageDocuments structure.
type PatchKnowledgeKnowledgebaseLanguageDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsOK creates a PatchKnowledgeKnowledgebaseLanguageDocumentsOK with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsOK() *PatchKnowledgeKnowledgebaseLanguageDocumentsOK {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsOK{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsOK handles this case with default header values.

successful operation
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsOK struct {
	Payload *models.DocumentListing
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsOK  %+v", 200, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) GetPayload() *models.DocumentListing {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest creates a PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest() *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized creates a PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized() *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsForbidden creates a PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsForbidden() *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsNotFound creates a PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsNotFound() *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge creates a PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge() *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType creates a PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType() *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests creates a PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests() *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError creates a PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError() *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable creates a PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable() *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout creates a PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout with default headers values
func NewPatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout() *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout {
	return &PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout{}
}

/*PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] patchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}