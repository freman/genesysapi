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

// GetIntegrationsSpeechTtsEnginesReader is a Reader for the GetIntegrationsSpeechTtsEngines structure.
type GetIntegrationsSpeechTtsEnginesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechTtsEnginesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechTtsEnginesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechTtsEnginesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechTtsEnginesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechTtsEnginesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechTtsEnginesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechTtsEnginesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechTtsEnginesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechTtsEnginesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechTtsEnginesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechTtsEnginesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechTtsEnginesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechTtsEnginesOK creates a GetIntegrationsSpeechTtsEnginesOK with default headers values
func NewGetIntegrationsSpeechTtsEnginesOK() *GetIntegrationsSpeechTtsEnginesOK {
	return &GetIntegrationsSpeechTtsEnginesOK{}
}

/*GetIntegrationsSpeechTtsEnginesOK handles this case with default header values.

successful operation
*/
type GetIntegrationsSpeechTtsEnginesOK struct {
	Payload *models.TtsEngineEntityListing
}

func (o *GetIntegrationsSpeechTtsEnginesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesOK) GetPayload() *models.TtsEngineEntityListing {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TtsEngineEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesBadRequest creates a GetIntegrationsSpeechTtsEnginesBadRequest with default headers values
func NewGetIntegrationsSpeechTtsEnginesBadRequest() *GetIntegrationsSpeechTtsEnginesBadRequest {
	return &GetIntegrationsSpeechTtsEnginesBadRequest{}
}

/*GetIntegrationsSpeechTtsEnginesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechTtsEnginesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesUnauthorized creates a GetIntegrationsSpeechTtsEnginesUnauthorized with default headers values
func NewGetIntegrationsSpeechTtsEnginesUnauthorized() *GetIntegrationsSpeechTtsEnginesUnauthorized {
	return &GetIntegrationsSpeechTtsEnginesUnauthorized{}
}

/*GetIntegrationsSpeechTtsEnginesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechTtsEnginesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesForbidden creates a GetIntegrationsSpeechTtsEnginesForbidden with default headers values
func NewGetIntegrationsSpeechTtsEnginesForbidden() *GetIntegrationsSpeechTtsEnginesForbidden {
	return &GetIntegrationsSpeechTtsEnginesForbidden{}
}

/*GetIntegrationsSpeechTtsEnginesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechTtsEnginesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesNotFound creates a GetIntegrationsSpeechTtsEnginesNotFound with default headers values
func NewGetIntegrationsSpeechTtsEnginesNotFound() *GetIntegrationsSpeechTtsEnginesNotFound {
	return &GetIntegrationsSpeechTtsEnginesNotFound{}
}

/*GetIntegrationsSpeechTtsEnginesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechTtsEnginesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesRequestEntityTooLarge creates a GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechTtsEnginesRequestEntityTooLarge() *GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge {
	return &GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge{}
}

/*GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesUnsupportedMediaType creates a GetIntegrationsSpeechTtsEnginesUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechTtsEnginesUnsupportedMediaType() *GetIntegrationsSpeechTtsEnginesUnsupportedMediaType {
	return &GetIntegrationsSpeechTtsEnginesUnsupportedMediaType{}
}

/*GetIntegrationsSpeechTtsEnginesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechTtsEnginesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesTooManyRequests creates a GetIntegrationsSpeechTtsEnginesTooManyRequests with default headers values
func NewGetIntegrationsSpeechTtsEnginesTooManyRequests() *GetIntegrationsSpeechTtsEnginesTooManyRequests {
	return &GetIntegrationsSpeechTtsEnginesTooManyRequests{}
}

/*GetIntegrationsSpeechTtsEnginesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsSpeechTtsEnginesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesInternalServerError creates a GetIntegrationsSpeechTtsEnginesInternalServerError with default headers values
func NewGetIntegrationsSpeechTtsEnginesInternalServerError() *GetIntegrationsSpeechTtsEnginesInternalServerError {
	return &GetIntegrationsSpeechTtsEnginesInternalServerError{}
}

/*GetIntegrationsSpeechTtsEnginesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechTtsEnginesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesServiceUnavailable creates a GetIntegrationsSpeechTtsEnginesServiceUnavailable with default headers values
func NewGetIntegrationsSpeechTtsEnginesServiceUnavailable() *GetIntegrationsSpeechTtsEnginesServiceUnavailable {
	return &GetIntegrationsSpeechTtsEnginesServiceUnavailable{}
}

/*GetIntegrationsSpeechTtsEnginesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechTtsEnginesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEnginesGatewayTimeout creates a GetIntegrationsSpeechTtsEnginesGatewayTimeout with default headers values
func NewGetIntegrationsSpeechTtsEnginesGatewayTimeout() *GetIntegrationsSpeechTtsEnginesGatewayTimeout {
	return &GetIntegrationsSpeechTtsEnginesGatewayTimeout{}
}

/*GetIntegrationsSpeechTtsEnginesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsSpeechTtsEnginesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsSpeechTtsEnginesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines][%d] getIntegrationsSpeechTtsEnginesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEnginesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEnginesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
