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

// PutOutboundCallabletimesetReader is a Reader for the PutOutboundCallabletimeset structure.
type PutOutboundCallabletimesetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundCallabletimesetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundCallabletimesetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundCallabletimesetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundCallabletimesetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundCallabletimesetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundCallabletimesetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundCallabletimesetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundCallabletimesetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundCallabletimesetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundCallabletimesetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundCallabletimesetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundCallabletimesetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundCallabletimesetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundCallabletimesetOK creates a PutOutboundCallabletimesetOK with default headers values
func NewPutOutboundCallabletimesetOK() *PutOutboundCallabletimesetOK {
	return &PutOutboundCallabletimesetOK{}
}

/*PutOutboundCallabletimesetOK handles this case with default header values.

successful operation
*/
type PutOutboundCallabletimesetOK struct {
	Payload *models.CallableTimeSet
}

func (o *PutOutboundCallabletimesetOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetOK  %+v", 200, o.Payload)
}

func (o *PutOutboundCallabletimesetOK) GetPayload() *models.CallableTimeSet {
	return o.Payload
}

func (o *PutOutboundCallabletimesetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallableTimeSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetBadRequest creates a PutOutboundCallabletimesetBadRequest with default headers values
func NewPutOutboundCallabletimesetBadRequest() *PutOutboundCallabletimesetBadRequest {
	return &PutOutboundCallabletimesetBadRequest{}
}

/*PutOutboundCallabletimesetBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundCallabletimesetBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundCallabletimesetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetUnauthorized creates a PutOutboundCallabletimesetUnauthorized with default headers values
func NewPutOutboundCallabletimesetUnauthorized() *PutOutboundCallabletimesetUnauthorized {
	return &PutOutboundCallabletimesetUnauthorized{}
}

/*PutOutboundCallabletimesetUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundCallabletimesetUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundCallabletimesetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetForbidden creates a PutOutboundCallabletimesetForbidden with default headers values
func NewPutOutboundCallabletimesetForbidden() *PutOutboundCallabletimesetForbidden {
	return &PutOutboundCallabletimesetForbidden{}
}

/*PutOutboundCallabletimesetForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundCallabletimesetForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundCallabletimesetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetNotFound creates a PutOutboundCallabletimesetNotFound with default headers values
func NewPutOutboundCallabletimesetNotFound() *PutOutboundCallabletimesetNotFound {
	return &PutOutboundCallabletimesetNotFound{}
}

/*PutOutboundCallabletimesetNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOutboundCallabletimesetNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundCallabletimesetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetConflict creates a PutOutboundCallabletimesetConflict with default headers values
func NewPutOutboundCallabletimesetConflict() *PutOutboundCallabletimesetConflict {
	return &PutOutboundCallabletimesetConflict{}
}

/*PutOutboundCallabletimesetConflict handles this case with default header values.

Conflict
*/
type PutOutboundCallabletimesetConflict struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundCallabletimesetConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetRequestEntityTooLarge creates a PutOutboundCallabletimesetRequestEntityTooLarge with default headers values
func NewPutOutboundCallabletimesetRequestEntityTooLarge() *PutOutboundCallabletimesetRequestEntityTooLarge {
	return &PutOutboundCallabletimesetRequestEntityTooLarge{}
}

/*PutOutboundCallabletimesetRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutOutboundCallabletimesetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundCallabletimesetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetUnsupportedMediaType creates a PutOutboundCallabletimesetUnsupportedMediaType with default headers values
func NewPutOutboundCallabletimesetUnsupportedMediaType() *PutOutboundCallabletimesetUnsupportedMediaType {
	return &PutOutboundCallabletimesetUnsupportedMediaType{}
}

/*PutOutboundCallabletimesetUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundCallabletimesetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundCallabletimesetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetTooManyRequests creates a PutOutboundCallabletimesetTooManyRequests with default headers values
func NewPutOutboundCallabletimesetTooManyRequests() *PutOutboundCallabletimesetTooManyRequests {
	return &PutOutboundCallabletimesetTooManyRequests{}
}

/*PutOutboundCallabletimesetTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutOutboundCallabletimesetTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundCallabletimesetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetInternalServerError creates a PutOutboundCallabletimesetInternalServerError with default headers values
func NewPutOutboundCallabletimesetInternalServerError() *PutOutboundCallabletimesetInternalServerError {
	return &PutOutboundCallabletimesetInternalServerError{}
}

/*PutOutboundCallabletimesetInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundCallabletimesetInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundCallabletimesetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetServiceUnavailable creates a PutOutboundCallabletimesetServiceUnavailable with default headers values
func NewPutOutboundCallabletimesetServiceUnavailable() *PutOutboundCallabletimesetServiceUnavailable {
	return &PutOutboundCallabletimesetServiceUnavailable{}
}

/*PutOutboundCallabletimesetServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundCallabletimesetServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundCallabletimesetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundCallabletimesetGatewayTimeout creates a PutOutboundCallabletimesetGatewayTimeout with default headers values
func NewPutOutboundCallabletimesetGatewayTimeout() *PutOutboundCallabletimesetGatewayTimeout {
	return &PutOutboundCallabletimesetGatewayTimeout{}
}

/*PutOutboundCallabletimesetGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOutboundCallabletimesetGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundCallabletimesetGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/callabletimesets/{callableTimeSetId}][%d] putOutboundCallabletimesetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundCallabletimesetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundCallabletimesetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
