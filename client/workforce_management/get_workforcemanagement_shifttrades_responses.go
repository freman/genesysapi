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

// GetWorkforcemanagementShifttradesReader is a Reader for the GetWorkforcemanagementShifttrades structure.
type GetWorkforcemanagementShifttradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementShifttradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementShifttradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementShifttradesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementShifttradesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementShifttradesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementShifttradesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementShifttradesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementShifttradesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementShifttradesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementShifttradesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementShifttradesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementShifttradesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementShifttradesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementShifttradesOK creates a GetWorkforcemanagementShifttradesOK with default headers values
func NewGetWorkforcemanagementShifttradesOK() *GetWorkforcemanagementShifttradesOK {
	return &GetWorkforcemanagementShifttradesOK{}
}

/*GetWorkforcemanagementShifttradesOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementShifttradesOK struct {
	Payload *models.ShiftTradeListResponse
}

func (o *GetWorkforcemanagementShifttradesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesOK) GetPayload() *models.ShiftTradeListResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShiftTradeListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesBadRequest creates a GetWorkforcemanagementShifttradesBadRequest with default headers values
func NewGetWorkforcemanagementShifttradesBadRequest() *GetWorkforcemanagementShifttradesBadRequest {
	return &GetWorkforcemanagementShifttradesBadRequest{}
}

/*GetWorkforcemanagementShifttradesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementShifttradesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesUnauthorized creates a GetWorkforcemanagementShifttradesUnauthorized with default headers values
func NewGetWorkforcemanagementShifttradesUnauthorized() *GetWorkforcemanagementShifttradesUnauthorized {
	return &GetWorkforcemanagementShifttradesUnauthorized{}
}

/*GetWorkforcemanagementShifttradesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementShifttradesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesForbidden creates a GetWorkforcemanagementShifttradesForbidden with default headers values
func NewGetWorkforcemanagementShifttradesForbidden() *GetWorkforcemanagementShifttradesForbidden {
	return &GetWorkforcemanagementShifttradesForbidden{}
}

/*GetWorkforcemanagementShifttradesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementShifttradesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesNotFound creates a GetWorkforcemanagementShifttradesNotFound with default headers values
func NewGetWorkforcemanagementShifttradesNotFound() *GetWorkforcemanagementShifttradesNotFound {
	return &GetWorkforcemanagementShifttradesNotFound{}
}

/*GetWorkforcemanagementShifttradesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementShifttradesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesRequestTimeout creates a GetWorkforcemanagementShifttradesRequestTimeout with default headers values
func NewGetWorkforcemanagementShifttradesRequestTimeout() *GetWorkforcemanagementShifttradesRequestTimeout {
	return &GetWorkforcemanagementShifttradesRequestTimeout{}
}

/*GetWorkforcemanagementShifttradesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementShifttradesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesRequestEntityTooLarge creates a GetWorkforcemanagementShifttradesRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementShifttradesRequestEntityTooLarge() *GetWorkforcemanagementShifttradesRequestEntityTooLarge {
	return &GetWorkforcemanagementShifttradesRequestEntityTooLarge{}
}

/*GetWorkforcemanagementShifttradesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementShifttradesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesUnsupportedMediaType creates a GetWorkforcemanagementShifttradesUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementShifttradesUnsupportedMediaType() *GetWorkforcemanagementShifttradesUnsupportedMediaType {
	return &GetWorkforcemanagementShifttradesUnsupportedMediaType{}
}

/*GetWorkforcemanagementShifttradesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementShifttradesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesTooManyRequests creates a GetWorkforcemanagementShifttradesTooManyRequests with default headers values
func NewGetWorkforcemanagementShifttradesTooManyRequests() *GetWorkforcemanagementShifttradesTooManyRequests {
	return &GetWorkforcemanagementShifttradesTooManyRequests{}
}

/*GetWorkforcemanagementShifttradesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementShifttradesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesInternalServerError creates a GetWorkforcemanagementShifttradesInternalServerError with default headers values
func NewGetWorkforcemanagementShifttradesInternalServerError() *GetWorkforcemanagementShifttradesInternalServerError {
	return &GetWorkforcemanagementShifttradesInternalServerError{}
}

/*GetWorkforcemanagementShifttradesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementShifttradesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesServiceUnavailable creates a GetWorkforcemanagementShifttradesServiceUnavailable with default headers values
func NewGetWorkforcemanagementShifttradesServiceUnavailable() *GetWorkforcemanagementShifttradesServiceUnavailable {
	return &GetWorkforcemanagementShifttradesServiceUnavailable{}
}

/*GetWorkforcemanagementShifttradesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementShifttradesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementShifttradesGatewayTimeout creates a GetWorkforcemanagementShifttradesGatewayTimeout with default headers values
func NewGetWorkforcemanagementShifttradesGatewayTimeout() *GetWorkforcemanagementShifttradesGatewayTimeout {
	return &GetWorkforcemanagementShifttradesGatewayTimeout{}
}

/*GetWorkforcemanagementShifttradesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementShifttradesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementShifttradesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/shifttrades][%d] getWorkforcemanagementShifttradesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementShifttradesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementShifttradesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
