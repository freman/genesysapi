// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetArchitectSystempromptResourceReader is a Reader for the GetArchitectSystempromptResource structure.
type GetArchitectSystempromptResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectSystempromptResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectSystempromptResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectSystempromptResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectSystempromptResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectSystempromptResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectSystempromptResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectSystempromptResourceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectSystempromptResourceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectSystempromptResourceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectSystempromptResourceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectSystempromptResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectSystempromptResourceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectSystempromptResourceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectSystempromptResourceOK creates a GetArchitectSystempromptResourceOK with default headers values
func NewGetArchitectSystempromptResourceOK() *GetArchitectSystempromptResourceOK {
	return &GetArchitectSystempromptResourceOK{}
}

/*
GetArchitectSystempromptResourceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectSystempromptResourceOK struct {
	Payload *models.SystemPromptAsset
}

// IsSuccess returns true when this get architect systemprompt resource o k response has a 2xx status code
func (o *GetArchitectSystempromptResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect systemprompt resource o k response has a 3xx status code
func (o *GetArchitectSystempromptResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource o k response has a 4xx status code
func (o *GetArchitectSystempromptResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompt resource o k response has a 5xx status code
func (o *GetArchitectSystempromptResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource o k response a status code equal to that given
func (o *GetArchitectSystempromptResourceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectSystempromptResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSystempromptResourceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSystempromptResourceOK) GetPayload() *models.SystemPromptAsset {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemPromptAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceBadRequest creates a GetArchitectSystempromptResourceBadRequest with default headers values
func NewGetArchitectSystempromptResourceBadRequest() *GetArchitectSystempromptResourceBadRequest {
	return &GetArchitectSystempromptResourceBadRequest{}
}

/*
GetArchitectSystempromptResourceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectSystempromptResourceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource bad request response has a 2xx status code
func (o *GetArchitectSystempromptResourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource bad request response has a 3xx status code
func (o *GetArchitectSystempromptResourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource bad request response has a 4xx status code
func (o *GetArchitectSystempromptResourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource bad request response has a 5xx status code
func (o *GetArchitectSystempromptResourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource bad request response a status code equal to that given
func (o *GetArchitectSystempromptResourceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectSystempromptResourceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSystempromptResourceBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSystempromptResourceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceUnauthorized creates a GetArchitectSystempromptResourceUnauthorized with default headers values
func NewGetArchitectSystempromptResourceUnauthorized() *GetArchitectSystempromptResourceUnauthorized {
	return &GetArchitectSystempromptResourceUnauthorized{}
}

/*
GetArchitectSystempromptResourceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectSystempromptResourceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource unauthorized response has a 2xx status code
func (o *GetArchitectSystempromptResourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource unauthorized response has a 3xx status code
func (o *GetArchitectSystempromptResourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource unauthorized response has a 4xx status code
func (o *GetArchitectSystempromptResourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource unauthorized response has a 5xx status code
func (o *GetArchitectSystempromptResourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource unauthorized response a status code equal to that given
func (o *GetArchitectSystempromptResourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectSystempromptResourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSystempromptResourceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSystempromptResourceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceForbidden creates a GetArchitectSystempromptResourceForbidden with default headers values
func NewGetArchitectSystempromptResourceForbidden() *GetArchitectSystempromptResourceForbidden {
	return &GetArchitectSystempromptResourceForbidden{}
}

/*
GetArchitectSystempromptResourceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectSystempromptResourceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource forbidden response has a 2xx status code
func (o *GetArchitectSystempromptResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource forbidden response has a 3xx status code
func (o *GetArchitectSystempromptResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource forbidden response has a 4xx status code
func (o *GetArchitectSystempromptResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource forbidden response has a 5xx status code
func (o *GetArchitectSystempromptResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource forbidden response a status code equal to that given
func (o *GetArchitectSystempromptResourceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectSystempromptResourceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSystempromptResourceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSystempromptResourceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceNotFound creates a GetArchitectSystempromptResourceNotFound with default headers values
func NewGetArchitectSystempromptResourceNotFound() *GetArchitectSystempromptResourceNotFound {
	return &GetArchitectSystempromptResourceNotFound{}
}

/*
GetArchitectSystempromptResourceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectSystempromptResourceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource not found response has a 2xx status code
func (o *GetArchitectSystempromptResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource not found response has a 3xx status code
func (o *GetArchitectSystempromptResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource not found response has a 4xx status code
func (o *GetArchitectSystempromptResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource not found response has a 5xx status code
func (o *GetArchitectSystempromptResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource not found response a status code equal to that given
func (o *GetArchitectSystempromptResourceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectSystempromptResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSystempromptResourceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSystempromptResourceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceRequestTimeout creates a GetArchitectSystempromptResourceRequestTimeout with default headers values
func NewGetArchitectSystempromptResourceRequestTimeout() *GetArchitectSystempromptResourceRequestTimeout {
	return &GetArchitectSystempromptResourceRequestTimeout{}
}

/*
GetArchitectSystempromptResourceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectSystempromptResourceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource request timeout response has a 2xx status code
func (o *GetArchitectSystempromptResourceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource request timeout response has a 3xx status code
func (o *GetArchitectSystempromptResourceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource request timeout response has a 4xx status code
func (o *GetArchitectSystempromptResourceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource request timeout response has a 5xx status code
func (o *GetArchitectSystempromptResourceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource request timeout response a status code equal to that given
func (o *GetArchitectSystempromptResourceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectSystempromptResourceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectSystempromptResourceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectSystempromptResourceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceRequestEntityTooLarge creates a GetArchitectSystempromptResourceRequestEntityTooLarge with default headers values
func NewGetArchitectSystempromptResourceRequestEntityTooLarge() *GetArchitectSystempromptResourceRequestEntityTooLarge {
	return &GetArchitectSystempromptResourceRequestEntityTooLarge{}
}

/*
GetArchitectSystempromptResourceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetArchitectSystempromptResourceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource request entity too large response has a 2xx status code
func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource request entity too large response has a 3xx status code
func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource request entity too large response has a 4xx status code
func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource request entity too large response has a 5xx status code
func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource request entity too large response a status code equal to that given
func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceUnsupportedMediaType creates a GetArchitectSystempromptResourceUnsupportedMediaType with default headers values
func NewGetArchitectSystempromptResourceUnsupportedMediaType() *GetArchitectSystempromptResourceUnsupportedMediaType {
	return &GetArchitectSystempromptResourceUnsupportedMediaType{}
}

/*
GetArchitectSystempromptResourceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectSystempromptResourceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource unsupported media type response has a 2xx status code
func (o *GetArchitectSystempromptResourceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource unsupported media type response has a 3xx status code
func (o *GetArchitectSystempromptResourceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource unsupported media type response has a 4xx status code
func (o *GetArchitectSystempromptResourceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource unsupported media type response has a 5xx status code
func (o *GetArchitectSystempromptResourceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource unsupported media type response a status code equal to that given
func (o *GetArchitectSystempromptResourceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectSystempromptResourceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSystempromptResourceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSystempromptResourceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceTooManyRequests creates a GetArchitectSystempromptResourceTooManyRequests with default headers values
func NewGetArchitectSystempromptResourceTooManyRequests() *GetArchitectSystempromptResourceTooManyRequests {
	return &GetArchitectSystempromptResourceTooManyRequests{}
}

/*
GetArchitectSystempromptResourceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectSystempromptResourceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource too many requests response has a 2xx status code
func (o *GetArchitectSystempromptResourceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource too many requests response has a 3xx status code
func (o *GetArchitectSystempromptResourceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource too many requests response has a 4xx status code
func (o *GetArchitectSystempromptResourceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect systemprompt resource too many requests response has a 5xx status code
func (o *GetArchitectSystempromptResourceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect systemprompt resource too many requests response a status code equal to that given
func (o *GetArchitectSystempromptResourceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectSystempromptResourceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSystempromptResourceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSystempromptResourceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceInternalServerError creates a GetArchitectSystempromptResourceInternalServerError with default headers values
func NewGetArchitectSystempromptResourceInternalServerError() *GetArchitectSystempromptResourceInternalServerError {
	return &GetArchitectSystempromptResourceInternalServerError{}
}

/*
GetArchitectSystempromptResourceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectSystempromptResourceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource internal server error response has a 2xx status code
func (o *GetArchitectSystempromptResourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource internal server error response has a 3xx status code
func (o *GetArchitectSystempromptResourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource internal server error response has a 4xx status code
func (o *GetArchitectSystempromptResourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompt resource internal server error response has a 5xx status code
func (o *GetArchitectSystempromptResourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect systemprompt resource internal server error response a status code equal to that given
func (o *GetArchitectSystempromptResourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectSystempromptResourceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSystempromptResourceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSystempromptResourceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceServiceUnavailable creates a GetArchitectSystempromptResourceServiceUnavailable with default headers values
func NewGetArchitectSystempromptResourceServiceUnavailable() *GetArchitectSystempromptResourceServiceUnavailable {
	return &GetArchitectSystempromptResourceServiceUnavailable{}
}

/*
GetArchitectSystempromptResourceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectSystempromptResourceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource service unavailable response has a 2xx status code
func (o *GetArchitectSystempromptResourceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource service unavailable response has a 3xx status code
func (o *GetArchitectSystempromptResourceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource service unavailable response has a 4xx status code
func (o *GetArchitectSystempromptResourceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompt resource service unavailable response has a 5xx status code
func (o *GetArchitectSystempromptResourceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect systemprompt resource service unavailable response a status code equal to that given
func (o *GetArchitectSystempromptResourceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectSystempromptResourceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSystempromptResourceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSystempromptResourceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourceGatewayTimeout creates a GetArchitectSystempromptResourceGatewayTimeout with default headers values
func NewGetArchitectSystempromptResourceGatewayTimeout() *GetArchitectSystempromptResourceGatewayTimeout {
	return &GetArchitectSystempromptResourceGatewayTimeout{}
}

/*
GetArchitectSystempromptResourceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectSystempromptResourceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect systemprompt resource gateway timeout response has a 2xx status code
func (o *GetArchitectSystempromptResourceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect systemprompt resource gateway timeout response has a 3xx status code
func (o *GetArchitectSystempromptResourceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect systemprompt resource gateway timeout response has a 4xx status code
func (o *GetArchitectSystempromptResourceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect systemprompt resource gateway timeout response has a 5xx status code
func (o *GetArchitectSystempromptResourceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect systemprompt resource gateway timeout response a status code equal to that given
func (o *GetArchitectSystempromptResourceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectSystempromptResourceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSystempromptResourceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources/{languageCode}][%d] getArchitectSystempromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSystempromptResourceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
