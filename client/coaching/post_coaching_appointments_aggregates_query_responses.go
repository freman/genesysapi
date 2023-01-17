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

/*
PostCoachingAppointmentsAggregatesQueryOK describes a response with status code 200, with default header values.

Query completed successfully
*/
type PostCoachingAppointmentsAggregatesQueryOK struct {
	Payload *models.CoachingAppointmentAggregateResponse
}

// IsSuccess returns true when this post coaching appointments aggregates query o k response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post coaching appointments aggregates query o k response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query o k response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post coaching appointments aggregates query o k response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query o k response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostCoachingAppointmentsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryOK) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostCoachingAppointmentsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query bad request response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query bad request response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query bad request response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query bad request response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query bad request response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryBadRequest) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostCoachingAppointmentsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query unauthorized response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query unauthorized response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query unauthorized response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query unauthorized response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query unauthorized response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryUnauthorized) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostCoachingAppointmentsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query forbidden response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query forbidden response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query forbidden response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query forbidden response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query forbidden response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostCoachingAppointmentsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryForbidden) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostCoachingAppointmentsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query not found response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query not found response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query not found response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query not found response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query not found response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostCoachingAppointmentsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryNotFound) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostCoachingAppointmentsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query request timeout response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query request timeout response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query request timeout response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query request timeout response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query request timeout response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestTimeout) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query request entity too large response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query request entity too large response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query request entity too large response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query request entity too large response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query request entity too large response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryRequestEntityTooLarge) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query unsupported media type response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query unsupported media type response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query unsupported media type response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query unsupported media type response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query unsupported media type response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryUnsupportedMediaType) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostCoachingAppointmentsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query too many requests response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query too many requests response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query too many requests response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post coaching appointments aggregates query too many requests response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post coaching appointments aggregates query too many requests response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryTooManyRequests) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostCoachingAppointmentsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query internal server error response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query internal server error response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query internal server error response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post coaching appointments aggregates query internal server error response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post coaching appointments aggregates query internal server error response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryInternalServerError) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostCoachingAppointmentsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query service unavailable response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query service unavailable response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query service unavailable response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post coaching appointments aggregates query service unavailable response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post coaching appointments aggregates query service unavailable response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryServiceUnavailable) String() string {
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

/*
PostCoachingAppointmentsAggregatesQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostCoachingAppointmentsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post coaching appointments aggregates query gateway timeout response has a 2xx status code
func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post coaching appointments aggregates query gateway timeout response has a 3xx status code
func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post coaching appointments aggregates query gateway timeout response has a 4xx status code
func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post coaching appointments aggregates query gateway timeout response has a 5xx status code
func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post coaching appointments aggregates query gateway timeout response a status code equal to that given
func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/coaching/appointments/aggregates/query][%d] postCoachingAppointmentsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCoachingAppointmentsAggregatesQueryGatewayTimeout) String() string {
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
