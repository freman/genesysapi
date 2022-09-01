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

// GetKnowledgeKnowledgebaseDocumentsReader is a Reader for the GetKnowledgeKnowledgebaseDocuments structure.
type GetKnowledgeKnowledgebaseDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseDocumentsOK creates a GetKnowledgeKnowledgebaseDocumentsOK with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsOK() *GetKnowledgeKnowledgebaseDocumentsOK {
	return &GetKnowledgeKnowledgebaseDocumentsOK{}
}

/*GetKnowledgeKnowledgebaseDocumentsOK handles this case with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseDocumentsOK struct {
	Payload *models.KnowledgeDocumentResponseListing
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) GetPayload() *models.KnowledgeDocumentResponseListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeDocumentResponseListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsBadRequest creates a GetKnowledgeKnowledgebaseDocumentsBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsBadRequest() *GetKnowledgeKnowledgebaseDocumentsBadRequest {
	return &GetKnowledgeKnowledgebaseDocumentsBadRequest{}
}

/*GetKnowledgeKnowledgebaseDocumentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsUnauthorized creates a GetKnowledgeKnowledgebaseDocumentsUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsUnauthorized() *GetKnowledgeKnowledgebaseDocumentsUnauthorized {
	return &GetKnowledgeKnowledgebaseDocumentsUnauthorized{}
}

/*GetKnowledgeKnowledgebaseDocumentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsForbidden creates a GetKnowledgeKnowledgebaseDocumentsForbidden with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsForbidden() *GetKnowledgeKnowledgebaseDocumentsForbidden {
	return &GetKnowledgeKnowledgebaseDocumentsForbidden{}
}

/*GetKnowledgeKnowledgebaseDocumentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseDocumentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsNotFound creates a GetKnowledgeKnowledgebaseDocumentsNotFound with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsNotFound() *GetKnowledgeKnowledgebaseDocumentsNotFound {
	return &GetKnowledgeKnowledgebaseDocumentsNotFound{}
}

/*GetKnowledgeKnowledgebaseDocumentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseDocumentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsRequestTimeout creates a GetKnowledgeKnowledgebaseDocumentsRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsRequestTimeout() *GetKnowledgeKnowledgebaseDocumentsRequestTimeout {
	return &GetKnowledgeKnowledgebaseDocumentsRequestTimeout{}
}

/*GetKnowledgeKnowledgebaseDocumentsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge() *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge{}
}

/*GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType creates a GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType() *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType{}
}

/*GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsTooManyRequests creates a GetKnowledgeKnowledgebaseDocumentsTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsTooManyRequests() *GetKnowledgeKnowledgebaseDocumentsTooManyRequests {
	return &GetKnowledgeKnowledgebaseDocumentsTooManyRequests{}
}

/*GetKnowledgeKnowledgebaseDocumentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsInternalServerError creates a GetKnowledgeKnowledgebaseDocumentsInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsInternalServerError() *GetKnowledgeKnowledgebaseDocumentsInternalServerError {
	return &GetKnowledgeKnowledgebaseDocumentsInternalServerError{}
}

/*GetKnowledgeKnowledgebaseDocumentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsServiceUnavailable creates a GetKnowledgeKnowledgebaseDocumentsServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsServiceUnavailable() *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable {
	return &GetKnowledgeKnowledgebaseDocumentsServiceUnavailable{}
}

/*GetKnowledgeKnowledgebaseDocumentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseDocumentsGatewayTimeout creates a GetKnowledgeKnowledgebaseDocumentsGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseDocumentsGatewayTimeout() *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout {
	return &GetKnowledgeKnowledgebaseDocumentsGatewayTimeout{}
}

/*GetKnowledgeKnowledgebaseDocumentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents][%d] getKnowledgeKnowledgebaseDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
