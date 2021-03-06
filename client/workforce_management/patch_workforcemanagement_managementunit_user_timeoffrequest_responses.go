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

// PatchWorkforcemanagementManagementunitUserTimeoffrequestReader is a Reader for the PatchWorkforcemanagementManagementunitUserTimeoffrequest structure.
type PatchWorkforcemanagementManagementunitUserTimeoffrequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestOK creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestOK with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestOK() *PatchWorkforcemanagementManagementunitUserTimeoffrequestOK {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestOK{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestOK handles this case with default header values.

successful operation
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestOK struct {
	Payload *models.TimeOffRequestResponse
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestOK  %+v", 200, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestOK) GetPayload() *models.TimeOffRequestResponse {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest() *PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest  %+v", 400, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized() *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden() *PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestForbidden  %+v", 403, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound() *PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestNotFound  %+v", 404, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestConflict creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestConflict() *PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict handles this case with default header values.

Conflict
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestConflict  %+v", 409, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge() *PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType() *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests() *PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError() *PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable() *PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout creates a PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout with default headers values
func NewPatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout() *PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout {
	return &PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout{}
}

/*PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}][%d] patchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchWorkforcemanagementManagementunitUserTimeoffrequestGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
