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

// PatchRecordingCrossplatformMediaretentionpolicyReader is a Reader for the PatchRecordingCrossplatformMediaretentionpolicy structure.
type PatchRecordingCrossplatformMediaretentionpolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRecordingCrossplatformMediaretentionpolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRecordingCrossplatformMediaretentionpolicyOK creates a PatchRecordingCrossplatformMediaretentionpolicyOK with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyOK() *PatchRecordingCrossplatformMediaretentionpolicyOK {
	return &PatchRecordingCrossplatformMediaretentionpolicyOK{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchRecordingCrossplatformMediaretentionpolicyOK struct {
	Payload *models.CrossPlatformPolicy
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy o k response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy o k response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy o k response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy o k response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy o k response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyOK  %+v", 200, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyOK  %+v", 200, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) GetPayload() *models.CrossPlatformPolicy {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrossPlatformPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyBadRequest creates a PatchRecordingCrossplatformMediaretentionpolicyBadRequest with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyBadRequest() *PatchRecordingCrossplatformMediaretentionpolicyBadRequest {
	return &PatchRecordingCrossplatformMediaretentionpolicyBadRequest{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRecordingCrossplatformMediaretentionpolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy bad request response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy bad request response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy bad request response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy bad request response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy bad request response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyUnauthorized creates a PatchRecordingCrossplatformMediaretentionpolicyUnauthorized with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyUnauthorized() *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized {
	return &PatchRecordingCrossplatformMediaretentionpolicyUnauthorized{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRecordingCrossplatformMediaretentionpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy unauthorized response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy unauthorized response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy unauthorized response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy unauthorized response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy unauthorized response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyForbidden creates a PatchRecordingCrossplatformMediaretentionpolicyForbidden with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyForbidden() *PatchRecordingCrossplatformMediaretentionpolicyForbidden {
	return &PatchRecordingCrossplatformMediaretentionpolicyForbidden{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchRecordingCrossplatformMediaretentionpolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy forbidden response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy forbidden response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy forbidden response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy forbidden response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy forbidden response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyNotFound creates a PatchRecordingCrossplatformMediaretentionpolicyNotFound with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyNotFound() *PatchRecordingCrossplatformMediaretentionpolicyNotFound {
	return &PatchRecordingCrossplatformMediaretentionpolicyNotFound{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchRecordingCrossplatformMediaretentionpolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy not found response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy not found response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy not found response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy not found response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy not found response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyRequestTimeout creates a PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyRequestTimeout() *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout {
	return &PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy request timeout response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy request timeout response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy request timeout response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy request timeout response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy request timeout response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge creates a PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge() *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge {
	return &PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy request entity too large response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy request entity too large response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy request entity too large response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy request entity too large response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy request entity too large response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType creates a PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType() *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType {
	return &PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy unsupported media type response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy unsupported media type response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy unsupported media type response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy unsupported media type response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy unsupported media type response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyTooManyRequests creates a PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyTooManyRequests() *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests {
	return &PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy too many requests response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy too many requests response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy too many requests response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy too many requests response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy too many requests response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyInternalServerError creates a PatchRecordingCrossplatformMediaretentionpolicyInternalServerError with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyInternalServerError() *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError {
	return &PatchRecordingCrossplatformMediaretentionpolicyInternalServerError{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRecordingCrossplatformMediaretentionpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy internal server error response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy internal server error response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy internal server error response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy internal server error response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy internal server error response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable creates a PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable() *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable {
	return &PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy service unavailable response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy service unavailable response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy service unavailable response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy service unavailable response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy service unavailable response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout creates a PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout with default headers values
func NewPatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout() *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout {
	return &PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout{}
}

/*
PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch recording crossplatform mediaretentionpolicy gateway timeout response has a 2xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch recording crossplatform mediaretentionpolicy gateway timeout response has a 3xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch recording crossplatform mediaretentionpolicy gateway timeout response has a 4xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch recording crossplatform mediaretentionpolicy gateway timeout response has a 5xx status code
func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch recording crossplatform mediaretentionpolicy gateway timeout response a status code equal to that given
func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] patchRecordingCrossplatformMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRecordingCrossplatformMediaretentionpolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
