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

// DeleteIntegrationReader is a Reader for the DeleteIntegration structure.
type DeleteIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIntegrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIntegrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteIntegrationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteIntegrationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteIntegrationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteIntegrationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIntegrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteIntegrationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteIntegrationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIntegrationOK creates a DeleteIntegrationOK with default headers values
func NewDeleteIntegrationOK() *DeleteIntegrationOK {
	return &DeleteIntegrationOK{}
}

/*
DeleteIntegrationOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteIntegrationOK struct {
	Payload *models.Integration
}

// IsSuccess returns true when this delete integration o k response has a 2xx status code
func (o *DeleteIntegrationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete integration o k response has a 3xx status code
func (o *DeleteIntegrationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration o k response has a 4xx status code
func (o *DeleteIntegrationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integration o k response has a 5xx status code
func (o *DeleteIntegrationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration o k response a status code equal to that given
func (o *DeleteIntegrationOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteIntegrationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteIntegrationOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationOK  %+v", 200, o.Payload)
}

func (o *DeleteIntegrationOK) GetPayload() *models.Integration {
	return o.Payload
}

func (o *DeleteIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Integration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationBadRequest creates a DeleteIntegrationBadRequest with default headers values
func NewDeleteIntegrationBadRequest() *DeleteIntegrationBadRequest {
	return &DeleteIntegrationBadRequest{}
}

/*
DeleteIntegrationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteIntegrationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration bad request response has a 2xx status code
func (o *DeleteIntegrationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration bad request response has a 3xx status code
func (o *DeleteIntegrationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration bad request response has a 4xx status code
func (o *DeleteIntegrationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration bad request response has a 5xx status code
func (o *DeleteIntegrationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration bad request response a status code equal to that given
func (o *DeleteIntegrationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIntegrationBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIntegrationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationUnauthorized creates a DeleteIntegrationUnauthorized with default headers values
func NewDeleteIntegrationUnauthorized() *DeleteIntegrationUnauthorized {
	return &DeleteIntegrationUnauthorized{}
}

/*
DeleteIntegrationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteIntegrationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration unauthorized response has a 2xx status code
func (o *DeleteIntegrationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration unauthorized response has a 3xx status code
func (o *DeleteIntegrationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration unauthorized response has a 4xx status code
func (o *DeleteIntegrationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration unauthorized response has a 5xx status code
func (o *DeleteIntegrationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration unauthorized response a status code equal to that given
func (o *DeleteIntegrationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteIntegrationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIntegrationUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIntegrationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationForbidden creates a DeleteIntegrationForbidden with default headers values
func NewDeleteIntegrationForbidden() *DeleteIntegrationForbidden {
	return &DeleteIntegrationForbidden{}
}

/*
DeleteIntegrationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteIntegrationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration forbidden response has a 2xx status code
func (o *DeleteIntegrationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration forbidden response has a 3xx status code
func (o *DeleteIntegrationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration forbidden response has a 4xx status code
func (o *DeleteIntegrationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration forbidden response has a 5xx status code
func (o *DeleteIntegrationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration forbidden response a status code equal to that given
func (o *DeleteIntegrationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteIntegrationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIntegrationForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIntegrationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationNotFound creates a DeleteIntegrationNotFound with default headers values
func NewDeleteIntegrationNotFound() *DeleteIntegrationNotFound {
	return &DeleteIntegrationNotFound{}
}

/*
DeleteIntegrationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteIntegrationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration not found response has a 2xx status code
func (o *DeleteIntegrationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration not found response has a 3xx status code
func (o *DeleteIntegrationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration not found response has a 4xx status code
func (o *DeleteIntegrationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration not found response has a 5xx status code
func (o *DeleteIntegrationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration not found response a status code equal to that given
func (o *DeleteIntegrationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteIntegrationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIntegrationNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIntegrationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationRequestTimeout creates a DeleteIntegrationRequestTimeout with default headers values
func NewDeleteIntegrationRequestTimeout() *DeleteIntegrationRequestTimeout {
	return &DeleteIntegrationRequestTimeout{}
}

/*
DeleteIntegrationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteIntegrationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration request timeout response has a 2xx status code
func (o *DeleteIntegrationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration request timeout response has a 3xx status code
func (o *DeleteIntegrationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration request timeout response has a 4xx status code
func (o *DeleteIntegrationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration request timeout response has a 5xx status code
func (o *DeleteIntegrationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration request timeout response a status code equal to that given
func (o *DeleteIntegrationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteIntegrationRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIntegrationRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteIntegrationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationRequestEntityTooLarge creates a DeleteIntegrationRequestEntityTooLarge with default headers values
func NewDeleteIntegrationRequestEntityTooLarge() *DeleteIntegrationRequestEntityTooLarge {
	return &DeleteIntegrationRequestEntityTooLarge{}
}

/*
DeleteIntegrationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteIntegrationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration request entity too large response has a 2xx status code
func (o *DeleteIntegrationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration request entity too large response has a 3xx status code
func (o *DeleteIntegrationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration request entity too large response has a 4xx status code
func (o *DeleteIntegrationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration request entity too large response has a 5xx status code
func (o *DeleteIntegrationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration request entity too large response a status code equal to that given
func (o *DeleteIntegrationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteIntegrationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIntegrationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteIntegrationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationUnsupportedMediaType creates a DeleteIntegrationUnsupportedMediaType with default headers values
func NewDeleteIntegrationUnsupportedMediaType() *DeleteIntegrationUnsupportedMediaType {
	return &DeleteIntegrationUnsupportedMediaType{}
}

/*
DeleteIntegrationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteIntegrationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration unsupported media type response has a 2xx status code
func (o *DeleteIntegrationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration unsupported media type response has a 3xx status code
func (o *DeleteIntegrationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration unsupported media type response has a 4xx status code
func (o *DeleteIntegrationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration unsupported media type response has a 5xx status code
func (o *DeleteIntegrationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration unsupported media type response a status code equal to that given
func (o *DeleteIntegrationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteIntegrationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIntegrationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteIntegrationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationTooManyRequests creates a DeleteIntegrationTooManyRequests with default headers values
func NewDeleteIntegrationTooManyRequests() *DeleteIntegrationTooManyRequests {
	return &DeleteIntegrationTooManyRequests{}
}

/*
DeleteIntegrationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteIntegrationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration too many requests response has a 2xx status code
func (o *DeleteIntegrationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration too many requests response has a 3xx status code
func (o *DeleteIntegrationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration too many requests response has a 4xx status code
func (o *DeleteIntegrationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete integration too many requests response has a 5xx status code
func (o *DeleteIntegrationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete integration too many requests response a status code equal to that given
func (o *DeleteIntegrationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteIntegrationTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIntegrationTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteIntegrationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationInternalServerError creates a DeleteIntegrationInternalServerError with default headers values
func NewDeleteIntegrationInternalServerError() *DeleteIntegrationInternalServerError {
	return &DeleteIntegrationInternalServerError{}
}

/*
DeleteIntegrationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteIntegrationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration internal server error response has a 2xx status code
func (o *DeleteIntegrationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration internal server error response has a 3xx status code
func (o *DeleteIntegrationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration internal server error response has a 4xx status code
func (o *DeleteIntegrationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integration internal server error response has a 5xx status code
func (o *DeleteIntegrationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete integration internal server error response a status code equal to that given
func (o *DeleteIntegrationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteIntegrationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIntegrationInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIntegrationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationServiceUnavailable creates a DeleteIntegrationServiceUnavailable with default headers values
func NewDeleteIntegrationServiceUnavailable() *DeleteIntegrationServiceUnavailable {
	return &DeleteIntegrationServiceUnavailable{}
}

/*
DeleteIntegrationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteIntegrationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration service unavailable response has a 2xx status code
func (o *DeleteIntegrationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration service unavailable response has a 3xx status code
func (o *DeleteIntegrationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration service unavailable response has a 4xx status code
func (o *DeleteIntegrationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integration service unavailable response has a 5xx status code
func (o *DeleteIntegrationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete integration service unavailable response a status code equal to that given
func (o *DeleteIntegrationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteIntegrationServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIntegrationServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteIntegrationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIntegrationGatewayTimeout creates a DeleteIntegrationGatewayTimeout with default headers values
func NewDeleteIntegrationGatewayTimeout() *DeleteIntegrationGatewayTimeout {
	return &DeleteIntegrationGatewayTimeout{}
}

/*
DeleteIntegrationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteIntegrationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete integration gateway timeout response has a 2xx status code
func (o *DeleteIntegrationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete integration gateway timeout response has a 3xx status code
func (o *DeleteIntegrationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete integration gateway timeout response has a 4xx status code
func (o *DeleteIntegrationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete integration gateway timeout response has a 5xx status code
func (o *DeleteIntegrationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete integration gateway timeout response a status code equal to that given
func (o *DeleteIntegrationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteIntegrationGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIntegrationGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/integrations/{integrationId}][%d] deleteIntegrationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteIntegrationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteIntegrationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
