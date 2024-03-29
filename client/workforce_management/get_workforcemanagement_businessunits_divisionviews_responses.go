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

// GetWorkforcemanagementBusinessunitsDivisionviewsReader is a Reader for the GetWorkforcemanagementBusinessunitsDivisionviews structure.
type GetWorkforcemanagementBusinessunitsDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsOK creates a GetWorkforcemanagementBusinessunitsDivisionviewsOK with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsOK() *GetWorkforcemanagementBusinessunitsDivisionviewsOK {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsOK{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsOK struct {
	Payload *models.BusinessUnitListing
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews o k response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews o k response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews o k response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews o k response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews o k response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) GetPayload() *models.BusinessUnitListing {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BusinessUnitListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsBadRequest creates a GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsBadRequest() *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews bad request response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews bad request response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews bad request response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews bad request response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews bad request response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized creates a GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized() *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsForbidden creates a GetWorkforcemanagementBusinessunitsDivisionviewsForbidden with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsForbidden() *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsForbidden{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews forbidden response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews forbidden response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews forbidden response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews forbidden response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews forbidden response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsNotFound creates a GetWorkforcemanagementBusinessunitsDivisionviewsNotFound with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsNotFound() *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsNotFound{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews not found response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews not found response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews not found response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews not found response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews not found response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout creates a GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout() *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews request timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews request timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews request timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews request timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews request timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge creates a GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge() *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType creates a GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType() *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests creates a GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests() *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews too many requests response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews too many requests response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews too many requests response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews too many requests response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews too many requests response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError creates a GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError() *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews internal server error response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews internal server error response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews internal server error response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews internal server error response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews internal server error response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable creates a GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable() *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout creates a GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout with default headers values
func NewGetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout() *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout {
	return &GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout{}
}

/*
GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement businessunits divisionviews gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement businessunits divisionviews gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement businessunits divisionviews gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement businessunits divisionviews gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement businessunits divisionviews gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/businessunits/divisionviews][%d] getWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementBusinessunitsDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
