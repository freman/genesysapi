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

// GetVoicemailMeMailboxReader is a Reader for the GetVoicemailMeMailbox structure.
type GetVoicemailMeMailboxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailMeMailboxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailMeMailboxOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailMeMailboxBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailMeMailboxUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailMeMailboxForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailMeMailboxNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetVoicemailMeMailboxRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailMeMailboxRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailMeMailboxUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailMeMailboxTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailMeMailboxInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailMeMailboxServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailMeMailboxGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailMeMailboxOK creates a GetVoicemailMeMailboxOK with default headers values
func NewGetVoicemailMeMailboxOK() *GetVoicemailMeMailboxOK {
	return &GetVoicemailMeMailboxOK{}
}

/*
GetVoicemailMeMailboxOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVoicemailMeMailboxOK struct {
	Payload *models.VoicemailMailboxInfo
}

// IsSuccess returns true when this get voicemail me mailbox o k response has a 2xx status code
func (o *GetVoicemailMeMailboxOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get voicemail me mailbox o k response has a 3xx status code
func (o *GetVoicemailMeMailboxOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox o k response has a 4xx status code
func (o *GetVoicemailMeMailboxOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me mailbox o k response has a 5xx status code
func (o *GetVoicemailMeMailboxOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox o k response a status code equal to that given
func (o *GetVoicemailMeMailboxOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVoicemailMeMailboxOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailMeMailboxOK) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailMeMailboxOK) GetPayload() *models.VoicemailMailboxInfo {
	return o.Payload
}

func (o *GetVoicemailMeMailboxOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailMailboxInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxBadRequest creates a GetVoicemailMeMailboxBadRequest with default headers values
func NewGetVoicemailMeMailboxBadRequest() *GetVoicemailMeMailboxBadRequest {
	return &GetVoicemailMeMailboxBadRequest{}
}

/*
GetVoicemailMeMailboxBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailMeMailboxBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox bad request response has a 2xx status code
func (o *GetVoicemailMeMailboxBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox bad request response has a 3xx status code
func (o *GetVoicemailMeMailboxBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox bad request response has a 4xx status code
func (o *GetVoicemailMeMailboxBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox bad request response has a 5xx status code
func (o *GetVoicemailMeMailboxBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox bad request response a status code equal to that given
func (o *GetVoicemailMeMailboxBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVoicemailMeMailboxBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailMeMailboxBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailMeMailboxBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxUnauthorized creates a GetVoicemailMeMailboxUnauthorized with default headers values
func NewGetVoicemailMeMailboxUnauthorized() *GetVoicemailMeMailboxUnauthorized {
	return &GetVoicemailMeMailboxUnauthorized{}
}

/*
GetVoicemailMeMailboxUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailMeMailboxUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox unauthorized response has a 2xx status code
func (o *GetVoicemailMeMailboxUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox unauthorized response has a 3xx status code
func (o *GetVoicemailMeMailboxUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox unauthorized response has a 4xx status code
func (o *GetVoicemailMeMailboxUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox unauthorized response has a 5xx status code
func (o *GetVoicemailMeMailboxUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox unauthorized response a status code equal to that given
func (o *GetVoicemailMeMailboxUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVoicemailMeMailboxUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailMeMailboxUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailMeMailboxUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxForbidden creates a GetVoicemailMeMailboxForbidden with default headers values
func NewGetVoicemailMeMailboxForbidden() *GetVoicemailMeMailboxForbidden {
	return &GetVoicemailMeMailboxForbidden{}
}

/*
GetVoicemailMeMailboxForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailMeMailboxForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox forbidden response has a 2xx status code
func (o *GetVoicemailMeMailboxForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox forbidden response has a 3xx status code
func (o *GetVoicemailMeMailboxForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox forbidden response has a 4xx status code
func (o *GetVoicemailMeMailboxForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox forbidden response has a 5xx status code
func (o *GetVoicemailMeMailboxForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox forbidden response a status code equal to that given
func (o *GetVoicemailMeMailboxForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetVoicemailMeMailboxForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailMeMailboxForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailMeMailboxForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxNotFound creates a GetVoicemailMeMailboxNotFound with default headers values
func NewGetVoicemailMeMailboxNotFound() *GetVoicemailMeMailboxNotFound {
	return &GetVoicemailMeMailboxNotFound{}
}

/*
GetVoicemailMeMailboxNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetVoicemailMeMailboxNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox not found response has a 2xx status code
func (o *GetVoicemailMeMailboxNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox not found response has a 3xx status code
func (o *GetVoicemailMeMailboxNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox not found response has a 4xx status code
func (o *GetVoicemailMeMailboxNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox not found response has a 5xx status code
func (o *GetVoicemailMeMailboxNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox not found response a status code equal to that given
func (o *GetVoicemailMeMailboxNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVoicemailMeMailboxNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailMeMailboxNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailMeMailboxNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxRequestTimeout creates a GetVoicemailMeMailboxRequestTimeout with default headers values
func NewGetVoicemailMeMailboxRequestTimeout() *GetVoicemailMeMailboxRequestTimeout {
	return &GetVoicemailMeMailboxRequestTimeout{}
}

/*
GetVoicemailMeMailboxRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetVoicemailMeMailboxRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox request timeout response has a 2xx status code
func (o *GetVoicemailMeMailboxRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox request timeout response has a 3xx status code
func (o *GetVoicemailMeMailboxRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox request timeout response has a 4xx status code
func (o *GetVoicemailMeMailboxRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox request timeout response has a 5xx status code
func (o *GetVoicemailMeMailboxRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox request timeout response a status code equal to that given
func (o *GetVoicemailMeMailboxRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetVoicemailMeMailboxRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailMeMailboxRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailMeMailboxRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxRequestEntityTooLarge creates a GetVoicemailMeMailboxRequestEntityTooLarge with default headers values
func NewGetVoicemailMeMailboxRequestEntityTooLarge() *GetVoicemailMeMailboxRequestEntityTooLarge {
	return &GetVoicemailMeMailboxRequestEntityTooLarge{}
}

/*
GetVoicemailMeMailboxRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetVoicemailMeMailboxRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox request entity too large response has a 2xx status code
func (o *GetVoicemailMeMailboxRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox request entity too large response has a 3xx status code
func (o *GetVoicemailMeMailboxRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox request entity too large response has a 4xx status code
func (o *GetVoicemailMeMailboxRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox request entity too large response has a 5xx status code
func (o *GetVoicemailMeMailboxRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox request entity too large response a status code equal to that given
func (o *GetVoicemailMeMailboxRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetVoicemailMeMailboxRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailMeMailboxRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailMeMailboxRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxUnsupportedMediaType creates a GetVoicemailMeMailboxUnsupportedMediaType with default headers values
func NewGetVoicemailMeMailboxUnsupportedMediaType() *GetVoicemailMeMailboxUnsupportedMediaType {
	return &GetVoicemailMeMailboxUnsupportedMediaType{}
}

/*
GetVoicemailMeMailboxUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailMeMailboxUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox unsupported media type response has a 2xx status code
func (o *GetVoicemailMeMailboxUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox unsupported media type response has a 3xx status code
func (o *GetVoicemailMeMailboxUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox unsupported media type response has a 4xx status code
func (o *GetVoicemailMeMailboxUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox unsupported media type response has a 5xx status code
func (o *GetVoicemailMeMailboxUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox unsupported media type response a status code equal to that given
func (o *GetVoicemailMeMailboxUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetVoicemailMeMailboxUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailMeMailboxUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailMeMailboxUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxTooManyRequests creates a GetVoicemailMeMailboxTooManyRequests with default headers values
func NewGetVoicemailMeMailboxTooManyRequests() *GetVoicemailMeMailboxTooManyRequests {
	return &GetVoicemailMeMailboxTooManyRequests{}
}

/*
GetVoicemailMeMailboxTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetVoicemailMeMailboxTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox too many requests response has a 2xx status code
func (o *GetVoicemailMeMailboxTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox too many requests response has a 3xx status code
func (o *GetVoicemailMeMailboxTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox too many requests response has a 4xx status code
func (o *GetVoicemailMeMailboxTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me mailbox too many requests response has a 5xx status code
func (o *GetVoicemailMeMailboxTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me mailbox too many requests response a status code equal to that given
func (o *GetVoicemailMeMailboxTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetVoicemailMeMailboxTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailMeMailboxTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailMeMailboxTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxInternalServerError creates a GetVoicemailMeMailboxInternalServerError with default headers values
func NewGetVoicemailMeMailboxInternalServerError() *GetVoicemailMeMailboxInternalServerError {
	return &GetVoicemailMeMailboxInternalServerError{}
}

/*
GetVoicemailMeMailboxInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailMeMailboxInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox internal server error response has a 2xx status code
func (o *GetVoicemailMeMailboxInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox internal server error response has a 3xx status code
func (o *GetVoicemailMeMailboxInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox internal server error response has a 4xx status code
func (o *GetVoicemailMeMailboxInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me mailbox internal server error response has a 5xx status code
func (o *GetVoicemailMeMailboxInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail me mailbox internal server error response a status code equal to that given
func (o *GetVoicemailMeMailboxInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVoicemailMeMailboxInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailMeMailboxInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailMeMailboxInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxServiceUnavailable creates a GetVoicemailMeMailboxServiceUnavailable with default headers values
func NewGetVoicemailMeMailboxServiceUnavailable() *GetVoicemailMeMailboxServiceUnavailable {
	return &GetVoicemailMeMailboxServiceUnavailable{}
}

/*
GetVoicemailMeMailboxServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailMeMailboxServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox service unavailable response has a 2xx status code
func (o *GetVoicemailMeMailboxServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox service unavailable response has a 3xx status code
func (o *GetVoicemailMeMailboxServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox service unavailable response has a 4xx status code
func (o *GetVoicemailMeMailboxServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me mailbox service unavailable response has a 5xx status code
func (o *GetVoicemailMeMailboxServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail me mailbox service unavailable response a status code equal to that given
func (o *GetVoicemailMeMailboxServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetVoicemailMeMailboxServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailMeMailboxServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailMeMailboxServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMeMailboxGatewayTimeout creates a GetVoicemailMeMailboxGatewayTimeout with default headers values
func NewGetVoicemailMeMailboxGatewayTimeout() *GetVoicemailMeMailboxGatewayTimeout {
	return &GetVoicemailMeMailboxGatewayTimeout{}
}

/*
GetVoicemailMeMailboxGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetVoicemailMeMailboxGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me mailbox gateway timeout response has a 2xx status code
func (o *GetVoicemailMeMailboxGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me mailbox gateway timeout response has a 3xx status code
func (o *GetVoicemailMeMailboxGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me mailbox gateway timeout response has a 4xx status code
func (o *GetVoicemailMeMailboxGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me mailbox gateway timeout response has a 5xx status code
func (o *GetVoicemailMeMailboxGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail me mailbox gateway timeout response a status code equal to that given
func (o *GetVoicemailMeMailboxGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetVoicemailMeMailboxGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailMeMailboxGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/mailbox][%d] getVoicemailMeMailboxGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailMeMailboxGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMeMailboxGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
