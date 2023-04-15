// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLocationReader is a Reader for the GetLocation structure.
type GetLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLocationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLocationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLocationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLocationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLocationOK creates a GetLocationOK with default headers values
func NewGetLocationOK() *GetLocationOK {
	return &GetLocationOK{}
}

/*
GetLocationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLocationOK struct {
	Payload *models.LocationDefinition
}

// IsSuccess returns true when this get location o k response has a 2xx status code
func (o *GetLocationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get location o k response has a 3xx status code
func (o *GetLocationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location o k response has a 4xx status code
func (o *GetLocationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get location o k response has a 5xx status code
func (o *GetLocationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get location o k response a status code equal to that given
func (o *GetLocationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLocationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationOK  %+v", 200, o.Payload)
}

func (o *GetLocationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationOK  %+v", 200, o.Payload)
}

func (o *GetLocationOK) GetPayload() *models.LocationDefinition {
	return o.Payload
}

func (o *GetLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationBadRequest creates a GetLocationBadRequest with default headers values
func NewGetLocationBadRequest() *GetLocationBadRequest {
	return &GetLocationBadRequest{}
}

/*
GetLocationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLocationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location bad request response has a 2xx status code
func (o *GetLocationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location bad request response has a 3xx status code
func (o *GetLocationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location bad request response has a 4xx status code
func (o *GetLocationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location bad request response has a 5xx status code
func (o *GetLocationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get location bad request response a status code equal to that given
func (o *GetLocationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLocationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationBadRequest  %+v", 400, o.Payload)
}

func (o *GetLocationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationUnauthorized creates a GetLocationUnauthorized with default headers values
func NewGetLocationUnauthorized() *GetLocationUnauthorized {
	return &GetLocationUnauthorized{}
}

/*
GetLocationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLocationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location unauthorized response has a 2xx status code
func (o *GetLocationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location unauthorized response has a 3xx status code
func (o *GetLocationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location unauthorized response has a 4xx status code
func (o *GetLocationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location unauthorized response has a 5xx status code
func (o *GetLocationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get location unauthorized response a status code equal to that given
func (o *GetLocationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLocationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationForbidden creates a GetLocationForbidden with default headers values
func NewGetLocationForbidden() *GetLocationForbidden {
	return &GetLocationForbidden{}
}

/*
GetLocationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLocationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location forbidden response has a 2xx status code
func (o *GetLocationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location forbidden response has a 3xx status code
func (o *GetLocationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location forbidden response has a 4xx status code
func (o *GetLocationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location forbidden response has a 5xx status code
func (o *GetLocationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get location forbidden response a status code equal to that given
func (o *GetLocationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLocationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationForbidden  %+v", 403, o.Payload)
}

func (o *GetLocationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationNotFound creates a GetLocationNotFound with default headers values
func NewGetLocationNotFound() *GetLocationNotFound {
	return &GetLocationNotFound{}
}

/*
GetLocationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLocationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location not found response has a 2xx status code
func (o *GetLocationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location not found response has a 3xx status code
func (o *GetLocationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location not found response has a 4xx status code
func (o *GetLocationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location not found response has a 5xx status code
func (o *GetLocationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get location not found response a status code equal to that given
func (o *GetLocationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLocationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationNotFound  %+v", 404, o.Payload)
}

func (o *GetLocationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationNotFound  %+v", 404, o.Payload)
}

func (o *GetLocationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationRequestTimeout creates a GetLocationRequestTimeout with default headers values
func NewGetLocationRequestTimeout() *GetLocationRequestTimeout {
	return &GetLocationRequestTimeout{}
}

/*
GetLocationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLocationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location request timeout response has a 2xx status code
func (o *GetLocationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location request timeout response has a 3xx status code
func (o *GetLocationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location request timeout response has a 4xx status code
func (o *GetLocationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location request timeout response has a 5xx status code
func (o *GetLocationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get location request timeout response a status code equal to that given
func (o *GetLocationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLocationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLocationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLocationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationRequestEntityTooLarge creates a GetLocationRequestEntityTooLarge with default headers values
func NewGetLocationRequestEntityTooLarge() *GetLocationRequestEntityTooLarge {
	return &GetLocationRequestEntityTooLarge{}
}

/*
GetLocationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLocationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location request entity too large response has a 2xx status code
func (o *GetLocationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location request entity too large response has a 3xx status code
func (o *GetLocationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location request entity too large response has a 4xx status code
func (o *GetLocationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location request entity too large response has a 5xx status code
func (o *GetLocationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get location request entity too large response a status code equal to that given
func (o *GetLocationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLocationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLocationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLocationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationUnsupportedMediaType creates a GetLocationUnsupportedMediaType with default headers values
func NewGetLocationUnsupportedMediaType() *GetLocationUnsupportedMediaType {
	return &GetLocationUnsupportedMediaType{}
}

/*
GetLocationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLocationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location unsupported media type response has a 2xx status code
func (o *GetLocationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location unsupported media type response has a 3xx status code
func (o *GetLocationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location unsupported media type response has a 4xx status code
func (o *GetLocationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location unsupported media type response has a 5xx status code
func (o *GetLocationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get location unsupported media type response a status code equal to that given
func (o *GetLocationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLocationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLocationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLocationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationTooManyRequests creates a GetLocationTooManyRequests with default headers values
func NewGetLocationTooManyRequests() *GetLocationTooManyRequests {
	return &GetLocationTooManyRequests{}
}

/*
GetLocationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLocationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location too many requests response has a 2xx status code
func (o *GetLocationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location too many requests response has a 3xx status code
func (o *GetLocationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location too many requests response has a 4xx status code
func (o *GetLocationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get location too many requests response has a 5xx status code
func (o *GetLocationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get location too many requests response a status code equal to that given
func (o *GetLocationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLocationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLocationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationInternalServerError creates a GetLocationInternalServerError with default headers values
func NewGetLocationInternalServerError() *GetLocationInternalServerError {
	return &GetLocationInternalServerError{}
}

/*
GetLocationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLocationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location internal server error response has a 2xx status code
func (o *GetLocationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location internal server error response has a 3xx status code
func (o *GetLocationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location internal server error response has a 4xx status code
func (o *GetLocationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get location internal server error response has a 5xx status code
func (o *GetLocationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get location internal server error response a status code equal to that given
func (o *GetLocationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLocationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLocationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationServiceUnavailable creates a GetLocationServiceUnavailable with default headers values
func NewGetLocationServiceUnavailable() *GetLocationServiceUnavailable {
	return &GetLocationServiceUnavailable{}
}

/*
GetLocationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLocationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location service unavailable response has a 2xx status code
func (o *GetLocationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location service unavailable response has a 3xx status code
func (o *GetLocationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location service unavailable response has a 4xx status code
func (o *GetLocationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get location service unavailable response has a 5xx status code
func (o *GetLocationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get location service unavailable response a status code equal to that given
func (o *GetLocationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLocationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLocationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationGatewayTimeout creates a GetLocationGatewayTimeout with default headers values
func NewGetLocationGatewayTimeout() *GetLocationGatewayTimeout {
	return &GetLocationGatewayTimeout{}
}

/*
GetLocationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLocationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get location gateway timeout response has a 2xx status code
func (o *GetLocationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get location gateway timeout response has a 3xx status code
func (o *GetLocationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get location gateway timeout response has a 4xx status code
func (o *GetLocationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get location gateway timeout response has a 5xx status code
func (o *GetLocationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get location gateway timeout response a status code equal to that given
func (o *GetLocationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLocationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/locations/{locationId}][%d] getLocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLocationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
