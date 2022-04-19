// Code generated by go-swagger; DO NOT EDIT.

package coaching

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostCoachingAppointmentsAggregatesQueryReader is a Reader for the PostCoachingAppointmentsAggregatesQuery structure.
type PostCoachingAppointmentsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCoachingAppointmentsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCoachingAppointmentsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCoachingAppointmentsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCoachingAppointmentsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCoachingAppointmentsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCoachingAppointmentsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostCoachingAppointmentsAggregatesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostCoachingAppointmentsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostCoachingAppointmentsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCoachingAppointmentsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCoachingAppointmentsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCoachingAppointmentsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCoachingAppointmentsAggregatesQueryOK creates a PostCoachingAppointmentsAggregatesQueryOK with default headers values
func NewPostCoachingAppointmentsAggregatesQueryOK() *PostCoachingAppointmentsAggregatesQueryOK {
	return &PostCoachingAppointmentsAggregatesQueryOK{}
}

/*PostCoachingAppointmentsAggregatesQueryOK handles this case with default header values.

Query completed successfully
*/
type PostCoachingAppointmentsAggregatesQueryOK struct {
	Payload *models.CoachingAppointmentAggregateResponse
}

func (o *PostCoachingAppointmentsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryOK) GetPayload() *models.CoachingAppointmentAggregateResponse {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoachingAppointmentAggregateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryBadRequest creates a PostCoachingAppointmentsAggregatesQueryBadRequest with default headers values
func NewPostCoachingAppointmentsAggregatesQueryBadRequest() *PostCoachingAppointmentsAggregatesQueryBadRequest {
	return &PostCoachingAppointmentsAggregatesQueryBadRequest{}
}

/*PostCoachingAppointmentsAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostCoachingAppointmentsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryUnauthorized creates a PostCoachingAppointmentsAggregatesQueryUnauthorized with default headers values
func NewPostCoachingAppointmentsAggregatesQueryUnauthorized() *PostCoachingAppointmentsAggregatesQueryUnauthorized {
	return &PostCoachingAppointmentsAggregatesQueryUnauthorized{}
}

/*PostCoachingAppointmentsAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostCoachingAppointmentsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryForbidden creates a PostCoachingAppointmentsAggregatesQueryForbidden with default headers values
func NewPostCoachingAppointmentsAggregatesQueryForbidden() *PostCoachingAppointmentsAggregatesQueryForbidden {
	return &PostCoachingAppointmentsAggregatesQueryForbidden{}
}

/*PostCoachingAppointmentsAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostCoachingAppointmentsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryNotFound creates a PostCoachingAppointmentsAggregatesQueryNotFound with default headers values
func NewPostCoachingAppointmentsAggregatesQueryNotFound() *PostCoachingAppointmentsAggregatesQueryNotFound {
	return &PostCoachingAppointmentsAggregatesQueryNotFound{}
}

/*PostCoachingAppointmentsAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostCoachingAppointmentsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryRequestTimeout creates a PostCoachingAppointmentsAggregatesQueryRequestTimeout with default headers values
func NewPostCoachingAppointmentsAggregatesQueryRequestTimeout() *PostCoachingAppointmentsAggregatesQueryRequestTimeout {
	return &PostCoachingAppointmentsAggregatesQueryRequestTimeout{}
}

/*PostCoachingAppointmentsAggregatesQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostCoachingAppointmentsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge creates a PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge() *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge {
	return &PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge{}
}

/*PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryUnsupportedMediaType creates a PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostCoachingAppointmentsAggregatesQueryUnsupportedMediaType() *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType {
	return &PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType{}
}

/*PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryTooManyRequests creates a PostCoachingAppointmentsAggregatesQueryTooManyRequests with default headers values
func NewPostCoachingAppointmentsAggregatesQueryTooManyRequests() *PostCoachingAppointmentsAggregatesQueryTooManyRequests {
	return &PostCoachingAppointmentsAggregatesQueryTooManyRequests{}
}

/*PostCoachingAppointmentsAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostCoachingAppointmentsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryInternalServerError creates a PostCoachingAppointmentsAggregatesQueryInternalServerError with default headers values
func NewPostCoachingAppointmentsAggregatesQueryInternalServerError() *PostCoachingAppointmentsAggregatesQueryInternalServerError {
	return &PostCoachingAppointmentsAggregatesQueryInternalServerError{}
}

/*PostCoachingAppointmentsAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostCoachingAppointmentsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryServiceUnavailable creates a PostCoachingAppointmentsAggregatesQueryServiceUnavailable with default headers values
func NewPostCoachingAppointmentsAggregatesQueryServiceUnavailable() *PostCoachingAppointmentsAggregatesQueryServiceUnavailable {
	return &PostCoachingAppointmentsAggregatesQueryServiceUnavailable{}
}

/*PostCoachingAppointmentsAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostCoachingAppointmentsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCoachingAppointmentsAggregatesQueryGatewayTimeout creates a PostCoachingAppointmentsAggregatesQueryGatewayTimeout with default headers values
func NewPostCoachingAppointmentsAggregatesQueryGatewayTimeout() *PostCoachingAppointmentsAggregatesQueryGatewayTimeout {
	return &PostCoachingAppointmentsAggregatesQueryGatewayTimeout{}
}

/*PostCoachingAppointmentsAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostCoachingAppointmentsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
