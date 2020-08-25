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

/*GetOutboundCampaignruleOK handles this case with default header values.

successful operation
*/
type GetOutboundCampaignruleOK struct {
	Payload *models.CampaignRule
}

func (o *GetOutboundCampaignruleOK) Error() string {
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

/*GetOutboundCampaignruleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundCampaignruleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleBadRequest) Error() string {
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

/*GetOutboundCampaignruleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundCampaignruleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleUnauthorized) Error() string {
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

/*GetOutboundCampaignruleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundCampaignruleForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleForbidden) Error() string {
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

/*GetOutboundCampaignruleNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundCampaignruleNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleNotFound) Error() string {
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

// NewGetOutboundCampaignruleRequestEntityTooLarge creates a GetOutboundCampaignruleRequestEntityTooLarge with default headers values
func NewGetOutboundCampaignruleRequestEntityTooLarge() *GetOutboundCampaignruleRequestEntityTooLarge {
	return &GetOutboundCampaignruleRequestEntityTooLarge{}
}

/*GetOutboundCampaignruleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundCampaignruleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleRequestEntityTooLarge) Error() string {
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

/*GetOutboundCampaignruleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundCampaignruleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleUnsupportedMediaType) Error() string {
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

/*GetOutboundCampaignruleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundCampaignruleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleTooManyRequests) Error() string {
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

/*GetOutboundCampaignruleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundCampaignruleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleInternalServerError) Error() string {
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

/*GetOutboundCampaignruleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundCampaignruleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleServiceUnavailable) Error() string {
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

/*GetOutboundCampaignruleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundCampaignruleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundCampaignruleGatewayTimeout) Error() string {
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