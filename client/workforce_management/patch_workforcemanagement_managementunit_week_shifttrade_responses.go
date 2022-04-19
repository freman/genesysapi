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

// PatchWorkforcemanagementManagementunitWeekShifttradeReader is a Reader for the PatchWorkforcemanagementManagementunitWeekShifttrade structure.
type PatchWorkforcemanagementManagementunitWeekShifttradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkforcemanagementManagementunitWeekShifttradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeOK creates a PatchWorkforcemanagementManagementunitWeekShifttradeOK with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeOK() *PatchWorkforcemanagementManagementunitWeekShifttradeOK {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeOK{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeOK handles this case with default header values.

successful operation
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeOK struct {
	Payload *models.ShiftTradeResponse
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeOK  %+v", 200, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeOK) GetPayload() *models.ShiftTradeResponse {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShiftTradeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeBadRequest creates a PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeBadRequest() *PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeBadRequest  %+v", 400, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized creates a PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized() *PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeForbidden creates a PatchWorkforcemanagementManagementunitWeekShifttradeForbidden with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeForbidden() *PatchWorkforcemanagementManagementunitWeekShifttradeForbidden {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeForbidden{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeForbidden  %+v", 403, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeNotFound creates a PatchWorkforcemanagementManagementunitWeekShifttradeNotFound with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeNotFound() *PatchWorkforcemanagementManagementunitWeekShifttradeNotFound {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeNotFound{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeNotFound  %+v", 404, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout creates a PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout() *PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeConflict creates a PatchWorkforcemanagementManagementunitWeekShifttradeConflict with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeConflict() *PatchWorkforcemanagementManagementunitWeekShifttradeConflict {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeConflict{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeConflict handles this case with default header values.

Conflict
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeConflict  %+v", 409, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge creates a PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge() *PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType creates a PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType() *PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests creates a PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests() *PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError creates a PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError() *PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable creates a PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable() *PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout creates a PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout() *PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout {
	return &PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout{}
}

/*PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekDateId}/shifttrades/{tradeId}][%d] patchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWeekShifttradeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
