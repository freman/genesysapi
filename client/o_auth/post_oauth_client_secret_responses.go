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

// PostOauthClientSecretReader is a Reader for the PostOauthClientSecret structure.
type PostOauthClientSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauthClientSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOauthClientSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOauthClientSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOauthClientSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOauthClientSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOauthClientSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOauthClientSecretRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOauthClientSecretUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOauthClientSecretTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOauthClientSecretInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOauthClientSecretServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOauthClientSecretGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOauthClientSecretOK creates a PostOauthClientSecretOK with default headers values
func NewPostOauthClientSecretOK() *PostOauthClientSecretOK {
	return &PostOauthClientSecretOK{}
}

/*PostOauthClientSecretOK handles this case with default header values.

successful operation
*/
type PostOauthClientSecretOK struct {
	Payload *models.OAuthClient
}

func (o *PostOauthClientSecretOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretOK  %+v", 200, o.Payload)
}

func (o *PostOauthClientSecretOK) GetPayload() *models.OAuthClient {
	return o.Payload
}

func (o *PostOauthClientSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthClient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretBadRequest creates a PostOauthClientSecretBadRequest with default headers values
func NewPostOauthClientSecretBadRequest() *PostOauthClientSecretBadRequest {
	return &PostOauthClientSecretBadRequest{}
}

/*PostOauthClientSecretBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOauthClientSecretBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretBadRequest  %+v", 400, o.Payload)
}

func (o *PostOauthClientSecretBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretUnauthorized creates a PostOauthClientSecretUnauthorized with default headers values
func NewPostOauthClientSecretUnauthorized() *PostOauthClientSecretUnauthorized {
	return &PostOauthClientSecretUnauthorized{}
}

/*PostOauthClientSecretUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOauthClientSecretUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOauthClientSecretUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretForbidden creates a PostOauthClientSecretForbidden with default headers values
func NewPostOauthClientSecretForbidden() *PostOauthClientSecretForbidden {
	return &PostOauthClientSecretForbidden{}
}

/*PostOauthClientSecretForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOauthClientSecretForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretForbidden  %+v", 403, o.Payload)
}

func (o *PostOauthClientSecretForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretNotFound creates a PostOauthClientSecretNotFound with default headers values
func NewPostOauthClientSecretNotFound() *PostOauthClientSecretNotFound {
	return &PostOauthClientSecretNotFound{}
}

/*PostOauthClientSecretNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOauthClientSecretNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretNotFound  %+v", 404, o.Payload)
}

func (o *PostOauthClientSecretNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretRequestEntityTooLarge creates a PostOauthClientSecretRequestEntityTooLarge with default headers values
func NewPostOauthClientSecretRequestEntityTooLarge() *PostOauthClientSecretRequestEntityTooLarge {
	return &PostOauthClientSecretRequestEntityTooLarge{}
}

/*PostOauthClientSecretRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOauthClientSecretRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOauthClientSecretRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretUnsupportedMediaType creates a PostOauthClientSecretUnsupportedMediaType with default headers values
func NewPostOauthClientSecretUnsupportedMediaType() *PostOauthClientSecretUnsupportedMediaType {
	return &PostOauthClientSecretUnsupportedMediaType{}
}

/*PostOauthClientSecretUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOauthClientSecretUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOauthClientSecretUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretTooManyRequests creates a PostOauthClientSecretTooManyRequests with default headers values
func NewPostOauthClientSecretTooManyRequests() *PostOauthClientSecretTooManyRequests {
	return &PostOauthClientSecretTooManyRequests{}
}

/*PostOauthClientSecretTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOauthClientSecretTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOauthClientSecretTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretInternalServerError creates a PostOauthClientSecretInternalServerError with default headers values
func NewPostOauthClientSecretInternalServerError() *PostOauthClientSecretInternalServerError {
	return &PostOauthClientSecretInternalServerError{}
}

/*PostOauthClientSecretInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOauthClientSecretInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOauthClientSecretInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretServiceUnavailable creates a PostOauthClientSecretServiceUnavailable with default headers values
func NewPostOauthClientSecretServiceUnavailable() *PostOauthClientSecretServiceUnavailable {
	return &PostOauthClientSecretServiceUnavailable{}
}

/*PostOauthClientSecretServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOauthClientSecretServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOauthClientSecretServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientSecretGatewayTimeout creates a PostOauthClientSecretGatewayTimeout with default headers values
func NewPostOauthClientSecretGatewayTimeout() *PostOauthClientSecretGatewayTimeout {
	return &PostOauthClientSecretGatewayTimeout{}
}

/*PostOauthClientSecretGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOauthClientSecretGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientSecretGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/secret][%d] postOauthClientSecretGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOauthClientSecretGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientSecretGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}