// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFlowVersionConfigurationReader is a Reader for the GetFlowVersionConfiguration structure.
type GetFlowVersionConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowVersionConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowVersionConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowVersionConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowVersionConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowVersionConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowVersionConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowVersionConfigurationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowVersionConfigurationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowVersionConfigurationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowVersionConfigurationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowVersionConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowVersionConfigurationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowVersionConfigurationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowVersionConfigurationOK creates a GetFlowVersionConfigurationOK with default headers values
func NewGetFlowVersionConfigurationOK() *GetFlowVersionConfigurationOK {
	return &GetFlowVersionConfigurationOK{}
}

/*
GetFlowVersionConfigurationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowVersionConfigurationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get flow version configuration o k response has a 2xx status code
func (o *GetFlowVersionConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flow version configuration o k response has a 3xx status code
func (o *GetFlowVersionConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration o k response has a 4xx status code
func (o *GetFlowVersionConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow version configuration o k response has a 5xx status code
func (o *GetFlowVersionConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration o k response a status code equal to that given
func (o *GetFlowVersionConfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFlowVersionConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFlowVersionConfigurationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFlowVersionConfigurationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetFlowVersionConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationBadRequest creates a GetFlowVersionConfigurationBadRequest with default headers values
func NewGetFlowVersionConfigurationBadRequest() *GetFlowVersionConfigurationBadRequest {
	return &GetFlowVersionConfigurationBadRequest{}
}

/*
GetFlowVersionConfigurationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowVersionConfigurationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration bad request response has a 2xx status code
func (o *GetFlowVersionConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration bad request response has a 3xx status code
func (o *GetFlowVersionConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration bad request response has a 4xx status code
func (o *GetFlowVersionConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration bad request response has a 5xx status code
func (o *GetFlowVersionConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration bad request response a status code equal to that given
func (o *GetFlowVersionConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFlowVersionConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowVersionConfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowVersionConfigurationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationUnauthorized creates a GetFlowVersionConfigurationUnauthorized with default headers values
func NewGetFlowVersionConfigurationUnauthorized() *GetFlowVersionConfigurationUnauthorized {
	return &GetFlowVersionConfigurationUnauthorized{}
}

/*
GetFlowVersionConfigurationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowVersionConfigurationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration unauthorized response has a 2xx status code
func (o *GetFlowVersionConfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration unauthorized response has a 3xx status code
func (o *GetFlowVersionConfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration unauthorized response has a 4xx status code
func (o *GetFlowVersionConfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration unauthorized response has a 5xx status code
func (o *GetFlowVersionConfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration unauthorized response a status code equal to that given
func (o *GetFlowVersionConfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFlowVersionConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowVersionConfigurationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowVersionConfigurationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationForbidden creates a GetFlowVersionConfigurationForbidden with default headers values
func NewGetFlowVersionConfigurationForbidden() *GetFlowVersionConfigurationForbidden {
	return &GetFlowVersionConfigurationForbidden{}
}

/*
GetFlowVersionConfigurationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowVersionConfigurationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration forbidden response has a 2xx status code
func (o *GetFlowVersionConfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration forbidden response has a 3xx status code
func (o *GetFlowVersionConfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration forbidden response has a 4xx status code
func (o *GetFlowVersionConfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration forbidden response has a 5xx status code
func (o *GetFlowVersionConfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration forbidden response a status code equal to that given
func (o *GetFlowVersionConfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFlowVersionConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowVersionConfigurationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowVersionConfigurationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationNotFound creates a GetFlowVersionConfigurationNotFound with default headers values
func NewGetFlowVersionConfigurationNotFound() *GetFlowVersionConfigurationNotFound {
	return &GetFlowVersionConfigurationNotFound{}
}

/*
GetFlowVersionConfigurationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFlowVersionConfigurationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration not found response has a 2xx status code
func (o *GetFlowVersionConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration not found response has a 3xx status code
func (o *GetFlowVersionConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration not found response has a 4xx status code
func (o *GetFlowVersionConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration not found response has a 5xx status code
func (o *GetFlowVersionConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration not found response a status code equal to that given
func (o *GetFlowVersionConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFlowVersionConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowVersionConfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowVersionConfigurationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationRequestTimeout creates a GetFlowVersionConfigurationRequestTimeout with default headers values
func NewGetFlowVersionConfigurationRequestTimeout() *GetFlowVersionConfigurationRequestTimeout {
	return &GetFlowVersionConfigurationRequestTimeout{}
}

/*
GetFlowVersionConfigurationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowVersionConfigurationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration request timeout response has a 2xx status code
func (o *GetFlowVersionConfigurationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration request timeout response has a 3xx status code
func (o *GetFlowVersionConfigurationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration request timeout response has a 4xx status code
func (o *GetFlowVersionConfigurationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration request timeout response has a 5xx status code
func (o *GetFlowVersionConfigurationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration request timeout response a status code equal to that given
func (o *GetFlowVersionConfigurationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFlowVersionConfigurationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowVersionConfigurationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowVersionConfigurationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationRequestEntityTooLarge creates a GetFlowVersionConfigurationRequestEntityTooLarge with default headers values
func NewGetFlowVersionConfigurationRequestEntityTooLarge() *GetFlowVersionConfigurationRequestEntityTooLarge {
	return &GetFlowVersionConfigurationRequestEntityTooLarge{}
}

/*
GetFlowVersionConfigurationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetFlowVersionConfigurationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration request entity too large response has a 2xx status code
func (o *GetFlowVersionConfigurationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration request entity too large response has a 3xx status code
func (o *GetFlowVersionConfigurationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration request entity too large response has a 4xx status code
func (o *GetFlowVersionConfigurationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration request entity too large response has a 5xx status code
func (o *GetFlowVersionConfigurationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration request entity too large response a status code equal to that given
func (o *GetFlowVersionConfigurationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFlowVersionConfigurationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowVersionConfigurationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowVersionConfigurationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationUnsupportedMediaType creates a GetFlowVersionConfigurationUnsupportedMediaType with default headers values
func NewGetFlowVersionConfigurationUnsupportedMediaType() *GetFlowVersionConfigurationUnsupportedMediaType {
	return &GetFlowVersionConfigurationUnsupportedMediaType{}
}

/*
GetFlowVersionConfigurationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowVersionConfigurationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration unsupported media type response has a 2xx status code
func (o *GetFlowVersionConfigurationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration unsupported media type response has a 3xx status code
func (o *GetFlowVersionConfigurationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration unsupported media type response has a 4xx status code
func (o *GetFlowVersionConfigurationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration unsupported media type response has a 5xx status code
func (o *GetFlowVersionConfigurationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration unsupported media type response a status code equal to that given
func (o *GetFlowVersionConfigurationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFlowVersionConfigurationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowVersionConfigurationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowVersionConfigurationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationTooManyRequests creates a GetFlowVersionConfigurationTooManyRequests with default headers values
func NewGetFlowVersionConfigurationTooManyRequests() *GetFlowVersionConfigurationTooManyRequests {
	return &GetFlowVersionConfigurationTooManyRequests{}
}

/*
GetFlowVersionConfigurationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowVersionConfigurationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration too many requests response has a 2xx status code
func (o *GetFlowVersionConfigurationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration too many requests response has a 3xx status code
func (o *GetFlowVersionConfigurationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration too many requests response has a 4xx status code
func (o *GetFlowVersionConfigurationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow version configuration too many requests response has a 5xx status code
func (o *GetFlowVersionConfigurationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow version configuration too many requests response a status code equal to that given
func (o *GetFlowVersionConfigurationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFlowVersionConfigurationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowVersionConfigurationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowVersionConfigurationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationInternalServerError creates a GetFlowVersionConfigurationInternalServerError with default headers values
func NewGetFlowVersionConfigurationInternalServerError() *GetFlowVersionConfigurationInternalServerError {
	return &GetFlowVersionConfigurationInternalServerError{}
}

/*
GetFlowVersionConfigurationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowVersionConfigurationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration internal server error response has a 2xx status code
func (o *GetFlowVersionConfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration internal server error response has a 3xx status code
func (o *GetFlowVersionConfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration internal server error response has a 4xx status code
func (o *GetFlowVersionConfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow version configuration internal server error response has a 5xx status code
func (o *GetFlowVersionConfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get flow version configuration internal server error response a status code equal to that given
func (o *GetFlowVersionConfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFlowVersionConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowVersionConfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowVersionConfigurationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationServiceUnavailable creates a GetFlowVersionConfigurationServiceUnavailable with default headers values
func NewGetFlowVersionConfigurationServiceUnavailable() *GetFlowVersionConfigurationServiceUnavailable {
	return &GetFlowVersionConfigurationServiceUnavailable{}
}

/*
GetFlowVersionConfigurationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowVersionConfigurationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration service unavailable response has a 2xx status code
func (o *GetFlowVersionConfigurationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration service unavailable response has a 3xx status code
func (o *GetFlowVersionConfigurationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration service unavailable response has a 4xx status code
func (o *GetFlowVersionConfigurationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow version configuration service unavailable response has a 5xx status code
func (o *GetFlowVersionConfigurationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get flow version configuration service unavailable response a status code equal to that given
func (o *GetFlowVersionConfigurationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFlowVersionConfigurationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowVersionConfigurationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowVersionConfigurationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowVersionConfigurationGatewayTimeout creates a GetFlowVersionConfigurationGatewayTimeout with default headers values
func NewGetFlowVersionConfigurationGatewayTimeout() *GetFlowVersionConfigurationGatewayTimeout {
	return &GetFlowVersionConfigurationGatewayTimeout{}
}

/*
GetFlowVersionConfigurationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFlowVersionConfigurationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow version configuration gateway timeout response has a 2xx status code
func (o *GetFlowVersionConfigurationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow version configuration gateway timeout response has a 3xx status code
func (o *GetFlowVersionConfigurationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow version configuration gateway timeout response has a 4xx status code
func (o *GetFlowVersionConfigurationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow version configuration gateway timeout response has a 5xx status code
func (o *GetFlowVersionConfigurationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get flow version configuration gateway timeout response a status code equal to that given
func (o *GetFlowVersionConfigurationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFlowVersionConfigurationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowVersionConfigurationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/versions/{versionId}/configuration][%d] getFlowVersionConfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowVersionConfigurationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowVersionConfigurationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
