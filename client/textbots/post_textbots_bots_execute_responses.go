// Code generated by go-swagger; DO NOT EDIT.

package textbots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostTextbotsBotsExecuteReader is a Reader for the PostTextbotsBotsExecute structure.
type PostTextbotsBotsExecuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTextbotsBotsExecuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTextbotsBotsExecuteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTextbotsBotsExecuteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTextbotsBotsExecuteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTextbotsBotsExecuteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTextbotsBotsExecuteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTextbotsBotsExecuteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTextbotsBotsExecuteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTextbotsBotsExecuteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTextbotsBotsExecuteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTextbotsBotsExecuteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTextbotsBotsExecuteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTextbotsBotsExecuteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTextbotsBotsExecuteOK creates a PostTextbotsBotsExecuteOK with default headers values
func NewPostTextbotsBotsExecuteOK() *PostTextbotsBotsExecuteOK {
	return &PostTextbotsBotsExecuteOK{}
}

/*
PostTextbotsBotsExecuteOK describes a response with status code 200, with default header values.

successful operation
*/
type PostTextbotsBotsExecuteOK struct {
	Payload *models.PostTextResponse
}

// IsSuccess returns true when this post textbots bots execute o k response has a 2xx status code
func (o *PostTextbotsBotsExecuteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post textbots bots execute o k response has a 3xx status code
func (o *PostTextbotsBotsExecuteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute o k response has a 4xx status code
func (o *PostTextbotsBotsExecuteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post textbots bots execute o k response has a 5xx status code
func (o *PostTextbotsBotsExecuteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute o k response a status code equal to that given
func (o *PostTextbotsBotsExecuteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostTextbotsBotsExecuteOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteOK  %+v", 200, o.Payload)
}

func (o *PostTextbotsBotsExecuteOK) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteOK  %+v", 200, o.Payload)
}

func (o *PostTextbotsBotsExecuteOK) GetPayload() *models.PostTextResponse {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostTextResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteBadRequest creates a PostTextbotsBotsExecuteBadRequest with default headers values
func NewPostTextbotsBotsExecuteBadRequest() *PostTextbotsBotsExecuteBadRequest {
	return &PostTextbotsBotsExecuteBadRequest{}
}

/*
PostTextbotsBotsExecuteBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTextbotsBotsExecuteBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute bad request response has a 2xx status code
func (o *PostTextbotsBotsExecuteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute bad request response has a 3xx status code
func (o *PostTextbotsBotsExecuteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute bad request response has a 4xx status code
func (o *PostTextbotsBotsExecuteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute bad request response has a 5xx status code
func (o *PostTextbotsBotsExecuteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute bad request response a status code equal to that given
func (o *PostTextbotsBotsExecuteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostTextbotsBotsExecuteBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteBadRequest  %+v", 400, o.Payload)
}

func (o *PostTextbotsBotsExecuteBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteBadRequest  %+v", 400, o.Payload)
}

func (o *PostTextbotsBotsExecuteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteUnauthorized creates a PostTextbotsBotsExecuteUnauthorized with default headers values
func NewPostTextbotsBotsExecuteUnauthorized() *PostTextbotsBotsExecuteUnauthorized {
	return &PostTextbotsBotsExecuteUnauthorized{}
}

/*
PostTextbotsBotsExecuteUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTextbotsBotsExecuteUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute unauthorized response has a 2xx status code
func (o *PostTextbotsBotsExecuteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute unauthorized response has a 3xx status code
func (o *PostTextbotsBotsExecuteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute unauthorized response has a 4xx status code
func (o *PostTextbotsBotsExecuteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute unauthorized response has a 5xx status code
func (o *PostTextbotsBotsExecuteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute unauthorized response a status code equal to that given
func (o *PostTextbotsBotsExecuteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostTextbotsBotsExecuteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTextbotsBotsExecuteUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTextbotsBotsExecuteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteForbidden creates a PostTextbotsBotsExecuteForbidden with default headers values
func NewPostTextbotsBotsExecuteForbidden() *PostTextbotsBotsExecuteForbidden {
	return &PostTextbotsBotsExecuteForbidden{}
}

/*
PostTextbotsBotsExecuteForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostTextbotsBotsExecuteForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute forbidden response has a 2xx status code
func (o *PostTextbotsBotsExecuteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute forbidden response has a 3xx status code
func (o *PostTextbotsBotsExecuteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute forbidden response has a 4xx status code
func (o *PostTextbotsBotsExecuteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute forbidden response has a 5xx status code
func (o *PostTextbotsBotsExecuteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute forbidden response a status code equal to that given
func (o *PostTextbotsBotsExecuteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostTextbotsBotsExecuteForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteForbidden  %+v", 403, o.Payload)
}

func (o *PostTextbotsBotsExecuteForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteForbidden  %+v", 403, o.Payload)
}

func (o *PostTextbotsBotsExecuteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteNotFound creates a PostTextbotsBotsExecuteNotFound with default headers values
func NewPostTextbotsBotsExecuteNotFound() *PostTextbotsBotsExecuteNotFound {
	return &PostTextbotsBotsExecuteNotFound{}
}

/*
PostTextbotsBotsExecuteNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostTextbotsBotsExecuteNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute not found response has a 2xx status code
func (o *PostTextbotsBotsExecuteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute not found response has a 3xx status code
func (o *PostTextbotsBotsExecuteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute not found response has a 4xx status code
func (o *PostTextbotsBotsExecuteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute not found response has a 5xx status code
func (o *PostTextbotsBotsExecuteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute not found response a status code equal to that given
func (o *PostTextbotsBotsExecuteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostTextbotsBotsExecuteNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteNotFound  %+v", 404, o.Payload)
}

func (o *PostTextbotsBotsExecuteNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteNotFound  %+v", 404, o.Payload)
}

func (o *PostTextbotsBotsExecuteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteRequestTimeout creates a PostTextbotsBotsExecuteRequestTimeout with default headers values
func NewPostTextbotsBotsExecuteRequestTimeout() *PostTextbotsBotsExecuteRequestTimeout {
	return &PostTextbotsBotsExecuteRequestTimeout{}
}

/*
PostTextbotsBotsExecuteRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTextbotsBotsExecuteRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute request timeout response has a 2xx status code
func (o *PostTextbotsBotsExecuteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute request timeout response has a 3xx status code
func (o *PostTextbotsBotsExecuteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute request timeout response has a 4xx status code
func (o *PostTextbotsBotsExecuteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute request timeout response has a 5xx status code
func (o *PostTextbotsBotsExecuteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute request timeout response a status code equal to that given
func (o *PostTextbotsBotsExecuteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostTextbotsBotsExecuteRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTextbotsBotsExecuteRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTextbotsBotsExecuteRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteRequestEntityTooLarge creates a PostTextbotsBotsExecuteRequestEntityTooLarge with default headers values
func NewPostTextbotsBotsExecuteRequestEntityTooLarge() *PostTextbotsBotsExecuteRequestEntityTooLarge {
	return &PostTextbotsBotsExecuteRequestEntityTooLarge{}
}

/*
PostTextbotsBotsExecuteRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostTextbotsBotsExecuteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute request entity too large response has a 2xx status code
func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute request entity too large response has a 3xx status code
func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute request entity too large response has a 4xx status code
func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute request entity too large response has a 5xx status code
func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute request entity too large response a status code equal to that given
func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteUnsupportedMediaType creates a PostTextbotsBotsExecuteUnsupportedMediaType with default headers values
func NewPostTextbotsBotsExecuteUnsupportedMediaType() *PostTextbotsBotsExecuteUnsupportedMediaType {
	return &PostTextbotsBotsExecuteUnsupportedMediaType{}
}

/*
PostTextbotsBotsExecuteUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTextbotsBotsExecuteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute unsupported media type response has a 2xx status code
func (o *PostTextbotsBotsExecuteUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute unsupported media type response has a 3xx status code
func (o *PostTextbotsBotsExecuteUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute unsupported media type response has a 4xx status code
func (o *PostTextbotsBotsExecuteUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute unsupported media type response has a 5xx status code
func (o *PostTextbotsBotsExecuteUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute unsupported media type response a status code equal to that given
func (o *PostTextbotsBotsExecuteUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostTextbotsBotsExecuteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTextbotsBotsExecuteUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTextbotsBotsExecuteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteTooManyRequests creates a PostTextbotsBotsExecuteTooManyRequests with default headers values
func NewPostTextbotsBotsExecuteTooManyRequests() *PostTextbotsBotsExecuteTooManyRequests {
	return &PostTextbotsBotsExecuteTooManyRequests{}
}

/*
PostTextbotsBotsExecuteTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTextbotsBotsExecuteTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute too many requests response has a 2xx status code
func (o *PostTextbotsBotsExecuteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute too many requests response has a 3xx status code
func (o *PostTextbotsBotsExecuteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute too many requests response has a 4xx status code
func (o *PostTextbotsBotsExecuteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post textbots bots execute too many requests response has a 5xx status code
func (o *PostTextbotsBotsExecuteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post textbots bots execute too many requests response a status code equal to that given
func (o *PostTextbotsBotsExecuteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostTextbotsBotsExecuteTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTextbotsBotsExecuteTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTextbotsBotsExecuteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteInternalServerError creates a PostTextbotsBotsExecuteInternalServerError with default headers values
func NewPostTextbotsBotsExecuteInternalServerError() *PostTextbotsBotsExecuteInternalServerError {
	return &PostTextbotsBotsExecuteInternalServerError{}
}

/*
PostTextbotsBotsExecuteInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTextbotsBotsExecuteInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute internal server error response has a 2xx status code
func (o *PostTextbotsBotsExecuteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute internal server error response has a 3xx status code
func (o *PostTextbotsBotsExecuteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute internal server error response has a 4xx status code
func (o *PostTextbotsBotsExecuteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post textbots bots execute internal server error response has a 5xx status code
func (o *PostTextbotsBotsExecuteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post textbots bots execute internal server error response a status code equal to that given
func (o *PostTextbotsBotsExecuteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostTextbotsBotsExecuteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTextbotsBotsExecuteInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTextbotsBotsExecuteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteServiceUnavailable creates a PostTextbotsBotsExecuteServiceUnavailable with default headers values
func NewPostTextbotsBotsExecuteServiceUnavailable() *PostTextbotsBotsExecuteServiceUnavailable {
	return &PostTextbotsBotsExecuteServiceUnavailable{}
}

/*
PostTextbotsBotsExecuteServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTextbotsBotsExecuteServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute service unavailable response has a 2xx status code
func (o *PostTextbotsBotsExecuteServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute service unavailable response has a 3xx status code
func (o *PostTextbotsBotsExecuteServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute service unavailable response has a 4xx status code
func (o *PostTextbotsBotsExecuteServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post textbots bots execute service unavailable response has a 5xx status code
func (o *PostTextbotsBotsExecuteServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post textbots bots execute service unavailable response a status code equal to that given
func (o *PostTextbotsBotsExecuteServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostTextbotsBotsExecuteServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTextbotsBotsExecuteServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTextbotsBotsExecuteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTextbotsBotsExecuteGatewayTimeout creates a PostTextbotsBotsExecuteGatewayTimeout with default headers values
func NewPostTextbotsBotsExecuteGatewayTimeout() *PostTextbotsBotsExecuteGatewayTimeout {
	return &PostTextbotsBotsExecuteGatewayTimeout{}
}

/*
PostTextbotsBotsExecuteGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostTextbotsBotsExecuteGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post textbots bots execute gateway timeout response has a 2xx status code
func (o *PostTextbotsBotsExecuteGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post textbots bots execute gateway timeout response has a 3xx status code
func (o *PostTextbotsBotsExecuteGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post textbots bots execute gateway timeout response has a 4xx status code
func (o *PostTextbotsBotsExecuteGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post textbots bots execute gateway timeout response has a 5xx status code
func (o *PostTextbotsBotsExecuteGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post textbots bots execute gateway timeout response a status code equal to that given
func (o *PostTextbotsBotsExecuteGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostTextbotsBotsExecuteGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTextbotsBotsExecuteGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/textbots/bots/execute][%d] postTextbotsBotsExecuteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTextbotsBotsExecuteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTextbotsBotsExecuteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
