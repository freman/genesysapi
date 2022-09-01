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

// GetKnowledgeGuestSessionDocumentsReader is a Reader for the GetKnowledgeGuestSessionDocuments structure.
type GetKnowledgeGuestSessionDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeGuestSessionDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeGuestSessionDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeGuestSessionDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeGuestSessionDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeGuestSessionDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeGuestSessionDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeGuestSessionDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeGuestSessionDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeGuestSessionDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeGuestSessionDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeGuestSessionDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeGuestSessionDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeGuestSessionDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeGuestSessionDocumentsOK creates a GetKnowledgeGuestSessionDocumentsOK with default headers values
func NewGetKnowledgeGuestSessionDocumentsOK() *GetKnowledgeGuestSessionDocumentsOK {
	return &GetKnowledgeGuestSessionDocumentsOK{}
}

/*GetKnowledgeGuestSessionDocumentsOK handles this case with default header values.

successful operation
*/
type GetKnowledgeGuestSessionDocumentsOK struct {
	Payload *models.KnowledgeGuestDocumentResponseListing
}

func (o *GetKnowledgeGuestSessionDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsOK) GetPayload() *models.KnowledgeGuestDocumentResponseListing {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeGuestDocumentResponseListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsBadRequest creates a GetKnowledgeGuestSessionDocumentsBadRequest with default headers values
func NewGetKnowledgeGuestSessionDocumentsBadRequest() *GetKnowledgeGuestSessionDocumentsBadRequest {
	return &GetKnowledgeGuestSessionDocumentsBadRequest{}
}

/*GetKnowledgeGuestSessionDocumentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeGuestSessionDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsUnauthorized creates a GetKnowledgeGuestSessionDocumentsUnauthorized with default headers values
func NewGetKnowledgeGuestSessionDocumentsUnauthorized() *GetKnowledgeGuestSessionDocumentsUnauthorized {
	return &GetKnowledgeGuestSessionDocumentsUnauthorized{}
}

/*GetKnowledgeGuestSessionDocumentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeGuestSessionDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsForbidden creates a GetKnowledgeGuestSessionDocumentsForbidden with default headers values
func NewGetKnowledgeGuestSessionDocumentsForbidden() *GetKnowledgeGuestSessionDocumentsForbidden {
	return &GetKnowledgeGuestSessionDocumentsForbidden{}
}

/*GetKnowledgeGuestSessionDocumentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeGuestSessionDocumentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsNotFound creates a GetKnowledgeGuestSessionDocumentsNotFound with default headers values
func NewGetKnowledgeGuestSessionDocumentsNotFound() *GetKnowledgeGuestSessionDocumentsNotFound {
	return &GetKnowledgeGuestSessionDocumentsNotFound{}
}

/*GetKnowledgeGuestSessionDocumentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetKnowledgeGuestSessionDocumentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsRequestTimeout creates a GetKnowledgeGuestSessionDocumentsRequestTimeout with default headers values
func NewGetKnowledgeGuestSessionDocumentsRequestTimeout() *GetKnowledgeGuestSessionDocumentsRequestTimeout {
	return &GetKnowledgeGuestSessionDocumentsRequestTimeout{}
}

/*GetKnowledgeGuestSessionDocumentsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeGuestSessionDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsRequestEntityTooLarge creates a GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge with default headers values
func NewGetKnowledgeGuestSessionDocumentsRequestEntityTooLarge() *GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge {
	return &GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge{}
}

/*GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsUnsupportedMediaType creates a GetKnowledgeGuestSessionDocumentsUnsupportedMediaType with default headers values
func NewGetKnowledgeGuestSessionDocumentsUnsupportedMediaType() *GetKnowledgeGuestSessionDocumentsUnsupportedMediaType {
	return &GetKnowledgeGuestSessionDocumentsUnsupportedMediaType{}
}

/*GetKnowledgeGuestSessionDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeGuestSessionDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsTooManyRequests creates a GetKnowledgeGuestSessionDocumentsTooManyRequests with default headers values
func NewGetKnowledgeGuestSessionDocumentsTooManyRequests() *GetKnowledgeGuestSessionDocumentsTooManyRequests {
	return &GetKnowledgeGuestSessionDocumentsTooManyRequests{}
}

/*GetKnowledgeGuestSessionDocumentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeGuestSessionDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsInternalServerError creates a GetKnowledgeGuestSessionDocumentsInternalServerError with default headers values
func NewGetKnowledgeGuestSessionDocumentsInternalServerError() *GetKnowledgeGuestSessionDocumentsInternalServerError {
	return &GetKnowledgeGuestSessionDocumentsInternalServerError{}
}

/*GetKnowledgeGuestSessionDocumentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeGuestSessionDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsServiceUnavailable creates a GetKnowledgeGuestSessionDocumentsServiceUnavailable with default headers values
func NewGetKnowledgeGuestSessionDocumentsServiceUnavailable() *GetKnowledgeGuestSessionDocumentsServiceUnavailable {
	return &GetKnowledgeGuestSessionDocumentsServiceUnavailable{}
}

/*GetKnowledgeGuestSessionDocumentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeGuestSessionDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeGuestSessionDocumentsGatewayTimeout creates a GetKnowledgeGuestSessionDocumentsGatewayTimeout with default headers values
func NewGetKnowledgeGuestSessionDocumentsGatewayTimeout() *GetKnowledgeGuestSessionDocumentsGatewayTimeout {
	return &GetKnowledgeGuestSessionDocumentsGatewayTimeout{}
}

/*GetKnowledgeGuestSessionDocumentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetKnowledgeGuestSessionDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeGuestSessionDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/guest/sessions/{sessionId}/documents][%d] getKnowledgeGuestSessionDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeGuestSessionDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeGuestSessionDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
