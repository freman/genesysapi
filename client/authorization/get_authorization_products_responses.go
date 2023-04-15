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

// GetAuthorizationProductsReader is a Reader for the GetAuthorizationProducts structure.
type GetAuthorizationProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationProductsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationProductsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationProductsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationProductsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationProductsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationProductsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationProductsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationProductsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationProductsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationProductsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationProductsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationProductsOK creates a GetAuthorizationProductsOK with default headers values
func NewGetAuthorizationProductsOK() *GetAuthorizationProductsOK {
	return &GetAuthorizationProductsOK{}
}

/*
GetAuthorizationProductsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuthorizationProductsOK struct {
	Payload *models.OrganizationProductEntityListing
}

// IsSuccess returns true when this get authorization products o k response has a 2xx status code
func (o *GetAuthorizationProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get authorization products o k response has a 3xx status code
func (o *GetAuthorizationProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products o k response has a 4xx status code
func (o *GetAuthorizationProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization products o k response has a 5xx status code
func (o *GetAuthorizationProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products o k response a status code equal to that given
func (o *GetAuthorizationProductsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAuthorizationProductsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationProductsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationProductsOK) GetPayload() *models.OrganizationProductEntityListing {
	return o.Payload
}

func (o *GetAuthorizationProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationProductEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsBadRequest creates a GetAuthorizationProductsBadRequest with default headers values
func NewGetAuthorizationProductsBadRequest() *GetAuthorizationProductsBadRequest {
	return &GetAuthorizationProductsBadRequest{}
}

/*
GetAuthorizationProductsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationProductsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products bad request response has a 2xx status code
func (o *GetAuthorizationProductsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products bad request response has a 3xx status code
func (o *GetAuthorizationProductsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products bad request response has a 4xx status code
func (o *GetAuthorizationProductsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products bad request response has a 5xx status code
func (o *GetAuthorizationProductsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products bad request response a status code equal to that given
func (o *GetAuthorizationProductsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAuthorizationProductsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationProductsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationProductsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsUnauthorized creates a GetAuthorizationProductsUnauthorized with default headers values
func NewGetAuthorizationProductsUnauthorized() *GetAuthorizationProductsUnauthorized {
	return &GetAuthorizationProductsUnauthorized{}
}

/*
GetAuthorizationProductsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationProductsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products unauthorized response has a 2xx status code
func (o *GetAuthorizationProductsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products unauthorized response has a 3xx status code
func (o *GetAuthorizationProductsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products unauthorized response has a 4xx status code
func (o *GetAuthorizationProductsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products unauthorized response has a 5xx status code
func (o *GetAuthorizationProductsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products unauthorized response a status code equal to that given
func (o *GetAuthorizationProductsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAuthorizationProductsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationProductsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationProductsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsForbidden creates a GetAuthorizationProductsForbidden with default headers values
func NewGetAuthorizationProductsForbidden() *GetAuthorizationProductsForbidden {
	return &GetAuthorizationProductsForbidden{}
}

/*
GetAuthorizationProductsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationProductsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products forbidden response has a 2xx status code
func (o *GetAuthorizationProductsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products forbidden response has a 3xx status code
func (o *GetAuthorizationProductsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products forbidden response has a 4xx status code
func (o *GetAuthorizationProductsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products forbidden response has a 5xx status code
func (o *GetAuthorizationProductsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products forbidden response a status code equal to that given
func (o *GetAuthorizationProductsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAuthorizationProductsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationProductsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationProductsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsNotFound creates a GetAuthorizationProductsNotFound with default headers values
func NewGetAuthorizationProductsNotFound() *GetAuthorizationProductsNotFound {
	return &GetAuthorizationProductsNotFound{}
}

/*
GetAuthorizationProductsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAuthorizationProductsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products not found response has a 2xx status code
func (o *GetAuthorizationProductsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products not found response has a 3xx status code
func (o *GetAuthorizationProductsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products not found response has a 4xx status code
func (o *GetAuthorizationProductsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products not found response has a 5xx status code
func (o *GetAuthorizationProductsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products not found response a status code equal to that given
func (o *GetAuthorizationProductsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAuthorizationProductsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationProductsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationProductsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsRequestTimeout creates a GetAuthorizationProductsRequestTimeout with default headers values
func NewGetAuthorizationProductsRequestTimeout() *GetAuthorizationProductsRequestTimeout {
	return &GetAuthorizationProductsRequestTimeout{}
}

/*
GetAuthorizationProductsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationProductsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products request timeout response has a 2xx status code
func (o *GetAuthorizationProductsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products request timeout response has a 3xx status code
func (o *GetAuthorizationProductsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products request timeout response has a 4xx status code
func (o *GetAuthorizationProductsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products request timeout response has a 5xx status code
func (o *GetAuthorizationProductsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products request timeout response a status code equal to that given
func (o *GetAuthorizationProductsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAuthorizationProductsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationProductsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationProductsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsRequestEntityTooLarge creates a GetAuthorizationProductsRequestEntityTooLarge with default headers values
func NewGetAuthorizationProductsRequestEntityTooLarge() *GetAuthorizationProductsRequestEntityTooLarge {
	return &GetAuthorizationProductsRequestEntityTooLarge{}
}

/*
GetAuthorizationProductsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAuthorizationProductsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products request entity too large response has a 2xx status code
func (o *GetAuthorizationProductsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products request entity too large response has a 3xx status code
func (o *GetAuthorizationProductsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products request entity too large response has a 4xx status code
func (o *GetAuthorizationProductsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products request entity too large response has a 5xx status code
func (o *GetAuthorizationProductsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products request entity too large response a status code equal to that given
func (o *GetAuthorizationProductsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsUnsupportedMediaType creates a GetAuthorizationProductsUnsupportedMediaType with default headers values
func NewGetAuthorizationProductsUnsupportedMediaType() *GetAuthorizationProductsUnsupportedMediaType {
	return &GetAuthorizationProductsUnsupportedMediaType{}
}

/*
GetAuthorizationProductsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationProductsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products unsupported media type response has a 2xx status code
func (o *GetAuthorizationProductsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products unsupported media type response has a 3xx status code
func (o *GetAuthorizationProductsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products unsupported media type response has a 4xx status code
func (o *GetAuthorizationProductsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products unsupported media type response has a 5xx status code
func (o *GetAuthorizationProductsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products unsupported media type response a status code equal to that given
func (o *GetAuthorizationProductsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAuthorizationProductsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationProductsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationProductsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsTooManyRequests creates a GetAuthorizationProductsTooManyRequests with default headers values
func NewGetAuthorizationProductsTooManyRequests() *GetAuthorizationProductsTooManyRequests {
	return &GetAuthorizationProductsTooManyRequests{}
}

/*
GetAuthorizationProductsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationProductsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products too many requests response has a 2xx status code
func (o *GetAuthorizationProductsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products too many requests response has a 3xx status code
func (o *GetAuthorizationProductsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products too many requests response has a 4xx status code
func (o *GetAuthorizationProductsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization products too many requests response has a 5xx status code
func (o *GetAuthorizationProductsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization products too many requests response a status code equal to that given
func (o *GetAuthorizationProductsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAuthorizationProductsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationProductsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationProductsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsInternalServerError creates a GetAuthorizationProductsInternalServerError with default headers values
func NewGetAuthorizationProductsInternalServerError() *GetAuthorizationProductsInternalServerError {
	return &GetAuthorizationProductsInternalServerError{}
}

/*
GetAuthorizationProductsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationProductsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products internal server error response has a 2xx status code
func (o *GetAuthorizationProductsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products internal server error response has a 3xx status code
func (o *GetAuthorizationProductsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products internal server error response has a 4xx status code
func (o *GetAuthorizationProductsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization products internal server error response has a 5xx status code
func (o *GetAuthorizationProductsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization products internal server error response a status code equal to that given
func (o *GetAuthorizationProductsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAuthorizationProductsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationProductsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationProductsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsServiceUnavailable creates a GetAuthorizationProductsServiceUnavailable with default headers values
func NewGetAuthorizationProductsServiceUnavailable() *GetAuthorizationProductsServiceUnavailable {
	return &GetAuthorizationProductsServiceUnavailable{}
}

/*
GetAuthorizationProductsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationProductsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products service unavailable response has a 2xx status code
func (o *GetAuthorizationProductsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products service unavailable response has a 3xx status code
func (o *GetAuthorizationProductsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products service unavailable response has a 4xx status code
func (o *GetAuthorizationProductsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization products service unavailable response has a 5xx status code
func (o *GetAuthorizationProductsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization products service unavailable response a status code equal to that given
func (o *GetAuthorizationProductsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAuthorizationProductsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationProductsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationProductsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsGatewayTimeout creates a GetAuthorizationProductsGatewayTimeout with default headers values
func NewGetAuthorizationProductsGatewayTimeout() *GetAuthorizationProductsGatewayTimeout {
	return &GetAuthorizationProductsGatewayTimeout{}
}

/*
GetAuthorizationProductsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAuthorizationProductsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization products gateway timeout response has a 2xx status code
func (o *GetAuthorizationProductsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization products gateway timeout response has a 3xx status code
func (o *GetAuthorizationProductsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization products gateway timeout response has a 4xx status code
func (o *GetAuthorizationProductsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization products gateway timeout response has a 5xx status code
func (o *GetAuthorizationProductsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization products gateway timeout response a status code equal to that given
func (o *GetAuthorizationProductsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAuthorizationProductsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationProductsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationProductsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
