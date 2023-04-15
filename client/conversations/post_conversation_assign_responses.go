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

// PostConversationAssignReader is a Reader for the PostConversationAssign structure.
type PostConversationAssignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationAssignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostConversationAssignAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationAssignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationAssignUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationAssignForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationAssignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationAssignRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationAssignRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationAssignUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationAssignTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationAssignInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationAssignServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationAssignGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationAssignAccepted creates a PostConversationAssignAccepted with default headers values
func NewPostConversationAssignAccepted() *PostConversationAssignAccepted {
	return &PostConversationAssignAccepted{}
}

/*
PostConversationAssignAccepted describes a response with status code 202, with default header values.

The manual assignment request was accepted
*/
type PostConversationAssignAccepted struct {
	Payload string
}

// IsSuccess returns true when this post conversation assign accepted response has a 2xx status code
func (o *PostConversationAssignAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversation assign accepted response has a 3xx status code
func (o *PostConversationAssignAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign accepted response has a 4xx status code
func (o *PostConversationAssignAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversation assign accepted response has a 5xx status code
func (o *PostConversationAssignAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign accepted response a status code equal to that given
func (o *PostConversationAssignAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostConversationAssignAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignAccepted  %+v", 202, o.Payload)
}

func (o *PostConversationAssignAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignAccepted  %+v", 202, o.Payload)
}

func (o *PostConversationAssignAccepted) GetPayload() string {
	return o.Payload
}

func (o *PostConversationAssignAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignBadRequest creates a PostConversationAssignBadRequest with default headers values
func NewPostConversationAssignBadRequest() *PostConversationAssignBadRequest {
	return &PostConversationAssignBadRequest{}
}

/*
PostConversationAssignBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationAssignBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign bad request response has a 2xx status code
func (o *PostConversationAssignBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign bad request response has a 3xx status code
func (o *PostConversationAssignBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign bad request response has a 4xx status code
func (o *PostConversationAssignBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign bad request response has a 5xx status code
func (o *PostConversationAssignBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign bad request response a status code equal to that given
func (o *PostConversationAssignBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationAssignBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationAssignBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationAssignBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignUnauthorized creates a PostConversationAssignUnauthorized with default headers values
func NewPostConversationAssignUnauthorized() *PostConversationAssignUnauthorized {
	return &PostConversationAssignUnauthorized{}
}

/*
PostConversationAssignUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationAssignUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign unauthorized response has a 2xx status code
func (o *PostConversationAssignUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign unauthorized response has a 3xx status code
func (o *PostConversationAssignUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign unauthorized response has a 4xx status code
func (o *PostConversationAssignUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign unauthorized response has a 5xx status code
func (o *PostConversationAssignUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign unauthorized response a status code equal to that given
func (o *PostConversationAssignUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationAssignUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationAssignUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationAssignUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignForbidden creates a PostConversationAssignForbidden with default headers values
func NewPostConversationAssignForbidden() *PostConversationAssignForbidden {
	return &PostConversationAssignForbidden{}
}

/*
PostConversationAssignForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationAssignForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign forbidden response has a 2xx status code
func (o *PostConversationAssignForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign forbidden response has a 3xx status code
func (o *PostConversationAssignForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign forbidden response has a 4xx status code
func (o *PostConversationAssignForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign forbidden response has a 5xx status code
func (o *PostConversationAssignForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign forbidden response a status code equal to that given
func (o *PostConversationAssignForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationAssignForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationAssignForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationAssignForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignNotFound creates a PostConversationAssignNotFound with default headers values
func NewPostConversationAssignNotFound() *PostConversationAssignNotFound {
	return &PostConversationAssignNotFound{}
}

/*
PostConversationAssignNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationAssignNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign not found response has a 2xx status code
func (o *PostConversationAssignNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign not found response has a 3xx status code
func (o *PostConversationAssignNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign not found response has a 4xx status code
func (o *PostConversationAssignNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign not found response has a 5xx status code
func (o *PostConversationAssignNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign not found response a status code equal to that given
func (o *PostConversationAssignNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationAssignNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationAssignNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationAssignNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignRequestTimeout creates a PostConversationAssignRequestTimeout with default headers values
func NewPostConversationAssignRequestTimeout() *PostConversationAssignRequestTimeout {
	return &PostConversationAssignRequestTimeout{}
}

/*
PostConversationAssignRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationAssignRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign request timeout response has a 2xx status code
func (o *PostConversationAssignRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign request timeout response has a 3xx status code
func (o *PostConversationAssignRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign request timeout response has a 4xx status code
func (o *PostConversationAssignRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign request timeout response has a 5xx status code
func (o *PostConversationAssignRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign request timeout response a status code equal to that given
func (o *PostConversationAssignRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationAssignRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationAssignRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationAssignRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignRequestEntityTooLarge creates a PostConversationAssignRequestEntityTooLarge with default headers values
func NewPostConversationAssignRequestEntityTooLarge() *PostConversationAssignRequestEntityTooLarge {
	return &PostConversationAssignRequestEntityTooLarge{}
}

/*
PostConversationAssignRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationAssignRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign request entity too large response has a 2xx status code
func (o *PostConversationAssignRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign request entity too large response has a 3xx status code
func (o *PostConversationAssignRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign request entity too large response has a 4xx status code
func (o *PostConversationAssignRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign request entity too large response has a 5xx status code
func (o *PostConversationAssignRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign request entity too large response a status code equal to that given
func (o *PostConversationAssignRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationAssignRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationAssignRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationAssignRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignUnsupportedMediaType creates a PostConversationAssignUnsupportedMediaType with default headers values
func NewPostConversationAssignUnsupportedMediaType() *PostConversationAssignUnsupportedMediaType {
	return &PostConversationAssignUnsupportedMediaType{}
}

/*
PostConversationAssignUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationAssignUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign unsupported media type response has a 2xx status code
func (o *PostConversationAssignUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign unsupported media type response has a 3xx status code
func (o *PostConversationAssignUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign unsupported media type response has a 4xx status code
func (o *PostConversationAssignUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign unsupported media type response has a 5xx status code
func (o *PostConversationAssignUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign unsupported media type response a status code equal to that given
func (o *PostConversationAssignUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationAssignUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationAssignUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationAssignUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignTooManyRequests creates a PostConversationAssignTooManyRequests with default headers values
func NewPostConversationAssignTooManyRequests() *PostConversationAssignTooManyRequests {
	return &PostConversationAssignTooManyRequests{}
}

/*
PostConversationAssignTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationAssignTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign too many requests response has a 2xx status code
func (o *PostConversationAssignTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign too many requests response has a 3xx status code
func (o *PostConversationAssignTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign too many requests response has a 4xx status code
func (o *PostConversationAssignTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversation assign too many requests response has a 5xx status code
func (o *PostConversationAssignTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversation assign too many requests response a status code equal to that given
func (o *PostConversationAssignTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationAssignTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationAssignTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationAssignTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignInternalServerError creates a PostConversationAssignInternalServerError with default headers values
func NewPostConversationAssignInternalServerError() *PostConversationAssignInternalServerError {
	return &PostConversationAssignInternalServerError{}
}

/*
PostConversationAssignInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationAssignInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign internal server error response has a 2xx status code
func (o *PostConversationAssignInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign internal server error response has a 3xx status code
func (o *PostConversationAssignInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign internal server error response has a 4xx status code
func (o *PostConversationAssignInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversation assign internal server error response has a 5xx status code
func (o *PostConversationAssignInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversation assign internal server error response a status code equal to that given
func (o *PostConversationAssignInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationAssignInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationAssignInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationAssignInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignServiceUnavailable creates a PostConversationAssignServiceUnavailable with default headers values
func NewPostConversationAssignServiceUnavailable() *PostConversationAssignServiceUnavailable {
	return &PostConversationAssignServiceUnavailable{}
}

/*
PostConversationAssignServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationAssignServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign service unavailable response has a 2xx status code
func (o *PostConversationAssignServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign service unavailable response has a 3xx status code
func (o *PostConversationAssignServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign service unavailable response has a 4xx status code
func (o *PostConversationAssignServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversation assign service unavailable response has a 5xx status code
func (o *PostConversationAssignServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversation assign service unavailable response a status code equal to that given
func (o *PostConversationAssignServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationAssignServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationAssignServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationAssignServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationAssignGatewayTimeout creates a PostConversationAssignGatewayTimeout with default headers values
func NewPostConversationAssignGatewayTimeout() *PostConversationAssignGatewayTimeout {
	return &PostConversationAssignGatewayTimeout{}
}

/*
PostConversationAssignGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationAssignGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversation assign gateway timeout response has a 2xx status code
func (o *PostConversationAssignGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversation assign gateway timeout response has a 3xx status code
func (o *PostConversationAssignGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversation assign gateway timeout response has a 4xx status code
func (o *PostConversationAssignGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversation assign gateway timeout response has a 5xx status code
func (o *PostConversationAssignGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversation assign gateway timeout response a status code equal to that given
func (o *PostConversationAssignGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationAssignGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationAssignGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/{conversationId}/assign][%d] postConversationAssignGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationAssignGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationAssignGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
