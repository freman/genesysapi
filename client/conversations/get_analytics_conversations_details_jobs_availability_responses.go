// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAnalyticsConversationsDetailsJobsAvailabilityReader is a Reader for the GetAnalyticsConversationsDetailsJobsAvailability structure.
type GetAnalyticsConversationsDetailsJobsAvailabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityOK creates a GetAnalyticsConversationsDetailsJobsAvailabilityOK with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityOK() *GetAnalyticsConversationsDetailsJobsAvailabilityOK {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityOK{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityOK handles this case with default header values.

successful operation
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityOK struct {
	Payload *models.DataAvailabilityResponse
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) GetPayload() *models.DataAvailabilityResponse {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataAvailabilityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityBadRequest creates a GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityBadRequest() *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized creates a GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized() *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityForbidden creates a GetAnalyticsConversationsDetailsJobsAvailabilityForbidden with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityForbidden() *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityForbidden{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityNotFound creates a GetAnalyticsConversationsDetailsJobsAvailabilityNotFound with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityNotFound() *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityNotFound{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge creates a GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge() *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType creates a GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType() *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests creates a GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests() *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError creates a GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError() *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable creates a GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable() *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout creates a GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout() *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout{}
}

/*GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
