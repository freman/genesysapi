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

// GetConversationsKeyconfigurationReader is a Reader for the GetConversationsKeyconfiguration structure.
type GetConversationsKeyconfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsKeyconfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsKeyconfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsKeyconfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsKeyconfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsKeyconfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsKeyconfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsKeyconfigurationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsKeyconfigurationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsKeyconfigurationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsKeyconfigurationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsKeyconfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsKeyconfigurationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsKeyconfigurationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsKeyconfigurationOK creates a GetConversationsKeyconfigurationOK with default headers values
func NewGetConversationsKeyconfigurationOK() *GetConversationsKeyconfigurationOK {
	return &GetConversationsKeyconfigurationOK{}
}

/*
GetConversationsKeyconfigurationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsKeyconfigurationOK struct {
	Payload *models.ConversationEncryptionConfiguration
}

// IsSuccess returns true when this get conversations keyconfiguration o k response has a 2xx status code
func (o *GetConversationsKeyconfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations keyconfiguration o k response has a 3xx status code
func (o *GetConversationsKeyconfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration o k response has a 4xx status code
func (o *GetConversationsKeyconfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations keyconfiguration o k response has a 5xx status code
func (o *GetConversationsKeyconfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration o k response a status code equal to that given
func (o *GetConversationsKeyconfigurationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsKeyconfigurationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationOK  %+v", 200, o.Payload)
}

func (o *GetConversationsKeyconfigurationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationOK  %+v", 200, o.Payload)
}

func (o *GetConversationsKeyconfigurationOK) GetPayload() *models.ConversationEncryptionConfiguration {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConversationEncryptionConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationBadRequest creates a GetConversationsKeyconfigurationBadRequest with default headers values
func NewGetConversationsKeyconfigurationBadRequest() *GetConversationsKeyconfigurationBadRequest {
	return &GetConversationsKeyconfigurationBadRequest{}
}

/*
GetConversationsKeyconfigurationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsKeyconfigurationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration bad request response has a 2xx status code
func (o *GetConversationsKeyconfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration bad request response has a 3xx status code
func (o *GetConversationsKeyconfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration bad request response has a 4xx status code
func (o *GetConversationsKeyconfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration bad request response has a 5xx status code
func (o *GetConversationsKeyconfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration bad request response a status code equal to that given
func (o *GetConversationsKeyconfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsKeyconfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsKeyconfigurationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsKeyconfigurationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationUnauthorized creates a GetConversationsKeyconfigurationUnauthorized with default headers values
func NewGetConversationsKeyconfigurationUnauthorized() *GetConversationsKeyconfigurationUnauthorized {
	return &GetConversationsKeyconfigurationUnauthorized{}
}

/*
GetConversationsKeyconfigurationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsKeyconfigurationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration unauthorized response has a 2xx status code
func (o *GetConversationsKeyconfigurationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration unauthorized response has a 3xx status code
func (o *GetConversationsKeyconfigurationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration unauthorized response has a 4xx status code
func (o *GetConversationsKeyconfigurationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration unauthorized response has a 5xx status code
func (o *GetConversationsKeyconfigurationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration unauthorized response a status code equal to that given
func (o *GetConversationsKeyconfigurationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsKeyconfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsKeyconfigurationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsKeyconfigurationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationForbidden creates a GetConversationsKeyconfigurationForbidden with default headers values
func NewGetConversationsKeyconfigurationForbidden() *GetConversationsKeyconfigurationForbidden {
	return &GetConversationsKeyconfigurationForbidden{}
}

/*
GetConversationsKeyconfigurationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsKeyconfigurationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration forbidden response has a 2xx status code
func (o *GetConversationsKeyconfigurationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration forbidden response has a 3xx status code
func (o *GetConversationsKeyconfigurationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration forbidden response has a 4xx status code
func (o *GetConversationsKeyconfigurationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration forbidden response has a 5xx status code
func (o *GetConversationsKeyconfigurationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration forbidden response a status code equal to that given
func (o *GetConversationsKeyconfigurationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsKeyconfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsKeyconfigurationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsKeyconfigurationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationNotFound creates a GetConversationsKeyconfigurationNotFound with default headers values
func NewGetConversationsKeyconfigurationNotFound() *GetConversationsKeyconfigurationNotFound {
	return &GetConversationsKeyconfigurationNotFound{}
}

/*
GetConversationsKeyconfigurationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsKeyconfigurationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration not found response has a 2xx status code
func (o *GetConversationsKeyconfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration not found response has a 3xx status code
func (o *GetConversationsKeyconfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration not found response has a 4xx status code
func (o *GetConversationsKeyconfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration not found response has a 5xx status code
func (o *GetConversationsKeyconfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration not found response a status code equal to that given
func (o *GetConversationsKeyconfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsKeyconfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsKeyconfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsKeyconfigurationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationRequestTimeout creates a GetConversationsKeyconfigurationRequestTimeout with default headers values
func NewGetConversationsKeyconfigurationRequestTimeout() *GetConversationsKeyconfigurationRequestTimeout {
	return &GetConversationsKeyconfigurationRequestTimeout{}
}

/*
GetConversationsKeyconfigurationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsKeyconfigurationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration request timeout response has a 2xx status code
func (o *GetConversationsKeyconfigurationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration request timeout response has a 3xx status code
func (o *GetConversationsKeyconfigurationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration request timeout response has a 4xx status code
func (o *GetConversationsKeyconfigurationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration request timeout response has a 5xx status code
func (o *GetConversationsKeyconfigurationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration request timeout response a status code equal to that given
func (o *GetConversationsKeyconfigurationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsKeyconfigurationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsKeyconfigurationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsKeyconfigurationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationRequestEntityTooLarge creates a GetConversationsKeyconfigurationRequestEntityTooLarge with default headers values
func NewGetConversationsKeyconfigurationRequestEntityTooLarge() *GetConversationsKeyconfigurationRequestEntityTooLarge {
	return &GetConversationsKeyconfigurationRequestEntityTooLarge{}
}

/*
GetConversationsKeyconfigurationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsKeyconfigurationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration request entity too large response has a 2xx status code
func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration request entity too large response has a 3xx status code
func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration request entity too large response has a 4xx status code
func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration request entity too large response has a 5xx status code
func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration request entity too large response a status code equal to that given
func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationUnsupportedMediaType creates a GetConversationsKeyconfigurationUnsupportedMediaType with default headers values
func NewGetConversationsKeyconfigurationUnsupportedMediaType() *GetConversationsKeyconfigurationUnsupportedMediaType {
	return &GetConversationsKeyconfigurationUnsupportedMediaType{}
}

/*
GetConversationsKeyconfigurationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsKeyconfigurationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration unsupported media type response has a 2xx status code
func (o *GetConversationsKeyconfigurationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration unsupported media type response has a 3xx status code
func (o *GetConversationsKeyconfigurationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration unsupported media type response has a 4xx status code
func (o *GetConversationsKeyconfigurationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration unsupported media type response has a 5xx status code
func (o *GetConversationsKeyconfigurationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration unsupported media type response a status code equal to that given
func (o *GetConversationsKeyconfigurationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsKeyconfigurationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsKeyconfigurationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsKeyconfigurationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationTooManyRequests creates a GetConversationsKeyconfigurationTooManyRequests with default headers values
func NewGetConversationsKeyconfigurationTooManyRequests() *GetConversationsKeyconfigurationTooManyRequests {
	return &GetConversationsKeyconfigurationTooManyRequests{}
}

/*
GetConversationsKeyconfigurationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsKeyconfigurationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration too many requests response has a 2xx status code
func (o *GetConversationsKeyconfigurationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration too many requests response has a 3xx status code
func (o *GetConversationsKeyconfigurationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration too many requests response has a 4xx status code
func (o *GetConversationsKeyconfigurationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations keyconfiguration too many requests response has a 5xx status code
func (o *GetConversationsKeyconfigurationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations keyconfiguration too many requests response a status code equal to that given
func (o *GetConversationsKeyconfigurationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsKeyconfigurationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsKeyconfigurationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsKeyconfigurationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationInternalServerError creates a GetConversationsKeyconfigurationInternalServerError with default headers values
func NewGetConversationsKeyconfigurationInternalServerError() *GetConversationsKeyconfigurationInternalServerError {
	return &GetConversationsKeyconfigurationInternalServerError{}
}

/*
GetConversationsKeyconfigurationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsKeyconfigurationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration internal server error response has a 2xx status code
func (o *GetConversationsKeyconfigurationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration internal server error response has a 3xx status code
func (o *GetConversationsKeyconfigurationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration internal server error response has a 4xx status code
func (o *GetConversationsKeyconfigurationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations keyconfiguration internal server error response has a 5xx status code
func (o *GetConversationsKeyconfigurationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations keyconfiguration internal server error response a status code equal to that given
func (o *GetConversationsKeyconfigurationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsKeyconfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsKeyconfigurationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsKeyconfigurationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationServiceUnavailable creates a GetConversationsKeyconfigurationServiceUnavailable with default headers values
func NewGetConversationsKeyconfigurationServiceUnavailable() *GetConversationsKeyconfigurationServiceUnavailable {
	return &GetConversationsKeyconfigurationServiceUnavailable{}
}

/*
GetConversationsKeyconfigurationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsKeyconfigurationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration service unavailable response has a 2xx status code
func (o *GetConversationsKeyconfigurationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration service unavailable response has a 3xx status code
func (o *GetConversationsKeyconfigurationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration service unavailable response has a 4xx status code
func (o *GetConversationsKeyconfigurationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations keyconfiguration service unavailable response has a 5xx status code
func (o *GetConversationsKeyconfigurationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations keyconfiguration service unavailable response a status code equal to that given
func (o *GetConversationsKeyconfigurationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsKeyconfigurationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsKeyconfigurationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsKeyconfigurationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsKeyconfigurationGatewayTimeout creates a GetConversationsKeyconfigurationGatewayTimeout with default headers values
func NewGetConversationsKeyconfigurationGatewayTimeout() *GetConversationsKeyconfigurationGatewayTimeout {
	return &GetConversationsKeyconfigurationGatewayTimeout{}
}

/*
GetConversationsKeyconfigurationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsKeyconfigurationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations keyconfiguration gateway timeout response has a 2xx status code
func (o *GetConversationsKeyconfigurationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations keyconfiguration gateway timeout response has a 3xx status code
func (o *GetConversationsKeyconfigurationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations keyconfiguration gateway timeout response has a 4xx status code
func (o *GetConversationsKeyconfigurationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations keyconfiguration gateway timeout response has a 5xx status code
func (o *GetConversationsKeyconfigurationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations keyconfiguration gateway timeout response a status code equal to that given
func (o *GetConversationsKeyconfigurationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsKeyconfigurationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsKeyconfigurationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/keyconfigurations/{keyconfigurationsId}][%d] getConversationsKeyconfigurationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsKeyconfigurationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsKeyconfigurationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
