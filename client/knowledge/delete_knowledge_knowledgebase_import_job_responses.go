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

// DeleteKnowledgeKnowledgebaseImportJobReader is a Reader for the DeleteKnowledgeKnowledgebaseImportJob structure.
type DeleteKnowledgeKnowledgebaseImportJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseImportJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseImportJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseImportJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseImportJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseImportJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseImportJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteKnowledgeKnowledgebaseImportJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseImportJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseImportJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseImportJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseImportJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseImportJobNoContent creates a DeleteKnowledgeKnowledgebaseImportJobNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobNoContent() *DeleteKnowledgeKnowledgebaseImportJobNoContent {
	return &DeleteKnowledgeKnowledgebaseImportJobNoContent{}
}

/*DeleteKnowledgeKnowledgebaseImportJobNoContent handles this case with default header values.

Import job deleted
*/
type DeleteKnowledgeKnowledgebaseImportJobNoContent struct {
}

func (o *DeleteKnowledgeKnowledgebaseImportJobNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobBadRequest creates a DeleteKnowledgeKnowledgebaseImportJobBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobBadRequest() *DeleteKnowledgeKnowledgebaseImportJobBadRequest {
	return &DeleteKnowledgeKnowledgebaseImportJobBadRequest{}
}

/*DeleteKnowledgeKnowledgebaseImportJobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseImportJobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobUnauthorized creates a DeleteKnowledgeKnowledgebaseImportJobUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobUnauthorized() *DeleteKnowledgeKnowledgebaseImportJobUnauthorized {
	return &DeleteKnowledgeKnowledgebaseImportJobUnauthorized{}
}

/*DeleteKnowledgeKnowledgebaseImportJobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseImportJobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobForbidden creates a DeleteKnowledgeKnowledgebaseImportJobForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobForbidden() *DeleteKnowledgeKnowledgebaseImportJobForbidden {
	return &DeleteKnowledgeKnowledgebaseImportJobForbidden{}
}

/*DeleteKnowledgeKnowledgebaseImportJobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseImportJobForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobNotFound creates a DeleteKnowledgeKnowledgebaseImportJobNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobNotFound() *DeleteKnowledgeKnowledgebaseImportJobNotFound {
	return &DeleteKnowledgeKnowledgebaseImportJobNotFound{}
}

/*DeleteKnowledgeKnowledgebaseImportJobNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseImportJobNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobRequestTimeout creates a DeleteKnowledgeKnowledgebaseImportJobRequestTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobRequestTimeout() *DeleteKnowledgeKnowledgebaseImportJobRequestTimeout {
	return &DeleteKnowledgeKnowledgebaseImportJobRequestTimeout{}
}

/*DeleteKnowledgeKnowledgebaseImportJobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteKnowledgeKnowledgebaseImportJobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge{}
}

/*DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType{}
}

/*DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobTooManyRequests creates a DeleteKnowledgeKnowledgebaseImportJobTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobTooManyRequests() *DeleteKnowledgeKnowledgebaseImportJobTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseImportJobTooManyRequests{}
}

/*DeleteKnowledgeKnowledgebaseImportJobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseImportJobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobInternalServerError creates a DeleteKnowledgeKnowledgebaseImportJobInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobInternalServerError() *DeleteKnowledgeKnowledgebaseImportJobInternalServerError {
	return &DeleteKnowledgeKnowledgebaseImportJobInternalServerError{}
}

/*DeleteKnowledgeKnowledgebaseImportJobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseImportJobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobServiceUnavailable creates a DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobServiceUnavailable() *DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable{}
}

/*DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseImportJobGatewayTimeout creates a DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseImportJobGatewayTimeout() *DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout{}
}

/*DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/import/jobs/{importJobId}][%d] deleteKnowledgeKnowledgebaseImportJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseImportJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
