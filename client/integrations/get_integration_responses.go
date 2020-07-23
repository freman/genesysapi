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

/*GetIntegrationOK handles this case with default header values.

successful operation
*/
type GetIntegrationOK struct {
	Payload *models.Integration
}

func (o *GetIntegrationOK) Error() string {
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

/*GetIntegrationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationBadRequest) Error() string {
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

/*GetIntegrationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationUnauthorized) Error() string {
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

/*GetIntegrationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationForbidden) Error() string {
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

/*GetIntegrationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationNotFound) Error() string {
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

// NewGetIntegrationRequestEntityTooLarge creates a GetIntegrationRequestEntityTooLarge with default headers values
func NewGetIntegrationRequestEntityTooLarge() *GetIntegrationRequestEntityTooLarge {
	return &GetIntegrationRequestEntityTooLarge{}
}

/*GetIntegrationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationRequestEntityTooLarge) Error() string {
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

/*GetIntegrationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationUnsupportedMediaType) Error() string {
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

/*GetIntegrationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationTooManyRequests) Error() string {
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

/*GetIntegrationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationInternalServerError) Error() string {
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

/*GetIntegrationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationServiceUnavailable) Error() string {
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

/*GetIntegrationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationGatewayTimeout) Error() string {
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
