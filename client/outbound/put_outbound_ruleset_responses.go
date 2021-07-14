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

// PutOutboundRulesetReader is a Reader for the PutOutboundRuleset structure.
type PutOutboundRulesetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundRulesetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundRulesetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundRulesetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundRulesetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundRulesetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundRulesetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundRulesetRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundRulesetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundRulesetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundRulesetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundRulesetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundRulesetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundRulesetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundRulesetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundRulesetOK creates a PutOutboundRulesetOK with default headers values
func NewPutOutboundRulesetOK() *PutOutboundRulesetOK {
	return &PutOutboundRulesetOK{}
}

/*PutOutboundRulesetOK handles this case with default header values.

successful operation
*/
type PutOutboundRulesetOK struct {
	Payload *models.RuleSet
}

func (o *PutOutboundRulesetOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetOK  %+v", 200, o.Payload)
}

func (o *PutOutboundRulesetOK) GetPayload() *models.RuleSet {
	return o.Payload
}

func (o *PutOutboundRulesetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuleSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetBadRequest creates a PutOutboundRulesetBadRequest with default headers values
func NewPutOutboundRulesetBadRequest() *PutOutboundRulesetBadRequest {
	return &PutOutboundRulesetBadRequest{}
}

/*PutOutboundRulesetBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundRulesetBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundRulesetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetUnauthorized creates a PutOutboundRulesetUnauthorized with default headers values
func NewPutOutboundRulesetUnauthorized() *PutOutboundRulesetUnauthorized {
	return &PutOutboundRulesetUnauthorized{}
}

/*PutOutboundRulesetUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundRulesetUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundRulesetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetForbidden creates a PutOutboundRulesetForbidden with default headers values
func NewPutOutboundRulesetForbidden() *PutOutboundRulesetForbidden {
	return &PutOutboundRulesetForbidden{}
}

/*PutOutboundRulesetForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundRulesetForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundRulesetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetNotFound creates a PutOutboundRulesetNotFound with default headers values
func NewPutOutboundRulesetNotFound() *PutOutboundRulesetNotFound {
	return &PutOutboundRulesetNotFound{}
}

/*PutOutboundRulesetNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOutboundRulesetNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundRulesetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetRequestTimeout creates a PutOutboundRulesetRequestTimeout with default headers values
func NewPutOutboundRulesetRequestTimeout() *PutOutboundRulesetRequestTimeout {
	return &PutOutboundRulesetRequestTimeout{}
}

/*PutOutboundRulesetRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundRulesetRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundRulesetRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetConflict creates a PutOutboundRulesetConflict with default headers values
func NewPutOutboundRulesetConflict() *PutOutboundRulesetConflict {
	return &PutOutboundRulesetConflict{}
}

/*PutOutboundRulesetConflict handles this case with default header values.

Conflict
*/
type PutOutboundRulesetConflict struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundRulesetConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetRequestEntityTooLarge creates a PutOutboundRulesetRequestEntityTooLarge with default headers values
func NewPutOutboundRulesetRequestEntityTooLarge() *PutOutboundRulesetRequestEntityTooLarge {
	return &PutOutboundRulesetRequestEntityTooLarge{}
}

/*PutOutboundRulesetRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOutboundRulesetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundRulesetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetUnsupportedMediaType creates a PutOutboundRulesetUnsupportedMediaType with default headers values
func NewPutOutboundRulesetUnsupportedMediaType() *PutOutboundRulesetUnsupportedMediaType {
	return &PutOutboundRulesetUnsupportedMediaType{}
}

/*PutOutboundRulesetUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundRulesetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundRulesetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetTooManyRequests creates a PutOutboundRulesetTooManyRequests with default headers values
func NewPutOutboundRulesetTooManyRequests() *PutOutboundRulesetTooManyRequests {
	return &PutOutboundRulesetTooManyRequests{}
}

/*PutOutboundRulesetTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundRulesetTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundRulesetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetInternalServerError creates a PutOutboundRulesetInternalServerError with default headers values
func NewPutOutboundRulesetInternalServerError() *PutOutboundRulesetInternalServerError {
	return &PutOutboundRulesetInternalServerError{}
}

/*PutOutboundRulesetInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundRulesetInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundRulesetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetServiceUnavailable creates a PutOutboundRulesetServiceUnavailable with default headers values
func NewPutOutboundRulesetServiceUnavailable() *PutOutboundRulesetServiceUnavailable {
	return &PutOutboundRulesetServiceUnavailable{}
}

/*PutOutboundRulesetServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundRulesetServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundRulesetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundRulesetGatewayTimeout creates a PutOutboundRulesetGatewayTimeout with default headers values
func NewPutOutboundRulesetGatewayTimeout() *PutOutboundRulesetGatewayTimeout {
	return &PutOutboundRulesetGatewayTimeout{}
}

/*PutOutboundRulesetGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOutboundRulesetGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundRulesetGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/rulesets/{ruleSetId}][%d] putOutboundRulesetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundRulesetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundRulesetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
