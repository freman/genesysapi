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

// PutIntegrationsCredentialReader is a Reader for the PutIntegrationsCredential structure.
type PutIntegrationsCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIntegrationsCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIntegrationsCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIntegrationsCredentialBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIntegrationsCredentialUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIntegrationsCredentialForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIntegrationsCredentialNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIntegrationsCredentialRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIntegrationsCredentialUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIntegrationsCredentialTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIntegrationsCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIntegrationsCredentialServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIntegrationsCredentialGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIntegrationsCredentialOK creates a PutIntegrationsCredentialOK with default headers values
func NewPutIntegrationsCredentialOK() *PutIntegrationsCredentialOK {
	return &PutIntegrationsCredentialOK{}
}

/*PutIntegrationsCredentialOK handles this case with default header values.

successful operation
*/
type PutIntegrationsCredentialOK struct {
	Payload *models.CredentialInfo
}

func (o *PutIntegrationsCredentialOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialOK  %+v", 200, o.Payload)
}

func (o *PutIntegrationsCredentialOK) GetPayload() *models.CredentialInfo {
	return o.Payload
}

func (o *PutIntegrationsCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialBadRequest creates a PutIntegrationsCredentialBadRequest with default headers values
func NewPutIntegrationsCredentialBadRequest() *PutIntegrationsCredentialBadRequest {
	return &PutIntegrationsCredentialBadRequest{}
}

/*PutIntegrationsCredentialBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIntegrationsCredentialBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialBadRequest  %+v", 400, o.Payload)
}

func (o *PutIntegrationsCredentialBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialUnauthorized creates a PutIntegrationsCredentialUnauthorized with default headers values
func NewPutIntegrationsCredentialUnauthorized() *PutIntegrationsCredentialUnauthorized {
	return &PutIntegrationsCredentialUnauthorized{}
}

/*PutIntegrationsCredentialUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIntegrationsCredentialUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIntegrationsCredentialUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialForbidden creates a PutIntegrationsCredentialForbidden with default headers values
func NewPutIntegrationsCredentialForbidden() *PutIntegrationsCredentialForbidden {
	return &PutIntegrationsCredentialForbidden{}
}

/*PutIntegrationsCredentialForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutIntegrationsCredentialForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialForbidden  %+v", 403, o.Payload)
}

func (o *PutIntegrationsCredentialForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialNotFound creates a PutIntegrationsCredentialNotFound with default headers values
func NewPutIntegrationsCredentialNotFound() *PutIntegrationsCredentialNotFound {
	return &PutIntegrationsCredentialNotFound{}
}

/*PutIntegrationsCredentialNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIntegrationsCredentialNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialNotFound  %+v", 404, o.Payload)
}

func (o *PutIntegrationsCredentialNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialRequestEntityTooLarge creates a PutIntegrationsCredentialRequestEntityTooLarge with default headers values
func NewPutIntegrationsCredentialRequestEntityTooLarge() *PutIntegrationsCredentialRequestEntityTooLarge {
	return &PutIntegrationsCredentialRequestEntityTooLarge{}
}

/*PutIntegrationsCredentialRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutIntegrationsCredentialRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialUnsupportedMediaType creates a PutIntegrationsCredentialUnsupportedMediaType with default headers values
func NewPutIntegrationsCredentialUnsupportedMediaType() *PutIntegrationsCredentialUnsupportedMediaType {
	return &PutIntegrationsCredentialUnsupportedMediaType{}
}

/*PutIntegrationsCredentialUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIntegrationsCredentialUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialTooManyRequests creates a PutIntegrationsCredentialTooManyRequests with default headers values
func NewPutIntegrationsCredentialTooManyRequests() *PutIntegrationsCredentialTooManyRequests {
	return &PutIntegrationsCredentialTooManyRequests{}
}

/*PutIntegrationsCredentialTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutIntegrationsCredentialTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIntegrationsCredentialTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialInternalServerError creates a PutIntegrationsCredentialInternalServerError with default headers values
func NewPutIntegrationsCredentialInternalServerError() *PutIntegrationsCredentialInternalServerError {
	return &PutIntegrationsCredentialInternalServerError{}
}

/*PutIntegrationsCredentialInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIntegrationsCredentialInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIntegrationsCredentialInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialServiceUnavailable creates a PutIntegrationsCredentialServiceUnavailable with default headers values
func NewPutIntegrationsCredentialServiceUnavailable() *PutIntegrationsCredentialServiceUnavailable {
	return &PutIntegrationsCredentialServiceUnavailable{}
}

/*PutIntegrationsCredentialServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIntegrationsCredentialServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIntegrationsCredentialServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIntegrationsCredentialGatewayTimeout creates a PutIntegrationsCredentialGatewayTimeout with default headers values
func NewPutIntegrationsCredentialGatewayTimeout() *PutIntegrationsCredentialGatewayTimeout {
	return &PutIntegrationsCredentialGatewayTimeout{}
}

/*PutIntegrationsCredentialGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutIntegrationsCredentialGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIntegrationsCredentialGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/integrations/credentials/{credentialId}][%d] putIntegrationsCredentialGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIntegrationsCredentialGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIntegrationsCredentialGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}