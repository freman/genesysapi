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

// GetIntegrationsSpeechLexBotsReader is a Reader for the GetIntegrationsSpeechLexBots structure.
type GetIntegrationsSpeechLexBotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechLexBotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechLexBotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechLexBotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechLexBotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechLexBotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechLexBotsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsSpeechLexBotsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechLexBotsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechLexBotsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechLexBotsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechLexBotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechLexBotsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechLexBotsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechLexBotsOK creates a GetIntegrationsSpeechLexBotsOK with default headers values
func NewGetIntegrationsSpeechLexBotsOK() *GetIntegrationsSpeechLexBotsOK {
	return &GetIntegrationsSpeechLexBotsOK{}
}

/*GetIntegrationsSpeechLexBotsOK handles this case with default header values.

successful operation
*/
type GetIntegrationsSpeechLexBotsOK struct {
	Payload *models.LexBotEntityListing
}

func (o *GetIntegrationsSpeechLexBotsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsOK) GetPayload() *models.LexBotEntityListing {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LexBotEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsBadRequest creates a GetIntegrationsSpeechLexBotsBadRequest with default headers values
func NewGetIntegrationsSpeechLexBotsBadRequest() *GetIntegrationsSpeechLexBotsBadRequest {
	return &GetIntegrationsSpeechLexBotsBadRequest{}
}

/*GetIntegrationsSpeechLexBotsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechLexBotsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsUnauthorized creates a GetIntegrationsSpeechLexBotsUnauthorized with default headers values
func NewGetIntegrationsSpeechLexBotsUnauthorized() *GetIntegrationsSpeechLexBotsUnauthorized {
	return &GetIntegrationsSpeechLexBotsUnauthorized{}
}

/*GetIntegrationsSpeechLexBotsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechLexBotsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsForbidden creates a GetIntegrationsSpeechLexBotsForbidden with default headers values
func NewGetIntegrationsSpeechLexBotsForbidden() *GetIntegrationsSpeechLexBotsForbidden {
	return &GetIntegrationsSpeechLexBotsForbidden{}
}

/*GetIntegrationsSpeechLexBotsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechLexBotsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsNotFound creates a GetIntegrationsSpeechLexBotsNotFound with default headers values
func NewGetIntegrationsSpeechLexBotsNotFound() *GetIntegrationsSpeechLexBotsNotFound {
	return &GetIntegrationsSpeechLexBotsNotFound{}
}

/*GetIntegrationsSpeechLexBotsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechLexBotsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsRequestTimeout creates a GetIntegrationsSpeechLexBotsRequestTimeout with default headers values
func NewGetIntegrationsSpeechLexBotsRequestTimeout() *GetIntegrationsSpeechLexBotsRequestTimeout {
	return &GetIntegrationsSpeechLexBotsRequestTimeout{}
}

/*GetIntegrationsSpeechLexBotsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsSpeechLexBotsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsRequestEntityTooLarge creates a GetIntegrationsSpeechLexBotsRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechLexBotsRequestEntityTooLarge() *GetIntegrationsSpeechLexBotsRequestEntityTooLarge {
	return &GetIntegrationsSpeechLexBotsRequestEntityTooLarge{}
}

/*GetIntegrationsSpeechLexBotsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetIntegrationsSpeechLexBotsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsUnsupportedMediaType creates a GetIntegrationsSpeechLexBotsUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechLexBotsUnsupportedMediaType() *GetIntegrationsSpeechLexBotsUnsupportedMediaType {
	return &GetIntegrationsSpeechLexBotsUnsupportedMediaType{}
}

/*GetIntegrationsSpeechLexBotsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechLexBotsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsTooManyRequests creates a GetIntegrationsSpeechLexBotsTooManyRequests with default headers values
func NewGetIntegrationsSpeechLexBotsTooManyRequests() *GetIntegrationsSpeechLexBotsTooManyRequests {
	return &GetIntegrationsSpeechLexBotsTooManyRequests{}
}

/*GetIntegrationsSpeechLexBotsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsSpeechLexBotsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsInternalServerError creates a GetIntegrationsSpeechLexBotsInternalServerError with default headers values
func NewGetIntegrationsSpeechLexBotsInternalServerError() *GetIntegrationsSpeechLexBotsInternalServerError {
	return &GetIntegrationsSpeechLexBotsInternalServerError{}
}

/*GetIntegrationsSpeechLexBotsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechLexBotsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsServiceUnavailable creates a GetIntegrationsSpeechLexBotsServiceUnavailable with default headers values
func NewGetIntegrationsSpeechLexBotsServiceUnavailable() *GetIntegrationsSpeechLexBotsServiceUnavailable {
	return &GetIntegrationsSpeechLexBotsServiceUnavailable{}
}

/*GetIntegrationsSpeechLexBotsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechLexBotsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotsGatewayTimeout creates a GetIntegrationsSpeechLexBotsGatewayTimeout with default headers values
func NewGetIntegrationsSpeechLexBotsGatewayTimeout() *GetIntegrationsSpeechLexBotsGatewayTimeout {
	return &GetIntegrationsSpeechLexBotsGatewayTimeout{}
}

/*GetIntegrationsSpeechLexBotsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsSpeechLexBotsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechLexBotsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bots][%d] getIntegrationsSpeechLexBotsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
