// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingEmailSetupReader is a Reader for the GetRoutingEmailSetup structure.
type GetRoutingEmailSetupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingEmailSetupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingEmailSetupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingEmailSetupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingEmailSetupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingEmailSetupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingEmailSetupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingEmailSetupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingEmailSetupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingEmailSetupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingEmailSetupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingEmailSetupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingEmailSetupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingEmailSetupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingEmailSetupOK creates a GetRoutingEmailSetupOK with default headers values
func NewGetRoutingEmailSetupOK() *GetRoutingEmailSetupOK {
	return &GetRoutingEmailSetupOK{}
}

/*GetRoutingEmailSetupOK handles this case with default header values.

successful operation
*/
type GetRoutingEmailSetupOK struct {
	Payload *models.EmailSetup
}

func (o *GetRoutingEmailSetupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupOK  %+v", 200, o.Payload)
}

func (o *GetRoutingEmailSetupOK) GetPayload() *models.EmailSetup {
	return o.Payload
}

func (o *GetRoutingEmailSetupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailSetup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupBadRequest creates a GetRoutingEmailSetupBadRequest with default headers values
func NewGetRoutingEmailSetupBadRequest() *GetRoutingEmailSetupBadRequest {
	return &GetRoutingEmailSetupBadRequest{}
}

/*GetRoutingEmailSetupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingEmailSetupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingEmailSetupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupUnauthorized creates a GetRoutingEmailSetupUnauthorized with default headers values
func NewGetRoutingEmailSetupUnauthorized() *GetRoutingEmailSetupUnauthorized {
	return &GetRoutingEmailSetupUnauthorized{}
}

/*GetRoutingEmailSetupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingEmailSetupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingEmailSetupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupForbidden creates a GetRoutingEmailSetupForbidden with default headers values
func NewGetRoutingEmailSetupForbidden() *GetRoutingEmailSetupForbidden {
	return &GetRoutingEmailSetupForbidden{}
}

/*GetRoutingEmailSetupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingEmailSetupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingEmailSetupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupNotFound creates a GetRoutingEmailSetupNotFound with default headers values
func NewGetRoutingEmailSetupNotFound() *GetRoutingEmailSetupNotFound {
	return &GetRoutingEmailSetupNotFound{}
}

/*GetRoutingEmailSetupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingEmailSetupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingEmailSetupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupRequestTimeout creates a GetRoutingEmailSetupRequestTimeout with default headers values
func NewGetRoutingEmailSetupRequestTimeout() *GetRoutingEmailSetupRequestTimeout {
	return &GetRoutingEmailSetupRequestTimeout{}
}

/*GetRoutingEmailSetupRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingEmailSetupRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingEmailSetupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupRequestEntityTooLarge creates a GetRoutingEmailSetupRequestEntityTooLarge with default headers values
func NewGetRoutingEmailSetupRequestEntityTooLarge() *GetRoutingEmailSetupRequestEntityTooLarge {
	return &GetRoutingEmailSetupRequestEntityTooLarge{}
}

/*GetRoutingEmailSetupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingEmailSetupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingEmailSetupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupUnsupportedMediaType creates a GetRoutingEmailSetupUnsupportedMediaType with default headers values
func NewGetRoutingEmailSetupUnsupportedMediaType() *GetRoutingEmailSetupUnsupportedMediaType {
	return &GetRoutingEmailSetupUnsupportedMediaType{}
}

/*GetRoutingEmailSetupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingEmailSetupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingEmailSetupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupTooManyRequests creates a GetRoutingEmailSetupTooManyRequests with default headers values
func NewGetRoutingEmailSetupTooManyRequests() *GetRoutingEmailSetupTooManyRequests {
	return &GetRoutingEmailSetupTooManyRequests{}
}

/*GetRoutingEmailSetupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingEmailSetupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingEmailSetupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupInternalServerError creates a GetRoutingEmailSetupInternalServerError with default headers values
func NewGetRoutingEmailSetupInternalServerError() *GetRoutingEmailSetupInternalServerError {
	return &GetRoutingEmailSetupInternalServerError{}
}

/*GetRoutingEmailSetupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingEmailSetupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingEmailSetupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupServiceUnavailable creates a GetRoutingEmailSetupServiceUnavailable with default headers values
func NewGetRoutingEmailSetupServiceUnavailable() *GetRoutingEmailSetupServiceUnavailable {
	return &GetRoutingEmailSetupServiceUnavailable{}
}

/*GetRoutingEmailSetupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingEmailSetupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingEmailSetupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingEmailSetupGatewayTimeout creates a GetRoutingEmailSetupGatewayTimeout with default headers values
func NewGetRoutingEmailSetupGatewayTimeout() *GetRoutingEmailSetupGatewayTimeout {
	return &GetRoutingEmailSetupGatewayTimeout{}
}

/*GetRoutingEmailSetupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingEmailSetupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingEmailSetupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/email/setup][%d] getRoutingEmailSetupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingEmailSetupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingEmailSetupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
