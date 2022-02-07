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

// DeleteWorkforcemanagementCalendarURLIcsReader is a Reader for the DeleteWorkforcemanagementCalendarURLIcs structure.
type DeleteWorkforcemanagementCalendarURLIcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWorkforcemanagementCalendarURLIcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWorkforcemanagementCalendarURLIcsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWorkforcemanagementCalendarURLIcsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteWorkforcemanagementCalendarURLIcsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteWorkforcemanagementCalendarURLIcsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWorkforcemanagementCalendarURLIcsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteWorkforcemanagementCalendarURLIcsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteWorkforcemanagementCalendarURLIcsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteWorkforcemanagementCalendarURLIcsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteWorkforcemanagementCalendarURLIcsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteWorkforcemanagementCalendarURLIcsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteWorkforcemanagementCalendarURLIcsNoContent creates a DeleteWorkforcemanagementCalendarURLIcsNoContent with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsNoContent() *DeleteWorkforcemanagementCalendarURLIcsNoContent {
	return &DeleteWorkforcemanagementCalendarURLIcsNoContent{}
}

/*DeleteWorkforcemanagementCalendarURLIcsNoContent handles this case with default header values.

Operation was successful.
*/
type DeleteWorkforcemanagementCalendarURLIcsNoContent struct {
}

func (o *DeleteWorkforcemanagementCalendarURLIcsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsNoContent ", 204)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsBadRequest creates a DeleteWorkforcemanagementCalendarURLIcsBadRequest with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsBadRequest() *DeleteWorkforcemanagementCalendarURLIcsBadRequest {
	return &DeleteWorkforcemanagementCalendarURLIcsBadRequest{}
}

/*DeleteWorkforcemanagementCalendarURLIcsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteWorkforcemanagementCalendarURLIcsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsUnauthorized creates a DeleteWorkforcemanagementCalendarURLIcsUnauthorized with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsUnauthorized() *DeleteWorkforcemanagementCalendarURLIcsUnauthorized {
	return &DeleteWorkforcemanagementCalendarURLIcsUnauthorized{}
}

/*DeleteWorkforcemanagementCalendarURLIcsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteWorkforcemanagementCalendarURLIcsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsForbidden creates a DeleteWorkforcemanagementCalendarURLIcsForbidden with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsForbidden() *DeleteWorkforcemanagementCalendarURLIcsForbidden {
	return &DeleteWorkforcemanagementCalendarURLIcsForbidden{}
}

/*DeleteWorkforcemanagementCalendarURLIcsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteWorkforcemanagementCalendarURLIcsForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsNotFound creates a DeleteWorkforcemanagementCalendarURLIcsNotFound with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsNotFound() *DeleteWorkforcemanagementCalendarURLIcsNotFound {
	return &DeleteWorkforcemanagementCalendarURLIcsNotFound{}
}

/*DeleteWorkforcemanagementCalendarURLIcsNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteWorkforcemanagementCalendarURLIcsNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsRequestTimeout creates a DeleteWorkforcemanagementCalendarURLIcsRequestTimeout with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsRequestTimeout() *DeleteWorkforcemanagementCalendarURLIcsRequestTimeout {
	return &DeleteWorkforcemanagementCalendarURLIcsRequestTimeout{}
}

/*DeleteWorkforcemanagementCalendarURLIcsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteWorkforcemanagementCalendarURLIcsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge creates a DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge() *DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge {
	return &DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge{}
}

/*DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType creates a DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType() *DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType {
	return &DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType{}
}

/*DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsTooManyRequests creates a DeleteWorkforcemanagementCalendarURLIcsTooManyRequests with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsTooManyRequests() *DeleteWorkforcemanagementCalendarURLIcsTooManyRequests {
	return &DeleteWorkforcemanagementCalendarURLIcsTooManyRequests{}
}

/*DeleteWorkforcemanagementCalendarURLIcsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteWorkforcemanagementCalendarURLIcsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsInternalServerError creates a DeleteWorkforcemanagementCalendarURLIcsInternalServerError with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsInternalServerError() *DeleteWorkforcemanagementCalendarURLIcsInternalServerError {
	return &DeleteWorkforcemanagementCalendarURLIcsInternalServerError{}
}

/*DeleteWorkforcemanagementCalendarURLIcsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteWorkforcemanagementCalendarURLIcsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsServiceUnavailable creates a DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsServiceUnavailable() *DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable {
	return &DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable{}
}

/*DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWorkforcemanagementCalendarURLIcsGatewayTimeout creates a DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout with default headers values
func NewDeleteWorkforcemanagementCalendarURLIcsGatewayTimeout() *DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout {
	return &DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout{}
}

/*DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/workforcemanagement/calendar/url/ics][%d] deleteWorkforcemanagementCalendarUrlIcsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteWorkforcemanagementCalendarURLIcsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}