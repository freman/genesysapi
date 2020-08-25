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

/*PutIntegrationConfigCurrentOK handles this case with default header values.

successful operation
*/
type PutIntegrationConfigCurrentOK struct {
	Payload *models.IntegrationConfiguration
}

func (o *PutIntegrationConfigCurrentOK) Error() string {
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

/*PutIntegrationConfigCurrentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIntegrationConfigCurrentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentBadRequest) Error() string {
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

/*PutIntegrationConfigCurrentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIntegrationConfigCurrentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentUnauthorized) Error() string {
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

/*PutIntegrationConfigCurrentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutIntegrationConfigCurrentForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentForbidden) Error() string {
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

/*PutIntegrationConfigCurrentNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIntegrationConfigCurrentNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentNotFound) Error() string {
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

// NewPutIntegrationConfigCurrentConflict creates a PutIntegrationConfigCurrentConflict with default headers values
func NewPutIntegrationConfigCurrentConflict() *PutIntegrationConfigCurrentConflict {
	return &PutIntegrationConfigCurrentConflict{}
}

/*PutIntegrationConfigCurrentConflict handles this case with default header values.

Conflict
*/
type PutIntegrationConfigCurrentConflict struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentConflict) Error() string {
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

/*PutIntegrationConfigCurrentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutIntegrationConfigCurrentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentRequestEntityTooLarge) Error() string {
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

/*PutIntegrationConfigCurrentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIntegrationConfigCurrentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentUnsupportedMediaType) Error() string {
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

/*PutIntegrationConfigCurrentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutIntegrationConfigCurrentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentTooManyRequests) Error() string {
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

/*PutIntegrationConfigCurrentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIntegrationConfigCurrentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentInternalServerError) Error() string {
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

/*PutIntegrationConfigCurrentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIntegrationConfigCurrentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentServiceUnavailable) Error() string {
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

/*PutIntegrationConfigCurrentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutIntegrationConfigCurrentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationConfigCurrentGatewayTimeout) Error() string {
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