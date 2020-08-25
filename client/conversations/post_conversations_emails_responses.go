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

// PostConversationsEmailsReader is a Reader for the PostConversationsEmails structure.
type PostConversationsEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsEmailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsEmailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsEmailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsEmailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsEmailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsEmailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsEmailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsEmailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsEmailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsEmailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsEmailsOK creates a PostConversationsEmailsOK with default headers values
func NewPostConversationsEmailsOK() *PostConversationsEmailsOK {
	return &PostConversationsEmailsOK{}
}

/*PostConversationsEmailsOK handles this case with default header values.

successful operation
*/
type PostConversationsEmailsOK struct {
	Payload *models.EmailConversation
}

func (o *PostConversationsEmailsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsOK  %+v", 200, o.Payload)
}

func (o *PostConversationsEmailsOK) GetPayload() *models.EmailConversation {
	return o.Payload
}

func (o *PostConversationsEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailConversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsBadRequest creates a PostConversationsEmailsBadRequest with default headers values
func NewPostConversationsEmailsBadRequest() *PostConversationsEmailsBadRequest {
	return &PostConversationsEmailsBadRequest{}
}

/*PostConversationsEmailsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsEmailsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsEmailsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsUnauthorized creates a PostConversationsEmailsUnauthorized with default headers values
func NewPostConversationsEmailsUnauthorized() *PostConversationsEmailsUnauthorized {
	return &PostConversationsEmailsUnauthorized{}
}

/*PostConversationsEmailsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsEmailsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsEmailsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsForbidden creates a PostConversationsEmailsForbidden with default headers values
func NewPostConversationsEmailsForbidden() *PostConversationsEmailsForbidden {
	return &PostConversationsEmailsForbidden{}
}

/*PostConversationsEmailsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsEmailsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsEmailsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsNotFound creates a PostConversationsEmailsNotFound with default headers values
func NewPostConversationsEmailsNotFound() *PostConversationsEmailsNotFound {
	return &PostConversationsEmailsNotFound{}
}

/*PostConversationsEmailsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsEmailsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsEmailsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsRequestEntityTooLarge creates a PostConversationsEmailsRequestEntityTooLarge with default headers values
func NewPostConversationsEmailsRequestEntityTooLarge() *PostConversationsEmailsRequestEntityTooLarge {
	return &PostConversationsEmailsRequestEntityTooLarge{}
}

/*PostConversationsEmailsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsEmailsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsEmailsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsUnsupportedMediaType creates a PostConversationsEmailsUnsupportedMediaType with default headers values
func NewPostConversationsEmailsUnsupportedMediaType() *PostConversationsEmailsUnsupportedMediaType {
	return &PostConversationsEmailsUnsupportedMediaType{}
}

/*PostConversationsEmailsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsEmailsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsEmailsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsTooManyRequests creates a PostConversationsEmailsTooManyRequests with default headers values
func NewPostConversationsEmailsTooManyRequests() *PostConversationsEmailsTooManyRequests {
	return &PostConversationsEmailsTooManyRequests{}
}

/*PostConversationsEmailsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsEmailsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsEmailsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsInternalServerError creates a PostConversationsEmailsInternalServerError with default headers values
func NewPostConversationsEmailsInternalServerError() *PostConversationsEmailsInternalServerError {
	return &PostConversationsEmailsInternalServerError{}
}

/*PostConversationsEmailsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsEmailsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsEmailsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsServiceUnavailable creates a PostConversationsEmailsServiceUnavailable with default headers values
func NewPostConversationsEmailsServiceUnavailable() *PostConversationsEmailsServiceUnavailable {
	return &PostConversationsEmailsServiceUnavailable{}
}

/*PostConversationsEmailsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsEmailsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsEmailsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsGatewayTimeout creates a PostConversationsEmailsGatewayTimeout with default headers values
func NewPostConversationsEmailsGatewayTimeout() *PostConversationsEmailsGatewayTimeout {
	return &PostConversationsEmailsGatewayTimeout{}
}

/*PostConversationsEmailsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsEmailsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails][%d] postConversationsEmailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsEmailsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}