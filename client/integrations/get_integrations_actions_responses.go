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

// GetIntegrationsActionsReader is a Reader for the GetIntegrationsActions structure.
type GetIntegrationsActionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionsOK creates a GetIntegrationsActionsOK with default headers values
func NewGetIntegrationsActionsOK() *GetIntegrationsActionsOK {
	return &GetIntegrationsActionsOK{}
}

/*GetIntegrationsActionsOK handles this case with default header values.

successful operation
*/
type GetIntegrationsActionsOK struct {
	Payload *models.ActionEntityListing
}

func (o *GetIntegrationsActionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionsOK) GetPayload() *models.ActionEntityListing {
	return o.Payload
}

func (o *GetIntegrationsActionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsBadRequest creates a GetIntegrationsActionsBadRequest with default headers values
func NewGetIntegrationsActionsBadRequest() *GetIntegrationsActionsBadRequest {
	return &GetIntegrationsActionsBadRequest{}
}

/*GetIntegrationsActionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsUnauthorized creates a GetIntegrationsActionsUnauthorized with default headers values
func NewGetIntegrationsActionsUnauthorized() *GetIntegrationsActionsUnauthorized {
	return &GetIntegrationsActionsUnauthorized{}
}

/*GetIntegrationsActionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsForbidden creates a GetIntegrationsActionsForbidden with default headers values
func NewGetIntegrationsActionsForbidden() *GetIntegrationsActionsForbidden {
	return &GetIntegrationsActionsForbidden{}
}

/*GetIntegrationsActionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsNotFound creates a GetIntegrationsActionsNotFound with default headers values
func NewGetIntegrationsActionsNotFound() *GetIntegrationsActionsNotFound {
	return &GetIntegrationsActionsNotFound{}
}

/*GetIntegrationsActionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsRequestEntityTooLarge creates a GetIntegrationsActionsRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionsRequestEntityTooLarge() *GetIntegrationsActionsRequestEntityTooLarge {
	return &GetIntegrationsActionsRequestEntityTooLarge{}
}

/*GetIntegrationsActionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsActionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsUnsupportedMediaType creates a GetIntegrationsActionsUnsupportedMediaType with default headers values
func NewGetIntegrationsActionsUnsupportedMediaType() *GetIntegrationsActionsUnsupportedMediaType {
	return &GetIntegrationsActionsUnsupportedMediaType{}
}

/*GetIntegrationsActionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsTooManyRequests creates a GetIntegrationsActionsTooManyRequests with default headers values
func NewGetIntegrationsActionsTooManyRequests() *GetIntegrationsActionsTooManyRequests {
	return &GetIntegrationsActionsTooManyRequests{}
}

/*GetIntegrationsActionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsActionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsInternalServerError creates a GetIntegrationsActionsInternalServerError with default headers values
func NewGetIntegrationsActionsInternalServerError() *GetIntegrationsActionsInternalServerError {
	return &GetIntegrationsActionsInternalServerError{}
}

/*GetIntegrationsActionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsServiceUnavailable creates a GetIntegrationsActionsServiceUnavailable with default headers values
func NewGetIntegrationsActionsServiceUnavailable() *GetIntegrationsActionsServiceUnavailable {
	return &GetIntegrationsActionsServiceUnavailable{}
}

/*GetIntegrationsActionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsGatewayTimeout creates a GetIntegrationsActionsGatewayTimeout with default headers values
func NewGetIntegrationsActionsGatewayTimeout() *GetIntegrationsActionsGatewayTimeout {
	return &GetIntegrationsActionsGatewayTimeout{}
}

/*GetIntegrationsActionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsActionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions][%d] getIntegrationsActionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
