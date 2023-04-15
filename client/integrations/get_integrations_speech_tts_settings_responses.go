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

// GetIntegrationsSpeechTtsSettingsReader is a Reader for the GetIntegrationsSpeechTtsSettings structure.
type GetIntegrationsSpeechTtsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechTtsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechTtsSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechTtsSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechTtsSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechTtsSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechTtsSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsSpeechTtsSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechTtsSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechTtsSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechTtsSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechTtsSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechTtsSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechTtsSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechTtsSettingsOK creates a GetIntegrationsSpeechTtsSettingsOK with default headers values
func NewGetIntegrationsSpeechTtsSettingsOK() *GetIntegrationsSpeechTtsSettingsOK {
	return &GetIntegrationsSpeechTtsSettingsOK{}
}

/*
GetIntegrationsSpeechTtsSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsSpeechTtsSettingsOK struct {
	Payload *models.TtsSettings
}

// IsSuccess returns true when this get integrations speech tts settings o k response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations speech tts settings o k response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings o k response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts settings o k response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings o k response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsSpeechTtsSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsOK) GetPayload() *models.TtsSettings {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TtsSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsBadRequest creates a GetIntegrationsSpeechTtsSettingsBadRequest with default headers values
func NewGetIntegrationsSpeechTtsSettingsBadRequest() *GetIntegrationsSpeechTtsSettingsBadRequest {
	return &GetIntegrationsSpeechTtsSettingsBadRequest{}
}

/*
GetIntegrationsSpeechTtsSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechTtsSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings bad request response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings bad request response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings bad request response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings bad request response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings bad request response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsSpeechTtsSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsUnauthorized creates a GetIntegrationsSpeechTtsSettingsUnauthorized with default headers values
func NewGetIntegrationsSpeechTtsSettingsUnauthorized() *GetIntegrationsSpeechTtsSettingsUnauthorized {
	return &GetIntegrationsSpeechTtsSettingsUnauthorized{}
}

/*
GetIntegrationsSpeechTtsSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechTtsSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings unauthorized response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings unauthorized response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings unauthorized response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings unauthorized response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings unauthorized response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsForbidden creates a GetIntegrationsSpeechTtsSettingsForbidden with default headers values
func NewGetIntegrationsSpeechTtsSettingsForbidden() *GetIntegrationsSpeechTtsSettingsForbidden {
	return &GetIntegrationsSpeechTtsSettingsForbidden{}
}

/*
GetIntegrationsSpeechTtsSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechTtsSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings forbidden response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings forbidden response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings forbidden response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings forbidden response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings forbidden response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsSpeechTtsSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsNotFound creates a GetIntegrationsSpeechTtsSettingsNotFound with default headers values
func NewGetIntegrationsSpeechTtsSettingsNotFound() *GetIntegrationsSpeechTtsSettingsNotFound {
	return &GetIntegrationsSpeechTtsSettingsNotFound{}
}

/*
GetIntegrationsSpeechTtsSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechTtsSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings not found response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings not found response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings not found response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings not found response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings not found response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsSpeechTtsSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsRequestTimeout creates a GetIntegrationsSpeechTtsSettingsRequestTimeout with default headers values
func NewGetIntegrationsSpeechTtsSettingsRequestTimeout() *GetIntegrationsSpeechTtsSettingsRequestTimeout {
	return &GetIntegrationsSpeechTtsSettingsRequestTimeout{}
}

/*
GetIntegrationsSpeechTtsSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsSpeechTtsSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings request timeout response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings request timeout response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings request timeout response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings request timeout response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings request timeout response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsRequestEntityTooLarge creates a GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechTtsSettingsRequestEntityTooLarge() *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge {
	return &GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge{}
}

/*
GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings request entity too large response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings request entity too large response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings request entity too large response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings request entity too large response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings request entity too large response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsUnsupportedMediaType creates a GetIntegrationsSpeechTtsSettingsUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechTtsSettingsUnsupportedMediaType() *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType {
	return &GetIntegrationsSpeechTtsSettingsUnsupportedMediaType{}
}

/*
GetIntegrationsSpeechTtsSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechTtsSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings unsupported media type response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings unsupported media type response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings unsupported media type response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings unsupported media type response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings unsupported media type response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsTooManyRequests creates a GetIntegrationsSpeechTtsSettingsTooManyRequests with default headers values
func NewGetIntegrationsSpeechTtsSettingsTooManyRequests() *GetIntegrationsSpeechTtsSettingsTooManyRequests {
	return &GetIntegrationsSpeechTtsSettingsTooManyRequests{}
}

/*
GetIntegrationsSpeechTtsSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsSpeechTtsSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings too many requests response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings too many requests response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings too many requests response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech tts settings too many requests response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech tts settings too many requests response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsInternalServerError creates a GetIntegrationsSpeechTtsSettingsInternalServerError with default headers values
func NewGetIntegrationsSpeechTtsSettingsInternalServerError() *GetIntegrationsSpeechTtsSettingsInternalServerError {
	return &GetIntegrationsSpeechTtsSettingsInternalServerError{}
}

/*
GetIntegrationsSpeechTtsSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechTtsSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings internal server error response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings internal server error response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings internal server error response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts settings internal server error response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech tts settings internal server error response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsServiceUnavailable creates a GetIntegrationsSpeechTtsSettingsServiceUnavailable with default headers values
func NewGetIntegrationsSpeechTtsSettingsServiceUnavailable() *GetIntegrationsSpeechTtsSettingsServiceUnavailable {
	return &GetIntegrationsSpeechTtsSettingsServiceUnavailable{}
}

/*
GetIntegrationsSpeechTtsSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechTtsSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings service unavailable response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings service unavailable response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings service unavailable response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts settings service unavailable response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech tts settings service unavailable response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechTtsSettingsGatewayTimeout creates a GetIntegrationsSpeechTtsSettingsGatewayTimeout with default headers values
func NewGetIntegrationsSpeechTtsSettingsGatewayTimeout() *GetIntegrationsSpeechTtsSettingsGatewayTimeout {
	return &GetIntegrationsSpeechTtsSettingsGatewayTimeout{}
}

/*
GetIntegrationsSpeechTtsSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsSpeechTtsSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech tts settings gateway timeout response has a 2xx status code
func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech tts settings gateway timeout response has a 3xx status code
func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech tts settings gateway timeout response has a 4xx status code
func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech tts settings gateway timeout response has a 5xx status code
func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech tts settings gateway timeout response a status code equal to that given
func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/tts/settings][%d] getIntegrationsSpeechTtsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechTtsSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
