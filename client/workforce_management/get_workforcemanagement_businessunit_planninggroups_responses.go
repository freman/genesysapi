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

// GetWorkforcemanagementBusinessunitPlanninggroupsReader is a Reader for the GetWorkforcemanagementBusinessunitPlanninggroups structure.
type GetWorkforcemanagementBusinessunitPlanninggroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsOK creates a GetWorkforcemanagementBusinessunitPlanninggroupsOK with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsOK() *GetWorkforcemanagementBusinessunitPlanninggroupsOK {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsOK{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsOK struct {
	Payload *models.PlanningGroupList
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups o k response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups o k response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups o k response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups o k response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups o k response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) GetPayload() *models.PlanningGroupList {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanningGroupList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsBadRequest creates a GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsBadRequest() *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups bad request response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups bad request response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups bad request response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups bad request response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups bad request response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized creates a GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized() *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsForbidden creates a GetWorkforcemanagementBusinessunitPlanninggroupsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsForbidden() *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsForbidden{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups forbidden response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups forbidden response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups forbidden response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups forbidden response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups forbidden response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsNotFound creates a GetWorkforcemanagementBusinessunitPlanninggroupsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsNotFound() *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsNotFound{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups not found response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups not found response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups not found response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups not found response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups not found response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout creates a GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout() *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups request timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups request timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups request timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups request timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups request timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests creates a GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests() *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups too many requests response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups too many requests response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups too many requests response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups too many requests response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups too many requests response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError creates a GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError() *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups internal server error response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups internal server error response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups internal server error response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups internal server error response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups internal server error response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable creates a GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable() *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout creates a GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout() *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout{}
}

/*
GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit planninggroups gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit planninggroups gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit planninggroups gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit planninggroups gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit planninggroups gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/planninggroups][%d] getWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitPlanninggroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
