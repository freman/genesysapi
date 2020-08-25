// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchWorkforcemanagementManagementunitActivitycodeReader is a Reader for the PatchWorkforcemanagementManagementunitActivitycode structure.
type PatchWorkforcemanagementManagementunitActivitycodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkforcemanagementManagementunitActivitycodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchWorkforcemanagementManagementunitActivitycodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchWorkforcemanagementManagementunitActivitycodeBadRequest creates a PatchWorkforcemanagementManagementunitActivitycodeBadRequest with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeBadRequest() *PatchWorkforcemanagementManagementunitActivitycodeBadRequest {
	return &PatchWorkforcemanagementManagementunitActivitycodeBadRequest{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchWorkforcemanagementManagementunitActivitycodeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeBadRequest  %+v", 400, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeUnauthorized creates a PatchWorkforcemanagementManagementunitActivitycodeUnauthorized with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeUnauthorized() *PatchWorkforcemanagementManagementunitActivitycodeUnauthorized {
	return &PatchWorkforcemanagementManagementunitActivitycodeUnauthorized{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchWorkforcemanagementManagementunitActivitycodeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeForbidden creates a PatchWorkforcemanagementManagementunitActivitycodeForbidden with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeForbidden() *PatchWorkforcemanagementManagementunitActivitycodeForbidden {
	return &PatchWorkforcemanagementManagementunitActivitycodeForbidden{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchWorkforcemanagementManagementunitActivitycodeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeForbidden  %+v", 403, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeNotFound creates a PatchWorkforcemanagementManagementunitActivitycodeNotFound with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeNotFound() *PatchWorkforcemanagementManagementunitActivitycodeNotFound {
	return &PatchWorkforcemanagementManagementunitActivitycodeNotFound{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchWorkforcemanagementManagementunitActivitycodeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeNotFound  %+v", 404, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeGone creates a PatchWorkforcemanagementManagementunitActivitycodeGone with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeGone() *PatchWorkforcemanagementManagementunitActivitycodeGone {
	return &PatchWorkforcemanagementManagementunitActivitycodeGone{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeGone handles this case with default header values.

Gone
*/
type PatchWorkforcemanagementManagementunitActivitycodeGone struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeGone) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeGone  %+v", 410, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge creates a PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge() *PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge {
	return &PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType creates a PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType() *PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType {
	return &PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeTooManyRequests creates a PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeTooManyRequests() *PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests {
	return &PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeInternalServerError creates a PatchWorkforcemanagementManagementunitActivitycodeInternalServerError with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeInternalServerError() *PatchWorkforcemanagementManagementunitActivitycodeInternalServerError {
	return &PatchWorkforcemanagementManagementunitActivitycodeInternalServerError{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchWorkforcemanagementManagementunitActivitycodeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable creates a PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable() *PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable {
	return &PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout creates a PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout() *PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout {
	return &PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout{}
}

/*PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitActivitycodeDefault creates a PatchWorkforcemanagementManagementunitActivitycodeDefault with default headers values
func NewPatchWorkforcemanagementManagementunitActivitycodeDefault(code int) *PatchWorkforcemanagementManagementunitActivitycodeDefault {
	return &PatchWorkforcemanagementManagementunitActivitycodeDefault{
		_statusCode: code,
	}
}

/*PatchWorkforcemanagementManagementunitActivitycodeDefault handles this case with default header values.

successful operation
*/
type PatchWorkforcemanagementManagementunitActivitycodeDefault struct {
	_statusCode int
}

// Code gets the status code for the patch workforcemanagement managementunit activitycode default response
func (o *PatchWorkforcemanagementManagementunitActivitycodeDefault) Code() int {
	return o._statusCode
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{muId}/activitycodes/{acId}][%d] patchWorkforcemanagementManagementunitActivitycode default ", o._statusCode)
}

func (o *PatchWorkforcemanagementManagementunitActivitycodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}