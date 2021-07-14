// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteLearningModuleReader is a Reader for the DeleteLearningModule structure.
type DeleteLearningModuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLearningModuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLearningModuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLearningModuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLearningModuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLearningModuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLearningModuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteLearningModuleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteLearningModuleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteLearningModuleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteLearningModuleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLearningModuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteLearningModuleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteLearningModuleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLearningModuleNoContent creates a DeleteLearningModuleNoContent with default headers values
func NewDeleteLearningModuleNoContent() *DeleteLearningModuleNoContent {
	return &DeleteLearningModuleNoContent{}
}

/*DeleteLearningModuleNoContent handles this case with default header values.

The learning module was deleted successfully
*/
type DeleteLearningModuleNoContent struct {
}

func (o *DeleteLearningModuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleNoContent ", 204)
}

func (o *DeleteLearningModuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLearningModuleBadRequest creates a DeleteLearningModuleBadRequest with default headers values
func NewDeleteLearningModuleBadRequest() *DeleteLearningModuleBadRequest {
	return &DeleteLearningModuleBadRequest{}
}

/*DeleteLearningModuleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteLearningModuleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLearningModuleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleUnauthorized creates a DeleteLearningModuleUnauthorized with default headers values
func NewDeleteLearningModuleUnauthorized() *DeleteLearningModuleUnauthorized {
	return &DeleteLearningModuleUnauthorized{}
}

/*DeleteLearningModuleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteLearningModuleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLearningModuleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleForbidden creates a DeleteLearningModuleForbidden with default headers values
func NewDeleteLearningModuleForbidden() *DeleteLearningModuleForbidden {
	return &DeleteLearningModuleForbidden{}
}

/*DeleteLearningModuleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteLearningModuleForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLearningModuleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleNotFound creates a DeleteLearningModuleNotFound with default headers values
func NewDeleteLearningModuleNotFound() *DeleteLearningModuleNotFound {
	return &DeleteLearningModuleNotFound{}
}

/*DeleteLearningModuleNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteLearningModuleNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLearningModuleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleRequestTimeout creates a DeleteLearningModuleRequestTimeout with default headers values
func NewDeleteLearningModuleRequestTimeout() *DeleteLearningModuleRequestTimeout {
	return &DeleteLearningModuleRequestTimeout{}
}

/*DeleteLearningModuleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteLearningModuleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteLearningModuleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleRequestEntityTooLarge creates a DeleteLearningModuleRequestEntityTooLarge with default headers values
func NewDeleteLearningModuleRequestEntityTooLarge() *DeleteLearningModuleRequestEntityTooLarge {
	return &DeleteLearningModuleRequestEntityTooLarge{}
}

/*DeleteLearningModuleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteLearningModuleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteLearningModuleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleUnsupportedMediaType creates a DeleteLearningModuleUnsupportedMediaType with default headers values
func NewDeleteLearningModuleUnsupportedMediaType() *DeleteLearningModuleUnsupportedMediaType {
	return &DeleteLearningModuleUnsupportedMediaType{}
}

/*DeleteLearningModuleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteLearningModuleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteLearningModuleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleTooManyRequests creates a DeleteLearningModuleTooManyRequests with default headers values
func NewDeleteLearningModuleTooManyRequests() *DeleteLearningModuleTooManyRequests {
	return &DeleteLearningModuleTooManyRequests{}
}

/*DeleteLearningModuleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteLearningModuleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteLearningModuleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleInternalServerError creates a DeleteLearningModuleInternalServerError with default headers values
func NewDeleteLearningModuleInternalServerError() *DeleteLearningModuleInternalServerError {
	return &DeleteLearningModuleInternalServerError{}
}

/*DeleteLearningModuleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteLearningModuleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLearningModuleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleServiceUnavailable creates a DeleteLearningModuleServiceUnavailable with default headers values
func NewDeleteLearningModuleServiceUnavailable() *DeleteLearningModuleServiceUnavailable {
	return &DeleteLearningModuleServiceUnavailable{}
}

/*DeleteLearningModuleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteLearningModuleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteLearningModuleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLearningModuleGatewayTimeout creates a DeleteLearningModuleGatewayTimeout with default headers values
func NewDeleteLearningModuleGatewayTimeout() *DeleteLearningModuleGatewayTimeout {
	return &DeleteLearningModuleGatewayTimeout{}
}

/*DeleteLearningModuleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteLearningModuleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteLearningModuleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/learning/modules/{moduleId}][%d] deleteLearningModuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteLearningModuleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteLearningModuleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}