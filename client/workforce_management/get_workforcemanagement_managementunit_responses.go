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

// GetWorkforcemanagementManagementunitReader is a Reader for the GetWorkforcemanagementManagementunit structure.
type GetWorkforcemanagementManagementunitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementManagementunitRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitOK creates a GetWorkforcemanagementManagementunitOK with default headers values
func NewGetWorkforcemanagementManagementunitOK() *GetWorkforcemanagementManagementunitOK {
	return &GetWorkforcemanagementManagementunitOK{}
}

/*
GetWorkforcemanagementManagementunitOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitOK struct {
	Payload *models.ManagementUnit
}

// IsSuccess returns true when this get workforcemanagement managementunit o k response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement managementunit o k response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit o k response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit o k response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit o k response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementManagementunitOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitOK) GetPayload() *models.ManagementUnit {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ManagementUnit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitBadRequest creates a GetWorkforcemanagementManagementunitBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitBadRequest() *GetWorkforcemanagementManagementunitBadRequest {
	return &GetWorkforcemanagementManagementunitBadRequest{}
}

/*
GetWorkforcemanagementManagementunitBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit bad request response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit bad request response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit bad request response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit bad request response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit bad request response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementManagementunitBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUnauthorized creates a GetWorkforcemanagementManagementunitUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitUnauthorized() *GetWorkforcemanagementManagementunitUnauthorized {
	return &GetWorkforcemanagementManagementunitUnauthorized{}
}

/*
GetWorkforcemanagementManagementunitUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitForbidden creates a GetWorkforcemanagementManagementunitForbidden with default headers values
func NewGetWorkforcemanagementManagementunitForbidden() *GetWorkforcemanagementManagementunitForbidden {
	return &GetWorkforcemanagementManagementunitForbidden{}
}

/*
GetWorkforcemanagementManagementunitForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit forbidden response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit forbidden response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit forbidden response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit forbidden response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit forbidden response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementManagementunitForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitNotFound creates a GetWorkforcemanagementManagementunitNotFound with default headers values
func NewGetWorkforcemanagementManagementunitNotFound() *GetWorkforcemanagementManagementunitNotFound {
	return &GetWorkforcemanagementManagementunitNotFound{}
}

/*
GetWorkforcemanagementManagementunitNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit not found response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit not found response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit not found response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit not found response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit not found response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementManagementunitNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitRequestTimeout creates a GetWorkforcemanagementManagementunitRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitRequestTimeout() *GetWorkforcemanagementManagementunitRequestTimeout {
	return &GetWorkforcemanagementManagementunitRequestTimeout{}
}

/*
GetWorkforcemanagementManagementunitRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit request timeout response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit request timeout response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit request timeout response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit request timeout response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit request timeout response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementManagementunitRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitRequestEntityTooLarge() *GetWorkforcemanagementManagementunitRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementManagementunitRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUnsupportedMediaType creates a GetWorkforcemanagementManagementunitUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitUnsupportedMediaType() *GetWorkforcemanagementManagementunitUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitUnsupportedMediaType{}
}

/*
GetWorkforcemanagementManagementunitUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitTooManyRequests creates a GetWorkforcemanagementManagementunitTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitTooManyRequests() *GetWorkforcemanagementManagementunitTooManyRequests {
	return &GetWorkforcemanagementManagementunitTooManyRequests{}
}

/*
GetWorkforcemanagementManagementunitTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit too many requests response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit too many requests response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit too many requests response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit too many requests response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit too many requests response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitInternalServerError creates a GetWorkforcemanagementManagementunitInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitInternalServerError() *GetWorkforcemanagementManagementunitInternalServerError {
	return &GetWorkforcemanagementManagementunitInternalServerError{}
}

/*
GetWorkforcemanagementManagementunitInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit internal server error response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit internal server error response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit internal server error response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit internal server error response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit internal server error response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitServiceUnavailable creates a GetWorkforcemanagementManagementunitServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitServiceUnavailable() *GetWorkforcemanagementManagementunitServiceUnavailable {
	return &GetWorkforcemanagementManagementunitServiceUnavailable{}
}

/*
GetWorkforcemanagementManagementunitServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitGatewayTimeout creates a GetWorkforcemanagementManagementunitGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitGatewayTimeout() *GetWorkforcemanagementManagementunitGatewayTimeout {
	return &GetWorkforcemanagementManagementunitGatewayTimeout{}
}

/*
GetWorkforcemanagementManagementunitGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}][%d] getWorkforcemanagementManagementunitGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
