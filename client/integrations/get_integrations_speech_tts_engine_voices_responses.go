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

// GetIntegrationsSpeechTtsEngineVoicesReader is a Reader for the GetIntegrationsSpeechTtsEngineVoices structure.
type GetIntegrationsSpeechTtsEngineVoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechTtsEngineVoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechTtsEngineVoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechTtsEngineVoicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechTtsEngineVoicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechTtsEngineVoicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechTtsEngineVoicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsSpeechTtsEngineVoicesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechTtsEngineVoicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechTtsEngineVoicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechTtsEngineVoicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechTtsEngineVoicesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechTtsEngineVoicesOK creates a GetIntegrationsSpeechTtsEngineVoicesOK with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesOK() *GetIntegrationsSpeechTtsEngineVoicesOK {
	return &GetIntegrationsSpeechTtsEngineVoicesOK{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsSpeechTtsEngineVoicesOK struct {
	Payload *models.TtsVoiceEntityListing
}

// IsSuccess returns true when this get integrations speech tts engine voices o k response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations speech tts engine voices o k response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices o k response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts engine voices o k response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices o k response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsSpeechTtsEngineVoicesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesOK) GetPayload() *models.TtsVoiceEntityListing {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TtsVoiceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesBadRequest creates a GetIntegrationsSpeechTtsEngineVoicesBadRequest with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesBadRequest() *GetIntegrationsSpeechTtsEngineVoicesBadRequest {
	return &GetIntegrationsSpeechTtsEngineVoicesBadRequest{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechTtsEngineVoicesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices bad request response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices bad request response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices bad request response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices bad request response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices bad request response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesUnauthorized creates a GetIntegrationsSpeechTtsEngineVoicesUnauthorized with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesUnauthorized() *GetIntegrationsSpeechTtsEngineVoicesUnauthorized {
	return &GetIntegrationsSpeechTtsEngineVoicesUnauthorized{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechTtsEngineVoicesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices unauthorized response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices unauthorized response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices unauthorized response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices unauthorized response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices unauthorized response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesForbidden creates a GetIntegrationsSpeechTtsEngineVoicesForbidden with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesForbidden() *GetIntegrationsSpeechTtsEngineVoicesForbidden {
	return &GetIntegrationsSpeechTtsEngineVoicesForbidden{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechTtsEngineVoicesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices forbidden response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices forbidden response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices forbidden response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices forbidden response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices forbidden response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesNotFound creates a GetIntegrationsSpeechTtsEngineVoicesNotFound with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesNotFound() *GetIntegrationsSpeechTtsEngineVoicesNotFound {
	return &GetIntegrationsSpeechTtsEngineVoicesNotFound{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechTtsEngineVoicesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices not found response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices not found response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices not found response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices not found response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices not found response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesRequestTimeout creates a GetIntegrationsSpeechTtsEngineVoicesRequestTimeout with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesRequestTimeout() *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout {
	return &GetIntegrationsSpeechTtsEngineVoicesRequestTimeout{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsSpeechTtsEngineVoicesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices request timeout response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices request timeout response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices request timeout response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices request timeout response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices request timeout response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge creates a GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge() *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge {
	return &GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices request entity too large response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices request entity too large response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices request entity too large response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices request entity too large response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices request entity too large response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType creates a GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType() *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType {
	return &GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices unsupported media type response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices unsupported media type response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices unsupported media type response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices unsupported media type response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices unsupported media type response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesTooManyRequests creates a GetIntegrationsSpeechTtsEngineVoicesTooManyRequests with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesTooManyRequests() *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests {
	return &GetIntegrationsSpeechTtsEngineVoicesTooManyRequests{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsSpeechTtsEngineVoicesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices too many requests response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices too many requests response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices too many requests response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts engine voices too many requests response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts engine voices too many requests response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesInternalServerError creates a GetIntegrationsSpeechTtsEngineVoicesInternalServerError with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesInternalServerError() *GetIntegrationsSpeechTtsEngineVoicesInternalServerError {
	return &GetIntegrationsSpeechTtsEngineVoicesInternalServerError{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechTtsEngineVoicesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices internal server error response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices internal server error response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices internal server error response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts engine voices internal server error response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech tts engine voices internal server error response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesServiceUnavailable creates a GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesServiceUnavailable() *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable {
	return &GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices service unavailable response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices service unavailable response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices service unavailable response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts engine voices service unavailable response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech tts engine voices service unavailable response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsEngineVoicesGatewayTimeout creates a GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout with default headers values
func NewGetIntegrationsSpeechTtsEngineVoicesGatewayTimeout() *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout {
	return &GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout{}
}

/*
GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts engine voices gateway timeout response has a 2xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts engine voices gateway timeout response has a 3xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts engine voices gateway timeout response has a 4xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts engine voices gateway timeout response has a 5xx status code
func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech tts engine voices gateway timeout response a status code equal to that given
func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/engines/{engineId}/voices][%d] getIntegrationsSpeechTtsEngineVoicesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsEngineVoicesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
