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

// PostLearningAssignmentsBulkaddReader is a Reader for the PostLearningAssignmentsBulkadd structure.
type PostLearningAssignmentsBulkaddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearningAssignmentsBulkaddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLearningAssignmentsBulkaddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearningAssignmentsBulkaddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLearningAssignmentsBulkaddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearningAssignmentsBulkaddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLearningAssignmentsBulkaddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLearningAssignmentsBulkaddRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLearningAssignmentsBulkaddRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLearningAssignmentsBulkaddUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLearningAssignmentsBulkaddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLearningAssignmentsBulkaddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLearningAssignmentsBulkaddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLearningAssignmentsBulkaddGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearningAssignmentsBulkaddOK creates a PostLearningAssignmentsBulkaddOK with default headers values
func NewPostLearningAssignmentsBulkaddOK() *PostLearningAssignmentsBulkaddOK {
	return &PostLearningAssignmentsBulkaddOK{}
}

/*PostLearningAssignmentsBulkaddOK handles this case with default header values.

successful operation
*/
type PostLearningAssignmentsBulkaddOK struct {
	Payload *models.LearningAssignmentBulkAddResponse
}

func (o *PostLearningAssignmentsBulkaddOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddOK) GetPayload() *models.LearningAssignmentBulkAddResponse {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningAssignmentBulkAddResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddBadRequest creates a PostLearningAssignmentsBulkaddBadRequest with default headers values
func NewPostLearningAssignmentsBulkaddBadRequest() *PostLearningAssignmentsBulkaddBadRequest {
	return &PostLearningAssignmentsBulkaddBadRequest{}
}

/*PostLearningAssignmentsBulkaddBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLearningAssignmentsBulkaddBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddUnauthorized creates a PostLearningAssignmentsBulkaddUnauthorized with default headers values
func NewPostLearningAssignmentsBulkaddUnauthorized() *PostLearningAssignmentsBulkaddUnauthorized {
	return &PostLearningAssignmentsBulkaddUnauthorized{}
}

/*PostLearningAssignmentsBulkaddUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLearningAssignmentsBulkaddUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddForbidden creates a PostLearningAssignmentsBulkaddForbidden with default headers values
func NewPostLearningAssignmentsBulkaddForbidden() *PostLearningAssignmentsBulkaddForbidden {
	return &PostLearningAssignmentsBulkaddForbidden{}
}

/*PostLearningAssignmentsBulkaddForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLearningAssignmentsBulkaddForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddNotFound creates a PostLearningAssignmentsBulkaddNotFound with default headers values
func NewPostLearningAssignmentsBulkaddNotFound() *PostLearningAssignmentsBulkaddNotFound {
	return &PostLearningAssignmentsBulkaddNotFound{}
}

/*PostLearningAssignmentsBulkaddNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLearningAssignmentsBulkaddNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddRequestTimeout creates a PostLearningAssignmentsBulkaddRequestTimeout with default headers values
func NewPostLearningAssignmentsBulkaddRequestTimeout() *PostLearningAssignmentsBulkaddRequestTimeout {
	return &PostLearningAssignmentsBulkaddRequestTimeout{}
}

/*PostLearningAssignmentsBulkaddRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLearningAssignmentsBulkaddRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddRequestEntityTooLarge creates a PostLearningAssignmentsBulkaddRequestEntityTooLarge with default headers values
func NewPostLearningAssignmentsBulkaddRequestEntityTooLarge() *PostLearningAssignmentsBulkaddRequestEntityTooLarge {
	return &PostLearningAssignmentsBulkaddRequestEntityTooLarge{}
}

/*PostLearningAssignmentsBulkaddRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLearningAssignmentsBulkaddRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddUnsupportedMediaType creates a PostLearningAssignmentsBulkaddUnsupportedMediaType with default headers values
func NewPostLearningAssignmentsBulkaddUnsupportedMediaType() *PostLearningAssignmentsBulkaddUnsupportedMediaType {
	return &PostLearningAssignmentsBulkaddUnsupportedMediaType{}
}

/*PostLearningAssignmentsBulkaddUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLearningAssignmentsBulkaddUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddTooManyRequests creates a PostLearningAssignmentsBulkaddTooManyRequests with default headers values
func NewPostLearningAssignmentsBulkaddTooManyRequests() *PostLearningAssignmentsBulkaddTooManyRequests {
	return &PostLearningAssignmentsBulkaddTooManyRequests{}
}

/*PostLearningAssignmentsBulkaddTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLearningAssignmentsBulkaddTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddInternalServerError creates a PostLearningAssignmentsBulkaddInternalServerError with default headers values
func NewPostLearningAssignmentsBulkaddInternalServerError() *PostLearningAssignmentsBulkaddInternalServerError {
	return &PostLearningAssignmentsBulkaddInternalServerError{}
}

/*PostLearningAssignmentsBulkaddInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLearningAssignmentsBulkaddInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddServiceUnavailable creates a PostLearningAssignmentsBulkaddServiceUnavailable with default headers values
func NewPostLearningAssignmentsBulkaddServiceUnavailable() *PostLearningAssignmentsBulkaddServiceUnavailable {
	return &PostLearningAssignmentsBulkaddServiceUnavailable{}
}

/*PostLearningAssignmentsBulkaddServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLearningAssignmentsBulkaddServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsBulkaddGatewayTimeout creates a PostLearningAssignmentsBulkaddGatewayTimeout with default headers values
func NewPostLearningAssignmentsBulkaddGatewayTimeout() *PostLearningAssignmentsBulkaddGatewayTimeout {
	return &PostLearningAssignmentsBulkaddGatewayTimeout{}
}

/*PostLearningAssignmentsBulkaddGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLearningAssignmentsBulkaddGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsBulkaddGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/bulkadd][%d] postLearningAssignmentsBulkaddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssignmentsBulkaddGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsBulkaddGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}