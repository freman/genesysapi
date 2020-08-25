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

// GetIntegrationsCredentialReader is a Reader for the GetIntegrationsCredential structure.
type GetIntegrationsCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsCredentialRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsCredentialUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsCredentialTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsCredentialServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsCredentialGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsCredentialOK creates a GetIntegrationsCredentialOK with default headers values
func NewGetIntegrationsCredentialOK() *GetIntegrationsCredentialOK {
	return &GetIntegrationsCredentialOK{}
}

/*GetIntegrationsCredentialOK handles this case with default header values.

successful operation
*/
type GetIntegrationsCredentialOK struct {
	Payload *models.Credential
}

func (o *GetIntegrationsCredentialOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsCredentialOK) GetPayload() *models.Credential {
	return o.Payload
}

func (o *GetIntegrationsCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Credential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialBadRequest creates a GetIntegrationsCredentialBadRequest with default headers values
func NewGetIntegrationsCredentialBadRequest() *GetIntegrationsCredentialBadRequest {
	return &GetIntegrationsCredentialBadRequest{}
}

/*GetIntegrationsCredentialBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsCredentialBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsCredentialBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialUnauthorized creates a GetIntegrationsCredentialUnauthorized with default headers values
func NewGetIntegrationsCredentialUnauthorized() *GetIntegrationsCredentialUnauthorized {
	return &GetIntegrationsCredentialUnauthorized{}
}

/*GetIntegrationsCredentialUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsCredentialUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsCredentialUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialForbidden creates a GetIntegrationsCredentialForbidden with default headers values
func NewGetIntegrationsCredentialForbidden() *GetIntegrationsCredentialForbidden {
	return &GetIntegrationsCredentialForbidden{}
}

/*GetIntegrationsCredentialForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsCredentialForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsCredentialForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialNotFound creates a GetIntegrationsCredentialNotFound with default headers values
func NewGetIntegrationsCredentialNotFound() *GetIntegrationsCredentialNotFound {
	return &GetIntegrationsCredentialNotFound{}
}

/*GetIntegrationsCredentialNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsCredentialNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsCredentialNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialRequestEntityTooLarge creates a GetIntegrationsCredentialRequestEntityTooLarge with default headers values
func NewGetIntegrationsCredentialRequestEntityTooLarge() *GetIntegrationsCredentialRequestEntityTooLarge {
	return &GetIntegrationsCredentialRequestEntityTooLarge{}
}

/*GetIntegrationsCredentialRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsCredentialRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsCredentialRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialUnsupportedMediaType creates a GetIntegrationsCredentialUnsupportedMediaType with default headers values
func NewGetIntegrationsCredentialUnsupportedMediaType() *GetIntegrationsCredentialUnsupportedMediaType {
	return &GetIntegrationsCredentialUnsupportedMediaType{}
}

/*GetIntegrationsCredentialUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsCredentialUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsCredentialUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialTooManyRequests creates a GetIntegrationsCredentialTooManyRequests with default headers values
func NewGetIntegrationsCredentialTooManyRequests() *GetIntegrationsCredentialTooManyRequests {
	return &GetIntegrationsCredentialTooManyRequests{}
}

/*GetIntegrationsCredentialTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsCredentialTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsCredentialTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialInternalServerError creates a GetIntegrationsCredentialInternalServerError with default headers values
func NewGetIntegrationsCredentialInternalServerError() *GetIntegrationsCredentialInternalServerError {
	return &GetIntegrationsCredentialInternalServerError{}
}

/*GetIntegrationsCredentialInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsCredentialInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsCredentialInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialServiceUnavailable creates a GetIntegrationsCredentialServiceUnavailable with default headers values
func NewGetIntegrationsCredentialServiceUnavailable() *GetIntegrationsCredentialServiceUnavailable {
	return &GetIntegrationsCredentialServiceUnavailable{}
}

/*GetIntegrationsCredentialServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsCredentialServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsCredentialServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsCredentialGatewayTimeout creates a GetIntegrationsCredentialGatewayTimeout with default headers values
func NewGetIntegrationsCredentialGatewayTimeout() *GetIntegrationsCredentialGatewayTimeout {
	return &GetIntegrationsCredentialGatewayTimeout{}
}

/*GetIntegrationsCredentialGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsCredentialGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsCredentialGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/credentials/{credentialId}][%d] getIntegrationsCredentialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsCredentialGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsCredentialGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}