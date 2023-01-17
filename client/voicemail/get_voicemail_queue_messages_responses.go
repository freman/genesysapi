// Code generated by go-swagger; DO NOT EDIT.

package voicemail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetVoicemailQueueMessagesReader is a Reader for the GetVoicemailQueueMessages structure.
type GetVoicemailQueueMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailQueueMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailQueueMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailQueueMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailQueueMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailQueueMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailQueueMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetVoicemailQueueMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailQueueMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailQueueMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailQueueMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailQueueMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailQueueMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailQueueMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailQueueMessagesOK creates a GetVoicemailQueueMessagesOK with default headers values
func NewGetVoicemailQueueMessagesOK() *GetVoicemailQueueMessagesOK {
	return &GetVoicemailQueueMessagesOK{}
}

/*
GetVoicemailQueueMessagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVoicemailQueueMessagesOK struct {
	Payload *models.VoicemailMessageEntityListing
}

// IsSuccess returns true when this get voicemail queue messages o k response has a 2xx status code
func (o *GetVoicemailQueueMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get voicemail queue messages o k response has a 3xx status code
func (o *GetVoicemailQueueMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages o k response has a 4xx status code
func (o *GetVoicemailQueueMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail queue messages o k response has a 5xx status code
func (o *GetVoicemailQueueMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages o k response a status code equal to that given
func (o *GetVoicemailQueueMessagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVoicemailQueueMessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailQueueMessagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailQueueMessagesOK) GetPayload() *models.VoicemailMessageEntityListing {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMessageEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesBadRequest creates a GetVoicemailQueueMessagesBadRequest with default headers values
func NewGetVoicemailQueueMessagesBadRequest() *GetVoicemailQueueMessagesBadRequest {
	return &GetVoicemailQueueMessagesBadRequest{}
}

/*
GetVoicemailQueueMessagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailQueueMessagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages bad request response has a 2xx status code
func (o *GetVoicemailQueueMessagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages bad request response has a 3xx status code
func (o *GetVoicemailQueueMessagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages bad request response has a 4xx status code
func (o *GetVoicemailQueueMessagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages bad request response has a 5xx status code
func (o *GetVoicemailQueueMessagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages bad request response a status code equal to that given
func (o *GetVoicemailQueueMessagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVoicemailQueueMessagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailQueueMessagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailQueueMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesUnauthorized creates a GetVoicemailQueueMessagesUnauthorized with default headers values
func NewGetVoicemailQueueMessagesUnauthorized() *GetVoicemailQueueMessagesUnauthorized {
	return &GetVoicemailQueueMessagesUnauthorized{}
}

/*
GetVoicemailQueueMessagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailQueueMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages unauthorized response has a 2xx status code
func (o *GetVoicemailQueueMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages unauthorized response has a 3xx status code
func (o *GetVoicemailQueueMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages unauthorized response has a 4xx status code
func (o *GetVoicemailQueueMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages unauthorized response has a 5xx status code
func (o *GetVoicemailQueueMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages unauthorized response a status code equal to that given
func (o *GetVoicemailQueueMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVoicemailQueueMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailQueueMessagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailQueueMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesForbidden creates a GetVoicemailQueueMessagesForbidden with default headers values
func NewGetVoicemailQueueMessagesForbidden() *GetVoicemailQueueMessagesForbidden {
	return &GetVoicemailQueueMessagesForbidden{}
}

/*
GetVoicemailQueueMessagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailQueueMessagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages forbidden response has a 2xx status code
func (o *GetVoicemailQueueMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages forbidden response has a 3xx status code
func (o *GetVoicemailQueueMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages forbidden response has a 4xx status code
func (o *GetVoicemailQueueMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages forbidden response has a 5xx status code
func (o *GetVoicemailQueueMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages forbidden response a status code equal to that given
func (o *GetVoicemailQueueMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetVoicemailQueueMessagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailQueueMessagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailQueueMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesNotFound creates a GetVoicemailQueueMessagesNotFound with default headers values
func NewGetVoicemailQueueMessagesNotFound() *GetVoicemailQueueMessagesNotFound {
	return &GetVoicemailQueueMessagesNotFound{}
}

/*
GetVoicemailQueueMessagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetVoicemailQueueMessagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages not found response has a 2xx status code
func (o *GetVoicemailQueueMessagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages not found response has a 3xx status code
func (o *GetVoicemailQueueMessagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages not found response has a 4xx status code
func (o *GetVoicemailQueueMessagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages not found response has a 5xx status code
func (o *GetVoicemailQueueMessagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages not found response a status code equal to that given
func (o *GetVoicemailQueueMessagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVoicemailQueueMessagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailQueueMessagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailQueueMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesRequestTimeout creates a GetVoicemailQueueMessagesRequestTimeout with default headers values
func NewGetVoicemailQueueMessagesRequestTimeout() *GetVoicemailQueueMessagesRequestTimeout {
	return &GetVoicemailQueueMessagesRequestTimeout{}
}

/*
GetVoicemailQueueMessagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetVoicemailQueueMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages request timeout response has a 2xx status code
func (o *GetVoicemailQueueMessagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages request timeout response has a 3xx status code
func (o *GetVoicemailQueueMessagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages request timeout response has a 4xx status code
func (o *GetVoicemailQueueMessagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages request timeout response has a 5xx status code
func (o *GetVoicemailQueueMessagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages request timeout response a status code equal to that given
func (o *GetVoicemailQueueMessagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetVoicemailQueueMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailQueueMessagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailQueueMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesRequestEntityTooLarge creates a GetVoicemailQueueMessagesRequestEntityTooLarge with default headers values
func NewGetVoicemailQueueMessagesRequestEntityTooLarge() *GetVoicemailQueueMessagesRequestEntityTooLarge {
	return &GetVoicemailQueueMessagesRequestEntityTooLarge{}
}

/*
GetVoicemailQueueMessagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetVoicemailQueueMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages request entity too large response has a 2xx status code
func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages request entity too large response has a 3xx status code
func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages request entity too large response has a 4xx status code
func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages request entity too large response has a 5xx status code
func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages request entity too large response a status code equal to that given
func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesUnsupportedMediaType creates a GetVoicemailQueueMessagesUnsupportedMediaType with default headers values
func NewGetVoicemailQueueMessagesUnsupportedMediaType() *GetVoicemailQueueMessagesUnsupportedMediaType {
	return &GetVoicemailQueueMessagesUnsupportedMediaType{}
}

/*
GetVoicemailQueueMessagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailQueueMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages unsupported media type response has a 2xx status code
func (o *GetVoicemailQueueMessagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages unsupported media type response has a 3xx status code
func (o *GetVoicemailQueueMessagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages unsupported media type response has a 4xx status code
func (o *GetVoicemailQueueMessagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages unsupported media type response has a 5xx status code
func (o *GetVoicemailQueueMessagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages unsupported media type response a status code equal to that given
func (o *GetVoicemailQueueMessagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesTooManyRequests creates a GetVoicemailQueueMessagesTooManyRequests with default headers values
func NewGetVoicemailQueueMessagesTooManyRequests() *GetVoicemailQueueMessagesTooManyRequests {
	return &GetVoicemailQueueMessagesTooManyRequests{}
}

/*
GetVoicemailQueueMessagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetVoicemailQueueMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages too many requests response has a 2xx status code
func (o *GetVoicemailQueueMessagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages too many requests response has a 3xx status code
func (o *GetVoicemailQueueMessagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages too many requests response has a 4xx status code
func (o *GetVoicemailQueueMessagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail queue messages too many requests response has a 5xx status code
func (o *GetVoicemailQueueMessagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail queue messages too many requests response a status code equal to that given
func (o *GetVoicemailQueueMessagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetVoicemailQueueMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailQueueMessagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailQueueMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesInternalServerError creates a GetVoicemailQueueMessagesInternalServerError with default headers values
func NewGetVoicemailQueueMessagesInternalServerError() *GetVoicemailQueueMessagesInternalServerError {
	return &GetVoicemailQueueMessagesInternalServerError{}
}

/*
GetVoicemailQueueMessagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailQueueMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages internal server error response has a 2xx status code
func (o *GetVoicemailQueueMessagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages internal server error response has a 3xx status code
func (o *GetVoicemailQueueMessagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages internal server error response has a 4xx status code
func (o *GetVoicemailQueueMessagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail queue messages internal server error response has a 5xx status code
func (o *GetVoicemailQueueMessagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail queue messages internal server error response a status code equal to that given
func (o *GetVoicemailQueueMessagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVoicemailQueueMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailQueueMessagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailQueueMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesServiceUnavailable creates a GetVoicemailQueueMessagesServiceUnavailable with default headers values
func NewGetVoicemailQueueMessagesServiceUnavailable() *GetVoicemailQueueMessagesServiceUnavailable {
	return &GetVoicemailQueueMessagesServiceUnavailable{}
}

/*
GetVoicemailQueueMessagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailQueueMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages service unavailable response has a 2xx status code
func (o *GetVoicemailQueueMessagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages service unavailable response has a 3xx status code
func (o *GetVoicemailQueueMessagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages service unavailable response has a 4xx status code
func (o *GetVoicemailQueueMessagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail queue messages service unavailable response has a 5xx status code
func (o *GetVoicemailQueueMessagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail queue messages service unavailable response a status code equal to that given
func (o *GetVoicemailQueueMessagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailQueueMessagesGatewayTimeout creates a GetVoicemailQueueMessagesGatewayTimeout with default headers values
func NewGetVoicemailQueueMessagesGatewayTimeout() *GetVoicemailQueueMessagesGatewayTimeout {
	return &GetVoicemailQueueMessagesGatewayTimeout{}
}

/*
GetVoicemailQueueMessagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetVoicemailQueueMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail queue messages gateway timeout response has a 2xx status code
func (o *GetVoicemailQueueMessagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail queue messages gateway timeout response has a 3xx status code
func (o *GetVoicemailQueueMessagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail queue messages gateway timeout response has a 4xx status code
func (o *GetVoicemailQueueMessagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail queue messages gateway timeout response has a 5xx status code
func (o *GetVoicemailQueueMessagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail queue messages gateway timeout response a status code equal to that given
func (o *GetVoicemailQueueMessagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/queues/{queueId}/messages][%d] getVoicemailQueueMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailQueueMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
