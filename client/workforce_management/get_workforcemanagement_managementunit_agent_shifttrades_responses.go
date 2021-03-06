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

// GetWorkforcemanagementManagementunitAgentShifttradesReader is a Reader for the GetWorkforcemanagementManagementunitAgentShifttrades structure.
type GetWorkforcemanagementManagementunitAgentShifttradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitAgentShifttradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesOK creates a GetWorkforcemanagementManagementunitAgentShifttradesOK with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesOK() *GetWorkforcemanagementManagementunitAgentShifttradesOK {
	return &GetWorkforcemanagementManagementunitAgentShifttradesOK{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitAgentShifttradesOK struct {
	Payload *models.ShiftTradeListResponse
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesOK) GetPayload() *models.ShiftTradeListResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShiftTradeListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesBadRequest creates a GetWorkforcemanagementManagementunitAgentShifttradesBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesBadRequest() *GetWorkforcemanagementManagementunitAgentShifttradesBadRequest {
	return &GetWorkforcemanagementManagementunitAgentShifttradesBadRequest{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesUnauthorized creates a GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesUnauthorized() *GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized {
	return &GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesForbidden creates a GetWorkforcemanagementManagementunitAgentShifttradesForbidden with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesForbidden() *GetWorkforcemanagementManagementunitAgentShifttradesForbidden {
	return &GetWorkforcemanagementManagementunitAgentShifttradesForbidden{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesNotFound creates a GetWorkforcemanagementManagementunitAgentShifttradesNotFound with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesNotFound() *GetWorkforcemanagementManagementunitAgentShifttradesNotFound {
	return &GetWorkforcemanagementManagementunitAgentShifttradesNotFound{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge() *GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType creates a GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType() *GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests creates a GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests() *GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests {
	return &GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesInternalServerError creates a GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesInternalServerError() *GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError {
	return &GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable creates a GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable() *GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable {
	return &GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout creates a GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout() *GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout {
	return &GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents/{agentId}/shifttrades][%d] getWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitAgentShifttradesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
