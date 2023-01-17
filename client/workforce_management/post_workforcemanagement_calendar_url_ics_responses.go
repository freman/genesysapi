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

// PostWorkforcemanagementCalendarURLIcsReader is a Reader for the PostWorkforcemanagementCalendarURLIcs structure.
type PostWorkforcemanagementCalendarURLIcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWorkforcemanagementCalendarURLIcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWorkforcemanagementCalendarURLIcsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostWorkforcemanagementCalendarURLIcsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWorkforcemanagementCalendarURLIcsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostWorkforcemanagementCalendarURLIcsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostWorkforcemanagementCalendarURLIcsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostWorkforcemanagementCalendarURLIcsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostWorkforcemanagementCalendarURLIcsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostWorkforcemanagementCalendarURLIcsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostWorkforcemanagementCalendarURLIcsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWorkforcemanagementCalendarURLIcsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostWorkforcemanagementCalendarURLIcsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostWorkforcemanagementCalendarURLIcsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWorkforcemanagementCalendarURLIcsOK creates a PostWorkforcemanagementCalendarURLIcsOK with default headers values
func NewPostWorkforcemanagementCalendarURLIcsOK() *PostWorkforcemanagementCalendarURLIcsOK {
	return &PostWorkforcemanagementCalendarURLIcsOK{}
}

/*
PostWorkforcemanagementCalendarURLIcsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostWorkforcemanagementCalendarURLIcsOK struct {
	Payload *models.CalendarURLResponse
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics o k response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics o k response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics o k response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement calendar Url ics o k response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics o k response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostWorkforcemanagementCalendarURLIcsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsOK  %+v", 200, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsOK) GetPayload() *models.CalendarURLResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CalendarURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsCreated creates a PostWorkforcemanagementCalendarURLIcsCreated with default headers values
func NewPostWorkforcemanagementCalendarURLIcsCreated() *PostWorkforcemanagementCalendarURLIcsCreated {
	return &PostWorkforcemanagementCalendarURLIcsCreated{}
}

/*
PostWorkforcemanagementCalendarURLIcsCreated describes a response with status code 201, with default header values.

PostWorkforcemanagementCalendarURLIcsCreated post workforcemanagement calendar Url ics created
*/
type PostWorkforcemanagementCalendarURLIcsCreated struct {
	Payload *models.CalendarURLResponse
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics created response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics created response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics created response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement calendar Url ics created response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics created response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostWorkforcemanagementCalendarURLIcsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsCreated  %+v", 201, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsCreated) GetPayload() *models.CalendarURLResponse {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CalendarURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsBadRequest creates a PostWorkforcemanagementCalendarURLIcsBadRequest with default headers values
func NewPostWorkforcemanagementCalendarURLIcsBadRequest() *PostWorkforcemanagementCalendarURLIcsBadRequest {
	return &PostWorkforcemanagementCalendarURLIcsBadRequest{}
}

/*
PostWorkforcemanagementCalendarURLIcsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostWorkforcemanagementCalendarURLIcsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics bad request response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics bad request response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics bad request response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics bad request response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics bad request response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsBadRequest  %+v", 400, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsUnauthorized creates a PostWorkforcemanagementCalendarURLIcsUnauthorized with default headers values
func NewPostWorkforcemanagementCalendarURLIcsUnauthorized() *PostWorkforcemanagementCalendarURLIcsUnauthorized {
	return &PostWorkforcemanagementCalendarURLIcsUnauthorized{}
}

/*
PostWorkforcemanagementCalendarURLIcsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostWorkforcemanagementCalendarURLIcsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics unauthorized response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics unauthorized response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics unauthorized response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics unauthorized response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics unauthorized response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsForbidden creates a PostWorkforcemanagementCalendarURLIcsForbidden with default headers values
func NewPostWorkforcemanagementCalendarURLIcsForbidden() *PostWorkforcemanagementCalendarURLIcsForbidden {
	return &PostWorkforcemanagementCalendarURLIcsForbidden{}
}

/*
PostWorkforcemanagementCalendarURLIcsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostWorkforcemanagementCalendarURLIcsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics forbidden response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics forbidden response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics forbidden response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics forbidden response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics forbidden response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostWorkforcemanagementCalendarURLIcsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsForbidden  %+v", 403, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsNotFound creates a PostWorkforcemanagementCalendarURLIcsNotFound with default headers values
func NewPostWorkforcemanagementCalendarURLIcsNotFound() *PostWorkforcemanagementCalendarURLIcsNotFound {
	return &PostWorkforcemanagementCalendarURLIcsNotFound{}
}

/*
PostWorkforcemanagementCalendarURLIcsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostWorkforcemanagementCalendarURLIcsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics not found response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics not found response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics not found response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics not found response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics not found response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostWorkforcemanagementCalendarURLIcsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsNotFound  %+v", 404, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsRequestTimeout creates a PostWorkforcemanagementCalendarURLIcsRequestTimeout with default headers values
func NewPostWorkforcemanagementCalendarURLIcsRequestTimeout() *PostWorkforcemanagementCalendarURLIcsRequestTimeout {
	return &PostWorkforcemanagementCalendarURLIcsRequestTimeout{}
}

/*
PostWorkforcemanagementCalendarURLIcsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostWorkforcemanagementCalendarURLIcsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics request timeout response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics request timeout response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics request timeout response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics request timeout response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics request timeout response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge creates a PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge with default headers values
func NewPostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge() *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge {
	return &PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge{}
}

/*
PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics request entity too large response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics request entity too large response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics request entity too large response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics request entity too large response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics request entity too large response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsUnsupportedMediaType creates a PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType with default headers values
func NewPostWorkforcemanagementCalendarURLIcsUnsupportedMediaType() *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType {
	return &PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType{}
}

/*
PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics unsupported media type response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics unsupported media type response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics unsupported media type response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics unsupported media type response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics unsupported media type response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsTooManyRequests creates a PostWorkforcemanagementCalendarURLIcsTooManyRequests with default headers values
func NewPostWorkforcemanagementCalendarURLIcsTooManyRequests() *PostWorkforcemanagementCalendarURLIcsTooManyRequests {
	return &PostWorkforcemanagementCalendarURLIcsTooManyRequests{}
}

/*
PostWorkforcemanagementCalendarURLIcsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostWorkforcemanagementCalendarURLIcsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics too many requests response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics too many requests response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics too many requests response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post workforcemanagement calendar Url ics too many requests response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post workforcemanagement calendar Url ics too many requests response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsInternalServerError creates a PostWorkforcemanagementCalendarURLIcsInternalServerError with default headers values
func NewPostWorkforcemanagementCalendarURLIcsInternalServerError() *PostWorkforcemanagementCalendarURLIcsInternalServerError {
	return &PostWorkforcemanagementCalendarURLIcsInternalServerError{}
}

/*
PostWorkforcemanagementCalendarURLIcsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostWorkforcemanagementCalendarURLIcsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics internal server error response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics internal server error response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics internal server error response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement calendar Url ics internal server error response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement calendar Url ics internal server error response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsServiceUnavailable creates a PostWorkforcemanagementCalendarURLIcsServiceUnavailable with default headers values
func NewPostWorkforcemanagementCalendarURLIcsServiceUnavailable() *PostWorkforcemanagementCalendarURLIcsServiceUnavailable {
	return &PostWorkforcemanagementCalendarURLIcsServiceUnavailable{}
}

/*
PostWorkforcemanagementCalendarURLIcsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostWorkforcemanagementCalendarURLIcsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics service unavailable response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics service unavailable response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics service unavailable response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement calendar Url ics service unavailable response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement calendar Url ics service unavailable response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWorkforcemanagementCalendarURLIcsGatewayTimeout creates a PostWorkforcemanagementCalendarURLIcsGatewayTimeout with default headers values
func NewPostWorkforcemanagementCalendarURLIcsGatewayTimeout() *PostWorkforcemanagementCalendarURLIcsGatewayTimeout {
	return &PostWorkforcemanagementCalendarURLIcsGatewayTimeout{}
}

/*
PostWorkforcemanagementCalendarURLIcsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostWorkforcemanagementCalendarURLIcsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post workforcemanagement calendar Url ics gateway timeout response has a 2xx status code
func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post workforcemanagement calendar Url ics gateway timeout response has a 3xx status code
func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post workforcemanagement calendar Url ics gateway timeout response has a 4xx status code
func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post workforcemanagement calendar Url ics gateway timeout response has a 5xx status code
func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post workforcemanagement calendar Url ics gateway timeout response a status code equal to that given
func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/workforcemanagement/calendar/url/ics][%d] postWorkforcemanagementCalendarUrlIcsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostWorkforcemanagementCalendarURLIcsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
