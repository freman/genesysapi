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

// PostAlertingInteractionstatsRulesReader is a Reader for the PostAlertingInteractionstatsRules structure.
type PostAlertingInteractionstatsRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAlertingInteractionstatsRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAlertingInteractionstatsRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAlertingInteractionstatsRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAlertingInteractionstatsRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAlertingInteractionstatsRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAlertingInteractionstatsRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAlertingInteractionstatsRulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAlertingInteractionstatsRulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAlertingInteractionstatsRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAlertingInteractionstatsRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAlertingInteractionstatsRulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAlertingInteractionstatsRulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAlertingInteractionstatsRulesOK creates a PostAlertingInteractionstatsRulesOK with default headers values
func NewPostAlertingInteractionstatsRulesOK() *PostAlertingInteractionstatsRulesOK {
	return &PostAlertingInteractionstatsRulesOK{}
}

/*PostAlertingInteractionstatsRulesOK handles this case with default header values.

successful operation
*/
type PostAlertingInteractionstatsRulesOK struct {
	Payload *models.InteractionStatsRule
}

func (o *PostAlertingInteractionstatsRulesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesOK  %+v", 200, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesOK) GetPayload() *models.InteractionStatsRule {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InteractionStatsRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesBadRequest creates a PostAlertingInteractionstatsRulesBadRequest with default headers values
func NewPostAlertingInteractionstatsRulesBadRequest() *PostAlertingInteractionstatsRulesBadRequest {
	return &PostAlertingInteractionstatsRulesBadRequest{}
}

/*PostAlertingInteractionstatsRulesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAlertingInteractionstatsRulesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesBadRequest  %+v", 400, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesUnauthorized creates a PostAlertingInteractionstatsRulesUnauthorized with default headers values
func NewPostAlertingInteractionstatsRulesUnauthorized() *PostAlertingInteractionstatsRulesUnauthorized {
	return &PostAlertingInteractionstatsRulesUnauthorized{}
}

/*PostAlertingInteractionstatsRulesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAlertingInteractionstatsRulesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesForbidden creates a PostAlertingInteractionstatsRulesForbidden with default headers values
func NewPostAlertingInteractionstatsRulesForbidden() *PostAlertingInteractionstatsRulesForbidden {
	return &PostAlertingInteractionstatsRulesForbidden{}
}

/*PostAlertingInteractionstatsRulesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAlertingInteractionstatsRulesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesForbidden  %+v", 403, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesNotFound creates a PostAlertingInteractionstatsRulesNotFound with default headers values
func NewPostAlertingInteractionstatsRulesNotFound() *PostAlertingInteractionstatsRulesNotFound {
	return &PostAlertingInteractionstatsRulesNotFound{}
}

/*PostAlertingInteractionstatsRulesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAlertingInteractionstatsRulesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesNotFound  %+v", 404, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesRequestEntityTooLarge creates a PostAlertingInteractionstatsRulesRequestEntityTooLarge with default headers values
func NewPostAlertingInteractionstatsRulesRequestEntityTooLarge() *PostAlertingInteractionstatsRulesRequestEntityTooLarge {
	return &PostAlertingInteractionstatsRulesRequestEntityTooLarge{}
}

/*PostAlertingInteractionstatsRulesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAlertingInteractionstatsRulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesUnsupportedMediaType creates a PostAlertingInteractionstatsRulesUnsupportedMediaType with default headers values
func NewPostAlertingInteractionstatsRulesUnsupportedMediaType() *PostAlertingInteractionstatsRulesUnsupportedMediaType {
	return &PostAlertingInteractionstatsRulesUnsupportedMediaType{}
}

/*PostAlertingInteractionstatsRulesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAlertingInteractionstatsRulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesTooManyRequests creates a PostAlertingInteractionstatsRulesTooManyRequests with default headers values
func NewPostAlertingInteractionstatsRulesTooManyRequests() *PostAlertingInteractionstatsRulesTooManyRequests {
	return &PostAlertingInteractionstatsRulesTooManyRequests{}
}

/*PostAlertingInteractionstatsRulesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAlertingInteractionstatsRulesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesInternalServerError creates a PostAlertingInteractionstatsRulesInternalServerError with default headers values
func NewPostAlertingInteractionstatsRulesInternalServerError() *PostAlertingInteractionstatsRulesInternalServerError {
	return &PostAlertingInteractionstatsRulesInternalServerError{}
}

/*PostAlertingInteractionstatsRulesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAlertingInteractionstatsRulesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesServiceUnavailable creates a PostAlertingInteractionstatsRulesServiceUnavailable with default headers values
func NewPostAlertingInteractionstatsRulesServiceUnavailable() *PostAlertingInteractionstatsRulesServiceUnavailable {
	return &PostAlertingInteractionstatsRulesServiceUnavailable{}
}

/*PostAlertingInteractionstatsRulesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAlertingInteractionstatsRulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAlertingInteractionstatsRulesGatewayTimeout creates a PostAlertingInteractionstatsRulesGatewayTimeout with default headers values
func NewPostAlertingInteractionstatsRulesGatewayTimeout() *PostAlertingInteractionstatsRulesGatewayTimeout {
	return &PostAlertingInteractionstatsRulesGatewayTimeout{}
}

/*PostAlertingInteractionstatsRulesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAlertingInteractionstatsRulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAlertingInteractionstatsRulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/alerting/interactionstats/rules][%d] postAlertingInteractionstatsRulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAlertingInteractionstatsRulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAlertingInteractionstatsRulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
