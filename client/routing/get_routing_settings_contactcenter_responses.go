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

// GetRoutingSettingsContactcenterReader is a Reader for the GetRoutingSettingsContactcenter structure.
type GetRoutingSettingsContactcenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSettingsContactcenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSettingsContactcenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSettingsContactcenterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSettingsContactcenterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSettingsContactcenterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSettingsContactcenterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSettingsContactcenterRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSettingsContactcenterRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSettingsContactcenterUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSettingsContactcenterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSettingsContactcenterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSettingsContactcenterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSettingsContactcenterGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSettingsContactcenterOK creates a GetRoutingSettingsContactcenterOK with default headers values
func NewGetRoutingSettingsContactcenterOK() *GetRoutingSettingsContactcenterOK {
	return &GetRoutingSettingsContactcenterOK{}
}

/*GetRoutingSettingsContactcenterOK handles this case with default header values.

successful operation
*/
type GetRoutingSettingsContactcenterOK struct {
	Payload *models.ContactCenterSettings
}

func (o *GetRoutingSettingsContactcenterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSettingsContactcenterOK) GetPayload() *models.ContactCenterSettings {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactCenterSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterBadRequest creates a GetRoutingSettingsContactcenterBadRequest with default headers values
func NewGetRoutingSettingsContactcenterBadRequest() *GetRoutingSettingsContactcenterBadRequest {
	return &GetRoutingSettingsContactcenterBadRequest{}
}

/*GetRoutingSettingsContactcenterBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSettingsContactcenterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSettingsContactcenterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterUnauthorized creates a GetRoutingSettingsContactcenterUnauthorized with default headers values
func NewGetRoutingSettingsContactcenterUnauthorized() *GetRoutingSettingsContactcenterUnauthorized {
	return &GetRoutingSettingsContactcenterUnauthorized{}
}

/*GetRoutingSettingsContactcenterUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSettingsContactcenterUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSettingsContactcenterUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterForbidden creates a GetRoutingSettingsContactcenterForbidden with default headers values
func NewGetRoutingSettingsContactcenterForbidden() *GetRoutingSettingsContactcenterForbidden {
	return &GetRoutingSettingsContactcenterForbidden{}
}

/*GetRoutingSettingsContactcenterForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSettingsContactcenterForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSettingsContactcenterForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterNotFound creates a GetRoutingSettingsContactcenterNotFound with default headers values
func NewGetRoutingSettingsContactcenterNotFound() *GetRoutingSettingsContactcenterNotFound {
	return &GetRoutingSettingsContactcenterNotFound{}
}

/*GetRoutingSettingsContactcenterNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingSettingsContactcenterNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSettingsContactcenterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterRequestTimeout creates a GetRoutingSettingsContactcenterRequestTimeout with default headers values
func NewGetRoutingSettingsContactcenterRequestTimeout() *GetRoutingSettingsContactcenterRequestTimeout {
	return &GetRoutingSettingsContactcenterRequestTimeout{}
}

/*GetRoutingSettingsContactcenterRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSettingsContactcenterRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterRequestEntityTooLarge creates a GetRoutingSettingsContactcenterRequestEntityTooLarge with default headers values
func NewGetRoutingSettingsContactcenterRequestEntityTooLarge() *GetRoutingSettingsContactcenterRequestEntityTooLarge {
	return &GetRoutingSettingsContactcenterRequestEntityTooLarge{}
}

/*GetRoutingSettingsContactcenterRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingSettingsContactcenterRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterUnsupportedMediaType creates a GetRoutingSettingsContactcenterUnsupportedMediaType with default headers values
func NewGetRoutingSettingsContactcenterUnsupportedMediaType() *GetRoutingSettingsContactcenterUnsupportedMediaType {
	return &GetRoutingSettingsContactcenterUnsupportedMediaType{}
}

/*GetRoutingSettingsContactcenterUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSettingsContactcenterUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterTooManyRequests creates a GetRoutingSettingsContactcenterTooManyRequests with default headers values
func NewGetRoutingSettingsContactcenterTooManyRequests() *GetRoutingSettingsContactcenterTooManyRequests {
	return &GetRoutingSettingsContactcenterTooManyRequests{}
}

/*GetRoutingSettingsContactcenterTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSettingsContactcenterTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterInternalServerError creates a GetRoutingSettingsContactcenterInternalServerError with default headers values
func NewGetRoutingSettingsContactcenterInternalServerError() *GetRoutingSettingsContactcenterInternalServerError {
	return &GetRoutingSettingsContactcenterInternalServerError{}
}

/*GetRoutingSettingsContactcenterInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSettingsContactcenterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSettingsContactcenterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterServiceUnavailable creates a GetRoutingSettingsContactcenterServiceUnavailable with default headers values
func NewGetRoutingSettingsContactcenterServiceUnavailable() *GetRoutingSettingsContactcenterServiceUnavailable {
	return &GetRoutingSettingsContactcenterServiceUnavailable{}
}

/*GetRoutingSettingsContactcenterServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSettingsContactcenterServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterGatewayTimeout creates a GetRoutingSettingsContactcenterGatewayTimeout with default headers values
func NewGetRoutingSettingsContactcenterGatewayTimeout() *GetRoutingSettingsContactcenterGatewayTimeout {
	return &GetRoutingSettingsContactcenterGatewayTimeout{}
}

/*GetRoutingSettingsContactcenterGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingSettingsContactcenterGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
