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

// PutRecordingCrossplatformMediaretentionpolicyReader is a Reader for the PutRecordingCrossplatformMediaretentionpolicy structure.
type PutRecordingCrossplatformMediaretentionpolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRecordingCrossplatformMediaretentionpolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRecordingCrossplatformMediaretentionpolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRecordingCrossplatformMediaretentionpolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRecordingCrossplatformMediaretentionpolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRecordingCrossplatformMediaretentionpolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRecordingCrossplatformMediaretentionpolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutRecordingCrossplatformMediaretentionpolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRecordingCrossplatformMediaretentionpolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRecordingCrossplatformMediaretentionpolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRecordingCrossplatformMediaretentionpolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRecordingCrossplatformMediaretentionpolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRecordingCrossplatformMediaretentionpolicyOK creates a PutRecordingCrossplatformMediaretentionpolicyOK with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyOK() *PutRecordingCrossplatformMediaretentionpolicyOK {
	return &PutRecordingCrossplatformMediaretentionpolicyOK{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type PutRecordingCrossplatformMediaretentionpolicyOK struct {
	Payload *models.CrossPlatformPolicy
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy o k response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy o k response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy o k response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy o k response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy o k response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutRecordingCrossplatformMediaretentionpolicyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyOK  %+v", 200, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyOK  %+v", 200, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyOK) GetPayload() *models.CrossPlatformPolicy {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrossPlatformPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyBadRequest creates a PutRecordingCrossplatformMediaretentionpolicyBadRequest with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyBadRequest() *PutRecordingCrossplatformMediaretentionpolicyBadRequest {
	return &PutRecordingCrossplatformMediaretentionpolicyBadRequest{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRecordingCrossplatformMediaretentionpolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy bad request response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy bad request response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy bad request response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy bad request response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy bad request response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyUnauthorized creates a PutRecordingCrossplatformMediaretentionpolicyUnauthorized with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyUnauthorized() *PutRecordingCrossplatformMediaretentionpolicyUnauthorized {
	return &PutRecordingCrossplatformMediaretentionpolicyUnauthorized{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRecordingCrossplatformMediaretentionpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy unauthorized response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy unauthorized response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy unauthorized response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy unauthorized response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy unauthorized response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyForbidden creates a PutRecordingCrossplatformMediaretentionpolicyForbidden with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyForbidden() *PutRecordingCrossplatformMediaretentionpolicyForbidden {
	return &PutRecordingCrossplatformMediaretentionpolicyForbidden{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutRecordingCrossplatformMediaretentionpolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy forbidden response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy forbidden response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy forbidden response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy forbidden response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy forbidden response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyNotFound creates a PutRecordingCrossplatformMediaretentionpolicyNotFound with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyNotFound() *PutRecordingCrossplatformMediaretentionpolicyNotFound {
	return &PutRecordingCrossplatformMediaretentionpolicyNotFound{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutRecordingCrossplatformMediaretentionpolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy not found response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy not found response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy not found response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy not found response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy not found response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyRequestTimeout creates a PutRecordingCrossplatformMediaretentionpolicyRequestTimeout with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyRequestTimeout() *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout {
	return &PutRecordingCrossplatformMediaretentionpolicyRequestTimeout{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutRecordingCrossplatformMediaretentionpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy request timeout response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy request timeout response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy request timeout response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy request timeout response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy request timeout response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge creates a PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge() *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge {
	return &PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy request entity too large response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy request entity too large response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy request entity too large response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy request entity too large response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy request entity too large response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType creates a PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType() *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType {
	return &PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy unsupported media type response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy unsupported media type response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy unsupported media type response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy unsupported media type response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy unsupported media type response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyTooManyRequests creates a PutRecordingCrossplatformMediaretentionpolicyTooManyRequests with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyTooManyRequests() *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests {
	return &PutRecordingCrossplatformMediaretentionpolicyTooManyRequests{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutRecordingCrossplatformMediaretentionpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy too many requests response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy too many requests response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy too many requests response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy too many requests response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy too many requests response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyInternalServerError creates a PutRecordingCrossplatformMediaretentionpolicyInternalServerError with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyInternalServerError() *PutRecordingCrossplatformMediaretentionpolicyInternalServerError {
	return &PutRecordingCrossplatformMediaretentionpolicyInternalServerError{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRecordingCrossplatformMediaretentionpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy internal server error response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy internal server error response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy internal server error response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy internal server error response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy internal server error response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyServiceUnavailable creates a PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyServiceUnavailable() *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable {
	return &PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy service unavailable response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy service unavailable response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy service unavailable response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy service unavailable response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy service unavailable response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingCrossplatformMediaretentionpolicyGatewayTimeout creates a PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout with default headers values
func NewPutRecordingCrossplatformMediaretentionpolicyGatewayTimeout() *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout {
	return &PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout{}
}

/*
PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put recording crossplatform mediaretentionpolicy gateway timeout response has a 2xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put recording crossplatform mediaretentionpolicy gateway timeout response has a 3xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put recording crossplatform mediaretentionpolicy gateway timeout response has a 4xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put recording crossplatform mediaretentionpolicy gateway timeout response has a 5xx status code
func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put recording crossplatform mediaretentionpolicy gateway timeout response a status code equal to that given
func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/recording/crossplatform/mediaretentionpolicies/{policyId}][%d] putRecordingCrossplatformMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingCrossplatformMediaretentionpolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
