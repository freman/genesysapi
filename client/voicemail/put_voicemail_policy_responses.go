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

// PutVoicemailPolicyReader is a Reader for the PutVoicemailPolicy structure.
type PutVoicemailPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutVoicemailPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutVoicemailPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutVoicemailPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutVoicemailPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutVoicemailPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutVoicemailPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutVoicemailPolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutVoicemailPolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutVoicemailPolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutVoicemailPolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewPutVoicemailPolicyFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutVoicemailPolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutVoicemailPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutVoicemailPolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutVoicemailPolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutVoicemailPolicyOK creates a PutVoicemailPolicyOK with default headers values
func NewPutVoicemailPolicyOK() *PutVoicemailPolicyOK {
	return &PutVoicemailPolicyOK{}
}

/*
PutVoicemailPolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type PutVoicemailPolicyOK struct {
	Payload *models.VoicemailOrganizationPolicy
}

// IsSuccess returns true when this put voicemail policy o k response has a 2xx status code
func (o *PutVoicemailPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put voicemail policy o k response has a 3xx status code
func (o *PutVoicemailPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy o k response has a 4xx status code
func (o *PutVoicemailPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail policy o k response has a 5xx status code
func (o *PutVoicemailPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy o k response a status code equal to that given
func (o *PutVoicemailPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutVoicemailPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyOK  %+v", 200, o.Payload)
}

func (o *PutVoicemailPolicyOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyOK  %+v", 200, o.Payload)
}

func (o *PutVoicemailPolicyOK) GetPayload() *models.VoicemailOrganizationPolicy {
	return o.Payload
}

func (o *PutVoicemailPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoicemailOrganizationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyBadRequest creates a PutVoicemailPolicyBadRequest with default headers values
func NewPutVoicemailPolicyBadRequest() *PutVoicemailPolicyBadRequest {
	return &PutVoicemailPolicyBadRequest{}
}

/*
PutVoicemailPolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutVoicemailPolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy bad request response has a 2xx status code
func (o *PutVoicemailPolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy bad request response has a 3xx status code
func (o *PutVoicemailPolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy bad request response has a 4xx status code
func (o *PutVoicemailPolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy bad request response has a 5xx status code
func (o *PutVoicemailPolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy bad request response a status code equal to that given
func (o *PutVoicemailPolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutVoicemailPolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutVoicemailPolicyBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutVoicemailPolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyUnauthorized creates a PutVoicemailPolicyUnauthorized with default headers values
func NewPutVoicemailPolicyUnauthorized() *PutVoicemailPolicyUnauthorized {
	return &PutVoicemailPolicyUnauthorized{}
}

/*
PutVoicemailPolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutVoicemailPolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy unauthorized response has a 2xx status code
func (o *PutVoicemailPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy unauthorized response has a 3xx status code
func (o *PutVoicemailPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy unauthorized response has a 4xx status code
func (o *PutVoicemailPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy unauthorized response has a 5xx status code
func (o *PutVoicemailPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy unauthorized response a status code equal to that given
func (o *PutVoicemailPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutVoicemailPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutVoicemailPolicyUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutVoicemailPolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyForbidden creates a PutVoicemailPolicyForbidden with default headers values
func NewPutVoicemailPolicyForbidden() *PutVoicemailPolicyForbidden {
	return &PutVoicemailPolicyForbidden{}
}

/*
PutVoicemailPolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutVoicemailPolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy forbidden response has a 2xx status code
func (o *PutVoicemailPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy forbidden response has a 3xx status code
func (o *PutVoicemailPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy forbidden response has a 4xx status code
func (o *PutVoicemailPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy forbidden response has a 5xx status code
func (o *PutVoicemailPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy forbidden response a status code equal to that given
func (o *PutVoicemailPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutVoicemailPolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutVoicemailPolicyForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutVoicemailPolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyNotFound creates a PutVoicemailPolicyNotFound with default headers values
func NewPutVoicemailPolicyNotFound() *PutVoicemailPolicyNotFound {
	return &PutVoicemailPolicyNotFound{}
}

/*
PutVoicemailPolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutVoicemailPolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy not found response has a 2xx status code
func (o *PutVoicemailPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy not found response has a 3xx status code
func (o *PutVoicemailPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy not found response has a 4xx status code
func (o *PutVoicemailPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy not found response has a 5xx status code
func (o *PutVoicemailPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy not found response a status code equal to that given
func (o *PutVoicemailPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutVoicemailPolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutVoicemailPolicyNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutVoicemailPolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyRequestTimeout creates a PutVoicemailPolicyRequestTimeout with default headers values
func NewPutVoicemailPolicyRequestTimeout() *PutVoicemailPolicyRequestTimeout {
	return &PutVoicemailPolicyRequestTimeout{}
}

/*
PutVoicemailPolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutVoicemailPolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy request timeout response has a 2xx status code
func (o *PutVoicemailPolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy request timeout response has a 3xx status code
func (o *PutVoicemailPolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy request timeout response has a 4xx status code
func (o *PutVoicemailPolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy request timeout response has a 5xx status code
func (o *PutVoicemailPolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy request timeout response a status code equal to that given
func (o *PutVoicemailPolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutVoicemailPolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutVoicemailPolicyRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutVoicemailPolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyConflict creates a PutVoicemailPolicyConflict with default headers values
func NewPutVoicemailPolicyConflict() *PutVoicemailPolicyConflict {
	return &PutVoicemailPolicyConflict{}
}

/*
PutVoicemailPolicyConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutVoicemailPolicyConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy conflict response has a 2xx status code
func (o *PutVoicemailPolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy conflict response has a 3xx status code
func (o *PutVoicemailPolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy conflict response has a 4xx status code
func (o *PutVoicemailPolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy conflict response has a 5xx status code
func (o *PutVoicemailPolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy conflict response a status code equal to that given
func (o *PutVoicemailPolicyConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutVoicemailPolicyConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyConflict  %+v", 409, o.Payload)
}

func (o *PutVoicemailPolicyConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyConflict  %+v", 409, o.Payload)
}

func (o *PutVoicemailPolicyConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyRequestEntityTooLarge creates a PutVoicemailPolicyRequestEntityTooLarge with default headers values
func NewPutVoicemailPolicyRequestEntityTooLarge() *PutVoicemailPolicyRequestEntityTooLarge {
	return &PutVoicemailPolicyRequestEntityTooLarge{}
}

/*
PutVoicemailPolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutVoicemailPolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy request entity too large response has a 2xx status code
func (o *PutVoicemailPolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy request entity too large response has a 3xx status code
func (o *PutVoicemailPolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy request entity too large response has a 4xx status code
func (o *PutVoicemailPolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy request entity too large response has a 5xx status code
func (o *PutVoicemailPolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy request entity too large response a status code equal to that given
func (o *PutVoicemailPolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyUnsupportedMediaType creates a PutVoicemailPolicyUnsupportedMediaType with default headers values
func NewPutVoicemailPolicyUnsupportedMediaType() *PutVoicemailPolicyUnsupportedMediaType {
	return &PutVoicemailPolicyUnsupportedMediaType{}
}

/*
PutVoicemailPolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutVoicemailPolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy unsupported media type response has a 2xx status code
func (o *PutVoicemailPolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy unsupported media type response has a 3xx status code
func (o *PutVoicemailPolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy unsupported media type response has a 4xx status code
func (o *PutVoicemailPolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy unsupported media type response has a 5xx status code
func (o *PutVoicemailPolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy unsupported media type response a status code equal to that given
func (o *PutVoicemailPolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutVoicemailPolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutVoicemailPolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutVoicemailPolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyFailedDependency creates a PutVoicemailPolicyFailedDependency with default headers values
func NewPutVoicemailPolicyFailedDependency() *PutVoicemailPolicyFailedDependency {
	return &PutVoicemailPolicyFailedDependency{}
}

/*
PutVoicemailPolicyFailedDependency describes a response with status code 424, with default header values.

PutVoicemailPolicyFailedDependency put voicemail policy failed dependency
*/
type PutVoicemailPolicyFailedDependency struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy failed dependency response has a 2xx status code
func (o *PutVoicemailPolicyFailedDependency) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy failed dependency response has a 3xx status code
func (o *PutVoicemailPolicyFailedDependency) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy failed dependency response has a 4xx status code
func (o *PutVoicemailPolicyFailedDependency) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy failed dependency response has a 5xx status code
func (o *PutVoicemailPolicyFailedDependency) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy failed dependency response a status code equal to that given
func (o *PutVoicemailPolicyFailedDependency) IsCode(code int) bool {
	return code == 424
}

func (o *PutVoicemailPolicyFailedDependency) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyFailedDependency  %+v", 424, o.Payload)
}

func (o *PutVoicemailPolicyFailedDependency) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyFailedDependency  %+v", 424, o.Payload)
}

func (o *PutVoicemailPolicyFailedDependency) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyTooManyRequests creates a PutVoicemailPolicyTooManyRequests with default headers values
func NewPutVoicemailPolicyTooManyRequests() *PutVoicemailPolicyTooManyRequests {
	return &PutVoicemailPolicyTooManyRequests{}
}

/*
PutVoicemailPolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutVoicemailPolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy too many requests response has a 2xx status code
func (o *PutVoicemailPolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy too many requests response has a 3xx status code
func (o *PutVoicemailPolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy too many requests response has a 4xx status code
func (o *PutVoicemailPolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put voicemail policy too many requests response has a 5xx status code
func (o *PutVoicemailPolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put voicemail policy too many requests response a status code equal to that given
func (o *PutVoicemailPolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutVoicemailPolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutVoicemailPolicyTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutVoicemailPolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyInternalServerError creates a PutVoicemailPolicyInternalServerError with default headers values
func NewPutVoicemailPolicyInternalServerError() *PutVoicemailPolicyInternalServerError {
	return &PutVoicemailPolicyInternalServerError{}
}

/*
PutVoicemailPolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutVoicemailPolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy internal server error response has a 2xx status code
func (o *PutVoicemailPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy internal server error response has a 3xx status code
func (o *PutVoicemailPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy internal server error response has a 4xx status code
func (o *PutVoicemailPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail policy internal server error response has a 5xx status code
func (o *PutVoicemailPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put voicemail policy internal server error response a status code equal to that given
func (o *PutVoicemailPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutVoicemailPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutVoicemailPolicyInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutVoicemailPolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyServiceUnavailable creates a PutVoicemailPolicyServiceUnavailable with default headers values
func NewPutVoicemailPolicyServiceUnavailable() *PutVoicemailPolicyServiceUnavailable {
	return &PutVoicemailPolicyServiceUnavailable{}
}

/*
PutVoicemailPolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutVoicemailPolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy service unavailable response has a 2xx status code
func (o *PutVoicemailPolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy service unavailable response has a 3xx status code
func (o *PutVoicemailPolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy service unavailable response has a 4xx status code
func (o *PutVoicemailPolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail policy service unavailable response has a 5xx status code
func (o *PutVoicemailPolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put voicemail policy service unavailable response a status code equal to that given
func (o *PutVoicemailPolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutVoicemailPolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutVoicemailPolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutVoicemailPolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutVoicemailPolicyGatewayTimeout creates a PutVoicemailPolicyGatewayTimeout with default headers values
func NewPutVoicemailPolicyGatewayTimeout() *PutVoicemailPolicyGatewayTimeout {
	return &PutVoicemailPolicyGatewayTimeout{}
}

/*
PutVoicemailPolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutVoicemailPolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put voicemail policy gateway timeout response has a 2xx status code
func (o *PutVoicemailPolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put voicemail policy gateway timeout response has a 3xx status code
func (o *PutVoicemailPolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put voicemail policy gateway timeout response has a 4xx status code
func (o *PutVoicemailPolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put voicemail policy gateway timeout response has a 5xx status code
func (o *PutVoicemailPolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put voicemail policy gateway timeout response a status code equal to that given
func (o *PutVoicemailPolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutVoicemailPolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutVoicemailPolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/voicemail/policy][%d] putVoicemailPolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutVoicemailPolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutVoicemailPolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
