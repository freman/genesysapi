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

// GetAuthorizationRoleComparedefaultRightRoleIDReader is a Reader for the GetAuthorizationRoleComparedefaultRightRoleID structure.
type GetAuthorizationRoleComparedefaultRightRoleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationRoleComparedefaultRightRoleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDOK creates a GetAuthorizationRoleComparedefaultRightRoleIDOK with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDOK() *GetAuthorizationRoleComparedefaultRightRoleIDOK {
	return &GetAuthorizationRoleComparedefaultRightRoleIDOK{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuthorizationRoleComparedefaultRightRoleIDOK struct {
	Payload *models.DomainOrgRoleDifference
}

// IsSuccess returns true when this get authorization role comparedefault right role Id o k response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get authorization role comparedefault right role Id o k response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id o k response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization role comparedefault right role Id o k response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id o k response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) GetPayload() *models.DomainOrgRoleDifference {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainOrgRoleDifference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDBadRequest creates a GetAuthorizationRoleComparedefaultRightRoleIDBadRequest with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDBadRequest() *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest {
	return &GetAuthorizationRoleComparedefaultRightRoleIDBadRequest{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id bad request response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id bad request response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id bad request response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id bad request response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id bad request response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDUnauthorized creates a GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDUnauthorized() *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized {
	return &GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id unauthorized response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id unauthorized response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id unauthorized response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id unauthorized response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id unauthorized response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDForbidden creates a GetAuthorizationRoleComparedefaultRightRoleIDForbidden with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDForbidden() *GetAuthorizationRoleComparedefaultRightRoleIDForbidden {
	return &GetAuthorizationRoleComparedefaultRightRoleIDForbidden{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id forbidden response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id forbidden response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id forbidden response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id forbidden response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id forbidden response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDNotFound creates a GetAuthorizationRoleComparedefaultRightRoleIDNotFound with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDNotFound() *GetAuthorizationRoleComparedefaultRightRoleIDNotFound {
	return &GetAuthorizationRoleComparedefaultRightRoleIDNotFound{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id not found response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id not found response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id not found response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id not found response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id not found response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout creates a GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout() *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout {
	return &GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id request timeout response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id request timeout response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id request timeout response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id request timeout response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id request timeout response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge creates a GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge() *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge {
	return &GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id request entity too large response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id request entity too large response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id request entity too large response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id request entity too large response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id request entity too large response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType creates a GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType() *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType {
	return &GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id unsupported media type response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id unsupported media type response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id unsupported media type response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id unsupported media type response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id unsupported media type response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests creates a GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests() *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests {
	return &GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id too many requests response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id too many requests response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id too many requests response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization role comparedefault right role Id too many requests response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization role comparedefault right role Id too many requests response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDInternalServerError creates a GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDInternalServerError() *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError {
	return &GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id internal server error response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id internal server error response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id internal server error response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization role comparedefault right role Id internal server error response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization role comparedefault right role Id internal server error response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable creates a GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable() *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable {
	return &GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id service unavailable response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id service unavailable response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id service unavailable response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization role comparedefault right role Id service unavailable response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization role comparedefault right role Id service unavailable response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout creates a GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout with default headers values
func NewGetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout() *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout {
	return &GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout{}
}

/*
GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization role comparedefault right role Id gateway timeout response has a 2xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization role comparedefault right role Id gateway timeout response has a 3xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization role comparedefault right role Id gateway timeout response has a 4xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization role comparedefault right role Id gateway timeout response has a 5xx status code
func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization role comparedefault right role Id gateway timeout response a status code equal to that given
func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{leftRoleId}/comparedefault/{rightRoleId}][%d] getAuthorizationRoleComparedefaultRightRoleIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleComparedefaultRightRoleIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
