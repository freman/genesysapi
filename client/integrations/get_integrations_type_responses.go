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

// GetIntegrationsTypeReader is a Reader for the GetIntegrationsType structure.
type GetIntegrationsTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsTypeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsTypeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsTypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsTypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsTypeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsTypeOK creates a GetIntegrationsTypeOK with default headers values
func NewGetIntegrationsTypeOK() *GetIntegrationsTypeOK {
	return &GetIntegrationsTypeOK{}
}

/*GetIntegrationsTypeOK handles this case with default header values.

successful operation
*/
type GetIntegrationsTypeOK struct {
	Payload *models.IntegrationType
}

func (o *GetIntegrationsTypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsTypeOK) GetPayload() *models.IntegrationType {
	return o.Payload
}

func (o *GetIntegrationsTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeBadRequest creates a GetIntegrationsTypeBadRequest with default headers values
func NewGetIntegrationsTypeBadRequest() *GetIntegrationsTypeBadRequest {
	return &GetIntegrationsTypeBadRequest{}
}

/*GetIntegrationsTypeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsTypeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsTypeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeUnauthorized creates a GetIntegrationsTypeUnauthorized with default headers values
func NewGetIntegrationsTypeUnauthorized() *GetIntegrationsTypeUnauthorized {
	return &GetIntegrationsTypeUnauthorized{}
}

/*GetIntegrationsTypeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsTypeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsTypeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeForbidden creates a GetIntegrationsTypeForbidden with default headers values
func NewGetIntegrationsTypeForbidden() *GetIntegrationsTypeForbidden {
	return &GetIntegrationsTypeForbidden{}
}

/*GetIntegrationsTypeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsTypeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsTypeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeNotFound creates a GetIntegrationsTypeNotFound with default headers values
func NewGetIntegrationsTypeNotFound() *GetIntegrationsTypeNotFound {
	return &GetIntegrationsTypeNotFound{}
}

/*GetIntegrationsTypeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsTypeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsTypeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeRequestEntityTooLarge creates a GetIntegrationsTypeRequestEntityTooLarge with default headers values
func NewGetIntegrationsTypeRequestEntityTooLarge() *GetIntegrationsTypeRequestEntityTooLarge {
	return &GetIntegrationsTypeRequestEntityTooLarge{}
}

/*GetIntegrationsTypeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsTypeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsTypeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeUnsupportedMediaType creates a GetIntegrationsTypeUnsupportedMediaType with default headers values
func NewGetIntegrationsTypeUnsupportedMediaType() *GetIntegrationsTypeUnsupportedMediaType {
	return &GetIntegrationsTypeUnsupportedMediaType{}
}

/*GetIntegrationsTypeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsTypeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsTypeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeTooManyRequests creates a GetIntegrationsTypeTooManyRequests with default headers values
func NewGetIntegrationsTypeTooManyRequests() *GetIntegrationsTypeTooManyRequests {
	return &GetIntegrationsTypeTooManyRequests{}
}

/*GetIntegrationsTypeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsTypeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsTypeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeInternalServerError creates a GetIntegrationsTypeInternalServerError with default headers values
func NewGetIntegrationsTypeInternalServerError() *GetIntegrationsTypeInternalServerError {
	return &GetIntegrationsTypeInternalServerError{}
}

/*GetIntegrationsTypeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsTypeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsTypeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeServiceUnavailable creates a GetIntegrationsTypeServiceUnavailable with default headers values
func NewGetIntegrationsTypeServiceUnavailable() *GetIntegrationsTypeServiceUnavailable {
	return &GetIntegrationsTypeServiceUnavailable{}
}

/*GetIntegrationsTypeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsTypeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsTypeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypeGatewayTimeout creates a GetIntegrationsTypeGatewayTimeout with default headers values
func NewGetIntegrationsTypeGatewayTimeout() *GetIntegrationsTypeGatewayTimeout {
	return &GetIntegrationsTypeGatewayTimeout{}
}

/*GetIntegrationsTypeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsTypeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsTypeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types/{typeId}][%d] getIntegrationsTypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsTypeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
