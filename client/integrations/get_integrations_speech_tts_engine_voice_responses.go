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

// GetIntegrationsSpeechTtsEngineVoiceReader is a Reader for the GetIntegrationsSpeechTtsEngineVoice structure.
type GetIntegrationsSpeechTtsEngineVoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechTtsEngineVoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechTtsEngineVoiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechTtsEngineVoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechTtsEngineVoiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechTtsEngineVoiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechTtsEngineVoiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsSpeechTtsEngineVoiceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechTtsEngineVoiceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechTtsEngineVoiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechTtsEngineVoiceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechTtsEngineVoiceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechTtsEngineVoiceOK creates a GetIntegrationsSpeechTtsEngineVoiceOK with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceOK() *GetIntegrationsSpeechTtsEngineVoiceOK {
	return &GetIntegrationsSpeechTtsEngineVoiceOK{}
}

/*GetIntegrationsSpeechTtsEngineVoiceOK handles this case with default header values.

successful operation
*/
type GetIntegrationsSpeechTtsEngineVoiceOK struct {
	Payload *models.TtsVoiceEntity
}

func (o *GetIntegrationsSpeechTtsEngineVoiceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceOK) GetPayload() *models.TtsVoiceEntity {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TtsVoiceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceBadRequest creates a GetIntegrationsSpeechTtsEngineVoiceBadRequest with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceBadRequest() *GetIntegrationsSpeechTtsEngineVoiceBadRequest {
	return &GetIntegrationsSpeechTtsEngineVoiceBadRequest{}
}

/*GetIntegrationsSpeechTtsEngineVoiceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechTtsEngineVoiceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceUnauthorized creates a GetIntegrationsSpeechTtsEngineVoiceUnauthorized with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceUnauthorized() *GetIntegrationsSpeechTtsEngineVoiceUnauthorized {
	return &GetIntegrationsSpeechTtsEngineVoiceUnauthorized{}
}

/*GetIntegrationsSpeechTtsEngineVoiceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechTtsEngineVoiceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceForbidden creates a GetIntegrationsSpeechTtsEngineVoiceForbidden with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceForbidden() *GetIntegrationsSpeechTtsEngineVoiceForbidden {
	return &GetIntegrationsSpeechTtsEngineVoiceForbidden{}
}

/*GetIntegrationsSpeechTtsEngineVoiceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechTtsEngineVoiceForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceNotFound creates a GetIntegrationsSpeechTtsEngineVoiceNotFound with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceNotFound() *GetIntegrationsSpeechTtsEngineVoiceNotFound {
	return &GetIntegrationsSpeechTtsEngineVoiceNotFound{}
}

/*GetIntegrationsSpeechTtsEngineVoiceNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechTtsEngineVoiceNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceRequestTimeout creates a GetIntegrationsSpeechTtsEngineVoiceRequestTimeout with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceRequestTimeout() *GetIntegrationsSpeechTtsEngineVoiceRequestTimeout {
	return &GetIntegrationsSpeechTtsEngineVoiceRequestTimeout{}
}

/*GetIntegrationsSpeechTtsEngineVoiceRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsSpeechTtsEngineVoiceRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge creates a GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge() *GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge {
	return &GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge{}
}

/*GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType creates a GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType() *GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType {
	return &GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType{}
}

/*GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceTooManyRequests creates a GetIntegrationsSpeechTtsEngineVoiceTooManyRequests with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceTooManyRequests() *GetIntegrationsSpeechTtsEngineVoiceTooManyRequests {
	return &GetIntegrationsSpeechTtsEngineVoiceTooManyRequests{}
}

/*GetIntegrationsSpeechTtsEngineVoiceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsSpeechTtsEngineVoiceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceInternalServerError creates a GetIntegrationsSpeechTtsEngineVoiceInternalServerError with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceInternalServerError() *GetIntegrationsSpeechTtsEngineVoiceInternalServerError {
	return &GetIntegrationsSpeechTtsEngineVoiceInternalServerError{}
}

/*GetIntegrationsSpeechTtsEngineVoiceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechTtsEngineVoiceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceServiceUnavailable creates a GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceServiceUnavailable() *GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable {
	return &GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable{}
}

/*GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoiceGatewayTimeout creates a GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout with default headers values
func NewGetIntegrationsSpeechTtsEngineVoiceGatewayTimeout() *GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout {
	return &GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout{}
}

/*GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices/{voiceId}][%d] getIntegrationsSpeechTtsEngineVoiceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoiceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
