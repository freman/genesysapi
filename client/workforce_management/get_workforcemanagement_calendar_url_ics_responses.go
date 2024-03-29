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

// GetWorkforcemanagementCalendarURLIcsReader is a Reader for the GetWorkforcemanagementCalendarURLIcs structure.
type GetWorkforcemanagementCalendarURLIcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementCalendarURLIcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementCalendarURLIcsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementCalendarURLIcsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementCalendarURLIcsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementCalendarURLIcsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementCalendarURLIcsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetWorkforcemanagementCalendarURLIcsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementCalendarURLIcsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementCalendarURLIcsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementCalendarURLIcsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementCalendarURLIcsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementCalendarURLIcsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementCalendarURLIcsOK creates a GetWorkforcemanagementCalendarURLIcsOK with default headers values
func NewGetWorkforcemanagementCalendarURLIcsOK() *GetWorkforcemanagementCalendarURLIcsOK {
	return &GetWorkforcemanagementCalendarURLIcsOK{}
}

/*
GetWorkforcemanagementCalendarURLIcsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementCalendarURLIcsOK struct {
	Payload *models.CalendarURLResponse
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics o k response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics o k response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics o k response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement calendar Url ics o k response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics o k response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementCalendarURLIcsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsOK) GetPayload() *models.CalendarURLResponse {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CalendarURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsBadRequest creates a GetWorkforcemanagementCalendarURLIcsBadRequest with default headers values
func NewGetWorkforcemanagementCalendarURLIcsBadRequest() *GetWorkforcemanagementCalendarURLIcsBadRequest {
	return &GetWorkforcemanagementCalendarURLIcsBadRequest{}
}

/*
GetWorkforcemanagementCalendarURLIcsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementCalendarURLIcsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics bad request response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics bad request response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics bad request response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics bad request response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics bad request response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsUnauthorized creates a GetWorkforcemanagementCalendarURLIcsUnauthorized with default headers values
func NewGetWorkforcemanagementCalendarURLIcsUnauthorized() *GetWorkforcemanagementCalendarURLIcsUnauthorized {
	return &GetWorkforcemanagementCalendarURLIcsUnauthorized{}
}

/*
GetWorkforcemanagementCalendarURLIcsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementCalendarURLIcsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsForbidden creates a GetWorkforcemanagementCalendarURLIcsForbidden with default headers values
func NewGetWorkforcemanagementCalendarURLIcsForbidden() *GetWorkforcemanagementCalendarURLIcsForbidden {
	return &GetWorkforcemanagementCalendarURLIcsForbidden{}
}

/*
GetWorkforcemanagementCalendarURLIcsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementCalendarURLIcsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics forbidden response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics forbidden response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics forbidden response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics forbidden response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics forbidden response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementCalendarURLIcsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsNotFound creates a GetWorkforcemanagementCalendarURLIcsNotFound with default headers values
func NewGetWorkforcemanagementCalendarURLIcsNotFound() *GetWorkforcemanagementCalendarURLIcsNotFound {
	return &GetWorkforcemanagementCalendarURLIcsNotFound{}
}

/*
GetWorkforcemanagementCalendarURLIcsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementCalendarURLIcsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics not found response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics not found response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics not found response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics not found response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics not found response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementCalendarURLIcsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsRequestTimeout creates a GetWorkforcemanagementCalendarURLIcsRequestTimeout with default headers values
func NewGetWorkforcemanagementCalendarURLIcsRequestTimeout() *GetWorkforcemanagementCalendarURLIcsRequestTimeout {
	return &GetWorkforcemanagementCalendarURLIcsRequestTimeout{}
}

/*
GetWorkforcemanagementCalendarURLIcsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementCalendarURLIcsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics request timeout response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics request timeout response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics request timeout response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics request timeout response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics request timeout response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge creates a GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge() *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge {
	return &GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge{}
}

/*
GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsUnsupportedMediaType creates a GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementCalendarURLIcsUnsupportedMediaType() *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType {
	return &GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType{}
}

/*
GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsTooManyRequests creates a GetWorkforcemanagementCalendarURLIcsTooManyRequests with default headers values
func NewGetWorkforcemanagementCalendarURLIcsTooManyRequests() *GetWorkforcemanagementCalendarURLIcsTooManyRequests {
	return &GetWorkforcemanagementCalendarURLIcsTooManyRequests{}
}

/*
GetWorkforcemanagementCalendarURLIcsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementCalendarURLIcsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics too many requests response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics too many requests response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics too many requests response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement calendar Url ics too many requests response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement calendar Url ics too many requests response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsInternalServerError creates a GetWorkforcemanagementCalendarURLIcsInternalServerError with default headers values
func NewGetWorkforcemanagementCalendarURLIcsInternalServerError() *GetWorkforcemanagementCalendarURLIcsInternalServerError {
	return &GetWorkforcemanagementCalendarURLIcsInternalServerError{}
}

/*
GetWorkforcemanagementCalendarURLIcsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementCalendarURLIcsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics internal server error response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics internal server error response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics internal server error response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement calendar Url ics internal server error response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement calendar Url ics internal server error response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsServiceUnavailable creates a GetWorkforcemanagementCalendarURLIcsServiceUnavailable with default headers values
func NewGetWorkforcemanagementCalendarURLIcsServiceUnavailable() *GetWorkforcemanagementCalendarURLIcsServiceUnavailable {
	return &GetWorkforcemanagementCalendarURLIcsServiceUnavailable{}
}

/*
GetWorkforcemanagementCalendarURLIcsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementCalendarURLIcsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement calendar Url ics service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement calendar Url ics service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementCalendarURLIcsGatewayTimeout creates a GetWorkforcemanagementCalendarURLIcsGatewayTimeout with default headers values
func NewGetWorkforcemanagementCalendarURLIcsGatewayTimeout() *GetWorkforcemanagementCalendarURLIcsGatewayTimeout {
	return &GetWorkforcemanagementCalendarURLIcsGatewayTimeout{}
}

/*
GetWorkforcemanagementCalendarURLIcsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementCalendarURLIcsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement calendar Url ics gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement calendar Url ics gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement calendar Url ics gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement calendar Url ics gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement calendar Url ics gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/calendar/url/ics][%d] getWorkforcemanagementCalendarUrlIcsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementCalendarURLIcsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
