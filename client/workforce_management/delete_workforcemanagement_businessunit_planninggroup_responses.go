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

// DeleteWorkforcemanagementBusinessunitPlanninggroupReader is a Reader for the DeleteWorkforcemanagementBusinessunitPlanninggroup structure.
type DeleteWorkforcemanagementBusinessunitPlanninggroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupNoContent creates a DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupNoContent() *DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent handles this case with default header values.

The planning group was deleted successfully
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent struct {
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupNoContent ", 204)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest creates a DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest() *DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized creates a DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized() *DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupForbidden creates a DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupForbidden() *DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupNotFound creates a DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupNotFound() *DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge creates a DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge() *DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType creates a DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType() *DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests creates a DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests() *DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError creates a DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError() *DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable creates a DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable() *DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout creates a DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout() *DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout {
	return &DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout{}
}

/*DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] deleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}