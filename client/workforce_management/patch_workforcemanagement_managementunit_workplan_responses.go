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

// PatchWorkforcemanagementManagementunitWorkplanReader is a Reader for the PatchWorkforcemanagementManagementunitWorkplan structure.
type PatchWorkforcemanagementManagementunitWorkplanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkforcemanagementManagementunitWorkplanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkforcemanagementManagementunitWorkplanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchWorkforcemanagementManagementunitWorkplanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkforcemanagementManagementunitWorkplanUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkforcemanagementManagementunitWorkplanForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkforcemanagementManagementunitWorkplanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchWorkforcemanagementManagementunitWorkplanConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchWorkforcemanagementManagementunitWorkplanTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchWorkforcemanagementManagementunitWorkplanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchWorkforcemanagementManagementunitWorkplanServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchWorkforcemanagementManagementunitWorkplanGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchWorkforcemanagementManagementunitWorkplanOK creates a PatchWorkforcemanagementManagementunitWorkplanOK with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanOK() *PatchWorkforcemanagementManagementunitWorkplanOK {
	return &PatchWorkforcemanagementManagementunitWorkplanOK{}
}

/*PatchWorkforcemanagementManagementunitWorkplanOK handles this case with default header values.

successful operation
*/
type PatchWorkforcemanagementManagementunitWorkplanOK struct {
	Payload *models.WorkPlan
}

func (o *PatchWorkforcemanagementManagementunitWorkplanOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanOK  %+v", 200, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanOK) GetPayload() *models.WorkPlan {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanBadRequest creates a PatchWorkforcemanagementManagementunitWorkplanBadRequest with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanBadRequest() *PatchWorkforcemanagementManagementunitWorkplanBadRequest {
	return &PatchWorkforcemanagementManagementunitWorkplanBadRequest{}
}

/*PatchWorkforcemanagementManagementunitWorkplanBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchWorkforcemanagementManagementunitWorkplanBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanBadRequest  %+v", 400, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanUnauthorized creates a PatchWorkforcemanagementManagementunitWorkplanUnauthorized with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanUnauthorized() *PatchWorkforcemanagementManagementunitWorkplanUnauthorized {
	return &PatchWorkforcemanagementManagementunitWorkplanUnauthorized{}
}

/*PatchWorkforcemanagementManagementunitWorkplanUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchWorkforcemanagementManagementunitWorkplanUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanForbidden creates a PatchWorkforcemanagementManagementunitWorkplanForbidden with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanForbidden() *PatchWorkforcemanagementManagementunitWorkplanForbidden {
	return &PatchWorkforcemanagementManagementunitWorkplanForbidden{}
}

/*PatchWorkforcemanagementManagementunitWorkplanForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchWorkforcemanagementManagementunitWorkplanForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanForbidden  %+v", 403, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanNotFound creates a PatchWorkforcemanagementManagementunitWorkplanNotFound with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanNotFound() *PatchWorkforcemanagementManagementunitWorkplanNotFound {
	return &PatchWorkforcemanagementManagementunitWorkplanNotFound{}
}

/*PatchWorkforcemanagementManagementunitWorkplanNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchWorkforcemanagementManagementunitWorkplanNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanNotFound  %+v", 404, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanConflict creates a PatchWorkforcemanagementManagementunitWorkplanConflict with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanConflict() *PatchWorkforcemanagementManagementunitWorkplanConflict {
	return &PatchWorkforcemanagementManagementunitWorkplanConflict{}
}

/*PatchWorkforcemanagementManagementunitWorkplanConflict handles this case with default header values.

Conflict
*/
type PatchWorkforcemanagementManagementunitWorkplanConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanConflict  %+v", 409, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge creates a PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge() *PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge {
	return &PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge{}
}

/*PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType creates a PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType() *PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType {
	return &PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType{}
}

/*PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanTooManyRequests creates a PatchWorkforcemanagementManagementunitWorkplanTooManyRequests with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanTooManyRequests() *PatchWorkforcemanagementManagementunitWorkplanTooManyRequests {
	return &PatchWorkforcemanagementManagementunitWorkplanTooManyRequests{}
}

/*PatchWorkforcemanagementManagementunitWorkplanTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchWorkforcemanagementManagementunitWorkplanTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanInternalServerError creates a PatchWorkforcemanagementManagementunitWorkplanInternalServerError with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanInternalServerError() *PatchWorkforcemanagementManagementunitWorkplanInternalServerError {
	return &PatchWorkforcemanagementManagementunitWorkplanInternalServerError{}
}

/*PatchWorkforcemanagementManagementunitWorkplanInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchWorkforcemanagementManagementunitWorkplanInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanServiceUnavailable creates a PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanServiceUnavailable() *PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable {
	return &PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable{}
}

/*PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanGatewayTimeout creates a PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanGatewayTimeout() *PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout {
	return &PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout{}
}

/*PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplans/{workPlanId}][%d] patchWorkforcemanagementManagementunitWorkplanGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
