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

// GetIntegrationReader is a Reader for the GetIntegration structure.
type GetIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationOK creates a GetIntegrationOK with default headers values
func NewGetIntegrationOK() *GetIntegrationOK {
	return &GetIntegrationOK{}
}

/*
GetIntegrationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationOK struct {
	Payload *models.Integration
}

// IsSuccess returns true when this get integration o k response has a 2xx status code
func (o *GetIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integration o k response has a 3xx status code
func (o *GetIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration o k response has a 4xx status code
func (o *GetIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration o k response has a 5xx status code
func (o *GetIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration o k response a status code equal to that given
func (o *GetIntegrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationOK) GetPayload() *models.Integration {
	return o.Payload
}

func (o *GetIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationBadRequest creates a GetIntegrationBadRequest with default headers values
func NewGetIntegrationBadRequest() *GetIntegrationBadRequest {
	return &GetIntegrationBadRequest{}
}

/*
GetIntegrationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration bad request response has a 2xx status code
func (o *GetIntegrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration bad request response has a 3xx status code
func (o *GetIntegrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration bad request response has a 4xx status code
func (o *GetIntegrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration bad request response has a 5xx status code
func (o *GetIntegrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration bad request response a status code equal to that given
func (o *GetIntegrationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationUnauthorized creates a GetIntegrationUnauthorized with default headers values
func NewGetIntegrationUnauthorized() *GetIntegrationUnauthorized {
	return &GetIntegrationUnauthorized{}
}

/*
GetIntegrationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration unauthorized response has a 2xx status code
func (o *GetIntegrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration unauthorized response has a 3xx status code
func (o *GetIntegrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration unauthorized response has a 4xx status code
func (o *GetIntegrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration unauthorized response has a 5xx status code
func (o *GetIntegrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration unauthorized response a status code equal to that given
func (o *GetIntegrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationForbidden creates a GetIntegrationForbidden with default headers values
func NewGetIntegrationForbidden() *GetIntegrationForbidden {
	return &GetIntegrationForbidden{}
}

/*
GetIntegrationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration forbidden response has a 2xx status code
func (o *GetIntegrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration forbidden response has a 3xx status code
func (o *GetIntegrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration forbidden response has a 4xx status code
func (o *GetIntegrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration forbidden response has a 5xx status code
func (o *GetIntegrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration forbidden response a status code equal to that given
func (o *GetIntegrationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationNotFound creates a GetIntegrationNotFound with default headers values
func NewGetIntegrationNotFound() *GetIntegrationNotFound {
	return &GetIntegrationNotFound{}
}

/*
GetIntegrationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration not found response has a 2xx status code
func (o *GetIntegrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration not found response has a 3xx status code
func (o *GetIntegrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration not found response has a 4xx status code
func (o *GetIntegrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration not found response has a 5xx status code
func (o *GetIntegrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration not found response a status code equal to that given
func (o *GetIntegrationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationRequestTimeout creates a GetIntegrationRequestTimeout with default headers values
func NewGetIntegrationRequestTimeout() *GetIntegrationRequestTimeout {
	return &GetIntegrationRequestTimeout{}
}

/*
GetIntegrationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration request timeout response has a 2xx status code
func (o *GetIntegrationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration request timeout response has a 3xx status code
func (o *GetIntegrationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration request timeout response has a 4xx status code
func (o *GetIntegrationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration request timeout response has a 5xx status code
func (o *GetIntegrationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration request timeout response a status code equal to that given
func (o *GetIntegrationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationRequestEntityTooLarge creates a GetIntegrationRequestEntityTooLarge with default headers values
func NewGetIntegrationRequestEntityTooLarge() *GetIntegrationRequestEntityTooLarge {
	return &GetIntegrationRequestEntityTooLarge{}
}

/*
GetIntegrationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration request entity too large response has a 2xx status code
func (o *GetIntegrationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration request entity too large response has a 3xx status code
func (o *GetIntegrationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration request entity too large response has a 4xx status code
func (o *GetIntegrationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration request entity too large response has a 5xx status code
func (o *GetIntegrationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration request entity too large response a status code equal to that given
func (o *GetIntegrationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationUnsupportedMediaType creates a GetIntegrationUnsupportedMediaType with default headers values
func NewGetIntegrationUnsupportedMediaType() *GetIntegrationUnsupportedMediaType {
	return &GetIntegrationUnsupportedMediaType{}
}

/*
GetIntegrationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration unsupported media type response has a 2xx status code
func (o *GetIntegrationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration unsupported media type response has a 3xx status code
func (o *GetIntegrationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration unsupported media type response has a 4xx status code
func (o *GetIntegrationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration unsupported media type response has a 5xx status code
func (o *GetIntegrationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration unsupported media type response a status code equal to that given
func (o *GetIntegrationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationTooManyRequests creates a GetIntegrationTooManyRequests with default headers values
func NewGetIntegrationTooManyRequests() *GetIntegrationTooManyRequests {
	return &GetIntegrationTooManyRequests{}
}

/*
GetIntegrationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration too many requests response has a 2xx status code
func (o *GetIntegrationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration too many requests response has a 3xx status code
func (o *GetIntegrationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration too many requests response has a 4xx status code
func (o *GetIntegrationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integration too many requests response has a 5xx status code
func (o *GetIntegrationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integration too many requests response a status code equal to that given
func (o *GetIntegrationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationInternalServerError creates a GetIntegrationInternalServerError with default headers values
func NewGetIntegrationInternalServerError() *GetIntegrationInternalServerError {
	return &GetIntegrationInternalServerError{}
}

/*
GetIntegrationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration internal server error response has a 2xx status code
func (o *GetIntegrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration internal server error response has a 3xx status code
func (o *GetIntegrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration internal server error response has a 4xx status code
func (o *GetIntegrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration internal server error response has a 5xx status code
func (o *GetIntegrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integration internal server error response a status code equal to that given
func (o *GetIntegrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationServiceUnavailable creates a GetIntegrationServiceUnavailable with default headers values
func NewGetIntegrationServiceUnavailable() *GetIntegrationServiceUnavailable {
	return &GetIntegrationServiceUnavailable{}
}

/*
GetIntegrationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration service unavailable response has a 2xx status code
func (o *GetIntegrationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration service unavailable response has a 3xx status code
func (o *GetIntegrationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration service unavailable response has a 4xx status code
func (o *GetIntegrationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration service unavailable response has a 5xx status code
func (o *GetIntegrationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integration service unavailable response a status code equal to that given
func (o *GetIntegrationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationGatewayTimeout creates a GetIntegrationGatewayTimeout with default headers values
func NewGetIntegrationGatewayTimeout() *GetIntegrationGatewayTimeout {
	return &GetIntegrationGatewayTimeout{}
}

/*
GetIntegrationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integration gateway timeout response has a 2xx status code
func (o *GetIntegrationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integration gateway timeout response has a 3xx status code
func (o *GetIntegrationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integration gateway timeout response has a 4xx status code
func (o *GetIntegrationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integration gateway timeout response has a 5xx status code
func (o *GetIntegrationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integration gateway timeout response a status code equal to that given
func (o *GetIntegrationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/{integrationId}][%d] getIntegrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
