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

// GetIdentityprovidersGenericReader is a Reader for the GetIdentityprovidersGeneric structure.
type GetIdentityprovidersGenericReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityprovidersGenericReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityprovidersGenericOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdentityprovidersGenericBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdentityprovidersGenericUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdentityprovidersGenericForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityprovidersGenericNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIdentityprovidersGenericRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIdentityprovidersGenericRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIdentityprovidersGenericUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdentityprovidersGenericTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityprovidersGenericInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIdentityprovidersGenericServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIdentityprovidersGenericGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdentityprovidersGenericOK creates a GetIdentityprovidersGenericOK with default headers values
func NewGetIdentityprovidersGenericOK() *GetIdentityprovidersGenericOK {
	return &GetIdentityprovidersGenericOK{}
}

/*GetIdentityprovidersGenericOK handles this case with default header values.

successful operation
*/
type GetIdentityprovidersGenericOK struct {
	Payload *models.GenericSAML
}

func (o *GetIdentityprovidersGenericOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericOK  %+v", 200, o.Payload)
}

func (o *GetIdentityprovidersGenericOK) GetPayload() *models.GenericSAML {
	return o.Payload
}

func (o *GetIdentityprovidersGenericOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericSAML)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericBadRequest creates a GetIdentityprovidersGenericBadRequest with default headers values
func NewGetIdentityprovidersGenericBadRequest() *GetIdentityprovidersGenericBadRequest {
	return &GetIdentityprovidersGenericBadRequest{}
}

/*GetIdentityprovidersGenericBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIdentityprovidersGenericBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityprovidersGenericBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericUnauthorized creates a GetIdentityprovidersGenericUnauthorized with default headers values
func NewGetIdentityprovidersGenericUnauthorized() *GetIdentityprovidersGenericUnauthorized {
	return &GetIdentityprovidersGenericUnauthorized{}
}

/*GetIdentityprovidersGenericUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIdentityprovidersGenericUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityprovidersGenericUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericForbidden creates a GetIdentityprovidersGenericForbidden with default headers values
func NewGetIdentityprovidersGenericForbidden() *GetIdentityprovidersGenericForbidden {
	return &GetIdentityprovidersGenericForbidden{}
}

/*GetIdentityprovidersGenericForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIdentityprovidersGenericForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericForbidden  %+v", 403, o.Payload)
}

func (o *GetIdentityprovidersGenericForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericNotFound creates a GetIdentityprovidersGenericNotFound with default headers values
func NewGetIdentityprovidersGenericNotFound() *GetIdentityprovidersGenericNotFound {
	return &GetIdentityprovidersGenericNotFound{}
}

/*GetIdentityprovidersGenericNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIdentityprovidersGenericNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityprovidersGenericNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericRequestTimeout creates a GetIdentityprovidersGenericRequestTimeout with default headers values
func NewGetIdentityprovidersGenericRequestTimeout() *GetIdentityprovidersGenericRequestTimeout {
	return &GetIdentityprovidersGenericRequestTimeout{}
}

/*GetIdentityprovidersGenericRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIdentityprovidersGenericRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIdentityprovidersGenericRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericRequestEntityTooLarge creates a GetIdentityprovidersGenericRequestEntityTooLarge with default headers values
func NewGetIdentityprovidersGenericRequestEntityTooLarge() *GetIdentityprovidersGenericRequestEntityTooLarge {
	return &GetIdentityprovidersGenericRequestEntityTooLarge{}
}

/*GetIdentityprovidersGenericRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIdentityprovidersGenericRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIdentityprovidersGenericRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericUnsupportedMediaType creates a GetIdentityprovidersGenericUnsupportedMediaType with default headers values
func NewGetIdentityprovidersGenericUnsupportedMediaType() *GetIdentityprovidersGenericUnsupportedMediaType {
	return &GetIdentityprovidersGenericUnsupportedMediaType{}
}

/*GetIdentityprovidersGenericUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIdentityprovidersGenericUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIdentityprovidersGenericUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericTooManyRequests creates a GetIdentityprovidersGenericTooManyRequests with default headers values
func NewGetIdentityprovidersGenericTooManyRequests() *GetIdentityprovidersGenericTooManyRequests {
	return &GetIdentityprovidersGenericTooManyRequests{}
}

/*GetIdentityprovidersGenericTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIdentityprovidersGenericTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIdentityprovidersGenericTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericInternalServerError creates a GetIdentityprovidersGenericInternalServerError with default headers values
func NewGetIdentityprovidersGenericInternalServerError() *GetIdentityprovidersGenericInternalServerError {
	return &GetIdentityprovidersGenericInternalServerError{}
}

/*GetIdentityprovidersGenericInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIdentityprovidersGenericInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityprovidersGenericInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericServiceUnavailable creates a GetIdentityprovidersGenericServiceUnavailable with default headers values
func NewGetIdentityprovidersGenericServiceUnavailable() *GetIdentityprovidersGenericServiceUnavailable {
	return &GetIdentityprovidersGenericServiceUnavailable{}
}

/*GetIdentityprovidersGenericServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIdentityprovidersGenericServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIdentityprovidersGenericServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersGenericGatewayTimeout creates a GetIdentityprovidersGenericGatewayTimeout with default headers values
func NewGetIdentityprovidersGenericGatewayTimeout() *GetIdentityprovidersGenericGatewayTimeout {
	return &GetIdentityprovidersGenericGatewayTimeout{}
}

/*GetIdentityprovidersGenericGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIdentityprovidersGenericGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersGenericGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/generic][%d] getIdentityprovidersGenericGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIdentityprovidersGenericGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersGenericGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
