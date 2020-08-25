// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIdentityprovidersIdentitynowReader is a Reader for the GetIdentityprovidersIdentitynow structure.
type GetIdentityprovidersIdentitynowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityprovidersIdentitynowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityprovidersIdentitynowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdentityprovidersIdentitynowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdentityprovidersIdentitynowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdentityprovidersIdentitynowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityprovidersIdentitynowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIdentityprovidersIdentitynowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIdentityprovidersIdentitynowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdentityprovidersIdentitynowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityprovidersIdentitynowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIdentityprovidersIdentitynowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIdentityprovidersIdentitynowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdentityprovidersIdentitynowOK creates a GetIdentityprovidersIdentitynowOK with default headers values
func NewGetIdentityprovidersIdentitynowOK() *GetIdentityprovidersIdentitynowOK {
	return &GetIdentityprovidersIdentitynowOK{}
}

/*GetIdentityprovidersIdentitynowOK handles this case with default header values.

successful operation
*/
type GetIdentityprovidersIdentitynowOK struct {
	Payload *models.IdentityNow
}

func (o *GetIdentityprovidersIdentitynowOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowOK  %+v", 200, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowOK) GetPayload() *models.IdentityNow {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityNow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowBadRequest creates a GetIdentityprovidersIdentitynowBadRequest with default headers values
func NewGetIdentityprovidersIdentitynowBadRequest() *GetIdentityprovidersIdentitynowBadRequest {
	return &GetIdentityprovidersIdentitynowBadRequest{}
}

/*GetIdentityprovidersIdentitynowBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIdentityprovidersIdentitynowBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowUnauthorized creates a GetIdentityprovidersIdentitynowUnauthorized with default headers values
func NewGetIdentityprovidersIdentitynowUnauthorized() *GetIdentityprovidersIdentitynowUnauthorized {
	return &GetIdentityprovidersIdentitynowUnauthorized{}
}

/*GetIdentityprovidersIdentitynowUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIdentityprovidersIdentitynowUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowForbidden creates a GetIdentityprovidersIdentitynowForbidden with default headers values
func NewGetIdentityprovidersIdentitynowForbidden() *GetIdentityprovidersIdentitynowForbidden {
	return &GetIdentityprovidersIdentitynowForbidden{}
}

/*GetIdentityprovidersIdentitynowForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIdentityprovidersIdentitynowForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowForbidden  %+v", 403, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowNotFound creates a GetIdentityprovidersIdentitynowNotFound with default headers values
func NewGetIdentityprovidersIdentitynowNotFound() *GetIdentityprovidersIdentitynowNotFound {
	return &GetIdentityprovidersIdentitynowNotFound{}
}

/*GetIdentityprovidersIdentitynowNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIdentityprovidersIdentitynowNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowRequestEntityTooLarge creates a GetIdentityprovidersIdentitynowRequestEntityTooLarge with default headers values
func NewGetIdentityprovidersIdentitynowRequestEntityTooLarge() *GetIdentityprovidersIdentitynowRequestEntityTooLarge {
	return &GetIdentityprovidersIdentitynowRequestEntityTooLarge{}
}

/*GetIdentityprovidersIdentitynowRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIdentityprovidersIdentitynowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowUnsupportedMediaType creates a GetIdentityprovidersIdentitynowUnsupportedMediaType with default headers values
func NewGetIdentityprovidersIdentitynowUnsupportedMediaType() *GetIdentityprovidersIdentitynowUnsupportedMediaType {
	return &GetIdentityprovidersIdentitynowUnsupportedMediaType{}
}

/*GetIdentityprovidersIdentitynowUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIdentityprovidersIdentitynowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowTooManyRequests creates a GetIdentityprovidersIdentitynowTooManyRequests with default headers values
func NewGetIdentityprovidersIdentitynowTooManyRequests() *GetIdentityprovidersIdentitynowTooManyRequests {
	return &GetIdentityprovidersIdentitynowTooManyRequests{}
}

/*GetIdentityprovidersIdentitynowTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIdentityprovidersIdentitynowTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowInternalServerError creates a GetIdentityprovidersIdentitynowInternalServerError with default headers values
func NewGetIdentityprovidersIdentitynowInternalServerError() *GetIdentityprovidersIdentitynowInternalServerError {
	return &GetIdentityprovidersIdentitynowInternalServerError{}
}

/*GetIdentityprovidersIdentitynowInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIdentityprovidersIdentitynowInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowServiceUnavailable creates a GetIdentityprovidersIdentitynowServiceUnavailable with default headers values
func NewGetIdentityprovidersIdentitynowServiceUnavailable() *GetIdentityprovidersIdentitynowServiceUnavailable {
	return &GetIdentityprovidersIdentitynowServiceUnavailable{}
}

/*GetIdentityprovidersIdentitynowServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIdentityprovidersIdentitynowServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersIdentitynowGatewayTimeout creates a GetIdentityprovidersIdentitynowGatewayTimeout with default headers values
func NewGetIdentityprovidersIdentitynowGatewayTimeout() *GetIdentityprovidersIdentitynowGatewayTimeout {
	return &GetIdentityprovidersIdentitynowGatewayTimeout{}
}

/*GetIdentityprovidersIdentitynowGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIdentityprovidersIdentitynowGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersIdentitynowGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/identitynow][%d] getIdentityprovidersIdentitynowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIdentityprovidersIdentitynowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersIdentitynowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}