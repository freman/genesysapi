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

// GetArchitectPromptResourceReader is a Reader for the GetArchitectPromptResource structure.
type GetArchitectPromptResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectPromptResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectPromptResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectPromptResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectPromptResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectPromptResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectPromptResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectPromptResourceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectPromptResourceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectPromptResourceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectPromptResourceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectPromptResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectPromptResourceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectPromptResourceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectPromptResourceOK creates a GetArchitectPromptResourceOK with default headers values
func NewGetArchitectPromptResourceOK() *GetArchitectPromptResourceOK {
	return &GetArchitectPromptResourceOK{}
}

/*
GetArchitectPromptResourceOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectPromptResourceOK struct {
	Payload *models.PromptAsset
}

// IsSuccess returns true when this get architect prompt resource o k response has a 2xx status code
func (o *GetArchitectPromptResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect prompt resource o k response has a 3xx status code
func (o *GetArchitectPromptResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource o k response has a 4xx status code
func (o *GetArchitectPromptResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resource o k response has a 5xx status code
func (o *GetArchitectPromptResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource o k response a status code equal to that given
func (o *GetArchitectPromptResourceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectPromptResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceOK  %+v", 200, o.Payload)
}

func (o *GetArchitectPromptResourceOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceOK  %+v", 200, o.Payload)
}

func (o *GetArchitectPromptResourceOK) GetPayload() *models.PromptAsset {
	return o.Payload
}

func (o *GetArchitectPromptResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromptAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceBadRequest creates a GetArchitectPromptResourceBadRequest with default headers values
func NewGetArchitectPromptResourceBadRequest() *GetArchitectPromptResourceBadRequest {
	return &GetArchitectPromptResourceBadRequest{}
}

/*
GetArchitectPromptResourceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectPromptResourceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource bad request response has a 2xx status code
func (o *GetArchitectPromptResourceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource bad request response has a 3xx status code
func (o *GetArchitectPromptResourceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource bad request response has a 4xx status code
func (o *GetArchitectPromptResourceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource bad request response has a 5xx status code
func (o *GetArchitectPromptResourceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource bad request response a status code equal to that given
func (o *GetArchitectPromptResourceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectPromptResourceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectPromptResourceBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectPromptResourceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceUnauthorized creates a GetArchitectPromptResourceUnauthorized with default headers values
func NewGetArchitectPromptResourceUnauthorized() *GetArchitectPromptResourceUnauthorized {
	return &GetArchitectPromptResourceUnauthorized{}
}

/*
GetArchitectPromptResourceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectPromptResourceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource unauthorized response has a 2xx status code
func (o *GetArchitectPromptResourceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource unauthorized response has a 3xx status code
func (o *GetArchitectPromptResourceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource unauthorized response has a 4xx status code
func (o *GetArchitectPromptResourceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource unauthorized response has a 5xx status code
func (o *GetArchitectPromptResourceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource unauthorized response a status code equal to that given
func (o *GetArchitectPromptResourceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectPromptResourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectPromptResourceUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectPromptResourceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceForbidden creates a GetArchitectPromptResourceForbidden with default headers values
func NewGetArchitectPromptResourceForbidden() *GetArchitectPromptResourceForbidden {
	return &GetArchitectPromptResourceForbidden{}
}

/*
GetArchitectPromptResourceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectPromptResourceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource forbidden response has a 2xx status code
func (o *GetArchitectPromptResourceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource forbidden response has a 3xx status code
func (o *GetArchitectPromptResourceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource forbidden response has a 4xx status code
func (o *GetArchitectPromptResourceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource forbidden response has a 5xx status code
func (o *GetArchitectPromptResourceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource forbidden response a status code equal to that given
func (o *GetArchitectPromptResourceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectPromptResourceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectPromptResourceForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectPromptResourceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceNotFound creates a GetArchitectPromptResourceNotFound with default headers values
func NewGetArchitectPromptResourceNotFound() *GetArchitectPromptResourceNotFound {
	return &GetArchitectPromptResourceNotFound{}
}

/*
GetArchitectPromptResourceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectPromptResourceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource not found response has a 2xx status code
func (o *GetArchitectPromptResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource not found response has a 3xx status code
func (o *GetArchitectPromptResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource not found response has a 4xx status code
func (o *GetArchitectPromptResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource not found response has a 5xx status code
func (o *GetArchitectPromptResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource not found response a status code equal to that given
func (o *GetArchitectPromptResourceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectPromptResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectPromptResourceNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectPromptResourceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceRequestTimeout creates a GetArchitectPromptResourceRequestTimeout with default headers values
func NewGetArchitectPromptResourceRequestTimeout() *GetArchitectPromptResourceRequestTimeout {
	return &GetArchitectPromptResourceRequestTimeout{}
}

/*
GetArchitectPromptResourceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectPromptResourceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource request timeout response has a 2xx status code
func (o *GetArchitectPromptResourceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource request timeout response has a 3xx status code
func (o *GetArchitectPromptResourceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource request timeout response has a 4xx status code
func (o *GetArchitectPromptResourceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource request timeout response has a 5xx status code
func (o *GetArchitectPromptResourceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource request timeout response a status code equal to that given
func (o *GetArchitectPromptResourceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectPromptResourceRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectPromptResourceRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectPromptResourceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceRequestEntityTooLarge creates a GetArchitectPromptResourceRequestEntityTooLarge with default headers values
func NewGetArchitectPromptResourceRequestEntityTooLarge() *GetArchitectPromptResourceRequestEntityTooLarge {
	return &GetArchitectPromptResourceRequestEntityTooLarge{}
}

/*
GetArchitectPromptResourceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetArchitectPromptResourceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource request entity too large response has a 2xx status code
func (o *GetArchitectPromptResourceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource request entity too large response has a 3xx status code
func (o *GetArchitectPromptResourceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource request entity too large response has a 4xx status code
func (o *GetArchitectPromptResourceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource request entity too large response has a 5xx status code
func (o *GetArchitectPromptResourceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource request entity too large response a status code equal to that given
func (o *GetArchitectPromptResourceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceUnsupportedMediaType creates a GetArchitectPromptResourceUnsupportedMediaType with default headers values
func NewGetArchitectPromptResourceUnsupportedMediaType() *GetArchitectPromptResourceUnsupportedMediaType {
	return &GetArchitectPromptResourceUnsupportedMediaType{}
}

/*
GetArchitectPromptResourceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectPromptResourceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource unsupported media type response has a 2xx status code
func (o *GetArchitectPromptResourceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource unsupported media type response has a 3xx status code
func (o *GetArchitectPromptResourceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource unsupported media type response has a 4xx status code
func (o *GetArchitectPromptResourceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource unsupported media type response has a 5xx status code
func (o *GetArchitectPromptResourceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource unsupported media type response a status code equal to that given
func (o *GetArchitectPromptResourceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceTooManyRequests creates a GetArchitectPromptResourceTooManyRequests with default headers values
func NewGetArchitectPromptResourceTooManyRequests() *GetArchitectPromptResourceTooManyRequests {
	return &GetArchitectPromptResourceTooManyRequests{}
}

/*
GetArchitectPromptResourceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectPromptResourceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource too many requests response has a 2xx status code
func (o *GetArchitectPromptResourceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource too many requests response has a 3xx status code
func (o *GetArchitectPromptResourceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource too many requests response has a 4xx status code
func (o *GetArchitectPromptResourceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect prompt resource too many requests response has a 5xx status code
func (o *GetArchitectPromptResourceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect prompt resource too many requests response a status code equal to that given
func (o *GetArchitectPromptResourceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectPromptResourceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectPromptResourceTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectPromptResourceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceInternalServerError creates a GetArchitectPromptResourceInternalServerError with default headers values
func NewGetArchitectPromptResourceInternalServerError() *GetArchitectPromptResourceInternalServerError {
	return &GetArchitectPromptResourceInternalServerError{}
}

/*
GetArchitectPromptResourceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectPromptResourceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource internal server error response has a 2xx status code
func (o *GetArchitectPromptResourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource internal server error response has a 3xx status code
func (o *GetArchitectPromptResourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource internal server error response has a 4xx status code
func (o *GetArchitectPromptResourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resource internal server error response has a 5xx status code
func (o *GetArchitectPromptResourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect prompt resource internal server error response a status code equal to that given
func (o *GetArchitectPromptResourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectPromptResourceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectPromptResourceInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectPromptResourceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceServiceUnavailable creates a GetArchitectPromptResourceServiceUnavailable with default headers values
func NewGetArchitectPromptResourceServiceUnavailable() *GetArchitectPromptResourceServiceUnavailable {
	return &GetArchitectPromptResourceServiceUnavailable{}
}

/*
GetArchitectPromptResourceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectPromptResourceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource service unavailable response has a 2xx status code
func (o *GetArchitectPromptResourceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource service unavailable response has a 3xx status code
func (o *GetArchitectPromptResourceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource service unavailable response has a 4xx status code
func (o *GetArchitectPromptResourceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resource service unavailable response has a 5xx status code
func (o *GetArchitectPromptResourceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect prompt resource service unavailable response a status code equal to that given
func (o *GetArchitectPromptResourceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectPromptResourceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectPromptResourceServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectPromptResourceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceGatewayTimeout creates a GetArchitectPromptResourceGatewayTimeout with default headers values
func NewGetArchitectPromptResourceGatewayTimeout() *GetArchitectPromptResourceGatewayTimeout {
	return &GetArchitectPromptResourceGatewayTimeout{}
}

/*
GetArchitectPromptResourceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectPromptResourceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect prompt resource gateway timeout response has a 2xx status code
func (o *GetArchitectPromptResourceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect prompt resource gateway timeout response has a 3xx status code
func (o *GetArchitectPromptResourceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect prompt resource gateway timeout response has a 4xx status code
func (o *GetArchitectPromptResourceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect prompt resource gateway timeout response has a 5xx status code
func (o *GetArchitectPromptResourceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect prompt resource gateway timeout response a status code equal to that given
func (o *GetArchitectPromptResourceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectPromptResourceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectPromptResourceGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectPromptResourceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
