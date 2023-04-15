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

// GetIntegrationsSpeechLexBotAliasReader is a Reader for the GetIntegrationsSpeechLexBotAlias structure.
type GetIntegrationsSpeechLexBotAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsSpeechLexBotAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsSpeechLexBotAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsSpeechLexBotAliasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsSpeechLexBotAliasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsSpeechLexBotAliasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsSpeechLexBotAliasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsSpeechLexBotAliasRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsSpeechLexBotAliasRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsSpeechLexBotAliasUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsSpeechLexBotAliasTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsSpeechLexBotAliasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsSpeechLexBotAliasServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsSpeechLexBotAliasGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsSpeechLexBotAliasOK creates a GetIntegrationsSpeechLexBotAliasOK with default headers values
func NewGetIntegrationsSpeechLexBotAliasOK() *GetIntegrationsSpeechLexBotAliasOK {
	return &GetIntegrationsSpeechLexBotAliasOK{}
}

/*
GetIntegrationsSpeechLexBotAliasOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsSpeechLexBotAliasOK struct {
	Payload *models.LexBotAlias
}

// IsSuccess returns true when this get integrations speech lex bot alias o k response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations speech lex bot alias o k response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias o k response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech lex bot alias o k response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias o k response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsSpeechLexBotAliasOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasOK) GetPayload() *models.LexBotAlias {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LexBotAlias)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasBadRequest creates a GetIntegrationsSpeechLexBotAliasBadRequest with default headers values
func NewGetIntegrationsSpeechLexBotAliasBadRequest() *GetIntegrationsSpeechLexBotAliasBadRequest {
	return &GetIntegrationsSpeechLexBotAliasBadRequest{}
}

/*
GetIntegrationsSpeechLexBotAliasBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsSpeechLexBotAliasBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias bad request response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias bad request response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias bad request response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias bad request response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias bad request response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsSpeechLexBotAliasBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasUnauthorized creates a GetIntegrationsSpeechLexBotAliasUnauthorized with default headers values
func NewGetIntegrationsSpeechLexBotAliasUnauthorized() *GetIntegrationsSpeechLexBotAliasUnauthorized {
	return &GetIntegrationsSpeechLexBotAliasUnauthorized{}
}

/*
GetIntegrationsSpeechLexBotAliasUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsSpeechLexBotAliasUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias unauthorized response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias unauthorized response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias unauthorized response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias unauthorized response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias unauthorized response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasForbidden creates a GetIntegrationsSpeechLexBotAliasForbidden with default headers values
func NewGetIntegrationsSpeechLexBotAliasForbidden() *GetIntegrationsSpeechLexBotAliasForbidden {
	return &GetIntegrationsSpeechLexBotAliasForbidden{}
}

/*
GetIntegrationsSpeechLexBotAliasForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsSpeechLexBotAliasForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias forbidden response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias forbidden response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias forbidden response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias forbidden response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias forbidden response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsSpeechLexBotAliasForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasNotFound creates a GetIntegrationsSpeechLexBotAliasNotFound with default headers values
func NewGetIntegrationsSpeechLexBotAliasNotFound() *GetIntegrationsSpeechLexBotAliasNotFound {
	return &GetIntegrationsSpeechLexBotAliasNotFound{}
}

/*
GetIntegrationsSpeechLexBotAliasNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsSpeechLexBotAliasNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias not found response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias not found response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias not found response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias not found response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias not found response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsSpeechLexBotAliasNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasRequestTimeout creates a GetIntegrationsSpeechLexBotAliasRequestTimeout with default headers values
func NewGetIntegrationsSpeechLexBotAliasRequestTimeout() *GetIntegrationsSpeechLexBotAliasRequestTimeout {
	return &GetIntegrationsSpeechLexBotAliasRequestTimeout{}
}

/*
GetIntegrationsSpeechLexBotAliasRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsSpeechLexBotAliasRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias request timeout response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias request timeout response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias request timeout response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias request timeout response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias request timeout response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasRequestEntityTooLarge creates a GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge with default headers values
func NewGetIntegrationsSpeechLexBotAliasRequestEntityTooLarge() *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge {
	return &GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge{}
}

/*
GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias request entity too large response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias request entity too large response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias request entity too large response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias request entity too large response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias request entity too large response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasUnsupportedMediaType creates a GetIntegrationsSpeechLexBotAliasUnsupportedMediaType with default headers values
func NewGetIntegrationsSpeechLexBotAliasUnsupportedMediaType() *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType {
	return &GetIntegrationsSpeechLexBotAliasUnsupportedMediaType{}
}

/*
GetIntegrationsSpeechLexBotAliasUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsSpeechLexBotAliasUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias unsupported media type response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias unsupported media type response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias unsupported media type response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias unsupported media type response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias unsupported media type response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasTooManyRequests creates a GetIntegrationsSpeechLexBotAliasTooManyRequests with default headers values
func NewGetIntegrationsSpeechLexBotAliasTooManyRequests() *GetIntegrationsSpeechLexBotAliasTooManyRequests {
	return &GetIntegrationsSpeechLexBotAliasTooManyRequests{}
}

/*
GetIntegrationsSpeechLexBotAliasTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsSpeechLexBotAliasTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias too many requests response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias too many requests response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias too many requests response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations speech lex bot alias too many requests response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations speech lex bot alias too many requests response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasInternalServerError creates a GetIntegrationsSpeechLexBotAliasInternalServerError with default headers values
func NewGetIntegrationsSpeechLexBotAliasInternalServerError() *GetIntegrationsSpeechLexBotAliasInternalServerError {
	return &GetIntegrationsSpeechLexBotAliasInternalServerError{}
}

/*
GetIntegrationsSpeechLexBotAliasInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsSpeechLexBotAliasInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias internal server error response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias internal server error response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias internal server error response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech lex bot alias internal server error response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech lex bot alias internal server error response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasServiceUnavailable creates a GetIntegrationsSpeechLexBotAliasServiceUnavailable with default headers values
func NewGetIntegrationsSpeechLexBotAliasServiceUnavailable() *GetIntegrationsSpeechLexBotAliasServiceUnavailable {
	return &GetIntegrationsSpeechLexBotAliasServiceUnavailable{}
}

/*
GetIntegrationsSpeechLexBotAliasServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsSpeechLexBotAliasServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias service unavailable response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias service unavailable response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias service unavailable response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech lex bot alias service unavailable response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech lex bot alias service unavailable response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsSpeechLexBotAliasGatewayTimeout creates a GetIntegrationsSpeechLexBotAliasGatewayTimeout with default headers values
func NewGetIntegrationsSpeechLexBotAliasGatewayTimeout() *GetIntegrationsSpeechLexBotAliasGatewayTimeout {
	return &GetIntegrationsSpeechLexBotAliasGatewayTimeout{}
}

/*
GetIntegrationsSpeechLexBotAliasGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsSpeechLexBotAliasGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations speech lex bot alias gateway timeout response has a 2xx status code
func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations speech lex bot alias gateway timeout response has a 3xx status code
func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations speech lex bot alias gateway timeout response has a 4xx status code
func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations speech lex bot alias gateway timeout response has a 5xx status code
func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations speech lex bot alias gateway timeout response a status code equal to that given
func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/speech/lex/bot/alias/{aliasId}][%d] getIntegrationsSpeechLexBotAliasGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsSpeechLexBotAliasGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
