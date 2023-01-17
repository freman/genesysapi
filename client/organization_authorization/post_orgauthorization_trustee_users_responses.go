// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOrgauthorizationTrusteeUsersReader is a Reader for the PostOrgauthorizationTrusteeUsers structure.
type PostOrgauthorizationTrusteeUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrgauthorizationTrusteeUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOrgauthorizationTrusteeUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOrgauthorizationTrusteeUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOrgauthorizationTrusteeUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOrgauthorizationTrusteeUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOrgauthorizationTrusteeUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOrgauthorizationTrusteeUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOrgauthorizationTrusteeUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOrgauthorizationTrusteeUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOrgauthorizationTrusteeUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOrgauthorizationTrusteeUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOrgauthorizationTrusteeUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOrgauthorizationTrusteeUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrgauthorizationTrusteeUsersOK creates a PostOrgauthorizationTrusteeUsersOK with default headers values
func NewPostOrgauthorizationTrusteeUsersOK() *PostOrgauthorizationTrusteeUsersOK {
	return &PostOrgauthorizationTrusteeUsersOK{}
}

/*
PostOrgauthorizationTrusteeUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOrgauthorizationTrusteeUsersOK struct {
	Payload *models.TrustUser
}

// IsSuccess returns true when this post orgauthorization trustee users o k response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post orgauthorization trustee users o k response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users o k response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee users o k response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users o k response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOrgauthorizationTrusteeUsersOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersOK) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersOK) GetPayload() *models.TrustUser {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrustUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersBadRequest creates a PostOrgauthorizationTrusteeUsersBadRequest with default headers values
func NewPostOrgauthorizationTrusteeUsersBadRequest() *PostOrgauthorizationTrusteeUsersBadRequest {
	return &PostOrgauthorizationTrusteeUsersBadRequest{}
}

/*
PostOrgauthorizationTrusteeUsersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOrgauthorizationTrusteeUsersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users bad request response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users bad request response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users bad request response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users bad request response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users bad request response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOrgauthorizationTrusteeUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersUnauthorized creates a PostOrgauthorizationTrusteeUsersUnauthorized with default headers values
func NewPostOrgauthorizationTrusteeUsersUnauthorized() *PostOrgauthorizationTrusteeUsersUnauthorized {
	return &PostOrgauthorizationTrusteeUsersUnauthorized{}
}

/*
PostOrgauthorizationTrusteeUsersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOrgauthorizationTrusteeUsersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users unauthorized response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users unauthorized response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users unauthorized response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users unauthorized response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users unauthorized response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOrgauthorizationTrusteeUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersForbidden creates a PostOrgauthorizationTrusteeUsersForbidden with default headers values
func NewPostOrgauthorizationTrusteeUsersForbidden() *PostOrgauthorizationTrusteeUsersForbidden {
	return &PostOrgauthorizationTrusteeUsersForbidden{}
}

/*
PostOrgauthorizationTrusteeUsersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOrgauthorizationTrusteeUsersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users forbidden response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users forbidden response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users forbidden response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users forbidden response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users forbidden response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOrgauthorizationTrusteeUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersNotFound creates a PostOrgauthorizationTrusteeUsersNotFound with default headers values
func NewPostOrgauthorizationTrusteeUsersNotFound() *PostOrgauthorizationTrusteeUsersNotFound {
	return &PostOrgauthorizationTrusteeUsersNotFound{}
}

/*
PostOrgauthorizationTrusteeUsersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOrgauthorizationTrusteeUsersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users not found response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users not found response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users not found response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users not found response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users not found response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOrgauthorizationTrusteeUsersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersRequestTimeout creates a PostOrgauthorizationTrusteeUsersRequestTimeout with default headers values
func NewPostOrgauthorizationTrusteeUsersRequestTimeout() *PostOrgauthorizationTrusteeUsersRequestTimeout {
	return &PostOrgauthorizationTrusteeUsersRequestTimeout{}
}

/*
PostOrgauthorizationTrusteeUsersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOrgauthorizationTrusteeUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users request timeout response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users request timeout response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users request timeout response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users request timeout response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users request timeout response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersRequestEntityTooLarge creates a PostOrgauthorizationTrusteeUsersRequestEntityTooLarge with default headers values
func NewPostOrgauthorizationTrusteeUsersRequestEntityTooLarge() *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge {
	return &PostOrgauthorizationTrusteeUsersRequestEntityTooLarge{}
}

/*
PostOrgauthorizationTrusteeUsersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostOrgauthorizationTrusteeUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users request entity too large response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users request entity too large response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users request entity too large response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users request entity too large response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users request entity too large response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersUnsupportedMediaType creates a PostOrgauthorizationTrusteeUsersUnsupportedMediaType with default headers values
func NewPostOrgauthorizationTrusteeUsersUnsupportedMediaType() *PostOrgauthorizationTrusteeUsersUnsupportedMediaType {
	return &PostOrgauthorizationTrusteeUsersUnsupportedMediaType{}
}

/*
PostOrgauthorizationTrusteeUsersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOrgauthorizationTrusteeUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users unsupported media type response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users unsupported media type response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users unsupported media type response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users unsupported media type response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users unsupported media type response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersTooManyRequests creates a PostOrgauthorizationTrusteeUsersTooManyRequests with default headers values
func NewPostOrgauthorizationTrusteeUsersTooManyRequests() *PostOrgauthorizationTrusteeUsersTooManyRequests {
	return &PostOrgauthorizationTrusteeUsersTooManyRequests{}
}

/*
PostOrgauthorizationTrusteeUsersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOrgauthorizationTrusteeUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users too many requests response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users too many requests response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users too many requests response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee users too many requests response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee users too many requests response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersInternalServerError creates a PostOrgauthorizationTrusteeUsersInternalServerError with default headers values
func NewPostOrgauthorizationTrusteeUsersInternalServerError() *PostOrgauthorizationTrusteeUsersInternalServerError {
	return &PostOrgauthorizationTrusteeUsersInternalServerError{}
}

/*
PostOrgauthorizationTrusteeUsersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOrgauthorizationTrusteeUsersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users internal server error response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users internal server error response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users internal server error response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee users internal server error response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post orgauthorization trustee users internal server error response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOrgauthorizationTrusteeUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersServiceUnavailable creates a PostOrgauthorizationTrusteeUsersServiceUnavailable with default headers values
func NewPostOrgauthorizationTrusteeUsersServiceUnavailable() *PostOrgauthorizationTrusteeUsersServiceUnavailable {
	return &PostOrgauthorizationTrusteeUsersServiceUnavailable{}
}

/*
PostOrgauthorizationTrusteeUsersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOrgauthorizationTrusteeUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users service unavailable response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users service unavailable response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users service unavailable response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee users service unavailable response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post orgauthorization trustee users service unavailable response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeUsersGatewayTimeout creates a PostOrgauthorizationTrusteeUsersGatewayTimeout with default headers values
func NewPostOrgauthorizationTrusteeUsersGatewayTimeout() *PostOrgauthorizationTrusteeUsersGatewayTimeout {
	return &PostOrgauthorizationTrusteeUsersGatewayTimeout{}
}

/*
PostOrgauthorizationTrusteeUsersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOrgauthorizationTrusteeUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee users gateway timeout response has a 2xx status code
func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee users gateway timeout response has a 3xx status code
func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee users gateway timeout response has a 4xx status code
func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee users gateway timeout response has a 5xx status code
func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post orgauthorization trustee users gateway timeout response a status code equal to that given
func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] postOrgauthorizationTrusteeUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
