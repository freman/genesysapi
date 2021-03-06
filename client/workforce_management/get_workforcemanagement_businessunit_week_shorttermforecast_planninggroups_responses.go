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

// GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsReader is a Reader for the GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroups structure.
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK struct {
	Payload *models.ForecastPlanningGroupsResponse
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK) GetPayload() *models.ForecastPlanningGroupsResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ForecastPlanningGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout creates a GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout() *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout{}
}

/*GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekDateId}/shorttermforecasts/{forecastId}/planninggroups][%d] getWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitWeekShorttermforecastPlanninggroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
