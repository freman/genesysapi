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

// GetGamificationTemplateReader is a Reader for the GetGamificationTemplate structure.
type GetGamificationTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationTemplateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationTemplateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationTemplateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationTemplateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationTemplateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationTemplateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationTemplateOK creates a GetGamificationTemplateOK with default headers values
func NewGetGamificationTemplateOK() *GetGamificationTemplateOK {
	return &GetGamificationTemplateOK{}
}

/*
GetGamificationTemplateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGamificationTemplateOK struct {
	Payload *models.ObjectiveTemplate
}

// IsSuccess returns true when this get gamification template o k response has a 2xx status code
func (o *GetGamificationTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get gamification template o k response has a 3xx status code
func (o *GetGamificationTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template o k response has a 4xx status code
func (o *GetGamificationTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification template o k response has a 5xx status code
func (o *GetGamificationTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template o k response a status code equal to that given
func (o *GetGamificationTemplateOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGamificationTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateOK  %+v", 200, o.Payload)
}

func (o *GetGamificationTemplateOK) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateOK  %+v", 200, o.Payload)
}

func (o *GetGamificationTemplateOK) GetPayload() *models.ObjectiveTemplate {
	return o.Payload
}

func (o *GetGamificationTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectiveTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateBadRequest creates a GetGamificationTemplateBadRequest with default headers values
func NewGetGamificationTemplateBadRequest() *GetGamificationTemplateBadRequest {
	return &GetGamificationTemplateBadRequest{}
}

/*
GetGamificationTemplateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationTemplateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template bad request response has a 2xx status code
func (o *GetGamificationTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template bad request response has a 3xx status code
func (o *GetGamificationTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template bad request response has a 4xx status code
func (o *GetGamificationTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template bad request response has a 5xx status code
func (o *GetGamificationTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template bad request response a status code equal to that given
func (o *GetGamificationTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGamificationTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationTemplateBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateUnauthorized creates a GetGamificationTemplateUnauthorized with default headers values
func NewGetGamificationTemplateUnauthorized() *GetGamificationTemplateUnauthorized {
	return &GetGamificationTemplateUnauthorized{}
}

/*
GetGamificationTemplateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationTemplateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template unauthorized response has a 2xx status code
func (o *GetGamificationTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template unauthorized response has a 3xx status code
func (o *GetGamificationTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template unauthorized response has a 4xx status code
func (o *GetGamificationTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template unauthorized response has a 5xx status code
func (o *GetGamificationTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template unauthorized response a status code equal to that given
func (o *GetGamificationTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGamificationTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationTemplateUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationTemplateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateForbidden creates a GetGamificationTemplateForbidden with default headers values
func NewGetGamificationTemplateForbidden() *GetGamificationTemplateForbidden {
	return &GetGamificationTemplateForbidden{}
}

/*
GetGamificationTemplateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationTemplateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template forbidden response has a 2xx status code
func (o *GetGamificationTemplateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template forbidden response has a 3xx status code
func (o *GetGamificationTemplateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template forbidden response has a 4xx status code
func (o *GetGamificationTemplateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template forbidden response has a 5xx status code
func (o *GetGamificationTemplateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template forbidden response a status code equal to that given
func (o *GetGamificationTemplateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGamificationTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationTemplateForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationTemplateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateNotFound creates a GetGamificationTemplateNotFound with default headers values
func NewGetGamificationTemplateNotFound() *GetGamificationTemplateNotFound {
	return &GetGamificationTemplateNotFound{}
}

/*
GetGamificationTemplateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGamificationTemplateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template not found response has a 2xx status code
func (o *GetGamificationTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template not found response has a 3xx status code
func (o *GetGamificationTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template not found response has a 4xx status code
func (o *GetGamificationTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template not found response has a 5xx status code
func (o *GetGamificationTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template not found response a status code equal to that given
func (o *GetGamificationTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGamificationTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationTemplateNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateRequestTimeout creates a GetGamificationTemplateRequestTimeout with default headers values
func NewGetGamificationTemplateRequestTimeout() *GetGamificationTemplateRequestTimeout {
	return &GetGamificationTemplateRequestTimeout{}
}

/*
GetGamificationTemplateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationTemplateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template request timeout response has a 2xx status code
func (o *GetGamificationTemplateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template request timeout response has a 3xx status code
func (o *GetGamificationTemplateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template request timeout response has a 4xx status code
func (o *GetGamificationTemplateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template request timeout response has a 5xx status code
func (o *GetGamificationTemplateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template request timeout response a status code equal to that given
func (o *GetGamificationTemplateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGamificationTemplateRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationTemplateRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationTemplateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateRequestEntityTooLarge creates a GetGamificationTemplateRequestEntityTooLarge with default headers values
func NewGetGamificationTemplateRequestEntityTooLarge() *GetGamificationTemplateRequestEntityTooLarge {
	return &GetGamificationTemplateRequestEntityTooLarge{}
}

/*
GetGamificationTemplateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetGamificationTemplateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template request entity too large response has a 2xx status code
func (o *GetGamificationTemplateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template request entity too large response has a 3xx status code
func (o *GetGamificationTemplateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template request entity too large response has a 4xx status code
func (o *GetGamificationTemplateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template request entity too large response has a 5xx status code
func (o *GetGamificationTemplateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template request entity too large response a status code equal to that given
func (o *GetGamificationTemplateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGamificationTemplateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationTemplateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationTemplateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateUnsupportedMediaType creates a GetGamificationTemplateUnsupportedMediaType with default headers values
func NewGetGamificationTemplateUnsupportedMediaType() *GetGamificationTemplateUnsupportedMediaType {
	return &GetGamificationTemplateUnsupportedMediaType{}
}

/*
GetGamificationTemplateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationTemplateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template unsupported media type response has a 2xx status code
func (o *GetGamificationTemplateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template unsupported media type response has a 3xx status code
func (o *GetGamificationTemplateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template unsupported media type response has a 4xx status code
func (o *GetGamificationTemplateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template unsupported media type response has a 5xx status code
func (o *GetGamificationTemplateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template unsupported media type response a status code equal to that given
func (o *GetGamificationTemplateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGamificationTemplateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationTemplateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationTemplateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateTooManyRequests creates a GetGamificationTemplateTooManyRequests with default headers values
func NewGetGamificationTemplateTooManyRequests() *GetGamificationTemplateTooManyRequests {
	return &GetGamificationTemplateTooManyRequests{}
}

/*
GetGamificationTemplateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationTemplateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template too many requests response has a 2xx status code
func (o *GetGamificationTemplateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template too many requests response has a 3xx status code
func (o *GetGamificationTemplateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template too many requests response has a 4xx status code
func (o *GetGamificationTemplateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get gamification template too many requests response has a 5xx status code
func (o *GetGamificationTemplateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get gamification template too many requests response a status code equal to that given
func (o *GetGamificationTemplateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGamificationTemplateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationTemplateTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationTemplateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateInternalServerError creates a GetGamificationTemplateInternalServerError with default headers values
func NewGetGamificationTemplateInternalServerError() *GetGamificationTemplateInternalServerError {
	return &GetGamificationTemplateInternalServerError{}
}

/*
GetGamificationTemplateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationTemplateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template internal server error response has a 2xx status code
func (o *GetGamificationTemplateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template internal server error response has a 3xx status code
func (o *GetGamificationTemplateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template internal server error response has a 4xx status code
func (o *GetGamificationTemplateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification template internal server error response has a 5xx status code
func (o *GetGamificationTemplateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification template internal server error response a status code equal to that given
func (o *GetGamificationTemplateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGamificationTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationTemplateInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateServiceUnavailable creates a GetGamificationTemplateServiceUnavailable with default headers values
func NewGetGamificationTemplateServiceUnavailable() *GetGamificationTemplateServiceUnavailable {
	return &GetGamificationTemplateServiceUnavailable{}
}

/*
GetGamificationTemplateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationTemplateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template service unavailable response has a 2xx status code
func (o *GetGamificationTemplateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template service unavailable response has a 3xx status code
func (o *GetGamificationTemplateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template service unavailable response has a 4xx status code
func (o *GetGamificationTemplateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification template service unavailable response has a 5xx status code
func (o *GetGamificationTemplateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification template service unavailable response a status code equal to that given
func (o *GetGamificationTemplateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGamificationTemplateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationTemplateServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationTemplateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationTemplateGatewayTimeout creates a GetGamificationTemplateGatewayTimeout with default headers values
func NewGetGamificationTemplateGatewayTimeout() *GetGamificationTemplateGatewayTimeout {
	return &GetGamificationTemplateGatewayTimeout{}
}

/*
GetGamificationTemplateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGamificationTemplateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get gamification template gateway timeout response has a 2xx status code
func (o *GetGamificationTemplateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get gamification template gateway timeout response has a 3xx status code
func (o *GetGamificationTemplateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get gamification template gateway timeout response has a 4xx status code
func (o *GetGamificationTemplateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get gamification template gateway timeout response has a 5xx status code
func (o *GetGamificationTemplateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get gamification template gateway timeout response a status code equal to that given
func (o *GetGamificationTemplateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGamificationTemplateGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationTemplateGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/gamification/templates/{templateId}][%d] getGamificationTemplateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationTemplateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationTemplateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
