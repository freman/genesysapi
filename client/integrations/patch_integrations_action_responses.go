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

// PatchIntegrationsActionReader is a Reader for the PatchIntegrationsAction structure.
type PatchIntegrationsActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchIntegrationsActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchIntegrationsActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchIntegrationsActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchIntegrationsActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchIntegrationsActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchIntegrationsActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchIntegrationsActionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchIntegrationsActionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchIntegrationsActionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchIntegrationsActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchIntegrationsActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchIntegrationsActionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchIntegrationsActionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchIntegrationsActionOK creates a PatchIntegrationsActionOK with default headers values
func NewPatchIntegrationsActionOK() *PatchIntegrationsActionOK {
	return &PatchIntegrationsActionOK{}
}

/*PatchIntegrationsActionOK handles this case with default header values.

successful operation
*/
type PatchIntegrationsActionOK struct {
	Payload *models.Action
}

func (o *PatchIntegrationsActionOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionOK  %+v", 200, o.Payload)
}

func (o *PatchIntegrationsActionOK) GetPayload() *models.Action {
	return o.Payload
}

func (o *PatchIntegrationsActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Action)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionBadRequest creates a PatchIntegrationsActionBadRequest with default headers values
func NewPatchIntegrationsActionBadRequest() *PatchIntegrationsActionBadRequest {
	return &PatchIntegrationsActionBadRequest{}
}

/*PatchIntegrationsActionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchIntegrationsActionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionBadRequest  %+v", 400, o.Payload)
}

func (o *PatchIntegrationsActionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionUnauthorized creates a PatchIntegrationsActionUnauthorized with default headers values
func NewPatchIntegrationsActionUnauthorized() *PatchIntegrationsActionUnauthorized {
	return &PatchIntegrationsActionUnauthorized{}
}

/*PatchIntegrationsActionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchIntegrationsActionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchIntegrationsActionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionForbidden creates a PatchIntegrationsActionForbidden with default headers values
func NewPatchIntegrationsActionForbidden() *PatchIntegrationsActionForbidden {
	return &PatchIntegrationsActionForbidden{}
}

/*PatchIntegrationsActionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchIntegrationsActionForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionForbidden  %+v", 403, o.Payload)
}

func (o *PatchIntegrationsActionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionNotFound creates a PatchIntegrationsActionNotFound with default headers values
func NewPatchIntegrationsActionNotFound() *PatchIntegrationsActionNotFound {
	return &PatchIntegrationsActionNotFound{}
}

/*PatchIntegrationsActionNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchIntegrationsActionNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionNotFound  %+v", 404, o.Payload)
}

func (o *PatchIntegrationsActionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionRequestTimeout creates a PatchIntegrationsActionRequestTimeout with default headers values
func NewPatchIntegrationsActionRequestTimeout() *PatchIntegrationsActionRequestTimeout {
	return &PatchIntegrationsActionRequestTimeout{}
}

/*PatchIntegrationsActionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchIntegrationsActionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchIntegrationsActionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionRequestEntityTooLarge creates a PatchIntegrationsActionRequestEntityTooLarge with default headers values
func NewPatchIntegrationsActionRequestEntityTooLarge() *PatchIntegrationsActionRequestEntityTooLarge {
	return &PatchIntegrationsActionRequestEntityTooLarge{}
}

/*PatchIntegrationsActionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchIntegrationsActionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchIntegrationsActionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionUnsupportedMediaType creates a PatchIntegrationsActionUnsupportedMediaType with default headers values
func NewPatchIntegrationsActionUnsupportedMediaType() *PatchIntegrationsActionUnsupportedMediaType {
	return &PatchIntegrationsActionUnsupportedMediaType{}
}

/*PatchIntegrationsActionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchIntegrationsActionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchIntegrationsActionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionTooManyRequests creates a PatchIntegrationsActionTooManyRequests with default headers values
func NewPatchIntegrationsActionTooManyRequests() *PatchIntegrationsActionTooManyRequests {
	return &PatchIntegrationsActionTooManyRequests{}
}

/*PatchIntegrationsActionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchIntegrationsActionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchIntegrationsActionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionInternalServerError creates a PatchIntegrationsActionInternalServerError with default headers values
func NewPatchIntegrationsActionInternalServerError() *PatchIntegrationsActionInternalServerError {
	return &PatchIntegrationsActionInternalServerError{}
}

/*PatchIntegrationsActionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchIntegrationsActionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchIntegrationsActionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionServiceUnavailable creates a PatchIntegrationsActionServiceUnavailable with default headers values
func NewPatchIntegrationsActionServiceUnavailable() *PatchIntegrationsActionServiceUnavailable {
	return &PatchIntegrationsActionServiceUnavailable{}
}

/*PatchIntegrationsActionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchIntegrationsActionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchIntegrationsActionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchIntegrationsActionGatewayTimeout creates a PatchIntegrationsActionGatewayTimeout with default headers values
func NewPatchIntegrationsActionGatewayTimeout() *PatchIntegrationsActionGatewayTimeout {
	return &PatchIntegrationsActionGatewayTimeout{}
}

/*PatchIntegrationsActionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchIntegrationsActionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchIntegrationsActionGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/integrations/actions/{actionId}][%d] patchIntegrationsActionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchIntegrationsActionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchIntegrationsActionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
