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

// GetRecordingKeyconfigurationsReader is a Reader for the GetRecordingKeyconfigurations structure.
type GetRecordingKeyconfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingKeyconfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingKeyconfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingKeyconfigurationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingKeyconfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingKeyconfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingKeyconfigurationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRecordingKeyconfigurationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingKeyconfigurationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingKeyconfigurationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingKeyconfigurationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingKeyconfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingKeyconfigurationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingKeyconfigurationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingKeyconfigurationsOK creates a GetRecordingKeyconfigurationsOK with default headers values
func NewGetRecordingKeyconfigurationsOK() *GetRecordingKeyconfigurationsOK {
	return &GetRecordingKeyconfigurationsOK{}
}

/*
GetRecordingKeyconfigurationsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRecordingKeyconfigurationsOK struct {
	Payload *models.RecordingEncryptionConfigurationListing
}

// IsSuccess returns true when this get recording keyconfigurations o k response has a 2xx status code
func (o *GetRecordingKeyconfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get recording keyconfigurations o k response has a 3xx status code
func (o *GetRecordingKeyconfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations o k response has a 4xx status code
func (o *GetRecordingKeyconfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording keyconfigurations o k response has a 5xx status code
func (o *GetRecordingKeyconfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations o k response a status code equal to that given
func (o *GetRecordingKeyconfigurationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRecordingKeyconfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingKeyconfigurationsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingKeyconfigurationsOK) GetPayload() *models.RecordingEncryptionConfigurationListing {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingEncryptionConfigurationListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsBadRequest creates a GetRecordingKeyconfigurationsBadRequest with default headers values
func NewGetRecordingKeyconfigurationsBadRequest() *GetRecordingKeyconfigurationsBadRequest {
	return &GetRecordingKeyconfigurationsBadRequest{}
}

/*
GetRecordingKeyconfigurationsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingKeyconfigurationsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations bad request response has a 2xx status code
func (o *GetRecordingKeyconfigurationsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations bad request response has a 3xx status code
func (o *GetRecordingKeyconfigurationsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations bad request response has a 4xx status code
func (o *GetRecordingKeyconfigurationsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations bad request response has a 5xx status code
func (o *GetRecordingKeyconfigurationsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations bad request response a status code equal to that given
func (o *GetRecordingKeyconfigurationsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRecordingKeyconfigurationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingKeyconfigurationsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingKeyconfigurationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsUnauthorized creates a GetRecordingKeyconfigurationsUnauthorized with default headers values
func NewGetRecordingKeyconfigurationsUnauthorized() *GetRecordingKeyconfigurationsUnauthorized {
	return &GetRecordingKeyconfigurationsUnauthorized{}
}

/*
GetRecordingKeyconfigurationsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingKeyconfigurationsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations unauthorized response has a 2xx status code
func (o *GetRecordingKeyconfigurationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations unauthorized response has a 3xx status code
func (o *GetRecordingKeyconfigurationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations unauthorized response has a 4xx status code
func (o *GetRecordingKeyconfigurationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations unauthorized response has a 5xx status code
func (o *GetRecordingKeyconfigurationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations unauthorized response a status code equal to that given
func (o *GetRecordingKeyconfigurationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRecordingKeyconfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingKeyconfigurationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingKeyconfigurationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsForbidden creates a GetRecordingKeyconfigurationsForbidden with default headers values
func NewGetRecordingKeyconfigurationsForbidden() *GetRecordingKeyconfigurationsForbidden {
	return &GetRecordingKeyconfigurationsForbidden{}
}

/*
GetRecordingKeyconfigurationsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingKeyconfigurationsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations forbidden response has a 2xx status code
func (o *GetRecordingKeyconfigurationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations forbidden response has a 3xx status code
func (o *GetRecordingKeyconfigurationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations forbidden response has a 4xx status code
func (o *GetRecordingKeyconfigurationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations forbidden response has a 5xx status code
func (o *GetRecordingKeyconfigurationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations forbidden response a status code equal to that given
func (o *GetRecordingKeyconfigurationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRecordingKeyconfigurationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingKeyconfigurationsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingKeyconfigurationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsNotFound creates a GetRecordingKeyconfigurationsNotFound with default headers values
func NewGetRecordingKeyconfigurationsNotFound() *GetRecordingKeyconfigurationsNotFound {
	return &GetRecordingKeyconfigurationsNotFound{}
}

/*
GetRecordingKeyconfigurationsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRecordingKeyconfigurationsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations not found response has a 2xx status code
func (o *GetRecordingKeyconfigurationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations not found response has a 3xx status code
func (o *GetRecordingKeyconfigurationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations not found response has a 4xx status code
func (o *GetRecordingKeyconfigurationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations not found response has a 5xx status code
func (o *GetRecordingKeyconfigurationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations not found response a status code equal to that given
func (o *GetRecordingKeyconfigurationsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRecordingKeyconfigurationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingKeyconfigurationsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingKeyconfigurationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsRequestTimeout creates a GetRecordingKeyconfigurationsRequestTimeout with default headers values
func NewGetRecordingKeyconfigurationsRequestTimeout() *GetRecordingKeyconfigurationsRequestTimeout {
	return &GetRecordingKeyconfigurationsRequestTimeout{}
}

/*
GetRecordingKeyconfigurationsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingKeyconfigurationsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations request timeout response has a 2xx status code
func (o *GetRecordingKeyconfigurationsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations request timeout response has a 3xx status code
func (o *GetRecordingKeyconfigurationsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations request timeout response has a 4xx status code
func (o *GetRecordingKeyconfigurationsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations request timeout response has a 5xx status code
func (o *GetRecordingKeyconfigurationsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations request timeout response a status code equal to that given
func (o *GetRecordingKeyconfigurationsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRecordingKeyconfigurationsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingKeyconfigurationsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingKeyconfigurationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsRequestEntityTooLarge creates a GetRecordingKeyconfigurationsRequestEntityTooLarge with default headers values
func NewGetRecordingKeyconfigurationsRequestEntityTooLarge() *GetRecordingKeyconfigurationsRequestEntityTooLarge {
	return &GetRecordingKeyconfigurationsRequestEntityTooLarge{}
}

/*
GetRecordingKeyconfigurationsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRecordingKeyconfigurationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations request entity too large response has a 2xx status code
func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations request entity too large response has a 3xx status code
func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations request entity too large response has a 4xx status code
func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations request entity too large response has a 5xx status code
func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations request entity too large response a status code equal to that given
func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsUnsupportedMediaType creates a GetRecordingKeyconfigurationsUnsupportedMediaType with default headers values
func NewGetRecordingKeyconfigurationsUnsupportedMediaType() *GetRecordingKeyconfigurationsUnsupportedMediaType {
	return &GetRecordingKeyconfigurationsUnsupportedMediaType{}
}

/*
GetRecordingKeyconfigurationsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingKeyconfigurationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations unsupported media type response has a 2xx status code
func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations unsupported media type response has a 3xx status code
func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations unsupported media type response has a 4xx status code
func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations unsupported media type response has a 5xx status code
func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations unsupported media type response a status code equal to that given
func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsTooManyRequests creates a GetRecordingKeyconfigurationsTooManyRequests with default headers values
func NewGetRecordingKeyconfigurationsTooManyRequests() *GetRecordingKeyconfigurationsTooManyRequests {
	return &GetRecordingKeyconfigurationsTooManyRequests{}
}

/*
GetRecordingKeyconfigurationsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingKeyconfigurationsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations too many requests response has a 2xx status code
func (o *GetRecordingKeyconfigurationsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations too many requests response has a 3xx status code
func (o *GetRecordingKeyconfigurationsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations too many requests response has a 4xx status code
func (o *GetRecordingKeyconfigurationsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording keyconfigurations too many requests response has a 5xx status code
func (o *GetRecordingKeyconfigurationsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording keyconfigurations too many requests response a status code equal to that given
func (o *GetRecordingKeyconfigurationsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRecordingKeyconfigurationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingKeyconfigurationsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingKeyconfigurationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsInternalServerError creates a GetRecordingKeyconfigurationsInternalServerError with default headers values
func NewGetRecordingKeyconfigurationsInternalServerError() *GetRecordingKeyconfigurationsInternalServerError {
	return &GetRecordingKeyconfigurationsInternalServerError{}
}

/*
GetRecordingKeyconfigurationsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingKeyconfigurationsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations internal server error response has a 2xx status code
func (o *GetRecordingKeyconfigurationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations internal server error response has a 3xx status code
func (o *GetRecordingKeyconfigurationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations internal server error response has a 4xx status code
func (o *GetRecordingKeyconfigurationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording keyconfigurations internal server error response has a 5xx status code
func (o *GetRecordingKeyconfigurationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording keyconfigurations internal server error response a status code equal to that given
func (o *GetRecordingKeyconfigurationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRecordingKeyconfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingKeyconfigurationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingKeyconfigurationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsServiceUnavailable creates a GetRecordingKeyconfigurationsServiceUnavailable with default headers values
func NewGetRecordingKeyconfigurationsServiceUnavailable() *GetRecordingKeyconfigurationsServiceUnavailable {
	return &GetRecordingKeyconfigurationsServiceUnavailable{}
}

/*
GetRecordingKeyconfigurationsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingKeyconfigurationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations service unavailable response has a 2xx status code
func (o *GetRecordingKeyconfigurationsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations service unavailable response has a 3xx status code
func (o *GetRecordingKeyconfigurationsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations service unavailable response has a 4xx status code
func (o *GetRecordingKeyconfigurationsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording keyconfigurations service unavailable response has a 5xx status code
func (o *GetRecordingKeyconfigurationsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording keyconfigurations service unavailable response a status code equal to that given
func (o *GetRecordingKeyconfigurationsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRecordingKeyconfigurationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingKeyconfigurationsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingKeyconfigurationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingKeyconfigurationsGatewayTimeout creates a GetRecordingKeyconfigurationsGatewayTimeout with default headers values
func NewGetRecordingKeyconfigurationsGatewayTimeout() *GetRecordingKeyconfigurationsGatewayTimeout {
	return &GetRecordingKeyconfigurationsGatewayTimeout{}
}

/*
GetRecordingKeyconfigurationsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRecordingKeyconfigurationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording keyconfigurations gateway timeout response has a 2xx status code
func (o *GetRecordingKeyconfigurationsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording keyconfigurations gateway timeout response has a 3xx status code
func (o *GetRecordingKeyconfigurationsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording keyconfigurations gateway timeout response has a 4xx status code
func (o *GetRecordingKeyconfigurationsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording keyconfigurations gateway timeout response has a 5xx status code
func (o *GetRecordingKeyconfigurationsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording keyconfigurations gateway timeout response a status code equal to that given
func (o *GetRecordingKeyconfigurationsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRecordingKeyconfigurationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingKeyconfigurationsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/keyconfigurations][%d] getRecordingKeyconfigurationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingKeyconfigurationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingKeyconfigurationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
