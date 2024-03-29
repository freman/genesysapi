// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetNotificationsChannelSubscriptionsReader is a Reader for the GetNotificationsChannelSubscriptions structure.
type GetNotificationsChannelSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotificationsChannelSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNotificationsChannelSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNotificationsChannelSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNotificationsChannelSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNotificationsChannelSubscriptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNotificationsChannelSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetNotificationsChannelSubscriptionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetNotificationsChannelSubscriptionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetNotificationsChannelSubscriptionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetNotificationsChannelSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNotificationsChannelSubscriptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetNotificationsChannelSubscriptionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetNotificationsChannelSubscriptionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNotificationsChannelSubscriptionsOK creates a GetNotificationsChannelSubscriptionsOK with default headers values
func NewGetNotificationsChannelSubscriptionsOK() *GetNotificationsChannelSubscriptionsOK {
	return &GetNotificationsChannelSubscriptionsOK{}
}

/*
GetNotificationsChannelSubscriptionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNotificationsChannelSubscriptionsOK struct {
	Payload *models.ChannelTopicEntityListing
}

// IsSuccess returns true when this get notifications channel subscriptions o k response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get notifications channel subscriptions o k response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions o k response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications channel subscriptions o k response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions o k response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNotificationsChannelSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsOK) GetPayload() *models.ChannelTopicEntityListing {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelTopicEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsBadRequest creates a GetNotificationsChannelSubscriptionsBadRequest with default headers values
func NewGetNotificationsChannelSubscriptionsBadRequest() *GetNotificationsChannelSubscriptionsBadRequest {
	return &GetNotificationsChannelSubscriptionsBadRequest{}
}

/*
GetNotificationsChannelSubscriptionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetNotificationsChannelSubscriptionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions bad request response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions bad request response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions bad request response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions bad request response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions bad request response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetNotificationsChannelSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsUnauthorized creates a GetNotificationsChannelSubscriptionsUnauthorized with default headers values
func NewGetNotificationsChannelSubscriptionsUnauthorized() *GetNotificationsChannelSubscriptionsUnauthorized {
	return &GetNotificationsChannelSubscriptionsUnauthorized{}
}

/*
GetNotificationsChannelSubscriptionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetNotificationsChannelSubscriptionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions unauthorized response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions unauthorized response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions unauthorized response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions unauthorized response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions unauthorized response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetNotificationsChannelSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsForbidden creates a GetNotificationsChannelSubscriptionsForbidden with default headers values
func NewGetNotificationsChannelSubscriptionsForbidden() *GetNotificationsChannelSubscriptionsForbidden {
	return &GetNotificationsChannelSubscriptionsForbidden{}
}

/*
GetNotificationsChannelSubscriptionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetNotificationsChannelSubscriptionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions forbidden response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions forbidden response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions forbidden response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions forbidden response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions forbidden response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetNotificationsChannelSubscriptionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsForbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsForbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsNotFound creates a GetNotificationsChannelSubscriptionsNotFound with default headers values
func NewGetNotificationsChannelSubscriptionsNotFound() *GetNotificationsChannelSubscriptionsNotFound {
	return &GetNotificationsChannelSubscriptionsNotFound{}
}

/*
GetNotificationsChannelSubscriptionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetNotificationsChannelSubscriptionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions not found response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions not found response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions not found response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions not found response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions not found response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetNotificationsChannelSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsNotFound  %+v", 404, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsNotFound  %+v", 404, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsRequestTimeout creates a GetNotificationsChannelSubscriptionsRequestTimeout with default headers values
func NewGetNotificationsChannelSubscriptionsRequestTimeout() *GetNotificationsChannelSubscriptionsRequestTimeout {
	return &GetNotificationsChannelSubscriptionsRequestTimeout{}
}

/*
GetNotificationsChannelSubscriptionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetNotificationsChannelSubscriptionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions request timeout response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions request timeout response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions request timeout response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions request timeout response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions request timeout response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetNotificationsChannelSubscriptionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsRequestEntityTooLarge creates a GetNotificationsChannelSubscriptionsRequestEntityTooLarge with default headers values
func NewGetNotificationsChannelSubscriptionsRequestEntityTooLarge() *GetNotificationsChannelSubscriptionsRequestEntityTooLarge {
	return &GetNotificationsChannelSubscriptionsRequestEntityTooLarge{}
}

/*
GetNotificationsChannelSubscriptionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetNotificationsChannelSubscriptionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions request entity too large response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions request entity too large response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions request entity too large response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions request entity too large response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions request entity too large response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsUnsupportedMediaType creates a GetNotificationsChannelSubscriptionsUnsupportedMediaType with default headers values
func NewGetNotificationsChannelSubscriptionsUnsupportedMediaType() *GetNotificationsChannelSubscriptionsUnsupportedMediaType {
	return &GetNotificationsChannelSubscriptionsUnsupportedMediaType{}
}

/*
GetNotificationsChannelSubscriptionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetNotificationsChannelSubscriptionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions unsupported media type response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions unsupported media type response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions unsupported media type response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions unsupported media type response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions unsupported media type response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsTooManyRequests creates a GetNotificationsChannelSubscriptionsTooManyRequests with default headers values
func NewGetNotificationsChannelSubscriptionsTooManyRequests() *GetNotificationsChannelSubscriptionsTooManyRequests {
	return &GetNotificationsChannelSubscriptionsTooManyRequests{}
}

/*
GetNotificationsChannelSubscriptionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetNotificationsChannelSubscriptionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions too many requests response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions too many requests response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions too many requests response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get notifications channel subscriptions too many requests response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get notifications channel subscriptions too many requests response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetNotificationsChannelSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsInternalServerError creates a GetNotificationsChannelSubscriptionsInternalServerError with default headers values
func NewGetNotificationsChannelSubscriptionsInternalServerError() *GetNotificationsChannelSubscriptionsInternalServerError {
	return &GetNotificationsChannelSubscriptionsInternalServerError{}
}

/*
GetNotificationsChannelSubscriptionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetNotificationsChannelSubscriptionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions internal server error response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions internal server error response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions internal server error response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications channel subscriptions internal server error response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications channel subscriptions internal server error response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetNotificationsChannelSubscriptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsServiceUnavailable creates a GetNotificationsChannelSubscriptionsServiceUnavailable with default headers values
func NewGetNotificationsChannelSubscriptionsServiceUnavailable() *GetNotificationsChannelSubscriptionsServiceUnavailable {
	return &GetNotificationsChannelSubscriptionsServiceUnavailable{}
}

/*
GetNotificationsChannelSubscriptionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetNotificationsChannelSubscriptionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions service unavailable response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions service unavailable response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions service unavailable response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications channel subscriptions service unavailable response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications channel subscriptions service unavailable response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsChannelSubscriptionsGatewayTimeout creates a GetNotificationsChannelSubscriptionsGatewayTimeout with default headers values
func NewGetNotificationsChannelSubscriptionsGatewayTimeout() *GetNotificationsChannelSubscriptionsGatewayTimeout {
	return &GetNotificationsChannelSubscriptionsGatewayTimeout{}
}

/*
GetNotificationsChannelSubscriptionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetNotificationsChannelSubscriptionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get notifications channel subscriptions gateway timeout response has a 2xx status code
func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get notifications channel subscriptions gateway timeout response has a 3xx status code
func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get notifications channel subscriptions gateway timeout response has a 4xx status code
func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get notifications channel subscriptions gateway timeout response has a 5xx status code
func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get notifications channel subscriptions gateway timeout response a status code equal to that given
func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/notifications/channels/{channelId}/subscriptions][%d] getNotificationsChannelSubscriptionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsChannelSubscriptionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
