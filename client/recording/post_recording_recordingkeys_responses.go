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

// PostRecordingRecordingkeysReader is a Reader for the PostRecordingRecordingkeys structure.
type PostRecordingRecordingkeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordingRecordingkeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRecordingRecordingkeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRecordingRecordingkeysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRecordingRecordingkeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRecordingRecordingkeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordingRecordingkeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRecordingRecordingkeysRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRecordingRecordingkeysRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRecordingRecordingkeysUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRecordingRecordingkeysTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordingRecordingkeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRecordingRecordingkeysServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRecordingRecordingkeysGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRecordingRecordingkeysOK creates a PostRecordingRecordingkeysOK with default headers values
func NewPostRecordingRecordingkeysOK() *PostRecordingRecordingkeysOK {
	return &PostRecordingRecordingkeysOK{}
}

/*
PostRecordingRecordingkeysOK describes a response with status code 200, with default header values.

successful operation
*/
type PostRecordingRecordingkeysOK struct {
	Payload *models.EncryptionKey
}

// IsSuccess returns true when this post recording recordingkeys o k response has a 2xx status code
func (o *PostRecordingRecordingkeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post recording recordingkeys o k response has a 3xx status code
func (o *PostRecordingRecordingkeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys o k response has a 4xx status code
func (o *PostRecordingRecordingkeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording recordingkeys o k response has a 5xx status code
func (o *PostRecordingRecordingkeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys o k response a status code equal to that given
func (o *PostRecordingRecordingkeysOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostRecordingRecordingkeysOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysOK  %+v", 200, o.Payload)
}

func (o *PostRecordingRecordingkeysOK) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysOK  %+v", 200, o.Payload)
}

func (o *PostRecordingRecordingkeysOK) GetPayload() *models.EncryptionKey {
	return o.Payload
}

func (o *PostRecordingRecordingkeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EncryptionKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysBadRequest creates a PostRecordingRecordingkeysBadRequest with default headers values
func NewPostRecordingRecordingkeysBadRequest() *PostRecordingRecordingkeysBadRequest {
	return &PostRecordingRecordingkeysBadRequest{}
}

/*
PostRecordingRecordingkeysBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRecordingRecordingkeysBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys bad request response has a 2xx status code
func (o *PostRecordingRecordingkeysBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys bad request response has a 3xx status code
func (o *PostRecordingRecordingkeysBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys bad request response has a 4xx status code
func (o *PostRecordingRecordingkeysBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys bad request response has a 5xx status code
func (o *PostRecordingRecordingkeysBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys bad request response a status code equal to that given
func (o *PostRecordingRecordingkeysBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRecordingRecordingkeysBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingRecordingkeysBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingRecordingkeysBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysUnauthorized creates a PostRecordingRecordingkeysUnauthorized with default headers values
func NewPostRecordingRecordingkeysUnauthorized() *PostRecordingRecordingkeysUnauthorized {
	return &PostRecordingRecordingkeysUnauthorized{}
}

/*
PostRecordingRecordingkeysUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRecordingRecordingkeysUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys unauthorized response has a 2xx status code
func (o *PostRecordingRecordingkeysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys unauthorized response has a 3xx status code
func (o *PostRecordingRecordingkeysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys unauthorized response has a 4xx status code
func (o *PostRecordingRecordingkeysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys unauthorized response has a 5xx status code
func (o *PostRecordingRecordingkeysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys unauthorized response a status code equal to that given
func (o *PostRecordingRecordingkeysUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRecordingRecordingkeysUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingRecordingkeysUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingRecordingkeysUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysForbidden creates a PostRecordingRecordingkeysForbidden with default headers values
func NewPostRecordingRecordingkeysForbidden() *PostRecordingRecordingkeysForbidden {
	return &PostRecordingRecordingkeysForbidden{}
}

/*
PostRecordingRecordingkeysForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRecordingRecordingkeysForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys forbidden response has a 2xx status code
func (o *PostRecordingRecordingkeysForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys forbidden response has a 3xx status code
func (o *PostRecordingRecordingkeysForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys forbidden response has a 4xx status code
func (o *PostRecordingRecordingkeysForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys forbidden response has a 5xx status code
func (o *PostRecordingRecordingkeysForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys forbidden response a status code equal to that given
func (o *PostRecordingRecordingkeysForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRecordingRecordingkeysForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingRecordingkeysForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingRecordingkeysForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysNotFound creates a PostRecordingRecordingkeysNotFound with default headers values
func NewPostRecordingRecordingkeysNotFound() *PostRecordingRecordingkeysNotFound {
	return &PostRecordingRecordingkeysNotFound{}
}

/*
PostRecordingRecordingkeysNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRecordingRecordingkeysNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys not found response has a 2xx status code
func (o *PostRecordingRecordingkeysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys not found response has a 3xx status code
func (o *PostRecordingRecordingkeysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys not found response has a 4xx status code
func (o *PostRecordingRecordingkeysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys not found response has a 5xx status code
func (o *PostRecordingRecordingkeysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys not found response a status code equal to that given
func (o *PostRecordingRecordingkeysNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRecordingRecordingkeysNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingRecordingkeysNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingRecordingkeysNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysRequestTimeout creates a PostRecordingRecordingkeysRequestTimeout with default headers values
func NewPostRecordingRecordingkeysRequestTimeout() *PostRecordingRecordingkeysRequestTimeout {
	return &PostRecordingRecordingkeysRequestTimeout{}
}

/*
PostRecordingRecordingkeysRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRecordingRecordingkeysRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys request timeout response has a 2xx status code
func (o *PostRecordingRecordingkeysRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys request timeout response has a 3xx status code
func (o *PostRecordingRecordingkeysRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys request timeout response has a 4xx status code
func (o *PostRecordingRecordingkeysRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys request timeout response has a 5xx status code
func (o *PostRecordingRecordingkeysRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys request timeout response a status code equal to that given
func (o *PostRecordingRecordingkeysRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRecordingRecordingkeysRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingRecordingkeysRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingRecordingkeysRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysRequestEntityTooLarge creates a PostRecordingRecordingkeysRequestEntityTooLarge with default headers values
func NewPostRecordingRecordingkeysRequestEntityTooLarge() *PostRecordingRecordingkeysRequestEntityTooLarge {
	return &PostRecordingRecordingkeysRequestEntityTooLarge{}
}

/*
PostRecordingRecordingkeysRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRecordingRecordingkeysRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys request entity too large response has a 2xx status code
func (o *PostRecordingRecordingkeysRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys request entity too large response has a 3xx status code
func (o *PostRecordingRecordingkeysRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys request entity too large response has a 4xx status code
func (o *PostRecordingRecordingkeysRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys request entity too large response has a 5xx status code
func (o *PostRecordingRecordingkeysRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys request entity too large response a status code equal to that given
func (o *PostRecordingRecordingkeysRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRecordingRecordingkeysRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingRecordingkeysRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingRecordingkeysRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysUnsupportedMediaType creates a PostRecordingRecordingkeysUnsupportedMediaType with default headers values
func NewPostRecordingRecordingkeysUnsupportedMediaType() *PostRecordingRecordingkeysUnsupportedMediaType {
	return &PostRecordingRecordingkeysUnsupportedMediaType{}
}

/*
PostRecordingRecordingkeysUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRecordingRecordingkeysUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys unsupported media type response has a 2xx status code
func (o *PostRecordingRecordingkeysUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys unsupported media type response has a 3xx status code
func (o *PostRecordingRecordingkeysUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys unsupported media type response has a 4xx status code
func (o *PostRecordingRecordingkeysUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys unsupported media type response has a 5xx status code
func (o *PostRecordingRecordingkeysUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys unsupported media type response a status code equal to that given
func (o *PostRecordingRecordingkeysUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRecordingRecordingkeysUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingRecordingkeysUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingRecordingkeysUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysTooManyRequests creates a PostRecordingRecordingkeysTooManyRequests with default headers values
func NewPostRecordingRecordingkeysTooManyRequests() *PostRecordingRecordingkeysTooManyRequests {
	return &PostRecordingRecordingkeysTooManyRequests{}
}

/*
PostRecordingRecordingkeysTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRecordingRecordingkeysTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys too many requests response has a 2xx status code
func (o *PostRecordingRecordingkeysTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys too many requests response has a 3xx status code
func (o *PostRecordingRecordingkeysTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys too many requests response has a 4xx status code
func (o *PostRecordingRecordingkeysTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording recordingkeys too many requests response has a 5xx status code
func (o *PostRecordingRecordingkeysTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording recordingkeys too many requests response a status code equal to that given
func (o *PostRecordingRecordingkeysTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRecordingRecordingkeysTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingRecordingkeysTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingRecordingkeysTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysInternalServerError creates a PostRecordingRecordingkeysInternalServerError with default headers values
func NewPostRecordingRecordingkeysInternalServerError() *PostRecordingRecordingkeysInternalServerError {
	return &PostRecordingRecordingkeysInternalServerError{}
}

/*
PostRecordingRecordingkeysInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRecordingRecordingkeysInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys internal server error response has a 2xx status code
func (o *PostRecordingRecordingkeysInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys internal server error response has a 3xx status code
func (o *PostRecordingRecordingkeysInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys internal server error response has a 4xx status code
func (o *PostRecordingRecordingkeysInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording recordingkeys internal server error response has a 5xx status code
func (o *PostRecordingRecordingkeysInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording recordingkeys internal server error response a status code equal to that given
func (o *PostRecordingRecordingkeysInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRecordingRecordingkeysInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingRecordingkeysInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingRecordingkeysInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysServiceUnavailable creates a PostRecordingRecordingkeysServiceUnavailable with default headers values
func NewPostRecordingRecordingkeysServiceUnavailable() *PostRecordingRecordingkeysServiceUnavailable {
	return &PostRecordingRecordingkeysServiceUnavailable{}
}

/*
PostRecordingRecordingkeysServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRecordingRecordingkeysServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys service unavailable response has a 2xx status code
func (o *PostRecordingRecordingkeysServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys service unavailable response has a 3xx status code
func (o *PostRecordingRecordingkeysServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys service unavailable response has a 4xx status code
func (o *PostRecordingRecordingkeysServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording recordingkeys service unavailable response has a 5xx status code
func (o *PostRecordingRecordingkeysServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording recordingkeys service unavailable response a status code equal to that given
func (o *PostRecordingRecordingkeysServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRecordingRecordingkeysServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingRecordingkeysServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingRecordingkeysServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingRecordingkeysGatewayTimeout creates a PostRecordingRecordingkeysGatewayTimeout with default headers values
func NewPostRecordingRecordingkeysGatewayTimeout() *PostRecordingRecordingkeysGatewayTimeout {
	return &PostRecordingRecordingkeysGatewayTimeout{}
}

/*
PostRecordingRecordingkeysGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRecordingRecordingkeysGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording recordingkeys gateway timeout response has a 2xx status code
func (o *PostRecordingRecordingkeysGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording recordingkeys gateway timeout response has a 3xx status code
func (o *PostRecordingRecordingkeysGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording recordingkeys gateway timeout response has a 4xx status code
func (o *PostRecordingRecordingkeysGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording recordingkeys gateway timeout response has a 5xx status code
func (o *PostRecordingRecordingkeysGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording recordingkeys gateway timeout response a status code equal to that given
func (o *PostRecordingRecordingkeysGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRecordingRecordingkeysGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingRecordingkeysGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/recordingkeys][%d] postRecordingRecordingkeysGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingRecordingkeysGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingRecordingkeysGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
