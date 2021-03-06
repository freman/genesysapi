// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutIdentityprovidersGsuiteReader is a Reader for the PutIdentityprovidersGsuite structure.
type PutIdentityprovidersGsuiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersGsuiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersGsuiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersGsuiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersGsuiteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersGsuiteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersGsuiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersGsuiteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersGsuiteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersGsuiteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersGsuiteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersGsuiteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersGsuiteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersGsuiteOK creates a PutIdentityprovidersGsuiteOK with default headers values
func NewPutIdentityprovidersGsuiteOK() *PutIdentityprovidersGsuiteOK {
	return &PutIdentityprovidersGsuiteOK{}
}

/*PutIdentityprovidersGsuiteOK handles this case with default header values.

successful operation
*/
type PutIdentityprovidersGsuiteOK struct {
	Payload *models.OAuthProvider
}

func (o *PutIdentityprovidersGsuiteOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersGsuiteOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteBadRequest creates a PutIdentityprovidersGsuiteBadRequest with default headers values
func NewPutIdentityprovidersGsuiteBadRequest() *PutIdentityprovidersGsuiteBadRequest {
	return &PutIdentityprovidersGsuiteBadRequest{}
}

/*PutIdentityprovidersGsuiteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersGsuiteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersGsuiteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteUnauthorized creates a PutIdentityprovidersGsuiteUnauthorized with default headers values
func NewPutIdentityprovidersGsuiteUnauthorized() *PutIdentityprovidersGsuiteUnauthorized {
	return &PutIdentityprovidersGsuiteUnauthorized{}
}

/*PutIdentityprovidersGsuiteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersGsuiteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersGsuiteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteForbidden creates a PutIdentityprovidersGsuiteForbidden with default headers values
func NewPutIdentityprovidersGsuiteForbidden() *PutIdentityprovidersGsuiteForbidden {
	return &PutIdentityprovidersGsuiteForbidden{}
}

/*PutIdentityprovidersGsuiteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersGsuiteForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersGsuiteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteNotFound creates a PutIdentityprovidersGsuiteNotFound with default headers values
func NewPutIdentityprovidersGsuiteNotFound() *PutIdentityprovidersGsuiteNotFound {
	return &PutIdentityprovidersGsuiteNotFound{}
}

/*PutIdentityprovidersGsuiteNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersGsuiteNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersGsuiteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteRequestEntityTooLarge creates a PutIdentityprovidersGsuiteRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersGsuiteRequestEntityTooLarge() *PutIdentityprovidersGsuiteRequestEntityTooLarge {
	return &PutIdentityprovidersGsuiteRequestEntityTooLarge{}
}

/*PutIdentityprovidersGsuiteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutIdentityprovidersGsuiteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteUnsupportedMediaType creates a PutIdentityprovidersGsuiteUnsupportedMediaType with default headers values
func NewPutIdentityprovidersGsuiteUnsupportedMediaType() *PutIdentityprovidersGsuiteUnsupportedMediaType {
	return &PutIdentityprovidersGsuiteUnsupportedMediaType{}
}

/*PutIdentityprovidersGsuiteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersGsuiteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteTooManyRequests creates a PutIdentityprovidersGsuiteTooManyRequests with default headers values
func NewPutIdentityprovidersGsuiteTooManyRequests() *PutIdentityprovidersGsuiteTooManyRequests {
	return &PutIdentityprovidersGsuiteTooManyRequests{}
}

/*PutIdentityprovidersGsuiteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutIdentityprovidersGsuiteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteInternalServerError creates a PutIdentityprovidersGsuiteInternalServerError with default headers values
func NewPutIdentityprovidersGsuiteInternalServerError() *PutIdentityprovidersGsuiteInternalServerError {
	return &PutIdentityprovidersGsuiteInternalServerError{}
}

/*PutIdentityprovidersGsuiteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersGsuiteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersGsuiteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteServiceUnavailable creates a PutIdentityprovidersGsuiteServiceUnavailable with default headers values
func NewPutIdentityprovidersGsuiteServiceUnavailable() *PutIdentityprovidersGsuiteServiceUnavailable {
	return &PutIdentityprovidersGsuiteServiceUnavailable{}
}

/*PutIdentityprovidersGsuiteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersGsuiteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGsuiteGatewayTimeout creates a PutIdentityprovidersGsuiteGatewayTimeout with default headers values
func NewPutIdentityprovidersGsuiteGatewayTimeout() *PutIdentityprovidersGsuiteGatewayTimeout {
	return &PutIdentityprovidersGsuiteGatewayTimeout{}
}

/*PutIdentityprovidersGsuiteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutIdentityprovidersGsuiteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/gsuite][%d] putIdentityprovidersGsuiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGsuiteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
