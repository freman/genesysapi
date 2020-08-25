// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAnalyticsConversationsDetailsJobsReader is a Reader for the PostAnalyticsConversationsDetailsJobs structure.
type PostAnalyticsConversationsDetailsJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsConversationsDetailsJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostAnalyticsConversationsDetailsJobsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsConversationsDetailsJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsConversationsDetailsJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsConversationsDetailsJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsConversationsDetailsJobsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsConversationsDetailsJobsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsConversationsDetailsJobsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsConversationsDetailsJobsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsConversationsDetailsJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsConversationsDetailsJobsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsConversationsDetailsJobsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsConversationsDetailsJobsAccepted creates a PostAnalyticsConversationsDetailsJobsAccepted with default headers values
func NewPostAnalyticsConversationsDetailsJobsAccepted() *PostAnalyticsConversationsDetailsJobsAccepted {
	return &PostAnalyticsConversationsDetailsJobsAccepted{}
}

/*PostAnalyticsConversationsDetailsJobsAccepted handles this case with default header values.

Accepted - Running query asynchronously
*/
type PostAnalyticsConversationsDetailsJobsAccepted struct {
	Payload *models.AsyncQueryResponse
}

func (o *PostAnalyticsConversationsDetailsJobsAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsAccepted  %+v", 202, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsAccepted) GetPayload() *models.AsyncQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsBadRequest creates a PostAnalyticsConversationsDetailsJobsBadRequest with default headers values
func NewPostAnalyticsConversationsDetailsJobsBadRequest() *PostAnalyticsConversationsDetailsJobsBadRequest {
	return &PostAnalyticsConversationsDetailsJobsBadRequest{}
}

/*PostAnalyticsConversationsDetailsJobsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsConversationsDetailsJobsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsUnauthorized creates a PostAnalyticsConversationsDetailsJobsUnauthorized with default headers values
func NewPostAnalyticsConversationsDetailsJobsUnauthorized() *PostAnalyticsConversationsDetailsJobsUnauthorized {
	return &PostAnalyticsConversationsDetailsJobsUnauthorized{}
}

/*PostAnalyticsConversationsDetailsJobsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsConversationsDetailsJobsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsForbidden creates a PostAnalyticsConversationsDetailsJobsForbidden with default headers values
func NewPostAnalyticsConversationsDetailsJobsForbidden() *PostAnalyticsConversationsDetailsJobsForbidden {
	return &PostAnalyticsConversationsDetailsJobsForbidden{}
}

/*PostAnalyticsConversationsDetailsJobsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsConversationsDetailsJobsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsNotFound creates a PostAnalyticsConversationsDetailsJobsNotFound with default headers values
func NewPostAnalyticsConversationsDetailsJobsNotFound() *PostAnalyticsConversationsDetailsJobsNotFound {
	return &PostAnalyticsConversationsDetailsJobsNotFound{}
}

/*PostAnalyticsConversationsDetailsJobsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsConversationsDetailsJobsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsRequestEntityTooLarge creates a PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge with default headers values
func NewPostAnalyticsConversationsDetailsJobsRequestEntityTooLarge() *PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge {
	return &PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge{}
}

/*PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsUnsupportedMediaType creates a PostAnalyticsConversationsDetailsJobsUnsupportedMediaType with default headers values
func NewPostAnalyticsConversationsDetailsJobsUnsupportedMediaType() *PostAnalyticsConversationsDetailsJobsUnsupportedMediaType {
	return &PostAnalyticsConversationsDetailsJobsUnsupportedMediaType{}
}

/*PostAnalyticsConversationsDetailsJobsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsConversationsDetailsJobsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsTooManyRequests creates a PostAnalyticsConversationsDetailsJobsTooManyRequests with default headers values
func NewPostAnalyticsConversationsDetailsJobsTooManyRequests() *PostAnalyticsConversationsDetailsJobsTooManyRequests {
	return &PostAnalyticsConversationsDetailsJobsTooManyRequests{}
}

/*PostAnalyticsConversationsDetailsJobsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsConversationsDetailsJobsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsInternalServerError creates a PostAnalyticsConversationsDetailsJobsInternalServerError with default headers values
func NewPostAnalyticsConversationsDetailsJobsInternalServerError() *PostAnalyticsConversationsDetailsJobsInternalServerError {
	return &PostAnalyticsConversationsDetailsJobsInternalServerError{}
}

/*PostAnalyticsConversationsDetailsJobsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsConversationsDetailsJobsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsServiceUnavailable creates a PostAnalyticsConversationsDetailsJobsServiceUnavailable with default headers values
func NewPostAnalyticsConversationsDetailsJobsServiceUnavailable() *PostAnalyticsConversationsDetailsJobsServiceUnavailable {
	return &PostAnalyticsConversationsDetailsJobsServiceUnavailable{}
}

/*PostAnalyticsConversationsDetailsJobsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsConversationsDetailsJobsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsDetailsJobsGatewayTimeout creates a PostAnalyticsConversationsDetailsJobsGatewayTimeout with default headers values
func NewPostAnalyticsConversationsDetailsJobsGatewayTimeout() *PostAnalyticsConversationsDetailsJobsGatewayTimeout {
	return &PostAnalyticsConversationsDetailsJobsGatewayTimeout{}
}

/*PostAnalyticsConversationsDetailsJobsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsConversationsDetailsJobsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsDetailsJobsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/details/jobs][%d] postAnalyticsConversationsDetailsJobsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsConversationsDetailsJobsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsDetailsJobsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}