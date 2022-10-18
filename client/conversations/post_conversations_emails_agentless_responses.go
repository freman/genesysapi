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

// PostConversationsEmailsAgentlessReader is a Reader for the PostConversationsEmailsAgentless structure.
type PostConversationsEmailsAgentlessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsEmailsAgentlessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsEmailsAgentlessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsEmailsAgentlessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsEmailsAgentlessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsEmailsAgentlessForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsEmailsAgentlessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsEmailsAgentlessRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsEmailsAgentlessRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsEmailsAgentlessUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsEmailsAgentlessTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsEmailsAgentlessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsEmailsAgentlessServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsEmailsAgentlessGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsEmailsAgentlessOK creates a PostConversationsEmailsAgentlessOK with default headers values
func NewPostConversationsEmailsAgentlessOK() *PostConversationsEmailsAgentlessOK {
	return &PostConversationsEmailsAgentlessOK{}
}

/*PostConversationsEmailsAgentlessOK handles this case with default header values.

successful operation
*/
type PostConversationsEmailsAgentlessOK struct {
	Payload *models.AgentlessEmailSendResponseDto
}

func (o *PostConversationsEmailsAgentlessOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessOK  %+v", 200, o.Payload)
}

func (o *PostConversationsEmailsAgentlessOK) GetPayload() *models.AgentlessEmailSendResponseDto {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentlessEmailSendResponseDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessBadRequest creates a PostConversationsEmailsAgentlessBadRequest with default headers values
func NewPostConversationsEmailsAgentlessBadRequest() *PostConversationsEmailsAgentlessBadRequest {
	return &PostConversationsEmailsAgentlessBadRequest{}
}

/*PostConversationsEmailsAgentlessBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsEmailsAgentlessBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsEmailsAgentlessBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessUnauthorized creates a PostConversationsEmailsAgentlessUnauthorized with default headers values
func NewPostConversationsEmailsAgentlessUnauthorized() *PostConversationsEmailsAgentlessUnauthorized {
	return &PostConversationsEmailsAgentlessUnauthorized{}
}

/*PostConversationsEmailsAgentlessUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsEmailsAgentlessUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsEmailsAgentlessUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessForbidden creates a PostConversationsEmailsAgentlessForbidden with default headers values
func NewPostConversationsEmailsAgentlessForbidden() *PostConversationsEmailsAgentlessForbidden {
	return &PostConversationsEmailsAgentlessForbidden{}
}

/*PostConversationsEmailsAgentlessForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsEmailsAgentlessForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsEmailsAgentlessForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessNotFound creates a PostConversationsEmailsAgentlessNotFound with default headers values
func NewPostConversationsEmailsAgentlessNotFound() *PostConversationsEmailsAgentlessNotFound {
	return &PostConversationsEmailsAgentlessNotFound{}
}

/*PostConversationsEmailsAgentlessNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsEmailsAgentlessNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsEmailsAgentlessNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessRequestTimeout creates a PostConversationsEmailsAgentlessRequestTimeout with default headers values
func NewPostConversationsEmailsAgentlessRequestTimeout() *PostConversationsEmailsAgentlessRequestTimeout {
	return &PostConversationsEmailsAgentlessRequestTimeout{}
}

/*PostConversationsEmailsAgentlessRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsEmailsAgentlessRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsEmailsAgentlessRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessRequestEntityTooLarge creates a PostConversationsEmailsAgentlessRequestEntityTooLarge with default headers values
func NewPostConversationsEmailsAgentlessRequestEntityTooLarge() *PostConversationsEmailsAgentlessRequestEntityTooLarge {
	return &PostConversationsEmailsAgentlessRequestEntityTooLarge{}
}

/*PostConversationsEmailsAgentlessRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsEmailsAgentlessRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessUnsupportedMediaType creates a PostConversationsEmailsAgentlessUnsupportedMediaType with default headers values
func NewPostConversationsEmailsAgentlessUnsupportedMediaType() *PostConversationsEmailsAgentlessUnsupportedMediaType {
	return &PostConversationsEmailsAgentlessUnsupportedMediaType{}
}

/*PostConversationsEmailsAgentlessUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsEmailsAgentlessUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessTooManyRequests creates a PostConversationsEmailsAgentlessTooManyRequests with default headers values
func NewPostConversationsEmailsAgentlessTooManyRequests() *PostConversationsEmailsAgentlessTooManyRequests {
	return &PostConversationsEmailsAgentlessTooManyRequests{}
}

/*PostConversationsEmailsAgentlessTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsEmailsAgentlessTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsEmailsAgentlessTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessInternalServerError creates a PostConversationsEmailsAgentlessInternalServerError with default headers values
func NewPostConversationsEmailsAgentlessInternalServerError() *PostConversationsEmailsAgentlessInternalServerError {
	return &PostConversationsEmailsAgentlessInternalServerError{}
}

/*PostConversationsEmailsAgentlessInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsEmailsAgentlessInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsEmailsAgentlessInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessServiceUnavailable creates a PostConversationsEmailsAgentlessServiceUnavailable with default headers values
func NewPostConversationsEmailsAgentlessServiceUnavailable() *PostConversationsEmailsAgentlessServiceUnavailable {
	return &PostConversationsEmailsAgentlessServiceUnavailable{}
}

/*PostConversationsEmailsAgentlessServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsEmailsAgentlessServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsEmailsAgentlessServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailsAgentlessGatewayTimeout creates a PostConversationsEmailsAgentlessGatewayTimeout with default headers values
func NewPostConversationsEmailsAgentlessGatewayTimeout() *PostConversationsEmailsAgentlessGatewayTimeout {
	return &PostConversationsEmailsAgentlessGatewayTimeout{}
}

/*PostConversationsEmailsAgentlessGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsEmailsAgentlessGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailsAgentlessGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsEmailsAgentlessGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailsAgentlessGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
