// Code generated by go-swagger; DO NOT EDIT.

package alerting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutAlertingInteractionstatsRuleReader is a Reader for the PutAlertingInteractionstatsRule structure.
type PutAlertingInteractionstatsRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAlertingInteractionstatsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAlertingInteractionstatsRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAlertingInteractionstatsRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAlertingInteractionstatsRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAlertingInteractionstatsRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAlertingInteractionstatsRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutAlertingInteractionstatsRuleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutAlertingInteractionstatsRuleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutAlertingInteractionstatsRuleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAlertingInteractionstatsRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutAlertingInteractionstatsRuleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutAlertingInteractionstatsRuleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAlertingInteractionstatsRuleOK creates a PutAlertingInteractionstatsRuleOK with default headers values
func NewPutAlertingInteractionstatsRuleOK() *PutAlertingInteractionstatsRuleOK {
	return &PutAlertingInteractionstatsRuleOK{}
}

/*PutAlertingInteractionstatsRuleOK handles this case with default header values.

successful operation
*/
type PutAlertingInteractionstatsRuleOK struct {
	Payload *models.InteractionStatsRule
}

func (o *PutAlertingInteractionstatsRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleOK  %+v", 200, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleOK) GetPayload() *models.InteractionStatsRule {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InteractionStatsRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleBadRequest creates a PutAlertingInteractionstatsRuleBadRequest with default headers values
func NewPutAlertingInteractionstatsRuleBadRequest() *PutAlertingInteractionstatsRuleBadRequest {
	return &PutAlertingInteractionstatsRuleBadRequest{}
}

/*PutAlertingInteractionstatsRuleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutAlertingInteractionstatsRuleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleBadRequest  %+v", 400, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleUnauthorized creates a PutAlertingInteractionstatsRuleUnauthorized with default headers values
func NewPutAlertingInteractionstatsRuleUnauthorized() *PutAlertingInteractionstatsRuleUnauthorized {
	return &PutAlertingInteractionstatsRuleUnauthorized{}
}

/*PutAlertingInteractionstatsRuleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutAlertingInteractionstatsRuleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleForbidden creates a PutAlertingInteractionstatsRuleForbidden with default headers values
func NewPutAlertingInteractionstatsRuleForbidden() *PutAlertingInteractionstatsRuleForbidden {
	return &PutAlertingInteractionstatsRuleForbidden{}
}

/*PutAlertingInteractionstatsRuleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutAlertingInteractionstatsRuleForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleForbidden  %+v", 403, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleNotFound creates a PutAlertingInteractionstatsRuleNotFound with default headers values
func NewPutAlertingInteractionstatsRuleNotFound() *PutAlertingInteractionstatsRuleNotFound {
	return &PutAlertingInteractionstatsRuleNotFound{}
}

/*PutAlertingInteractionstatsRuleNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutAlertingInteractionstatsRuleNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleNotFound  %+v", 404, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleRequestEntityTooLarge creates a PutAlertingInteractionstatsRuleRequestEntityTooLarge with default headers values
func NewPutAlertingInteractionstatsRuleRequestEntityTooLarge() *PutAlertingInteractionstatsRuleRequestEntityTooLarge {
	return &PutAlertingInteractionstatsRuleRequestEntityTooLarge{}
}

/*PutAlertingInteractionstatsRuleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutAlertingInteractionstatsRuleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleUnsupportedMediaType creates a PutAlertingInteractionstatsRuleUnsupportedMediaType with default headers values
func NewPutAlertingInteractionstatsRuleUnsupportedMediaType() *PutAlertingInteractionstatsRuleUnsupportedMediaType {
	return &PutAlertingInteractionstatsRuleUnsupportedMediaType{}
}

/*PutAlertingInteractionstatsRuleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutAlertingInteractionstatsRuleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleTooManyRequests creates a PutAlertingInteractionstatsRuleTooManyRequests with default headers values
func NewPutAlertingInteractionstatsRuleTooManyRequests() *PutAlertingInteractionstatsRuleTooManyRequests {
	return &PutAlertingInteractionstatsRuleTooManyRequests{}
}

/*PutAlertingInteractionstatsRuleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutAlertingInteractionstatsRuleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleInternalServerError creates a PutAlertingInteractionstatsRuleInternalServerError with default headers values
func NewPutAlertingInteractionstatsRuleInternalServerError() *PutAlertingInteractionstatsRuleInternalServerError {
	return &PutAlertingInteractionstatsRuleInternalServerError{}
}

/*PutAlertingInteractionstatsRuleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutAlertingInteractionstatsRuleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleServiceUnavailable creates a PutAlertingInteractionstatsRuleServiceUnavailable with default headers values
func NewPutAlertingInteractionstatsRuleServiceUnavailable() *PutAlertingInteractionstatsRuleServiceUnavailable {
	return &PutAlertingInteractionstatsRuleServiceUnavailable{}
}

/*PutAlertingInteractionstatsRuleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutAlertingInteractionstatsRuleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertingInteractionstatsRuleGatewayTimeout creates a PutAlertingInteractionstatsRuleGatewayTimeout with default headers values
func NewPutAlertingInteractionstatsRuleGatewayTimeout() *PutAlertingInteractionstatsRuleGatewayTimeout {
	return &PutAlertingInteractionstatsRuleGatewayTimeout{}
}

/*PutAlertingInteractionstatsRuleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutAlertingInteractionstatsRuleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutAlertingInteractionstatsRuleGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/alerting/interactionstats/rules/{ruleId}][%d] putAlertingInteractionstatsRuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutAlertingInteractionstatsRuleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAlertingInteractionstatsRuleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}