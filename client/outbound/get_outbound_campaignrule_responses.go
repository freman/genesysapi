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

// GetOutboundCampaignruleReader is a Reader for the GetOutboundCampaignrule structure.
type GetOutboundCampaignruleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundCampaignruleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundCampaignruleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundCampaignruleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundCampaignruleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundCampaignruleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundCampaignruleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundCampaignruleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundCampaignruleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundCampaignruleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundCampaignruleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundCampaignruleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundCampaignruleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundCampaignruleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundCampaignruleOK creates a GetOutboundCampaignruleOK with default headers values
func NewGetOutboundCampaignruleOK() *GetOutboundCampaignruleOK {
	return &GetOutboundCampaignruleOK{}
}

/*
GetOutboundCampaignruleOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOutboundCampaignruleOK struct {
	Payload *models.CampaignRule
}

// IsSuccess returns true when this get outbound campaignrule o k response has a 2xx status code
func (o *GetOutboundCampaignruleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get outbound campaignrule o k response has a 3xx status code
func (o *GetOutboundCampaignruleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule o k response has a 4xx status code
func (o *GetOutboundCampaignruleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaignrule o k response has a 5xx status code
func (o *GetOutboundCampaignruleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule o k response a status code equal to that given
func (o *GetOutboundCampaignruleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOutboundCampaignruleOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignruleOK) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleOK  %+v", 200, o.Payload)
}

func (o *GetOutboundCampaignruleOK) GetPayload() *models.CampaignRule {
	return o.Payload
}

func (o *GetOutboundCampaignruleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CampaignRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleBadRequest creates a GetOutboundCampaignruleBadRequest with default headers values
func NewGetOutboundCampaignruleBadRequest() *GetOutboundCampaignruleBadRequest {
	return &GetOutboundCampaignruleBadRequest{}
}

/*
GetOutboundCampaignruleBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignruleBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule bad request response has a 2xx status code
func (o *GetOutboundCampaignruleBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule bad request response has a 3xx status code
func (o *GetOutboundCampaignruleBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule bad request response has a 4xx status code
func (o *GetOutboundCampaignruleBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule bad request response has a 5xx status code
func (o *GetOutboundCampaignruleBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule bad request response a status code equal to that given
func (o *GetOutboundCampaignruleBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOutboundCampaignruleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignruleBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundCampaignruleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleUnauthorized creates a GetOutboundCampaignruleUnauthorized with default headers values
func NewGetOutboundCampaignruleUnauthorized() *GetOutboundCampaignruleUnauthorized {
	return &GetOutboundCampaignruleUnauthorized{}
}

/*
GetOutboundCampaignruleUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignruleUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule unauthorized response has a 2xx status code
func (o *GetOutboundCampaignruleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule unauthorized response has a 3xx status code
func (o *GetOutboundCampaignruleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule unauthorized response has a 4xx status code
func (o *GetOutboundCampaignruleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule unauthorized response has a 5xx status code
func (o *GetOutboundCampaignruleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule unauthorized response a status code equal to that given
func (o *GetOutboundCampaignruleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOutboundCampaignruleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignruleUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundCampaignruleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleForbidden creates a GetOutboundCampaignruleForbidden with default headers values
func NewGetOutboundCampaignruleForbidden() *GetOutboundCampaignruleForbidden {
	return &GetOutboundCampaignruleForbidden{}
}

/*
GetOutboundCampaignruleForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignruleForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule forbidden response has a 2xx status code
func (o *GetOutboundCampaignruleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule forbidden response has a 3xx status code
func (o *GetOutboundCampaignruleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule forbidden response has a 4xx status code
func (o *GetOutboundCampaignruleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule forbidden response has a 5xx status code
func (o *GetOutboundCampaignruleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule forbidden response a status code equal to that given
func (o *GetOutboundCampaignruleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOutboundCampaignruleForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignruleForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundCampaignruleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleNotFound creates a GetOutboundCampaignruleNotFound with default headers values
func NewGetOutboundCampaignruleNotFound() *GetOutboundCampaignruleNotFound {
	return &GetOutboundCampaignruleNotFound{}
}

/*
GetOutboundCampaignruleNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignruleNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule not found response has a 2xx status code
func (o *GetOutboundCampaignruleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule not found response has a 3xx status code
func (o *GetOutboundCampaignruleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule not found response has a 4xx status code
func (o *GetOutboundCampaignruleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule not found response has a 5xx status code
func (o *GetOutboundCampaignruleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule not found response a status code equal to that given
func (o *GetOutboundCampaignruleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOutboundCampaignruleNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignruleNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundCampaignruleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleRequestTimeout creates a GetOutboundCampaignruleRequestTimeout with default headers values
func NewGetOutboundCampaignruleRequestTimeout() *GetOutboundCampaignruleRequestTimeout {
	return &GetOutboundCampaignruleRequestTimeout{}
}

/*
GetOutboundCampaignruleRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundCampaignruleRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule request timeout response has a 2xx status code
func (o *GetOutboundCampaignruleRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule request timeout response has a 3xx status code
func (o *GetOutboundCampaignruleRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule request timeout response has a 4xx status code
func (o *GetOutboundCampaignruleRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule request timeout response has a 5xx status code
func (o *GetOutboundCampaignruleRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule request timeout response a status code equal to that given
func (o *GetOutboundCampaignruleRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOutboundCampaignruleRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignruleRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundCampaignruleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleRequestEntityTooLarge creates a GetOutboundCampaignruleRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignruleRequestEntityTooLarge() *GetOutboundCampaignruleRequestEntityTooLarge {
	return &GetOutboundCampaignruleRequestEntityTooLarge{}
}

/*
GetOutboundCampaignruleRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOutboundCampaignruleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule request entity too large response has a 2xx status code
func (o *GetOutboundCampaignruleRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule request entity too large response has a 3xx status code
func (o *GetOutboundCampaignruleRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule request entity too large response has a 4xx status code
func (o *GetOutboundCampaignruleRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule request entity too large response has a 5xx status code
func (o *GetOutboundCampaignruleRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule request entity too large response a status code equal to that given
func (o *GetOutboundCampaignruleRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOutboundCampaignruleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignruleRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundCampaignruleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleUnsupportedMediaType creates a GetOutboundCampaignruleUnsupportedMediaType with default headers values
func NewGetOutboundCampaignruleUnsupportedMediaType() *GetOutboundCampaignruleUnsupportedMediaType {
	return &GetOutboundCampaignruleUnsupportedMediaType{}
}

/*
GetOutboundCampaignruleUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignruleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule unsupported media type response has a 2xx status code
func (o *GetOutboundCampaignruleUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule unsupported media type response has a 3xx status code
func (o *GetOutboundCampaignruleUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule unsupported media type response has a 4xx status code
func (o *GetOutboundCampaignruleUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule unsupported media type response has a 5xx status code
func (o *GetOutboundCampaignruleUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule unsupported media type response a status code equal to that given
func (o *GetOutboundCampaignruleUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOutboundCampaignruleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignruleUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundCampaignruleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleTooManyRequests creates a GetOutboundCampaignruleTooManyRequests with default headers values
func NewGetOutboundCampaignruleTooManyRequests() *GetOutboundCampaignruleTooManyRequests {
	return &GetOutboundCampaignruleTooManyRequests{}
}

/*
GetOutboundCampaignruleTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundCampaignruleTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule too many requests response has a 2xx status code
func (o *GetOutboundCampaignruleTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule too many requests response has a 3xx status code
func (o *GetOutboundCampaignruleTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule too many requests response has a 4xx status code
func (o *GetOutboundCampaignruleTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get outbound campaignrule too many requests response has a 5xx status code
func (o *GetOutboundCampaignruleTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get outbound campaignrule too many requests response a status code equal to that given
func (o *GetOutboundCampaignruleTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOutboundCampaignruleTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignruleTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundCampaignruleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleInternalServerError creates a GetOutboundCampaignruleInternalServerError with default headers values
func NewGetOutboundCampaignruleInternalServerError() *GetOutboundCampaignruleInternalServerError {
	return &GetOutboundCampaignruleInternalServerError{}
}

/*
GetOutboundCampaignruleInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignruleInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule internal server error response has a 2xx status code
func (o *GetOutboundCampaignruleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule internal server error response has a 3xx status code
func (o *GetOutboundCampaignruleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule internal server error response has a 4xx status code
func (o *GetOutboundCampaignruleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaignrule internal server error response has a 5xx status code
func (o *GetOutboundCampaignruleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaignrule internal server error response a status code equal to that given
func (o *GetOutboundCampaignruleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOutboundCampaignruleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignruleInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundCampaignruleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleServiceUnavailable creates a GetOutboundCampaignruleServiceUnavailable with default headers values
func NewGetOutboundCampaignruleServiceUnavailable() *GetOutboundCampaignruleServiceUnavailable {
	return &GetOutboundCampaignruleServiceUnavailable{}
}

/*
GetOutboundCampaignruleServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignruleServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule service unavailable response has a 2xx status code
func (o *GetOutboundCampaignruleServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule service unavailable response has a 3xx status code
func (o *GetOutboundCampaignruleServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule service unavailable response has a 4xx status code
func (o *GetOutboundCampaignruleServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaignrule service unavailable response has a 5xx status code
func (o *GetOutboundCampaignruleServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaignrule service unavailable response a status code equal to that given
func (o *GetOutboundCampaignruleServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOutboundCampaignruleServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignruleServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundCampaignruleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundCampaignruleGatewayTimeout creates a GetOutboundCampaignruleGatewayTimeout with default headers values
func NewGetOutboundCampaignruleGatewayTimeout() *GetOutboundCampaignruleGatewayTimeout {
	return &GetOutboundCampaignruleGatewayTimeout{}
}

/*
GetOutboundCampaignruleGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOutboundCampaignruleGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get outbound campaignrule gateway timeout response has a 2xx status code
func (o *GetOutboundCampaignruleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get outbound campaignrule gateway timeout response has a 3xx status code
func (o *GetOutboundCampaignruleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get outbound campaignrule gateway timeout response has a 4xx status code
func (o *GetOutboundCampaignruleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get outbound campaignrule gateway timeout response has a 5xx status code
func (o *GetOutboundCampaignruleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get outbound campaignrule gateway timeout response a status code equal to that given
func (o *GetOutboundCampaignruleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOutboundCampaignruleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignruleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/outbound/campaignrules/{campaignRuleId}][%d] getOutboundCampaignruleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundCampaignruleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundCampaignruleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
