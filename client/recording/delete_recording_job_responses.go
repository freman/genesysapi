// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteRecordingJobReader is a Reader for the DeleteRecordingJob structure.
type DeleteRecordingJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordingJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRecordingJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRecordingJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRecordingJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordingJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordingJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRecordingJobRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRecordingJobRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRecordingJobUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRecordingJobTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordingJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRecordingJobServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRecordingJobGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRecordingJobNoContent creates a DeleteRecordingJobNoContent with default headers values
func NewDeleteRecordingJobNoContent() *DeleteRecordingJobNoContent {
	return &DeleteRecordingJobNoContent{}
}

/*DeleteRecordingJobNoContent handles this case with default header values.

Operation was successful.
*/
type DeleteRecordingJobNoContent struct {
}

func (o *DeleteRecordingJobNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobNoContent ", 204)
}

func (o *DeleteRecordingJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecordingJobBadRequest creates a DeleteRecordingJobBadRequest with default headers values
func NewDeleteRecordingJobBadRequest() *DeleteRecordingJobBadRequest {
	return &DeleteRecordingJobBadRequest{}
}

/*DeleteRecordingJobBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRecordingJobBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRecordingJobBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobUnauthorized creates a DeleteRecordingJobUnauthorized with default headers values
func NewDeleteRecordingJobUnauthorized() *DeleteRecordingJobUnauthorized {
	return &DeleteRecordingJobUnauthorized{}
}

/*DeleteRecordingJobUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRecordingJobUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRecordingJobUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobForbidden creates a DeleteRecordingJobForbidden with default headers values
func NewDeleteRecordingJobForbidden() *DeleteRecordingJobForbidden {
	return &DeleteRecordingJobForbidden{}
}

/*DeleteRecordingJobForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRecordingJobForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRecordingJobForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobNotFound creates a DeleteRecordingJobNotFound with default headers values
func NewDeleteRecordingJobNotFound() *DeleteRecordingJobNotFound {
	return &DeleteRecordingJobNotFound{}
}

/*DeleteRecordingJobNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRecordingJobNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRecordingJobNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobRequestTimeout creates a DeleteRecordingJobRequestTimeout with default headers values
func NewDeleteRecordingJobRequestTimeout() *DeleteRecordingJobRequestTimeout {
	return &DeleteRecordingJobRequestTimeout{}
}

/*DeleteRecordingJobRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRecordingJobRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRecordingJobRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobRequestEntityTooLarge creates a DeleteRecordingJobRequestEntityTooLarge with default headers values
func NewDeleteRecordingJobRequestEntityTooLarge() *DeleteRecordingJobRequestEntityTooLarge {
	return &DeleteRecordingJobRequestEntityTooLarge{}
}

/*DeleteRecordingJobRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRecordingJobRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRecordingJobRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobUnsupportedMediaType creates a DeleteRecordingJobUnsupportedMediaType with default headers values
func NewDeleteRecordingJobUnsupportedMediaType() *DeleteRecordingJobUnsupportedMediaType {
	return &DeleteRecordingJobUnsupportedMediaType{}
}

/*DeleteRecordingJobUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRecordingJobUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRecordingJobUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobTooManyRequests creates a DeleteRecordingJobTooManyRequests with default headers values
func NewDeleteRecordingJobTooManyRequests() *DeleteRecordingJobTooManyRequests {
	return &DeleteRecordingJobTooManyRequests{}
}

/*DeleteRecordingJobTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRecordingJobTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRecordingJobTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobInternalServerError creates a DeleteRecordingJobInternalServerError with default headers values
func NewDeleteRecordingJobInternalServerError() *DeleteRecordingJobInternalServerError {
	return &DeleteRecordingJobInternalServerError{}
}

/*DeleteRecordingJobInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRecordingJobInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRecordingJobInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobServiceUnavailable creates a DeleteRecordingJobServiceUnavailable with default headers values
func NewDeleteRecordingJobServiceUnavailable() *DeleteRecordingJobServiceUnavailable {
	return &DeleteRecordingJobServiceUnavailable{}
}

/*DeleteRecordingJobServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRecordingJobServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRecordingJobServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingJobGatewayTimeout creates a DeleteRecordingJobGatewayTimeout with default headers values
func NewDeleteRecordingJobGatewayTimeout() *DeleteRecordingJobGatewayTimeout {
	return &DeleteRecordingJobGatewayTimeout{}
}

/*DeleteRecordingJobGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRecordingJobGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingJobGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/jobs/{jobId}][%d] deleteRecordingJobGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRecordingJobGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingJobGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
