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

// PatchRoutingSettingsContactcenterReader is a Reader for the PatchRoutingSettingsContactcenter structure.
type PatchRoutingSettingsContactcenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingSettingsContactcenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPatchRoutingSettingsContactcenterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingSettingsContactcenterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingSettingsContactcenterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingSettingsContactcenterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingSettingsContactcenterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingSettingsContactcenterRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingSettingsContactcenterUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingSettingsContactcenterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingSettingsContactcenterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingSettingsContactcenterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingSettingsContactcenterGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingSettingsContactcenterAccepted creates a PatchRoutingSettingsContactcenterAccepted with default headers values
func NewPatchRoutingSettingsContactcenterAccepted() *PatchRoutingSettingsContactcenterAccepted {
	return &PatchRoutingSettingsContactcenterAccepted{}
}

/*PatchRoutingSettingsContactcenterAccepted handles this case with default header values.

Accepted
*/
type PatchRoutingSettingsContactcenterAccepted struct {
}

func (o *PatchRoutingSettingsContactcenterAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterAccepted ", 202)
}

func (o *PatchRoutingSettingsContactcenterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRoutingSettingsContactcenterBadRequest creates a PatchRoutingSettingsContactcenterBadRequest with default headers values
func NewPatchRoutingSettingsContactcenterBadRequest() *PatchRoutingSettingsContactcenterBadRequest {
	return &PatchRoutingSettingsContactcenterBadRequest{}
}

/*PatchRoutingSettingsContactcenterBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingSettingsContactcenterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterUnauthorized creates a PatchRoutingSettingsContactcenterUnauthorized with default headers values
func NewPatchRoutingSettingsContactcenterUnauthorized() *PatchRoutingSettingsContactcenterUnauthorized {
	return &PatchRoutingSettingsContactcenterUnauthorized{}
}

/*PatchRoutingSettingsContactcenterUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingSettingsContactcenterUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterForbidden creates a PatchRoutingSettingsContactcenterForbidden with default headers values
func NewPatchRoutingSettingsContactcenterForbidden() *PatchRoutingSettingsContactcenterForbidden {
	return &PatchRoutingSettingsContactcenterForbidden{}
}

/*PatchRoutingSettingsContactcenterForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingSettingsContactcenterForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterNotFound creates a PatchRoutingSettingsContactcenterNotFound with default headers values
func NewPatchRoutingSettingsContactcenterNotFound() *PatchRoutingSettingsContactcenterNotFound {
	return &PatchRoutingSettingsContactcenterNotFound{}
}

/*PatchRoutingSettingsContactcenterNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRoutingSettingsContactcenterNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterRequestEntityTooLarge creates a PatchRoutingSettingsContactcenterRequestEntityTooLarge with default headers values
func NewPatchRoutingSettingsContactcenterRequestEntityTooLarge() *PatchRoutingSettingsContactcenterRequestEntityTooLarge {
	return &PatchRoutingSettingsContactcenterRequestEntityTooLarge{}
}

/*PatchRoutingSettingsContactcenterRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchRoutingSettingsContactcenterRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterUnsupportedMediaType creates a PatchRoutingSettingsContactcenterUnsupportedMediaType with default headers values
func NewPatchRoutingSettingsContactcenterUnsupportedMediaType() *PatchRoutingSettingsContactcenterUnsupportedMediaType {
	return &PatchRoutingSettingsContactcenterUnsupportedMediaType{}
}

/*PatchRoutingSettingsContactcenterUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingSettingsContactcenterUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterTooManyRequests creates a PatchRoutingSettingsContactcenterTooManyRequests with default headers values
func NewPatchRoutingSettingsContactcenterTooManyRequests() *PatchRoutingSettingsContactcenterTooManyRequests {
	return &PatchRoutingSettingsContactcenterTooManyRequests{}
}

/*PatchRoutingSettingsContactcenterTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchRoutingSettingsContactcenterTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterInternalServerError creates a PatchRoutingSettingsContactcenterInternalServerError with default headers values
func NewPatchRoutingSettingsContactcenterInternalServerError() *PatchRoutingSettingsContactcenterInternalServerError {
	return &PatchRoutingSettingsContactcenterInternalServerError{}
}

/*PatchRoutingSettingsContactcenterInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingSettingsContactcenterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterServiceUnavailable creates a PatchRoutingSettingsContactcenterServiceUnavailable with default headers values
func NewPatchRoutingSettingsContactcenterServiceUnavailable() *PatchRoutingSettingsContactcenterServiceUnavailable {
	return &PatchRoutingSettingsContactcenterServiceUnavailable{}
}

/*PatchRoutingSettingsContactcenterServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingSettingsContactcenterServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingSettingsContactcenterGatewayTimeout creates a PatchRoutingSettingsContactcenterGatewayTimeout with default headers values
func NewPatchRoutingSettingsContactcenterGatewayTimeout() *PatchRoutingSettingsContactcenterGatewayTimeout {
	return &PatchRoutingSettingsContactcenterGatewayTimeout{}
}

/*PatchRoutingSettingsContactcenterGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRoutingSettingsContactcenterGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingSettingsContactcenterGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/settings/contactcenter][%d] patchRoutingSettingsContactcenterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingSettingsContactcenterGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingSettingsContactcenterGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
