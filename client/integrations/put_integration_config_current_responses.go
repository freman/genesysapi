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

// PutIntegrationConfigCurrentReader is a Reader for the PutIntegrationConfigCurrent structure.
type PutIntegrationConfigCurrentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIntegrationConfigCurrentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIntegrationConfigCurrentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIntegrationConfigCurrentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIntegrationConfigCurrentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIntegrationConfigCurrentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIntegrationConfigCurrentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIntegrationConfigCurrentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutIntegrationConfigCurrentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIntegrationConfigCurrentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIntegrationConfigCurrentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIntegrationConfigCurrentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIntegrationConfigCurrentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIntegrationConfigCurrentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIntegrationConfigCurrentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIntegrationConfigCurrentOK creates a PutIntegrationConfigCurrentOK with default headers values
func NewPutIntegrationConfigCurrentOK() *PutIntegrationConfigCurrentOK {
	return &PutIntegrationConfigCurrentOK{}
}

/*
PutIntegrationConfigCurrentOK describes a response with status code 200, with default header values.

successful operation
*/
type PutIntegrationConfigCurrentOK struct {
	Payload *models.IntegrationConfiguration
}

// IsSuccess returns true when this put integration config current o k response has a 2xx status code
func (o *PutIntegrationConfigCurrentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put integration config current o k response has a 3xx status code
func (o *PutIntegrationConfigCurrentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current o k response has a 4xx status code
func (o *PutIntegrationConfigCurrentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integration config current o k response has a 5xx status code
func (o *PutIntegrationConfigCurrentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current o k response a status code equal to that given
func (o *PutIntegrationConfigCurrentOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutIntegrationConfigCurrentOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentOK  %+v", 200, o.Payload)
}

func (o *PutIntegrationConfigCurrentOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentOK  %+v", 200, o.Payload)
}

func (o *PutIntegrationConfigCurrentOK) GetPayload() *models.IntegrationConfiguration {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentBadRequest creates a PutIntegrationConfigCurrentBadRequest with default headers values
func NewPutIntegrationConfigCurrentBadRequest() *PutIntegrationConfigCurrentBadRequest {
	return &PutIntegrationConfigCurrentBadRequest{}
}

/*
PutIntegrationConfigCurrentBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIntegrationConfigCurrentBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current bad request response has a 2xx status code
func (o *PutIntegrationConfigCurrentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current bad request response has a 3xx status code
func (o *PutIntegrationConfigCurrentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current bad request response has a 4xx status code
func (o *PutIntegrationConfigCurrentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current bad request response has a 5xx status code
func (o *PutIntegrationConfigCurrentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current bad request response a status code equal to that given
func (o *PutIntegrationConfigCurrentBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutIntegrationConfigCurrentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentBadRequest  %+v", 400, o.Payload)
}

func (o *PutIntegrationConfigCurrentBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentBadRequest  %+v", 400, o.Payload)
}

func (o *PutIntegrationConfigCurrentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentUnauthorized creates a PutIntegrationConfigCurrentUnauthorized with default headers values
func NewPutIntegrationConfigCurrentUnauthorized() *PutIntegrationConfigCurrentUnauthorized {
	return &PutIntegrationConfigCurrentUnauthorized{}
}

/*
PutIntegrationConfigCurrentUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIntegrationConfigCurrentUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current unauthorized response has a 2xx status code
func (o *PutIntegrationConfigCurrentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current unauthorized response has a 3xx status code
func (o *PutIntegrationConfigCurrentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current unauthorized response has a 4xx status code
func (o *PutIntegrationConfigCurrentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current unauthorized response has a 5xx status code
func (o *PutIntegrationConfigCurrentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current unauthorized response a status code equal to that given
func (o *PutIntegrationConfigCurrentUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutIntegrationConfigCurrentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIntegrationConfigCurrentUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIntegrationConfigCurrentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentForbidden creates a PutIntegrationConfigCurrentForbidden with default headers values
func NewPutIntegrationConfigCurrentForbidden() *PutIntegrationConfigCurrentForbidden {
	return &PutIntegrationConfigCurrentForbidden{}
}

/*
PutIntegrationConfigCurrentForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutIntegrationConfigCurrentForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current forbidden response has a 2xx status code
func (o *PutIntegrationConfigCurrentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current forbidden response has a 3xx status code
func (o *PutIntegrationConfigCurrentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current forbidden response has a 4xx status code
func (o *PutIntegrationConfigCurrentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current forbidden response has a 5xx status code
func (o *PutIntegrationConfigCurrentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current forbidden response a status code equal to that given
func (o *PutIntegrationConfigCurrentForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutIntegrationConfigCurrentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentForbidden  %+v", 403, o.Payload)
}

func (o *PutIntegrationConfigCurrentForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentForbidden  %+v", 403, o.Payload)
}

func (o *PutIntegrationConfigCurrentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentNotFound creates a PutIntegrationConfigCurrentNotFound with default headers values
func NewPutIntegrationConfigCurrentNotFound() *PutIntegrationConfigCurrentNotFound {
	return &PutIntegrationConfigCurrentNotFound{}
}

/*
PutIntegrationConfigCurrentNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutIntegrationConfigCurrentNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current not found response has a 2xx status code
func (o *PutIntegrationConfigCurrentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current not found response has a 3xx status code
func (o *PutIntegrationConfigCurrentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current not found response has a 4xx status code
func (o *PutIntegrationConfigCurrentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current not found response has a 5xx status code
func (o *PutIntegrationConfigCurrentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current not found response a status code equal to that given
func (o *PutIntegrationConfigCurrentNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutIntegrationConfigCurrentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentNotFound  %+v", 404, o.Payload)
}

func (o *PutIntegrationConfigCurrentNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentNotFound  %+v", 404, o.Payload)
}

func (o *PutIntegrationConfigCurrentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentRequestTimeout creates a PutIntegrationConfigCurrentRequestTimeout with default headers values
func NewPutIntegrationConfigCurrentRequestTimeout() *PutIntegrationConfigCurrentRequestTimeout {
	return &PutIntegrationConfigCurrentRequestTimeout{}
}

/*
PutIntegrationConfigCurrentRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIntegrationConfigCurrentRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current request timeout response has a 2xx status code
func (o *PutIntegrationConfigCurrentRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current request timeout response has a 3xx status code
func (o *PutIntegrationConfigCurrentRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current request timeout response has a 4xx status code
func (o *PutIntegrationConfigCurrentRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current request timeout response has a 5xx status code
func (o *PutIntegrationConfigCurrentRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current request timeout response a status code equal to that given
func (o *PutIntegrationConfigCurrentRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutIntegrationConfigCurrentRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIntegrationConfigCurrentRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIntegrationConfigCurrentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentConflict creates a PutIntegrationConfigCurrentConflict with default headers values
func NewPutIntegrationConfigCurrentConflict() *PutIntegrationConfigCurrentConflict {
	return &PutIntegrationConfigCurrentConflict{}
}

/*
PutIntegrationConfigCurrentConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutIntegrationConfigCurrentConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current conflict response has a 2xx status code
func (o *PutIntegrationConfigCurrentConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current conflict response has a 3xx status code
func (o *PutIntegrationConfigCurrentConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current conflict response has a 4xx status code
func (o *PutIntegrationConfigCurrentConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current conflict response has a 5xx status code
func (o *PutIntegrationConfigCurrentConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current conflict response a status code equal to that given
func (o *PutIntegrationConfigCurrentConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutIntegrationConfigCurrentConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentConflict  %+v", 409, o.Payload)
}

func (o *PutIntegrationConfigCurrentConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentConflict  %+v", 409, o.Payload)
}

func (o *PutIntegrationConfigCurrentConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentRequestEntityTooLarge creates a PutIntegrationConfigCurrentRequestEntityTooLarge with default headers values
func NewPutIntegrationConfigCurrentRequestEntityTooLarge() *PutIntegrationConfigCurrentRequestEntityTooLarge {
	return &PutIntegrationConfigCurrentRequestEntityTooLarge{}
}

/*
PutIntegrationConfigCurrentRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutIntegrationConfigCurrentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current request entity too large response has a 2xx status code
func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current request entity too large response has a 3xx status code
func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current request entity too large response has a 4xx status code
func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current request entity too large response has a 5xx status code
func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current request entity too large response a status code equal to that given
func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentUnsupportedMediaType creates a PutIntegrationConfigCurrentUnsupportedMediaType with default headers values
func NewPutIntegrationConfigCurrentUnsupportedMediaType() *PutIntegrationConfigCurrentUnsupportedMediaType {
	return &PutIntegrationConfigCurrentUnsupportedMediaType{}
}

/*
PutIntegrationConfigCurrentUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIntegrationConfigCurrentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current unsupported media type response has a 2xx status code
func (o *PutIntegrationConfigCurrentUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current unsupported media type response has a 3xx status code
func (o *PutIntegrationConfigCurrentUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current unsupported media type response has a 4xx status code
func (o *PutIntegrationConfigCurrentUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current unsupported media type response has a 5xx status code
func (o *PutIntegrationConfigCurrentUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current unsupported media type response a status code equal to that given
func (o *PutIntegrationConfigCurrentUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutIntegrationConfigCurrentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIntegrationConfigCurrentUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIntegrationConfigCurrentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentTooManyRequests creates a PutIntegrationConfigCurrentTooManyRequests with default headers values
func NewPutIntegrationConfigCurrentTooManyRequests() *PutIntegrationConfigCurrentTooManyRequests {
	return &PutIntegrationConfigCurrentTooManyRequests{}
}

/*
PutIntegrationConfigCurrentTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIntegrationConfigCurrentTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current too many requests response has a 2xx status code
func (o *PutIntegrationConfigCurrentTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current too many requests response has a 3xx status code
func (o *PutIntegrationConfigCurrentTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current too many requests response has a 4xx status code
func (o *PutIntegrationConfigCurrentTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put integration config current too many requests response has a 5xx status code
func (o *PutIntegrationConfigCurrentTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put integration config current too many requests response a status code equal to that given
func (o *PutIntegrationConfigCurrentTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutIntegrationConfigCurrentTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIntegrationConfigCurrentTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIntegrationConfigCurrentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentInternalServerError creates a PutIntegrationConfigCurrentInternalServerError with default headers values
func NewPutIntegrationConfigCurrentInternalServerError() *PutIntegrationConfigCurrentInternalServerError {
	return &PutIntegrationConfigCurrentInternalServerError{}
}

/*
PutIntegrationConfigCurrentInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIntegrationConfigCurrentInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current internal server error response has a 2xx status code
func (o *PutIntegrationConfigCurrentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current internal server error response has a 3xx status code
func (o *PutIntegrationConfigCurrentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current internal server error response has a 4xx status code
func (o *PutIntegrationConfigCurrentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integration config current internal server error response has a 5xx status code
func (o *PutIntegrationConfigCurrentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put integration config current internal server error response a status code equal to that given
func (o *PutIntegrationConfigCurrentInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutIntegrationConfigCurrentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIntegrationConfigCurrentInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIntegrationConfigCurrentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentServiceUnavailable creates a PutIntegrationConfigCurrentServiceUnavailable with default headers values
func NewPutIntegrationConfigCurrentServiceUnavailable() *PutIntegrationConfigCurrentServiceUnavailable {
	return &PutIntegrationConfigCurrentServiceUnavailable{}
}

/*
PutIntegrationConfigCurrentServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIntegrationConfigCurrentServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current service unavailable response has a 2xx status code
func (o *PutIntegrationConfigCurrentServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current service unavailable response has a 3xx status code
func (o *PutIntegrationConfigCurrentServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current service unavailable response has a 4xx status code
func (o *PutIntegrationConfigCurrentServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integration config current service unavailable response has a 5xx status code
func (o *PutIntegrationConfigCurrentServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put integration config current service unavailable response a status code equal to that given
func (o *PutIntegrationConfigCurrentServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutIntegrationConfigCurrentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIntegrationConfigCurrentServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIntegrationConfigCurrentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationConfigCurrentGatewayTimeout creates a PutIntegrationConfigCurrentGatewayTimeout with default headers values
func NewPutIntegrationConfigCurrentGatewayTimeout() *PutIntegrationConfigCurrentGatewayTimeout {
	return &PutIntegrationConfigCurrentGatewayTimeout{}
}

/*
PutIntegrationConfigCurrentGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutIntegrationConfigCurrentGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put integration config current gateway timeout response has a 2xx status code
func (o *PutIntegrationConfigCurrentGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put integration config current gateway timeout response has a 3xx status code
func (o *PutIntegrationConfigCurrentGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put integration config current gateway timeout response has a 4xx status code
func (o *PutIntegrationConfigCurrentGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put integration config current gateway timeout response has a 5xx status code
func (o *PutIntegrationConfigCurrentGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put integration config current gateway timeout response a status code equal to that given
func (o *PutIntegrationConfigCurrentGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutIntegrationConfigCurrentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIntegrationConfigCurrentGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/{integrationId}/config/current][%d] putIntegrationConfigCurrentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIntegrationConfigCurrentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationConfigCurrentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
