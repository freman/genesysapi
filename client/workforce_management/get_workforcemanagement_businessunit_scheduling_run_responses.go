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

// GetWorkforcemanagementBusinessunitSchedulingRunReader is a Reader for the GetWorkforcemanagementBusinessunitSchedulingRun structure.
type GetWorkforcemanagementBusinessunitSchedulingRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitSchedulingRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunOK creates a GetWorkforcemanagementBusinessunitSchedulingRunOK with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunOK() *GetWorkforcemanagementBusinessunitSchedulingRunOK {
	return &GetWorkforcemanagementBusinessunitSchedulingRunOK{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitSchedulingRunOK struct {
	Payload *models.BuScheduleRun
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run o k response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run o k response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run o k response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run o k response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run o k response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) GetPayload() *models.BuScheduleRun {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuScheduleRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunBadRequest creates a GetWorkforcemanagementBusinessunitSchedulingRunBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunBadRequest() *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest {
	return &GetWorkforcemanagementBusinessunitSchedulingRunBadRequest{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run bad request response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run bad request response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run bad request response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run bad request response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run bad request response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunUnauthorized creates a GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunUnauthorized() *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized {
	return &GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunForbidden creates a GetWorkforcemanagementBusinessunitSchedulingRunForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunForbidden() *GetWorkforcemanagementBusinessunitSchedulingRunForbidden {
	return &GetWorkforcemanagementBusinessunitSchedulingRunForbidden{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run forbidden response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run forbidden response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run forbidden response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run forbidden response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run forbidden response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunNotFound creates a GetWorkforcemanagementBusinessunitSchedulingRunNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunNotFound() *GetWorkforcemanagementBusinessunitSchedulingRunNotFound {
	return &GetWorkforcemanagementBusinessunitSchedulingRunNotFound{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run not found response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run not found response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run not found response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run not found response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run not found response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout creates a GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout() *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout {
	return &GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run request timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run request timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run request timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run request timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run request timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType() *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests creates a GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests() *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests {
	return &GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run too many requests response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run too many requests response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run too many requests response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run too many requests response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run too many requests response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunInternalServerError creates a GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunInternalServerError() *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError {
	return &GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run internal server error response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run internal server error response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run internal server error response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run internal server error response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run internal server error response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable creates a GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable() *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout creates a GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout() *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout{}
}

/*
GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunit scheduling run gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunit scheduling run gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunit scheduling run gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunit scheduling run gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunit scheduling run gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}][%d] getWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitSchedulingRunGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
