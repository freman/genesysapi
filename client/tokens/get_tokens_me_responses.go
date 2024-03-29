// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTokensMeReader is a Reader for the GetTokensMe structure.
type GetTokensMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokensMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTokensMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTokensMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTokensMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTokensMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTokensMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTokensMeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTokensMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTokensMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTokensMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTokensMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTokensMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTokensMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTokensMeOK creates a GetTokensMeOK with default headers values
func NewGetTokensMeOK() *GetTokensMeOK {
	return &GetTokensMeOK{}
}

/*
GetTokensMeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTokensMeOK struct {
	Payload *models.TokenInfo
}

// IsSuccess returns true when this get tokens me o k response has a 2xx status code
func (o *GetTokensMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tokens me o k response has a 3xx status code
func (o *GetTokensMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me o k response has a 4xx status code
func (o *GetTokensMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tokens me o k response has a 5xx status code
func (o *GetTokensMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me o k response a status code equal to that given
func (o *GetTokensMeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTokensMeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeOK  %+v", 200, o.Payload)
}

func (o *GetTokensMeOK) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeOK  %+v", 200, o.Payload)
}

func (o *GetTokensMeOK) GetPayload() *models.TokenInfo {
	return o.Payload
}

func (o *GetTokensMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeBadRequest creates a GetTokensMeBadRequest with default headers values
func NewGetTokensMeBadRequest() *GetTokensMeBadRequest {
	return &GetTokensMeBadRequest{}
}

/*
GetTokensMeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTokensMeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me bad request response has a 2xx status code
func (o *GetTokensMeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me bad request response has a 3xx status code
func (o *GetTokensMeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me bad request response has a 4xx status code
func (o *GetTokensMeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me bad request response has a 5xx status code
func (o *GetTokensMeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me bad request response a status code equal to that given
func (o *GetTokensMeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetTokensMeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetTokensMeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetTokensMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeUnauthorized creates a GetTokensMeUnauthorized with default headers values
func NewGetTokensMeUnauthorized() *GetTokensMeUnauthorized {
	return &GetTokensMeUnauthorized{}
}

/*
GetTokensMeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTokensMeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me unauthorized response has a 2xx status code
func (o *GetTokensMeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me unauthorized response has a 3xx status code
func (o *GetTokensMeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me unauthorized response has a 4xx status code
func (o *GetTokensMeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me unauthorized response has a 5xx status code
func (o *GetTokensMeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me unauthorized response a status code equal to that given
func (o *GetTokensMeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetTokensMeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTokensMeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTokensMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeForbidden creates a GetTokensMeForbidden with default headers values
func NewGetTokensMeForbidden() *GetTokensMeForbidden {
	return &GetTokensMeForbidden{}
}

/*
GetTokensMeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetTokensMeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me forbidden response has a 2xx status code
func (o *GetTokensMeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me forbidden response has a 3xx status code
func (o *GetTokensMeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me forbidden response has a 4xx status code
func (o *GetTokensMeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me forbidden response has a 5xx status code
func (o *GetTokensMeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me forbidden response a status code equal to that given
func (o *GetTokensMeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetTokensMeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeForbidden  %+v", 403, o.Payload)
}

func (o *GetTokensMeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeForbidden  %+v", 403, o.Payload)
}

func (o *GetTokensMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeNotFound creates a GetTokensMeNotFound with default headers values
func NewGetTokensMeNotFound() *GetTokensMeNotFound {
	return &GetTokensMeNotFound{}
}

/*
GetTokensMeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetTokensMeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me not found response has a 2xx status code
func (o *GetTokensMeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me not found response has a 3xx status code
func (o *GetTokensMeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me not found response has a 4xx status code
func (o *GetTokensMeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me not found response has a 5xx status code
func (o *GetTokensMeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me not found response a status code equal to that given
func (o *GetTokensMeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetTokensMeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeNotFound  %+v", 404, o.Payload)
}

func (o *GetTokensMeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeNotFound  %+v", 404, o.Payload)
}

func (o *GetTokensMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeRequestTimeout creates a GetTokensMeRequestTimeout with default headers values
func NewGetTokensMeRequestTimeout() *GetTokensMeRequestTimeout {
	return &GetTokensMeRequestTimeout{}
}

/*
GetTokensMeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTokensMeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me request timeout response has a 2xx status code
func (o *GetTokensMeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me request timeout response has a 3xx status code
func (o *GetTokensMeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me request timeout response has a 4xx status code
func (o *GetTokensMeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me request timeout response has a 5xx status code
func (o *GetTokensMeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me request timeout response a status code equal to that given
func (o *GetTokensMeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetTokensMeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTokensMeRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTokensMeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeRequestEntityTooLarge creates a GetTokensMeRequestEntityTooLarge with default headers values
func NewGetTokensMeRequestEntityTooLarge() *GetTokensMeRequestEntityTooLarge {
	return &GetTokensMeRequestEntityTooLarge{}
}

/*
GetTokensMeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetTokensMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me request entity too large response has a 2xx status code
func (o *GetTokensMeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me request entity too large response has a 3xx status code
func (o *GetTokensMeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me request entity too large response has a 4xx status code
func (o *GetTokensMeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me request entity too large response has a 5xx status code
func (o *GetTokensMeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me request entity too large response a status code equal to that given
func (o *GetTokensMeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetTokensMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTokensMeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTokensMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeUnsupportedMediaType creates a GetTokensMeUnsupportedMediaType with default headers values
func NewGetTokensMeUnsupportedMediaType() *GetTokensMeUnsupportedMediaType {
	return &GetTokensMeUnsupportedMediaType{}
}

/*
GetTokensMeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTokensMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me unsupported media type response has a 2xx status code
func (o *GetTokensMeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me unsupported media type response has a 3xx status code
func (o *GetTokensMeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me unsupported media type response has a 4xx status code
func (o *GetTokensMeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me unsupported media type response has a 5xx status code
func (o *GetTokensMeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me unsupported media type response a status code equal to that given
func (o *GetTokensMeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetTokensMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTokensMeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTokensMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeTooManyRequests creates a GetTokensMeTooManyRequests with default headers values
func NewGetTokensMeTooManyRequests() *GetTokensMeTooManyRequests {
	return &GetTokensMeTooManyRequests{}
}

/*
GetTokensMeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTokensMeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me too many requests response has a 2xx status code
func (o *GetTokensMeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me too many requests response has a 3xx status code
func (o *GetTokensMeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me too many requests response has a 4xx status code
func (o *GetTokensMeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get tokens me too many requests response has a 5xx status code
func (o *GetTokensMeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get tokens me too many requests response a status code equal to that given
func (o *GetTokensMeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetTokensMeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTokensMeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTokensMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeInternalServerError creates a GetTokensMeInternalServerError with default headers values
func NewGetTokensMeInternalServerError() *GetTokensMeInternalServerError {
	return &GetTokensMeInternalServerError{}
}

/*
GetTokensMeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTokensMeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me internal server error response has a 2xx status code
func (o *GetTokensMeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me internal server error response has a 3xx status code
func (o *GetTokensMeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me internal server error response has a 4xx status code
func (o *GetTokensMeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tokens me internal server error response has a 5xx status code
func (o *GetTokensMeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get tokens me internal server error response a status code equal to that given
func (o *GetTokensMeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTokensMeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTokensMeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTokensMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeServiceUnavailable creates a GetTokensMeServiceUnavailable with default headers values
func NewGetTokensMeServiceUnavailable() *GetTokensMeServiceUnavailable {
	return &GetTokensMeServiceUnavailable{}
}

/*
GetTokensMeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTokensMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me service unavailable response has a 2xx status code
func (o *GetTokensMeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me service unavailable response has a 3xx status code
func (o *GetTokensMeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me service unavailable response has a 4xx status code
func (o *GetTokensMeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tokens me service unavailable response has a 5xx status code
func (o *GetTokensMeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get tokens me service unavailable response a status code equal to that given
func (o *GetTokensMeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetTokensMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTokensMeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTokensMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokensMeGatewayTimeout creates a GetTokensMeGatewayTimeout with default headers values
func NewGetTokensMeGatewayTimeout() *GetTokensMeGatewayTimeout {
	return &GetTokensMeGatewayTimeout{}
}

/*
GetTokensMeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetTokensMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get tokens me gateway timeout response has a 2xx status code
func (o *GetTokensMeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get tokens me gateway timeout response has a 3xx status code
func (o *GetTokensMeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tokens me gateway timeout response has a 4xx status code
func (o *GetTokensMeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tokens me gateway timeout response has a 5xx status code
func (o *GetTokensMeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get tokens me gateway timeout response a status code equal to that given
func (o *GetTokensMeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetTokensMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTokensMeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/tokens/me][%d] getTokensMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTokensMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTokensMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
