// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsUsersDetailsJobsAvailabilityReader is a Reader for the GetAnalyticsUsersDetailsJobsAvailability structure.
type GetAnalyticsUsersDetailsJobsAvailabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsUsersDetailsJobsAvailabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityOK creates a GetAnalyticsUsersDetailsJobsAvailabilityOK with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityOK() *GetAnalyticsUsersDetailsJobsAvailabilityOK {
	return &GetAnalyticsUsersDetailsJobsAvailabilityOK{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityOK handles this case with default header values.

successful operation
*/
type GetAnalyticsUsersDetailsJobsAvailabilityOK struct {
	Payload *models.DataAvailabilityResponse
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityOK) GetPayload() *models.DataAvailabilityResponse {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataAvailabilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityBadRequest creates a GetAnalyticsUsersDetailsJobsAvailabilityBadRequest with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityBadRequest() *GetAnalyticsUsersDetailsJobsAvailabilityBadRequest {
	return &GetAnalyticsUsersDetailsJobsAvailabilityBadRequest{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityUnauthorized creates a GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityUnauthorized() *GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized {
	return &GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityForbidden creates a GetAnalyticsUsersDetailsJobsAvailabilityForbidden with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityForbidden() *GetAnalyticsUsersDetailsJobsAvailabilityForbidden {
	return &GetAnalyticsUsersDetailsJobsAvailabilityForbidden{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityNotFound creates a GetAnalyticsUsersDetailsJobsAvailabilityNotFound with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityNotFound() *GetAnalyticsUsersDetailsJobsAvailabilityNotFound {
	return &GetAnalyticsUsersDetailsJobsAvailabilityNotFound{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout creates a GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout() *GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout {
	return &GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge creates a GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge() *GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge {
	return &GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType creates a GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType() *GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType {
	return &GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests creates a GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests() *GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests {
	return &GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityInternalServerError creates a GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityInternalServerError() *GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError {
	return &GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable creates a GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable() *GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable {
	return &GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout creates a GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout with default headers values
func NewGetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout() *GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout {
	return &GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout{}
}

/*GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/users/details/jobs/availability][%d] getAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsUsersDetailsJobsAvailabilityGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
