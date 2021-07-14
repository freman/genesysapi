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

// PostOauthClientsReader is a Reader for the PostOauthClients structure.
type PostOauthClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauthClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOauthClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOauthClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOauthClientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOauthClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOauthClientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOauthClientsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOauthClientsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOauthClientsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOauthClientsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOauthClientsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOauthClientsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOauthClientsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOauthClientsOK creates a PostOauthClientsOK with default headers values
func NewPostOauthClientsOK() *PostOauthClientsOK {
	return &PostOauthClientsOK{}
}

/*PostOauthClientsOK handles this case with default header values.

successful operation
*/
type PostOauthClientsOK struct {
	Payload *models.OAuthClient
}

func (o *PostOauthClientsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsOK  %+v", 200, o.Payload)
}

func (o *PostOauthClientsOK) GetPayload() *models.OAuthClient {
	return o.Payload
}

func (o *PostOauthClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthClient)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsBadRequest creates a PostOauthClientsBadRequest with default headers values
func NewPostOauthClientsBadRequest() *PostOauthClientsBadRequest {
	return &PostOauthClientsBadRequest{}
}

/*PostOauthClientsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOauthClientsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOauthClientsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsUnauthorized creates a PostOauthClientsUnauthorized with default headers values
func NewPostOauthClientsUnauthorized() *PostOauthClientsUnauthorized {
	return &PostOauthClientsUnauthorized{}
}

/*PostOauthClientsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOauthClientsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOauthClientsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsForbidden creates a PostOauthClientsForbidden with default headers values
func NewPostOauthClientsForbidden() *PostOauthClientsForbidden {
	return &PostOauthClientsForbidden{}
}

/*PostOauthClientsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOauthClientsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsForbidden  %+v", 403, o.Payload)
}

func (o *PostOauthClientsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsNotFound creates a PostOauthClientsNotFound with default headers values
func NewPostOauthClientsNotFound() *PostOauthClientsNotFound {
	return &PostOauthClientsNotFound{}
}

/*PostOauthClientsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOauthClientsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsNotFound  %+v", 404, o.Payload)
}

func (o *PostOauthClientsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsRequestTimeout creates a PostOauthClientsRequestTimeout with default headers values
func NewPostOauthClientsRequestTimeout() *PostOauthClientsRequestTimeout {
	return &PostOauthClientsRequestTimeout{}
}

/*PostOauthClientsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOauthClientsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOauthClientsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsRequestEntityTooLarge creates a PostOauthClientsRequestEntityTooLarge with default headers values
func NewPostOauthClientsRequestEntityTooLarge() *PostOauthClientsRequestEntityTooLarge {
	return &PostOauthClientsRequestEntityTooLarge{}
}

/*PostOauthClientsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOauthClientsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOauthClientsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsUnsupportedMediaType creates a PostOauthClientsUnsupportedMediaType with default headers values
func NewPostOauthClientsUnsupportedMediaType() *PostOauthClientsUnsupportedMediaType {
	return &PostOauthClientsUnsupportedMediaType{}
}

/*PostOauthClientsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOauthClientsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOauthClientsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsTooManyRequests creates a PostOauthClientsTooManyRequests with default headers values
func NewPostOauthClientsTooManyRequests() *PostOauthClientsTooManyRequests {
	return &PostOauthClientsTooManyRequests{}
}

/*PostOauthClientsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOauthClientsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOauthClientsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsInternalServerError creates a PostOauthClientsInternalServerError with default headers values
func NewPostOauthClientsInternalServerError() *PostOauthClientsInternalServerError {
	return &PostOauthClientsInternalServerError{}
}

/*PostOauthClientsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOauthClientsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOauthClientsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsServiceUnavailable creates a PostOauthClientsServiceUnavailable with default headers values
func NewPostOauthClientsServiceUnavailable() *PostOauthClientsServiceUnavailable {
	return &PostOauthClientsServiceUnavailable{}
}

/*PostOauthClientsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOauthClientsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOauthClientsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientsGatewayTimeout creates a PostOauthClientsGatewayTimeout with default headers values
func NewPostOauthClientsGatewayTimeout() *PostOauthClientsGatewayTimeout {
	return &PostOauthClientsGatewayTimeout{}
}

/*PostOauthClientsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOauthClientsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOauthClientsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients][%d] postOauthClientsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOauthClientsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
