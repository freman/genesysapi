// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingUtilizationReader is a Reader for the GetRoutingUtilization structure.
type GetRoutingUtilizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingUtilizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingUtilizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingUtilizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingUtilizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingUtilizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingUtilizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingUtilizationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingUtilizationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingUtilizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingUtilizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingUtilizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingUtilizationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingUtilizationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingUtilizationOK creates a GetRoutingUtilizationOK with default headers values
func NewGetRoutingUtilizationOK() *GetRoutingUtilizationOK {
	return &GetRoutingUtilizationOK{}
}

/*
GetRoutingUtilizationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingUtilizationOK struct {
	Payload *models.Utilization
}

// IsSuccess returns true when this get routing utilization o k response has a 2xx status code
func (o *GetRoutingUtilizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing utilization o k response has a 3xx status code
func (o *GetRoutingUtilizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization o k response has a 4xx status code
func (o *GetRoutingUtilizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing utilization o k response has a 5xx status code
func (o *GetRoutingUtilizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization o k response a status code equal to that given
func (o *GetRoutingUtilizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingUtilizationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationOK  %+v", 200, o.Payload)
}

func (o *GetRoutingUtilizationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationOK  %+v", 200, o.Payload)
}

func (o *GetRoutingUtilizationOK) GetPayload() *models.Utilization {
	return o.Payload
}

func (o *GetRoutingUtilizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Utilization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationBadRequest creates a GetRoutingUtilizationBadRequest with default headers values
func NewGetRoutingUtilizationBadRequest() *GetRoutingUtilizationBadRequest {
	return &GetRoutingUtilizationBadRequest{}
}

/*
GetRoutingUtilizationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingUtilizationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization bad request response has a 2xx status code
func (o *GetRoutingUtilizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization bad request response has a 3xx status code
func (o *GetRoutingUtilizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization bad request response has a 4xx status code
func (o *GetRoutingUtilizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization bad request response has a 5xx status code
func (o *GetRoutingUtilizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization bad request response a status code equal to that given
func (o *GetRoutingUtilizationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingUtilizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingUtilizationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingUtilizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationUnauthorized creates a GetRoutingUtilizationUnauthorized with default headers values
func NewGetRoutingUtilizationUnauthorized() *GetRoutingUtilizationUnauthorized {
	return &GetRoutingUtilizationUnauthorized{}
}

/*
GetRoutingUtilizationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingUtilizationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization unauthorized response has a 2xx status code
func (o *GetRoutingUtilizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization unauthorized response has a 3xx status code
func (o *GetRoutingUtilizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization unauthorized response has a 4xx status code
func (o *GetRoutingUtilizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization unauthorized response has a 5xx status code
func (o *GetRoutingUtilizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization unauthorized response a status code equal to that given
func (o *GetRoutingUtilizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingUtilizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingUtilizationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingUtilizationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationForbidden creates a GetRoutingUtilizationForbidden with default headers values
func NewGetRoutingUtilizationForbidden() *GetRoutingUtilizationForbidden {
	return &GetRoutingUtilizationForbidden{}
}

/*
GetRoutingUtilizationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingUtilizationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization forbidden response has a 2xx status code
func (o *GetRoutingUtilizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization forbidden response has a 3xx status code
func (o *GetRoutingUtilizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization forbidden response has a 4xx status code
func (o *GetRoutingUtilizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization forbidden response has a 5xx status code
func (o *GetRoutingUtilizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization forbidden response a status code equal to that given
func (o *GetRoutingUtilizationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingUtilizationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingUtilizationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingUtilizationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationNotFound creates a GetRoutingUtilizationNotFound with default headers values
func NewGetRoutingUtilizationNotFound() *GetRoutingUtilizationNotFound {
	return &GetRoutingUtilizationNotFound{}
}

/*
GetRoutingUtilizationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingUtilizationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization not found response has a 2xx status code
func (o *GetRoutingUtilizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization not found response has a 3xx status code
func (o *GetRoutingUtilizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization not found response has a 4xx status code
func (o *GetRoutingUtilizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization not found response has a 5xx status code
func (o *GetRoutingUtilizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization not found response a status code equal to that given
func (o *GetRoutingUtilizationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingUtilizationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingUtilizationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingUtilizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationRequestTimeout creates a GetRoutingUtilizationRequestTimeout with default headers values
func NewGetRoutingUtilizationRequestTimeout() *GetRoutingUtilizationRequestTimeout {
	return &GetRoutingUtilizationRequestTimeout{}
}

/*
GetRoutingUtilizationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingUtilizationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization request timeout response has a 2xx status code
func (o *GetRoutingUtilizationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization request timeout response has a 3xx status code
func (o *GetRoutingUtilizationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization request timeout response has a 4xx status code
func (o *GetRoutingUtilizationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization request timeout response has a 5xx status code
func (o *GetRoutingUtilizationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization request timeout response a status code equal to that given
func (o *GetRoutingUtilizationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingUtilizationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingUtilizationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingUtilizationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationRequestEntityTooLarge creates a GetRoutingUtilizationRequestEntityTooLarge with default headers values
func NewGetRoutingUtilizationRequestEntityTooLarge() *GetRoutingUtilizationRequestEntityTooLarge {
	return &GetRoutingUtilizationRequestEntityTooLarge{}
}

/*
GetRoutingUtilizationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingUtilizationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization request entity too large response has a 2xx status code
func (o *GetRoutingUtilizationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization request entity too large response has a 3xx status code
func (o *GetRoutingUtilizationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization request entity too large response has a 4xx status code
func (o *GetRoutingUtilizationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization request entity too large response has a 5xx status code
func (o *GetRoutingUtilizationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization request entity too large response a status code equal to that given
func (o *GetRoutingUtilizationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingUtilizationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingUtilizationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingUtilizationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationUnsupportedMediaType creates a GetRoutingUtilizationUnsupportedMediaType with default headers values
func NewGetRoutingUtilizationUnsupportedMediaType() *GetRoutingUtilizationUnsupportedMediaType {
	return &GetRoutingUtilizationUnsupportedMediaType{}
}

/*
GetRoutingUtilizationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingUtilizationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization unsupported media type response has a 2xx status code
func (o *GetRoutingUtilizationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization unsupported media type response has a 3xx status code
func (o *GetRoutingUtilizationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization unsupported media type response has a 4xx status code
func (o *GetRoutingUtilizationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization unsupported media type response has a 5xx status code
func (o *GetRoutingUtilizationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization unsupported media type response a status code equal to that given
func (o *GetRoutingUtilizationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingUtilizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingUtilizationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingUtilizationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationTooManyRequests creates a GetRoutingUtilizationTooManyRequests with default headers values
func NewGetRoutingUtilizationTooManyRequests() *GetRoutingUtilizationTooManyRequests {
	return &GetRoutingUtilizationTooManyRequests{}
}

/*
GetRoutingUtilizationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingUtilizationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization too many requests response has a 2xx status code
func (o *GetRoutingUtilizationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization too many requests response has a 3xx status code
func (o *GetRoutingUtilizationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization too many requests response has a 4xx status code
func (o *GetRoutingUtilizationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing utilization too many requests response has a 5xx status code
func (o *GetRoutingUtilizationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing utilization too many requests response a status code equal to that given
func (o *GetRoutingUtilizationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingUtilizationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingUtilizationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingUtilizationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationInternalServerError creates a GetRoutingUtilizationInternalServerError with default headers values
func NewGetRoutingUtilizationInternalServerError() *GetRoutingUtilizationInternalServerError {
	return &GetRoutingUtilizationInternalServerError{}
}

/*
GetRoutingUtilizationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingUtilizationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization internal server error response has a 2xx status code
func (o *GetRoutingUtilizationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization internal server error response has a 3xx status code
func (o *GetRoutingUtilizationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization internal server error response has a 4xx status code
func (o *GetRoutingUtilizationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing utilization internal server error response has a 5xx status code
func (o *GetRoutingUtilizationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing utilization internal server error response a status code equal to that given
func (o *GetRoutingUtilizationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingUtilizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingUtilizationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingUtilizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationServiceUnavailable creates a GetRoutingUtilizationServiceUnavailable with default headers values
func NewGetRoutingUtilizationServiceUnavailable() *GetRoutingUtilizationServiceUnavailable {
	return &GetRoutingUtilizationServiceUnavailable{}
}

/*
GetRoutingUtilizationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingUtilizationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization service unavailable response has a 2xx status code
func (o *GetRoutingUtilizationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization service unavailable response has a 3xx status code
func (o *GetRoutingUtilizationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization service unavailable response has a 4xx status code
func (o *GetRoutingUtilizationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing utilization service unavailable response has a 5xx status code
func (o *GetRoutingUtilizationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing utilization service unavailable response a status code equal to that given
func (o *GetRoutingUtilizationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingUtilizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingUtilizationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingUtilizationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingUtilizationGatewayTimeout creates a GetRoutingUtilizationGatewayTimeout with default headers values
func NewGetRoutingUtilizationGatewayTimeout() *GetRoutingUtilizationGatewayTimeout {
	return &GetRoutingUtilizationGatewayTimeout{}
}

/*
GetRoutingUtilizationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingUtilizationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing utilization gateway timeout response has a 2xx status code
func (o *GetRoutingUtilizationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing utilization gateway timeout response has a 3xx status code
func (o *GetRoutingUtilizationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing utilization gateway timeout response has a 4xx status code
func (o *GetRoutingUtilizationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing utilization gateway timeout response has a 5xx status code
func (o *GetRoutingUtilizationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing utilization gateway timeout response a status code equal to that given
func (o *GetRoutingUtilizationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingUtilizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingUtilizationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/utilization][%d] getRoutingUtilizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingUtilizationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingUtilizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
