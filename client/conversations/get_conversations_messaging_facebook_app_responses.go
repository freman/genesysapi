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

// GetConversationsMessagingFacebookAppReader is a Reader for the GetConversationsMessagingFacebookApp structure.
type GetConversationsMessagingFacebookAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingFacebookAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingFacebookAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingFacebookAppBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingFacebookAppUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingFacebookAppForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingFacebookAppNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingFacebookAppRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingFacebookAppUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingFacebookAppTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingFacebookAppInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingFacebookAppServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingFacebookAppGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingFacebookAppOK creates a GetConversationsMessagingFacebookAppOK with default headers values
func NewGetConversationsMessagingFacebookAppOK() *GetConversationsMessagingFacebookAppOK {
	return &GetConversationsMessagingFacebookAppOK{}
}

/*GetConversationsMessagingFacebookAppOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingFacebookAppOK struct {
	Payload *models.FacebookAppCredentials
}

func (o *GetConversationsMessagingFacebookAppOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppOK) GetPayload() *models.FacebookAppCredentials {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FacebookAppCredentials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppBadRequest creates a GetConversationsMessagingFacebookAppBadRequest with default headers values
func NewGetConversationsMessagingFacebookAppBadRequest() *GetConversationsMessagingFacebookAppBadRequest {
	return &GetConversationsMessagingFacebookAppBadRequest{}
}

/*GetConversationsMessagingFacebookAppBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingFacebookAppBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppUnauthorized creates a GetConversationsMessagingFacebookAppUnauthorized with default headers values
func NewGetConversationsMessagingFacebookAppUnauthorized() *GetConversationsMessagingFacebookAppUnauthorized {
	return &GetConversationsMessagingFacebookAppUnauthorized{}
}

/*GetConversationsMessagingFacebookAppUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingFacebookAppUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppForbidden creates a GetConversationsMessagingFacebookAppForbidden with default headers values
func NewGetConversationsMessagingFacebookAppForbidden() *GetConversationsMessagingFacebookAppForbidden {
	return &GetConversationsMessagingFacebookAppForbidden{}
}

/*GetConversationsMessagingFacebookAppForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingFacebookAppForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppNotFound creates a GetConversationsMessagingFacebookAppNotFound with default headers values
func NewGetConversationsMessagingFacebookAppNotFound() *GetConversationsMessagingFacebookAppNotFound {
	return &GetConversationsMessagingFacebookAppNotFound{}
}

/*GetConversationsMessagingFacebookAppNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingFacebookAppNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppRequestEntityTooLarge creates a GetConversationsMessagingFacebookAppRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingFacebookAppRequestEntityTooLarge() *GetConversationsMessagingFacebookAppRequestEntityTooLarge {
	return &GetConversationsMessagingFacebookAppRequestEntityTooLarge{}
}

/*GetConversationsMessagingFacebookAppRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessagingFacebookAppRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppUnsupportedMediaType creates a GetConversationsMessagingFacebookAppUnsupportedMediaType with default headers values
func NewGetConversationsMessagingFacebookAppUnsupportedMediaType() *GetConversationsMessagingFacebookAppUnsupportedMediaType {
	return &GetConversationsMessagingFacebookAppUnsupportedMediaType{}
}

/*GetConversationsMessagingFacebookAppUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingFacebookAppUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppTooManyRequests creates a GetConversationsMessagingFacebookAppTooManyRequests with default headers values
func NewGetConversationsMessagingFacebookAppTooManyRequests() *GetConversationsMessagingFacebookAppTooManyRequests {
	return &GetConversationsMessagingFacebookAppTooManyRequests{}
}

/*GetConversationsMessagingFacebookAppTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsMessagingFacebookAppTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppInternalServerError creates a GetConversationsMessagingFacebookAppInternalServerError with default headers values
func NewGetConversationsMessagingFacebookAppInternalServerError() *GetConversationsMessagingFacebookAppInternalServerError {
	return &GetConversationsMessagingFacebookAppInternalServerError{}
}

/*GetConversationsMessagingFacebookAppInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingFacebookAppInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppServiceUnavailable creates a GetConversationsMessagingFacebookAppServiceUnavailable with default headers values
func NewGetConversationsMessagingFacebookAppServiceUnavailable() *GetConversationsMessagingFacebookAppServiceUnavailable {
	return &GetConversationsMessagingFacebookAppServiceUnavailable{}
}

/*GetConversationsMessagingFacebookAppServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingFacebookAppServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingFacebookAppGatewayTimeout creates a GetConversationsMessagingFacebookAppGatewayTimeout with default headers values
func NewGetConversationsMessagingFacebookAppGatewayTimeout() *GetConversationsMessagingFacebookAppGatewayTimeout {
	return &GetConversationsMessagingFacebookAppGatewayTimeout{}
}

/*GetConversationsMessagingFacebookAppGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingFacebookAppGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingFacebookAppGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/facebook/app][%d] getConversationsMessagingFacebookAppGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingFacebookAppGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingFacebookAppGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
