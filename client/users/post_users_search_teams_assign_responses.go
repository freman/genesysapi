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

// PostUsersSearchTeamsAssignReader is a Reader for the PostUsersSearchTeamsAssign structure.
type PostUsersSearchTeamsAssignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersSearchTeamsAssignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUsersSearchTeamsAssignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersSearchTeamsAssignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsersSearchTeamsAssignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersSearchTeamsAssignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUsersSearchTeamsAssignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUsersSearchTeamsAssignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUsersSearchTeamsAssignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUsersSearchTeamsAssignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUsersSearchTeamsAssignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersSearchTeamsAssignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUsersSearchTeamsAssignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUsersSearchTeamsAssignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsersSearchTeamsAssignOK creates a PostUsersSearchTeamsAssignOK with default headers values
func NewPostUsersSearchTeamsAssignOK() *PostUsersSearchTeamsAssignOK {
	return &PostUsersSearchTeamsAssignOK{}
}

/*
PostUsersSearchTeamsAssignOK describes a response with status code 200, with default header values.

successful operation
*/
type PostUsersSearchTeamsAssignOK struct {
	Payload *models.UsersSearchResponse
}

// IsSuccess returns true when this post users search teams assign o k response has a 2xx status code
func (o *PostUsersSearchTeamsAssignOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post users search teams assign o k response has a 3xx status code
func (o *PostUsersSearchTeamsAssignOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign o k response has a 4xx status code
func (o *PostUsersSearchTeamsAssignOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search teams assign o k response has a 5xx status code
func (o *PostUsersSearchTeamsAssignOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign o k response a status code equal to that given
func (o *PostUsersSearchTeamsAssignOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostUsersSearchTeamsAssignOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignOK  %+v", 200, o.Payload)
}

func (o *PostUsersSearchTeamsAssignOK) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignOK  %+v", 200, o.Payload)
}

func (o *PostUsersSearchTeamsAssignOK) GetPayload() *models.UsersSearchResponse {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignBadRequest creates a PostUsersSearchTeamsAssignBadRequest with default headers values
func NewPostUsersSearchTeamsAssignBadRequest() *PostUsersSearchTeamsAssignBadRequest {
	return &PostUsersSearchTeamsAssignBadRequest{}
}

/*
PostUsersSearchTeamsAssignBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUsersSearchTeamsAssignBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign bad request response has a 2xx status code
func (o *PostUsersSearchTeamsAssignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign bad request response has a 3xx status code
func (o *PostUsersSearchTeamsAssignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign bad request response has a 4xx status code
func (o *PostUsersSearchTeamsAssignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign bad request response has a 5xx status code
func (o *PostUsersSearchTeamsAssignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign bad request response a status code equal to that given
func (o *PostUsersSearchTeamsAssignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUsersSearchTeamsAssignBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersSearchTeamsAssignBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersSearchTeamsAssignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignUnauthorized creates a PostUsersSearchTeamsAssignUnauthorized with default headers values
func NewPostUsersSearchTeamsAssignUnauthorized() *PostUsersSearchTeamsAssignUnauthorized {
	return &PostUsersSearchTeamsAssignUnauthorized{}
}

/*
PostUsersSearchTeamsAssignUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUsersSearchTeamsAssignUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign unauthorized response has a 2xx status code
func (o *PostUsersSearchTeamsAssignUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign unauthorized response has a 3xx status code
func (o *PostUsersSearchTeamsAssignUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign unauthorized response has a 4xx status code
func (o *PostUsersSearchTeamsAssignUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign unauthorized response has a 5xx status code
func (o *PostUsersSearchTeamsAssignUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign unauthorized response a status code equal to that given
func (o *PostUsersSearchTeamsAssignUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUsersSearchTeamsAssignUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersSearchTeamsAssignUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersSearchTeamsAssignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignForbidden creates a PostUsersSearchTeamsAssignForbidden with default headers values
func NewPostUsersSearchTeamsAssignForbidden() *PostUsersSearchTeamsAssignForbidden {
	return &PostUsersSearchTeamsAssignForbidden{}
}

/*
PostUsersSearchTeamsAssignForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostUsersSearchTeamsAssignForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign forbidden response has a 2xx status code
func (o *PostUsersSearchTeamsAssignForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign forbidden response has a 3xx status code
func (o *PostUsersSearchTeamsAssignForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign forbidden response has a 4xx status code
func (o *PostUsersSearchTeamsAssignForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign forbidden response has a 5xx status code
func (o *PostUsersSearchTeamsAssignForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign forbidden response a status code equal to that given
func (o *PostUsersSearchTeamsAssignForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUsersSearchTeamsAssignForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersSearchTeamsAssignForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersSearchTeamsAssignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignNotFound creates a PostUsersSearchTeamsAssignNotFound with default headers values
func NewPostUsersSearchTeamsAssignNotFound() *PostUsersSearchTeamsAssignNotFound {
	return &PostUsersSearchTeamsAssignNotFound{}
}

/*
PostUsersSearchTeamsAssignNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostUsersSearchTeamsAssignNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign not found response has a 2xx status code
func (o *PostUsersSearchTeamsAssignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign not found response has a 3xx status code
func (o *PostUsersSearchTeamsAssignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign not found response has a 4xx status code
func (o *PostUsersSearchTeamsAssignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign not found response has a 5xx status code
func (o *PostUsersSearchTeamsAssignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign not found response a status code equal to that given
func (o *PostUsersSearchTeamsAssignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUsersSearchTeamsAssignNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersSearchTeamsAssignNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersSearchTeamsAssignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignRequestTimeout creates a PostUsersSearchTeamsAssignRequestTimeout with default headers values
func NewPostUsersSearchTeamsAssignRequestTimeout() *PostUsersSearchTeamsAssignRequestTimeout {
	return &PostUsersSearchTeamsAssignRequestTimeout{}
}

/*
PostUsersSearchTeamsAssignRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUsersSearchTeamsAssignRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign request timeout response has a 2xx status code
func (o *PostUsersSearchTeamsAssignRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign request timeout response has a 3xx status code
func (o *PostUsersSearchTeamsAssignRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign request timeout response has a 4xx status code
func (o *PostUsersSearchTeamsAssignRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign request timeout response has a 5xx status code
func (o *PostUsersSearchTeamsAssignRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign request timeout response a status code equal to that given
func (o *PostUsersSearchTeamsAssignRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostUsersSearchTeamsAssignRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUsersSearchTeamsAssignRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUsersSearchTeamsAssignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignRequestEntityTooLarge creates a PostUsersSearchTeamsAssignRequestEntityTooLarge with default headers values
func NewPostUsersSearchTeamsAssignRequestEntityTooLarge() *PostUsersSearchTeamsAssignRequestEntityTooLarge {
	return &PostUsersSearchTeamsAssignRequestEntityTooLarge{}
}

/*
PostUsersSearchTeamsAssignRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostUsersSearchTeamsAssignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign request entity too large response has a 2xx status code
func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign request entity too large response has a 3xx status code
func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign request entity too large response has a 4xx status code
func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign request entity too large response has a 5xx status code
func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign request entity too large response a status code equal to that given
func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignUnsupportedMediaType creates a PostUsersSearchTeamsAssignUnsupportedMediaType with default headers values
func NewPostUsersSearchTeamsAssignUnsupportedMediaType() *PostUsersSearchTeamsAssignUnsupportedMediaType {
	return &PostUsersSearchTeamsAssignUnsupportedMediaType{}
}

/*
PostUsersSearchTeamsAssignUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUsersSearchTeamsAssignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign unsupported media type response has a 2xx status code
func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign unsupported media type response has a 3xx status code
func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign unsupported media type response has a 4xx status code
func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign unsupported media type response has a 5xx status code
func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign unsupported media type response a status code equal to that given
func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignTooManyRequests creates a PostUsersSearchTeamsAssignTooManyRequests with default headers values
func NewPostUsersSearchTeamsAssignTooManyRequests() *PostUsersSearchTeamsAssignTooManyRequests {
	return &PostUsersSearchTeamsAssignTooManyRequests{}
}

/*
PostUsersSearchTeamsAssignTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUsersSearchTeamsAssignTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign too many requests response has a 2xx status code
func (o *PostUsersSearchTeamsAssignTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign too many requests response has a 3xx status code
func (o *PostUsersSearchTeamsAssignTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign too many requests response has a 4xx status code
func (o *PostUsersSearchTeamsAssignTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users search teams assign too many requests response has a 5xx status code
func (o *PostUsersSearchTeamsAssignTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post users search teams assign too many requests response a status code equal to that given
func (o *PostUsersSearchTeamsAssignTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostUsersSearchTeamsAssignTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersSearchTeamsAssignTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersSearchTeamsAssignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignInternalServerError creates a PostUsersSearchTeamsAssignInternalServerError with default headers values
func NewPostUsersSearchTeamsAssignInternalServerError() *PostUsersSearchTeamsAssignInternalServerError {
	return &PostUsersSearchTeamsAssignInternalServerError{}
}

/*
PostUsersSearchTeamsAssignInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUsersSearchTeamsAssignInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign internal server error response has a 2xx status code
func (o *PostUsersSearchTeamsAssignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign internal server error response has a 3xx status code
func (o *PostUsersSearchTeamsAssignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign internal server error response has a 4xx status code
func (o *PostUsersSearchTeamsAssignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search teams assign internal server error response has a 5xx status code
func (o *PostUsersSearchTeamsAssignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post users search teams assign internal server error response a status code equal to that given
func (o *PostUsersSearchTeamsAssignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUsersSearchTeamsAssignInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersSearchTeamsAssignInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersSearchTeamsAssignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignServiceUnavailable creates a PostUsersSearchTeamsAssignServiceUnavailable with default headers values
func NewPostUsersSearchTeamsAssignServiceUnavailable() *PostUsersSearchTeamsAssignServiceUnavailable {
	return &PostUsersSearchTeamsAssignServiceUnavailable{}
}

/*
PostUsersSearchTeamsAssignServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUsersSearchTeamsAssignServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign service unavailable response has a 2xx status code
func (o *PostUsersSearchTeamsAssignServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign service unavailable response has a 3xx status code
func (o *PostUsersSearchTeamsAssignServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign service unavailable response has a 4xx status code
func (o *PostUsersSearchTeamsAssignServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search teams assign service unavailable response has a 5xx status code
func (o *PostUsersSearchTeamsAssignServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post users search teams assign service unavailable response a status code equal to that given
func (o *PostUsersSearchTeamsAssignServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUsersSearchTeamsAssignServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersSearchTeamsAssignServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersSearchTeamsAssignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersSearchTeamsAssignGatewayTimeout creates a PostUsersSearchTeamsAssignGatewayTimeout with default headers values
func NewPostUsersSearchTeamsAssignGatewayTimeout() *PostUsersSearchTeamsAssignGatewayTimeout {
	return &PostUsersSearchTeamsAssignGatewayTimeout{}
}

/*
PostUsersSearchTeamsAssignGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostUsersSearchTeamsAssignGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users search teams assign gateway timeout response has a 2xx status code
func (o *PostUsersSearchTeamsAssignGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users search teams assign gateway timeout response has a 3xx status code
func (o *PostUsersSearchTeamsAssignGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users search teams assign gateway timeout response has a 4xx status code
func (o *PostUsersSearchTeamsAssignGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users search teams assign gateway timeout response has a 5xx status code
func (o *PostUsersSearchTeamsAssignGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post users search teams assign gateway timeout response a status code equal to that given
func (o *PostUsersSearchTeamsAssignGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUsersSearchTeamsAssignGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUsersSearchTeamsAssignGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/search/teams/assign][%d] postUsersSearchTeamsAssignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUsersSearchTeamsAssignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersSearchTeamsAssignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}