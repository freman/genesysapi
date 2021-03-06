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

// DeleteKnowledgeKnowledgebaseLanguageDocumentReader is a Reader for the DeleteKnowledgeKnowledgebaseLanguageDocument structure.
type DeleteKnowledgeKnowledgebaseLanguageDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentOK creates a DeleteKnowledgeKnowledgebaseLanguageDocumentOK with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentOK() *DeleteKnowledgeKnowledgebaseLanguageDocumentOK {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentOK{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentOK handles this case with default header values.

successful operation
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentOK struct {
	Payload *models.KnowledgeDocument
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentOK  %+v", 200, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentOK) GetPayload() *models.KnowledgeDocument {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentNoContent creates a DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentNoContent() *DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent handles this case with default header values.

Document deleted
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent struct {
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest creates a DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest() *DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized creates a DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized() *DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentForbidden creates a DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentForbidden() *DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentNotFound creates a DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentNotFound() *DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests creates a DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests() *DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError creates a DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError() *DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable creates a DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable() *DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout creates a DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout() *DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout{}
}

/*DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLanguageDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
