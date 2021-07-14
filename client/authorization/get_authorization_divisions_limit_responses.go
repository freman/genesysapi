// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuthorizationDivisionsLimitReader is a Reader for the GetAuthorizationDivisionsLimit structure.
type GetAuthorizationDivisionsLimitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationDivisionsLimitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationDivisionsLimitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationDivisionsLimitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationDivisionsLimitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationDivisionsLimitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationDivisionsLimitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationDivisionsLimitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationDivisionsLimitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationDivisionsLimitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationDivisionsLimitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationDivisionsLimitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationDivisionsLimitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationDivisionsLimitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationDivisionsLimitOK creates a GetAuthorizationDivisionsLimitOK with default headers values
func NewGetAuthorizationDivisionsLimitOK() *GetAuthorizationDivisionsLimitOK {
	return &GetAuthorizationDivisionsLimitOK{}
}

/*GetAuthorizationDivisionsLimitOK handles this case with default header values.

successful operation
*/
type GetAuthorizationDivisionsLimitOK struct {
	Payload int32
}

func (o *GetAuthorizationDivisionsLimitOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitOK) GetPayload() int32 {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitBadRequest creates a GetAuthorizationDivisionsLimitBadRequest with default headers values
func NewGetAuthorizationDivisionsLimitBadRequest() *GetAuthorizationDivisionsLimitBadRequest {
	return &GetAuthorizationDivisionsLimitBadRequest{}
}

/*GetAuthorizationDivisionsLimitBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationDivisionsLimitBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitUnauthorized creates a GetAuthorizationDivisionsLimitUnauthorized with default headers values
func NewGetAuthorizationDivisionsLimitUnauthorized() *GetAuthorizationDivisionsLimitUnauthorized {
	return &GetAuthorizationDivisionsLimitUnauthorized{}
}

/*GetAuthorizationDivisionsLimitUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationDivisionsLimitUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitForbidden creates a GetAuthorizationDivisionsLimitForbidden with default headers values
func NewGetAuthorizationDivisionsLimitForbidden() *GetAuthorizationDivisionsLimitForbidden {
	return &GetAuthorizationDivisionsLimitForbidden{}
}

/*GetAuthorizationDivisionsLimitForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationDivisionsLimitForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitNotFound creates a GetAuthorizationDivisionsLimitNotFound with default headers values
func NewGetAuthorizationDivisionsLimitNotFound() *GetAuthorizationDivisionsLimitNotFound {
	return &GetAuthorizationDivisionsLimitNotFound{}
}

/*GetAuthorizationDivisionsLimitNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationDivisionsLimitNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitRequestTimeout creates a GetAuthorizationDivisionsLimitRequestTimeout with default headers values
func NewGetAuthorizationDivisionsLimitRequestTimeout() *GetAuthorizationDivisionsLimitRequestTimeout {
	return &GetAuthorizationDivisionsLimitRequestTimeout{}
}

/*GetAuthorizationDivisionsLimitRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationDivisionsLimitRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitRequestEntityTooLarge creates a GetAuthorizationDivisionsLimitRequestEntityTooLarge with default headers values
func NewGetAuthorizationDivisionsLimitRequestEntityTooLarge() *GetAuthorizationDivisionsLimitRequestEntityTooLarge {
	return &GetAuthorizationDivisionsLimitRequestEntityTooLarge{}
}

/*GetAuthorizationDivisionsLimitRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAuthorizationDivisionsLimitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitUnsupportedMediaType creates a GetAuthorizationDivisionsLimitUnsupportedMediaType with default headers values
func NewGetAuthorizationDivisionsLimitUnsupportedMediaType() *GetAuthorizationDivisionsLimitUnsupportedMediaType {
	return &GetAuthorizationDivisionsLimitUnsupportedMediaType{}
}

/*GetAuthorizationDivisionsLimitUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationDivisionsLimitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitTooManyRequests creates a GetAuthorizationDivisionsLimitTooManyRequests with default headers values
func NewGetAuthorizationDivisionsLimitTooManyRequests() *GetAuthorizationDivisionsLimitTooManyRequests {
	return &GetAuthorizationDivisionsLimitTooManyRequests{}
}

/*GetAuthorizationDivisionsLimitTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationDivisionsLimitTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitInternalServerError creates a GetAuthorizationDivisionsLimitInternalServerError with default headers values
func NewGetAuthorizationDivisionsLimitInternalServerError() *GetAuthorizationDivisionsLimitInternalServerError {
	return &GetAuthorizationDivisionsLimitInternalServerError{}
}

/*GetAuthorizationDivisionsLimitInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationDivisionsLimitInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitServiceUnavailable creates a GetAuthorizationDivisionsLimitServiceUnavailable with default headers values
func NewGetAuthorizationDivisionsLimitServiceUnavailable() *GetAuthorizationDivisionsLimitServiceUnavailable {
	return &GetAuthorizationDivisionsLimitServiceUnavailable{}
}

/*GetAuthorizationDivisionsLimitServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationDivisionsLimitServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionsLimitGatewayTimeout creates a GetAuthorizationDivisionsLimitGatewayTimeout with default headers values
func NewGetAuthorizationDivisionsLimitGatewayTimeout() *GetAuthorizationDivisionsLimitGatewayTimeout {
	return &GetAuthorizationDivisionsLimitGatewayTimeout{}
}

/*GetAuthorizationDivisionsLimitGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationDivisionsLimitGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionsLimitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/limit][%d] getAuthorizationDivisionsLimitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationDivisionsLimitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionsLimitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
