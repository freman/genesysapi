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

// GetScimV2UserReader is a Reader for the GetScimV2User structure.
type GetScimV2UserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimV2UserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimV2UserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 304:
		result := NewGetScimV2UserNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewGetScimV2UserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimV2UserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimV2UserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimV2UserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScimV2UserRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetScimV2UserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimV2UserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimV2UserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimV2UserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimV2UserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimV2UserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimV2UserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimV2UserOK creates a GetScimV2UserOK with default headers values
func NewGetScimV2UserOK() *GetScimV2UserOK {
	return &GetScimV2UserOK{}
}

/*
GetScimV2UserOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScimV2UserOK struct {
	Payload *models.ScimV2User
}

// IsSuccess returns true when this get scim v2 user o k response has a 2xx status code
func (o *GetScimV2UserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scim v2 user o k response has a 3xx status code
func (o *GetScimV2UserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user o k response has a 4xx status code
func (o *GetScimV2UserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 user o k response has a 5xx status code
func (o *GetScimV2UserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user o k response a status code equal to that given
func (o *GetScimV2UserOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScimV2UserOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserOK  %+v", 200, o.Payload)
}

func (o *GetScimV2UserOK) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserOK  %+v", 200, o.Payload)
}

func (o *GetScimV2UserOK) GetPayload() *models.ScimV2User {
	return o.Payload
}

func (o *GetScimV2UserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserNotModified creates a GetScimV2UserNotModified with default headers values
func NewGetScimV2UserNotModified() *GetScimV2UserNotModified {
	return &GetScimV2UserNotModified{}
}

/*
GetScimV2UserNotModified describes a response with status code 304, with default header values.

If-Non-Match header matches current version. No content returned.
*/
type GetScimV2UserNotModified struct {
}

// IsSuccess returns true when this get scim v2 user not modified response has a 2xx status code
func (o *GetScimV2UserNotModified) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user not modified response has a 3xx status code
func (o *GetScimV2UserNotModified) IsRedirect() bool {
	return true
}

// IsClientError returns true when this get scim v2 user not modified response has a 4xx status code
func (o *GetScimV2UserNotModified) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 user not modified response has a 5xx status code
func (o *GetScimV2UserNotModified) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user not modified response a status code equal to that given
func (o *GetScimV2UserNotModified) IsCode(code int) bool {
	return code == 304
}

func (o *GetScimV2UserNotModified) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserNotModified ", 304)
}

func (o *GetScimV2UserNotModified) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserNotModified ", 304)
}

func (o *GetScimV2UserNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetScimV2UserBadRequest creates a GetScimV2UserBadRequest with default headers values
func NewGetScimV2UserBadRequest() *GetScimV2UserBadRequest {
	return &GetScimV2UserBadRequest{}
}

/*
GetScimV2UserBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimV2UserBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user bad request response has a 2xx status code
func (o *GetScimV2UserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user bad request response has a 3xx status code
func (o *GetScimV2UserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user bad request response has a 4xx status code
func (o *GetScimV2UserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user bad request response has a 5xx status code
func (o *GetScimV2UserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user bad request response a status code equal to that given
func (o *GetScimV2UserBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScimV2UserBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2UserBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2UserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserUnauthorized creates a GetScimV2UserUnauthorized with default headers values
func NewGetScimV2UserUnauthorized() *GetScimV2UserUnauthorized {
	return &GetScimV2UserUnauthorized{}
}

/*
GetScimV2UserUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimV2UserUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user unauthorized response has a 2xx status code
func (o *GetScimV2UserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user unauthorized response has a 3xx status code
func (o *GetScimV2UserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user unauthorized response has a 4xx status code
func (o *GetScimV2UserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user unauthorized response has a 5xx status code
func (o *GetScimV2UserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user unauthorized response a status code equal to that given
func (o *GetScimV2UserUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScimV2UserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2UserUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2UserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserForbidden creates a GetScimV2UserForbidden with default headers values
func NewGetScimV2UserForbidden() *GetScimV2UserForbidden {
	return &GetScimV2UserForbidden{}
}

/*
GetScimV2UserForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetScimV2UserForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user forbidden response has a 2xx status code
func (o *GetScimV2UserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user forbidden response has a 3xx status code
func (o *GetScimV2UserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user forbidden response has a 4xx status code
func (o *GetScimV2UserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user forbidden response has a 5xx status code
func (o *GetScimV2UserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user forbidden response a status code equal to that given
func (o *GetScimV2UserForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScimV2UserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2UserForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2UserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserNotFound creates a GetScimV2UserNotFound with default headers values
func NewGetScimV2UserNotFound() *GetScimV2UserNotFound {
	return &GetScimV2UserNotFound{}
}

/*
GetScimV2UserNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetScimV2UserNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user not found response has a 2xx status code
func (o *GetScimV2UserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user not found response has a 3xx status code
func (o *GetScimV2UserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user not found response has a 4xx status code
func (o *GetScimV2UserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user not found response has a 5xx status code
func (o *GetScimV2UserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user not found response a status code equal to that given
func (o *GetScimV2UserNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScimV2UserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2UserNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2UserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserRequestTimeout creates a GetScimV2UserRequestTimeout with default headers values
func NewGetScimV2UserRequestTimeout() *GetScimV2UserRequestTimeout {
	return &GetScimV2UserRequestTimeout{}
}

/*
GetScimV2UserRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScimV2UserRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user request timeout response has a 2xx status code
func (o *GetScimV2UserRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user request timeout response has a 3xx status code
func (o *GetScimV2UserRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user request timeout response has a 4xx status code
func (o *GetScimV2UserRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user request timeout response has a 5xx status code
func (o *GetScimV2UserRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user request timeout response a status code equal to that given
func (o *GetScimV2UserRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetScimV2UserRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimV2UserRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimV2UserRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserConflict creates a GetScimV2UserConflict with default headers values
func NewGetScimV2UserConflict() *GetScimV2UserConflict {
	return &GetScimV2UserConflict{}
}

/*
GetScimV2UserConflict describes a response with status code 409, with default header values.

Conflict
*/
type GetScimV2UserConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user conflict response has a 2xx status code
func (o *GetScimV2UserConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user conflict response has a 3xx status code
func (o *GetScimV2UserConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user conflict response has a 4xx status code
func (o *GetScimV2UserConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user conflict response has a 5xx status code
func (o *GetScimV2UserConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user conflict response a status code equal to that given
func (o *GetScimV2UserConflict) IsCode(code int) bool {
	return code == 409
}

func (o *GetScimV2UserConflict) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserConflict  %+v", 409, o.Payload)
}

func (o *GetScimV2UserConflict) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserConflict  %+v", 409, o.Payload)
}

func (o *GetScimV2UserConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserRequestEntityTooLarge creates a GetScimV2UserRequestEntityTooLarge with default headers values
func NewGetScimV2UserRequestEntityTooLarge() *GetScimV2UserRequestEntityTooLarge {
	return &GetScimV2UserRequestEntityTooLarge{}
}

/*
GetScimV2UserRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetScimV2UserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user request entity too large response has a 2xx status code
func (o *GetScimV2UserRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user request entity too large response has a 3xx status code
func (o *GetScimV2UserRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user request entity too large response has a 4xx status code
func (o *GetScimV2UserRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user request entity too large response has a 5xx status code
func (o *GetScimV2UserRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user request entity too large response a status code equal to that given
func (o *GetScimV2UserRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetScimV2UserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2UserRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2UserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserUnsupportedMediaType creates a GetScimV2UserUnsupportedMediaType with default headers values
func NewGetScimV2UserUnsupportedMediaType() *GetScimV2UserUnsupportedMediaType {
	return &GetScimV2UserUnsupportedMediaType{}
}

/*
GetScimV2UserUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimV2UserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user unsupported media type response has a 2xx status code
func (o *GetScimV2UserUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user unsupported media type response has a 3xx status code
func (o *GetScimV2UserUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user unsupported media type response has a 4xx status code
func (o *GetScimV2UserUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user unsupported media type response has a 5xx status code
func (o *GetScimV2UserUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user unsupported media type response a status code equal to that given
func (o *GetScimV2UserUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetScimV2UserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2UserUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2UserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserTooManyRequests creates a GetScimV2UserTooManyRequests with default headers values
func NewGetScimV2UserTooManyRequests() *GetScimV2UserTooManyRequests {
	return &GetScimV2UserTooManyRequests{}
}

/*
GetScimV2UserTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScimV2UserTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user too many requests response has a 2xx status code
func (o *GetScimV2UserTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user too many requests response has a 3xx status code
func (o *GetScimV2UserTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user too many requests response has a 4xx status code
func (o *GetScimV2UserTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 user too many requests response has a 5xx status code
func (o *GetScimV2UserTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 user too many requests response a status code equal to that given
func (o *GetScimV2UserTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetScimV2UserTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2UserTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2UserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserInternalServerError creates a GetScimV2UserInternalServerError with default headers values
func NewGetScimV2UserInternalServerError() *GetScimV2UserInternalServerError {
	return &GetScimV2UserInternalServerError{}
}

/*
GetScimV2UserInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimV2UserInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user internal server error response has a 2xx status code
func (o *GetScimV2UserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user internal server error response has a 3xx status code
func (o *GetScimV2UserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user internal server error response has a 4xx status code
func (o *GetScimV2UserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 user internal server error response has a 5xx status code
func (o *GetScimV2UserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 user internal server error response a status code equal to that given
func (o *GetScimV2UserInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScimV2UserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2UserInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2UserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserServiceUnavailable creates a GetScimV2UserServiceUnavailable with default headers values
func NewGetScimV2UserServiceUnavailable() *GetScimV2UserServiceUnavailable {
	return &GetScimV2UserServiceUnavailable{}
}

/*
GetScimV2UserServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimV2UserServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user service unavailable response has a 2xx status code
func (o *GetScimV2UserServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user service unavailable response has a 3xx status code
func (o *GetScimV2UserServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user service unavailable response has a 4xx status code
func (o *GetScimV2UserServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 user service unavailable response has a 5xx status code
func (o *GetScimV2UserServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 user service unavailable response a status code equal to that given
func (o *GetScimV2UserServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetScimV2UserServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2UserServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2UserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2UserGatewayTimeout creates a GetScimV2UserGatewayTimeout with default headers values
func NewGetScimV2UserGatewayTimeout() *GetScimV2UserGatewayTimeout {
	return &GetScimV2UserGatewayTimeout{}
}

/*
GetScimV2UserGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetScimV2UserGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 user gateway timeout response has a 2xx status code
func (o *GetScimV2UserGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 user gateway timeout response has a 3xx status code
func (o *GetScimV2UserGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 user gateway timeout response has a 4xx status code
func (o *GetScimV2UserGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 user gateway timeout response has a 5xx status code
func (o *GetScimV2UserGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 user gateway timeout response a status code equal to that given
func (o *GetScimV2UserGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetScimV2UserGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2UserGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/users/{userId}][%d] getScimV2UserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2UserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2UserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
