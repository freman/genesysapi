// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostUploadsRecordingsReader is a Reader for the PostUploadsRecordings structure.
type PostUploadsRecordingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUploadsRecordingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUploadsRecordingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUploadsRecordingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUploadsRecordingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUploadsRecordingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUploadsRecordingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUploadsRecordingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUploadsRecordingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUploadsRecordingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUploadsRecordingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUploadsRecordingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUploadsRecordingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUploadsRecordingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUploadsRecordingsOK creates a PostUploadsRecordingsOK with default headers values
func NewPostUploadsRecordingsOK() *PostUploadsRecordingsOK {
	return &PostUploadsRecordingsOK{}
}

/*
PostUploadsRecordingsOK describes a response with status code 200, with default header values.

Presigned url successfully created.
*/
type PostUploadsRecordingsOK struct {
	Payload *models.UploadURLResponse
}

// IsSuccess returns true when this post uploads recordings o k response has a 2xx status code
func (o *PostUploadsRecordingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post uploads recordings o k response has a 3xx status code
func (o *PostUploadsRecordingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings o k response has a 4xx status code
func (o *PostUploadsRecordingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post uploads recordings o k response has a 5xx status code
func (o *PostUploadsRecordingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings o k response a status code equal to that given
func (o *PostUploadsRecordingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostUploadsRecordingsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsOK  %+v", 200, o.Payload)
}

func (o *PostUploadsRecordingsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsOK  %+v", 200, o.Payload)
}

func (o *PostUploadsRecordingsOK) GetPayload() *models.UploadURLResponse {
	return o.Payload
}

func (o *PostUploadsRecordingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsBadRequest creates a PostUploadsRecordingsBadRequest with default headers values
func NewPostUploadsRecordingsBadRequest() *PostUploadsRecordingsBadRequest {
	return &PostUploadsRecordingsBadRequest{}
}

/*
PostUploadsRecordingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUploadsRecordingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings bad request response has a 2xx status code
func (o *PostUploadsRecordingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings bad request response has a 3xx status code
func (o *PostUploadsRecordingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings bad request response has a 4xx status code
func (o *PostUploadsRecordingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings bad request response has a 5xx status code
func (o *PostUploadsRecordingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings bad request response a status code equal to that given
func (o *PostUploadsRecordingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUploadsRecordingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUploadsRecordingsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostUploadsRecordingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsUnauthorized creates a PostUploadsRecordingsUnauthorized with default headers values
func NewPostUploadsRecordingsUnauthorized() *PostUploadsRecordingsUnauthorized {
	return &PostUploadsRecordingsUnauthorized{}
}

/*
PostUploadsRecordingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUploadsRecordingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings unauthorized response has a 2xx status code
func (o *PostUploadsRecordingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings unauthorized response has a 3xx status code
func (o *PostUploadsRecordingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings unauthorized response has a 4xx status code
func (o *PostUploadsRecordingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings unauthorized response has a 5xx status code
func (o *PostUploadsRecordingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings unauthorized response a status code equal to that given
func (o *PostUploadsRecordingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUploadsRecordingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUploadsRecordingsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUploadsRecordingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsForbidden creates a PostUploadsRecordingsForbidden with default headers values
func NewPostUploadsRecordingsForbidden() *PostUploadsRecordingsForbidden {
	return &PostUploadsRecordingsForbidden{}
}

/*
PostUploadsRecordingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostUploadsRecordingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings forbidden response has a 2xx status code
func (o *PostUploadsRecordingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings forbidden response has a 3xx status code
func (o *PostUploadsRecordingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings forbidden response has a 4xx status code
func (o *PostUploadsRecordingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings forbidden response has a 5xx status code
func (o *PostUploadsRecordingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings forbidden response a status code equal to that given
func (o *PostUploadsRecordingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUploadsRecordingsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsForbidden  %+v", 403, o.Payload)
}

func (o *PostUploadsRecordingsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsForbidden  %+v", 403, o.Payload)
}

func (o *PostUploadsRecordingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsNotFound creates a PostUploadsRecordingsNotFound with default headers values
func NewPostUploadsRecordingsNotFound() *PostUploadsRecordingsNotFound {
	return &PostUploadsRecordingsNotFound{}
}

/*
PostUploadsRecordingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostUploadsRecordingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings not found response has a 2xx status code
func (o *PostUploadsRecordingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings not found response has a 3xx status code
func (o *PostUploadsRecordingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings not found response has a 4xx status code
func (o *PostUploadsRecordingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings not found response has a 5xx status code
func (o *PostUploadsRecordingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings not found response a status code equal to that given
func (o *PostUploadsRecordingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUploadsRecordingsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsNotFound  %+v", 404, o.Payload)
}

func (o *PostUploadsRecordingsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsNotFound  %+v", 404, o.Payload)
}

func (o *PostUploadsRecordingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsRequestTimeout creates a PostUploadsRecordingsRequestTimeout with default headers values
func NewPostUploadsRecordingsRequestTimeout() *PostUploadsRecordingsRequestTimeout {
	return &PostUploadsRecordingsRequestTimeout{}
}

/*
PostUploadsRecordingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUploadsRecordingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings request timeout response has a 2xx status code
func (o *PostUploadsRecordingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings request timeout response has a 3xx status code
func (o *PostUploadsRecordingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings request timeout response has a 4xx status code
func (o *PostUploadsRecordingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings request timeout response has a 5xx status code
func (o *PostUploadsRecordingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings request timeout response a status code equal to that given
func (o *PostUploadsRecordingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostUploadsRecordingsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUploadsRecordingsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUploadsRecordingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsRequestEntityTooLarge creates a PostUploadsRecordingsRequestEntityTooLarge with default headers values
func NewPostUploadsRecordingsRequestEntityTooLarge() *PostUploadsRecordingsRequestEntityTooLarge {
	return &PostUploadsRecordingsRequestEntityTooLarge{}
}

/*
PostUploadsRecordingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostUploadsRecordingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings request entity too large response has a 2xx status code
func (o *PostUploadsRecordingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings request entity too large response has a 3xx status code
func (o *PostUploadsRecordingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings request entity too large response has a 4xx status code
func (o *PostUploadsRecordingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings request entity too large response has a 5xx status code
func (o *PostUploadsRecordingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings request entity too large response a status code equal to that given
func (o *PostUploadsRecordingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostUploadsRecordingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUploadsRecordingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUploadsRecordingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsUnsupportedMediaType creates a PostUploadsRecordingsUnsupportedMediaType with default headers values
func NewPostUploadsRecordingsUnsupportedMediaType() *PostUploadsRecordingsUnsupportedMediaType {
	return &PostUploadsRecordingsUnsupportedMediaType{}
}

/*
PostUploadsRecordingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUploadsRecordingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings unsupported media type response has a 2xx status code
func (o *PostUploadsRecordingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings unsupported media type response has a 3xx status code
func (o *PostUploadsRecordingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings unsupported media type response has a 4xx status code
func (o *PostUploadsRecordingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings unsupported media type response has a 5xx status code
func (o *PostUploadsRecordingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings unsupported media type response a status code equal to that given
func (o *PostUploadsRecordingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostUploadsRecordingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUploadsRecordingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUploadsRecordingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsTooManyRequests creates a PostUploadsRecordingsTooManyRequests with default headers values
func NewPostUploadsRecordingsTooManyRequests() *PostUploadsRecordingsTooManyRequests {
	return &PostUploadsRecordingsTooManyRequests{}
}

/*
PostUploadsRecordingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUploadsRecordingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings too many requests response has a 2xx status code
func (o *PostUploadsRecordingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings too many requests response has a 3xx status code
func (o *PostUploadsRecordingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings too many requests response has a 4xx status code
func (o *PostUploadsRecordingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post uploads recordings too many requests response has a 5xx status code
func (o *PostUploadsRecordingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post uploads recordings too many requests response a status code equal to that given
func (o *PostUploadsRecordingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostUploadsRecordingsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUploadsRecordingsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUploadsRecordingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsInternalServerError creates a PostUploadsRecordingsInternalServerError with default headers values
func NewPostUploadsRecordingsInternalServerError() *PostUploadsRecordingsInternalServerError {
	return &PostUploadsRecordingsInternalServerError{}
}

/*
PostUploadsRecordingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUploadsRecordingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings internal server error response has a 2xx status code
func (o *PostUploadsRecordingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings internal server error response has a 3xx status code
func (o *PostUploadsRecordingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings internal server error response has a 4xx status code
func (o *PostUploadsRecordingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post uploads recordings internal server error response has a 5xx status code
func (o *PostUploadsRecordingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post uploads recordings internal server error response a status code equal to that given
func (o *PostUploadsRecordingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUploadsRecordingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUploadsRecordingsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUploadsRecordingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsServiceUnavailable creates a PostUploadsRecordingsServiceUnavailable with default headers values
func NewPostUploadsRecordingsServiceUnavailable() *PostUploadsRecordingsServiceUnavailable {
	return &PostUploadsRecordingsServiceUnavailable{}
}

/*
PostUploadsRecordingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUploadsRecordingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings service unavailable response has a 2xx status code
func (o *PostUploadsRecordingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings service unavailable response has a 3xx status code
func (o *PostUploadsRecordingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings service unavailable response has a 4xx status code
func (o *PostUploadsRecordingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post uploads recordings service unavailable response has a 5xx status code
func (o *PostUploadsRecordingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post uploads recordings service unavailable response a status code equal to that given
func (o *PostUploadsRecordingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUploadsRecordingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUploadsRecordingsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUploadsRecordingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUploadsRecordingsGatewayTimeout creates a PostUploadsRecordingsGatewayTimeout with default headers values
func NewPostUploadsRecordingsGatewayTimeout() *PostUploadsRecordingsGatewayTimeout {
	return &PostUploadsRecordingsGatewayTimeout{}
}

/*
PostUploadsRecordingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostUploadsRecordingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post uploads recordings gateway timeout response has a 2xx status code
func (o *PostUploadsRecordingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post uploads recordings gateway timeout response has a 3xx status code
func (o *PostUploadsRecordingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post uploads recordings gateway timeout response has a 4xx status code
func (o *PostUploadsRecordingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post uploads recordings gateway timeout response has a 5xx status code
func (o *PostUploadsRecordingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post uploads recordings gateway timeout response a status code equal to that given
func (o *PostUploadsRecordingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUploadsRecordingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUploadsRecordingsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/uploads/recordings][%d] postUploadsRecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUploadsRecordingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUploadsRecordingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
