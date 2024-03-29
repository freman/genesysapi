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

// GetRecordingMediaretentionpolicyReader is a Reader for the GetRecordingMediaretentionpolicy structure.
type GetRecordingMediaretentionpolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingMediaretentionpolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingMediaretentionpolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingMediaretentionpolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingMediaretentionpolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingMediaretentionpolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingMediaretentionpolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRecordingMediaretentionpolicyRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingMediaretentionpolicyRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingMediaretentionpolicyUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingMediaretentionpolicyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingMediaretentionpolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingMediaretentionpolicyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingMediaretentionpolicyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingMediaretentionpolicyOK creates a GetRecordingMediaretentionpolicyOK with default headers values
func NewGetRecordingMediaretentionpolicyOK() *GetRecordingMediaretentionpolicyOK {
	return &GetRecordingMediaretentionpolicyOK{}
}

/*
GetRecordingMediaretentionpolicyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRecordingMediaretentionpolicyOK struct {
	Payload *models.Policy
}

// IsSuccess returns true when this get recording mediaretentionpolicy o k response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get recording mediaretentionpolicy o k response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy o k response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording mediaretentionpolicy o k response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy o k response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRecordingMediaretentionpolicyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyOK  %+v", 200, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyOK) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyOK  %+v", 200, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyOK) GetPayload() *models.Policy {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Policy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyBadRequest creates a GetRecordingMediaretentionpolicyBadRequest with default headers values
func NewGetRecordingMediaretentionpolicyBadRequest() *GetRecordingMediaretentionpolicyBadRequest {
	return &GetRecordingMediaretentionpolicyBadRequest{}
}

/*
GetRecordingMediaretentionpolicyBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingMediaretentionpolicyBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy bad request response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy bad request response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy bad request response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy bad request response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy bad request response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRecordingMediaretentionpolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyUnauthorized creates a GetRecordingMediaretentionpolicyUnauthorized with default headers values
func NewGetRecordingMediaretentionpolicyUnauthorized() *GetRecordingMediaretentionpolicyUnauthorized {
	return &GetRecordingMediaretentionpolicyUnauthorized{}
}

/*
GetRecordingMediaretentionpolicyUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingMediaretentionpolicyUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy unauthorized response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy unauthorized response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy unauthorized response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy unauthorized response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy unauthorized response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRecordingMediaretentionpolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyForbidden creates a GetRecordingMediaretentionpolicyForbidden with default headers values
func NewGetRecordingMediaretentionpolicyForbidden() *GetRecordingMediaretentionpolicyForbidden {
	return &GetRecordingMediaretentionpolicyForbidden{}
}

/*
GetRecordingMediaretentionpolicyForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingMediaretentionpolicyForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy forbidden response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy forbidden response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy forbidden response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy forbidden response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy forbidden response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRecordingMediaretentionpolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyNotFound creates a GetRecordingMediaretentionpolicyNotFound with default headers values
func NewGetRecordingMediaretentionpolicyNotFound() *GetRecordingMediaretentionpolicyNotFound {
	return &GetRecordingMediaretentionpolicyNotFound{}
}

/*
GetRecordingMediaretentionpolicyNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRecordingMediaretentionpolicyNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy not found response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy not found response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy not found response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy not found response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy not found response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRecordingMediaretentionpolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyRequestTimeout creates a GetRecordingMediaretentionpolicyRequestTimeout with default headers values
func NewGetRecordingMediaretentionpolicyRequestTimeout() *GetRecordingMediaretentionpolicyRequestTimeout {
	return &GetRecordingMediaretentionpolicyRequestTimeout{}
}

/*
GetRecordingMediaretentionpolicyRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingMediaretentionpolicyRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy request timeout response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy request timeout response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy request timeout response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy request timeout response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy request timeout response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRecordingMediaretentionpolicyRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyRequestEntityTooLarge creates a GetRecordingMediaretentionpolicyRequestEntityTooLarge with default headers values
func NewGetRecordingMediaretentionpolicyRequestEntityTooLarge() *GetRecordingMediaretentionpolicyRequestEntityTooLarge {
	return &GetRecordingMediaretentionpolicyRequestEntityTooLarge{}
}

/*
GetRecordingMediaretentionpolicyRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRecordingMediaretentionpolicyRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy request entity too large response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy request entity too large response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy request entity too large response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy request entity too large response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy request entity too large response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyUnsupportedMediaType creates a GetRecordingMediaretentionpolicyUnsupportedMediaType with default headers values
func NewGetRecordingMediaretentionpolicyUnsupportedMediaType() *GetRecordingMediaretentionpolicyUnsupportedMediaType {
	return &GetRecordingMediaretentionpolicyUnsupportedMediaType{}
}

/*
GetRecordingMediaretentionpolicyUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingMediaretentionpolicyUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy unsupported media type response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy unsupported media type response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy unsupported media type response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy unsupported media type response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy unsupported media type response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyTooManyRequests creates a GetRecordingMediaretentionpolicyTooManyRequests with default headers values
func NewGetRecordingMediaretentionpolicyTooManyRequests() *GetRecordingMediaretentionpolicyTooManyRequests {
	return &GetRecordingMediaretentionpolicyTooManyRequests{}
}

/*
GetRecordingMediaretentionpolicyTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingMediaretentionpolicyTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy too many requests response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy too many requests response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy too many requests response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording mediaretentionpolicy too many requests response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording mediaretentionpolicy too many requests response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRecordingMediaretentionpolicyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyInternalServerError creates a GetRecordingMediaretentionpolicyInternalServerError with default headers values
func NewGetRecordingMediaretentionpolicyInternalServerError() *GetRecordingMediaretentionpolicyInternalServerError {
	return &GetRecordingMediaretentionpolicyInternalServerError{}
}

/*
GetRecordingMediaretentionpolicyInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingMediaretentionpolicyInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy internal server error response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy internal server error response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy internal server error response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording mediaretentionpolicy internal server error response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording mediaretentionpolicy internal server error response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRecordingMediaretentionpolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyServiceUnavailable creates a GetRecordingMediaretentionpolicyServiceUnavailable with default headers values
func NewGetRecordingMediaretentionpolicyServiceUnavailable() *GetRecordingMediaretentionpolicyServiceUnavailable {
	return &GetRecordingMediaretentionpolicyServiceUnavailable{}
}

/*
GetRecordingMediaretentionpolicyServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingMediaretentionpolicyServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy service unavailable response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy service unavailable response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy service unavailable response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording mediaretentionpolicy service unavailable response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording mediaretentionpolicy service unavailable response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRecordingMediaretentionpolicyServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpolicyGatewayTimeout creates a GetRecordingMediaretentionpolicyGatewayTimeout with default headers values
func NewGetRecordingMediaretentionpolicyGatewayTimeout() *GetRecordingMediaretentionpolicyGatewayTimeout {
	return &GetRecordingMediaretentionpolicyGatewayTimeout{}
}

/*
GetRecordingMediaretentionpolicyGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRecordingMediaretentionpolicyGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording mediaretentionpolicy gateway timeout response has a 2xx status code
func (o *GetRecordingMediaretentionpolicyGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording mediaretentionpolicy gateway timeout response has a 3xx status code
func (o *GetRecordingMediaretentionpolicyGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording mediaretentionpolicy gateway timeout response has a 4xx status code
func (o *GetRecordingMediaretentionpolicyGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording mediaretentionpolicy gateway timeout response has a 5xx status code
func (o *GetRecordingMediaretentionpolicyGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording mediaretentionpolicy gateway timeout response a status code equal to that given
func (o *GetRecordingMediaretentionpolicyGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRecordingMediaretentionpolicyGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies/{policyId}][%d] getRecordingMediaretentionpolicyGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingMediaretentionpolicyGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpolicyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
