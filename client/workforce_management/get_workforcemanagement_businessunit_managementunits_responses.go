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

// GetWorkforcemanagementBusinessunitManagementunitsReader is a Reader for the GetWorkforcemanagementBusinessunitManagementunits structure.
type GetWorkforcemanagementBusinessunitManagementunitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitManagementunitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitManagementunitsOK creates a GetWorkforcemanagementBusinessunitManagementunitsOK with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsOK() *GetWorkforcemanagementBusinessunitManagementunitsOK {
	return &GetWorkforcemanagementBusinessunitManagementunitsOK{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitManagementunitsOK struct {
	Payload *models.ManagementUnitListing
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits o k response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits o k response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits o k response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits o k response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits o k response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) GetPayload() *models.ManagementUnitListing {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnitListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsBadRequest creates a GetWorkforcemanagementBusinessunitManagementunitsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsBadRequest() *GetWorkforcemanagementBusinessunitManagementunitsBadRequest {
	return &GetWorkforcemanagementBusinessunitManagementunitsBadRequest{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitManagementunitsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits bad request response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits bad request response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits bad request response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits bad request response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits bad request response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsUnauthorized creates a GetWorkforcemanagementBusinessunitManagementunitsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsUnauthorized() *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized {
	return &GetWorkforcemanagementBusinessunitManagementunitsUnauthorized{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitManagementunitsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsForbidden creates a GetWorkforcemanagementBusinessunitManagementunitsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsForbidden() *GetWorkforcemanagementBusinessunitManagementunitsForbidden {
	return &GetWorkforcemanagementBusinessunitManagementunitsForbidden{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitManagementunitsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits forbidden response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits forbidden response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits forbidden response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits forbidden response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits forbidden response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsNotFound creates a GetWorkforcemanagementBusinessunitManagementunitsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsNotFound() *GetWorkforcemanagementBusinessunitManagementunitsNotFound {
	return &GetWorkforcemanagementBusinessunitManagementunitsNotFound{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitManagementunitsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits not found response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits not found response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits not found response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits not found response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits not found response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsRequestTimeout creates a GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsRequestTimeout() *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout {
	return &GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits request timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits request timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits request timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits request timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits request timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsTooManyRequests creates a GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsTooManyRequests() *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits too many requests response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits too many requests response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits too many requests response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits too many requests response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit managementunits too many requests response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsInternalServerError creates a GetWorkforcemanagementBusinessunitManagementunitsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsInternalServerError() *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError {
	return &GetWorkforcemanagementBusinessunitManagementunitsInternalServerError{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitManagementunitsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits internal server error response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits internal server error response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits internal server error response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits internal server error response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit managementunits internal server error response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable creates a GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable() *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit managementunits service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout creates a GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout() *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout{}
}

/*
GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit managementunits gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit managementunits gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit managementunits gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit managementunits gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit managementunits gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/managementunits][%d] getWorkforcemanagementBusinessunitManagementunitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitManagementunitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
