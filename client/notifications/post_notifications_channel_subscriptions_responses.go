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

// PostNotificationsChannelSubscriptionsReader is a Reader for the PostNotificationsChannelSubscriptions structure.
type PostNotificationsChannelSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNotificationsChannelSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNotificationsChannelSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostNotificationsChannelSubscriptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostNotificationsChannelSubscriptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostNotificationsChannelSubscriptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostNotificationsChannelSubscriptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostNotificationsChannelSubscriptionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostNotificationsChannelSubscriptionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostNotificationsChannelSubscriptionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostNotificationsChannelSubscriptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostNotificationsChannelSubscriptionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostNotificationsChannelSubscriptionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostNotificationsChannelSubscriptionsOK creates a PostNotificationsChannelSubscriptionsOK with default headers values
func NewPostNotificationsChannelSubscriptionsOK() *PostNotificationsChannelSubscriptionsOK {
	return &PostNotificationsChannelSubscriptionsOK{}
}

/*PostNotificationsChannelSubscriptionsOK handles this case with default header values.

successful operation
*/
type PostNotificationsChannelSubscriptionsOK struct {
	Payload *models.ChannelTopicEntityListing
}

func (o *PostNotificationsChannelSubscriptionsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsOK) GetPayload() *models.ChannelTopicEntityListing {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelTopicEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsBadRequest creates a PostNotificationsChannelSubscriptionsBadRequest with default headers values
func NewPostNotificationsChannelSubscriptionsBadRequest() *PostNotificationsChannelSubscriptionsBadRequest {
	return &PostNotificationsChannelSubscriptionsBadRequest{}
}

/*PostNotificationsChannelSubscriptionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostNotificationsChannelSubscriptionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsBadRequest  %+v", 400, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsUnauthorized creates a PostNotificationsChannelSubscriptionsUnauthorized with default headers values
func NewPostNotificationsChannelSubscriptionsUnauthorized() *PostNotificationsChannelSubscriptionsUnauthorized {
	return &PostNotificationsChannelSubscriptionsUnauthorized{}
}

/*PostNotificationsChannelSubscriptionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostNotificationsChannelSubscriptionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsForbidden creates a PostNotificationsChannelSubscriptionsForbidden with default headers values
func NewPostNotificationsChannelSubscriptionsForbidden() *PostNotificationsChannelSubscriptionsForbidden {
	return &PostNotificationsChannelSubscriptionsForbidden{}
}

/*PostNotificationsChannelSubscriptionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostNotificationsChannelSubscriptionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsForbidden  %+v", 403, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsNotFound creates a PostNotificationsChannelSubscriptionsNotFound with default headers values
func NewPostNotificationsChannelSubscriptionsNotFound() *PostNotificationsChannelSubscriptionsNotFound {
	return &PostNotificationsChannelSubscriptionsNotFound{}
}

/*PostNotificationsChannelSubscriptionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostNotificationsChannelSubscriptionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsNotFound  %+v", 404, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsRequestEntityTooLarge creates a PostNotificationsChannelSubscriptionsRequestEntityTooLarge with default headers values
func NewPostNotificationsChannelSubscriptionsRequestEntityTooLarge() *PostNotificationsChannelSubscriptionsRequestEntityTooLarge {
	return &PostNotificationsChannelSubscriptionsRequestEntityTooLarge{}
}

/*PostNotificationsChannelSubscriptionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostNotificationsChannelSubscriptionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsUnsupportedMediaType creates a PostNotificationsChannelSubscriptionsUnsupportedMediaType with default headers values
func NewPostNotificationsChannelSubscriptionsUnsupportedMediaType() *PostNotificationsChannelSubscriptionsUnsupportedMediaType {
	return &PostNotificationsChannelSubscriptionsUnsupportedMediaType{}
}

/*PostNotificationsChannelSubscriptionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostNotificationsChannelSubscriptionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsTooManyRequests creates a PostNotificationsChannelSubscriptionsTooManyRequests with default headers values
func NewPostNotificationsChannelSubscriptionsTooManyRequests() *PostNotificationsChannelSubscriptionsTooManyRequests {
	return &PostNotificationsChannelSubscriptionsTooManyRequests{}
}

/*PostNotificationsChannelSubscriptionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostNotificationsChannelSubscriptionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsInternalServerError creates a PostNotificationsChannelSubscriptionsInternalServerError with default headers values
func NewPostNotificationsChannelSubscriptionsInternalServerError() *PostNotificationsChannelSubscriptionsInternalServerError {
	return &PostNotificationsChannelSubscriptionsInternalServerError{}
}

/*PostNotificationsChannelSubscriptionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostNotificationsChannelSubscriptionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsServiceUnavailable creates a PostNotificationsChannelSubscriptionsServiceUnavailable with default headers values
func NewPostNotificationsChannelSubscriptionsServiceUnavailable() *PostNotificationsChannelSubscriptionsServiceUnavailable {
	return &PostNotificationsChannelSubscriptionsServiceUnavailable{}
}

/*PostNotificationsChannelSubscriptionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostNotificationsChannelSubscriptionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNotificationsChannelSubscriptionsGatewayTimeout creates a PostNotificationsChannelSubscriptionsGatewayTimeout with default headers values
func NewPostNotificationsChannelSubscriptionsGatewayTimeout() *PostNotificationsChannelSubscriptionsGatewayTimeout {
	return &PostNotificationsChannelSubscriptionsGatewayTimeout{}
}

/*PostNotificationsChannelSubscriptionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostNotificationsChannelSubscriptionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostNotificationsChannelSubscriptionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/notifications/channels/{channelId}/subscriptions][%d] postNotificationsChannelSubscriptionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostNotificationsChannelSubscriptionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostNotificationsChannelSubscriptionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
