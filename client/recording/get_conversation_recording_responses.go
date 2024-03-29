// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetConversationRecordingReader is a Reader for the GetConversationRecording structure.
type GetConversationRecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationRecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewGetConversationRecordingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationRecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationRecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationRecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationRecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationRecordingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationRecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationRecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationRecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationRecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationRecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationRecordingAccepted creates a GetConversationRecordingAccepted with default headers values
func NewGetConversationRecordingAccepted() *GetConversationRecordingAccepted {
	return &GetConversationRecordingAccepted{}
}

/*
GetConversationRecordingAccepted describes a response with status code 202, with default header values.

Success - recording is transcoding
*/
type GetConversationRecordingAccepted struct {
	Payload *models.Recording
}

// IsSuccess returns true when this get conversation recording accepted response has a 2xx status code
func (o *GetConversationRecordingAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversation recording accepted response has a 3xx status code
func (o *GetConversationRecordingAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording accepted response has a 4xx status code
func (o *GetConversationRecordingAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recording accepted response has a 5xx status code
func (o *GetConversationRecordingAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording accepted response a status code equal to that given
func (o *GetConversationRecordingAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GetConversationRecordingAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingAccepted  %+v", 202, o.Payload)
}

func (o *GetConversationRecordingAccepted) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingAccepted  %+v", 202, o.Payload)
}

func (o *GetConversationRecordingAccepted) GetPayload() *models.Recording {
	return o.Payload
}

func (o *GetConversationRecordingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recording)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingBadRequest creates a GetConversationRecordingBadRequest with default headers values
func NewGetConversationRecordingBadRequest() *GetConversationRecordingBadRequest {
	return &GetConversationRecordingBadRequest{}
}

/*
GetConversationRecordingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording bad request response has a 2xx status code
func (o *GetConversationRecordingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording bad request response has a 3xx status code
func (o *GetConversationRecordingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording bad request response has a 4xx status code
func (o *GetConversationRecordingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording bad request response has a 5xx status code
func (o *GetConversationRecordingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording bad request response a status code equal to that given
func (o *GetConversationRecordingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationRecordingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingUnauthorized creates a GetConversationRecordingUnauthorized with default headers values
func NewGetConversationRecordingUnauthorized() *GetConversationRecordingUnauthorized {
	return &GetConversationRecordingUnauthorized{}
}

/*
GetConversationRecordingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording unauthorized response has a 2xx status code
func (o *GetConversationRecordingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording unauthorized response has a 3xx status code
func (o *GetConversationRecordingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording unauthorized response has a 4xx status code
func (o *GetConversationRecordingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording unauthorized response has a 5xx status code
func (o *GetConversationRecordingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording unauthorized response a status code equal to that given
func (o *GetConversationRecordingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationRecordingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingForbidden creates a GetConversationRecordingForbidden with default headers values
func NewGetConversationRecordingForbidden() *GetConversationRecordingForbidden {
	return &GetConversationRecordingForbidden{}
}

/*
GetConversationRecordingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording forbidden response has a 2xx status code
func (o *GetConversationRecordingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording forbidden response has a 3xx status code
func (o *GetConversationRecordingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording forbidden response has a 4xx status code
func (o *GetConversationRecordingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording forbidden response has a 5xx status code
func (o *GetConversationRecordingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording forbidden response a status code equal to that given
func (o *GetConversationRecordingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationRecordingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingNotFound creates a GetConversationRecordingNotFound with default headers values
func NewGetConversationRecordingNotFound() *GetConversationRecordingNotFound {
	return &GetConversationRecordingNotFound{}
}

/*
GetConversationRecordingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationRecordingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording not found response has a 2xx status code
func (o *GetConversationRecordingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording not found response has a 3xx status code
func (o *GetConversationRecordingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording not found response has a 4xx status code
func (o *GetConversationRecordingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording not found response has a 5xx status code
func (o *GetConversationRecordingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording not found response a status code equal to that given
func (o *GetConversationRecordingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationRecordingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingRequestTimeout creates a GetConversationRecordingRequestTimeout with default headers values
func NewGetConversationRecordingRequestTimeout() *GetConversationRecordingRequestTimeout {
	return &GetConversationRecordingRequestTimeout{}
}

/*
GetConversationRecordingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationRecordingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording request timeout response has a 2xx status code
func (o *GetConversationRecordingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording request timeout response has a 3xx status code
func (o *GetConversationRecordingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording request timeout response has a 4xx status code
func (o *GetConversationRecordingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording request timeout response has a 5xx status code
func (o *GetConversationRecordingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording request timeout response a status code equal to that given
func (o *GetConversationRecordingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationRecordingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingRequestEntityTooLarge creates a GetConversationRecordingRequestEntityTooLarge with default headers values
func NewGetConversationRecordingRequestEntityTooLarge() *GetConversationRecordingRequestEntityTooLarge {
	return &GetConversationRecordingRequestEntityTooLarge{}
}

/*
GetConversationRecordingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationRecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording request entity too large response has a 2xx status code
func (o *GetConversationRecordingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording request entity too large response has a 3xx status code
func (o *GetConversationRecordingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording request entity too large response has a 4xx status code
func (o *GetConversationRecordingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording request entity too large response has a 5xx status code
func (o *GetConversationRecordingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording request entity too large response a status code equal to that given
func (o *GetConversationRecordingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationRecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingUnsupportedMediaType creates a GetConversationRecordingUnsupportedMediaType with default headers values
func NewGetConversationRecordingUnsupportedMediaType() *GetConversationRecordingUnsupportedMediaType {
	return &GetConversationRecordingUnsupportedMediaType{}
}

/*
GetConversationRecordingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording unsupported media type response has a 2xx status code
func (o *GetConversationRecordingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording unsupported media type response has a 3xx status code
func (o *GetConversationRecordingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording unsupported media type response has a 4xx status code
func (o *GetConversationRecordingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording unsupported media type response has a 5xx status code
func (o *GetConversationRecordingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording unsupported media type response a status code equal to that given
func (o *GetConversationRecordingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationRecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingTooManyRequests creates a GetConversationRecordingTooManyRequests with default headers values
func NewGetConversationRecordingTooManyRequests() *GetConversationRecordingTooManyRequests {
	return &GetConversationRecordingTooManyRequests{}
}

/*
GetConversationRecordingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationRecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording too many requests response has a 2xx status code
func (o *GetConversationRecordingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording too many requests response has a 3xx status code
func (o *GetConversationRecordingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording too many requests response has a 4xx status code
func (o *GetConversationRecordingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recording too many requests response has a 5xx status code
func (o *GetConversationRecordingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recording too many requests response a status code equal to that given
func (o *GetConversationRecordingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationRecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingInternalServerError creates a GetConversationRecordingInternalServerError with default headers values
func NewGetConversationRecordingInternalServerError() *GetConversationRecordingInternalServerError {
	return &GetConversationRecordingInternalServerError{}
}

/*
GetConversationRecordingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording internal server error response has a 2xx status code
func (o *GetConversationRecordingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording internal server error response has a 3xx status code
func (o *GetConversationRecordingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording internal server error response has a 4xx status code
func (o *GetConversationRecordingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recording internal server error response has a 5xx status code
func (o *GetConversationRecordingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recording internal server error response a status code equal to that given
func (o *GetConversationRecordingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationRecordingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingServiceUnavailable creates a GetConversationRecordingServiceUnavailable with default headers values
func NewGetConversationRecordingServiceUnavailable() *GetConversationRecordingServiceUnavailable {
	return &GetConversationRecordingServiceUnavailable{}
}

/*
GetConversationRecordingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording service unavailable response has a 2xx status code
func (o *GetConversationRecordingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording service unavailable response has a 3xx status code
func (o *GetConversationRecordingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording service unavailable response has a 4xx status code
func (o *GetConversationRecordingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recording service unavailable response has a 5xx status code
func (o *GetConversationRecordingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recording service unavailable response a status code equal to that given
func (o *GetConversationRecordingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationRecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingGatewayTimeout creates a GetConversationRecordingGatewayTimeout with default headers values
func NewGetConversationRecordingGatewayTimeout() *GetConversationRecordingGatewayTimeout {
	return &GetConversationRecordingGatewayTimeout{}
}

/*
GetConversationRecordingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationRecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recording gateway timeout response has a 2xx status code
func (o *GetConversationRecordingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recording gateway timeout response has a 3xx status code
func (o *GetConversationRecordingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recording gateway timeout response has a 4xx status code
func (o *GetConversationRecordingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recording gateway timeout response has a 5xx status code
func (o *GetConversationRecordingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recording gateway timeout response a status code equal to that given
func (o *GetConversationRecordingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationRecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings/{recordingId}][%d] getConversationRecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
