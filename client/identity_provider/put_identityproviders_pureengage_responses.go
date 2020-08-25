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

// PutIdentityprovidersPureengageReader is a Reader for the PutIdentityprovidersPureengage structure.
type PutIdentityprovidersPureengageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutIdentityprovidersPureengageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutIdentityprovidersPureengageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutIdentityprovidersPureengageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutIdentityprovidersPureengageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutIdentityprovidersPureengageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutIdentityprovidersPureengageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutIdentityprovidersPureengageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutIdentityprovidersPureengageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutIdentityprovidersPureengageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutIdentityprovidersPureengageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutIdentityprovidersPureengageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutIdentityprovidersPureengageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutIdentityprovidersPureengageOK creates a PutIdentityprovidersPureengageOK with default headers values
func NewPutIdentityprovidersPureengageOK() *PutIdentityprovidersPureengageOK {
	return &PutIdentityprovidersPureengageOK{}
}

/*PutIdentityprovidersPureengageOK handles this case with default header values.

successful operation
*/
type PutIdentityprovidersPureengageOK struct {
	Payload *models.OAuthProvider
}

func (o *PutIdentityprovidersPureengageOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageOK  %+v", 200, o.Payload)
}

func (o *PutIdentityprovidersPureengageOK) GetPayload() *models.OAuthProvider {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageBadRequest creates a PutIdentityprovidersPureengageBadRequest with default headers values
func NewPutIdentityprovidersPureengageBadRequest() *PutIdentityprovidersPureengageBadRequest {
	return &PutIdentityprovidersPureengageBadRequest{}
}

/*PutIdentityprovidersPureengageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutIdentityprovidersPureengageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageBadRequest  %+v", 400, o.Payload)
}

func (o *PutIdentityprovidersPureengageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageUnauthorized creates a PutIdentityprovidersPureengageUnauthorized with default headers values
func NewPutIdentityprovidersPureengageUnauthorized() *PutIdentityprovidersPureengageUnauthorized {
	return &PutIdentityprovidersPureengageUnauthorized{}
}

/*PutIdentityprovidersPureengageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutIdentityprovidersPureengageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageUnauthorized  %+v", 401, o.Payload)
}

func (o *PutIdentityprovidersPureengageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageForbidden creates a PutIdentityprovidersPureengageForbidden with default headers values
func NewPutIdentityprovidersPureengageForbidden() *PutIdentityprovidersPureengageForbidden {
	return &PutIdentityprovidersPureengageForbidden{}
}

/*PutIdentityprovidersPureengageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutIdentityprovidersPureengageForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageForbidden  %+v", 403, o.Payload)
}

func (o *PutIdentityprovidersPureengageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageNotFound creates a PutIdentityprovidersPureengageNotFound with default headers values
func NewPutIdentityprovidersPureengageNotFound() *PutIdentityprovidersPureengageNotFound {
	return &PutIdentityprovidersPureengageNotFound{}
}

/*PutIdentityprovidersPureengageNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutIdentityprovidersPureengageNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageNotFound  %+v", 404, o.Payload)
}

func (o *PutIdentityprovidersPureengageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageRequestEntityTooLarge creates a PutIdentityprovidersPureengageRequestEntityTooLarge with default headers values
func NewPutIdentityprovidersPureengageRequestEntityTooLarge() *PutIdentityprovidersPureengageRequestEntityTooLarge {
	return &PutIdentityprovidersPureengageRequestEntityTooLarge{}
}

/*PutIdentityprovidersPureengageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutIdentityprovidersPureengageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageUnsupportedMediaType creates a PutIdentityprovidersPureengageUnsupportedMediaType with default headers values
func NewPutIdentityprovidersPureengageUnsupportedMediaType() *PutIdentityprovidersPureengageUnsupportedMediaType {
	return &PutIdentityprovidersPureengageUnsupportedMediaType{}
}

/*PutIdentityprovidersPureengageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutIdentityprovidersPureengageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageTooManyRequests creates a PutIdentityprovidersPureengageTooManyRequests with default headers values
func NewPutIdentityprovidersPureengageTooManyRequests() *PutIdentityprovidersPureengageTooManyRequests {
	return &PutIdentityprovidersPureengageTooManyRequests{}
}

/*PutIdentityprovidersPureengageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutIdentityprovidersPureengageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutIdentityprovidersPureengageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageInternalServerError creates a PutIdentityprovidersPureengageInternalServerError with default headers values
func NewPutIdentityprovidersPureengageInternalServerError() *PutIdentityprovidersPureengageInternalServerError {
	return &PutIdentityprovidersPureengageInternalServerError{}
}

/*PutIdentityprovidersPureengageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutIdentityprovidersPureengageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageInternalServerError  %+v", 500, o.Payload)
}

func (o *PutIdentityprovidersPureengageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageServiceUnavailable creates a PutIdentityprovidersPureengageServiceUnavailable with default headers values
func NewPutIdentityprovidersPureengageServiceUnavailable() *PutIdentityprovidersPureengageServiceUnavailable {
	return &PutIdentityprovidersPureengageServiceUnavailable{}
}

/*PutIdentityprovidersPureengageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutIdentityprovidersPureengageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutIdentityprovidersPureengageGatewayTimeout creates a PutIdentityprovidersPureengageGatewayTimeout with default headers values
func NewPutIdentityprovidersPureengageGatewayTimeout() *PutIdentityprovidersPureengageGatewayTimeout {
	return &PutIdentityprovidersPureengageGatewayTimeout{}
}

/*PutIdentityprovidersPureengageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutIdentityprovidersPureengageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/identityproviders/pureengage][%d] putIdentityprovidersPureengageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutIdentityprovidersPureengageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}