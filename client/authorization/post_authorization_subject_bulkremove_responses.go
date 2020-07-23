// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuthorizationSubjectBulkremoveReader is a Reader for the PostAuthorizationSubjectBulkremove structure.
type PostAuthorizationSubjectBulkremoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthorizationSubjectBulkremoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAuthorizationSubjectBulkremoveNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthorizationSubjectBulkremoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuthorizationSubjectBulkremoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuthorizationSubjectBulkremoveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuthorizationSubjectBulkremoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuthorizationSubjectBulkremoveRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuthorizationSubjectBulkremoveUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuthorizationSubjectBulkremoveTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthorizationSubjectBulkremoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuthorizationSubjectBulkremoveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuthorizationSubjectBulkremoveGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthorizationSubjectBulkremoveNoContent creates a PostAuthorizationSubjectBulkremoveNoContent with default headers values
func NewPostAuthorizationSubjectBulkremoveNoContent() *PostAuthorizationSubjectBulkremoveNoContent {
	return &PostAuthorizationSubjectBulkremoveNoContent{}
}

/*PostAuthorizationSubjectBulkremoveNoContent handles this case with default header values.

Bulk Grants Deleted
*/
type PostAuthorizationSubjectBulkremoveNoContent struct {
}

func (o *PostAuthorizationSubjectBulkremoveNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveNoContent ", 204)
}

func (o *PostAuthorizationSubjectBulkremoveNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAuthorizationSubjectBulkremoveBadRequest creates a PostAuthorizationSubjectBulkremoveBadRequest with default headers values
func NewPostAuthorizationSubjectBulkremoveBadRequest() *PostAuthorizationSubjectBulkremoveBadRequest {
	return &PostAuthorizationSubjectBulkremoveBadRequest{}
}

/*PostAuthorizationSubjectBulkremoveBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuthorizationSubjectBulkremoveBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveUnauthorized creates a PostAuthorizationSubjectBulkremoveUnauthorized with default headers values
func NewPostAuthorizationSubjectBulkremoveUnauthorized() *PostAuthorizationSubjectBulkremoveUnauthorized {
	return &PostAuthorizationSubjectBulkremoveUnauthorized{}
}

/*PostAuthorizationSubjectBulkremoveUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuthorizationSubjectBulkremoveUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveForbidden creates a PostAuthorizationSubjectBulkremoveForbidden with default headers values
func NewPostAuthorizationSubjectBulkremoveForbidden() *PostAuthorizationSubjectBulkremoveForbidden {
	return &PostAuthorizationSubjectBulkremoveForbidden{}
}

/*PostAuthorizationSubjectBulkremoveForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAuthorizationSubjectBulkremoveForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveNotFound creates a PostAuthorizationSubjectBulkremoveNotFound with default headers values
func NewPostAuthorizationSubjectBulkremoveNotFound() *PostAuthorizationSubjectBulkremoveNotFound {
	return &PostAuthorizationSubjectBulkremoveNotFound{}
}

/*PostAuthorizationSubjectBulkremoveNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAuthorizationSubjectBulkremoveNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveRequestEntityTooLarge creates a PostAuthorizationSubjectBulkremoveRequestEntityTooLarge with default headers values
func NewPostAuthorizationSubjectBulkremoveRequestEntityTooLarge() *PostAuthorizationSubjectBulkremoveRequestEntityTooLarge {
	return &PostAuthorizationSubjectBulkremoveRequestEntityTooLarge{}
}

/*PostAuthorizationSubjectBulkremoveRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAuthorizationSubjectBulkremoveRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveUnsupportedMediaType creates a PostAuthorizationSubjectBulkremoveUnsupportedMediaType with default headers values
func NewPostAuthorizationSubjectBulkremoveUnsupportedMediaType() *PostAuthorizationSubjectBulkremoveUnsupportedMediaType {
	return &PostAuthorizationSubjectBulkremoveUnsupportedMediaType{}
}

/*PostAuthorizationSubjectBulkremoveUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuthorizationSubjectBulkremoveUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveTooManyRequests creates a PostAuthorizationSubjectBulkremoveTooManyRequests with default headers values
func NewPostAuthorizationSubjectBulkremoveTooManyRequests() *PostAuthorizationSubjectBulkremoveTooManyRequests {
	return &PostAuthorizationSubjectBulkremoveTooManyRequests{}
}

/*PostAuthorizationSubjectBulkremoveTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAuthorizationSubjectBulkremoveTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveInternalServerError creates a PostAuthorizationSubjectBulkremoveInternalServerError with default headers values
func NewPostAuthorizationSubjectBulkremoveInternalServerError() *PostAuthorizationSubjectBulkremoveInternalServerError {
	return &PostAuthorizationSubjectBulkremoveInternalServerError{}
}

/*PostAuthorizationSubjectBulkremoveInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuthorizationSubjectBulkremoveInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveServiceUnavailable creates a PostAuthorizationSubjectBulkremoveServiceUnavailable with default headers values
func NewPostAuthorizationSubjectBulkremoveServiceUnavailable() *PostAuthorizationSubjectBulkremoveServiceUnavailable {
	return &PostAuthorizationSubjectBulkremoveServiceUnavailable{}
}

/*PostAuthorizationSubjectBulkremoveServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuthorizationSubjectBulkremoveServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkremoveGatewayTimeout creates a PostAuthorizationSubjectBulkremoveGatewayTimeout with default headers values
func NewPostAuthorizationSubjectBulkremoveGatewayTimeout() *PostAuthorizationSubjectBulkremoveGatewayTimeout {
	return &PostAuthorizationSubjectBulkremoveGatewayTimeout{}
}

/*PostAuthorizationSubjectBulkremoveGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAuthorizationSubjectBulkremoveGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuthorizationSubjectBulkremoveGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkremove][%d] postAuthorizationSubjectBulkremoveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationSubjectBulkremoveGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkremoveGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
