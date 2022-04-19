// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchIntegrationsActionDraftReader is a Reader for the PatchIntegrationsActionDraft structure.
type PatchIntegrationsActionDraftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIntegrationsActionDraftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchIntegrationsActionDraftOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchIntegrationsActionDraftBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchIntegrationsActionDraftUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchIntegrationsActionDraftForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchIntegrationsActionDraftNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchIntegrationsActionDraftRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchIntegrationsActionDraftRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchIntegrationsActionDraftUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchIntegrationsActionDraftTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchIntegrationsActionDraftInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchIntegrationsActionDraftServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchIntegrationsActionDraftGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchIntegrationsActionDraftOK creates a PatchIntegrationsActionDraftOK with default headers values
func NewPatchIntegrationsActionDraftOK() *PatchIntegrationsActionDraftOK {
	return &PatchIntegrationsActionDraftOK{}
}

/*PatchIntegrationsActionDraftOK handles this case with default header values.

successful operation
*/
type PatchIntegrationsActionDraftOK struct {
	Payload *models.Action
}

func (o *PatchIntegrationsActionDraftOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftOK  %+v", 200, o.Payload)
}

func (o *PatchIntegrationsActionDraftOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftBadRequest creates a PatchIntegrationsActionDraftBadRequest with default headers values
func NewPatchIntegrationsActionDraftBadRequest() *PatchIntegrationsActionDraftBadRequest {
	return &PatchIntegrationsActionDraftBadRequest{}
}

/*PatchIntegrationsActionDraftBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchIntegrationsActionDraftBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftBadRequest  %+v", 400, o.Payload)
}

func (o *PatchIntegrationsActionDraftBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftUnauthorized creates a PatchIntegrationsActionDraftUnauthorized with default headers values
func NewPatchIntegrationsActionDraftUnauthorized() *PatchIntegrationsActionDraftUnauthorized {
	return &PatchIntegrationsActionDraftUnauthorized{}
}

/*PatchIntegrationsActionDraftUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchIntegrationsActionDraftUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchIntegrationsActionDraftUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftForbidden creates a PatchIntegrationsActionDraftForbidden with default headers values
func NewPatchIntegrationsActionDraftForbidden() *PatchIntegrationsActionDraftForbidden {
	return &PatchIntegrationsActionDraftForbidden{}
}

/*PatchIntegrationsActionDraftForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchIntegrationsActionDraftForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftForbidden  %+v", 403, o.Payload)
}

func (o *PatchIntegrationsActionDraftForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftNotFound creates a PatchIntegrationsActionDraftNotFound with default headers values
func NewPatchIntegrationsActionDraftNotFound() *PatchIntegrationsActionDraftNotFound {
	return &PatchIntegrationsActionDraftNotFound{}
}

/*PatchIntegrationsActionDraftNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchIntegrationsActionDraftNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftNotFound  %+v", 404, o.Payload)
}

func (o *PatchIntegrationsActionDraftNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftRequestTimeout creates a PatchIntegrationsActionDraftRequestTimeout with default headers values
func NewPatchIntegrationsActionDraftRequestTimeout() *PatchIntegrationsActionDraftRequestTimeout {
	return &PatchIntegrationsActionDraftRequestTimeout{}
}

/*PatchIntegrationsActionDraftRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchIntegrationsActionDraftRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchIntegrationsActionDraftRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftRequestEntityTooLarge creates a PatchIntegrationsActionDraftRequestEntityTooLarge with default headers values
func NewPatchIntegrationsActionDraftRequestEntityTooLarge() *PatchIntegrationsActionDraftRequestEntityTooLarge {
	return &PatchIntegrationsActionDraftRequestEntityTooLarge{}
}

/*PatchIntegrationsActionDraftRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchIntegrationsActionDraftRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftUnsupportedMediaType creates a PatchIntegrationsActionDraftUnsupportedMediaType with default headers values
func NewPatchIntegrationsActionDraftUnsupportedMediaType() *PatchIntegrationsActionDraftUnsupportedMediaType {
	return &PatchIntegrationsActionDraftUnsupportedMediaType{}
}

/*PatchIntegrationsActionDraftUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchIntegrationsActionDraftUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftTooManyRequests creates a PatchIntegrationsActionDraftTooManyRequests with default headers values
func NewPatchIntegrationsActionDraftTooManyRequests() *PatchIntegrationsActionDraftTooManyRequests {
	return &PatchIntegrationsActionDraftTooManyRequests{}
}

/*PatchIntegrationsActionDraftTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchIntegrationsActionDraftTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchIntegrationsActionDraftTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftInternalServerError creates a PatchIntegrationsActionDraftInternalServerError with default headers values
func NewPatchIntegrationsActionDraftInternalServerError() *PatchIntegrationsActionDraftInternalServerError {
	return &PatchIntegrationsActionDraftInternalServerError{}
}

/*PatchIntegrationsActionDraftInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchIntegrationsActionDraftInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchIntegrationsActionDraftInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftServiceUnavailable creates a PatchIntegrationsActionDraftServiceUnavailable with default headers values
func NewPatchIntegrationsActionDraftServiceUnavailable() *PatchIntegrationsActionDraftServiceUnavailable {
	return &PatchIntegrationsActionDraftServiceUnavailable{}
}

/*PatchIntegrationsActionDraftServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchIntegrationsActionDraftServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionDraftGatewayTimeout creates a PatchIntegrationsActionDraftGatewayTimeout with default headers values
func NewPatchIntegrationsActionDraftGatewayTimeout() *PatchIntegrationsActionDraftGatewayTimeout {
	return &PatchIntegrationsActionDraftGatewayTimeout{}
}

/*PatchIntegrationsActionDraftGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchIntegrationsActionDraftGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}/draft][%d] patchIntegrationsActionDraftGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionDraftGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
