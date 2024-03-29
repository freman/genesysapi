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

/*
PostConversationsEmailsAgentlessOK describes a response with status code 200, with default header values.

successful operation
*/
type PostConversationsEmailsAgentlessOK struct {
	Payload *models.AgentlessEmailSendResponseDto
}

// IsSuccess returns true when this post conversations emails agentless o k response has a 2xx status code
func (o *PostConversationsEmailsAgentlessOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations emails agentless o k response has a 3xx status code
func (o *PostConversationsEmailsAgentlessOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless o k response has a 4xx status code
func (o *PostConversationsEmailsAgentlessOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations emails agentless o k response has a 5xx status code
func (o *PostConversationsEmailsAgentlessOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless o k response a status code equal to that given
func (o *PostConversationsEmailsAgentlessOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostConversationsEmailsAgentlessOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessOK  %+v", 200, o.Payload)
}

func (o *PostConversationsEmailsAgentlessOK) String() string {
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

/*
PostConversationsEmailsAgentlessBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsEmailsAgentlessBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless bad request response has a 2xx status code
func (o *PostConversationsEmailsAgentlessBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless bad request response has a 3xx status code
func (o *PostConversationsEmailsAgentlessBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless bad request response has a 4xx status code
func (o *PostConversationsEmailsAgentlessBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless bad request response has a 5xx status code
func (o *PostConversationsEmailsAgentlessBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless bad request response a status code equal to that given
func (o *PostConversationsEmailsAgentlessBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsEmailsAgentlessBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsEmailsAgentlessBadRequest) String() string {
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

/*
PostConversationsEmailsAgentlessUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsEmailsAgentlessUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless unauthorized response has a 2xx status code
func (o *PostConversationsEmailsAgentlessUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless unauthorized response has a 3xx status code
func (o *PostConversationsEmailsAgentlessUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless unauthorized response has a 4xx status code
func (o *PostConversationsEmailsAgentlessUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless unauthorized response has a 5xx status code
func (o *PostConversationsEmailsAgentlessUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless unauthorized response a status code equal to that given
func (o *PostConversationsEmailsAgentlessUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsEmailsAgentlessUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsEmailsAgentlessUnauthorized) String() string {
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

/*
PostConversationsEmailsAgentlessForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsEmailsAgentlessForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless forbidden response has a 2xx status code
func (o *PostConversationsEmailsAgentlessForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless forbidden response has a 3xx status code
func (o *PostConversationsEmailsAgentlessForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless forbidden response has a 4xx status code
func (o *PostConversationsEmailsAgentlessForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless forbidden response has a 5xx status code
func (o *PostConversationsEmailsAgentlessForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless forbidden response a status code equal to that given
func (o *PostConversationsEmailsAgentlessForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsEmailsAgentlessForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsEmailsAgentlessForbidden) String() string {
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

/*
PostConversationsEmailsAgentlessNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsEmailsAgentlessNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless not found response has a 2xx status code
func (o *PostConversationsEmailsAgentlessNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless not found response has a 3xx status code
func (o *PostConversationsEmailsAgentlessNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless not found response has a 4xx status code
func (o *PostConversationsEmailsAgentlessNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless not found response has a 5xx status code
func (o *PostConversationsEmailsAgentlessNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless not found response a status code equal to that given
func (o *PostConversationsEmailsAgentlessNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsEmailsAgentlessNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsEmailsAgentlessNotFound) String() string {
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

/*
PostConversationsEmailsAgentlessRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsEmailsAgentlessRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless request timeout response has a 2xx status code
func (o *PostConversationsEmailsAgentlessRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless request timeout response has a 3xx status code
func (o *PostConversationsEmailsAgentlessRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless request timeout response has a 4xx status code
func (o *PostConversationsEmailsAgentlessRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless request timeout response has a 5xx status code
func (o *PostConversationsEmailsAgentlessRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless request timeout response a status code equal to that given
func (o *PostConversationsEmailsAgentlessRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsEmailsAgentlessRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsEmailsAgentlessRequestTimeout) String() string {
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

/*
PostConversationsEmailsAgentlessRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationsEmailsAgentlessRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless request entity too large response has a 2xx status code
func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless request entity too large response has a 3xx status code
func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless request entity too large response has a 4xx status code
func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless request entity too large response has a 5xx status code
func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless request entity too large response a status code equal to that given
func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsEmailsAgentlessRequestEntityTooLarge) String() string {
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

/*
PostConversationsEmailsAgentlessUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsEmailsAgentlessUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless unsupported media type response has a 2xx status code
func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless unsupported media type response has a 3xx status code
func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless unsupported media type response has a 4xx status code
func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless unsupported media type response has a 5xx status code
func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless unsupported media type response a status code equal to that given
func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsEmailsAgentlessUnsupportedMediaType) String() string {
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

/*
PostConversationsEmailsAgentlessTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsEmailsAgentlessTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless too many requests response has a 2xx status code
func (o *PostConversationsEmailsAgentlessTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless too many requests response has a 3xx status code
func (o *PostConversationsEmailsAgentlessTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless too many requests response has a 4xx status code
func (o *PostConversationsEmailsAgentlessTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations emails agentless too many requests response has a 5xx status code
func (o *PostConversationsEmailsAgentlessTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations emails agentless too many requests response a status code equal to that given
func (o *PostConversationsEmailsAgentlessTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsEmailsAgentlessTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsEmailsAgentlessTooManyRequests) String() string {
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

/*
PostConversationsEmailsAgentlessInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsEmailsAgentlessInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless internal server error response has a 2xx status code
func (o *PostConversationsEmailsAgentlessInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless internal server error response has a 3xx status code
func (o *PostConversationsEmailsAgentlessInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless internal server error response has a 4xx status code
func (o *PostConversationsEmailsAgentlessInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations emails agentless internal server error response has a 5xx status code
func (o *PostConversationsEmailsAgentlessInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations emails agentless internal server error response a status code equal to that given
func (o *PostConversationsEmailsAgentlessInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsEmailsAgentlessInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsEmailsAgentlessInternalServerError) String() string {
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

/*
PostConversationsEmailsAgentlessServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsEmailsAgentlessServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless service unavailable response has a 2xx status code
func (o *PostConversationsEmailsAgentlessServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless service unavailable response has a 3xx status code
func (o *PostConversationsEmailsAgentlessServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless service unavailable response has a 4xx status code
func (o *PostConversationsEmailsAgentlessServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations emails agentless service unavailable response has a 5xx status code
func (o *PostConversationsEmailsAgentlessServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations emails agentless service unavailable response a status code equal to that given
func (o *PostConversationsEmailsAgentlessServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsEmailsAgentlessServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsEmailsAgentlessServiceUnavailable) String() string {
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

/*
PostConversationsEmailsAgentlessGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsEmailsAgentlessGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations emails agentless gateway timeout response has a 2xx status code
func (o *PostConversationsEmailsAgentlessGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations emails agentless gateway timeout response has a 3xx status code
func (o *PostConversationsEmailsAgentlessGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations emails agentless gateway timeout response has a 4xx status code
func (o *PostConversationsEmailsAgentlessGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations emails agentless gateway timeout response has a 5xx status code
func (o *PostConversationsEmailsAgentlessGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations emails agentless gateway timeout response a status code equal to that given
func (o *PostConversationsEmailsAgentlessGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsEmailsAgentlessGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/agentless][%d] postConversationsEmailsAgentlessGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsEmailsAgentlessGatewayTimeout) String() string {
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
