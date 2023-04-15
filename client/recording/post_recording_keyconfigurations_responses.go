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

// PostRecordingKeyconfigurationsReader is a Reader for the PostRecordingKeyconfigurations structure.
type PostRecordingKeyconfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordingKeyconfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRecordingKeyconfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRecordingKeyconfigurationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRecordingKeyconfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRecordingKeyconfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordingKeyconfigurationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRecordingKeyconfigurationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRecordingKeyconfigurationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRecordingKeyconfigurationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRecordingKeyconfigurationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordingKeyconfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRecordingKeyconfigurationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRecordingKeyconfigurationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRecordingKeyconfigurationsOK creates a PostRecordingKeyconfigurationsOK with default headers values
func NewPostRecordingKeyconfigurationsOK() *PostRecordingKeyconfigurationsOK {
	return &PostRecordingKeyconfigurationsOK{}
}

/*
PostRecordingKeyconfigurationsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostRecordingKeyconfigurationsOK struct {
	Payload *models.RecordingEncryptionConfiguration
}

// IsSuccess returns true when this post recording keyconfigurations o k response has a 2xx status code
func (o *PostRecordingKeyconfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post recording keyconfigurations o k response has a 3xx status code
func (o *PostRecordingKeyconfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations o k response has a 4xx status code
func (o *PostRecordingKeyconfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations o k response has a 5xx status code
func (o *PostRecordingKeyconfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations o k response a status code equal to that given
func (o *PostRecordingKeyconfigurationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostRecordingKeyconfigurationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsOK  %+v", 200, o.Payload)
}

func (o *PostRecordingKeyconfigurationsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsOK  %+v", 200, o.Payload)
}

func (o *PostRecordingKeyconfigurationsOK) GetPayload() *models.RecordingEncryptionConfiguration {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingEncryptionConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsBadRequest creates a PostRecordingKeyconfigurationsBadRequest with default headers values
func NewPostRecordingKeyconfigurationsBadRequest() *PostRecordingKeyconfigurationsBadRequest {
	return &PostRecordingKeyconfigurationsBadRequest{}
}

/*
PostRecordingKeyconfigurationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRecordingKeyconfigurationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations bad request response has a 2xx status code
func (o *PostRecordingKeyconfigurationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations bad request response has a 3xx status code
func (o *PostRecordingKeyconfigurationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations bad request response has a 4xx status code
func (o *PostRecordingKeyconfigurationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations bad request response has a 5xx status code
func (o *PostRecordingKeyconfigurationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations bad request response a status code equal to that given
func (o *PostRecordingKeyconfigurationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRecordingKeyconfigurationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingKeyconfigurationsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingKeyconfigurationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsUnauthorized creates a PostRecordingKeyconfigurationsUnauthorized with default headers values
func NewPostRecordingKeyconfigurationsUnauthorized() *PostRecordingKeyconfigurationsUnauthorized {
	return &PostRecordingKeyconfigurationsUnauthorized{}
}

/*
PostRecordingKeyconfigurationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRecordingKeyconfigurationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations unauthorized response has a 2xx status code
func (o *PostRecordingKeyconfigurationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations unauthorized response has a 3xx status code
func (o *PostRecordingKeyconfigurationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations unauthorized response has a 4xx status code
func (o *PostRecordingKeyconfigurationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations unauthorized response has a 5xx status code
func (o *PostRecordingKeyconfigurationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations unauthorized response a status code equal to that given
func (o *PostRecordingKeyconfigurationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRecordingKeyconfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingKeyconfigurationsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingKeyconfigurationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsForbidden creates a PostRecordingKeyconfigurationsForbidden with default headers values
func NewPostRecordingKeyconfigurationsForbidden() *PostRecordingKeyconfigurationsForbidden {
	return &PostRecordingKeyconfigurationsForbidden{}
}

/*
PostRecordingKeyconfigurationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRecordingKeyconfigurationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations forbidden response has a 2xx status code
func (o *PostRecordingKeyconfigurationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations forbidden response has a 3xx status code
func (o *PostRecordingKeyconfigurationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations forbidden response has a 4xx status code
func (o *PostRecordingKeyconfigurationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations forbidden response has a 5xx status code
func (o *PostRecordingKeyconfigurationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations forbidden response a status code equal to that given
func (o *PostRecordingKeyconfigurationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRecordingKeyconfigurationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingKeyconfigurationsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingKeyconfigurationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsNotFound creates a PostRecordingKeyconfigurationsNotFound with default headers values
func NewPostRecordingKeyconfigurationsNotFound() *PostRecordingKeyconfigurationsNotFound {
	return &PostRecordingKeyconfigurationsNotFound{}
}

/*
PostRecordingKeyconfigurationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRecordingKeyconfigurationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations not found response has a 2xx status code
func (o *PostRecordingKeyconfigurationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations not found response has a 3xx status code
func (o *PostRecordingKeyconfigurationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations not found response has a 4xx status code
func (o *PostRecordingKeyconfigurationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations not found response has a 5xx status code
func (o *PostRecordingKeyconfigurationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations not found response a status code equal to that given
func (o *PostRecordingKeyconfigurationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRecordingKeyconfigurationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingKeyconfigurationsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingKeyconfigurationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsRequestTimeout creates a PostRecordingKeyconfigurationsRequestTimeout with default headers values
func NewPostRecordingKeyconfigurationsRequestTimeout() *PostRecordingKeyconfigurationsRequestTimeout {
	return &PostRecordingKeyconfigurationsRequestTimeout{}
}

/*
PostRecordingKeyconfigurationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRecordingKeyconfigurationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations request timeout response has a 2xx status code
func (o *PostRecordingKeyconfigurationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations request timeout response has a 3xx status code
func (o *PostRecordingKeyconfigurationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations request timeout response has a 4xx status code
func (o *PostRecordingKeyconfigurationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations request timeout response has a 5xx status code
func (o *PostRecordingKeyconfigurationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations request timeout response a status code equal to that given
func (o *PostRecordingKeyconfigurationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRecordingKeyconfigurationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingKeyconfigurationsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingKeyconfigurationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsRequestEntityTooLarge creates a PostRecordingKeyconfigurationsRequestEntityTooLarge with default headers values
func NewPostRecordingKeyconfigurationsRequestEntityTooLarge() *PostRecordingKeyconfigurationsRequestEntityTooLarge {
	return &PostRecordingKeyconfigurationsRequestEntityTooLarge{}
}

/*
PostRecordingKeyconfigurationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostRecordingKeyconfigurationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations request entity too large response has a 2xx status code
func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations request entity too large response has a 3xx status code
func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations request entity too large response has a 4xx status code
func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations request entity too large response has a 5xx status code
func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations request entity too large response a status code equal to that given
func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsUnsupportedMediaType creates a PostRecordingKeyconfigurationsUnsupportedMediaType with default headers values
func NewPostRecordingKeyconfigurationsUnsupportedMediaType() *PostRecordingKeyconfigurationsUnsupportedMediaType {
	return &PostRecordingKeyconfigurationsUnsupportedMediaType{}
}

/*
PostRecordingKeyconfigurationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRecordingKeyconfigurationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations unsupported media type response has a 2xx status code
func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations unsupported media type response has a 3xx status code
func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations unsupported media type response has a 4xx status code
func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations unsupported media type response has a 5xx status code
func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations unsupported media type response a status code equal to that given
func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsTooManyRequests creates a PostRecordingKeyconfigurationsTooManyRequests with default headers values
func NewPostRecordingKeyconfigurationsTooManyRequests() *PostRecordingKeyconfigurationsTooManyRequests {
	return &PostRecordingKeyconfigurationsTooManyRequests{}
}

/*
PostRecordingKeyconfigurationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRecordingKeyconfigurationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations too many requests response has a 2xx status code
func (o *PostRecordingKeyconfigurationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations too many requests response has a 3xx status code
func (o *PostRecordingKeyconfigurationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations too many requests response has a 4xx status code
func (o *PostRecordingKeyconfigurationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations too many requests response has a 5xx status code
func (o *PostRecordingKeyconfigurationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations too many requests response a status code equal to that given
func (o *PostRecordingKeyconfigurationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRecordingKeyconfigurationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingKeyconfigurationsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingKeyconfigurationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsInternalServerError creates a PostRecordingKeyconfigurationsInternalServerError with default headers values
func NewPostRecordingKeyconfigurationsInternalServerError() *PostRecordingKeyconfigurationsInternalServerError {
	return &PostRecordingKeyconfigurationsInternalServerError{}
}

/*
PostRecordingKeyconfigurationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRecordingKeyconfigurationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations internal server error response has a 2xx status code
func (o *PostRecordingKeyconfigurationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations internal server error response has a 3xx status code
func (o *PostRecordingKeyconfigurationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations internal server error response has a 4xx status code
func (o *PostRecordingKeyconfigurationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations internal server error response has a 5xx status code
func (o *PostRecordingKeyconfigurationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording keyconfigurations internal server error response a status code equal to that given
func (o *PostRecordingKeyconfigurationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRecordingKeyconfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingKeyconfigurationsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingKeyconfigurationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsServiceUnavailable creates a PostRecordingKeyconfigurationsServiceUnavailable with default headers values
func NewPostRecordingKeyconfigurationsServiceUnavailable() *PostRecordingKeyconfigurationsServiceUnavailable {
	return &PostRecordingKeyconfigurationsServiceUnavailable{}
}

/*
PostRecordingKeyconfigurationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRecordingKeyconfigurationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations service unavailable response has a 2xx status code
func (o *PostRecordingKeyconfigurationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations service unavailable response has a 3xx status code
func (o *PostRecordingKeyconfigurationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations service unavailable response has a 4xx status code
func (o *PostRecordingKeyconfigurationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations service unavailable response has a 5xx status code
func (o *PostRecordingKeyconfigurationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording keyconfigurations service unavailable response a status code equal to that given
func (o *PostRecordingKeyconfigurationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRecordingKeyconfigurationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingKeyconfigurationsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingKeyconfigurationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsGatewayTimeout creates a PostRecordingKeyconfigurationsGatewayTimeout with default headers values
func NewPostRecordingKeyconfigurationsGatewayTimeout() *PostRecordingKeyconfigurationsGatewayTimeout {
	return &PostRecordingKeyconfigurationsGatewayTimeout{}
}

/*
PostRecordingKeyconfigurationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRecordingKeyconfigurationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations gateway timeout response has a 2xx status code
func (o *PostRecordingKeyconfigurationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations gateway timeout response has a 3xx status code
func (o *PostRecordingKeyconfigurationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations gateway timeout response has a 4xx status code
func (o *PostRecordingKeyconfigurationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations gateway timeout response has a 5xx status code
func (o *PostRecordingKeyconfigurationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording keyconfigurations gateway timeout response a status code equal to that given
func (o *PostRecordingKeyconfigurationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRecordingKeyconfigurationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingKeyconfigurationsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations][%d] postRecordingKeyconfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingKeyconfigurationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
