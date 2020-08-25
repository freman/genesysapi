// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutTelephonyProvidersEdgesDidpoolReader is a Reader for the PutTelephonyProvidersEdgesDidpool structure.
type PutTelephonyProvidersEdgesDidpoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgesDidpoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgesDidpoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgesDidpoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgesDidpoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgesDidpoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgesDidpoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutTelephonyProvidersEdgesDidpoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgesDidpoolUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgesDidpoolTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgesDidpoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgesDidpoolServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgesDidpoolGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgesDidpoolOK creates a PutTelephonyProvidersEdgesDidpoolOK with default headers values
func NewPutTelephonyProvidersEdgesDidpoolOK() *PutTelephonyProvidersEdgesDidpoolOK {
	return &PutTelephonyProvidersEdgesDidpoolOK{}
}

/*PutTelephonyProvidersEdgesDidpoolOK handles this case with default header values.

successful operation
*/
type PutTelephonyProvidersEdgesDidpoolOK struct {
	Payload *models.DIDPool
}

func (o *PutTelephonyProvidersEdgesDidpoolOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolOK) GetPayload() *models.DIDPool {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DIDPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolBadRequest creates a PutTelephonyProvidersEdgesDidpoolBadRequest with default headers values
func NewPutTelephonyProvidersEdgesDidpoolBadRequest() *PutTelephonyProvidersEdgesDidpoolBadRequest {
	return &PutTelephonyProvidersEdgesDidpoolBadRequest{}
}

/*PutTelephonyProvidersEdgesDidpoolBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgesDidpoolBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolUnauthorized creates a PutTelephonyProvidersEdgesDidpoolUnauthorized with default headers values
func NewPutTelephonyProvidersEdgesDidpoolUnauthorized() *PutTelephonyProvidersEdgesDidpoolUnauthorized {
	return &PutTelephonyProvidersEdgesDidpoolUnauthorized{}
}

/*PutTelephonyProvidersEdgesDidpoolUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgesDidpoolUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolForbidden creates a PutTelephonyProvidersEdgesDidpoolForbidden with default headers values
func NewPutTelephonyProvidersEdgesDidpoolForbidden() *PutTelephonyProvidersEdgesDidpoolForbidden {
	return &PutTelephonyProvidersEdgesDidpoolForbidden{}
}

/*PutTelephonyProvidersEdgesDidpoolForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgesDidpoolForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolNotFound creates a PutTelephonyProvidersEdgesDidpoolNotFound with default headers values
func NewPutTelephonyProvidersEdgesDidpoolNotFound() *PutTelephonyProvidersEdgesDidpoolNotFound {
	return &PutTelephonyProvidersEdgesDidpoolNotFound{}
}

/*PutTelephonyProvidersEdgesDidpoolNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgesDidpoolNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolConflict creates a PutTelephonyProvidersEdgesDidpoolConflict with default headers values
func NewPutTelephonyProvidersEdgesDidpoolConflict() *PutTelephonyProvidersEdgesDidpoolConflict {
	return &PutTelephonyProvidersEdgesDidpoolConflict{}
}

/*PutTelephonyProvidersEdgesDidpoolConflict handles this case with default header values.

Conflict
*/
type PutTelephonyProvidersEdgesDidpoolConflict struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolConflict  %+v", 409, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge creates a PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge() *PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge{}
}

/*PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolUnsupportedMediaType creates a PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgesDidpoolUnsupportedMediaType() *PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType {
	return &PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType{}
}

/*PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolTooManyRequests creates a PutTelephonyProvidersEdgesDidpoolTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgesDidpoolTooManyRequests() *PutTelephonyProvidersEdgesDidpoolTooManyRequests {
	return &PutTelephonyProvidersEdgesDidpoolTooManyRequests{}
}

/*PutTelephonyProvidersEdgesDidpoolTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutTelephonyProvidersEdgesDidpoolTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolInternalServerError creates a PutTelephonyProvidersEdgesDidpoolInternalServerError with default headers values
func NewPutTelephonyProvidersEdgesDidpoolInternalServerError() *PutTelephonyProvidersEdgesDidpoolInternalServerError {
	return &PutTelephonyProvidersEdgesDidpoolInternalServerError{}
}

/*PutTelephonyProvidersEdgesDidpoolInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgesDidpoolInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolServiceUnavailable creates a PutTelephonyProvidersEdgesDidpoolServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgesDidpoolServiceUnavailable() *PutTelephonyProvidersEdgesDidpoolServiceUnavailable {
	return &PutTelephonyProvidersEdgesDidpoolServiceUnavailable{}
}

/*PutTelephonyProvidersEdgesDidpoolServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgesDidpoolServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesDidpoolGatewayTimeout creates a PutTelephonyProvidersEdgesDidpoolGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgesDidpoolGatewayTimeout() *PutTelephonyProvidersEdgesDidpoolGatewayTimeout {
	return &PutTelephonyProvidersEdgesDidpoolGatewayTimeout{}
}

/*PutTelephonyProvidersEdgesDidpoolGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgesDidpoolGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesDidpoolGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/didpools/{didPoolId}][%d] putTelephonyProvidersEdgesDidpoolGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgesDidpoolGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesDidpoolGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}