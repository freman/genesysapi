// Code generated by go-swagger; DO NOT EDIT.

package license

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLicenseUsersReader is a Reader for the GetLicenseUsers structure.
type GetLicenseUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLicenseUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLicenseUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLicenseUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLicenseUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLicenseUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLicenseUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLicenseUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLicenseUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLicenseUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLicenseUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLicenseUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseUsersOK creates a GetLicenseUsersOK with default headers values
func NewGetLicenseUsersOK() *GetLicenseUsersOK {
	return &GetLicenseUsersOK{}
}

/*
GetLicenseUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLicenseUsersOK struct {
	Payload *models.UserLicensesEntityListing
}

// IsSuccess returns true when this get license users o k response has a 2xx status code
func (o *GetLicenseUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get license users o k response has a 3xx status code
func (o *GetLicenseUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users o k response has a 4xx status code
func (o *GetLicenseUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license users o k response has a 5xx status code
func (o *GetLicenseUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users o k response a status code equal to that given
func (o *GetLicenseUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLicenseUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersOK  %+v", 200, o.Payload)
}

func (o *GetLicenseUsersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersOK  %+v", 200, o.Payload)
}

func (o *GetLicenseUsersOK) GetPayload() *models.UserLicensesEntityListing {
	return o.Payload
}

func (o *GetLicenseUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserLicensesEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersBadRequest creates a GetLicenseUsersBadRequest with default headers values
func NewGetLicenseUsersBadRequest() *GetLicenseUsersBadRequest {
	return &GetLicenseUsersBadRequest{}
}

/*
GetLicenseUsersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLicenseUsersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users bad request response has a 2xx status code
func (o *GetLicenseUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users bad request response has a 3xx status code
func (o *GetLicenseUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users bad request response has a 4xx status code
func (o *GetLicenseUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users bad request response has a 5xx status code
func (o *GetLicenseUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users bad request response a status code equal to that given
func (o *GetLicenseUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLicenseUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetLicenseUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersUnauthorized creates a GetLicenseUsersUnauthorized with default headers values
func NewGetLicenseUsersUnauthorized() *GetLicenseUsersUnauthorized {
	return &GetLicenseUsersUnauthorized{}
}

/*
GetLicenseUsersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLicenseUsersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users unauthorized response has a 2xx status code
func (o *GetLicenseUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users unauthorized response has a 3xx status code
func (o *GetLicenseUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users unauthorized response has a 4xx status code
func (o *GetLicenseUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users unauthorized response has a 5xx status code
func (o *GetLicenseUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users unauthorized response a status code equal to that given
func (o *GetLicenseUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLicenseUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLicenseUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersForbidden creates a GetLicenseUsersForbidden with default headers values
func NewGetLicenseUsersForbidden() *GetLicenseUsersForbidden {
	return &GetLicenseUsersForbidden{}
}

/*
GetLicenseUsersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLicenseUsersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users forbidden response has a 2xx status code
func (o *GetLicenseUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users forbidden response has a 3xx status code
func (o *GetLicenseUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users forbidden response has a 4xx status code
func (o *GetLicenseUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users forbidden response has a 5xx status code
func (o *GetLicenseUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users forbidden response a status code equal to that given
func (o *GetLicenseUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLicenseUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseUsersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetLicenseUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersNotFound creates a GetLicenseUsersNotFound with default headers values
func NewGetLicenseUsersNotFound() *GetLicenseUsersNotFound {
	return &GetLicenseUsersNotFound{}
}

/*
GetLicenseUsersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLicenseUsersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users not found response has a 2xx status code
func (o *GetLicenseUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users not found response has a 3xx status code
func (o *GetLicenseUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users not found response has a 4xx status code
func (o *GetLicenseUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users not found response has a 5xx status code
func (o *GetLicenseUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users not found response a status code equal to that given
func (o *GetLicenseUsersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLicenseUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseUsersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersRequestTimeout creates a GetLicenseUsersRequestTimeout with default headers values
func NewGetLicenseUsersRequestTimeout() *GetLicenseUsersRequestTimeout {
	return &GetLicenseUsersRequestTimeout{}
}

/*
GetLicenseUsersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLicenseUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users request timeout response has a 2xx status code
func (o *GetLicenseUsersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users request timeout response has a 3xx status code
func (o *GetLicenseUsersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users request timeout response has a 4xx status code
func (o *GetLicenseUsersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users request timeout response has a 5xx status code
func (o *GetLicenseUsersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users request timeout response a status code equal to that given
func (o *GetLicenseUsersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLicenseUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLicenseUsersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLicenseUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersRequestEntityTooLarge creates a GetLicenseUsersRequestEntityTooLarge with default headers values
func NewGetLicenseUsersRequestEntityTooLarge() *GetLicenseUsersRequestEntityTooLarge {
	return &GetLicenseUsersRequestEntityTooLarge{}
}

/*
GetLicenseUsersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetLicenseUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users request entity too large response has a 2xx status code
func (o *GetLicenseUsersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users request entity too large response has a 3xx status code
func (o *GetLicenseUsersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users request entity too large response has a 4xx status code
func (o *GetLicenseUsersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users request entity too large response has a 5xx status code
func (o *GetLicenseUsersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users request entity too large response a status code equal to that given
func (o *GetLicenseUsersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLicenseUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseUsersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLicenseUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersUnsupportedMediaType creates a GetLicenseUsersUnsupportedMediaType with default headers values
func NewGetLicenseUsersUnsupportedMediaType() *GetLicenseUsersUnsupportedMediaType {
	return &GetLicenseUsersUnsupportedMediaType{}
}

/*
GetLicenseUsersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLicenseUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users unsupported media type response has a 2xx status code
func (o *GetLicenseUsersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users unsupported media type response has a 3xx status code
func (o *GetLicenseUsersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users unsupported media type response has a 4xx status code
func (o *GetLicenseUsersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users unsupported media type response has a 5xx status code
func (o *GetLicenseUsersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users unsupported media type response a status code equal to that given
func (o *GetLicenseUsersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLicenseUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseUsersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLicenseUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersTooManyRequests creates a GetLicenseUsersTooManyRequests with default headers values
func NewGetLicenseUsersTooManyRequests() *GetLicenseUsersTooManyRequests {
	return &GetLicenseUsersTooManyRequests{}
}

/*
GetLicenseUsersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLicenseUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users too many requests response has a 2xx status code
func (o *GetLicenseUsersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users too many requests response has a 3xx status code
func (o *GetLicenseUsersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users too many requests response has a 4xx status code
func (o *GetLicenseUsersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license users too many requests response has a 5xx status code
func (o *GetLicenseUsersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get license users too many requests response a status code equal to that given
func (o *GetLicenseUsersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLicenseUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseUsersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLicenseUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersInternalServerError creates a GetLicenseUsersInternalServerError with default headers values
func NewGetLicenseUsersInternalServerError() *GetLicenseUsersInternalServerError {
	return &GetLicenseUsersInternalServerError{}
}

/*
GetLicenseUsersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLicenseUsersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users internal server error response has a 2xx status code
func (o *GetLicenseUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users internal server error response has a 3xx status code
func (o *GetLicenseUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users internal server error response has a 4xx status code
func (o *GetLicenseUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license users internal server error response has a 5xx status code
func (o *GetLicenseUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get license users internal server error response a status code equal to that given
func (o *GetLicenseUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLicenseUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLicenseUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersServiceUnavailable creates a GetLicenseUsersServiceUnavailable with default headers values
func NewGetLicenseUsersServiceUnavailable() *GetLicenseUsersServiceUnavailable {
	return &GetLicenseUsersServiceUnavailable{}
}

/*
GetLicenseUsersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLicenseUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users service unavailable response has a 2xx status code
func (o *GetLicenseUsersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users service unavailable response has a 3xx status code
func (o *GetLicenseUsersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users service unavailable response has a 4xx status code
func (o *GetLicenseUsersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license users service unavailable response has a 5xx status code
func (o *GetLicenseUsersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get license users service unavailable response a status code equal to that given
func (o *GetLicenseUsersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLicenseUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseUsersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLicenseUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseUsersGatewayTimeout creates a GetLicenseUsersGatewayTimeout with default headers values
func NewGetLicenseUsersGatewayTimeout() *GetLicenseUsersGatewayTimeout {
	return &GetLicenseUsersGatewayTimeout{}
}

/*
GetLicenseUsersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLicenseUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get license users gateway timeout response has a 2xx status code
func (o *GetLicenseUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license users gateway timeout response has a 3xx status code
func (o *GetLicenseUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license users gateway timeout response has a 4xx status code
func (o *GetLicenseUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license users gateway timeout response has a 5xx status code
func (o *GetLicenseUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get license users gateway timeout response a status code equal to that given
func (o *GetLicenseUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLicenseUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseUsersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/license/users][%d] getLicenseUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLicenseUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLicenseUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
