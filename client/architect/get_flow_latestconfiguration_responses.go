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

// GetFlowLatestconfigurationReader is a Reader for the GetFlowLatestconfiguration structure.
type GetFlowLatestconfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowLatestconfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowLatestconfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowLatestconfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowLatestconfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowLatestconfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowLatestconfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowLatestconfigurationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetFlowLatestconfigurationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowLatestconfigurationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowLatestconfigurationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowLatestconfigurationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowLatestconfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowLatestconfigurationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowLatestconfigurationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowLatestconfigurationOK creates a GetFlowLatestconfigurationOK with default headers values
func NewGetFlowLatestconfigurationOK() *GetFlowLatestconfigurationOK {
	return &GetFlowLatestconfigurationOK{}
}

/*
GetFlowLatestconfigurationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFlowLatestconfigurationOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get flow latestconfiguration o k response has a 2xx status code
func (o *GetFlowLatestconfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get flow latestconfiguration o k response has a 3xx status code
func (o *GetFlowLatestconfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration o k response has a 4xx status code
func (o *GetFlowLatestconfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow latestconfiguration o k response has a 5xx status code
func (o *GetFlowLatestconfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration o k response a status code equal to that given
func (o *GetFlowLatestconfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetFlowLatestconfigurationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFlowLatestconfigurationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFlowLatestconfigurationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetFlowLatestconfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationBadRequest creates a GetFlowLatestconfigurationBadRequest with default headers values
func NewGetFlowLatestconfigurationBadRequest() *GetFlowLatestconfigurationBadRequest {
	return &GetFlowLatestconfigurationBadRequest{}
}

/*
GetFlowLatestconfigurationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowLatestconfigurationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration bad request response has a 2xx status code
func (o *GetFlowLatestconfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration bad request response has a 3xx status code
func (o *GetFlowLatestconfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration bad request response has a 4xx status code
func (o *GetFlowLatestconfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration bad request response has a 5xx status code
func (o *GetFlowLatestconfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration bad request response a status code equal to that given
func (o *GetFlowLatestconfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetFlowLatestconfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowLatestconfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowLatestconfigurationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationUnauthorized creates a GetFlowLatestconfigurationUnauthorized with default headers values
func NewGetFlowLatestconfigurationUnauthorized() *GetFlowLatestconfigurationUnauthorized {
	return &GetFlowLatestconfigurationUnauthorized{}
}

/*
GetFlowLatestconfigurationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowLatestconfigurationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration unauthorized response has a 2xx status code
func (o *GetFlowLatestconfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration unauthorized response has a 3xx status code
func (o *GetFlowLatestconfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration unauthorized response has a 4xx status code
func (o *GetFlowLatestconfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration unauthorized response has a 5xx status code
func (o *GetFlowLatestconfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration unauthorized response a status code equal to that given
func (o *GetFlowLatestconfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetFlowLatestconfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowLatestconfigurationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowLatestconfigurationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationForbidden creates a GetFlowLatestconfigurationForbidden with default headers values
func NewGetFlowLatestconfigurationForbidden() *GetFlowLatestconfigurationForbidden {
	return &GetFlowLatestconfigurationForbidden{}
}

/*
GetFlowLatestconfigurationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowLatestconfigurationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration forbidden response has a 2xx status code
func (o *GetFlowLatestconfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration forbidden response has a 3xx status code
func (o *GetFlowLatestconfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration forbidden response has a 4xx status code
func (o *GetFlowLatestconfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration forbidden response has a 5xx status code
func (o *GetFlowLatestconfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration forbidden response a status code equal to that given
func (o *GetFlowLatestconfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetFlowLatestconfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowLatestconfigurationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowLatestconfigurationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationNotFound creates a GetFlowLatestconfigurationNotFound with default headers values
func NewGetFlowLatestconfigurationNotFound() *GetFlowLatestconfigurationNotFound {
	return &GetFlowLatestconfigurationNotFound{}
}

/*
GetFlowLatestconfigurationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetFlowLatestconfigurationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration not found response has a 2xx status code
func (o *GetFlowLatestconfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration not found response has a 3xx status code
func (o *GetFlowLatestconfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration not found response has a 4xx status code
func (o *GetFlowLatestconfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration not found response has a 5xx status code
func (o *GetFlowLatestconfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration not found response a status code equal to that given
func (o *GetFlowLatestconfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetFlowLatestconfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowLatestconfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowLatestconfigurationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationMethodNotAllowed creates a GetFlowLatestconfigurationMethodNotAllowed with default headers values
func NewGetFlowLatestconfigurationMethodNotAllowed() *GetFlowLatestconfigurationMethodNotAllowed {
	return &GetFlowLatestconfigurationMethodNotAllowed{}
}

/*
GetFlowLatestconfigurationMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetFlowLatestconfigurationMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration method not allowed response has a 2xx status code
func (o *GetFlowLatestconfigurationMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration method not allowed response has a 3xx status code
func (o *GetFlowLatestconfigurationMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration method not allowed response has a 4xx status code
func (o *GetFlowLatestconfigurationMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration method not allowed response has a 5xx status code
func (o *GetFlowLatestconfigurationMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration method not allowed response a status code equal to that given
func (o *GetFlowLatestconfigurationMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationRequestTimeout creates a GetFlowLatestconfigurationRequestTimeout with default headers values
func NewGetFlowLatestconfigurationRequestTimeout() *GetFlowLatestconfigurationRequestTimeout {
	return &GetFlowLatestconfigurationRequestTimeout{}
}

/*
GetFlowLatestconfigurationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetFlowLatestconfigurationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration request timeout response has a 2xx status code
func (o *GetFlowLatestconfigurationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration request timeout response has a 3xx status code
func (o *GetFlowLatestconfigurationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration request timeout response has a 4xx status code
func (o *GetFlowLatestconfigurationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration request timeout response has a 5xx status code
func (o *GetFlowLatestconfigurationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration request timeout response a status code equal to that given
func (o *GetFlowLatestconfigurationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetFlowLatestconfigurationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowLatestconfigurationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetFlowLatestconfigurationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationRequestEntityTooLarge creates a GetFlowLatestconfigurationRequestEntityTooLarge with default headers values
func NewGetFlowLatestconfigurationRequestEntityTooLarge() *GetFlowLatestconfigurationRequestEntityTooLarge {
	return &GetFlowLatestconfigurationRequestEntityTooLarge{}
}

/*
GetFlowLatestconfigurationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetFlowLatestconfigurationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration request entity too large response has a 2xx status code
func (o *GetFlowLatestconfigurationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration request entity too large response has a 3xx status code
func (o *GetFlowLatestconfigurationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration request entity too large response has a 4xx status code
func (o *GetFlowLatestconfigurationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration request entity too large response has a 5xx status code
func (o *GetFlowLatestconfigurationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration request entity too large response a status code equal to that given
func (o *GetFlowLatestconfigurationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationUnsupportedMediaType creates a GetFlowLatestconfigurationUnsupportedMediaType with default headers values
func NewGetFlowLatestconfigurationUnsupportedMediaType() *GetFlowLatestconfigurationUnsupportedMediaType {
	return &GetFlowLatestconfigurationUnsupportedMediaType{}
}

/*
GetFlowLatestconfigurationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowLatestconfigurationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration unsupported media type response has a 2xx status code
func (o *GetFlowLatestconfigurationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration unsupported media type response has a 3xx status code
func (o *GetFlowLatestconfigurationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration unsupported media type response has a 4xx status code
func (o *GetFlowLatestconfigurationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration unsupported media type response has a 5xx status code
func (o *GetFlowLatestconfigurationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration unsupported media type response a status code equal to that given
func (o *GetFlowLatestconfigurationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationTooManyRequests creates a GetFlowLatestconfigurationTooManyRequests with default headers values
func NewGetFlowLatestconfigurationTooManyRequests() *GetFlowLatestconfigurationTooManyRequests {
	return &GetFlowLatestconfigurationTooManyRequests{}
}

/*
GetFlowLatestconfigurationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetFlowLatestconfigurationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration too many requests response has a 2xx status code
func (o *GetFlowLatestconfigurationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration too many requests response has a 3xx status code
func (o *GetFlowLatestconfigurationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration too many requests response has a 4xx status code
func (o *GetFlowLatestconfigurationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get flow latestconfiguration too many requests response has a 5xx status code
func (o *GetFlowLatestconfigurationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get flow latestconfiguration too many requests response a status code equal to that given
func (o *GetFlowLatestconfigurationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetFlowLatestconfigurationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowLatestconfigurationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowLatestconfigurationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationInternalServerError creates a GetFlowLatestconfigurationInternalServerError with default headers values
func NewGetFlowLatestconfigurationInternalServerError() *GetFlowLatestconfigurationInternalServerError {
	return &GetFlowLatestconfigurationInternalServerError{}
}

/*
GetFlowLatestconfigurationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowLatestconfigurationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration internal server error response has a 2xx status code
func (o *GetFlowLatestconfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration internal server error response has a 3xx status code
func (o *GetFlowLatestconfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration internal server error response has a 4xx status code
func (o *GetFlowLatestconfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow latestconfiguration internal server error response has a 5xx status code
func (o *GetFlowLatestconfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get flow latestconfiguration internal server error response a status code equal to that given
func (o *GetFlowLatestconfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetFlowLatestconfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowLatestconfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowLatestconfigurationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationServiceUnavailable creates a GetFlowLatestconfigurationServiceUnavailable with default headers values
func NewGetFlowLatestconfigurationServiceUnavailable() *GetFlowLatestconfigurationServiceUnavailable {
	return &GetFlowLatestconfigurationServiceUnavailable{}
}

/*
GetFlowLatestconfigurationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowLatestconfigurationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration service unavailable response has a 2xx status code
func (o *GetFlowLatestconfigurationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration service unavailable response has a 3xx status code
func (o *GetFlowLatestconfigurationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration service unavailable response has a 4xx status code
func (o *GetFlowLatestconfigurationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow latestconfiguration service unavailable response has a 5xx status code
func (o *GetFlowLatestconfigurationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get flow latestconfiguration service unavailable response a status code equal to that given
func (o *GetFlowLatestconfigurationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetFlowLatestconfigurationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowLatestconfigurationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowLatestconfigurationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowLatestconfigurationGatewayTimeout creates a GetFlowLatestconfigurationGatewayTimeout with default headers values
func NewGetFlowLatestconfigurationGatewayTimeout() *GetFlowLatestconfigurationGatewayTimeout {
	return &GetFlowLatestconfigurationGatewayTimeout{}
}

/*
GetFlowLatestconfigurationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetFlowLatestconfigurationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get flow latestconfiguration gateway timeout response has a 2xx status code
func (o *GetFlowLatestconfigurationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get flow latestconfiguration gateway timeout response has a 3xx status code
func (o *GetFlowLatestconfigurationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get flow latestconfiguration gateway timeout response has a 4xx status code
func (o *GetFlowLatestconfigurationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get flow latestconfiguration gateway timeout response has a 5xx status code
func (o *GetFlowLatestconfigurationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get flow latestconfiguration gateway timeout response a status code equal to that given
func (o *GetFlowLatestconfigurationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetFlowLatestconfigurationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowLatestconfigurationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/flows/{flowId}/latestconfiguration][%d] getFlowLatestconfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowLatestconfigurationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowLatestconfigurationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
