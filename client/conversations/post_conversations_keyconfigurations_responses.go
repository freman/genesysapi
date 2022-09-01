// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostConversationsKeyconfigurationsReader is a Reader for the PostConversationsKeyconfigurations structure.
type PostConversationsKeyconfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsKeyconfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsKeyconfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsKeyconfigurationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsKeyconfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsKeyconfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsKeyconfigurationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsKeyconfigurationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsKeyconfigurationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsKeyconfigurationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsKeyconfigurationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsKeyconfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsKeyconfigurationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsKeyconfigurationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsKeyconfigurationsOK creates a PostConversationsKeyconfigurationsOK with default headers values
func NewPostConversationsKeyconfigurationsOK() *PostConversationsKeyconfigurationsOK {
	return &PostConversationsKeyconfigurationsOK{}
}

/*PostConversationsKeyconfigurationsOK handles this case with default header values.

successful operation
*/
type PostConversationsKeyconfigurationsOK struct {
	Payload *models.ConversationEncryptionConfiguration
}

func (o *PostConversationsKeyconfigurationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsOK  %+v", 200, o.Payload)
}

func (o *PostConversationsKeyconfigurationsOK) GetPayload() *models.ConversationEncryptionConfiguration {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConversationEncryptionConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsBadRequest creates a PostConversationsKeyconfigurationsBadRequest with default headers values
func NewPostConversationsKeyconfigurationsBadRequest() *PostConversationsKeyconfigurationsBadRequest {
	return &PostConversationsKeyconfigurationsBadRequest{}
}

/*PostConversationsKeyconfigurationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsKeyconfigurationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsKeyconfigurationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsUnauthorized creates a PostConversationsKeyconfigurationsUnauthorized with default headers values
func NewPostConversationsKeyconfigurationsUnauthorized() *PostConversationsKeyconfigurationsUnauthorized {
	return &PostConversationsKeyconfigurationsUnauthorized{}
}

/*PostConversationsKeyconfigurationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsKeyconfigurationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsKeyconfigurationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsForbidden creates a PostConversationsKeyconfigurationsForbidden with default headers values
func NewPostConversationsKeyconfigurationsForbidden() *PostConversationsKeyconfigurationsForbidden {
	return &PostConversationsKeyconfigurationsForbidden{}
}

/*PostConversationsKeyconfigurationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsKeyconfigurationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsKeyconfigurationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsNotFound creates a PostConversationsKeyconfigurationsNotFound with default headers values
func NewPostConversationsKeyconfigurationsNotFound() *PostConversationsKeyconfigurationsNotFound {
	return &PostConversationsKeyconfigurationsNotFound{}
}

/*PostConversationsKeyconfigurationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsKeyconfigurationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsKeyconfigurationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsRequestTimeout creates a PostConversationsKeyconfigurationsRequestTimeout with default headers values
func NewPostConversationsKeyconfigurationsRequestTimeout() *PostConversationsKeyconfigurationsRequestTimeout {
	return &PostConversationsKeyconfigurationsRequestTimeout{}
}

/*PostConversationsKeyconfigurationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsKeyconfigurationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsKeyconfigurationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsRequestEntityTooLarge creates a PostConversationsKeyconfigurationsRequestEntityTooLarge with default headers values
func NewPostConversationsKeyconfigurationsRequestEntityTooLarge() *PostConversationsKeyconfigurationsRequestEntityTooLarge {
	return &PostConversationsKeyconfigurationsRequestEntityTooLarge{}
}

/*PostConversationsKeyconfigurationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsKeyconfigurationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsKeyconfigurationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsUnsupportedMediaType creates a PostConversationsKeyconfigurationsUnsupportedMediaType with default headers values
func NewPostConversationsKeyconfigurationsUnsupportedMediaType() *PostConversationsKeyconfigurationsUnsupportedMediaType {
	return &PostConversationsKeyconfigurationsUnsupportedMediaType{}
}

/*PostConversationsKeyconfigurationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsKeyconfigurationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsKeyconfigurationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsTooManyRequests creates a PostConversationsKeyconfigurationsTooManyRequests with default headers values
func NewPostConversationsKeyconfigurationsTooManyRequests() *PostConversationsKeyconfigurationsTooManyRequests {
	return &PostConversationsKeyconfigurationsTooManyRequests{}
}

/*PostConversationsKeyconfigurationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsKeyconfigurationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsKeyconfigurationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsInternalServerError creates a PostConversationsKeyconfigurationsInternalServerError with default headers values
func NewPostConversationsKeyconfigurationsInternalServerError() *PostConversationsKeyconfigurationsInternalServerError {
	return &PostConversationsKeyconfigurationsInternalServerError{}
}

/*PostConversationsKeyconfigurationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsKeyconfigurationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsKeyconfigurationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsServiceUnavailable creates a PostConversationsKeyconfigurationsServiceUnavailable with default headers values
func NewPostConversationsKeyconfigurationsServiceUnavailable() *PostConversationsKeyconfigurationsServiceUnavailable {
	return &PostConversationsKeyconfigurationsServiceUnavailable{}
}

/*PostConversationsKeyconfigurationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsKeyconfigurationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsKeyconfigurationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsKeyconfigurationsGatewayTimeout creates a PostConversationsKeyconfigurationsGatewayTimeout with default headers values
func NewPostConversationsKeyconfigurationsGatewayTimeout() *PostConversationsKeyconfigurationsGatewayTimeout {
	return &PostConversationsKeyconfigurationsGatewayTimeout{}
}

/*PostConversationsKeyconfigurationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsKeyconfigurationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsKeyconfigurationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/keyconfigurations][%d] postConversationsKeyconfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsKeyconfigurationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsKeyconfigurationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}