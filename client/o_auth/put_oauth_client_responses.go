// Code generated by go-swagger; DO NOT EDIT.

package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOauthClientReader is a Reader for the PutOauthClient structure.
type PutOauthClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOauthClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOauthClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOauthClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOauthClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOauthClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOauthClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOauthClientRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOauthClientUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOauthClientTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOauthClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOauthClientServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOauthClientGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOauthClientOK creates a PutOauthClientOK with default headers values
func NewPutOauthClientOK() *PutOauthClientOK {
	return &PutOauthClientOK{}
}

/*PutOauthClientOK handles this case with default header values.

successful operation
*/
type PutOauthClientOK struct {
	Payload *models.OAuthClient
}

func (o *PutOauthClientOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientOK  %+v", 200, o.Payload)
}

func (o *PutOauthClientOK) GetPayload() *models.OAuthClient {
	return o.Payload
}

func (o *PutOauthClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthClient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientBadRequest creates a PutOauthClientBadRequest with default headers values
func NewPutOauthClientBadRequest() *PutOauthClientBadRequest {
	return &PutOauthClientBadRequest{}
}

/*PutOauthClientBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOauthClientBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientBadRequest  %+v", 400, o.Payload)
}

func (o *PutOauthClientBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientUnauthorized creates a PutOauthClientUnauthorized with default headers values
func NewPutOauthClientUnauthorized() *PutOauthClientUnauthorized {
	return &PutOauthClientUnauthorized{}
}

/*PutOauthClientUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOauthClientUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOauthClientUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientForbidden creates a PutOauthClientForbidden with default headers values
func NewPutOauthClientForbidden() *PutOauthClientForbidden {
	return &PutOauthClientForbidden{}
}

/*PutOauthClientForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOauthClientForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientForbidden  %+v", 403, o.Payload)
}

func (o *PutOauthClientForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientNotFound creates a PutOauthClientNotFound with default headers values
func NewPutOauthClientNotFound() *PutOauthClientNotFound {
	return &PutOauthClientNotFound{}
}

/*PutOauthClientNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOauthClientNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientNotFound  %+v", 404, o.Payload)
}

func (o *PutOauthClientNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientRequestEntityTooLarge creates a PutOauthClientRequestEntityTooLarge with default headers values
func NewPutOauthClientRequestEntityTooLarge() *PutOauthClientRequestEntityTooLarge {
	return &PutOauthClientRequestEntityTooLarge{}
}

/*PutOauthClientRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOauthClientRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOauthClientRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientUnsupportedMediaType creates a PutOauthClientUnsupportedMediaType with default headers values
func NewPutOauthClientUnsupportedMediaType() *PutOauthClientUnsupportedMediaType {
	return &PutOauthClientUnsupportedMediaType{}
}

/*PutOauthClientUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOauthClientUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOauthClientUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientTooManyRequests creates a PutOauthClientTooManyRequests with default headers values
func NewPutOauthClientTooManyRequests() *PutOauthClientTooManyRequests {
	return &PutOauthClientTooManyRequests{}
}

/*PutOauthClientTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutOauthClientTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOauthClientTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientInternalServerError creates a PutOauthClientInternalServerError with default headers values
func NewPutOauthClientInternalServerError() *PutOauthClientInternalServerError {
	return &PutOauthClientInternalServerError{}
}

/*PutOauthClientInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOauthClientInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOauthClientInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientServiceUnavailable creates a PutOauthClientServiceUnavailable with default headers values
func NewPutOauthClientServiceUnavailable() *PutOauthClientServiceUnavailable {
	return &PutOauthClientServiceUnavailable{}
}

/*PutOauthClientServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOauthClientServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOauthClientServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOauthClientGatewayTimeout creates a PutOauthClientGatewayTimeout with default headers values
func NewPutOauthClientGatewayTimeout() *PutOauthClientGatewayTimeout {
	return &PutOauthClientGatewayTimeout{}
}

/*PutOauthClientGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOauthClientGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOauthClientGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/oauth/clients/{clientId}][%d] putOauthClientGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOauthClientGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOauthClientGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}