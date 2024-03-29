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

// GetVoicemailMePolicyReader is a Reader for the GetVoicemailMePolicy structure.
type GetVoicemailMePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoicemailMePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoicemailMePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVoicemailMePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetVoicemailMePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVoicemailMePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVoicemailMePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetVoicemailMePolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetVoicemailMePolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetVoicemailMePolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetVoicemailMePolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVoicemailMePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetVoicemailMePolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetVoicemailMePolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVoicemailMePolicyOK creates a GetVoicemailMePolicyOK with default headers values
func NewGetVoicemailMePolicyOK() *GetVoicemailMePolicyOK {
	return &GetVoicemailMePolicyOK{}
}

/*
GetVoicemailMePolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVoicemailMePolicyOK struct {
	Payload *models.VoicemailUserPolicy
}

// IsSuccess returns true when this get voicemail me policy o k response has a 2xx status code
func (o *GetVoicemailMePolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get voicemail me policy o k response has a 3xx status code
func (o *GetVoicemailMePolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy o k response has a 4xx status code
func (o *GetVoicemailMePolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me policy o k response has a 5xx status code
func (o *GetVoicemailMePolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy o k response a status code equal to that given
func (o *GetVoicemailMePolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVoicemailMePolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailMePolicyOK) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyOK  %+v", 200, o.Payload)
}

func (o *GetVoicemailMePolicyOK) GetPayload() *models.VoicemailUserPolicy {
	return o.Payload
}

func (o *GetVoicemailMePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailUserPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyBadRequest creates a GetVoicemailMePolicyBadRequest with default headers values
func NewGetVoicemailMePolicyBadRequest() *GetVoicemailMePolicyBadRequest {
	return &GetVoicemailMePolicyBadRequest{}
}

/*
GetVoicemailMePolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetVoicemailMePolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy bad request response has a 2xx status code
func (o *GetVoicemailMePolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy bad request response has a 3xx status code
func (o *GetVoicemailMePolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy bad request response has a 4xx status code
func (o *GetVoicemailMePolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy bad request response has a 5xx status code
func (o *GetVoicemailMePolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy bad request response a status code equal to that given
func (o *GetVoicemailMePolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetVoicemailMePolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailMePolicyBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetVoicemailMePolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyUnauthorized creates a GetVoicemailMePolicyUnauthorized with default headers values
func NewGetVoicemailMePolicyUnauthorized() *GetVoicemailMePolicyUnauthorized {
	return &GetVoicemailMePolicyUnauthorized{}
}

/*
GetVoicemailMePolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetVoicemailMePolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy unauthorized response has a 2xx status code
func (o *GetVoicemailMePolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy unauthorized response has a 3xx status code
func (o *GetVoicemailMePolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy unauthorized response has a 4xx status code
func (o *GetVoicemailMePolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy unauthorized response has a 5xx status code
func (o *GetVoicemailMePolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy unauthorized response a status code equal to that given
func (o *GetVoicemailMePolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVoicemailMePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailMePolicyUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVoicemailMePolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyForbidden creates a GetVoicemailMePolicyForbidden with default headers values
func NewGetVoicemailMePolicyForbidden() *GetVoicemailMePolicyForbidden {
	return &GetVoicemailMePolicyForbidden{}
}

/*
GetVoicemailMePolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetVoicemailMePolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy forbidden response has a 2xx status code
func (o *GetVoicemailMePolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy forbidden response has a 3xx status code
func (o *GetVoicemailMePolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy forbidden response has a 4xx status code
func (o *GetVoicemailMePolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy forbidden response has a 5xx status code
func (o *GetVoicemailMePolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy forbidden response a status code equal to that given
func (o *GetVoicemailMePolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetVoicemailMePolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailMePolicyForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetVoicemailMePolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyNotFound creates a GetVoicemailMePolicyNotFound with default headers values
func NewGetVoicemailMePolicyNotFound() *GetVoicemailMePolicyNotFound {
	return &GetVoicemailMePolicyNotFound{}
}

/*
GetVoicemailMePolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetVoicemailMePolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy not found response has a 2xx status code
func (o *GetVoicemailMePolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy not found response has a 3xx status code
func (o *GetVoicemailMePolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy not found response has a 4xx status code
func (o *GetVoicemailMePolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy not found response has a 5xx status code
func (o *GetVoicemailMePolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy not found response a status code equal to that given
func (o *GetVoicemailMePolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVoicemailMePolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailMePolicyNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetVoicemailMePolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyRequestTimeout creates a GetVoicemailMePolicyRequestTimeout with default headers values
func NewGetVoicemailMePolicyRequestTimeout() *GetVoicemailMePolicyRequestTimeout {
	return &GetVoicemailMePolicyRequestTimeout{}
}

/*
GetVoicemailMePolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetVoicemailMePolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy request timeout response has a 2xx status code
func (o *GetVoicemailMePolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy request timeout response has a 3xx status code
func (o *GetVoicemailMePolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy request timeout response has a 4xx status code
func (o *GetVoicemailMePolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy request timeout response has a 5xx status code
func (o *GetVoicemailMePolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy request timeout response a status code equal to that given
func (o *GetVoicemailMePolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetVoicemailMePolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailMePolicyRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetVoicemailMePolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyRequestEntityTooLarge creates a GetVoicemailMePolicyRequestEntityTooLarge with default headers values
func NewGetVoicemailMePolicyRequestEntityTooLarge() *GetVoicemailMePolicyRequestEntityTooLarge {
	return &GetVoicemailMePolicyRequestEntityTooLarge{}
}

/*
GetVoicemailMePolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetVoicemailMePolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy request entity too large response has a 2xx status code
func (o *GetVoicemailMePolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy request entity too large response has a 3xx status code
func (o *GetVoicemailMePolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy request entity too large response has a 4xx status code
func (o *GetVoicemailMePolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy request entity too large response has a 5xx status code
func (o *GetVoicemailMePolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy request entity too large response a status code equal to that given
func (o *GetVoicemailMePolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetVoicemailMePolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailMePolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetVoicemailMePolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyUnsupportedMediaType creates a GetVoicemailMePolicyUnsupportedMediaType with default headers values
func NewGetVoicemailMePolicyUnsupportedMediaType() *GetVoicemailMePolicyUnsupportedMediaType {
	return &GetVoicemailMePolicyUnsupportedMediaType{}
}

/*
GetVoicemailMePolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetVoicemailMePolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy unsupported media type response has a 2xx status code
func (o *GetVoicemailMePolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy unsupported media type response has a 3xx status code
func (o *GetVoicemailMePolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy unsupported media type response has a 4xx status code
func (o *GetVoicemailMePolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy unsupported media type response has a 5xx status code
func (o *GetVoicemailMePolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy unsupported media type response a status code equal to that given
func (o *GetVoicemailMePolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetVoicemailMePolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailMePolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetVoicemailMePolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyTooManyRequests creates a GetVoicemailMePolicyTooManyRequests with default headers values
func NewGetVoicemailMePolicyTooManyRequests() *GetVoicemailMePolicyTooManyRequests {
	return &GetVoicemailMePolicyTooManyRequests{}
}

/*
GetVoicemailMePolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetVoicemailMePolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy too many requests response has a 2xx status code
func (o *GetVoicemailMePolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy too many requests response has a 3xx status code
func (o *GetVoicemailMePolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy too many requests response has a 4xx status code
func (o *GetVoicemailMePolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get voicemail me policy too many requests response has a 5xx status code
func (o *GetVoicemailMePolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get voicemail me policy too many requests response a status code equal to that given
func (o *GetVoicemailMePolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetVoicemailMePolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailMePolicyTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetVoicemailMePolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyInternalServerError creates a GetVoicemailMePolicyInternalServerError with default headers values
func NewGetVoicemailMePolicyInternalServerError() *GetVoicemailMePolicyInternalServerError {
	return &GetVoicemailMePolicyInternalServerError{}
}

/*
GetVoicemailMePolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetVoicemailMePolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy internal server error response has a 2xx status code
func (o *GetVoicemailMePolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy internal server error response has a 3xx status code
func (o *GetVoicemailMePolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy internal server error response has a 4xx status code
func (o *GetVoicemailMePolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me policy internal server error response has a 5xx status code
func (o *GetVoicemailMePolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail me policy internal server error response a status code equal to that given
func (o *GetVoicemailMePolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVoicemailMePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailMePolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVoicemailMePolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyServiceUnavailable creates a GetVoicemailMePolicyServiceUnavailable with default headers values
func NewGetVoicemailMePolicyServiceUnavailable() *GetVoicemailMePolicyServiceUnavailable {
	return &GetVoicemailMePolicyServiceUnavailable{}
}

/*
GetVoicemailMePolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetVoicemailMePolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy service unavailable response has a 2xx status code
func (o *GetVoicemailMePolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy service unavailable response has a 3xx status code
func (o *GetVoicemailMePolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy service unavailable response has a 4xx status code
func (o *GetVoicemailMePolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me policy service unavailable response has a 5xx status code
func (o *GetVoicemailMePolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail me policy service unavailable response a status code equal to that given
func (o *GetVoicemailMePolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetVoicemailMePolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailMePolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetVoicemailMePolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVoicemailMePolicyGatewayTimeout creates a GetVoicemailMePolicyGatewayTimeout with default headers values
func NewGetVoicemailMePolicyGatewayTimeout() *GetVoicemailMePolicyGatewayTimeout {
	return &GetVoicemailMePolicyGatewayTimeout{}
}

/*
GetVoicemailMePolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetVoicemailMePolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get voicemail me policy gateway timeout response has a 2xx status code
func (o *GetVoicemailMePolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get voicemail me policy gateway timeout response has a 3xx status code
func (o *GetVoicemailMePolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get voicemail me policy gateway timeout response has a 4xx status code
func (o *GetVoicemailMePolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get voicemail me policy gateway timeout response has a 5xx status code
func (o *GetVoicemailMePolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get voicemail me policy gateway timeout response a status code equal to that given
func (o *GetVoicemailMePolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetVoicemailMePolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailMePolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/voicemail/me/policy][%d] getVoicemailMePolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetVoicemailMePolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetVoicemailMePolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
