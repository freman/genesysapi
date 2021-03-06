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

// PutIdentityprovidersCicReader is a Reader for the PutIdentityprovidersCic structure.
type PutIdentityprovidersCicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersCicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersCicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersCicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersCicUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersCicForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersCicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersCicRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersCicUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersCicTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersCicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersCicServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersCicGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersCicOK creates a PutIdentityprovidersCicOK with default headers values
func NewPutIdentityprovidersCicOK() *PutIdentityprovidersCicOK {
	return &PutIdentityprovidersCicOK{}
}

/*PutIdentityprovidersCicOK handles this case with default header values.

successful operation
*/
type PutIdentityprovidersCicOK struct {
	Payload *models.OAuthProvider
}

func (o *PutIdentityprovidersCicOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersCicOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersCicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicBadRequest creates a PutIdentityprovidersCicBadRequest with default headers values
func NewPutIdentityprovidersCicBadRequest() *PutIdentityprovidersCicBadRequest {
	return &PutIdentityprovidersCicBadRequest{}
}

/*PutIdentityprovidersCicBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersCicBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersCicBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicUnauthorized creates a PutIdentityprovidersCicUnauthorized with default headers values
func NewPutIdentityprovidersCicUnauthorized() *PutIdentityprovidersCicUnauthorized {
	return &PutIdentityprovidersCicUnauthorized{}
}

/*PutIdentityprovidersCicUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersCicUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersCicUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicForbidden creates a PutIdentityprovidersCicForbidden with default headers values
func NewPutIdentityprovidersCicForbidden() *PutIdentityprovidersCicForbidden {
	return &PutIdentityprovidersCicForbidden{}
}

/*PutIdentityprovidersCicForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersCicForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersCicForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicNotFound creates a PutIdentityprovidersCicNotFound with default headers values
func NewPutIdentityprovidersCicNotFound() *PutIdentityprovidersCicNotFound {
	return &PutIdentityprovidersCicNotFound{}
}

/*PutIdentityprovidersCicNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersCicNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersCicNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicRequestEntityTooLarge creates a PutIdentityprovidersCicRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersCicRequestEntityTooLarge() *PutIdentityprovidersCicRequestEntityTooLarge {
	return &PutIdentityprovidersCicRequestEntityTooLarge{}
}

/*PutIdentityprovidersCicRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutIdentityprovidersCicRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersCicRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicUnsupportedMediaType creates a PutIdentityprovidersCicUnsupportedMediaType with default headers values
func NewPutIdentityprovidersCicUnsupportedMediaType() *PutIdentityprovidersCicUnsupportedMediaType {
	return &PutIdentityprovidersCicUnsupportedMediaType{}
}

/*PutIdentityprovidersCicUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersCicUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersCicUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicTooManyRequests creates a PutIdentityprovidersCicTooManyRequests with default headers values
func NewPutIdentityprovidersCicTooManyRequests() *PutIdentityprovidersCicTooManyRequests {
	return &PutIdentityprovidersCicTooManyRequests{}
}

/*PutIdentityprovidersCicTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutIdentityprovidersCicTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersCicTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicInternalServerError creates a PutIdentityprovidersCicInternalServerError with default headers values
func NewPutIdentityprovidersCicInternalServerError() *PutIdentityprovidersCicInternalServerError {
	return &PutIdentityprovidersCicInternalServerError{}
}

/*PutIdentityprovidersCicInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersCicInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersCicInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicServiceUnavailable creates a PutIdentityprovidersCicServiceUnavailable with default headers values
func NewPutIdentityprovidersCicServiceUnavailable() *PutIdentityprovidersCicServiceUnavailable {
	return &PutIdentityprovidersCicServiceUnavailable{}
}

/*PutIdentityprovidersCicServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersCicServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersCicServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersCicGatewayTimeout creates a PutIdentityprovidersCicGatewayTimeout with default headers values
func NewPutIdentityprovidersCicGatewayTimeout() *PutIdentityprovidersCicGatewayTimeout {
	return &PutIdentityprovidersCicGatewayTimeout{}
}

/*PutIdentityprovidersCicGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutIdentityprovidersCicGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersCicGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/cic][%d] putIdentityprovidersCicGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersCicGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersCicGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
