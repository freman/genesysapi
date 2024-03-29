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

// GetIntegrationsActionTemplateReader is a Reader for the GetIntegrationsActionTemplate structure.
type GetIntegrationsActionTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsActionTemplateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionTemplateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionTemplateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionTemplateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionTemplateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionTemplateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionTemplateOK creates a GetIntegrationsActionTemplateOK with default headers values
func NewGetIntegrationsActionTemplateOK() *GetIntegrationsActionTemplateOK {
	return &GetIntegrationsActionTemplateOK{}
}

/*
GetIntegrationsActionTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsActionTemplateOK struct {
	Payload string
}

// IsSuccess returns true when this get integrations action template o k response has a 2xx status code
func (o *GetIntegrationsActionTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations action template o k response has a 3xx status code
func (o *GetIntegrationsActionTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template o k response has a 4xx status code
func (o *GetIntegrationsActionTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations action template o k response has a 5xx status code
func (o *GetIntegrationsActionTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template o k response a status code equal to that given
func (o *GetIntegrationsActionTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsActionTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionTemplateOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateBadRequest creates a GetIntegrationsActionTemplateBadRequest with default headers values
func NewGetIntegrationsActionTemplateBadRequest() *GetIntegrationsActionTemplateBadRequest {
	return &GetIntegrationsActionTemplateBadRequest{}
}

/*
GetIntegrationsActionTemplateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionTemplateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template bad request response has a 2xx status code
func (o *GetIntegrationsActionTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template bad request response has a 3xx status code
func (o *GetIntegrationsActionTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template bad request response has a 4xx status code
func (o *GetIntegrationsActionTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template bad request response has a 5xx status code
func (o *GetIntegrationsActionTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template bad request response a status code equal to that given
func (o *GetIntegrationsActionTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsActionTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionTemplateBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateUnauthorized creates a GetIntegrationsActionTemplateUnauthorized with default headers values
func NewGetIntegrationsActionTemplateUnauthorized() *GetIntegrationsActionTemplateUnauthorized {
	return &GetIntegrationsActionTemplateUnauthorized{}
}

/*
GetIntegrationsActionTemplateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionTemplateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template unauthorized response has a 2xx status code
func (o *GetIntegrationsActionTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template unauthorized response has a 3xx status code
func (o *GetIntegrationsActionTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template unauthorized response has a 4xx status code
func (o *GetIntegrationsActionTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template unauthorized response has a 5xx status code
func (o *GetIntegrationsActionTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template unauthorized response a status code equal to that given
func (o *GetIntegrationsActionTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsActionTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionTemplateUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionTemplateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateForbidden creates a GetIntegrationsActionTemplateForbidden with default headers values
func NewGetIntegrationsActionTemplateForbidden() *GetIntegrationsActionTemplateForbidden {
	return &GetIntegrationsActionTemplateForbidden{}
}

/*
GetIntegrationsActionTemplateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionTemplateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template forbidden response has a 2xx status code
func (o *GetIntegrationsActionTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template forbidden response has a 3xx status code
func (o *GetIntegrationsActionTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template forbidden response has a 4xx status code
func (o *GetIntegrationsActionTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template forbidden response has a 5xx status code
func (o *GetIntegrationsActionTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template forbidden response a status code equal to that given
func (o *GetIntegrationsActionTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsActionTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionTemplateForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionTemplateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateNotFound creates a GetIntegrationsActionTemplateNotFound with default headers values
func NewGetIntegrationsActionTemplateNotFound() *GetIntegrationsActionTemplateNotFound {
	return &GetIntegrationsActionTemplateNotFound{}
}

/*
GetIntegrationsActionTemplateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionTemplateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template not found response has a 2xx status code
func (o *GetIntegrationsActionTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template not found response has a 3xx status code
func (o *GetIntegrationsActionTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template not found response has a 4xx status code
func (o *GetIntegrationsActionTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template not found response has a 5xx status code
func (o *GetIntegrationsActionTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template not found response a status code equal to that given
func (o *GetIntegrationsActionTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsActionTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionTemplateNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateRequestTimeout creates a GetIntegrationsActionTemplateRequestTimeout with default headers values
func NewGetIntegrationsActionTemplateRequestTimeout() *GetIntegrationsActionTemplateRequestTimeout {
	return &GetIntegrationsActionTemplateRequestTimeout{}
}

/*
GetIntegrationsActionTemplateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsActionTemplateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template request timeout response has a 2xx status code
func (o *GetIntegrationsActionTemplateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template request timeout response has a 3xx status code
func (o *GetIntegrationsActionTemplateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template request timeout response has a 4xx status code
func (o *GetIntegrationsActionTemplateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template request timeout response has a 5xx status code
func (o *GetIntegrationsActionTemplateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template request timeout response a status code equal to that given
func (o *GetIntegrationsActionTemplateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsActionTemplateRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionTemplateRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionTemplateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateRequestEntityTooLarge creates a GetIntegrationsActionTemplateRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionTemplateRequestEntityTooLarge() *GetIntegrationsActionTemplateRequestEntityTooLarge {
	return &GetIntegrationsActionTemplateRequestEntityTooLarge{}
}

/*
GetIntegrationsActionTemplateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsActionTemplateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template request entity too large response has a 2xx status code
func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template request entity too large response has a 3xx status code
func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template request entity too large response has a 4xx status code
func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template request entity too large response has a 5xx status code
func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template request entity too large response a status code equal to that given
func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateUnsupportedMediaType creates a GetIntegrationsActionTemplateUnsupportedMediaType with default headers values
func NewGetIntegrationsActionTemplateUnsupportedMediaType() *GetIntegrationsActionTemplateUnsupportedMediaType {
	return &GetIntegrationsActionTemplateUnsupportedMediaType{}
}

/*
GetIntegrationsActionTemplateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionTemplateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template unsupported media type response has a 2xx status code
func (o *GetIntegrationsActionTemplateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template unsupported media type response has a 3xx status code
func (o *GetIntegrationsActionTemplateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template unsupported media type response has a 4xx status code
func (o *GetIntegrationsActionTemplateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template unsupported media type response has a 5xx status code
func (o *GetIntegrationsActionTemplateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template unsupported media type response a status code equal to that given
func (o *GetIntegrationsActionTemplateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsActionTemplateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionTemplateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionTemplateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateTooManyRequests creates a GetIntegrationsActionTemplateTooManyRequests with default headers values
func NewGetIntegrationsActionTemplateTooManyRequests() *GetIntegrationsActionTemplateTooManyRequests {
	return &GetIntegrationsActionTemplateTooManyRequests{}
}

/*
GetIntegrationsActionTemplateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsActionTemplateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template too many requests response has a 2xx status code
func (o *GetIntegrationsActionTemplateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template too many requests response has a 3xx status code
func (o *GetIntegrationsActionTemplateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template too many requests response has a 4xx status code
func (o *GetIntegrationsActionTemplateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations action template too many requests response has a 5xx status code
func (o *GetIntegrationsActionTemplateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations action template too many requests response a status code equal to that given
func (o *GetIntegrationsActionTemplateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsActionTemplateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionTemplateTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionTemplateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateInternalServerError creates a GetIntegrationsActionTemplateInternalServerError with default headers values
func NewGetIntegrationsActionTemplateInternalServerError() *GetIntegrationsActionTemplateInternalServerError {
	return &GetIntegrationsActionTemplateInternalServerError{}
}

/*
GetIntegrationsActionTemplateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionTemplateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template internal server error response has a 2xx status code
func (o *GetIntegrationsActionTemplateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template internal server error response has a 3xx status code
func (o *GetIntegrationsActionTemplateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template internal server error response has a 4xx status code
func (o *GetIntegrationsActionTemplateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations action template internal server error response has a 5xx status code
func (o *GetIntegrationsActionTemplateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations action template internal server error response a status code equal to that given
func (o *GetIntegrationsActionTemplateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsActionTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionTemplateInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateServiceUnavailable creates a GetIntegrationsActionTemplateServiceUnavailable with default headers values
func NewGetIntegrationsActionTemplateServiceUnavailable() *GetIntegrationsActionTemplateServiceUnavailable {
	return &GetIntegrationsActionTemplateServiceUnavailable{}
}

/*
GetIntegrationsActionTemplateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionTemplateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template service unavailable response has a 2xx status code
func (o *GetIntegrationsActionTemplateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template service unavailable response has a 3xx status code
func (o *GetIntegrationsActionTemplateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template service unavailable response has a 4xx status code
func (o *GetIntegrationsActionTemplateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations action template service unavailable response has a 5xx status code
func (o *GetIntegrationsActionTemplateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations action template service unavailable response a status code equal to that given
func (o *GetIntegrationsActionTemplateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsActionTemplateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionTemplateServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionTemplateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionTemplateGatewayTimeout creates a GetIntegrationsActionTemplateGatewayTimeout with default headers values
func NewGetIntegrationsActionTemplateGatewayTimeout() *GetIntegrationsActionTemplateGatewayTimeout {
	return &GetIntegrationsActionTemplateGatewayTimeout{}
}

/*
GetIntegrationsActionTemplateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsActionTemplateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations action template gateway timeout response has a 2xx status code
func (o *GetIntegrationsActionTemplateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations action template gateway timeout response has a 3xx status code
func (o *GetIntegrationsActionTemplateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations action template gateway timeout response has a 4xx status code
func (o *GetIntegrationsActionTemplateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations action template gateway timeout response has a 5xx status code
func (o *GetIntegrationsActionTemplateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations action template gateway timeout response a status code equal to that given
func (o *GetIntegrationsActionTemplateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsActionTemplateGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionTemplateGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/templates/{fileName}][%d] getIntegrationsActionTemplateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionTemplateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionTemplateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
