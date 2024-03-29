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

// PostRecordingKeyconfigurationsValidateReader is a Reader for the PostRecordingKeyconfigurationsValidate structure.
type PostRecordingKeyconfigurationsValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordingKeyconfigurationsValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRecordingKeyconfigurationsValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRecordingKeyconfigurationsValidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRecordingKeyconfigurationsValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRecordingKeyconfigurationsValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordingKeyconfigurationsValidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRecordingKeyconfigurationsValidateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRecordingKeyconfigurationsValidateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRecordingKeyconfigurationsValidateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRecordingKeyconfigurationsValidateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordingKeyconfigurationsValidateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRecordingKeyconfigurationsValidateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRecordingKeyconfigurationsValidateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRecordingKeyconfigurationsValidateOK creates a PostRecordingKeyconfigurationsValidateOK with default headers values
func NewPostRecordingKeyconfigurationsValidateOK() *PostRecordingKeyconfigurationsValidateOK {
	return &PostRecordingKeyconfigurationsValidateOK{}
}

/*
PostRecordingKeyconfigurationsValidateOK describes a response with status code 200, with default header values.

successful operation
*/
type PostRecordingKeyconfigurationsValidateOK struct {
	Payload *models.RecordingEncryptionConfiguration
}

// IsSuccess returns true when this post recording keyconfigurations validate o k response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post recording keyconfigurations validate o k response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate o k response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations validate o k response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate o k response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostRecordingKeyconfigurationsValidateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateOK  %+v", 200, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateOK) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateOK  %+v", 200, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateOK) GetPayload() *models.RecordingEncryptionConfiguration {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingEncryptionConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateBadRequest creates a PostRecordingKeyconfigurationsValidateBadRequest with default headers values
func NewPostRecordingKeyconfigurationsValidateBadRequest() *PostRecordingKeyconfigurationsValidateBadRequest {
	return &PostRecordingKeyconfigurationsValidateBadRequest{}
}

/*
PostRecordingKeyconfigurationsValidateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRecordingKeyconfigurationsValidateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate bad request response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate bad request response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate bad request response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate bad request response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate bad request response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRecordingKeyconfigurationsValidateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateUnauthorized creates a PostRecordingKeyconfigurationsValidateUnauthorized with default headers values
func NewPostRecordingKeyconfigurationsValidateUnauthorized() *PostRecordingKeyconfigurationsValidateUnauthorized {
	return &PostRecordingKeyconfigurationsValidateUnauthorized{}
}

/*
PostRecordingKeyconfigurationsValidateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRecordingKeyconfigurationsValidateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate unauthorized response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate unauthorized response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate unauthorized response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate unauthorized response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate unauthorized response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRecordingKeyconfigurationsValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateForbidden creates a PostRecordingKeyconfigurationsValidateForbidden with default headers values
func NewPostRecordingKeyconfigurationsValidateForbidden() *PostRecordingKeyconfigurationsValidateForbidden {
	return &PostRecordingKeyconfigurationsValidateForbidden{}
}

/*
PostRecordingKeyconfigurationsValidateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRecordingKeyconfigurationsValidateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate forbidden response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate forbidden response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate forbidden response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate forbidden response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate forbidden response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRecordingKeyconfigurationsValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateNotFound creates a PostRecordingKeyconfigurationsValidateNotFound with default headers values
func NewPostRecordingKeyconfigurationsValidateNotFound() *PostRecordingKeyconfigurationsValidateNotFound {
	return &PostRecordingKeyconfigurationsValidateNotFound{}
}

/*
PostRecordingKeyconfigurationsValidateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRecordingKeyconfigurationsValidateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate not found response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate not found response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate not found response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate not found response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate not found response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRecordingKeyconfigurationsValidateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateRequestTimeout creates a PostRecordingKeyconfigurationsValidateRequestTimeout with default headers values
func NewPostRecordingKeyconfigurationsValidateRequestTimeout() *PostRecordingKeyconfigurationsValidateRequestTimeout {
	return &PostRecordingKeyconfigurationsValidateRequestTimeout{}
}

/*
PostRecordingKeyconfigurationsValidateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRecordingKeyconfigurationsValidateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate request timeout response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate request timeout response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate request timeout response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate request timeout response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate request timeout response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateRequestEntityTooLarge creates a PostRecordingKeyconfigurationsValidateRequestEntityTooLarge with default headers values
func NewPostRecordingKeyconfigurationsValidateRequestEntityTooLarge() *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge {
	return &PostRecordingKeyconfigurationsValidateRequestEntityTooLarge{}
}

/*
PostRecordingKeyconfigurationsValidateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostRecordingKeyconfigurationsValidateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate request entity too large response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate request entity too large response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate request entity too large response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate request entity too large response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate request entity too large response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateUnsupportedMediaType creates a PostRecordingKeyconfigurationsValidateUnsupportedMediaType with default headers values
func NewPostRecordingKeyconfigurationsValidateUnsupportedMediaType() *PostRecordingKeyconfigurationsValidateUnsupportedMediaType {
	return &PostRecordingKeyconfigurationsValidateUnsupportedMediaType{}
}

/*
PostRecordingKeyconfigurationsValidateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRecordingKeyconfigurationsValidateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate unsupported media type response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate unsupported media type response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate unsupported media type response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate unsupported media type response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate unsupported media type response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateTooManyRequests creates a PostRecordingKeyconfigurationsValidateTooManyRequests with default headers values
func NewPostRecordingKeyconfigurationsValidateTooManyRequests() *PostRecordingKeyconfigurationsValidateTooManyRequests {
	return &PostRecordingKeyconfigurationsValidateTooManyRequests{}
}

/*
PostRecordingKeyconfigurationsValidateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRecordingKeyconfigurationsValidateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate too many requests response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate too many requests response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate too many requests response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post recording keyconfigurations validate too many requests response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post recording keyconfigurations validate too many requests response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateInternalServerError creates a PostRecordingKeyconfigurationsValidateInternalServerError with default headers values
func NewPostRecordingKeyconfigurationsValidateInternalServerError() *PostRecordingKeyconfigurationsValidateInternalServerError {
	return &PostRecordingKeyconfigurationsValidateInternalServerError{}
}

/*
PostRecordingKeyconfigurationsValidateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRecordingKeyconfigurationsValidateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate internal server error response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate internal server error response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate internal server error response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations validate internal server error response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording keyconfigurations validate internal server error response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRecordingKeyconfigurationsValidateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateServiceUnavailable creates a PostRecordingKeyconfigurationsValidateServiceUnavailable with default headers values
func NewPostRecordingKeyconfigurationsValidateServiceUnavailable() *PostRecordingKeyconfigurationsValidateServiceUnavailable {
	return &PostRecordingKeyconfigurationsValidateServiceUnavailable{}
}

/*
PostRecordingKeyconfigurationsValidateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRecordingKeyconfigurationsValidateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate service unavailable response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate service unavailable response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate service unavailable response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations validate service unavailable response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording keyconfigurations validate service unavailable response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingKeyconfigurationsValidateGatewayTimeout creates a PostRecordingKeyconfigurationsValidateGatewayTimeout with default headers values
func NewPostRecordingKeyconfigurationsValidateGatewayTimeout() *PostRecordingKeyconfigurationsValidateGatewayTimeout {
	return &PostRecordingKeyconfigurationsValidateGatewayTimeout{}
}

/*
PostRecordingKeyconfigurationsValidateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRecordingKeyconfigurationsValidateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post recording keyconfigurations validate gateway timeout response has a 2xx status code
func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post recording keyconfigurations validate gateway timeout response has a 3xx status code
func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post recording keyconfigurations validate gateway timeout response has a 4xx status code
func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post recording keyconfigurations validate gateway timeout response has a 5xx status code
func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post recording keyconfigurations validate gateway timeout response a status code equal to that given
func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/recording/keyconfigurations/validate][%d] postRecordingKeyconfigurationsValidateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingKeyconfigurationsValidateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
