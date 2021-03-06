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

// PatchWorkforcemanagementBusinessunitActivitycodeReader is a Reader for the PatchWorkforcemanagementBusinessunitActivitycode structure.
type PatchWorkforcemanagementBusinessunitActivitycodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkforcemanagementBusinessunitActivitycodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeOK creates a PatchWorkforcemanagementBusinessunitActivitycodeOK with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeOK() *PatchWorkforcemanagementBusinessunitActivitycodeOK {
	return &PatchWorkforcemanagementBusinessunitActivitycodeOK{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeOK handles this case with default header values.

successful operation
*/
type PatchWorkforcemanagementBusinessunitActivitycodeOK struct {
	Payload *models.BusinessUnitActivityCode
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeOK  %+v", 200, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeOK) GetPayload() *models.BusinessUnitActivityCode {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessUnitActivityCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeBadRequest creates a PatchWorkforcemanagementBusinessunitActivitycodeBadRequest with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeBadRequest() *PatchWorkforcemanagementBusinessunitActivitycodeBadRequest {
	return &PatchWorkforcemanagementBusinessunitActivitycodeBadRequest{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeBadRequest  %+v", 400, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeUnauthorized creates a PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeUnauthorized() *PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized {
	return &PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeForbidden creates a PatchWorkforcemanagementBusinessunitActivitycodeForbidden with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeForbidden() *PatchWorkforcemanagementBusinessunitActivitycodeForbidden {
	return &PatchWorkforcemanagementBusinessunitActivitycodeForbidden{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeForbidden  %+v", 403, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeNotFound creates a PatchWorkforcemanagementBusinessunitActivitycodeNotFound with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeNotFound() *PatchWorkforcemanagementBusinessunitActivitycodeNotFound {
	return &PatchWorkforcemanagementBusinessunitActivitycodeNotFound{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeNotFound  %+v", 404, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeConflict creates a PatchWorkforcemanagementBusinessunitActivitycodeConflict with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeConflict() *PatchWorkforcemanagementBusinessunitActivitycodeConflict {
	return &PatchWorkforcemanagementBusinessunitActivitycodeConflict{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeConflict handles this case with default header values.

Conflict
*/
type PatchWorkforcemanagementBusinessunitActivitycodeConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeConflict  %+v", 409, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge creates a PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge() *PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge {
	return &PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType creates a PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType() *PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType {
	return &PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests creates a PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests() *PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests {
	return &PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeInternalServerError creates a PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeInternalServerError() *PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError {
	return &PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable creates a PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable() *PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable {
	return &PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout creates a PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout with default headers values
func NewPatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout() *PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout {
	return &PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout{}
}

/*PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/businessunits/{businessUnitId}/activitycodes/{activityCodeId}][%d] patchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementBusinessunitActivitycodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
