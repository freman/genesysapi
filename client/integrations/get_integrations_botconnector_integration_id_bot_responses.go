// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIntegrationsBotconnectorIntegrationIDBotReader is a Reader for the GetIntegrationsBotconnectorIntegrationIDBot structure.
type GetIntegrationsBotconnectorIntegrationIDBotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsBotconnectorIntegrationIDBotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsBotconnectorIntegrationIDBotOK creates a GetIntegrationsBotconnectorIntegrationIDBotOK with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotOK() *GetIntegrationsBotconnectorIntegrationIDBotOK {
	return &GetIntegrationsBotconnectorIntegrationIDBotOK{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotOK handles this case with default header values.

successful operation
*/
type GetIntegrationsBotconnectorIntegrationIDBotOK struct {
	Payload *models.BotConnectorBot
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotOK) GetPayload() *models.BotConnectorBot {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BotConnectorBot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotBadRequest creates a GetIntegrationsBotconnectorIntegrationIDBotBadRequest with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotBadRequest() *GetIntegrationsBotconnectorIntegrationIDBotBadRequest {
	return &GetIntegrationsBotconnectorIntegrationIDBotBadRequest{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsBotconnectorIntegrationIDBotBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotUnauthorized creates a GetIntegrationsBotconnectorIntegrationIDBotUnauthorized with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotUnauthorized() *GetIntegrationsBotconnectorIntegrationIDBotUnauthorized {
	return &GetIntegrationsBotconnectorIntegrationIDBotUnauthorized{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsBotconnectorIntegrationIDBotUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotForbidden creates a GetIntegrationsBotconnectorIntegrationIDBotForbidden with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotForbidden() *GetIntegrationsBotconnectorIntegrationIDBotForbidden {
	return &GetIntegrationsBotconnectorIntegrationIDBotForbidden{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsBotconnectorIntegrationIDBotForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotNotFound creates a GetIntegrationsBotconnectorIntegrationIDBotNotFound with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotNotFound() *GetIntegrationsBotconnectorIntegrationIDBotNotFound {
	return &GetIntegrationsBotconnectorIntegrationIDBotNotFound{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsBotconnectorIntegrationIDBotNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotRequestTimeout creates a GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotRequestTimeout() *GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout {
	return &GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge creates a GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge() *GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge {
	return &GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType creates a GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType() *GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType {
	return &GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotTooManyRequests creates a GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotTooManyRequests() *GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests {
	return &GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotInternalServerError creates a GetIntegrationsBotconnectorIntegrationIDBotInternalServerError with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotInternalServerError() *GetIntegrationsBotconnectorIntegrationIDBotInternalServerError {
	return &GetIntegrationsBotconnectorIntegrationIDBotInternalServerError{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsBotconnectorIntegrationIDBotInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable creates a GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable() *GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable {
	return &GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout creates a GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout() *GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout {
	return &GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/{botId}][%d] getIntegrationsBotconnectorIntegrationIdBotGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
