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

// GetResponsemanagementResponseReader is a Reader for the GetResponsemanagementResponse structure.
type GetResponsemanagementResponseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResponsemanagementResponseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResponsemanagementResponseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetResponsemanagementResponseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetResponsemanagementResponseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetResponsemanagementResponseForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResponsemanagementResponseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetResponsemanagementResponseRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetResponsemanagementResponseRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetResponsemanagementResponseUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetResponsemanagementResponseTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetResponsemanagementResponseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetResponsemanagementResponseServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetResponsemanagementResponseGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResponsemanagementResponseOK creates a GetResponsemanagementResponseOK with default headers values
func NewGetResponsemanagementResponseOK() *GetResponsemanagementResponseOK {
	return &GetResponsemanagementResponseOK{}
}

/*
GetResponsemanagementResponseOK describes a response with status code 200, with default header values.

successful operation
*/
type GetResponsemanagementResponseOK struct {
	Payload *models.Response
}

// IsSuccess returns true when this get responsemanagement response o k response has a 2xx status code
func (o *GetResponsemanagementResponseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get responsemanagement response o k response has a 3xx status code
func (o *GetResponsemanagementResponseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response o k response has a 4xx status code
func (o *GetResponsemanagementResponseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement response o k response has a 5xx status code
func (o *GetResponsemanagementResponseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response o k response a status code equal to that given
func (o *GetResponsemanagementResponseOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetResponsemanagementResponseOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponseOK) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseOK  %+v", 200, o.Payload)
}

func (o *GetResponsemanagementResponseOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *GetResponsemanagementResponseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseBadRequest creates a GetResponsemanagementResponseBadRequest with default headers values
func NewGetResponsemanagementResponseBadRequest() *GetResponsemanagementResponseBadRequest {
	return &GetResponsemanagementResponseBadRequest{}
}

/*
GetResponsemanagementResponseBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetResponsemanagementResponseBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response bad request response has a 2xx status code
func (o *GetResponsemanagementResponseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response bad request response has a 3xx status code
func (o *GetResponsemanagementResponseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response bad request response has a 4xx status code
func (o *GetResponsemanagementResponseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response bad request response has a 5xx status code
func (o *GetResponsemanagementResponseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response bad request response a status code equal to that given
func (o *GetResponsemanagementResponseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetResponsemanagementResponseBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponseBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseBadRequest  %+v", 400, o.Payload)
}

func (o *GetResponsemanagementResponseBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseUnauthorized creates a GetResponsemanagementResponseUnauthorized with default headers values
func NewGetResponsemanagementResponseUnauthorized() *GetResponsemanagementResponseUnauthorized {
	return &GetResponsemanagementResponseUnauthorized{}
}

/*
GetResponsemanagementResponseUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetResponsemanagementResponseUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response unauthorized response has a 2xx status code
func (o *GetResponsemanagementResponseUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response unauthorized response has a 3xx status code
func (o *GetResponsemanagementResponseUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response unauthorized response has a 4xx status code
func (o *GetResponsemanagementResponseUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response unauthorized response has a 5xx status code
func (o *GetResponsemanagementResponseUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response unauthorized response a status code equal to that given
func (o *GetResponsemanagementResponseUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetResponsemanagementResponseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponseUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseUnauthorized  %+v", 401, o.Payload)
}

func (o *GetResponsemanagementResponseUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseForbidden creates a GetResponsemanagementResponseForbidden with default headers values
func NewGetResponsemanagementResponseForbidden() *GetResponsemanagementResponseForbidden {
	return &GetResponsemanagementResponseForbidden{}
}

/*
GetResponsemanagementResponseForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetResponsemanagementResponseForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response forbidden response has a 2xx status code
func (o *GetResponsemanagementResponseForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response forbidden response has a 3xx status code
func (o *GetResponsemanagementResponseForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response forbidden response has a 4xx status code
func (o *GetResponsemanagementResponseForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response forbidden response has a 5xx status code
func (o *GetResponsemanagementResponseForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response forbidden response a status code equal to that given
func (o *GetResponsemanagementResponseForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetResponsemanagementResponseForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponseForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseForbidden  %+v", 403, o.Payload)
}

func (o *GetResponsemanagementResponseForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseNotFound creates a GetResponsemanagementResponseNotFound with default headers values
func NewGetResponsemanagementResponseNotFound() *GetResponsemanagementResponseNotFound {
	return &GetResponsemanagementResponseNotFound{}
}

/*
GetResponsemanagementResponseNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetResponsemanagementResponseNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response not found response has a 2xx status code
func (o *GetResponsemanagementResponseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response not found response has a 3xx status code
func (o *GetResponsemanagementResponseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response not found response has a 4xx status code
func (o *GetResponsemanagementResponseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response not found response has a 5xx status code
func (o *GetResponsemanagementResponseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response not found response a status code equal to that given
func (o *GetResponsemanagementResponseNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetResponsemanagementResponseNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponseNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseNotFound  %+v", 404, o.Payload)
}

func (o *GetResponsemanagementResponseNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseRequestTimeout creates a GetResponsemanagementResponseRequestTimeout with default headers values
func NewGetResponsemanagementResponseRequestTimeout() *GetResponsemanagementResponseRequestTimeout {
	return &GetResponsemanagementResponseRequestTimeout{}
}

/*
GetResponsemanagementResponseRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetResponsemanagementResponseRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response request timeout response has a 2xx status code
func (o *GetResponsemanagementResponseRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response request timeout response has a 3xx status code
func (o *GetResponsemanagementResponseRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response request timeout response has a 4xx status code
func (o *GetResponsemanagementResponseRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response request timeout response has a 5xx status code
func (o *GetResponsemanagementResponseRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response request timeout response a status code equal to that given
func (o *GetResponsemanagementResponseRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetResponsemanagementResponseRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponseRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetResponsemanagementResponseRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseRequestEntityTooLarge creates a GetResponsemanagementResponseRequestEntityTooLarge with default headers values
func NewGetResponsemanagementResponseRequestEntityTooLarge() *GetResponsemanagementResponseRequestEntityTooLarge {
	return &GetResponsemanagementResponseRequestEntityTooLarge{}
}

/*
GetResponsemanagementResponseRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetResponsemanagementResponseRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response request entity too large response has a 2xx status code
func (o *GetResponsemanagementResponseRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response request entity too large response has a 3xx status code
func (o *GetResponsemanagementResponseRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response request entity too large response has a 4xx status code
func (o *GetResponsemanagementResponseRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response request entity too large response has a 5xx status code
func (o *GetResponsemanagementResponseRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response request entity too large response a status code equal to that given
func (o *GetResponsemanagementResponseRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseUnsupportedMediaType creates a GetResponsemanagementResponseUnsupportedMediaType with default headers values
func NewGetResponsemanagementResponseUnsupportedMediaType() *GetResponsemanagementResponseUnsupportedMediaType {
	return &GetResponsemanagementResponseUnsupportedMediaType{}
}

/*
GetResponsemanagementResponseUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetResponsemanagementResponseUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response unsupported media type response has a 2xx status code
func (o *GetResponsemanagementResponseUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response unsupported media type response has a 3xx status code
func (o *GetResponsemanagementResponseUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response unsupported media type response has a 4xx status code
func (o *GetResponsemanagementResponseUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response unsupported media type response has a 5xx status code
func (o *GetResponsemanagementResponseUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response unsupported media type response a status code equal to that given
func (o *GetResponsemanagementResponseUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseTooManyRequests creates a GetResponsemanagementResponseTooManyRequests with default headers values
func NewGetResponsemanagementResponseTooManyRequests() *GetResponsemanagementResponseTooManyRequests {
	return &GetResponsemanagementResponseTooManyRequests{}
}

/*
GetResponsemanagementResponseTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetResponsemanagementResponseTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response too many requests response has a 2xx status code
func (o *GetResponsemanagementResponseTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response too many requests response has a 3xx status code
func (o *GetResponsemanagementResponseTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response too many requests response has a 4xx status code
func (o *GetResponsemanagementResponseTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get responsemanagement response too many requests response has a 5xx status code
func (o *GetResponsemanagementResponseTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get responsemanagement response too many requests response a status code equal to that given
func (o *GetResponsemanagementResponseTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetResponsemanagementResponseTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponseTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetResponsemanagementResponseTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseInternalServerError creates a GetResponsemanagementResponseInternalServerError with default headers values
func NewGetResponsemanagementResponseInternalServerError() *GetResponsemanagementResponseInternalServerError {
	return &GetResponsemanagementResponseInternalServerError{}
}

/*
GetResponsemanagementResponseInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetResponsemanagementResponseInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response internal server error response has a 2xx status code
func (o *GetResponsemanagementResponseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response internal server error response has a 3xx status code
func (o *GetResponsemanagementResponseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response internal server error response has a 4xx status code
func (o *GetResponsemanagementResponseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement response internal server error response has a 5xx status code
func (o *GetResponsemanagementResponseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement response internal server error response a status code equal to that given
func (o *GetResponsemanagementResponseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetResponsemanagementResponseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponseInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseInternalServerError  %+v", 500, o.Payload)
}

func (o *GetResponsemanagementResponseInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseServiceUnavailable creates a GetResponsemanagementResponseServiceUnavailable with default headers values
func NewGetResponsemanagementResponseServiceUnavailable() *GetResponsemanagementResponseServiceUnavailable {
	return &GetResponsemanagementResponseServiceUnavailable{}
}

/*
GetResponsemanagementResponseServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetResponsemanagementResponseServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response service unavailable response has a 2xx status code
func (o *GetResponsemanagementResponseServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response service unavailable response has a 3xx status code
func (o *GetResponsemanagementResponseServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response service unavailable response has a 4xx status code
func (o *GetResponsemanagementResponseServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement response service unavailable response has a 5xx status code
func (o *GetResponsemanagementResponseServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement response service unavailable response a status code equal to that given
func (o *GetResponsemanagementResponseServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetResponsemanagementResponseServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponseServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetResponsemanagementResponseServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResponsemanagementResponseGatewayTimeout creates a GetResponsemanagementResponseGatewayTimeout with default headers values
func NewGetResponsemanagementResponseGatewayTimeout() *GetResponsemanagementResponseGatewayTimeout {
	return &GetResponsemanagementResponseGatewayTimeout{}
}

/*
GetResponsemanagementResponseGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetResponsemanagementResponseGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get responsemanagement response gateway timeout response has a 2xx status code
func (o *GetResponsemanagementResponseGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get responsemanagement response gateway timeout response has a 3xx status code
func (o *GetResponsemanagementResponseGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get responsemanagement response gateway timeout response has a 4xx status code
func (o *GetResponsemanagementResponseGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get responsemanagement response gateway timeout response has a 5xx status code
func (o *GetResponsemanagementResponseGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get responsemanagement response gateway timeout response a status code equal to that given
func (o *GetResponsemanagementResponseGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetResponsemanagementResponseGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponseGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/responsemanagement/responses/{responseId}][%d] getResponsemanagementResponseGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetResponsemanagementResponseGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetResponsemanagementResponseGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
