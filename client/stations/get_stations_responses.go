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

// GetStationsReader is a Reader for the GetStations structure.
type GetStationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetStationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetStationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetStationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetStationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetStationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStationsOK creates a GetStationsOK with default headers values
func NewGetStationsOK() *GetStationsOK {
	return &GetStationsOK{}
}

/*
GetStationsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetStationsOK struct {
	Payload *models.StationEntityListing
}

// IsSuccess returns true when this get stations o k response has a 2xx status code
func (o *GetStationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get stations o k response has a 3xx status code
func (o *GetStationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations o k response has a 4xx status code
func (o *GetStationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations o k response has a 5xx status code
func (o *GetStationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations o k response a status code equal to that given
func (o *GetStationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetStationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsOK  %+v", 200, o.Payload)
}

func (o *GetStationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsOK  %+v", 200, o.Payload)
}

func (o *GetStationsOK) GetPayload() *models.StationEntityListing {
	return o.Payload
}

func (o *GetStationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsBadRequest creates a GetStationsBadRequest with default headers values
func NewGetStationsBadRequest() *GetStationsBadRequest {
	return &GetStationsBadRequest{}
}

/*
GetStationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetStationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations bad request response has a 2xx status code
func (o *GetStationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations bad request response has a 3xx status code
func (o *GetStationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations bad request response has a 4xx status code
func (o *GetStationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations bad request response has a 5xx status code
func (o *GetStationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations bad request response a status code equal to that given
func (o *GetStationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetStationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetStationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetStationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsUnauthorized creates a GetStationsUnauthorized with default headers values
func NewGetStationsUnauthorized() *GetStationsUnauthorized {
	return &GetStationsUnauthorized{}
}

/*
GetStationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetStationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations unauthorized response has a 2xx status code
func (o *GetStationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations unauthorized response has a 3xx status code
func (o *GetStationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations unauthorized response has a 4xx status code
func (o *GetStationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations unauthorized response has a 5xx status code
func (o *GetStationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations unauthorized response a status code equal to that given
func (o *GetStationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetStationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsForbidden creates a GetStationsForbidden with default headers values
func NewGetStationsForbidden() *GetStationsForbidden {
	return &GetStationsForbidden{}
}

/*
GetStationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetStationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations forbidden response has a 2xx status code
func (o *GetStationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations forbidden response has a 3xx status code
func (o *GetStationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations forbidden response has a 4xx status code
func (o *GetStationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations forbidden response has a 5xx status code
func (o *GetStationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations forbidden response a status code equal to that given
func (o *GetStationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetStationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsForbidden  %+v", 403, o.Payload)
}

func (o *GetStationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsForbidden  %+v", 403, o.Payload)
}

func (o *GetStationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsNotFound creates a GetStationsNotFound with default headers values
func NewGetStationsNotFound() *GetStationsNotFound {
	return &GetStationsNotFound{}
}

/*
GetStationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetStationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations not found response has a 2xx status code
func (o *GetStationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations not found response has a 3xx status code
func (o *GetStationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations not found response has a 4xx status code
func (o *GetStationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations not found response has a 5xx status code
func (o *GetStationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations not found response a status code equal to that given
func (o *GetStationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetStationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsNotFound  %+v", 404, o.Payload)
}

func (o *GetStationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsNotFound  %+v", 404, o.Payload)
}

func (o *GetStationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsRequestTimeout creates a GetStationsRequestTimeout with default headers values
func NewGetStationsRequestTimeout() *GetStationsRequestTimeout {
	return &GetStationsRequestTimeout{}
}

/*
GetStationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetStationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations request timeout response has a 2xx status code
func (o *GetStationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations request timeout response has a 3xx status code
func (o *GetStationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations request timeout response has a 4xx status code
func (o *GetStationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations request timeout response has a 5xx status code
func (o *GetStationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations request timeout response a status code equal to that given
func (o *GetStationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetStationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetStationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetStationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsRequestEntityTooLarge creates a GetStationsRequestEntityTooLarge with default headers values
func NewGetStationsRequestEntityTooLarge() *GetStationsRequestEntityTooLarge {
	return &GetStationsRequestEntityTooLarge{}
}

/*
GetStationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetStationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations request entity too large response has a 2xx status code
func (o *GetStationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations request entity too large response has a 3xx status code
func (o *GetStationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations request entity too large response has a 4xx status code
func (o *GetStationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations request entity too large response has a 5xx status code
func (o *GetStationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations request entity too large response a status code equal to that given
func (o *GetStationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetStationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetStationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetStationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsUnsupportedMediaType creates a GetStationsUnsupportedMediaType with default headers values
func NewGetStationsUnsupportedMediaType() *GetStationsUnsupportedMediaType {
	return &GetStationsUnsupportedMediaType{}
}

/*
GetStationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetStationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations unsupported media type response has a 2xx status code
func (o *GetStationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations unsupported media type response has a 3xx status code
func (o *GetStationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations unsupported media type response has a 4xx status code
func (o *GetStationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations unsupported media type response has a 5xx status code
func (o *GetStationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations unsupported media type response a status code equal to that given
func (o *GetStationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetStationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetStationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetStationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsTooManyRequests creates a GetStationsTooManyRequests with default headers values
func NewGetStationsTooManyRequests() *GetStationsTooManyRequests {
	return &GetStationsTooManyRequests{}
}

/*
GetStationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetStationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations too many requests response has a 2xx status code
func (o *GetStationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations too many requests response has a 3xx status code
func (o *GetStationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations too many requests response has a 4xx status code
func (o *GetStationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations too many requests response has a 5xx status code
func (o *GetStationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations too many requests response a status code equal to that given
func (o *GetStationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetStationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsInternalServerError creates a GetStationsInternalServerError with default headers values
func NewGetStationsInternalServerError() *GetStationsInternalServerError {
	return &GetStationsInternalServerError{}
}

/*
GetStationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetStationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations internal server error response has a 2xx status code
func (o *GetStationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations internal server error response has a 3xx status code
func (o *GetStationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations internal server error response has a 4xx status code
func (o *GetStationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations internal server error response has a 5xx status code
func (o *GetStationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get stations internal server error response a status code equal to that given
func (o *GetStationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetStationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsServiceUnavailable creates a GetStationsServiceUnavailable with default headers values
func NewGetStationsServiceUnavailable() *GetStationsServiceUnavailable {
	return &GetStationsServiceUnavailable{}
}

/*
GetStationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetStationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations service unavailable response has a 2xx status code
func (o *GetStationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations service unavailable response has a 3xx status code
func (o *GetStationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations service unavailable response has a 4xx status code
func (o *GetStationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations service unavailable response has a 5xx status code
func (o *GetStationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get stations service unavailable response a status code equal to that given
func (o *GetStationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetStationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsGatewayTimeout creates a GetStationsGatewayTimeout with default headers values
func NewGetStationsGatewayTimeout() *GetStationsGatewayTimeout {
	return &GetStationsGatewayTimeout{}
}

/*
GetStationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetStationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations gateway timeout response has a 2xx status code
func (o *GetStationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations gateway timeout response has a 3xx status code
func (o *GetStationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations gateway timeout response has a 4xx status code
func (o *GetStationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations gateway timeout response has a 5xx status code
func (o *GetStationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get stations gateway timeout response a status code equal to that given
func (o *GetStationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetStationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetStationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/stations][%d] getStationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetStationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
