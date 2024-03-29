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

// GetVoicemailGroupMailboxReader is a Reader for the GetVoicemailGroupMailbox structure.
type GetVoicemailGroupMailboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailGroupMailboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailGroupMailboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailGroupMailboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailGroupMailboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailGroupMailboxForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailGroupMailboxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetVoicemailGroupMailboxRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailGroupMailboxRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailGroupMailboxUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailGroupMailboxTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailGroupMailboxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailGroupMailboxServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailGroupMailboxGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailGroupMailboxOK creates a GetVoicemailGroupMailboxOK with default headers values
func NewGetVoicemailGroupMailboxOK() *GetVoicemailGroupMailboxOK {
	return &GetVoicemailGroupMailboxOK{}
}

/*
GetVoicemailGroupMailboxOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVoicemailGroupMailboxOK struct {
	Payload *models.VoicemailMailboxInfo
}

// IsSuccess returns true when this get voicemail group mailbox o k response has a 2xx status code
func (o *GetVoicemailGroupMailboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get voicemail group mailbox o k response has a 3xx status code
func (o *GetVoicemailGroupMailboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox o k response has a 4xx status code
func (o *GetVoicemailGroupMailboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail group mailbox o k response has a 5xx status code
func (o *GetVoicemailGroupMailboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox o k response a status code equal to that given
func (o *GetVoicemailGroupMailboxOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVoicemailGroupMailboxOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailGroupMailboxOK) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailGroupMailboxOK) GetPayload() *models.VoicemailMailboxInfo {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMailboxInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxBadRequest creates a GetVoicemailGroupMailboxBadRequest with default headers values
func NewGetVoicemailGroupMailboxBadRequest() *GetVoicemailGroupMailboxBadRequest {
	return &GetVoicemailGroupMailboxBadRequest{}
}

/*
GetVoicemailGroupMailboxBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailGroupMailboxBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox bad request response has a 2xx status code
func (o *GetVoicemailGroupMailboxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox bad request response has a 3xx status code
func (o *GetVoicemailGroupMailboxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox bad request response has a 4xx status code
func (o *GetVoicemailGroupMailboxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox bad request response has a 5xx status code
func (o *GetVoicemailGroupMailboxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox bad request response a status code equal to that given
func (o *GetVoicemailGroupMailboxBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVoicemailGroupMailboxBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailGroupMailboxBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailGroupMailboxBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxUnauthorized creates a GetVoicemailGroupMailboxUnauthorized with default headers values
func NewGetVoicemailGroupMailboxUnauthorized() *GetVoicemailGroupMailboxUnauthorized {
	return &GetVoicemailGroupMailboxUnauthorized{}
}

/*
GetVoicemailGroupMailboxUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailGroupMailboxUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox unauthorized response has a 2xx status code
func (o *GetVoicemailGroupMailboxUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox unauthorized response has a 3xx status code
func (o *GetVoicemailGroupMailboxUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox unauthorized response has a 4xx status code
func (o *GetVoicemailGroupMailboxUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox unauthorized response has a 5xx status code
func (o *GetVoicemailGroupMailboxUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox unauthorized response a status code equal to that given
func (o *GetVoicemailGroupMailboxUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVoicemailGroupMailboxUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailGroupMailboxUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailGroupMailboxUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxForbidden creates a GetVoicemailGroupMailboxForbidden with default headers values
func NewGetVoicemailGroupMailboxForbidden() *GetVoicemailGroupMailboxForbidden {
	return &GetVoicemailGroupMailboxForbidden{}
}

/*
GetVoicemailGroupMailboxForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailGroupMailboxForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox forbidden response has a 2xx status code
func (o *GetVoicemailGroupMailboxForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox forbidden response has a 3xx status code
func (o *GetVoicemailGroupMailboxForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox forbidden response has a 4xx status code
func (o *GetVoicemailGroupMailboxForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox forbidden response has a 5xx status code
func (o *GetVoicemailGroupMailboxForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox forbidden response a status code equal to that given
func (o *GetVoicemailGroupMailboxForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetVoicemailGroupMailboxForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailGroupMailboxForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailGroupMailboxForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxNotFound creates a GetVoicemailGroupMailboxNotFound with default headers values
func NewGetVoicemailGroupMailboxNotFound() *GetVoicemailGroupMailboxNotFound {
	return &GetVoicemailGroupMailboxNotFound{}
}

/*
GetVoicemailGroupMailboxNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetVoicemailGroupMailboxNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox not found response has a 2xx status code
func (o *GetVoicemailGroupMailboxNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox not found response has a 3xx status code
func (o *GetVoicemailGroupMailboxNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox not found response has a 4xx status code
func (o *GetVoicemailGroupMailboxNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox not found response has a 5xx status code
func (o *GetVoicemailGroupMailboxNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox not found response a status code equal to that given
func (o *GetVoicemailGroupMailboxNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVoicemailGroupMailboxNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailGroupMailboxNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailGroupMailboxNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxRequestTimeout creates a GetVoicemailGroupMailboxRequestTimeout with default headers values
func NewGetVoicemailGroupMailboxRequestTimeout() *GetVoicemailGroupMailboxRequestTimeout {
	return &GetVoicemailGroupMailboxRequestTimeout{}
}

/*
GetVoicemailGroupMailboxRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetVoicemailGroupMailboxRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox request timeout response has a 2xx status code
func (o *GetVoicemailGroupMailboxRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox request timeout response has a 3xx status code
func (o *GetVoicemailGroupMailboxRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox request timeout response has a 4xx status code
func (o *GetVoicemailGroupMailboxRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox request timeout response has a 5xx status code
func (o *GetVoicemailGroupMailboxRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox request timeout response a status code equal to that given
func (o *GetVoicemailGroupMailboxRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetVoicemailGroupMailboxRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailGroupMailboxRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailGroupMailboxRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxRequestEntityTooLarge creates a GetVoicemailGroupMailboxRequestEntityTooLarge with default headers values
func NewGetVoicemailGroupMailboxRequestEntityTooLarge() *GetVoicemailGroupMailboxRequestEntityTooLarge {
	return &GetVoicemailGroupMailboxRequestEntityTooLarge{}
}

/*
GetVoicemailGroupMailboxRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetVoicemailGroupMailboxRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox request entity too large response has a 2xx status code
func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox request entity too large response has a 3xx status code
func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox request entity too large response has a 4xx status code
func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox request entity too large response has a 5xx status code
func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox request entity too large response a status code equal to that given
func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxUnsupportedMediaType creates a GetVoicemailGroupMailboxUnsupportedMediaType with default headers values
func NewGetVoicemailGroupMailboxUnsupportedMediaType() *GetVoicemailGroupMailboxUnsupportedMediaType {
	return &GetVoicemailGroupMailboxUnsupportedMediaType{}
}

/*
GetVoicemailGroupMailboxUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailGroupMailboxUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox unsupported media type response has a 2xx status code
func (o *GetVoicemailGroupMailboxUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox unsupported media type response has a 3xx status code
func (o *GetVoicemailGroupMailboxUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox unsupported media type response has a 4xx status code
func (o *GetVoicemailGroupMailboxUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox unsupported media type response has a 5xx status code
func (o *GetVoicemailGroupMailboxUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox unsupported media type response a status code equal to that given
func (o *GetVoicemailGroupMailboxUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxTooManyRequests creates a GetVoicemailGroupMailboxTooManyRequests with default headers values
func NewGetVoicemailGroupMailboxTooManyRequests() *GetVoicemailGroupMailboxTooManyRequests {
	return &GetVoicemailGroupMailboxTooManyRequests{}
}

/*
GetVoicemailGroupMailboxTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetVoicemailGroupMailboxTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox too many requests response has a 2xx status code
func (o *GetVoicemailGroupMailboxTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox too many requests response has a 3xx status code
func (o *GetVoicemailGroupMailboxTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox too many requests response has a 4xx status code
func (o *GetVoicemailGroupMailboxTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail group mailbox too many requests response has a 5xx status code
func (o *GetVoicemailGroupMailboxTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail group mailbox too many requests response a status code equal to that given
func (o *GetVoicemailGroupMailboxTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetVoicemailGroupMailboxTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailGroupMailboxTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailGroupMailboxTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxInternalServerError creates a GetVoicemailGroupMailboxInternalServerError with default headers values
func NewGetVoicemailGroupMailboxInternalServerError() *GetVoicemailGroupMailboxInternalServerError {
	return &GetVoicemailGroupMailboxInternalServerError{}
}

/*
GetVoicemailGroupMailboxInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailGroupMailboxInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox internal server error response has a 2xx status code
func (o *GetVoicemailGroupMailboxInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox internal server error response has a 3xx status code
func (o *GetVoicemailGroupMailboxInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox internal server error response has a 4xx status code
func (o *GetVoicemailGroupMailboxInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail group mailbox internal server error response has a 5xx status code
func (o *GetVoicemailGroupMailboxInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail group mailbox internal server error response a status code equal to that given
func (o *GetVoicemailGroupMailboxInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVoicemailGroupMailboxInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailGroupMailboxInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailGroupMailboxInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxServiceUnavailable creates a GetVoicemailGroupMailboxServiceUnavailable with default headers values
func NewGetVoicemailGroupMailboxServiceUnavailable() *GetVoicemailGroupMailboxServiceUnavailable {
	return &GetVoicemailGroupMailboxServiceUnavailable{}
}

/*
GetVoicemailGroupMailboxServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailGroupMailboxServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox service unavailable response has a 2xx status code
func (o *GetVoicemailGroupMailboxServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox service unavailable response has a 3xx status code
func (o *GetVoicemailGroupMailboxServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox service unavailable response has a 4xx status code
func (o *GetVoicemailGroupMailboxServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail group mailbox service unavailable response has a 5xx status code
func (o *GetVoicemailGroupMailboxServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail group mailbox service unavailable response a status code equal to that given
func (o *GetVoicemailGroupMailboxServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailGroupMailboxGatewayTimeout creates a GetVoicemailGroupMailboxGatewayTimeout with default headers values
func NewGetVoicemailGroupMailboxGatewayTimeout() *GetVoicemailGroupMailboxGatewayTimeout {
	return &GetVoicemailGroupMailboxGatewayTimeout{}
}

/*
GetVoicemailGroupMailboxGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetVoicemailGroupMailboxGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail group mailbox gateway timeout response has a 2xx status code
func (o *GetVoicemailGroupMailboxGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail group mailbox gateway timeout response has a 3xx status code
func (o *GetVoicemailGroupMailboxGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail group mailbox gateway timeout response has a 4xx status code
func (o *GetVoicemailGroupMailboxGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail group mailbox gateway timeout response has a 5xx status code
func (o *GetVoicemailGroupMailboxGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail group mailbox gateway timeout response a status code equal to that given
func (o *GetVoicemailGroupMailboxGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/groups/{groupId}/mailbox][%d] getVoicemailGroupMailboxGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailGroupMailboxGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
