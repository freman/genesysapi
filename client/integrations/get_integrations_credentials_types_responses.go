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

// GetIntegrationsCredentialsTypesReader is a Reader for the GetIntegrationsCredentialsTypes structure.
type GetIntegrationsCredentialsTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsCredentialsTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsCredentialsTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsCredentialsTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsCredentialsTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsCredentialsTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsCredentialsTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsCredentialsTypesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsCredentialsTypesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsCredentialsTypesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsCredentialsTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsCredentialsTypesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsCredentialsTypesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsCredentialsTypesOK creates a GetIntegrationsCredentialsTypesOK with default headers values
func NewGetIntegrationsCredentialsTypesOK() *GetIntegrationsCredentialsTypesOK {
	return &GetIntegrationsCredentialsTypesOK{}
}

/*GetIntegrationsCredentialsTypesOK handles this case with default header values.

successful operation
*/
type GetIntegrationsCredentialsTypesOK struct {
	Payload *models.CredentialTypeListing
}

func (o *GetIntegrationsCredentialsTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesOK) GetPayload() *models.CredentialTypeListing {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialTypeListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesBadRequest creates a GetIntegrationsCredentialsTypesBadRequest with default headers values
func NewGetIntegrationsCredentialsTypesBadRequest() *GetIntegrationsCredentialsTypesBadRequest {
	return &GetIntegrationsCredentialsTypesBadRequest{}
}

/*GetIntegrationsCredentialsTypesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsCredentialsTypesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesUnauthorized creates a GetIntegrationsCredentialsTypesUnauthorized with default headers values
func NewGetIntegrationsCredentialsTypesUnauthorized() *GetIntegrationsCredentialsTypesUnauthorized {
	return &GetIntegrationsCredentialsTypesUnauthorized{}
}

/*GetIntegrationsCredentialsTypesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsCredentialsTypesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesForbidden creates a GetIntegrationsCredentialsTypesForbidden with default headers values
func NewGetIntegrationsCredentialsTypesForbidden() *GetIntegrationsCredentialsTypesForbidden {
	return &GetIntegrationsCredentialsTypesForbidden{}
}

/*GetIntegrationsCredentialsTypesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsCredentialsTypesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesNotFound creates a GetIntegrationsCredentialsTypesNotFound with default headers values
func NewGetIntegrationsCredentialsTypesNotFound() *GetIntegrationsCredentialsTypesNotFound {
	return &GetIntegrationsCredentialsTypesNotFound{}
}

/*GetIntegrationsCredentialsTypesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsCredentialsTypesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesRequestEntityTooLarge creates a GetIntegrationsCredentialsTypesRequestEntityTooLarge with default headers values
func NewGetIntegrationsCredentialsTypesRequestEntityTooLarge() *GetIntegrationsCredentialsTypesRequestEntityTooLarge {
	return &GetIntegrationsCredentialsTypesRequestEntityTooLarge{}
}

/*GetIntegrationsCredentialsTypesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsCredentialsTypesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesUnsupportedMediaType creates a GetIntegrationsCredentialsTypesUnsupportedMediaType with default headers values
func NewGetIntegrationsCredentialsTypesUnsupportedMediaType() *GetIntegrationsCredentialsTypesUnsupportedMediaType {
	return &GetIntegrationsCredentialsTypesUnsupportedMediaType{}
}

/*GetIntegrationsCredentialsTypesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsCredentialsTypesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesTooManyRequests creates a GetIntegrationsCredentialsTypesTooManyRequests with default headers values
func NewGetIntegrationsCredentialsTypesTooManyRequests() *GetIntegrationsCredentialsTypesTooManyRequests {
	return &GetIntegrationsCredentialsTypesTooManyRequests{}
}

/*GetIntegrationsCredentialsTypesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsCredentialsTypesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesInternalServerError creates a GetIntegrationsCredentialsTypesInternalServerError with default headers values
func NewGetIntegrationsCredentialsTypesInternalServerError() *GetIntegrationsCredentialsTypesInternalServerError {
	return &GetIntegrationsCredentialsTypesInternalServerError{}
}

/*GetIntegrationsCredentialsTypesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsCredentialsTypesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesServiceUnavailable creates a GetIntegrationsCredentialsTypesServiceUnavailable with default headers values
func NewGetIntegrationsCredentialsTypesServiceUnavailable() *GetIntegrationsCredentialsTypesServiceUnavailable {
	return &GetIntegrationsCredentialsTypesServiceUnavailable{}
}

/*GetIntegrationsCredentialsTypesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsCredentialsTypesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialsTypesGatewayTimeout creates a GetIntegrationsCredentialsTypesGatewayTimeout with default headers values
func NewGetIntegrationsCredentialsTypesGatewayTimeout() *GetIntegrationsCredentialsTypesGatewayTimeout {
	return &GetIntegrationsCredentialsTypesGatewayTimeout{}
}

/*GetIntegrationsCredentialsTypesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsCredentialsTypesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialsTypesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/types][%d] getIntegrationsCredentialsTypesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsCredentialsTypesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialsTypesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
