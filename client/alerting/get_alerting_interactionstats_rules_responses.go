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

// GetAlertingInteractionstatsRulesReader is a Reader for the GetAlertingInteractionstatsRules structure.
type GetAlertingInteractionstatsRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertingInteractionstatsRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertingInteractionstatsRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertingInteractionstatsRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAlertingInteractionstatsRulesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAlertingInteractionstatsRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAlertingInteractionstatsRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAlertingInteractionstatsRulesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAlertingInteractionstatsRulesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAlertingInteractionstatsRulesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAlertingInteractionstatsRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAlertingInteractionstatsRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAlertingInteractionstatsRulesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAlertingInteractionstatsRulesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertingInteractionstatsRulesOK creates a GetAlertingInteractionstatsRulesOK with default headers values
func NewGetAlertingInteractionstatsRulesOK() *GetAlertingInteractionstatsRulesOK {
	return &GetAlertingInteractionstatsRulesOK{}
}

/*GetAlertingInteractionstatsRulesOK handles this case with default header values.

successful operation
*/
type GetAlertingInteractionstatsRulesOK struct {
	Payload *models.InteractionStatsRuleContainer
}

func (o *GetAlertingInteractionstatsRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesOK) GetPayload() *models.InteractionStatsRuleContainer {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InteractionStatsRuleContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesBadRequest creates a GetAlertingInteractionstatsRulesBadRequest with default headers values
func NewGetAlertingInteractionstatsRulesBadRequest() *GetAlertingInteractionstatsRulesBadRequest {
	return &GetAlertingInteractionstatsRulesBadRequest{}
}

/*GetAlertingInteractionstatsRulesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingInteractionstatsRulesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesUnauthorized creates a GetAlertingInteractionstatsRulesUnauthorized with default headers values
func NewGetAlertingInteractionstatsRulesUnauthorized() *GetAlertingInteractionstatsRulesUnauthorized {
	return &GetAlertingInteractionstatsRulesUnauthorized{}
}

/*GetAlertingInteractionstatsRulesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingInteractionstatsRulesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesForbidden creates a GetAlertingInteractionstatsRulesForbidden with default headers values
func NewGetAlertingInteractionstatsRulesForbidden() *GetAlertingInteractionstatsRulesForbidden {
	return &GetAlertingInteractionstatsRulesForbidden{}
}

/*GetAlertingInteractionstatsRulesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingInteractionstatsRulesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesNotFound creates a GetAlertingInteractionstatsRulesNotFound with default headers values
func NewGetAlertingInteractionstatsRulesNotFound() *GetAlertingInteractionstatsRulesNotFound {
	return &GetAlertingInteractionstatsRulesNotFound{}
}

/*GetAlertingInteractionstatsRulesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAlertingInteractionstatsRulesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesRequestTimeout creates a GetAlertingInteractionstatsRulesRequestTimeout with default headers values
func NewGetAlertingInteractionstatsRulesRequestTimeout() *GetAlertingInteractionstatsRulesRequestTimeout {
	return &GetAlertingInteractionstatsRulesRequestTimeout{}
}

/*GetAlertingInteractionstatsRulesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAlertingInteractionstatsRulesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesRequestEntityTooLarge creates a GetAlertingInteractionstatsRulesRequestEntityTooLarge with default headers values
func NewGetAlertingInteractionstatsRulesRequestEntityTooLarge() *GetAlertingInteractionstatsRulesRequestEntityTooLarge {
	return &GetAlertingInteractionstatsRulesRequestEntityTooLarge{}
}

/*GetAlertingInteractionstatsRulesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAlertingInteractionstatsRulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesUnsupportedMediaType creates a GetAlertingInteractionstatsRulesUnsupportedMediaType with default headers values
func NewGetAlertingInteractionstatsRulesUnsupportedMediaType() *GetAlertingInteractionstatsRulesUnsupportedMediaType {
	return &GetAlertingInteractionstatsRulesUnsupportedMediaType{}
}

/*GetAlertingInteractionstatsRulesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingInteractionstatsRulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesTooManyRequests creates a GetAlertingInteractionstatsRulesTooManyRequests with default headers values
func NewGetAlertingInteractionstatsRulesTooManyRequests() *GetAlertingInteractionstatsRulesTooManyRequests {
	return &GetAlertingInteractionstatsRulesTooManyRequests{}
}

/*GetAlertingInteractionstatsRulesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAlertingInteractionstatsRulesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesInternalServerError creates a GetAlertingInteractionstatsRulesInternalServerError with default headers values
func NewGetAlertingInteractionstatsRulesInternalServerError() *GetAlertingInteractionstatsRulesInternalServerError {
	return &GetAlertingInteractionstatsRulesInternalServerError{}
}

/*GetAlertingInteractionstatsRulesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingInteractionstatsRulesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesServiceUnavailable creates a GetAlertingInteractionstatsRulesServiceUnavailable with default headers values
func NewGetAlertingInteractionstatsRulesServiceUnavailable() *GetAlertingInteractionstatsRulesServiceUnavailable {
	return &GetAlertingInteractionstatsRulesServiceUnavailable{}
}

/*GetAlertingInteractionstatsRulesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingInteractionstatsRulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertingInteractionstatsRulesGatewayTimeout creates a GetAlertingInteractionstatsRulesGatewayTimeout with default headers values
func NewGetAlertingInteractionstatsRulesGatewayTimeout() *GetAlertingInteractionstatsRulesGatewayTimeout {
	return &GetAlertingInteractionstatsRulesGatewayTimeout{}
}

/*GetAlertingInteractionstatsRulesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAlertingInteractionstatsRulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAlertingInteractionstatsRulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAlertingInteractionstatsRulesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
