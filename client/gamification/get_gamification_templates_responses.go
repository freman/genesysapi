// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGamificationTemplatesReader is a Reader for the GetGamificationTemplates structure.
type GetGamificationTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationTemplatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationTemplatesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationTemplatesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationTemplatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationTemplatesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationTemplatesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationTemplatesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationTemplatesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationTemplatesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationTemplatesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationTemplatesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationTemplatesOK creates a GetGamificationTemplatesOK with default headers values
func NewGetGamificationTemplatesOK() *GetGamificationTemplatesOK {
	return &GetGamificationTemplatesOK{}
}

/*
GetGamificationTemplatesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGamificationTemplatesOK struct {
	Payload *models.GetTemplatesResponse
}

// IsSuccess returns true when this get gamification templates o k response has a 2xx status code
func (o *GetGamificationTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gamification templates o k response has a 3xx status code
func (o *GetGamificationTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates o k response has a 4xx status code
func (o *GetGamificationTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification templates o k response has a 5xx status code
func (o *GetGamificationTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates o k response a status code equal to that given
func (o *GetGamificationTemplatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGamificationTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetGamificationTemplatesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetGamificationTemplatesOK) GetPayload() *models.GetTemplatesResponse {
	return o.Payload
}

func (o *GetGamificationTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTemplatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesBadRequest creates a GetGamificationTemplatesBadRequest with default headers values
func NewGetGamificationTemplatesBadRequest() *GetGamificationTemplatesBadRequest {
	return &GetGamificationTemplatesBadRequest{}
}

/*
GetGamificationTemplatesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationTemplatesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates bad request response has a 2xx status code
func (o *GetGamificationTemplatesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates bad request response has a 3xx status code
func (o *GetGamificationTemplatesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates bad request response has a 4xx status code
func (o *GetGamificationTemplatesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates bad request response has a 5xx status code
func (o *GetGamificationTemplatesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates bad request response a status code equal to that given
func (o *GetGamificationTemplatesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGamificationTemplatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationTemplatesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationTemplatesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesUnauthorized creates a GetGamificationTemplatesUnauthorized with default headers values
func NewGetGamificationTemplatesUnauthorized() *GetGamificationTemplatesUnauthorized {
	return &GetGamificationTemplatesUnauthorized{}
}

/*
GetGamificationTemplatesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationTemplatesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates unauthorized response has a 2xx status code
func (o *GetGamificationTemplatesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates unauthorized response has a 3xx status code
func (o *GetGamificationTemplatesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates unauthorized response has a 4xx status code
func (o *GetGamificationTemplatesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates unauthorized response has a 5xx status code
func (o *GetGamificationTemplatesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates unauthorized response a status code equal to that given
func (o *GetGamificationTemplatesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGamificationTemplatesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationTemplatesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationTemplatesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesForbidden creates a GetGamificationTemplatesForbidden with default headers values
func NewGetGamificationTemplatesForbidden() *GetGamificationTemplatesForbidden {
	return &GetGamificationTemplatesForbidden{}
}

/*
GetGamificationTemplatesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationTemplatesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates forbidden response has a 2xx status code
func (o *GetGamificationTemplatesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates forbidden response has a 3xx status code
func (o *GetGamificationTemplatesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates forbidden response has a 4xx status code
func (o *GetGamificationTemplatesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates forbidden response has a 5xx status code
func (o *GetGamificationTemplatesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates forbidden response a status code equal to that given
func (o *GetGamificationTemplatesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGamificationTemplatesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationTemplatesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationTemplatesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesNotFound creates a GetGamificationTemplatesNotFound with default headers values
func NewGetGamificationTemplatesNotFound() *GetGamificationTemplatesNotFound {
	return &GetGamificationTemplatesNotFound{}
}

/*
GetGamificationTemplatesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGamificationTemplatesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates not found response has a 2xx status code
func (o *GetGamificationTemplatesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates not found response has a 3xx status code
func (o *GetGamificationTemplatesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates not found response has a 4xx status code
func (o *GetGamificationTemplatesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates not found response has a 5xx status code
func (o *GetGamificationTemplatesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates not found response a status code equal to that given
func (o *GetGamificationTemplatesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGamificationTemplatesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationTemplatesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationTemplatesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesRequestTimeout creates a GetGamificationTemplatesRequestTimeout with default headers values
func NewGetGamificationTemplatesRequestTimeout() *GetGamificationTemplatesRequestTimeout {
	return &GetGamificationTemplatesRequestTimeout{}
}

/*
GetGamificationTemplatesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationTemplatesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates request timeout response has a 2xx status code
func (o *GetGamificationTemplatesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates request timeout response has a 3xx status code
func (o *GetGamificationTemplatesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates request timeout response has a 4xx status code
func (o *GetGamificationTemplatesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates request timeout response has a 5xx status code
func (o *GetGamificationTemplatesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates request timeout response a status code equal to that given
func (o *GetGamificationTemplatesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGamificationTemplatesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationTemplatesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationTemplatesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesRequestEntityTooLarge creates a GetGamificationTemplatesRequestEntityTooLarge with default headers values
func NewGetGamificationTemplatesRequestEntityTooLarge() *GetGamificationTemplatesRequestEntityTooLarge {
	return &GetGamificationTemplatesRequestEntityTooLarge{}
}

/*
GetGamificationTemplatesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetGamificationTemplatesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates request entity too large response has a 2xx status code
func (o *GetGamificationTemplatesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates request entity too large response has a 3xx status code
func (o *GetGamificationTemplatesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates request entity too large response has a 4xx status code
func (o *GetGamificationTemplatesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates request entity too large response has a 5xx status code
func (o *GetGamificationTemplatesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates request entity too large response a status code equal to that given
func (o *GetGamificationTemplatesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGamificationTemplatesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationTemplatesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationTemplatesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesUnsupportedMediaType creates a GetGamificationTemplatesUnsupportedMediaType with default headers values
func NewGetGamificationTemplatesUnsupportedMediaType() *GetGamificationTemplatesUnsupportedMediaType {
	return &GetGamificationTemplatesUnsupportedMediaType{}
}

/*
GetGamificationTemplatesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationTemplatesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates unsupported media type response has a 2xx status code
func (o *GetGamificationTemplatesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates unsupported media type response has a 3xx status code
func (o *GetGamificationTemplatesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates unsupported media type response has a 4xx status code
func (o *GetGamificationTemplatesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates unsupported media type response has a 5xx status code
func (o *GetGamificationTemplatesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates unsupported media type response a status code equal to that given
func (o *GetGamificationTemplatesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGamificationTemplatesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationTemplatesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationTemplatesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesTooManyRequests creates a GetGamificationTemplatesTooManyRequests with default headers values
func NewGetGamificationTemplatesTooManyRequests() *GetGamificationTemplatesTooManyRequests {
	return &GetGamificationTemplatesTooManyRequests{}
}

/*
GetGamificationTemplatesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationTemplatesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates too many requests response has a 2xx status code
func (o *GetGamificationTemplatesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates too many requests response has a 3xx status code
func (o *GetGamificationTemplatesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates too many requests response has a 4xx status code
func (o *GetGamificationTemplatesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification templates too many requests response has a 5xx status code
func (o *GetGamificationTemplatesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification templates too many requests response a status code equal to that given
func (o *GetGamificationTemplatesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGamificationTemplatesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationTemplatesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationTemplatesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesInternalServerError creates a GetGamificationTemplatesInternalServerError with default headers values
func NewGetGamificationTemplatesInternalServerError() *GetGamificationTemplatesInternalServerError {
	return &GetGamificationTemplatesInternalServerError{}
}

/*
GetGamificationTemplatesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationTemplatesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates internal server error response has a 2xx status code
func (o *GetGamificationTemplatesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates internal server error response has a 3xx status code
func (o *GetGamificationTemplatesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates internal server error response has a 4xx status code
func (o *GetGamificationTemplatesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification templates internal server error response has a 5xx status code
func (o *GetGamificationTemplatesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification templates internal server error response a status code equal to that given
func (o *GetGamificationTemplatesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGamificationTemplatesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationTemplatesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationTemplatesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesServiceUnavailable creates a GetGamificationTemplatesServiceUnavailable with default headers values
func NewGetGamificationTemplatesServiceUnavailable() *GetGamificationTemplatesServiceUnavailable {
	return &GetGamificationTemplatesServiceUnavailable{}
}

/*
GetGamificationTemplatesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationTemplatesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates service unavailable response has a 2xx status code
func (o *GetGamificationTemplatesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates service unavailable response has a 3xx status code
func (o *GetGamificationTemplatesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates service unavailable response has a 4xx status code
func (o *GetGamificationTemplatesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification templates service unavailable response has a 5xx status code
func (o *GetGamificationTemplatesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification templates service unavailable response a status code equal to that given
func (o *GetGamificationTemplatesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGamificationTemplatesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationTemplatesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationTemplatesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplatesGatewayTimeout creates a GetGamificationTemplatesGatewayTimeout with default headers values
func NewGetGamificationTemplatesGatewayTimeout() *GetGamificationTemplatesGatewayTimeout {
	return &GetGamificationTemplatesGatewayTimeout{}
}

/*
GetGamificationTemplatesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGamificationTemplatesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification templates gateway timeout response has a 2xx status code
func (o *GetGamificationTemplatesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification templates gateway timeout response has a 3xx status code
func (o *GetGamificationTemplatesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification templates gateway timeout response has a 4xx status code
func (o *GetGamificationTemplatesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification templates gateway timeout response has a 5xx status code
func (o *GetGamificationTemplatesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification templates gateway timeout response a status code equal to that given
func (o *GetGamificationTemplatesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGamificationTemplatesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationTemplatesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates][%d] getGamificationTemplatesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationTemplatesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplatesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
