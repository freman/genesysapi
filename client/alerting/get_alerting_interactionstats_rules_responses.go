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

/*
GetAlertingInteractionstatsRulesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAlertingInteractionstatsRulesOK struct {
	Payload *models.InteractionStatsRuleContainer
}

// IsSuccess returns true when this get alerting interactionstats rules o k response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alerting interactionstats rules o k response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules o k response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats rules o k response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules o k response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAlertingInteractionstatsRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesOK  %+v", 200, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesOK) String() string {
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

/*
GetAlertingInteractionstatsRulesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAlertingInteractionstatsRulesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules bad request response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules bad request response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules bad request response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules bad request response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules bad request response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAlertingInteractionstatsRulesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesBadRequest) String() string {
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

/*
GetAlertingInteractionstatsRulesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAlertingInteractionstatsRulesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules unauthorized response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules unauthorized response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules unauthorized response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules unauthorized response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules unauthorized response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAlertingInteractionstatsRulesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesUnauthorized) String() string {
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

/*
GetAlertingInteractionstatsRulesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAlertingInteractionstatsRulesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules forbidden response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules forbidden response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules forbidden response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules forbidden response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules forbidden response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAlertingInteractionstatsRulesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesForbidden  %+v", 403, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesForbidden) String() string {
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

/*
GetAlertingInteractionstatsRulesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAlertingInteractionstatsRulesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules not found response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules not found response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules not found response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules not found response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules not found response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAlertingInteractionstatsRulesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesNotFound  %+v", 404, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesNotFound) String() string {
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

/*
GetAlertingInteractionstatsRulesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAlertingInteractionstatsRulesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules request timeout response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules request timeout response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules request timeout response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules request timeout response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules request timeout response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAlertingInteractionstatsRulesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesRequestTimeout) String() string {
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

/*
GetAlertingInteractionstatsRulesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAlertingInteractionstatsRulesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules request entity too large response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules request entity too large response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules request entity too large response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules request entity too large response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules request entity too large response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesRequestEntityTooLarge) String() string {
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

/*
GetAlertingInteractionstatsRulesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAlertingInteractionstatsRulesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules unsupported media type response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules unsupported media type response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules unsupported media type response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules unsupported media type response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules unsupported media type response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesUnsupportedMediaType) String() string {
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

/*
GetAlertingInteractionstatsRulesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAlertingInteractionstatsRulesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules too many requests response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules too many requests response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules too many requests response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alerting interactionstats rules too many requests response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get alerting interactionstats rules too many requests response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAlertingInteractionstatsRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesTooManyRequests) String() string {
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

/*
GetAlertingInteractionstatsRulesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAlertingInteractionstatsRulesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules internal server error response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules internal server error response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules internal server error response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats rules internal server error response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats rules internal server error response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAlertingInteractionstatsRulesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesInternalServerError) String() string {
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

/*
GetAlertingInteractionstatsRulesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAlertingInteractionstatsRulesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules service unavailable response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules service unavailable response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules service unavailable response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats rules service unavailable response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats rules service unavailable response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAlertingInteractionstatsRulesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesServiceUnavailable) String() string {
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

/*
GetAlertingInteractionstatsRulesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAlertingInteractionstatsRulesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get alerting interactionstats rules gateway timeout response has a 2xx status code
func (o *GetAlertingInteractionstatsRulesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alerting interactionstats rules gateway timeout response has a 3xx status code
func (o *GetAlertingInteractionstatsRulesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alerting interactionstats rules gateway timeout response has a 4xx status code
func (o *GetAlertingInteractionstatsRulesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alerting interactionstats rules gateway timeout response has a 5xx status code
func (o *GetAlertingInteractionstatsRulesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get alerting interactionstats rules gateway timeout response a status code equal to that given
func (o *GetAlertingInteractionstatsRulesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAlertingInteractionstatsRulesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/alerting/interactionstats/rules][%d] getAlertingInteractionstatsRulesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlertingInteractionstatsRulesGatewayTimeout) String() string {
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
