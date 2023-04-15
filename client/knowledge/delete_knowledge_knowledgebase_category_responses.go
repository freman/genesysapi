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

// DeleteKnowledgeKnowledgebaseCategoryReader is a Reader for the DeleteKnowledgeKnowledgebaseCategory structure.
type DeleteKnowledgeKnowledgebaseCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteKnowledgeKnowledgebaseCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteKnowledgeKnowledgebaseCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteKnowledgeKnowledgebaseCategoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteKnowledgeKnowledgebaseCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteKnowledgeKnowledgebaseCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteKnowledgeKnowledgebaseCategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteKnowledgeKnowledgebaseCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteKnowledgeKnowledgebaseCategoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteKnowledgeKnowledgebaseCategoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteKnowledgeKnowledgebaseCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteKnowledgeKnowledgebaseCategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteKnowledgeKnowledgebaseCategoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteKnowledgeKnowledgebaseCategoryOK creates a DeleteKnowledgeKnowledgebaseCategoryOK with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryOK() *DeleteKnowledgeKnowledgebaseCategoryOK {
	return &DeleteKnowledgeKnowledgebaseCategoryOK{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteKnowledgeKnowledgebaseCategoryOK struct {
	Payload *models.CategoryResponse
}

// IsSuccess returns true when this delete knowledge knowledgebase category o k response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete knowledge knowledgebase category o k response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category o k response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase category o k response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category o k response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteKnowledgeKnowledgebaseCategoryOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryOK  %+v", 200, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryOK  %+v", 200, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryOK) GetPayload() *models.CategoryResponse {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryNoContent creates a DeleteKnowledgeKnowledgebaseCategoryNoContent with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryNoContent() *DeleteKnowledgeKnowledgebaseCategoryNoContent {
	return &DeleteKnowledgeKnowledgebaseCategoryNoContent{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryNoContent describes a response with status code 204, with default header values.

Category deleted
*/
type DeleteKnowledgeKnowledgebaseCategoryNoContent struct {
}

// IsSuccess returns true when this delete knowledge knowledgebase category no content response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete knowledge knowledgebase category no content response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category no content response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase category no content response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category no content response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryNoContent ", 204)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryBadRequest creates a DeleteKnowledgeKnowledgebaseCategoryBadRequest with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryBadRequest() *DeleteKnowledgeKnowledgebaseCategoryBadRequest {
	return &DeleteKnowledgeKnowledgebaseCategoryBadRequest{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteKnowledgeKnowledgebaseCategoryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category bad request response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category bad request response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category bad request response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category bad request response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category bad request response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryUnauthorized creates a DeleteKnowledgeKnowledgebaseCategoryUnauthorized with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryUnauthorized() *DeleteKnowledgeKnowledgebaseCategoryUnauthorized {
	return &DeleteKnowledgeKnowledgebaseCategoryUnauthorized{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteKnowledgeKnowledgebaseCategoryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category unauthorized response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category unauthorized response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category unauthorized response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category unauthorized response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category unauthorized response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryForbidden creates a DeleteKnowledgeKnowledgebaseCategoryForbidden with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryForbidden() *DeleteKnowledgeKnowledgebaseCategoryForbidden {
	return &DeleteKnowledgeKnowledgebaseCategoryForbidden{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteKnowledgeKnowledgebaseCategoryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category forbidden response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category forbidden response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category forbidden response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category forbidden response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category forbidden response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryNotFound creates a DeleteKnowledgeKnowledgebaseCategoryNotFound with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryNotFound() *DeleteKnowledgeKnowledgebaseCategoryNotFound {
	return &DeleteKnowledgeKnowledgebaseCategoryNotFound{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteKnowledgeKnowledgebaseCategoryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category not found response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category not found response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category not found response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category not found response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category not found response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryRequestTimeout creates a DeleteKnowledgeKnowledgebaseCategoryRequestTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryRequestTimeout() *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout {
	return &DeleteKnowledgeKnowledgebaseCategoryRequestTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteKnowledgeKnowledgebaseCategoryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category request timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category request timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category request timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category request timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category request timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge creates a DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge() *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge {
	return &DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category request entity too large response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category request entity too large response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category request entity too large response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category request entity too large response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category request entity too large response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType creates a DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType() *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType {
	return &DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category unsupported media type response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category unsupported media type response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category unsupported media type response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category unsupported media type response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category unsupported media type response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryTooManyRequests creates a DeleteKnowledgeKnowledgebaseCategoryTooManyRequests with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryTooManyRequests() *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests {
	return &DeleteKnowledgeKnowledgebaseCategoryTooManyRequests{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteKnowledgeKnowledgebaseCategoryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category too many requests response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category too many requests response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category too many requests response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete knowledge knowledgebase category too many requests response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete knowledge knowledgebase category too many requests response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryInternalServerError creates a DeleteKnowledgeKnowledgebaseCategoryInternalServerError with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryInternalServerError() *DeleteKnowledgeKnowledgebaseCategoryInternalServerError {
	return &DeleteKnowledgeKnowledgebaseCategoryInternalServerError{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteKnowledgeKnowledgebaseCategoryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category internal server error response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category internal server error response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category internal server error response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase category internal server error response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase category internal server error response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryServiceUnavailable creates a DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryServiceUnavailable() *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable {
	return &DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category service unavailable response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category service unavailable response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category service unavailable response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase category service unavailable response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase category service unavailable response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteKnowledgeKnowledgebaseCategoryGatewayTimeout creates a DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout with default headers values
func NewDeleteKnowledgeKnowledgebaseCategoryGatewayTimeout() *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout {
	return &DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout{}
}

/*
DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete knowledge knowledgebase category gateway timeout response has a 2xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete knowledge knowledgebase category gateway timeout response has a 3xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete knowledge knowledgebase category gateway timeout response has a 4xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete knowledge knowledgebase category gateway timeout response has a 5xx status code
func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete knowledge knowledgebase category gateway timeout response a status code equal to that given
func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/categories/{categoryId}][%d] deleteKnowledgeKnowledgebaseCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteKnowledgeKnowledgebaseCategoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
