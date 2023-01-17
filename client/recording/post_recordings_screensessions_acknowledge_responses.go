// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostRecordingsScreensessionsAcknowledgeReader is a Reader for the PostRecordingsScreensessionsAcknowledge structure.
type PostRecordingsScreensessionsAcknowledgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordingsScreensessionsAcknowledgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostRecordingsScreensessionsAcknowledgeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRecordingsScreensessionsAcknowledgeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRecordingsScreensessionsAcknowledgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRecordingsScreensessionsAcknowledgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordingsScreensessionsAcknowledgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRecordingsScreensessionsAcknowledgeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRecordingsScreensessionsAcknowledgeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRecordingsScreensessionsAcknowledgeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordingsScreensessionsAcknowledgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRecordingsScreensessionsAcknowledgeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRecordingsScreensessionsAcknowledgeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRecordingsScreensessionsAcknowledgeNoContent creates a PostRecordingsScreensessionsAcknowledgeNoContent with default headers values
func NewPostRecordingsScreensessionsAcknowledgeNoContent() *PostRecordingsScreensessionsAcknowledgeNoContent {
	return &PostRecordingsScreensessionsAcknowledgeNoContent{}
}

/*
PostRecordingsScreensessionsAcknowledgeNoContent describes a response with status code 204, with default header values.

Recording acknowledged
*/
type PostRecordingsScreensessionsAcknowledgeNoContent struct {
}

// IsSuccess returns true when this post recordings screensessions acknowledge no content response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post recordings screensessions acknowledge no content response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge no content response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recordings screensessions acknowledge no content response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge no content response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PostRecordingsScreensessionsAcknowledgeNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeNoContent ", 204)
}

func (o *PostRecordingsScreensessionsAcknowledgeNoContent) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeNoContent ", 204)
}

func (o *PostRecordingsScreensessionsAcknowledgeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeBadRequest creates a PostRecordingsScreensessionsAcknowledgeBadRequest with default headers values
func NewPostRecordingsScreensessionsAcknowledgeBadRequest() *PostRecordingsScreensessionsAcknowledgeBadRequest {
	return &PostRecordingsScreensessionsAcknowledgeBadRequest{}
}

/*
PostRecordingsScreensessionsAcknowledgeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRecordingsScreensessionsAcknowledgeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge bad request response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge bad request response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge bad request response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge bad request response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge bad request response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeUnauthorized creates a PostRecordingsScreensessionsAcknowledgeUnauthorized with default headers values
func NewPostRecordingsScreensessionsAcknowledgeUnauthorized() *PostRecordingsScreensessionsAcknowledgeUnauthorized {
	return &PostRecordingsScreensessionsAcknowledgeUnauthorized{}
}

/*
PostRecordingsScreensessionsAcknowledgeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRecordingsScreensessionsAcknowledgeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge unauthorized response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge unauthorized response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge unauthorized response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge unauthorized response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge unauthorized response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeForbidden creates a PostRecordingsScreensessionsAcknowledgeForbidden with default headers values
func NewPostRecordingsScreensessionsAcknowledgeForbidden() *PostRecordingsScreensessionsAcknowledgeForbidden {
	return &PostRecordingsScreensessionsAcknowledgeForbidden{}
}

/*
PostRecordingsScreensessionsAcknowledgeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRecordingsScreensessionsAcknowledgeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge forbidden response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge forbidden response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge forbidden response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge forbidden response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge forbidden response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRecordingsScreensessionsAcknowledgeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeNotFound creates a PostRecordingsScreensessionsAcknowledgeNotFound with default headers values
func NewPostRecordingsScreensessionsAcknowledgeNotFound() *PostRecordingsScreensessionsAcknowledgeNotFound {
	return &PostRecordingsScreensessionsAcknowledgeNotFound{}
}

/*
PostRecordingsScreensessionsAcknowledgeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRecordingsScreensessionsAcknowledgeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge not found response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge not found response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge not found response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge not found response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge not found response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRecordingsScreensessionsAcknowledgeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeRequestTimeout creates a PostRecordingsScreensessionsAcknowledgeRequestTimeout with default headers values
func NewPostRecordingsScreensessionsAcknowledgeRequestTimeout() *PostRecordingsScreensessionsAcknowledgeRequestTimeout {
	return &PostRecordingsScreensessionsAcknowledgeRequestTimeout{}
}

/*
PostRecordingsScreensessionsAcknowledgeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRecordingsScreensessionsAcknowledgeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge request timeout response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge request timeout response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge request timeout response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge request timeout response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge request timeout response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge creates a PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge with default headers values
func NewPostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge() *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge {
	return &PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge{}
}

/*
PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge request entity too large response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge request entity too large response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge request entity too large response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge request entity too large response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge request entity too large response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeUnsupportedMediaType creates a PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType with default headers values
func NewPostRecordingsScreensessionsAcknowledgeUnsupportedMediaType() *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType {
	return &PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType{}
}

/*
PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge unsupported media type response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge unsupported media type response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge unsupported media type response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge unsupported media type response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge unsupported media type response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeTooManyRequests creates a PostRecordingsScreensessionsAcknowledgeTooManyRequests with default headers values
func NewPostRecordingsScreensessionsAcknowledgeTooManyRequests() *PostRecordingsScreensessionsAcknowledgeTooManyRequests {
	return &PostRecordingsScreensessionsAcknowledgeTooManyRequests{}
}

/*
PostRecordingsScreensessionsAcknowledgeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRecordingsScreensessionsAcknowledgeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge too many requests response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge too many requests response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge too many requests response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recordings screensessions acknowledge too many requests response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post recordings screensessions acknowledge too many requests response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeInternalServerError creates a PostRecordingsScreensessionsAcknowledgeInternalServerError with default headers values
func NewPostRecordingsScreensessionsAcknowledgeInternalServerError() *PostRecordingsScreensessionsAcknowledgeInternalServerError {
	return &PostRecordingsScreensessionsAcknowledgeInternalServerError{}
}

/*
PostRecordingsScreensessionsAcknowledgeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRecordingsScreensessionsAcknowledgeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge internal server error response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge internal server error response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge internal server error response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recordings screensessions acknowledge internal server error response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post recordings screensessions acknowledge internal server error response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeServiceUnavailable creates a PostRecordingsScreensessionsAcknowledgeServiceUnavailable with default headers values
func NewPostRecordingsScreensessionsAcknowledgeServiceUnavailable() *PostRecordingsScreensessionsAcknowledgeServiceUnavailable {
	return &PostRecordingsScreensessionsAcknowledgeServiceUnavailable{}
}

/*
PostRecordingsScreensessionsAcknowledgeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRecordingsScreensessionsAcknowledgeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge service unavailable response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge service unavailable response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge service unavailable response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recordings screensessions acknowledge service unavailable response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post recordings screensessions acknowledge service unavailable response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsAcknowledgeGatewayTimeout creates a PostRecordingsScreensessionsAcknowledgeGatewayTimeout with default headers values
func NewPostRecordingsScreensessionsAcknowledgeGatewayTimeout() *PostRecordingsScreensessionsAcknowledgeGatewayTimeout {
	return &PostRecordingsScreensessionsAcknowledgeGatewayTimeout{}
}

/*
PostRecordingsScreensessionsAcknowledgeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRecordingsScreensessionsAcknowledgeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recordings screensessions acknowledge gateway timeout response has a 2xx status code
func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recordings screensessions acknowledge gateway timeout response has a 3xx status code
func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recordings screensessions acknowledge gateway timeout response has a 4xx status code
func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recordings screensessions acknowledge gateway timeout response has a 5xx status code
func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post recordings screensessions acknowledge gateway timeout response a status code equal to that given
func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/acknowledge][%d] postRecordingsScreensessionsAcknowledgeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsAcknowledgeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
