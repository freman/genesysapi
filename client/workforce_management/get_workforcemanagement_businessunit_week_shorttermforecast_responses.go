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

// GetWorkforcemanagementBusinessunitWeekShorttermforecastReader is a Reader for the GetWorkforcemanagementBusinessunitWeekShorttermforecast structure.
type GetWorkforcemanagementBusinessunitWeekShorttermforecastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastOK creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastOK with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastOK() *GetWorkforcemanagementBusinessunitWeekShorttermforecastOK {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastOK{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastOK struct {
	Payload *models.BuShortTermForecast
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastOK) GetPayload() *models.BuShortTermForecast {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuShortTermForecast)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest() *GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized() *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden() *GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound() *GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType() *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests() *GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError() *GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable() *GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout() *GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
