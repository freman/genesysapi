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

// GetIdentityprovidersPingReader is a Reader for the GetIdentityprovidersPing structure.
type GetIdentityprovidersPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentityprovidersPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIdentityprovidersPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIdentityprovidersPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIdentityprovidersPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIdentityprovidersPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIdentityprovidersPingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIdentityprovidersPingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIdentityprovidersPingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIdentityprovidersPingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIdentityprovidersPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIdentityprovidersPingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIdentityprovidersPingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIdentityprovidersPingOK creates a GetIdentityprovidersPingOK with default headers values
func NewGetIdentityprovidersPingOK() *GetIdentityprovidersPingOK {
	return &GetIdentityprovidersPingOK{}
}

/*GetIdentityprovidersPingOK handles this case with default header values.

successful operation
*/
type GetIdentityprovidersPingOK struct {
	Payload *models.PingIdentity
}

func (o *GetIdentityprovidersPingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingOK  %+v", 200, o.Payload)
}

func (o *GetIdentityprovidersPingOK) GetPayload() *models.PingIdentity {
	return o.Payload
}

func (o *GetIdentityprovidersPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PingIdentity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingBadRequest creates a GetIdentityprovidersPingBadRequest with default headers values
func NewGetIdentityprovidersPingBadRequest() *GetIdentityprovidersPingBadRequest {
	return &GetIdentityprovidersPingBadRequest{}
}

/*GetIdentityprovidersPingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIdentityprovidersPingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingBadRequest  %+v", 400, o.Payload)
}

func (o *GetIdentityprovidersPingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingUnauthorized creates a GetIdentityprovidersPingUnauthorized with default headers values
func NewGetIdentityprovidersPingUnauthorized() *GetIdentityprovidersPingUnauthorized {
	return &GetIdentityprovidersPingUnauthorized{}
}

/*GetIdentityprovidersPingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIdentityprovidersPingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIdentityprovidersPingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingForbidden creates a GetIdentityprovidersPingForbidden with default headers values
func NewGetIdentityprovidersPingForbidden() *GetIdentityprovidersPingForbidden {
	return &GetIdentityprovidersPingForbidden{}
}

/*GetIdentityprovidersPingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIdentityprovidersPingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingForbidden  %+v", 403, o.Payload)
}

func (o *GetIdentityprovidersPingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingNotFound creates a GetIdentityprovidersPingNotFound with default headers values
func NewGetIdentityprovidersPingNotFound() *GetIdentityprovidersPingNotFound {
	return &GetIdentityprovidersPingNotFound{}
}

/*GetIdentityprovidersPingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIdentityprovidersPingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingNotFound  %+v", 404, o.Payload)
}

func (o *GetIdentityprovidersPingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingRequestEntityTooLarge creates a GetIdentityprovidersPingRequestEntityTooLarge with default headers values
func NewGetIdentityprovidersPingRequestEntityTooLarge() *GetIdentityprovidersPingRequestEntityTooLarge {
	return &GetIdentityprovidersPingRequestEntityTooLarge{}
}

/*GetIdentityprovidersPingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIdentityprovidersPingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIdentityprovidersPingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingUnsupportedMediaType creates a GetIdentityprovidersPingUnsupportedMediaType with default headers values
func NewGetIdentityprovidersPingUnsupportedMediaType() *GetIdentityprovidersPingUnsupportedMediaType {
	return &GetIdentityprovidersPingUnsupportedMediaType{}
}

/*GetIdentityprovidersPingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIdentityprovidersPingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIdentityprovidersPingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingTooManyRequests creates a GetIdentityprovidersPingTooManyRequests with default headers values
func NewGetIdentityprovidersPingTooManyRequests() *GetIdentityprovidersPingTooManyRequests {
	return &GetIdentityprovidersPingTooManyRequests{}
}

/*GetIdentityprovidersPingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIdentityprovidersPingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIdentityprovidersPingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingInternalServerError creates a GetIdentityprovidersPingInternalServerError with default headers values
func NewGetIdentityprovidersPingInternalServerError() *GetIdentityprovidersPingInternalServerError {
	return &GetIdentityprovidersPingInternalServerError{}
}

/*GetIdentityprovidersPingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIdentityprovidersPingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIdentityprovidersPingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingServiceUnavailable creates a GetIdentityprovidersPingServiceUnavailable with default headers values
func NewGetIdentityprovidersPingServiceUnavailable() *GetIdentityprovidersPingServiceUnavailable {
	return &GetIdentityprovidersPingServiceUnavailable{}
}

/*GetIdentityprovidersPingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIdentityprovidersPingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIdentityprovidersPingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentityprovidersPingGatewayTimeout creates a GetIdentityprovidersPingGatewayTimeout with default headers values
func NewGetIdentityprovidersPingGatewayTimeout() *GetIdentityprovidersPingGatewayTimeout {
	return &GetIdentityprovidersPingGatewayTimeout{}
}

/*GetIdentityprovidersPingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIdentityprovidersPingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIdentityprovidersPingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/identityproviders/ping][%d] getIdentityprovidersPingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIdentityprovidersPingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIdentityprovidersPingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}