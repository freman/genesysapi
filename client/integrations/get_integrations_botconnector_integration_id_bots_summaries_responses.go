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

// GetIntegrationsBotconnectorIntegrationIDBotsSummariesReader is a Reader for the GetIntegrationsBotconnectorIntegrationIDBotsSummaries structure.
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesOK creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesOK() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK handles this case with default header values.

successful operation
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK struct {
	Payload *models.BotConnectorBotSummaryEntityListing
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK) GetPayload() *models.BotConnectorBotSummaryEntityListing {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BotConnectorBotSummaryEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout creates a GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout with default headers values
func NewGetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout() *GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout {
	return &GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout{}
}

/*GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/botconnector/{integrationId}/bots/summaries][%d] getIntegrationsBotconnectorIntegrationIdBotsSummariesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsBotconnectorIntegrationIDBotsSummariesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
