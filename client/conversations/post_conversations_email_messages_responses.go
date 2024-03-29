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

// PostConversationsEmailMessagesReader is a Reader for the PostConversationsEmailMessages structure.
type PostConversationsEmailMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsEmailMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsEmailMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsEmailMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsEmailMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsEmailMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsEmailMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsEmailMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostConversationsEmailMessagesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsEmailMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsEmailMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsEmailMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsEmailMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsEmailMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsEmailMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsEmailMessagesOK creates a PostConversationsEmailMessagesOK with default headers values
func NewPostConversationsEmailMessagesOK() *PostConversationsEmailMessagesOK {
	return &PostConversationsEmailMessagesOK{}
}

/*
PostConversationsEmailMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostConversationsEmailMessagesOK struct {
	Payload *models.EmailMessageReply
}

// IsSuccess returns true when this post conversations email messages o k response has a 2xx status code
func (o *PostConversationsEmailMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations email messages o k response has a 3xx status code
func (o *PostConversationsEmailMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages o k response has a 4xx status code
func (o *PostConversationsEmailMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations email messages o k response has a 5xx status code
func (o *PostConversationsEmailMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages o k response a status code equal to that given
func (o *PostConversationsEmailMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostConversationsEmailMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsEmailMessagesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsEmailMessagesOK) GetPayload() *models.EmailMessageReply {
	return o.Payload
}

func (o *PostConversationsEmailMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailMessageReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesBadRequest creates a PostConversationsEmailMessagesBadRequest with default headers values
func NewPostConversationsEmailMessagesBadRequest() *PostConversationsEmailMessagesBadRequest {
	return &PostConversationsEmailMessagesBadRequest{}
}

/*
PostConversationsEmailMessagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsEmailMessagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages bad request response has a 2xx status code
func (o *PostConversationsEmailMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages bad request response has a 3xx status code
func (o *PostConversationsEmailMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages bad request response has a 4xx status code
func (o *PostConversationsEmailMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages bad request response has a 5xx status code
func (o *PostConversationsEmailMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages bad request response a status code equal to that given
func (o *PostConversationsEmailMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsEmailMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsEmailMessagesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsEmailMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesUnauthorized creates a PostConversationsEmailMessagesUnauthorized with default headers values
func NewPostConversationsEmailMessagesUnauthorized() *PostConversationsEmailMessagesUnauthorized {
	return &PostConversationsEmailMessagesUnauthorized{}
}

/*
PostConversationsEmailMessagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsEmailMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages unauthorized response has a 2xx status code
func (o *PostConversationsEmailMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages unauthorized response has a 3xx status code
func (o *PostConversationsEmailMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages unauthorized response has a 4xx status code
func (o *PostConversationsEmailMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages unauthorized response has a 5xx status code
func (o *PostConversationsEmailMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages unauthorized response a status code equal to that given
func (o *PostConversationsEmailMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsEmailMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsEmailMessagesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsEmailMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesForbidden creates a PostConversationsEmailMessagesForbidden with default headers values
func NewPostConversationsEmailMessagesForbidden() *PostConversationsEmailMessagesForbidden {
	return &PostConversationsEmailMessagesForbidden{}
}

/*
PostConversationsEmailMessagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsEmailMessagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages forbidden response has a 2xx status code
func (o *PostConversationsEmailMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages forbidden response has a 3xx status code
func (o *PostConversationsEmailMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages forbidden response has a 4xx status code
func (o *PostConversationsEmailMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages forbidden response has a 5xx status code
func (o *PostConversationsEmailMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages forbidden response a status code equal to that given
func (o *PostConversationsEmailMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsEmailMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsEmailMessagesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsEmailMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesNotFound creates a PostConversationsEmailMessagesNotFound with default headers values
func NewPostConversationsEmailMessagesNotFound() *PostConversationsEmailMessagesNotFound {
	return &PostConversationsEmailMessagesNotFound{}
}

/*
PostConversationsEmailMessagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsEmailMessagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages not found response has a 2xx status code
func (o *PostConversationsEmailMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages not found response has a 3xx status code
func (o *PostConversationsEmailMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages not found response has a 4xx status code
func (o *PostConversationsEmailMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages not found response has a 5xx status code
func (o *PostConversationsEmailMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages not found response a status code equal to that given
func (o *PostConversationsEmailMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsEmailMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsEmailMessagesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsEmailMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesRequestTimeout creates a PostConversationsEmailMessagesRequestTimeout with default headers values
func NewPostConversationsEmailMessagesRequestTimeout() *PostConversationsEmailMessagesRequestTimeout {
	return &PostConversationsEmailMessagesRequestTimeout{}
}

/*
PostConversationsEmailMessagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsEmailMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages request timeout response has a 2xx status code
func (o *PostConversationsEmailMessagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages request timeout response has a 3xx status code
func (o *PostConversationsEmailMessagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages request timeout response has a 4xx status code
func (o *PostConversationsEmailMessagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages request timeout response has a 5xx status code
func (o *PostConversationsEmailMessagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages request timeout response a status code equal to that given
func (o *PostConversationsEmailMessagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsEmailMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsEmailMessagesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsEmailMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesConflict creates a PostConversationsEmailMessagesConflict with default headers values
func NewPostConversationsEmailMessagesConflict() *PostConversationsEmailMessagesConflict {
	return &PostConversationsEmailMessagesConflict{}
}

/*
PostConversationsEmailMessagesConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostConversationsEmailMessagesConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages conflict response has a 2xx status code
func (o *PostConversationsEmailMessagesConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages conflict response has a 3xx status code
func (o *PostConversationsEmailMessagesConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages conflict response has a 4xx status code
func (o *PostConversationsEmailMessagesConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages conflict response has a 5xx status code
func (o *PostConversationsEmailMessagesConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages conflict response a status code equal to that given
func (o *PostConversationsEmailMessagesConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostConversationsEmailMessagesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesConflict  %+v", 409, o.Payload)
}

func (o *PostConversationsEmailMessagesConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesConflict  %+v", 409, o.Payload)
}

func (o *PostConversationsEmailMessagesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesRequestEntityTooLarge creates a PostConversationsEmailMessagesRequestEntityTooLarge with default headers values
func NewPostConversationsEmailMessagesRequestEntityTooLarge() *PostConversationsEmailMessagesRequestEntityTooLarge {
	return &PostConversationsEmailMessagesRequestEntityTooLarge{}
}

/*
PostConversationsEmailMessagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationsEmailMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages request entity too large response has a 2xx status code
func (o *PostConversationsEmailMessagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages request entity too large response has a 3xx status code
func (o *PostConversationsEmailMessagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages request entity too large response has a 4xx status code
func (o *PostConversationsEmailMessagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages request entity too large response has a 5xx status code
func (o *PostConversationsEmailMessagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages request entity too large response a status code equal to that given
func (o *PostConversationsEmailMessagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesUnsupportedMediaType creates a PostConversationsEmailMessagesUnsupportedMediaType with default headers values
func NewPostConversationsEmailMessagesUnsupportedMediaType() *PostConversationsEmailMessagesUnsupportedMediaType {
	return &PostConversationsEmailMessagesUnsupportedMediaType{}
}

/*
PostConversationsEmailMessagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsEmailMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages unsupported media type response has a 2xx status code
func (o *PostConversationsEmailMessagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages unsupported media type response has a 3xx status code
func (o *PostConversationsEmailMessagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages unsupported media type response has a 4xx status code
func (o *PostConversationsEmailMessagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages unsupported media type response has a 5xx status code
func (o *PostConversationsEmailMessagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages unsupported media type response a status code equal to that given
func (o *PostConversationsEmailMessagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesTooManyRequests creates a PostConversationsEmailMessagesTooManyRequests with default headers values
func NewPostConversationsEmailMessagesTooManyRequests() *PostConversationsEmailMessagesTooManyRequests {
	return &PostConversationsEmailMessagesTooManyRequests{}
}

/*
PostConversationsEmailMessagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsEmailMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages too many requests response has a 2xx status code
func (o *PostConversationsEmailMessagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages too many requests response has a 3xx status code
func (o *PostConversationsEmailMessagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages too many requests response has a 4xx status code
func (o *PostConversationsEmailMessagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations email messages too many requests response has a 5xx status code
func (o *PostConversationsEmailMessagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations email messages too many requests response a status code equal to that given
func (o *PostConversationsEmailMessagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsEmailMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsEmailMessagesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsEmailMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesInternalServerError creates a PostConversationsEmailMessagesInternalServerError with default headers values
func NewPostConversationsEmailMessagesInternalServerError() *PostConversationsEmailMessagesInternalServerError {
	return &PostConversationsEmailMessagesInternalServerError{}
}

/*
PostConversationsEmailMessagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsEmailMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages internal server error response has a 2xx status code
func (o *PostConversationsEmailMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages internal server error response has a 3xx status code
func (o *PostConversationsEmailMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages internal server error response has a 4xx status code
func (o *PostConversationsEmailMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations email messages internal server error response has a 5xx status code
func (o *PostConversationsEmailMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations email messages internal server error response a status code equal to that given
func (o *PostConversationsEmailMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsEmailMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsEmailMessagesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsEmailMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesServiceUnavailable creates a PostConversationsEmailMessagesServiceUnavailable with default headers values
func NewPostConversationsEmailMessagesServiceUnavailable() *PostConversationsEmailMessagesServiceUnavailable {
	return &PostConversationsEmailMessagesServiceUnavailable{}
}

/*
PostConversationsEmailMessagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsEmailMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages service unavailable response has a 2xx status code
func (o *PostConversationsEmailMessagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages service unavailable response has a 3xx status code
func (o *PostConversationsEmailMessagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages service unavailable response has a 4xx status code
func (o *PostConversationsEmailMessagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations email messages service unavailable response has a 5xx status code
func (o *PostConversationsEmailMessagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations email messages service unavailable response a status code equal to that given
func (o *PostConversationsEmailMessagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsEmailMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsEmailMessagesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsEmailMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesGatewayTimeout creates a PostConversationsEmailMessagesGatewayTimeout with default headers values
func NewPostConversationsEmailMessagesGatewayTimeout() *PostConversationsEmailMessagesGatewayTimeout {
	return &PostConversationsEmailMessagesGatewayTimeout{}
}

/*
PostConversationsEmailMessagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsEmailMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations email messages gateway timeout response has a 2xx status code
func (o *PostConversationsEmailMessagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations email messages gateway timeout response has a 3xx status code
func (o *PostConversationsEmailMessagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations email messages gateway timeout response has a 4xx status code
func (o *PostConversationsEmailMessagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations email messages gateway timeout response has a 5xx status code
func (o *PostConversationsEmailMessagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations email messages gateway timeout response a status code equal to that given
func (o *PostConversationsEmailMessagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsEmailMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsEmailMessagesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsEmailMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
