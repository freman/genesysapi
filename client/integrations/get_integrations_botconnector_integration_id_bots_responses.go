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

// GetIntegrationsBotconnectorIntegrationIDBotsReader is a Reader for the GetIntegrationsBotconnectorIntegrationIDBots structure.
type GetIntegrationsBotconnectorIntegrationIDBotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsBotconnectorIntegrationIDBotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsOK creates a GetIntegrationsBotconnectorIntegrationIDBotsOK with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsOK() *GetIntegrationsBotconnectorIntegrationIDBotsOK {
	return &GetIntegrationsBotconnectorIntegrationIDBotsOK{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsOK handles this case with default header values.

successful operation
*/
type GetIntegrationsBotconnectorIntegrationIDBotsOK struct {
	Payload *models.BotList
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsOK) GetPayload() *models.BotList {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BotList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsBadRequest creates a GetIntegrationsBotconnectorIntegrationIDBotsBadRequest with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsBadRequest() *GetIntegrationsBotconnectorIntegrationIDBotsBadRequest {
	return &GetIntegrationsBotconnectorIntegrationIDBotsBadRequest{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsUnauthorized creates a GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsUnauthorized() *GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized {
	return &GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsForbidden creates a GetIntegrationsBotconnectorIntegrationIDBotsForbidden with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsForbidden() *GetIntegrationsBotconnectorIntegrationIDBotsForbidden {
	return &GetIntegrationsBotconnectorIntegrationIDBotsForbidden{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsNotFound creates a GetIntegrationsBotconnectorIntegrationIDBotsNotFound with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsNotFound() *GetIntegrationsBotconnectorIntegrationIDBotsNotFound {
	return &GetIntegrationsBotconnectorIntegrationIDBotsNotFound{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout creates a GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout() *GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout {
	return &GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge creates a GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge() *GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge {
	return &GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType creates a GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType() *GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType {
	return &GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests creates a GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests() *GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests {
	return &GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsInternalServerError creates a GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsInternalServerError() *GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError {
	return &GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable creates a GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable() *GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable {
	return &GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout creates a GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout() *GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout {
	return &GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots][%d] getIntegrationsBotconnectorIntegrationIdBotsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}