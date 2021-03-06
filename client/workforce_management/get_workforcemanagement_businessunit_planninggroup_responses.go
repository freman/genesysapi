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

// GetWorkforcemanagementBusinessunitPlanninggroupReader is a Reader for the GetWorkforcemanagementBusinessunitPlanninggroup structure.
type GetWorkforcemanagementBusinessunitPlanninggroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitPlanninggroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupOK creates a GetWorkforcemanagementBusinessunitPlanninggroupOK with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupOK() *GetWorkforcemanagementBusinessunitPlanninggroupOK {
	return &GetWorkforcemanagementBusinessunitPlanninggroupOK{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitPlanninggroupOK struct {
	Payload *models.PlanningGroup
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupOK) GetPayload() *models.PlanningGroup {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanningGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupBadRequest creates a GetWorkforcemanagementBusinessunitPlanninggroupBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupBadRequest() *GetWorkforcemanagementBusinessunitPlanninggroupBadRequest {
	return &GetWorkforcemanagementBusinessunitPlanninggroupBadRequest{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupUnauthorized creates a GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupUnauthorized() *GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized {
	return &GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupForbidden creates a GetWorkforcemanagementBusinessunitPlanninggroupForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupForbidden() *GetWorkforcemanagementBusinessunitPlanninggroupForbidden {
	return &GetWorkforcemanagementBusinessunitPlanninggroupForbidden{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupNotFound creates a GetWorkforcemanagementBusinessunitPlanninggroupNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupNotFound() *GetWorkforcemanagementBusinessunitPlanninggroupNotFound {
	return &GetWorkforcemanagementBusinessunitPlanninggroupNotFound{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType() *GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests creates a GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests() *GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests {
	return &GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupInternalServerError creates a GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupInternalServerError() *GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError {
	return &GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable creates a GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable() *GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout creates a GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout() *GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups/{planningGroupId}][%d] getWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
