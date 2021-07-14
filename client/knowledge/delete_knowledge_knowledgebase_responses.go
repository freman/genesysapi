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

// DeleteKnowledgeKnowledgebaseReader is a Reader for the DeleteKnowledgeKnowledgebase structure.
type DeleteKnowledgeKnowledgebaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteKnowledgeKnowledgebaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteKnowledgeKnowledgebaseRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseOK creates a DeleteKnowledgeKnowledgebaseOK with default headers values
func NewDeleteKnowledgeKnowledgebaseOK() *DeleteKnowledgeKnowledgebaseOK {
	return &DeleteKnowledgeKnowledgebaseOK{}
}

/*DeleteKnowledgeKnowledgebaseOK handles this case with default header values.

successful operation
*/
type DeleteKnowledgeKnowledgebaseOK struct {
	Payload *models.KnowledgeBase
}

func (o *DeleteKnowledgeKnowledgebaseOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseOK  %+v", 200, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseOK) GetPayload() *models.KnowledgeBase {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseNoContent creates a DeleteKnowledgeKnowledgebaseNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseNoContent() *DeleteKnowledgeKnowledgebaseNoContent {
	return &DeleteKnowledgeKnowledgebaseNoContent{}
}

/*DeleteKnowledgeKnowledgebaseNoContent handles this case with default header values.

Knowledge base deleted
*/
type DeleteKnowledgeKnowledgebaseNoContent struct {
}

func (o *DeleteKnowledgeKnowledgebaseNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseBadRequest creates a DeleteKnowledgeKnowledgebaseBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseBadRequest() *DeleteKnowledgeKnowledgebaseBadRequest {
	return &DeleteKnowledgeKnowledgebaseBadRequest{}
}

/*DeleteKnowledgeKnowledgebaseBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseUnauthorized creates a DeleteKnowledgeKnowledgebaseUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseUnauthorized() *DeleteKnowledgeKnowledgebaseUnauthorized {
	return &DeleteKnowledgeKnowledgebaseUnauthorized{}
}

/*DeleteKnowledgeKnowledgebaseUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseForbidden creates a DeleteKnowledgeKnowledgebaseForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseForbidden() *DeleteKnowledgeKnowledgebaseForbidden {
	return &DeleteKnowledgeKnowledgebaseForbidden{}
}

/*DeleteKnowledgeKnowledgebaseForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseNotFound creates a DeleteKnowledgeKnowledgebaseNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseNotFound() *DeleteKnowledgeKnowledgebaseNotFound {
	return &DeleteKnowledgeKnowledgebaseNotFound{}
}

/*DeleteKnowledgeKnowledgebaseNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseRequestTimeout creates a DeleteKnowledgeKnowledgebaseRequestTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseRequestTimeout() *DeleteKnowledgeKnowledgebaseRequestTimeout {
	return &DeleteKnowledgeKnowledgebaseRequestTimeout{}
}

/*DeleteKnowledgeKnowledgebaseRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteKnowledgeKnowledgebaseRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseRequestEntityTooLarge{}
}

/*DeleteKnowledgeKnowledgebaseRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteKnowledgeKnowledgebaseRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseUnsupportedMediaType{}
}

/*DeleteKnowledgeKnowledgebaseUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseTooManyRequests creates a DeleteKnowledgeKnowledgebaseTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseTooManyRequests() *DeleteKnowledgeKnowledgebaseTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseTooManyRequests{}
}

/*DeleteKnowledgeKnowledgebaseTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseInternalServerError creates a DeleteKnowledgeKnowledgebaseInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseInternalServerError() *DeleteKnowledgeKnowledgebaseInternalServerError {
	return &DeleteKnowledgeKnowledgebaseInternalServerError{}
}

/*DeleteKnowledgeKnowledgebaseInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseServiceUnavailable creates a DeleteKnowledgeKnowledgebaseServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseServiceUnavailable() *DeleteKnowledgeKnowledgebaseServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseServiceUnavailable{}
}

/*DeleteKnowledgeKnowledgebaseServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseGatewayTimeout creates a DeleteKnowledgeKnowledgebaseGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseGatewayTimeout() *DeleteKnowledgeKnowledgebaseGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseGatewayTimeout{}
}

/*DeleteKnowledgeKnowledgebaseGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}][%d] deleteKnowledgeKnowledgebaseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
