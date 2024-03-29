// Code generated by go-swagger; DO NOT EDIT.

package user_recordings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteUserrecordingReader is a Reader for the DeleteUserrecording structure.
type DeleteUserrecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserrecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteUserrecordingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserrecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserrecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserrecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserrecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteUserrecordingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteUserrecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteUserrecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserrecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserrecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteUserrecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteUserrecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserrecordingAccepted creates a DeleteUserrecordingAccepted with default headers values
func NewDeleteUserrecordingAccepted() *DeleteUserrecordingAccepted {
	return &DeleteUserrecordingAccepted{}
}

/*
DeleteUserrecordingAccepted describes a response with status code 202, with default header values.

Accepted - Processing Delete
*/
type DeleteUserrecordingAccepted struct {
}

// IsSuccess returns true when this delete userrecording accepted response has a 2xx status code
func (o *DeleteUserrecordingAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete userrecording accepted response has a 3xx status code
func (o *DeleteUserrecordingAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording accepted response has a 4xx status code
func (o *DeleteUserrecordingAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete userrecording accepted response has a 5xx status code
func (o *DeleteUserrecordingAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording accepted response a status code equal to that given
func (o *DeleteUserrecordingAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *DeleteUserrecordingAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingAccepted ", 202)
}

func (o *DeleteUserrecordingAccepted) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingAccepted ", 202)
}

func (o *DeleteUserrecordingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserrecordingBadRequest creates a DeleteUserrecordingBadRequest with default headers values
func NewDeleteUserrecordingBadRequest() *DeleteUserrecordingBadRequest {
	return &DeleteUserrecordingBadRequest{}
}

/*
DeleteUserrecordingBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteUserrecordingBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording bad request response has a 2xx status code
func (o *DeleteUserrecordingBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording bad request response has a 3xx status code
func (o *DeleteUserrecordingBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording bad request response has a 4xx status code
func (o *DeleteUserrecordingBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording bad request response has a 5xx status code
func (o *DeleteUserrecordingBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording bad request response a status code equal to that given
func (o *DeleteUserrecordingBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteUserrecordingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserrecordingBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserrecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingUnauthorized creates a DeleteUserrecordingUnauthorized with default headers values
func NewDeleteUserrecordingUnauthorized() *DeleteUserrecordingUnauthorized {
	return &DeleteUserrecordingUnauthorized{}
}

/*
DeleteUserrecordingUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteUserrecordingUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording unauthorized response has a 2xx status code
func (o *DeleteUserrecordingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording unauthorized response has a 3xx status code
func (o *DeleteUserrecordingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording unauthorized response has a 4xx status code
func (o *DeleteUserrecordingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording unauthorized response has a 5xx status code
func (o *DeleteUserrecordingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording unauthorized response a status code equal to that given
func (o *DeleteUserrecordingUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteUserrecordingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserrecordingUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserrecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingForbidden creates a DeleteUserrecordingForbidden with default headers values
func NewDeleteUserrecordingForbidden() *DeleteUserrecordingForbidden {
	return &DeleteUserrecordingForbidden{}
}

/*
DeleteUserrecordingForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteUserrecordingForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording forbidden response has a 2xx status code
func (o *DeleteUserrecordingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording forbidden response has a 3xx status code
func (o *DeleteUserrecordingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording forbidden response has a 4xx status code
func (o *DeleteUserrecordingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording forbidden response has a 5xx status code
func (o *DeleteUserrecordingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording forbidden response a status code equal to that given
func (o *DeleteUserrecordingForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteUserrecordingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserrecordingForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserrecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingNotFound creates a DeleteUserrecordingNotFound with default headers values
func NewDeleteUserrecordingNotFound() *DeleteUserrecordingNotFound {
	return &DeleteUserrecordingNotFound{}
}

/*
DeleteUserrecordingNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteUserrecordingNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording not found response has a 2xx status code
func (o *DeleteUserrecordingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording not found response has a 3xx status code
func (o *DeleteUserrecordingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording not found response has a 4xx status code
func (o *DeleteUserrecordingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording not found response has a 5xx status code
func (o *DeleteUserrecordingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording not found response a status code equal to that given
func (o *DeleteUserrecordingNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteUserrecordingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserrecordingNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserrecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingRequestTimeout creates a DeleteUserrecordingRequestTimeout with default headers values
func NewDeleteUserrecordingRequestTimeout() *DeleteUserrecordingRequestTimeout {
	return &DeleteUserrecordingRequestTimeout{}
}

/*
DeleteUserrecordingRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteUserrecordingRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording request timeout response has a 2xx status code
func (o *DeleteUserrecordingRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording request timeout response has a 3xx status code
func (o *DeleteUserrecordingRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording request timeout response has a 4xx status code
func (o *DeleteUserrecordingRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording request timeout response has a 5xx status code
func (o *DeleteUserrecordingRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording request timeout response a status code equal to that given
func (o *DeleteUserrecordingRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteUserrecordingRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteUserrecordingRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteUserrecordingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingRequestEntityTooLarge creates a DeleteUserrecordingRequestEntityTooLarge with default headers values
func NewDeleteUserrecordingRequestEntityTooLarge() *DeleteUserrecordingRequestEntityTooLarge {
	return &DeleteUserrecordingRequestEntityTooLarge{}
}

/*
DeleteUserrecordingRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteUserrecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording request entity too large response has a 2xx status code
func (o *DeleteUserrecordingRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording request entity too large response has a 3xx status code
func (o *DeleteUserrecordingRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording request entity too large response has a 4xx status code
func (o *DeleteUserrecordingRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording request entity too large response has a 5xx status code
func (o *DeleteUserrecordingRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording request entity too large response a status code equal to that given
func (o *DeleteUserrecordingRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteUserrecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserrecordingRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserrecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingUnsupportedMediaType creates a DeleteUserrecordingUnsupportedMediaType with default headers values
func NewDeleteUserrecordingUnsupportedMediaType() *DeleteUserrecordingUnsupportedMediaType {
	return &DeleteUserrecordingUnsupportedMediaType{}
}

/*
DeleteUserrecordingUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteUserrecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording unsupported media type response has a 2xx status code
func (o *DeleteUserrecordingUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording unsupported media type response has a 3xx status code
func (o *DeleteUserrecordingUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording unsupported media type response has a 4xx status code
func (o *DeleteUserrecordingUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording unsupported media type response has a 5xx status code
func (o *DeleteUserrecordingUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording unsupported media type response a status code equal to that given
func (o *DeleteUserrecordingUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteUserrecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserrecordingUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserrecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingTooManyRequests creates a DeleteUserrecordingTooManyRequests with default headers values
func NewDeleteUserrecordingTooManyRequests() *DeleteUserrecordingTooManyRequests {
	return &DeleteUserrecordingTooManyRequests{}
}

/*
DeleteUserrecordingTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteUserrecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording too many requests response has a 2xx status code
func (o *DeleteUserrecordingTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording too many requests response has a 3xx status code
func (o *DeleteUserrecordingTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording too many requests response has a 4xx status code
func (o *DeleteUserrecordingTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete userrecording too many requests response has a 5xx status code
func (o *DeleteUserrecordingTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete userrecording too many requests response a status code equal to that given
func (o *DeleteUserrecordingTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteUserrecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserrecordingTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserrecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingInternalServerError creates a DeleteUserrecordingInternalServerError with default headers values
func NewDeleteUserrecordingInternalServerError() *DeleteUserrecordingInternalServerError {
	return &DeleteUserrecordingInternalServerError{}
}

/*
DeleteUserrecordingInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteUserrecordingInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording internal server error response has a 2xx status code
func (o *DeleteUserrecordingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording internal server error response has a 3xx status code
func (o *DeleteUserrecordingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording internal server error response has a 4xx status code
func (o *DeleteUserrecordingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete userrecording internal server error response has a 5xx status code
func (o *DeleteUserrecordingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete userrecording internal server error response a status code equal to that given
func (o *DeleteUserrecordingInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteUserrecordingInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserrecordingInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserrecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingServiceUnavailable creates a DeleteUserrecordingServiceUnavailable with default headers values
func NewDeleteUserrecordingServiceUnavailable() *DeleteUserrecordingServiceUnavailable {
	return &DeleteUserrecordingServiceUnavailable{}
}

/*
DeleteUserrecordingServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteUserrecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording service unavailable response has a 2xx status code
func (o *DeleteUserrecordingServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording service unavailable response has a 3xx status code
func (o *DeleteUserrecordingServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording service unavailable response has a 4xx status code
func (o *DeleteUserrecordingServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete userrecording service unavailable response has a 5xx status code
func (o *DeleteUserrecordingServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete userrecording service unavailable response a status code equal to that given
func (o *DeleteUserrecordingServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteUserrecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserrecordingServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserrecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserrecordingGatewayTimeout creates a DeleteUserrecordingGatewayTimeout with default headers values
func NewDeleteUserrecordingGatewayTimeout() *DeleteUserrecordingGatewayTimeout {
	return &DeleteUserrecordingGatewayTimeout{}
}

/*
DeleteUserrecordingGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteUserrecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete userrecording gateway timeout response has a 2xx status code
func (o *DeleteUserrecordingGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete userrecording gateway timeout response has a 3xx status code
func (o *DeleteUserrecordingGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete userrecording gateway timeout response has a 4xx status code
func (o *DeleteUserrecordingGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete userrecording gateway timeout response has a 5xx status code
func (o *DeleteUserrecordingGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete userrecording gateway timeout response a status code equal to that given
func (o *DeleteUserrecordingGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteUserrecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserrecordingGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/userrecordings/{recordingId}][%d] deleteUserrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserrecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserrecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
