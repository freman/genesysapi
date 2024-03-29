// Code generated by go-swagger; DO NOT EDIT.

package stations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetStationReader is a Reader for the GetStation structure.
type GetStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetStationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetStationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetStationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetStationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetStationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStationOK creates a GetStationOK with default headers values
func NewGetStationOK() *GetStationOK {
	return &GetStationOK{}
}

/*
GetStationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetStationOK struct {
	Payload *models.Station
}

// IsSuccess returns true when this get station o k response has a 2xx status code
func (o *GetStationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get station o k response has a 3xx status code
func (o *GetStationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station o k response has a 4xx status code
func (o *GetStationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get station o k response has a 5xx status code
func (o *GetStationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get station o k response a status code equal to that given
func (o *GetStationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetStationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationOK  %+v", 200, o.Payload)
}

func (o *GetStationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationOK  %+v", 200, o.Payload)
}

func (o *GetStationOK) GetPayload() *models.Station {
	return o.Payload
}

func (o *GetStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Station)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationBadRequest creates a GetStationBadRequest with default headers values
func NewGetStationBadRequest() *GetStationBadRequest {
	return &GetStationBadRequest{}
}

/*
GetStationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetStationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station bad request response has a 2xx status code
func (o *GetStationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station bad request response has a 3xx status code
func (o *GetStationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station bad request response has a 4xx status code
func (o *GetStationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station bad request response has a 5xx status code
func (o *GetStationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get station bad request response a status code equal to that given
func (o *GetStationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetStationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationBadRequest  %+v", 400, o.Payload)
}

func (o *GetStationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationBadRequest  %+v", 400, o.Payload)
}

func (o *GetStationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationUnauthorized creates a GetStationUnauthorized with default headers values
func NewGetStationUnauthorized() *GetStationUnauthorized {
	return &GetStationUnauthorized{}
}

/*
GetStationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetStationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station unauthorized response has a 2xx status code
func (o *GetStationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station unauthorized response has a 3xx status code
func (o *GetStationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station unauthorized response has a 4xx status code
func (o *GetStationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station unauthorized response has a 5xx status code
func (o *GetStationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get station unauthorized response a status code equal to that given
func (o *GetStationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetStationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationForbidden creates a GetStationForbidden with default headers values
func NewGetStationForbidden() *GetStationForbidden {
	return &GetStationForbidden{}
}

/*
GetStationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetStationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station forbidden response has a 2xx status code
func (o *GetStationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station forbidden response has a 3xx status code
func (o *GetStationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station forbidden response has a 4xx status code
func (o *GetStationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station forbidden response has a 5xx status code
func (o *GetStationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get station forbidden response a status code equal to that given
func (o *GetStationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetStationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationForbidden  %+v", 403, o.Payload)
}

func (o *GetStationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationForbidden  %+v", 403, o.Payload)
}

func (o *GetStationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationNotFound creates a GetStationNotFound with default headers values
func NewGetStationNotFound() *GetStationNotFound {
	return &GetStationNotFound{}
}

/*
GetStationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetStationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station not found response has a 2xx status code
func (o *GetStationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station not found response has a 3xx status code
func (o *GetStationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station not found response has a 4xx status code
func (o *GetStationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station not found response has a 5xx status code
func (o *GetStationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get station not found response a status code equal to that given
func (o *GetStationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetStationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationNotFound  %+v", 404, o.Payload)
}

func (o *GetStationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationNotFound  %+v", 404, o.Payload)
}

func (o *GetStationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationRequestTimeout creates a GetStationRequestTimeout with default headers values
func NewGetStationRequestTimeout() *GetStationRequestTimeout {
	return &GetStationRequestTimeout{}
}

/*
GetStationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetStationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station request timeout response has a 2xx status code
func (o *GetStationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station request timeout response has a 3xx status code
func (o *GetStationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station request timeout response has a 4xx status code
func (o *GetStationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station request timeout response has a 5xx status code
func (o *GetStationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get station request timeout response a status code equal to that given
func (o *GetStationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetStationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetStationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetStationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationRequestEntityTooLarge creates a GetStationRequestEntityTooLarge with default headers values
func NewGetStationRequestEntityTooLarge() *GetStationRequestEntityTooLarge {
	return &GetStationRequestEntityTooLarge{}
}

/*
GetStationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetStationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station request entity too large response has a 2xx status code
func (o *GetStationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station request entity too large response has a 3xx status code
func (o *GetStationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station request entity too large response has a 4xx status code
func (o *GetStationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station request entity too large response has a 5xx status code
func (o *GetStationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get station request entity too large response a status code equal to that given
func (o *GetStationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetStationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetStationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetStationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationUnsupportedMediaType creates a GetStationUnsupportedMediaType with default headers values
func NewGetStationUnsupportedMediaType() *GetStationUnsupportedMediaType {
	return &GetStationUnsupportedMediaType{}
}

/*
GetStationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetStationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station unsupported media type response has a 2xx status code
func (o *GetStationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station unsupported media type response has a 3xx status code
func (o *GetStationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station unsupported media type response has a 4xx status code
func (o *GetStationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station unsupported media type response has a 5xx status code
func (o *GetStationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get station unsupported media type response a status code equal to that given
func (o *GetStationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetStationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetStationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetStationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationTooManyRequests creates a GetStationTooManyRequests with default headers values
func NewGetStationTooManyRequests() *GetStationTooManyRequests {
	return &GetStationTooManyRequests{}
}

/*
GetStationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetStationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station too many requests response has a 2xx status code
func (o *GetStationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station too many requests response has a 3xx status code
func (o *GetStationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station too many requests response has a 4xx status code
func (o *GetStationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get station too many requests response has a 5xx status code
func (o *GetStationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get station too many requests response a status code equal to that given
func (o *GetStationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetStationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationInternalServerError creates a GetStationInternalServerError with default headers values
func NewGetStationInternalServerError() *GetStationInternalServerError {
	return &GetStationInternalServerError{}
}

/*
GetStationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetStationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station internal server error response has a 2xx status code
func (o *GetStationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station internal server error response has a 3xx status code
func (o *GetStationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station internal server error response has a 4xx status code
func (o *GetStationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get station internal server error response has a 5xx status code
func (o *GetStationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get station internal server error response a status code equal to that given
func (o *GetStationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetStationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationServiceUnavailable creates a GetStationServiceUnavailable with default headers values
func NewGetStationServiceUnavailable() *GetStationServiceUnavailable {
	return &GetStationServiceUnavailable{}
}

/*
GetStationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetStationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station service unavailable response has a 2xx status code
func (o *GetStationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station service unavailable response has a 3xx status code
func (o *GetStationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station service unavailable response has a 4xx status code
func (o *GetStationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get station service unavailable response has a 5xx status code
func (o *GetStationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get station service unavailable response a status code equal to that given
func (o *GetStationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetStationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationGatewayTimeout creates a GetStationGatewayTimeout with default headers values
func NewGetStationGatewayTimeout() *GetStationGatewayTimeout {
	return &GetStationGatewayTimeout{}
}

/*
GetStationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetStationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get station gateway timeout response has a 2xx status code
func (o *GetStationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get station gateway timeout response has a 3xx status code
func (o *GetStationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get station gateway timeout response has a 4xx status code
func (o *GetStationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get station gateway timeout response has a 5xx status code
func (o *GetStationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get station gateway timeout response a status code equal to that given
func (o *GetStationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetStationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetStationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/stations/{stationId}][%d] getStationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetStationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
