// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundRulesetsReader is a Reader for the GetOutboundRulesets structure.
type GetOutboundRulesetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundRulesetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundRulesetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundRulesetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundRulesetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundRulesetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundRulesetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundRulesetsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundRulesetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundRulesetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundRulesetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundRulesetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundRulesetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundRulesetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundRulesetsOK creates a GetOutboundRulesetsOK with default headers values
func NewGetOutboundRulesetsOK() *GetOutboundRulesetsOK {
	return &GetOutboundRulesetsOK{}
}

/*
GetOutboundRulesetsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundRulesetsOK struct {
	Payload *models.RuleSetEntityListing
}

// IsSuccess returns true when this get outbound rulesets o k response has a 2xx status code
func (o *GetOutboundRulesetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound rulesets o k response has a 3xx status code
func (o *GetOutboundRulesetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets o k response has a 4xx status code
func (o *GetOutboundRulesetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound rulesets o k response has a 5xx status code
func (o *GetOutboundRulesetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets o k response a status code equal to that given
func (o *GetOutboundRulesetsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundRulesetsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundRulesetsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundRulesetsOK) GetPayload() *models.RuleSetEntityListing {
	return o.Payload
}

func (o *GetOutboundRulesetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleSetEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsBadRequest creates a GetOutboundRulesetsBadRequest with default headers values
func NewGetOutboundRulesetsBadRequest() *GetOutboundRulesetsBadRequest {
	return &GetOutboundRulesetsBadRequest{}
}

/*
GetOutboundRulesetsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundRulesetsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets bad request response has a 2xx status code
func (o *GetOutboundRulesetsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets bad request response has a 3xx status code
func (o *GetOutboundRulesetsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets bad request response has a 4xx status code
func (o *GetOutboundRulesetsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets bad request response has a 5xx status code
func (o *GetOutboundRulesetsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets bad request response a status code equal to that given
func (o *GetOutboundRulesetsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundRulesetsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundRulesetsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundRulesetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsUnauthorized creates a GetOutboundRulesetsUnauthorized with default headers values
func NewGetOutboundRulesetsUnauthorized() *GetOutboundRulesetsUnauthorized {
	return &GetOutboundRulesetsUnauthorized{}
}

/*
GetOutboundRulesetsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundRulesetsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets unauthorized response has a 2xx status code
func (o *GetOutboundRulesetsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets unauthorized response has a 3xx status code
func (o *GetOutboundRulesetsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets unauthorized response has a 4xx status code
func (o *GetOutboundRulesetsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets unauthorized response has a 5xx status code
func (o *GetOutboundRulesetsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets unauthorized response a status code equal to that given
func (o *GetOutboundRulesetsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundRulesetsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundRulesetsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundRulesetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsForbidden creates a GetOutboundRulesetsForbidden with default headers values
func NewGetOutboundRulesetsForbidden() *GetOutboundRulesetsForbidden {
	return &GetOutboundRulesetsForbidden{}
}

/*
GetOutboundRulesetsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundRulesetsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets forbidden response has a 2xx status code
func (o *GetOutboundRulesetsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets forbidden response has a 3xx status code
func (o *GetOutboundRulesetsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets forbidden response has a 4xx status code
func (o *GetOutboundRulesetsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets forbidden response has a 5xx status code
func (o *GetOutboundRulesetsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets forbidden response a status code equal to that given
func (o *GetOutboundRulesetsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundRulesetsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundRulesetsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundRulesetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsNotFound creates a GetOutboundRulesetsNotFound with default headers values
func NewGetOutboundRulesetsNotFound() *GetOutboundRulesetsNotFound {
	return &GetOutboundRulesetsNotFound{}
}

/*
GetOutboundRulesetsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundRulesetsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets not found response has a 2xx status code
func (o *GetOutboundRulesetsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets not found response has a 3xx status code
func (o *GetOutboundRulesetsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets not found response has a 4xx status code
func (o *GetOutboundRulesetsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets not found response has a 5xx status code
func (o *GetOutboundRulesetsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets not found response a status code equal to that given
func (o *GetOutboundRulesetsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundRulesetsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundRulesetsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundRulesetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsRequestTimeout creates a GetOutboundRulesetsRequestTimeout with default headers values
func NewGetOutboundRulesetsRequestTimeout() *GetOutboundRulesetsRequestTimeout {
	return &GetOutboundRulesetsRequestTimeout{}
}

/*
GetOutboundRulesetsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundRulesetsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets request timeout response has a 2xx status code
func (o *GetOutboundRulesetsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets request timeout response has a 3xx status code
func (o *GetOutboundRulesetsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets request timeout response has a 4xx status code
func (o *GetOutboundRulesetsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets request timeout response has a 5xx status code
func (o *GetOutboundRulesetsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets request timeout response a status code equal to that given
func (o *GetOutboundRulesetsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundRulesetsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundRulesetsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundRulesetsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsRequestEntityTooLarge creates a GetOutboundRulesetsRequestEntityTooLarge with default headers values
func NewGetOutboundRulesetsRequestEntityTooLarge() *GetOutboundRulesetsRequestEntityTooLarge {
	return &GetOutboundRulesetsRequestEntityTooLarge{}
}

/*
GetOutboundRulesetsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundRulesetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets request entity too large response has a 2xx status code
func (o *GetOutboundRulesetsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets request entity too large response has a 3xx status code
func (o *GetOutboundRulesetsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets request entity too large response has a 4xx status code
func (o *GetOutboundRulesetsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets request entity too large response has a 5xx status code
func (o *GetOutboundRulesetsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets request entity too large response a status code equal to that given
func (o *GetOutboundRulesetsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundRulesetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundRulesetsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundRulesetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsUnsupportedMediaType creates a GetOutboundRulesetsUnsupportedMediaType with default headers values
func NewGetOutboundRulesetsUnsupportedMediaType() *GetOutboundRulesetsUnsupportedMediaType {
	return &GetOutboundRulesetsUnsupportedMediaType{}
}

/*
GetOutboundRulesetsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundRulesetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets unsupported media type response has a 2xx status code
func (o *GetOutboundRulesetsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets unsupported media type response has a 3xx status code
func (o *GetOutboundRulesetsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets unsupported media type response has a 4xx status code
func (o *GetOutboundRulesetsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets unsupported media type response has a 5xx status code
func (o *GetOutboundRulesetsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets unsupported media type response a status code equal to that given
func (o *GetOutboundRulesetsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundRulesetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundRulesetsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundRulesetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsTooManyRequests creates a GetOutboundRulesetsTooManyRequests with default headers values
func NewGetOutboundRulesetsTooManyRequests() *GetOutboundRulesetsTooManyRequests {
	return &GetOutboundRulesetsTooManyRequests{}
}

/*
GetOutboundRulesetsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundRulesetsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets too many requests response has a 2xx status code
func (o *GetOutboundRulesetsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets too many requests response has a 3xx status code
func (o *GetOutboundRulesetsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets too many requests response has a 4xx status code
func (o *GetOutboundRulesetsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound rulesets too many requests response has a 5xx status code
func (o *GetOutboundRulesetsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound rulesets too many requests response a status code equal to that given
func (o *GetOutboundRulesetsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundRulesetsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundRulesetsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundRulesetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsInternalServerError creates a GetOutboundRulesetsInternalServerError with default headers values
func NewGetOutboundRulesetsInternalServerError() *GetOutboundRulesetsInternalServerError {
	return &GetOutboundRulesetsInternalServerError{}
}

/*
GetOutboundRulesetsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundRulesetsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets internal server error response has a 2xx status code
func (o *GetOutboundRulesetsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets internal server error response has a 3xx status code
func (o *GetOutboundRulesetsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets internal server error response has a 4xx status code
func (o *GetOutboundRulesetsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound rulesets internal server error response has a 5xx status code
func (o *GetOutboundRulesetsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound rulesets internal server error response a status code equal to that given
func (o *GetOutboundRulesetsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundRulesetsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundRulesetsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundRulesetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsServiceUnavailable creates a GetOutboundRulesetsServiceUnavailable with default headers values
func NewGetOutboundRulesetsServiceUnavailable() *GetOutboundRulesetsServiceUnavailable {
	return &GetOutboundRulesetsServiceUnavailable{}
}

/*
GetOutboundRulesetsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundRulesetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets service unavailable response has a 2xx status code
func (o *GetOutboundRulesetsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets service unavailable response has a 3xx status code
func (o *GetOutboundRulesetsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets service unavailable response has a 4xx status code
func (o *GetOutboundRulesetsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound rulesets service unavailable response has a 5xx status code
func (o *GetOutboundRulesetsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound rulesets service unavailable response a status code equal to that given
func (o *GetOutboundRulesetsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundRulesetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundRulesetsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundRulesetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundRulesetsGatewayTimeout creates a GetOutboundRulesetsGatewayTimeout with default headers values
func NewGetOutboundRulesetsGatewayTimeout() *GetOutboundRulesetsGatewayTimeout {
	return &GetOutboundRulesetsGatewayTimeout{}
}

/*
GetOutboundRulesetsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundRulesetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound rulesets gateway timeout response has a 2xx status code
func (o *GetOutboundRulesetsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound rulesets gateway timeout response has a 3xx status code
func (o *GetOutboundRulesetsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound rulesets gateway timeout response has a 4xx status code
func (o *GetOutboundRulesetsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound rulesets gateway timeout response has a 5xx status code
func (o *GetOutboundRulesetsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound rulesets gateway timeout response a status code equal to that given
func (o *GetOutboundRulesetsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundRulesetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundRulesetsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/rulesets][%d] getOutboundRulesetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundRulesetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundRulesetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
