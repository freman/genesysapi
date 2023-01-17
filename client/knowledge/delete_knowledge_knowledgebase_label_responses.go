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

// DeleteKnowledgeKnowledgebaseLabelReader is a Reader for the DeleteKnowledgeKnowledgebaseLabel structure.
type DeleteKnowledgeKnowledgebaseLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteKnowledgeKnowledgebaseLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseLabelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteKnowledgeKnowledgebaseLabelRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseLabelTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseLabelServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseLabelGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseLabelOK creates a DeleteKnowledgeKnowledgebaseLabelOK with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelOK() *DeleteKnowledgeKnowledgebaseLabelOK {
	return &DeleteKnowledgeKnowledgebaseLabelOK{}
}

/*
DeleteKnowledgeKnowledgebaseLabelOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteKnowledgeKnowledgebaseLabelOK struct {
	Payload *models.LabelResponse
}

// IsSuccess returns true when this delete knowledge knowledgebase label o k response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete knowledge knowledgebase label o k response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label o k response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase label o k response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label o k response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteKnowledgeKnowledgebaseLabelOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelOK  %+v", 200, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelOK  %+v", 200, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelOK) GetPayload() *models.LabelResponse {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelNoContent creates a DeleteKnowledgeKnowledgebaseLabelNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelNoContent() *DeleteKnowledgeKnowledgebaseLabelNoContent {
	return &DeleteKnowledgeKnowledgebaseLabelNoContent{}
}

/*
DeleteKnowledgeKnowledgebaseLabelNoContent describes a response with status code 204, with default header values.

Label deleted
*/
type DeleteKnowledgeKnowledgebaseLabelNoContent struct {
}

// IsSuccess returns true when this delete knowledge knowledgebase label no content response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete knowledge knowledgebase label no content response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label no content response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase label no content response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label no content response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseLabelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelBadRequest creates a DeleteKnowledgeKnowledgebaseLabelBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelBadRequest() *DeleteKnowledgeKnowledgebaseLabelBadRequest {
	return &DeleteKnowledgeKnowledgebaseLabelBadRequest{}
}

/*
DeleteKnowledgeKnowledgebaseLabelBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseLabelBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label bad request response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label bad request response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label bad request response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label bad request response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label bad request response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelUnauthorized creates a DeleteKnowledgeKnowledgebaseLabelUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelUnauthorized() *DeleteKnowledgeKnowledgebaseLabelUnauthorized {
	return &DeleteKnowledgeKnowledgebaseLabelUnauthorized{}
}

/*
DeleteKnowledgeKnowledgebaseLabelUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseLabelUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label unauthorized response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label unauthorized response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label unauthorized response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label unauthorized response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label unauthorized response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelForbidden creates a DeleteKnowledgeKnowledgebaseLabelForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelForbidden() *DeleteKnowledgeKnowledgebaseLabelForbidden {
	return &DeleteKnowledgeKnowledgebaseLabelForbidden{}
}

/*
DeleteKnowledgeKnowledgebaseLabelForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseLabelForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label forbidden response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label forbidden response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label forbidden response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label forbidden response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label forbidden response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelNotFound creates a DeleteKnowledgeKnowledgebaseLabelNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelNotFound() *DeleteKnowledgeKnowledgebaseLabelNotFound {
	return &DeleteKnowledgeKnowledgebaseLabelNotFound{}
}

/*
DeleteKnowledgeKnowledgebaseLabelNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseLabelNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label not found response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label not found response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label not found response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label not found response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label not found response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelRequestTimeout creates a DeleteKnowledgeKnowledgebaseLabelRequestTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelRequestTimeout() *DeleteKnowledgeKnowledgebaseLabelRequestTimeout {
	return &DeleteKnowledgeKnowledgebaseLabelRequestTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseLabelRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteKnowledgeKnowledgebaseLabelRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label request timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label request timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label request timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label request timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label request timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge{}
}

/*
DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label request entity too large response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label request entity too large response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label request entity too large response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label request entity too large response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label request entity too large response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType{}
}

/*
DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label unsupported media type response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label unsupported media type response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label unsupported media type response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label unsupported media type response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label unsupported media type response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelTooManyRequests creates a DeleteKnowledgeKnowledgebaseLabelTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelTooManyRequests() *DeleteKnowledgeKnowledgebaseLabelTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseLabelTooManyRequests{}
}

/*
DeleteKnowledgeKnowledgebaseLabelTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseLabelTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label too many requests response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label too many requests response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label too many requests response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase label too many requests response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase label too many requests response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelInternalServerError creates a DeleteKnowledgeKnowledgebaseLabelInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelInternalServerError() *DeleteKnowledgeKnowledgebaseLabelInternalServerError {
	return &DeleteKnowledgeKnowledgebaseLabelInternalServerError{}
}

/*
DeleteKnowledgeKnowledgebaseLabelInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseLabelInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label internal server error response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label internal server error response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label internal server error response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase label internal server error response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase label internal server error response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelServiceUnavailable creates a DeleteKnowledgeKnowledgebaseLabelServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelServiceUnavailable() *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseLabelServiceUnavailable{}
}

/*
DeleteKnowledgeKnowledgebaseLabelServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseLabelServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label service unavailable response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label service unavailable response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label service unavailable response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase label service unavailable response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase label service unavailable response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseLabelGatewayTimeout creates a DeleteKnowledgeKnowledgebaseLabelGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseLabelGatewayTimeout() *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseLabelGatewayTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseLabelGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseLabelGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase label gateway timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase label gateway timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase label gateway timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase label gateway timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase label gateway timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels/{labelId}][%d] deleteKnowledgeKnowledgebaseLabelGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseLabelGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
