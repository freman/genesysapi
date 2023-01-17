// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostUserExternalidReader is a Reader for the PostUserExternalid structure.
type PostUserExternalidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserExternalidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserExternalidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostUserExternalidCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserExternalidBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserExternalidUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserExternalidForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserExternalidNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUserExternalidRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostUserExternalidConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUserExternalidRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUserExternalidUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUserExternalidTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserExternalidInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUserExternalidServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUserExternalidGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUserExternalidOK creates a PostUserExternalidOK with default headers values
func NewPostUserExternalidOK() *PostUserExternalidOK {
	return &PostUserExternalidOK{}
}

/*
PostUserExternalidOK describes a response with status code 200, with default header values.

successful operation
*/
type PostUserExternalidOK struct {
	Payload []*models.UserExternalIdentifier
}

// IsSuccess returns true when this post user externalid o k response has a 2xx status code
func (o *PostUserExternalidOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post user externalid o k response has a 3xx status code
func (o *PostUserExternalidOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid o k response has a 4xx status code
func (o *PostUserExternalidOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user externalid o k response has a 5xx status code
func (o *PostUserExternalidOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid o k response a status code equal to that given
func (o *PostUserExternalidOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostUserExternalidOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidOK  %+v", 200, o.Payload)
}

func (o *PostUserExternalidOK) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidOK  %+v", 200, o.Payload)
}

func (o *PostUserExternalidOK) GetPayload() []*models.UserExternalIdentifier {
	return o.Payload
}

func (o *PostUserExternalidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidCreated creates a PostUserExternalidCreated with default headers values
func NewPostUserExternalidCreated() *PostUserExternalidCreated {
	return &PostUserExternalidCreated{}
}

/*
PostUserExternalidCreated describes a response with status code 201, with default header values.

External Identifier Created
*/
type PostUserExternalidCreated struct {
	Payload []*models.UserExternalIdentifier
}

// IsSuccess returns true when this post user externalid created response has a 2xx status code
func (o *PostUserExternalidCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post user externalid created response has a 3xx status code
func (o *PostUserExternalidCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid created response has a 4xx status code
func (o *PostUserExternalidCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user externalid created response has a 5xx status code
func (o *PostUserExternalidCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid created response a status code equal to that given
func (o *PostUserExternalidCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostUserExternalidCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidCreated  %+v", 201, o.Payload)
}

func (o *PostUserExternalidCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidCreated  %+v", 201, o.Payload)
}

func (o *PostUserExternalidCreated) GetPayload() []*models.UserExternalIdentifier {
	return o.Payload
}

func (o *PostUserExternalidCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidBadRequest creates a PostUserExternalidBadRequest with default headers values
func NewPostUserExternalidBadRequest() *PostUserExternalidBadRequest {
	return &PostUserExternalidBadRequest{}
}

/*
PostUserExternalidBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUserExternalidBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid bad request response has a 2xx status code
func (o *PostUserExternalidBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid bad request response has a 3xx status code
func (o *PostUserExternalidBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid bad request response has a 4xx status code
func (o *PostUserExternalidBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid bad request response has a 5xx status code
func (o *PostUserExternalidBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid bad request response a status code equal to that given
func (o *PostUserExternalidBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUserExternalidBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserExternalidBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserExternalidBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidUnauthorized creates a PostUserExternalidUnauthorized with default headers values
func NewPostUserExternalidUnauthorized() *PostUserExternalidUnauthorized {
	return &PostUserExternalidUnauthorized{}
}

/*
PostUserExternalidUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUserExternalidUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid unauthorized response has a 2xx status code
func (o *PostUserExternalidUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid unauthorized response has a 3xx status code
func (o *PostUserExternalidUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid unauthorized response has a 4xx status code
func (o *PostUserExternalidUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid unauthorized response has a 5xx status code
func (o *PostUserExternalidUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid unauthorized response a status code equal to that given
func (o *PostUserExternalidUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUserExternalidUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserExternalidUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserExternalidUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidForbidden creates a PostUserExternalidForbidden with default headers values
func NewPostUserExternalidForbidden() *PostUserExternalidForbidden {
	return &PostUserExternalidForbidden{}
}

/*
PostUserExternalidForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostUserExternalidForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid forbidden response has a 2xx status code
func (o *PostUserExternalidForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid forbidden response has a 3xx status code
func (o *PostUserExternalidForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid forbidden response has a 4xx status code
func (o *PostUserExternalidForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid forbidden response has a 5xx status code
func (o *PostUserExternalidForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid forbidden response a status code equal to that given
func (o *PostUserExternalidForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUserExternalidForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidForbidden  %+v", 403, o.Payload)
}

func (o *PostUserExternalidForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidForbidden  %+v", 403, o.Payload)
}

func (o *PostUserExternalidForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidNotFound creates a PostUserExternalidNotFound with default headers values
func NewPostUserExternalidNotFound() *PostUserExternalidNotFound {
	return &PostUserExternalidNotFound{}
}

/*
PostUserExternalidNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostUserExternalidNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid not found response has a 2xx status code
func (o *PostUserExternalidNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid not found response has a 3xx status code
func (o *PostUserExternalidNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid not found response has a 4xx status code
func (o *PostUserExternalidNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid not found response has a 5xx status code
func (o *PostUserExternalidNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid not found response a status code equal to that given
func (o *PostUserExternalidNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUserExternalidNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidNotFound  %+v", 404, o.Payload)
}

func (o *PostUserExternalidNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidNotFound  %+v", 404, o.Payload)
}

func (o *PostUserExternalidNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidRequestTimeout creates a PostUserExternalidRequestTimeout with default headers values
func NewPostUserExternalidRequestTimeout() *PostUserExternalidRequestTimeout {
	return &PostUserExternalidRequestTimeout{}
}

/*
PostUserExternalidRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUserExternalidRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid request timeout response has a 2xx status code
func (o *PostUserExternalidRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid request timeout response has a 3xx status code
func (o *PostUserExternalidRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid request timeout response has a 4xx status code
func (o *PostUserExternalidRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid request timeout response has a 5xx status code
func (o *PostUserExternalidRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid request timeout response a status code equal to that given
func (o *PostUserExternalidRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostUserExternalidRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUserExternalidRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUserExternalidRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidConflict creates a PostUserExternalidConflict with default headers values
func NewPostUserExternalidConflict() *PostUserExternalidConflict {
	return &PostUserExternalidConflict{}
}

/*
PostUserExternalidConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostUserExternalidConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid conflict response has a 2xx status code
func (o *PostUserExternalidConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid conflict response has a 3xx status code
func (o *PostUserExternalidConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid conflict response has a 4xx status code
func (o *PostUserExternalidConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid conflict response has a 5xx status code
func (o *PostUserExternalidConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid conflict response a status code equal to that given
func (o *PostUserExternalidConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostUserExternalidConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidConflict  %+v", 409, o.Payload)
}

func (o *PostUserExternalidConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidConflict  %+v", 409, o.Payload)
}

func (o *PostUserExternalidConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidRequestEntityTooLarge creates a PostUserExternalidRequestEntityTooLarge with default headers values
func NewPostUserExternalidRequestEntityTooLarge() *PostUserExternalidRequestEntityTooLarge {
	return &PostUserExternalidRequestEntityTooLarge{}
}

/*
PostUserExternalidRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostUserExternalidRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid request entity too large response has a 2xx status code
func (o *PostUserExternalidRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid request entity too large response has a 3xx status code
func (o *PostUserExternalidRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid request entity too large response has a 4xx status code
func (o *PostUserExternalidRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid request entity too large response has a 5xx status code
func (o *PostUserExternalidRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid request entity too large response a status code equal to that given
func (o *PostUserExternalidRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostUserExternalidRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserExternalidRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUserExternalidRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidUnsupportedMediaType creates a PostUserExternalidUnsupportedMediaType with default headers values
func NewPostUserExternalidUnsupportedMediaType() *PostUserExternalidUnsupportedMediaType {
	return &PostUserExternalidUnsupportedMediaType{}
}

/*
PostUserExternalidUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUserExternalidUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid unsupported media type response has a 2xx status code
func (o *PostUserExternalidUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid unsupported media type response has a 3xx status code
func (o *PostUserExternalidUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid unsupported media type response has a 4xx status code
func (o *PostUserExternalidUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid unsupported media type response has a 5xx status code
func (o *PostUserExternalidUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid unsupported media type response a status code equal to that given
func (o *PostUserExternalidUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostUserExternalidUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserExternalidUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUserExternalidUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidTooManyRequests creates a PostUserExternalidTooManyRequests with default headers values
func NewPostUserExternalidTooManyRequests() *PostUserExternalidTooManyRequests {
	return &PostUserExternalidTooManyRequests{}
}

/*
PostUserExternalidTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUserExternalidTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid too many requests response has a 2xx status code
func (o *PostUserExternalidTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid too many requests response has a 3xx status code
func (o *PostUserExternalidTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid too many requests response has a 4xx status code
func (o *PostUserExternalidTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post user externalid too many requests response has a 5xx status code
func (o *PostUserExternalidTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post user externalid too many requests response a status code equal to that given
func (o *PostUserExternalidTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostUserExternalidTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserExternalidTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUserExternalidTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidInternalServerError creates a PostUserExternalidInternalServerError with default headers values
func NewPostUserExternalidInternalServerError() *PostUserExternalidInternalServerError {
	return &PostUserExternalidInternalServerError{}
}

/*
PostUserExternalidInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUserExternalidInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid internal server error response has a 2xx status code
func (o *PostUserExternalidInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid internal server error response has a 3xx status code
func (o *PostUserExternalidInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid internal server error response has a 4xx status code
func (o *PostUserExternalidInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user externalid internal server error response has a 5xx status code
func (o *PostUserExternalidInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post user externalid internal server error response a status code equal to that given
func (o *PostUserExternalidInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUserExternalidInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserExternalidInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserExternalidInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidServiceUnavailable creates a PostUserExternalidServiceUnavailable with default headers values
func NewPostUserExternalidServiceUnavailable() *PostUserExternalidServiceUnavailable {
	return &PostUserExternalidServiceUnavailable{}
}

/*
PostUserExternalidServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUserExternalidServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid service unavailable response has a 2xx status code
func (o *PostUserExternalidServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid service unavailable response has a 3xx status code
func (o *PostUserExternalidServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid service unavailable response has a 4xx status code
func (o *PostUserExternalidServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user externalid service unavailable response has a 5xx status code
func (o *PostUserExternalidServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post user externalid service unavailable response a status code equal to that given
func (o *PostUserExternalidServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUserExternalidServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserExternalidServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUserExternalidServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserExternalidGatewayTimeout creates a PostUserExternalidGatewayTimeout with default headers values
func NewPostUserExternalidGatewayTimeout() *PostUserExternalidGatewayTimeout {
	return &PostUserExternalidGatewayTimeout{}
}

/*
PostUserExternalidGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostUserExternalidGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post user externalid gateway timeout response has a 2xx status code
func (o *PostUserExternalidGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post user externalid gateway timeout response has a 3xx status code
func (o *PostUserExternalidGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post user externalid gateway timeout response has a 4xx status code
func (o *PostUserExternalidGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post user externalid gateway timeout response has a 5xx status code
func (o *PostUserExternalidGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post user externalid gateway timeout response a status code equal to that given
func (o *PostUserExternalidGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUserExternalidGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserExternalidGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/{userId}/externalid][%d] postUserExternalidGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUserExternalidGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUserExternalidGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
