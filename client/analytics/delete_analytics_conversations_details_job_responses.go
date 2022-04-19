// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteAnalyticsConversationsDetailsJobReader is a Reader for the DeleteAnalyticsConversationsDetailsJob structure.
type DeleteAnalyticsConversationsDetailsJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAnalyticsConversationsDetailsJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAnalyticsConversationsDetailsJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAnalyticsConversationsDetailsJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAnalyticsConversationsDetailsJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAnalyticsConversationsDetailsJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAnalyticsConversationsDetailsJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteAnalyticsConversationsDetailsJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteAnalyticsConversationsDetailsJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAnalyticsConversationsDetailsJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAnalyticsConversationsDetailsJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteAnalyticsConversationsDetailsJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteAnalyticsConversationsDetailsJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAnalyticsConversationsDetailsJobNoContent creates a DeleteAnalyticsConversationsDetailsJobNoContent with default headers values
func NewDeleteAnalyticsConversationsDetailsJobNoContent() *DeleteAnalyticsConversationsDetailsJobNoContent {
	return &DeleteAnalyticsConversationsDetailsJobNoContent{}
}

/*DeleteAnalyticsConversationsDetailsJobNoContent handles this case with default header values.

Deleted
*/
type DeleteAnalyticsConversationsDetailsJobNoContent struct {
}

func (o *DeleteAnalyticsConversationsDetailsJobNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobNoContent ", 204)
}

func (o *DeleteAnalyticsConversationsDetailsJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobBadRequest creates a DeleteAnalyticsConversationsDetailsJobBadRequest with default headers values
func NewDeleteAnalyticsConversationsDetailsJobBadRequest() *DeleteAnalyticsConversationsDetailsJobBadRequest {
	return &DeleteAnalyticsConversationsDetailsJobBadRequest{}
}

/*DeleteAnalyticsConversationsDetailsJobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteAnalyticsConversationsDetailsJobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobUnauthorized creates a DeleteAnalyticsConversationsDetailsJobUnauthorized with default headers values
func NewDeleteAnalyticsConversationsDetailsJobUnauthorized() *DeleteAnalyticsConversationsDetailsJobUnauthorized {
	return &DeleteAnalyticsConversationsDetailsJobUnauthorized{}
}

/*DeleteAnalyticsConversationsDetailsJobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteAnalyticsConversationsDetailsJobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobForbidden creates a DeleteAnalyticsConversationsDetailsJobForbidden with default headers values
func NewDeleteAnalyticsConversationsDetailsJobForbidden() *DeleteAnalyticsConversationsDetailsJobForbidden {
	return &DeleteAnalyticsConversationsDetailsJobForbidden{}
}

/*DeleteAnalyticsConversationsDetailsJobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteAnalyticsConversationsDetailsJobForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobNotFound creates a DeleteAnalyticsConversationsDetailsJobNotFound with default headers values
func NewDeleteAnalyticsConversationsDetailsJobNotFound() *DeleteAnalyticsConversationsDetailsJobNotFound {
	return &DeleteAnalyticsConversationsDetailsJobNotFound{}
}

/*DeleteAnalyticsConversationsDetailsJobNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteAnalyticsConversationsDetailsJobNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobRequestTimeout creates a DeleteAnalyticsConversationsDetailsJobRequestTimeout with default headers values
func NewDeleteAnalyticsConversationsDetailsJobRequestTimeout() *DeleteAnalyticsConversationsDetailsJobRequestTimeout {
	return &DeleteAnalyticsConversationsDetailsJobRequestTimeout{}
}

/*DeleteAnalyticsConversationsDetailsJobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteAnalyticsConversationsDetailsJobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge creates a DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge with default headers values
func NewDeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge() *DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge {
	return &DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge{}
}

/*DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobUnsupportedMediaType creates a DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType with default headers values
func NewDeleteAnalyticsConversationsDetailsJobUnsupportedMediaType() *DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType {
	return &DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType{}
}

/*DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobTooManyRequests creates a DeleteAnalyticsConversationsDetailsJobTooManyRequests with default headers values
func NewDeleteAnalyticsConversationsDetailsJobTooManyRequests() *DeleteAnalyticsConversationsDetailsJobTooManyRequests {
	return &DeleteAnalyticsConversationsDetailsJobTooManyRequests{}
}

/*DeleteAnalyticsConversationsDetailsJobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteAnalyticsConversationsDetailsJobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobInternalServerError creates a DeleteAnalyticsConversationsDetailsJobInternalServerError with default headers values
func NewDeleteAnalyticsConversationsDetailsJobInternalServerError() *DeleteAnalyticsConversationsDetailsJobInternalServerError {
	return &DeleteAnalyticsConversationsDetailsJobInternalServerError{}
}

/*DeleteAnalyticsConversationsDetailsJobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteAnalyticsConversationsDetailsJobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobServiceUnavailable creates a DeleteAnalyticsConversationsDetailsJobServiceUnavailable with default headers values
func NewDeleteAnalyticsConversationsDetailsJobServiceUnavailable() *DeleteAnalyticsConversationsDetailsJobServiceUnavailable {
	return &DeleteAnalyticsConversationsDetailsJobServiceUnavailable{}
}

/*DeleteAnalyticsConversationsDetailsJobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteAnalyticsConversationsDetailsJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalyticsConversationsDetailsJobGatewayTimeout creates a DeleteAnalyticsConversationsDetailsJobGatewayTimeout with default headers values
func NewDeleteAnalyticsConversationsDetailsJobGatewayTimeout() *DeleteAnalyticsConversationsDetailsJobGatewayTimeout {
	return &DeleteAnalyticsConversationsDetailsJobGatewayTimeout{}
}

/*DeleteAnalyticsConversationsDetailsJobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteAnalyticsConversationsDetailsJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteAnalyticsConversationsDetailsJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/analytics/conversations/details/jobs/{jobId}][%d] deleteAnalyticsConversationsDetailsJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteAnalyticsConversationsDetailsJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteAnalyticsConversationsDetailsJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
