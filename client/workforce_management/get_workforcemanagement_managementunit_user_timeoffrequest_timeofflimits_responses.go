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

// GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsReader is a Reader for the GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimits structure.
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK struct {
	Payload *models.QueryTimeOffLimitValuesResponse
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK) GetPayload() *models.QueryTimeOffLimitValuesResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryTimeOffLimitValuesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout creates a GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout() *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout {
	return &GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users/{userId}/timeoffrequests/{timeOffRequestId}/timeofflimits][%d] getWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUserTimeoffrequestTimeofflimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
