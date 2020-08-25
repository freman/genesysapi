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

// GetWorkforcemanagementManagementunitShifttradesMatchedReader is a Reader for the GetWorkforcemanagementManagementunitShifttradesMatched structure.
type GetWorkforcemanagementManagementunitShifttradesMatchedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitShifttradesMatchedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedOK creates a GetWorkforcemanagementManagementunitShifttradesMatchedOK with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedOK() *GetWorkforcemanagementManagementunitShifttradesMatchedOK {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedOK{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedOK struct {
	Payload *models.ShiftTradeMatchesSummaryResponse
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedOK) GetPayload() *models.ShiftTradeMatchesSummaryResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ShiftTradeMatchesSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedBadRequest creates a GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedBadRequest() *GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized creates a GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized() *GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedForbidden creates a GetWorkforcemanagementManagementunitShifttradesMatchedForbidden with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedForbidden() *GetWorkforcemanagementManagementunitShifttradesMatchedForbidden {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedForbidden{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedNotFound creates a GetWorkforcemanagementManagementunitShifttradesMatchedNotFound with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedNotFound() *GetWorkforcemanagementManagementunitShifttradesMatchedNotFound {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedNotFound{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge() *GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType creates a GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType() *GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests creates a GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests() *GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError creates a GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError() *GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable creates a GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable() *GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout creates a GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout() *GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout {
	return &GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/matched][%d] getWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesMatchedGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}