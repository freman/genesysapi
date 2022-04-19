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

// DeleteAlertingInteractionstatsRuleReader is a Reader for the DeleteAlertingInteractionstatsRule structure.
type DeleteAlertingInteractionstatsRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertingInteractionstatsRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAlertingInteractionstatsRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAlertingInteractionstatsRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAlertingInteractionstatsRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAlertingInteractionstatsRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAlertingInteractionstatsRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteAlertingInteractionstatsRuleRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteAlertingInteractionstatsRuleRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteAlertingInteractionstatsRuleUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAlertingInteractionstatsRuleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAlertingInteractionstatsRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAlertingInteractionstatsRuleServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteAlertingInteractionstatsRuleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAlertingInteractionstatsRuleNoContent creates a DeleteAlertingInteractionstatsRuleNoContent with default headers values
func NewDeleteAlertingInteractionstatsRuleNoContent() *DeleteAlertingInteractionstatsRuleNoContent {
	return &DeleteAlertingInteractionstatsRuleNoContent{}
}

/*DeleteAlertingInteractionstatsRuleNoContent handles this case with default header values.

Interaction stats rule deleted
*/
type DeleteAlertingInteractionstatsRuleNoContent struct {
}

func (o *DeleteAlertingInteractionstatsRuleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleNoContent ", 204)
}

func (o *DeleteAlertingInteractionstatsRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAlertingInteractionstatsRuleBadRequest creates a DeleteAlertingInteractionstatsRuleBadRequest with default headers values
func NewDeleteAlertingInteractionstatsRuleBadRequest() *DeleteAlertingInteractionstatsRuleBadRequest {
	return &DeleteAlertingInteractionstatsRuleBadRequest{}
}

/*DeleteAlertingInteractionstatsRuleBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteAlertingInteractionstatsRuleBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleUnauthorized creates a DeleteAlertingInteractionstatsRuleUnauthorized with default headers values
func NewDeleteAlertingInteractionstatsRuleUnauthorized() *DeleteAlertingInteractionstatsRuleUnauthorized {
	return &DeleteAlertingInteractionstatsRuleUnauthorized{}
}

/*DeleteAlertingInteractionstatsRuleUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteAlertingInteractionstatsRuleUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleForbidden creates a DeleteAlertingInteractionstatsRuleForbidden with default headers values
func NewDeleteAlertingInteractionstatsRuleForbidden() *DeleteAlertingInteractionstatsRuleForbidden {
	return &DeleteAlertingInteractionstatsRuleForbidden{}
}

/*DeleteAlertingInteractionstatsRuleForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteAlertingInteractionstatsRuleForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleNotFound creates a DeleteAlertingInteractionstatsRuleNotFound with default headers values
func NewDeleteAlertingInteractionstatsRuleNotFound() *DeleteAlertingInteractionstatsRuleNotFound {
	return &DeleteAlertingInteractionstatsRuleNotFound{}
}

/*DeleteAlertingInteractionstatsRuleNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteAlertingInteractionstatsRuleNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleRequestTimeout creates a DeleteAlertingInteractionstatsRuleRequestTimeout with default headers values
func NewDeleteAlertingInteractionstatsRuleRequestTimeout() *DeleteAlertingInteractionstatsRuleRequestTimeout {
	return &DeleteAlertingInteractionstatsRuleRequestTimeout{}
}

/*DeleteAlertingInteractionstatsRuleRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteAlertingInteractionstatsRuleRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleRequestEntityTooLarge creates a DeleteAlertingInteractionstatsRuleRequestEntityTooLarge with default headers values
func NewDeleteAlertingInteractionstatsRuleRequestEntityTooLarge() *DeleteAlertingInteractionstatsRuleRequestEntityTooLarge {
	return &DeleteAlertingInteractionstatsRuleRequestEntityTooLarge{}
}

/*DeleteAlertingInteractionstatsRuleRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteAlertingInteractionstatsRuleRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleUnsupportedMediaType creates a DeleteAlertingInteractionstatsRuleUnsupportedMediaType with default headers values
func NewDeleteAlertingInteractionstatsRuleUnsupportedMediaType() *DeleteAlertingInteractionstatsRuleUnsupportedMediaType {
	return &DeleteAlertingInteractionstatsRuleUnsupportedMediaType{}
}

/*DeleteAlertingInteractionstatsRuleUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteAlertingInteractionstatsRuleUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleTooManyRequests creates a DeleteAlertingInteractionstatsRuleTooManyRequests with default headers values
func NewDeleteAlertingInteractionstatsRuleTooManyRequests() *DeleteAlertingInteractionstatsRuleTooManyRequests {
	return &DeleteAlertingInteractionstatsRuleTooManyRequests{}
}

/*DeleteAlertingInteractionstatsRuleTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteAlertingInteractionstatsRuleTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleInternalServerError creates a DeleteAlertingInteractionstatsRuleInternalServerError with default headers values
func NewDeleteAlertingInteractionstatsRuleInternalServerError() *DeleteAlertingInteractionstatsRuleInternalServerError {
	return &DeleteAlertingInteractionstatsRuleInternalServerError{}
}

/*DeleteAlertingInteractionstatsRuleInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteAlertingInteractionstatsRuleInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleServiceUnavailable creates a DeleteAlertingInteractionstatsRuleServiceUnavailable with default headers values
func NewDeleteAlertingInteractionstatsRuleServiceUnavailable() *DeleteAlertingInteractionstatsRuleServiceUnavailable {
	return &DeleteAlertingInteractionstatsRuleServiceUnavailable{}
}

/*DeleteAlertingInteractionstatsRuleServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteAlertingInteractionstatsRuleServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertingInteractionstatsRuleGatewayTimeout creates a DeleteAlertingInteractionstatsRuleGatewayTimeout with default headers values
func NewDeleteAlertingInteractionstatsRuleGatewayTimeout() *DeleteAlertingInteractionstatsRuleGatewayTimeout {
	return &DeleteAlertingInteractionstatsRuleGatewayTimeout{}
}

/*DeleteAlertingInteractionstatsRuleGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteAlertingInteractionstatsRuleGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAlertingInteractionstatsRuleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/alerting/interactionstats/rules/{ruleId}][%d] deleteAlertingInteractionstatsRuleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteAlertingInteractionstatsRuleGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAlertingInteractionstatsRuleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
