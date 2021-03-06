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

/*GetOutboundRulesetsOK handles this case with default header values.

successful operation
*/
type GetOutboundRulesetsOK struct {
	Payload *models.RuleSetEntityListing
}

func (o *GetOutboundRulesetsOK) Error() string {
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

/*GetOutboundRulesetsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundRulesetsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsBadRequest) Error() string {
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

/*GetOutboundRulesetsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundRulesetsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsUnauthorized) Error() string {
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

/*GetOutboundRulesetsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundRulesetsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsForbidden) Error() string {
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

/*GetOutboundRulesetsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundRulesetsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsNotFound) Error() string {
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

// NewGetOutboundRulesetsRequestEntityTooLarge creates a GetOutboundRulesetsRequestEntityTooLarge with default headers values
func NewGetOutboundRulesetsRequestEntityTooLarge() *GetOutboundRulesetsRequestEntityTooLarge {
	return &GetOutboundRulesetsRequestEntityTooLarge{}
}

/*GetOutboundRulesetsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundRulesetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsRequestEntityTooLarge) Error() string {
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

/*GetOutboundRulesetsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundRulesetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsUnsupportedMediaType) Error() string {
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

/*GetOutboundRulesetsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundRulesetsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsTooManyRequests) Error() string {
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

/*GetOutboundRulesetsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundRulesetsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsInternalServerError) Error() string {
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

/*GetOutboundRulesetsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundRulesetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsServiceUnavailable) Error() string {
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

/*GetOutboundRulesetsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundRulesetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundRulesetsGatewayTimeout) Error() string {
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
