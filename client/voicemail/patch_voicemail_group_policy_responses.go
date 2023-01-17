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

// PatchVoicemailGroupPolicyReader is a Reader for the PatchVoicemailGroupPolicy structure.
type PatchVoicemailGroupPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchVoicemailGroupPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchVoicemailGroupPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchVoicemailGroupPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchVoicemailGroupPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchVoicemailGroupPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchVoicemailGroupPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchVoicemailGroupPolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchVoicemailGroupPolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchVoicemailGroupPolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchVoicemailGroupPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchVoicemailGroupPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchVoicemailGroupPolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchVoicemailGroupPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchVoicemailGroupPolicyOK creates a PatchVoicemailGroupPolicyOK with default headers values
func NewPatchVoicemailGroupPolicyOK() *PatchVoicemailGroupPolicyOK {
	return &PatchVoicemailGroupPolicyOK{}
}

/*
PatchVoicemailGroupPolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchVoicemailGroupPolicyOK struct {
	Payload *models.VoicemailGroupPolicy
}

// IsSuccess returns true when this patch voicemail group policy o k response has a 2xx status code
func (o *PatchVoicemailGroupPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch voicemail group policy o k response has a 3xx status code
func (o *PatchVoicemailGroupPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy o k response has a 4xx status code
func (o *PatchVoicemailGroupPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail group policy o k response has a 5xx status code
func (o *PatchVoicemailGroupPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy o k response a status code equal to that given
func (o *PatchVoicemailGroupPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchVoicemailGroupPolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyOK  %+v", 200, o.Payload)
}

func (o *PatchVoicemailGroupPolicyOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyOK  %+v", 200, o.Payload)
}

func (o *PatchVoicemailGroupPolicyOK) GetPayload() *models.VoicemailGroupPolicy {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailGroupPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyBadRequest creates a PatchVoicemailGroupPolicyBadRequest with default headers values
func NewPatchVoicemailGroupPolicyBadRequest() *PatchVoicemailGroupPolicyBadRequest {
	return &PatchVoicemailGroupPolicyBadRequest{}
}

/*
PatchVoicemailGroupPolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchVoicemailGroupPolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy bad request response has a 2xx status code
func (o *PatchVoicemailGroupPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy bad request response has a 3xx status code
func (o *PatchVoicemailGroupPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy bad request response has a 4xx status code
func (o *PatchVoicemailGroupPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy bad request response has a 5xx status code
func (o *PatchVoicemailGroupPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy bad request response a status code equal to that given
func (o *PatchVoicemailGroupPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchVoicemailGroupPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVoicemailGroupPolicyBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PatchVoicemailGroupPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyUnauthorized creates a PatchVoicemailGroupPolicyUnauthorized with default headers values
func NewPatchVoicemailGroupPolicyUnauthorized() *PatchVoicemailGroupPolicyUnauthorized {
	return &PatchVoicemailGroupPolicyUnauthorized{}
}

/*
PatchVoicemailGroupPolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchVoicemailGroupPolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy unauthorized response has a 2xx status code
func (o *PatchVoicemailGroupPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy unauthorized response has a 3xx status code
func (o *PatchVoicemailGroupPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy unauthorized response has a 4xx status code
func (o *PatchVoicemailGroupPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy unauthorized response has a 5xx status code
func (o *PatchVoicemailGroupPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy unauthorized response a status code equal to that given
func (o *PatchVoicemailGroupPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchVoicemailGroupPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchVoicemailGroupPolicyUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchVoicemailGroupPolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyForbidden creates a PatchVoicemailGroupPolicyForbidden with default headers values
func NewPatchVoicemailGroupPolicyForbidden() *PatchVoicemailGroupPolicyForbidden {
	return &PatchVoicemailGroupPolicyForbidden{}
}

/*
PatchVoicemailGroupPolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchVoicemailGroupPolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy forbidden response has a 2xx status code
func (o *PatchVoicemailGroupPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy forbidden response has a 3xx status code
func (o *PatchVoicemailGroupPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy forbidden response has a 4xx status code
func (o *PatchVoicemailGroupPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy forbidden response has a 5xx status code
func (o *PatchVoicemailGroupPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy forbidden response a status code equal to that given
func (o *PatchVoicemailGroupPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchVoicemailGroupPolicyForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyForbidden  %+v", 403, o.Payload)
}

func (o *PatchVoicemailGroupPolicyForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyForbidden  %+v", 403, o.Payload)
}

func (o *PatchVoicemailGroupPolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyNotFound creates a PatchVoicemailGroupPolicyNotFound with default headers values
func NewPatchVoicemailGroupPolicyNotFound() *PatchVoicemailGroupPolicyNotFound {
	return &PatchVoicemailGroupPolicyNotFound{}
}

/*
PatchVoicemailGroupPolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchVoicemailGroupPolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy not found response has a 2xx status code
func (o *PatchVoicemailGroupPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy not found response has a 3xx status code
func (o *PatchVoicemailGroupPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy not found response has a 4xx status code
func (o *PatchVoicemailGroupPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy not found response has a 5xx status code
func (o *PatchVoicemailGroupPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy not found response a status code equal to that given
func (o *PatchVoicemailGroupPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchVoicemailGroupPolicyNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyNotFound  %+v", 404, o.Payload)
}

func (o *PatchVoicemailGroupPolicyNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyNotFound  %+v", 404, o.Payload)
}

func (o *PatchVoicemailGroupPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyRequestTimeout creates a PatchVoicemailGroupPolicyRequestTimeout with default headers values
func NewPatchVoicemailGroupPolicyRequestTimeout() *PatchVoicemailGroupPolicyRequestTimeout {
	return &PatchVoicemailGroupPolicyRequestTimeout{}
}

/*
PatchVoicemailGroupPolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchVoicemailGroupPolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy request timeout response has a 2xx status code
func (o *PatchVoicemailGroupPolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy request timeout response has a 3xx status code
func (o *PatchVoicemailGroupPolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy request timeout response has a 4xx status code
func (o *PatchVoicemailGroupPolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy request timeout response has a 5xx status code
func (o *PatchVoicemailGroupPolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy request timeout response a status code equal to that given
func (o *PatchVoicemailGroupPolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchVoicemailGroupPolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchVoicemailGroupPolicyRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchVoicemailGroupPolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyRequestEntityTooLarge creates a PatchVoicemailGroupPolicyRequestEntityTooLarge with default headers values
func NewPatchVoicemailGroupPolicyRequestEntityTooLarge() *PatchVoicemailGroupPolicyRequestEntityTooLarge {
	return &PatchVoicemailGroupPolicyRequestEntityTooLarge{}
}

/*
PatchVoicemailGroupPolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchVoicemailGroupPolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy request entity too large response has a 2xx status code
func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy request entity too large response has a 3xx status code
func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy request entity too large response has a 4xx status code
func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy request entity too large response has a 5xx status code
func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy request entity too large response a status code equal to that given
func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyUnsupportedMediaType creates a PatchVoicemailGroupPolicyUnsupportedMediaType with default headers values
func NewPatchVoicemailGroupPolicyUnsupportedMediaType() *PatchVoicemailGroupPolicyUnsupportedMediaType {
	return &PatchVoicemailGroupPolicyUnsupportedMediaType{}
}

/*
PatchVoicemailGroupPolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchVoicemailGroupPolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy unsupported media type response has a 2xx status code
func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy unsupported media type response has a 3xx status code
func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy unsupported media type response has a 4xx status code
func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy unsupported media type response has a 5xx status code
func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy unsupported media type response a status code equal to that given
func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyTooManyRequests creates a PatchVoicemailGroupPolicyTooManyRequests with default headers values
func NewPatchVoicemailGroupPolicyTooManyRequests() *PatchVoicemailGroupPolicyTooManyRequests {
	return &PatchVoicemailGroupPolicyTooManyRequests{}
}

/*
PatchVoicemailGroupPolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchVoicemailGroupPolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy too many requests response has a 2xx status code
func (o *PatchVoicemailGroupPolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy too many requests response has a 3xx status code
func (o *PatchVoicemailGroupPolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy too many requests response has a 4xx status code
func (o *PatchVoicemailGroupPolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch voicemail group policy too many requests response has a 5xx status code
func (o *PatchVoicemailGroupPolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch voicemail group policy too many requests response a status code equal to that given
func (o *PatchVoicemailGroupPolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchVoicemailGroupPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchVoicemailGroupPolicyTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchVoicemailGroupPolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyInternalServerError creates a PatchVoicemailGroupPolicyInternalServerError with default headers values
func NewPatchVoicemailGroupPolicyInternalServerError() *PatchVoicemailGroupPolicyInternalServerError {
	return &PatchVoicemailGroupPolicyInternalServerError{}
}

/*
PatchVoicemailGroupPolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchVoicemailGroupPolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy internal server error response has a 2xx status code
func (o *PatchVoicemailGroupPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy internal server error response has a 3xx status code
func (o *PatchVoicemailGroupPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy internal server error response has a 4xx status code
func (o *PatchVoicemailGroupPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail group policy internal server error response has a 5xx status code
func (o *PatchVoicemailGroupPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch voicemail group policy internal server error response a status code equal to that given
func (o *PatchVoicemailGroupPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchVoicemailGroupPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchVoicemailGroupPolicyInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchVoicemailGroupPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyServiceUnavailable creates a PatchVoicemailGroupPolicyServiceUnavailable with default headers values
func NewPatchVoicemailGroupPolicyServiceUnavailable() *PatchVoicemailGroupPolicyServiceUnavailable {
	return &PatchVoicemailGroupPolicyServiceUnavailable{}
}

/*
PatchVoicemailGroupPolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchVoicemailGroupPolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy service unavailable response has a 2xx status code
func (o *PatchVoicemailGroupPolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy service unavailable response has a 3xx status code
func (o *PatchVoicemailGroupPolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy service unavailable response has a 4xx status code
func (o *PatchVoicemailGroupPolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail group policy service unavailable response has a 5xx status code
func (o *PatchVoicemailGroupPolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch voicemail group policy service unavailable response a status code equal to that given
func (o *PatchVoicemailGroupPolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchVoicemailGroupPolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchVoicemailGroupPolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchVoicemailGroupPolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchVoicemailGroupPolicyGatewayTimeout creates a PatchVoicemailGroupPolicyGatewayTimeout with default headers values
func NewPatchVoicemailGroupPolicyGatewayTimeout() *PatchVoicemailGroupPolicyGatewayTimeout {
	return &PatchVoicemailGroupPolicyGatewayTimeout{}
}

/*
PatchVoicemailGroupPolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchVoicemailGroupPolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch voicemail group policy gateway timeout response has a 2xx status code
func (o *PatchVoicemailGroupPolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch voicemail group policy gateway timeout response has a 3xx status code
func (o *PatchVoicemailGroupPolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch voicemail group policy gateway timeout response has a 4xx status code
func (o *PatchVoicemailGroupPolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch voicemail group policy gateway timeout response has a 5xx status code
func (o *PatchVoicemailGroupPolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch voicemail group policy gateway timeout response a status code equal to that given
func (o *PatchVoicemailGroupPolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchVoicemailGroupPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchVoicemailGroupPolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/voicemail/groups/{groupId}/policy][%d] patchVoicemailGroupPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchVoicemailGroupPolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchVoicemailGroupPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
