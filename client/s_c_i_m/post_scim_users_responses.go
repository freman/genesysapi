// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostScimUsersReader is a Reader for the PostScimUsers structure.
type PostScimUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostScimUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostScimUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostScimUsersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostScimUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostScimUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostScimUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostScimUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostScimUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostScimUsersConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostScimUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostScimUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostScimUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostScimUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostScimUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostScimUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostScimUsersOK creates a PostScimUsersOK with default headers values
func NewPostScimUsersOK() *PostScimUsersOK {
	return &PostScimUsersOK{}
}

/*
PostScimUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type PostScimUsersOK struct {
	Payload *models.ScimV2User
}

// IsSuccess returns true when this post scim users o k response has a 2xx status code
func (o *PostScimUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post scim users o k response has a 3xx status code
func (o *PostScimUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users o k response has a 4xx status code
func (o *PostScimUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post scim users o k response has a 5xx status code
func (o *PostScimUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users o k response a status code equal to that given
func (o *PostScimUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostScimUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersOK  %+v", 200, o.Payload)
}

func (o *PostScimUsersOK) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersOK  %+v", 200, o.Payload)
}

func (o *PostScimUsersOK) GetPayload() *models.ScimV2User {
	return o.Payload
}

func (o *PostScimUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersCreated creates a PostScimUsersCreated with default headers values
func NewPostScimUsersCreated() *PostScimUsersCreated {
	return &PostScimUsersCreated{}
}

/*
PostScimUsersCreated describes a response with status code 201, with default header values.

User Created.
*/
type PostScimUsersCreated struct {
	Payload *models.ScimV2User
}

// IsSuccess returns true when this post scim users created response has a 2xx status code
func (o *PostScimUsersCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post scim users created response has a 3xx status code
func (o *PostScimUsersCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users created response has a 4xx status code
func (o *PostScimUsersCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post scim users created response has a 5xx status code
func (o *PostScimUsersCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users created response a status code equal to that given
func (o *PostScimUsersCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostScimUsersCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersCreated  %+v", 201, o.Payload)
}

func (o *PostScimUsersCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersCreated  %+v", 201, o.Payload)
}

func (o *PostScimUsersCreated) GetPayload() *models.ScimV2User {
	return o.Payload
}

func (o *PostScimUsersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersBadRequest creates a PostScimUsersBadRequest with default headers values
func NewPostScimUsersBadRequest() *PostScimUsersBadRequest {
	return &PostScimUsersBadRequest{}
}

/*
PostScimUsersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostScimUsersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users bad request response has a 2xx status code
func (o *PostScimUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users bad request response has a 3xx status code
func (o *PostScimUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users bad request response has a 4xx status code
func (o *PostScimUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users bad request response has a 5xx status code
func (o *PostScimUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users bad request response a status code equal to that given
func (o *PostScimUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostScimUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostScimUsersBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostScimUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersUnauthorized creates a PostScimUsersUnauthorized with default headers values
func NewPostScimUsersUnauthorized() *PostScimUsersUnauthorized {
	return &PostScimUsersUnauthorized{}
}

/*
PostScimUsersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostScimUsersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users unauthorized response has a 2xx status code
func (o *PostScimUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users unauthorized response has a 3xx status code
func (o *PostScimUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users unauthorized response has a 4xx status code
func (o *PostScimUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users unauthorized response has a 5xx status code
func (o *PostScimUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users unauthorized response a status code equal to that given
func (o *PostScimUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostScimUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostScimUsersUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostScimUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersForbidden creates a PostScimUsersForbidden with default headers values
func NewPostScimUsersForbidden() *PostScimUsersForbidden {
	return &PostScimUsersForbidden{}
}

/*
PostScimUsersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostScimUsersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users forbidden response has a 2xx status code
func (o *PostScimUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users forbidden response has a 3xx status code
func (o *PostScimUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users forbidden response has a 4xx status code
func (o *PostScimUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users forbidden response has a 5xx status code
func (o *PostScimUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users forbidden response a status code equal to that given
func (o *PostScimUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostScimUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostScimUsersForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostScimUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersNotFound creates a PostScimUsersNotFound with default headers values
func NewPostScimUsersNotFound() *PostScimUsersNotFound {
	return &PostScimUsersNotFound{}
}

/*
PostScimUsersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostScimUsersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users not found response has a 2xx status code
func (o *PostScimUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users not found response has a 3xx status code
func (o *PostScimUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users not found response has a 4xx status code
func (o *PostScimUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users not found response has a 5xx status code
func (o *PostScimUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users not found response a status code equal to that given
func (o *PostScimUsersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostScimUsersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersNotFound  %+v", 404, o.Payload)
}

func (o *PostScimUsersNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersNotFound  %+v", 404, o.Payload)
}

func (o *PostScimUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersRequestTimeout creates a PostScimUsersRequestTimeout with default headers values
func NewPostScimUsersRequestTimeout() *PostScimUsersRequestTimeout {
	return &PostScimUsersRequestTimeout{}
}

/*
PostScimUsersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostScimUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users request timeout response has a 2xx status code
func (o *PostScimUsersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users request timeout response has a 3xx status code
func (o *PostScimUsersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users request timeout response has a 4xx status code
func (o *PostScimUsersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users request timeout response has a 5xx status code
func (o *PostScimUsersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users request timeout response a status code equal to that given
func (o *PostScimUsersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostScimUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostScimUsersRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostScimUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersConflict creates a PostScimUsersConflict with default headers values
func NewPostScimUsersConflict() *PostScimUsersConflict {
	return &PostScimUsersConflict{}
}

/*
PostScimUsersConflict describes a response with status code 409, with default header values.

User name already in use by non-deleted user.
*/
type PostScimUsersConflict struct {
	Payload *models.ScimError
}

// IsSuccess returns true when this post scim users conflict response has a 2xx status code
func (o *PostScimUsersConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users conflict response has a 3xx status code
func (o *PostScimUsersConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users conflict response has a 4xx status code
func (o *PostScimUsersConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users conflict response has a 5xx status code
func (o *PostScimUsersConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users conflict response a status code equal to that given
func (o *PostScimUsersConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostScimUsersConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersConflict  %+v", 409, o.Payload)
}

func (o *PostScimUsersConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersConflict  %+v", 409, o.Payload)
}

func (o *PostScimUsersConflict) GetPayload() *models.ScimError {
	return o.Payload
}

func (o *PostScimUsersConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersRequestEntityTooLarge creates a PostScimUsersRequestEntityTooLarge with default headers values
func NewPostScimUsersRequestEntityTooLarge() *PostScimUsersRequestEntityTooLarge {
	return &PostScimUsersRequestEntityTooLarge{}
}

/*
PostScimUsersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostScimUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users request entity too large response has a 2xx status code
func (o *PostScimUsersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users request entity too large response has a 3xx status code
func (o *PostScimUsersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users request entity too large response has a 4xx status code
func (o *PostScimUsersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users request entity too large response has a 5xx status code
func (o *PostScimUsersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users request entity too large response a status code equal to that given
func (o *PostScimUsersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostScimUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostScimUsersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostScimUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersUnsupportedMediaType creates a PostScimUsersUnsupportedMediaType with default headers values
func NewPostScimUsersUnsupportedMediaType() *PostScimUsersUnsupportedMediaType {
	return &PostScimUsersUnsupportedMediaType{}
}

/*
PostScimUsersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostScimUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users unsupported media type response has a 2xx status code
func (o *PostScimUsersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users unsupported media type response has a 3xx status code
func (o *PostScimUsersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users unsupported media type response has a 4xx status code
func (o *PostScimUsersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users unsupported media type response has a 5xx status code
func (o *PostScimUsersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users unsupported media type response a status code equal to that given
func (o *PostScimUsersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostScimUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostScimUsersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostScimUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersTooManyRequests creates a PostScimUsersTooManyRequests with default headers values
func NewPostScimUsersTooManyRequests() *PostScimUsersTooManyRequests {
	return &PostScimUsersTooManyRequests{}
}

/*
PostScimUsersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostScimUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users too many requests response has a 2xx status code
func (o *PostScimUsersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users too many requests response has a 3xx status code
func (o *PostScimUsersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users too many requests response has a 4xx status code
func (o *PostScimUsersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post scim users too many requests response has a 5xx status code
func (o *PostScimUsersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post scim users too many requests response a status code equal to that given
func (o *PostScimUsersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostScimUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostScimUsersTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostScimUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersInternalServerError creates a PostScimUsersInternalServerError with default headers values
func NewPostScimUsersInternalServerError() *PostScimUsersInternalServerError {
	return &PostScimUsersInternalServerError{}
}

/*
PostScimUsersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostScimUsersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users internal server error response has a 2xx status code
func (o *PostScimUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users internal server error response has a 3xx status code
func (o *PostScimUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users internal server error response has a 4xx status code
func (o *PostScimUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post scim users internal server error response has a 5xx status code
func (o *PostScimUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post scim users internal server error response a status code equal to that given
func (o *PostScimUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostScimUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostScimUsersInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostScimUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersServiceUnavailable creates a PostScimUsersServiceUnavailable with default headers values
func NewPostScimUsersServiceUnavailable() *PostScimUsersServiceUnavailable {
	return &PostScimUsersServiceUnavailable{}
}

/*
PostScimUsersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostScimUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users service unavailable response has a 2xx status code
func (o *PostScimUsersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users service unavailable response has a 3xx status code
func (o *PostScimUsersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users service unavailable response has a 4xx status code
func (o *PostScimUsersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post scim users service unavailable response has a 5xx status code
func (o *PostScimUsersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post scim users service unavailable response a status code equal to that given
func (o *PostScimUsersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostScimUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostScimUsersServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostScimUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostScimUsersGatewayTimeout creates a PostScimUsersGatewayTimeout with default headers values
func NewPostScimUsersGatewayTimeout() *PostScimUsersGatewayTimeout {
	return &PostScimUsersGatewayTimeout{}
}

/*
PostScimUsersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostScimUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post scim users gateway timeout response has a 2xx status code
func (o *PostScimUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post scim users gateway timeout response has a 3xx status code
func (o *PostScimUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post scim users gateway timeout response has a 4xx status code
func (o *PostScimUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post scim users gateway timeout response has a 5xx status code
func (o *PostScimUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post scim users gateway timeout response a status code equal to that given
func (o *PostScimUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostScimUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostScimUsersGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/scim/users][%d] postScimUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostScimUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostScimUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
