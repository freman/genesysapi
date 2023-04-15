// Code generated by go-swagger; DO NOT EDIT.

package analytics

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
	case 408:
		result := NewGetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout()
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityOK struct {
	Payload *models.DataAvailabilityResponse
}

// IsSuccess returns true when this get analytics conversations details jobs availability o k response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get analytics conversations details jobs availability o k response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability o k response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics conversations details jobs availability o k response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability o k response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityOK) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability bad request response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability bad request response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability bad request response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability bad request response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability bad request response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityBadRequest) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability unauthorized response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability unauthorized response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability unauthorized response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability unauthorized response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability unauthorized response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnauthorized) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability forbidden response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability forbidden response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability forbidden response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability forbidden response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability forbidden response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityForbidden) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability not found response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability not found response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability not found response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability not found response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability not found response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityNotFound) String() string {
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

// NewGetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout creates a GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout with default headers values
func NewGetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout() *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout {
	return &GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout{}
}

/*
GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability request timeout response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability request timeout response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability request timeout response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability request timeout response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability request timeout response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability request entity too large response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability request entity too large response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability request entity too large response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability request entity too large response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability request entity too large response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityRequestEntityTooLarge) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability unsupported media type response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability unsupported media type response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability unsupported media type response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability unsupported media type response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability unsupported media type response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityUnsupportedMediaType) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability too many requests response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability too many requests response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability too many requests response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics conversations details jobs availability too many requests response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics conversations details jobs availability too many requests response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityTooManyRequests) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability internal server error response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability internal server error response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability internal server error response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics conversations details jobs availability internal server error response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics conversations details jobs availability internal server error response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityInternalServerError) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability service unavailable response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability service unavailable response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability service unavailable response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics conversations details jobs availability service unavailable response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics conversations details jobs availability service unavailable response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityServiceUnavailable) String() string {
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

/*
GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics conversations details jobs availability gateway timeout response has a 2xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics conversations details jobs availability gateway timeout response has a 3xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics conversations details jobs availability gateway timeout response has a 4xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics conversations details jobs availability gateway timeout response has a 5xx status code
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics conversations details jobs availability gateway timeout response a status code equal to that given
func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/conversations/details/jobs/availability][%d] getAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsConversationsDetailsJobsAvailabilityGatewayTimeout) String() string {
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
