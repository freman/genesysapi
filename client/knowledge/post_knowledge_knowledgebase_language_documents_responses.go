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

// PostKnowledgeKnowledgebaseLanguageDocumentsReader is a Reader for the PostKnowledgeKnowledgebaseLanguageDocuments structure.
type PostKnowledgeKnowledgebaseLanguageDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsOK creates a PostKnowledgeKnowledgebaseLanguageDocumentsOK with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsOK() *PostKnowledgeKnowledgebaseLanguageDocumentsOK {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsOK{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsOK handles this case with default header values.

successful operation
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsOK struct {
	Payload *models.KnowledgeDocument
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsOK) GetPayload() *models.KnowledgeDocument {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsBadRequest creates a PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsBadRequest() *PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized creates a PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized() *PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsForbidden creates a PostKnowledgeKnowledgebaseLanguageDocumentsForbidden with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsForbidden() *PostKnowledgeKnowledgebaseLanguageDocumentsForbidden {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsForbidden{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsNotFound creates a PostKnowledgeKnowledgebaseLanguageDocumentsNotFound with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsNotFound() *PostKnowledgeKnowledgebaseLanguageDocumentsNotFound {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsNotFound{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout creates a PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout() *PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge creates a PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge() *PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType creates a PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType() *PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests creates a PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests() *PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError creates a PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError() *PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable creates a PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable() *PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout creates a PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout() *PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout {
	return &PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout{}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents][%d] postKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLanguageDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
