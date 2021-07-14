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

// PatchWorkforcemanagementManagementunitWorkplanrotationReader is a Reader for the PatchWorkforcemanagementManagementunitWorkplanrotation structure.
type PatchWorkforcemanagementManagementunitWorkplanrotationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkforcemanagementManagementunitWorkplanrotationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationOK creates a PatchWorkforcemanagementManagementunitWorkplanrotationOK with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationOK() *PatchWorkforcemanagementManagementunitWorkplanrotationOK {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationOK{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationOK handles this case with default header values.

successful operation
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationOK struct {
	Payload *models.WorkPlanRotationResponse
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationOK  %+v", 200, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationOK) GetPayload() *models.WorkPlanRotationResponse {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkPlanRotationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationBadRequest creates a PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationBadRequest() *PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized creates a PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized() *PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationForbidden creates a PatchWorkforcemanagementManagementunitWorkplanrotationForbidden with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationForbidden() *PatchWorkforcemanagementManagementunitWorkplanrotationForbidden {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationForbidden{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationForbidden  %+v", 403, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationNotFound creates a PatchWorkforcemanagementManagementunitWorkplanrotationNotFound with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationNotFound() *PatchWorkforcemanagementManagementunitWorkplanrotationNotFound {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationNotFound{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationNotFound  %+v", 404, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout creates a PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout() *PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationConflict creates a PatchWorkforcemanagementManagementunitWorkplanrotationConflict with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationConflict() *PatchWorkforcemanagementManagementunitWorkplanrotationConflict {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationConflict{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationConflict handles this case with default header values.

Conflict
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationConflict  %+v", 409, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge creates a PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge() *PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType creates a PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType() *PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests creates a PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests() *PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError creates a PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError() *PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable creates a PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable() *PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout creates a PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout() *PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout {
	return &PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout{}
}

/*PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/workplanrotations/{workPlanRotationId}][%d] patchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitWorkplanrotationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
