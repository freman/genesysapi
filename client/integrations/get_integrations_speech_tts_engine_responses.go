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

// GetIntegrationsSpeechTtsEngineReader is a Reader for the GetIntegrationsSpeechTtsEngine structure.
type GetIntegrationsSpeechTtsEngineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechTtsEngineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechTtsEngineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechTtsEngineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechTtsEngineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechTtsEngineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechTtsEngineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechTtsEngineRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechTtsEngineUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechTtsEngineTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechTtsEngineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechTtsEngineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechTtsEngineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechTtsEngineOK creates a GetIntegrationsSpeechTtsEngineOK with default headers values
func NewGetIntegrationsSpeechTtsEngineOK() *GetIntegrationsSpeechTtsEngineOK {
	return &GetIntegrationsSpeechTtsEngineOK{}
}

/*GetIntegrationsSpeechTtsEngineOK handles this case with default header values.

successful operation
*/
type GetIntegrationsSpeechTtsEngineOK struct {
	Payload *models.TtsEngineEntity
}

func (o *GetIntegrationsSpeechTtsEngineOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineOK) GetPayload() *models.TtsEngineEntity {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TtsEngineEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineBadRequest creates a GetIntegrationsSpeechTtsEngineBadRequest with default headers values
func NewGetIntegrationsSpeechTtsEngineBadRequest() *GetIntegrationsSpeechTtsEngineBadRequest {
	return &GetIntegrationsSpeechTtsEngineBadRequest{}
}

/*GetIntegrationsSpeechTtsEngineBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechTtsEngineBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineUnauthorized creates a GetIntegrationsSpeechTtsEngineUnauthorized with default headers values
func NewGetIntegrationsSpeechTtsEngineUnauthorized() *GetIntegrationsSpeechTtsEngineUnauthorized {
	return &GetIntegrationsSpeechTtsEngineUnauthorized{}
}

/*GetIntegrationsSpeechTtsEngineUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechTtsEngineUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineForbidden creates a GetIntegrationsSpeechTtsEngineForbidden with default headers values
func NewGetIntegrationsSpeechTtsEngineForbidden() *GetIntegrationsSpeechTtsEngineForbidden {
	return &GetIntegrationsSpeechTtsEngineForbidden{}
}

/*GetIntegrationsSpeechTtsEngineForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechTtsEngineForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineNotFound creates a GetIntegrationsSpeechTtsEngineNotFound with default headers values
func NewGetIntegrationsSpeechTtsEngineNotFound() *GetIntegrationsSpeechTtsEngineNotFound {
	return &GetIntegrationsSpeechTtsEngineNotFound{}
}

/*GetIntegrationsSpeechTtsEngineNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechTtsEngineNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineRequestEntityTooLarge creates a GetIntegrationsSpeechTtsEngineRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechTtsEngineRequestEntityTooLarge() *GetIntegrationsSpeechTtsEngineRequestEntityTooLarge {
	return &GetIntegrationsSpeechTtsEngineRequestEntityTooLarge{}
}

/*GetIntegrationsSpeechTtsEngineRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsSpeechTtsEngineRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineUnsupportedMediaType creates a GetIntegrationsSpeechTtsEngineUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechTtsEngineUnsupportedMediaType() *GetIntegrationsSpeechTtsEngineUnsupportedMediaType {
	return &GetIntegrationsSpeechTtsEngineUnsupportedMediaType{}
}

/*GetIntegrationsSpeechTtsEngineUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechTtsEngineUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineTooManyRequests creates a GetIntegrationsSpeechTtsEngineTooManyRequests with default headers values
func NewGetIntegrationsSpeechTtsEngineTooManyRequests() *GetIntegrationsSpeechTtsEngineTooManyRequests {
	return &GetIntegrationsSpeechTtsEngineTooManyRequests{}
}

/*GetIntegrationsSpeechTtsEngineTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsSpeechTtsEngineTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineInternalServerError creates a GetIntegrationsSpeechTtsEngineInternalServerError with default headers values
func NewGetIntegrationsSpeechTtsEngineInternalServerError() *GetIntegrationsSpeechTtsEngineInternalServerError {
	return &GetIntegrationsSpeechTtsEngineInternalServerError{}
}

/*GetIntegrationsSpeechTtsEngineInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechTtsEngineInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineServiceUnavailable creates a GetIntegrationsSpeechTtsEngineServiceUnavailable with default headers values
func NewGetIntegrationsSpeechTtsEngineServiceUnavailable() *GetIntegrationsSpeechTtsEngineServiceUnavailable {
	return &GetIntegrationsSpeechTtsEngineServiceUnavailable{}
}

/*GetIntegrationsSpeechTtsEngineServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechTtsEngineServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineGatewayTimeout creates a GetIntegrationsSpeechTtsEngineGatewayTimeout with default headers values
func NewGetIntegrationsSpeechTtsEngineGatewayTimeout() *GetIntegrationsSpeechTtsEngineGatewayTimeout {
	return &GetIntegrationsSpeechTtsEngineGatewayTimeout{}
}

/*GetIntegrationsSpeechTtsEngineGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsSpeechTtsEngineGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEngineGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}][%d] getIntegrationsSpeechTtsEngineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
