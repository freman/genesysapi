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

// DeleteKnowledgeKnowledgebaseDocumentReader is a Reader for the DeleteKnowledgeKnowledgebaseDocument structure.
type DeleteKnowledgeKnowledgebaseDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseDocumentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteKnowledgeKnowledgebaseDocumentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseDocumentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseDocumentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseDocumentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseDocumentNoContent creates a DeleteKnowledgeKnowledgebaseDocumentNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentNoContent() *DeleteKnowledgeKnowledgebaseDocumentNoContent {
	return &DeleteKnowledgeKnowledgebaseDocumentNoContent{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentNoContent describes a response with status code 204, with default header values.

Document deleted.
*/
type DeleteKnowledgeKnowledgebaseDocumentNoContent struct {
}

// IsSuccess returns true when this delete knowledge knowledgebase document no content response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete knowledge knowledgebase document no content response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document no content response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase document no content response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document no content response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentBadRequest creates a DeleteKnowledgeKnowledgebaseDocumentBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentBadRequest() *DeleteKnowledgeKnowledgebaseDocumentBadRequest {
	return &DeleteKnowledgeKnowledgebaseDocumentBadRequest{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseDocumentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document bad request response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document bad request response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document bad request response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document bad request response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document bad request response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentUnauthorized creates a DeleteKnowledgeKnowledgebaseDocumentUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentUnauthorized() *DeleteKnowledgeKnowledgebaseDocumentUnauthorized {
	return &DeleteKnowledgeKnowledgebaseDocumentUnauthorized{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseDocumentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document unauthorized response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document unauthorized response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document unauthorized response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document unauthorized response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document unauthorized response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentForbidden creates a DeleteKnowledgeKnowledgebaseDocumentForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentForbidden() *DeleteKnowledgeKnowledgebaseDocumentForbidden {
	return &DeleteKnowledgeKnowledgebaseDocumentForbidden{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseDocumentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document forbidden response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document forbidden response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document forbidden response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document forbidden response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document forbidden response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentNotFound creates a DeleteKnowledgeKnowledgebaseDocumentNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentNotFound() *DeleteKnowledgeKnowledgebaseDocumentNotFound {
	return &DeleteKnowledgeKnowledgebaseDocumentNotFound{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseDocumentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document not found response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document not found response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document not found response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document not found response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document not found response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentRequestTimeout creates a DeleteKnowledgeKnowledgebaseDocumentRequestTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentRequestTimeout() *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout {
	return &DeleteKnowledgeKnowledgebaseDocumentRequestTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteKnowledgeKnowledgebaseDocumentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document request timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document request timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document request timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document request timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document request timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document request entity too large response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document request entity too large response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document request entity too large response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document request entity too large response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document request entity too large response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document unsupported media type response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document unsupported media type response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document unsupported media type response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document unsupported media type response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document unsupported media type response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentTooManyRequests creates a DeleteKnowledgeKnowledgebaseDocumentTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentTooManyRequests() *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseDocumentTooManyRequests{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseDocumentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document too many requests response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document too many requests response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document too many requests response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase document too many requests response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase document too many requests response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentInternalServerError creates a DeleteKnowledgeKnowledgebaseDocumentInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentInternalServerError() *DeleteKnowledgeKnowledgebaseDocumentInternalServerError {
	return &DeleteKnowledgeKnowledgebaseDocumentInternalServerError{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseDocumentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document internal server error response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document internal server error response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document internal server error response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase document internal server error response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase document internal server error response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentServiceUnavailable creates a DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentServiceUnavailable() *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document service unavailable response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document service unavailable response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document service unavailable response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase document service unavailable response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase document service unavailable response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseDocumentGatewayTimeout creates a DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseDocumentGatewayTimeout() *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase document gateway timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase document gateway timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase document gateway timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase document gateway timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase document gateway timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/documents/{documentId}][%d] deleteKnowledgeKnowledgebaseDocumentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseDocumentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
