// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetResponsemanagementResponsesReader is a Reader for the GetResponsemanagementResponses structure.
type GetResponsemanagementResponsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponsemanagementResponsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponsemanagementResponsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResponsemanagementResponsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResponsemanagementResponsesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResponsemanagementResponsesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResponsemanagementResponsesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetResponsemanagementResponsesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetResponsemanagementResponsesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetResponsemanagementResponsesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetResponsemanagementResponsesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResponsemanagementResponsesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetResponsemanagementResponsesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResponsemanagementResponsesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResponsemanagementResponsesOK creates a GetResponsemanagementResponsesOK with default headers values
func NewGetResponsemanagementResponsesOK() *GetResponsemanagementResponsesOK {
	return &GetResponsemanagementResponsesOK{}
}

/*
GetResponsemanagementResponsesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetResponsemanagementResponsesOK struct {
	Payload *models.ResponseEntityListing
}

// IsSuccess returns true when this get responsemanagement responses o k response has a 2xx status code
func (o *GetResponsemanagementResponsesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get responsemanagement responses o k response has a 3xx status code
func (o *GetResponsemanagementResponsesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses o k response has a 4xx status code
func (o *GetResponsemanagementResponsesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responses o k response has a 5xx status code
func (o *GetResponsemanagementResponsesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses o k response a status code equal to that given
func (o *GetResponsemanagementResponsesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResponsemanagementResponsesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponsesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponsesOK) GetPayload() *models.ResponseEntityListing {
	return o.Payload
}

func (o *GetResponsemanagementResponsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesBadRequest creates a GetResponsemanagementResponsesBadRequest with default headers values
func NewGetResponsemanagementResponsesBadRequest() *GetResponsemanagementResponsesBadRequest {
	return &GetResponsemanagementResponsesBadRequest{}
}

/*
GetResponsemanagementResponsesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetResponsemanagementResponsesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses bad request response has a 2xx status code
func (o *GetResponsemanagementResponsesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses bad request response has a 3xx status code
func (o *GetResponsemanagementResponsesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses bad request response has a 4xx status code
func (o *GetResponsemanagementResponsesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses bad request response has a 5xx status code
func (o *GetResponsemanagementResponsesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses bad request response a status code equal to that given
func (o *GetResponsemanagementResponsesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetResponsemanagementResponsesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponsesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponsesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesUnauthorized creates a GetResponsemanagementResponsesUnauthorized with default headers values
func NewGetResponsemanagementResponsesUnauthorized() *GetResponsemanagementResponsesUnauthorized {
	return &GetResponsemanagementResponsesUnauthorized{}
}

/*
GetResponsemanagementResponsesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetResponsemanagementResponsesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses unauthorized response has a 2xx status code
func (o *GetResponsemanagementResponsesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses unauthorized response has a 3xx status code
func (o *GetResponsemanagementResponsesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses unauthorized response has a 4xx status code
func (o *GetResponsemanagementResponsesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses unauthorized response has a 5xx status code
func (o *GetResponsemanagementResponsesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses unauthorized response a status code equal to that given
func (o *GetResponsemanagementResponsesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResponsemanagementResponsesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponsesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponsesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesForbidden creates a GetResponsemanagementResponsesForbidden with default headers values
func NewGetResponsemanagementResponsesForbidden() *GetResponsemanagementResponsesForbidden {
	return &GetResponsemanagementResponsesForbidden{}
}

/*
GetResponsemanagementResponsesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetResponsemanagementResponsesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses forbidden response has a 2xx status code
func (o *GetResponsemanagementResponsesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses forbidden response has a 3xx status code
func (o *GetResponsemanagementResponsesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses forbidden response has a 4xx status code
func (o *GetResponsemanagementResponsesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses forbidden response has a 5xx status code
func (o *GetResponsemanagementResponsesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses forbidden response a status code equal to that given
func (o *GetResponsemanagementResponsesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResponsemanagementResponsesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponsesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponsesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesNotFound creates a GetResponsemanagementResponsesNotFound with default headers values
func NewGetResponsemanagementResponsesNotFound() *GetResponsemanagementResponsesNotFound {
	return &GetResponsemanagementResponsesNotFound{}
}

/*
GetResponsemanagementResponsesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetResponsemanagementResponsesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses not found response has a 2xx status code
func (o *GetResponsemanagementResponsesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses not found response has a 3xx status code
func (o *GetResponsemanagementResponsesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses not found response has a 4xx status code
func (o *GetResponsemanagementResponsesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses not found response has a 5xx status code
func (o *GetResponsemanagementResponsesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses not found response a status code equal to that given
func (o *GetResponsemanagementResponsesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResponsemanagementResponsesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponsesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponsesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesRequestTimeout creates a GetResponsemanagementResponsesRequestTimeout with default headers values
func NewGetResponsemanagementResponsesRequestTimeout() *GetResponsemanagementResponsesRequestTimeout {
	return &GetResponsemanagementResponsesRequestTimeout{}
}

/*
GetResponsemanagementResponsesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetResponsemanagementResponsesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses request timeout response has a 2xx status code
func (o *GetResponsemanagementResponsesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses request timeout response has a 3xx status code
func (o *GetResponsemanagementResponsesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses request timeout response has a 4xx status code
func (o *GetResponsemanagementResponsesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses request timeout response has a 5xx status code
func (o *GetResponsemanagementResponsesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses request timeout response a status code equal to that given
func (o *GetResponsemanagementResponsesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetResponsemanagementResponsesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponsesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponsesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesRequestEntityTooLarge creates a GetResponsemanagementResponsesRequestEntityTooLarge with default headers values
func NewGetResponsemanagementResponsesRequestEntityTooLarge() *GetResponsemanagementResponsesRequestEntityTooLarge {
	return &GetResponsemanagementResponsesRequestEntityTooLarge{}
}

/*
GetResponsemanagementResponsesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetResponsemanagementResponsesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses request entity too large response has a 2xx status code
func (o *GetResponsemanagementResponsesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses request entity too large response has a 3xx status code
func (o *GetResponsemanagementResponsesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses request entity too large response has a 4xx status code
func (o *GetResponsemanagementResponsesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses request entity too large response has a 5xx status code
func (o *GetResponsemanagementResponsesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses request entity too large response a status code equal to that given
func (o *GetResponsemanagementResponsesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetResponsemanagementResponsesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponsesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponsesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesUnsupportedMediaType creates a GetResponsemanagementResponsesUnsupportedMediaType with default headers values
func NewGetResponsemanagementResponsesUnsupportedMediaType() *GetResponsemanagementResponsesUnsupportedMediaType {
	return &GetResponsemanagementResponsesUnsupportedMediaType{}
}

/*
GetResponsemanagementResponsesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetResponsemanagementResponsesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses unsupported media type response has a 2xx status code
func (o *GetResponsemanagementResponsesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses unsupported media type response has a 3xx status code
func (o *GetResponsemanagementResponsesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses unsupported media type response has a 4xx status code
func (o *GetResponsemanagementResponsesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses unsupported media type response has a 5xx status code
func (o *GetResponsemanagementResponsesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses unsupported media type response a status code equal to that given
func (o *GetResponsemanagementResponsesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetResponsemanagementResponsesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponsesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponsesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesTooManyRequests creates a GetResponsemanagementResponsesTooManyRequests with default headers values
func NewGetResponsemanagementResponsesTooManyRequests() *GetResponsemanagementResponsesTooManyRequests {
	return &GetResponsemanagementResponsesTooManyRequests{}
}

/*
GetResponsemanagementResponsesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetResponsemanagementResponsesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses too many requests response has a 2xx status code
func (o *GetResponsemanagementResponsesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses too many requests response has a 3xx status code
func (o *GetResponsemanagementResponsesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses too many requests response has a 4xx status code
func (o *GetResponsemanagementResponsesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement responses too many requests response has a 5xx status code
func (o *GetResponsemanagementResponsesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement responses too many requests response a status code equal to that given
func (o *GetResponsemanagementResponsesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetResponsemanagementResponsesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponsesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponsesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesInternalServerError creates a GetResponsemanagementResponsesInternalServerError with default headers values
func NewGetResponsemanagementResponsesInternalServerError() *GetResponsemanagementResponsesInternalServerError {
	return &GetResponsemanagementResponsesInternalServerError{}
}

/*
GetResponsemanagementResponsesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetResponsemanagementResponsesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses internal server error response has a 2xx status code
func (o *GetResponsemanagementResponsesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses internal server error response has a 3xx status code
func (o *GetResponsemanagementResponsesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses internal server error response has a 4xx status code
func (o *GetResponsemanagementResponsesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responses internal server error response has a 5xx status code
func (o *GetResponsemanagementResponsesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement responses internal server error response a status code equal to that given
func (o *GetResponsemanagementResponsesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetResponsemanagementResponsesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponsesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponsesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesServiceUnavailable creates a GetResponsemanagementResponsesServiceUnavailable with default headers values
func NewGetResponsemanagementResponsesServiceUnavailable() *GetResponsemanagementResponsesServiceUnavailable {
	return &GetResponsemanagementResponsesServiceUnavailable{}
}

/*
GetResponsemanagementResponsesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetResponsemanagementResponsesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses service unavailable response has a 2xx status code
func (o *GetResponsemanagementResponsesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses service unavailable response has a 3xx status code
func (o *GetResponsemanagementResponsesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses service unavailable response has a 4xx status code
func (o *GetResponsemanagementResponsesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responses service unavailable response has a 5xx status code
func (o *GetResponsemanagementResponsesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement responses service unavailable response a status code equal to that given
func (o *GetResponsemanagementResponsesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetResponsemanagementResponsesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponsesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponsesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponsesGatewayTimeout creates a GetResponsemanagementResponsesGatewayTimeout with default headers values
func NewGetResponsemanagementResponsesGatewayTimeout() *GetResponsemanagementResponsesGatewayTimeout {
	return &GetResponsemanagementResponsesGatewayTimeout{}
}

/*
GetResponsemanagementResponsesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetResponsemanagementResponsesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement responses gateway timeout response has a 2xx status code
func (o *GetResponsemanagementResponsesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement responses gateway timeout response has a 3xx status code
func (o *GetResponsemanagementResponsesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement responses gateway timeout response has a 4xx status code
func (o *GetResponsemanagementResponsesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement responses gateway timeout response has a 5xx status code
func (o *GetResponsemanagementResponsesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement responses gateway timeout response a status code equal to that given
func (o *GetResponsemanagementResponsesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetResponsemanagementResponsesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponsesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses][%d] getResponsemanagementResponsesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponsesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponsesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
