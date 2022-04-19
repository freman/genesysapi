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

// PutIdentityprovidersGenericReader is a Reader for the PutIdentityprovidersGeneric structure.
type PutIdentityprovidersGenericReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersGenericReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersGenericOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersGenericBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersGenericUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersGenericForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersGenericNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutIdentityprovidersGenericRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersGenericRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersGenericUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersGenericTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersGenericInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersGenericServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersGenericGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersGenericOK creates a PutIdentityprovidersGenericOK with default headers values
func NewPutIdentityprovidersGenericOK() *PutIdentityprovidersGenericOK {
	return &PutIdentityprovidersGenericOK{}
}

/*PutIdentityprovidersGenericOK handles this case with default header values.

successful operation
*/
type PutIdentityprovidersGenericOK struct {
	Payload *models.OAuthProvider
}

func (o *PutIdentityprovidersGenericOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersGenericOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersGenericOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericBadRequest creates a PutIdentityprovidersGenericBadRequest with default headers values
func NewPutIdentityprovidersGenericBadRequest() *PutIdentityprovidersGenericBadRequest {
	return &PutIdentityprovidersGenericBadRequest{}
}

/*PutIdentityprovidersGenericBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersGenericBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersGenericBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericUnauthorized creates a PutIdentityprovidersGenericUnauthorized with default headers values
func NewPutIdentityprovidersGenericUnauthorized() *PutIdentityprovidersGenericUnauthorized {
	return &PutIdentityprovidersGenericUnauthorized{}
}

/*PutIdentityprovidersGenericUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersGenericUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersGenericUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericForbidden creates a PutIdentityprovidersGenericForbidden with default headers values
func NewPutIdentityprovidersGenericForbidden() *PutIdentityprovidersGenericForbidden {
	return &PutIdentityprovidersGenericForbidden{}
}

/*PutIdentityprovidersGenericForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersGenericForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersGenericForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericNotFound creates a PutIdentityprovidersGenericNotFound with default headers values
func NewPutIdentityprovidersGenericNotFound() *PutIdentityprovidersGenericNotFound {
	return &PutIdentityprovidersGenericNotFound{}
}

/*PutIdentityprovidersGenericNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersGenericNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersGenericNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericRequestTimeout creates a PutIdentityprovidersGenericRequestTimeout with default headers values
func NewPutIdentityprovidersGenericRequestTimeout() *PutIdentityprovidersGenericRequestTimeout {
	return &PutIdentityprovidersGenericRequestTimeout{}
}

/*PutIdentityprovidersGenericRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutIdentityprovidersGenericRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutIdentityprovidersGenericRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericRequestEntityTooLarge creates a PutIdentityprovidersGenericRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersGenericRequestEntityTooLarge() *PutIdentityprovidersGenericRequestEntityTooLarge {
	return &PutIdentityprovidersGenericRequestEntityTooLarge{}
}

/*PutIdentityprovidersGenericRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutIdentityprovidersGenericRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersGenericRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericUnsupportedMediaType creates a PutIdentityprovidersGenericUnsupportedMediaType with default headers values
func NewPutIdentityprovidersGenericUnsupportedMediaType() *PutIdentityprovidersGenericUnsupportedMediaType {
	return &PutIdentityprovidersGenericUnsupportedMediaType{}
}

/*PutIdentityprovidersGenericUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersGenericUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersGenericUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericTooManyRequests creates a PutIdentityprovidersGenericTooManyRequests with default headers values
func NewPutIdentityprovidersGenericTooManyRequests() *PutIdentityprovidersGenericTooManyRequests {
	return &PutIdentityprovidersGenericTooManyRequests{}
}

/*PutIdentityprovidersGenericTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutIdentityprovidersGenericTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersGenericTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericInternalServerError creates a PutIdentityprovidersGenericInternalServerError with default headers values
func NewPutIdentityprovidersGenericInternalServerError() *PutIdentityprovidersGenericInternalServerError {
	return &PutIdentityprovidersGenericInternalServerError{}
}

/*PutIdentityprovidersGenericInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersGenericInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersGenericInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericServiceUnavailable creates a PutIdentityprovidersGenericServiceUnavailable with default headers values
func NewPutIdentityprovidersGenericServiceUnavailable() *PutIdentityprovidersGenericServiceUnavailable {
	return &PutIdentityprovidersGenericServiceUnavailable{}
}

/*PutIdentityprovidersGenericServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersGenericServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersGenericServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersGenericGatewayTimeout creates a PutIdentityprovidersGenericGatewayTimeout with default headers values
func NewPutIdentityprovidersGenericGatewayTimeout() *PutIdentityprovidersGenericGatewayTimeout {
	return &PutIdentityprovidersGenericGatewayTimeout{}
}

/*PutIdentityprovidersGenericGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutIdentityprovidersGenericGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersGenericGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/generic][%d] putIdentityprovidersGenericGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersGenericGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersGenericGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
