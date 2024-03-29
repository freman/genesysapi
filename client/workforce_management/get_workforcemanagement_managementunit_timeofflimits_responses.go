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

// GetWorkforcemanagementManagementunitTimeofflimitsReader is a Reader for the GetWorkforcemanagementManagementunitTimeofflimits structure.
type GetWorkforcemanagementManagementunitTimeofflimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitTimeofflimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsOK creates a GetWorkforcemanagementManagementunitTimeofflimitsOK with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsOK() *GetWorkforcemanagementManagementunitTimeofflimitsOK {
	return &GetWorkforcemanagementManagementunitTimeofflimitsOK{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitTimeofflimitsOK struct {
	Payload *models.TimeOffLimitListing
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits o k response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits o k response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits o k response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits o k response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits o k response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) GetPayload() *models.TimeOffLimitListing {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeOffLimitListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsBadRequest creates a GetWorkforcemanagementManagementunitTimeofflimitsBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsBadRequest() *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest {
	return &GetWorkforcemanagementManagementunitTimeofflimitsBadRequest{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits bad request response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits bad request response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits bad request response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits bad request response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits bad request response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsUnauthorized creates a GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsUnauthorized() *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized {
	return &GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsForbidden creates a GetWorkforcemanagementManagementunitTimeofflimitsForbidden with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsForbidden() *GetWorkforcemanagementManagementunitTimeofflimitsForbidden {
	return &GetWorkforcemanagementManagementunitTimeofflimitsForbidden{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits forbidden response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits forbidden response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits forbidden response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits forbidden response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits forbidden response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsNotFound creates a GetWorkforcemanagementManagementunitTimeofflimitsNotFound with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsNotFound() *GetWorkforcemanagementManagementunitTimeofflimitsNotFound {
	return &GetWorkforcemanagementManagementunitTimeofflimitsNotFound{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits not found response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits not found response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits not found response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits not found response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits not found response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout creates a GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout() *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout {
	return &GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits request timeout response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits request timeout response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits request timeout response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits request timeout response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits request timeout response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge() *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType creates a GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType() *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests creates a GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests() *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests {
	return &GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits too many requests response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits too many requests response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits too many requests response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits too many requests response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits too many requests response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsInternalServerError creates a GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsInternalServerError() *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError {
	return &GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits internal server error response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits internal server error response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits internal server error response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits internal server error response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits internal server error response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable creates a GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable() *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable {
	return &GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout creates a GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout() *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout {
	return &GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout{}
}

/*
GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit timeofflimits gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit timeofflimits gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit timeofflimits gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit timeofflimits gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit timeofflimits gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/timeofflimits][%d] getWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTimeofflimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
