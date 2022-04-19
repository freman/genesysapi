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

// DeleteOrphanrecordingReader is a Reader for the DeleteOrphanrecording structure.
type DeleteOrphanrecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrphanrecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrphanrecordingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOrphanrecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOrphanrecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrphanrecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrphanrecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOrphanrecordingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOrphanrecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOrphanrecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOrphanrecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrphanrecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOrphanrecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOrphanrecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOrphanrecordingOK creates a DeleteOrphanrecordingOK with default headers values
func NewDeleteOrphanrecordingOK() *DeleteOrphanrecordingOK {
	return &DeleteOrphanrecordingOK{}
}

/*DeleteOrphanrecordingOK handles this case with default header values.

successful operation
*/
type DeleteOrphanrecordingOK struct {
	Payload *models.OrphanRecording
}

func (o *DeleteOrphanrecordingOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingOK  %+v", 200, o.Payload)
}

func (o *DeleteOrphanrecordingOK) GetPayload() *models.OrphanRecording {
	return o.Payload
}

func (o *DeleteOrphanrecordingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrphanRecording)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingBadRequest creates a DeleteOrphanrecordingBadRequest with default headers values
func NewDeleteOrphanrecordingBadRequest() *DeleteOrphanrecordingBadRequest {
	return &DeleteOrphanrecordingBadRequest{}
}

/*DeleteOrphanrecordingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOrphanrecordingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrphanrecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingUnauthorized creates a DeleteOrphanrecordingUnauthorized with default headers values
func NewDeleteOrphanrecordingUnauthorized() *DeleteOrphanrecordingUnauthorized {
	return &DeleteOrphanrecordingUnauthorized{}
}

/*DeleteOrphanrecordingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOrphanrecordingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrphanrecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingForbidden creates a DeleteOrphanrecordingForbidden with default headers values
func NewDeleteOrphanrecordingForbidden() *DeleteOrphanrecordingForbidden {
	return &DeleteOrphanrecordingForbidden{}
}

/*DeleteOrphanrecordingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOrphanrecordingForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrphanrecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingNotFound creates a DeleteOrphanrecordingNotFound with default headers values
func NewDeleteOrphanrecordingNotFound() *DeleteOrphanrecordingNotFound {
	return &DeleteOrphanrecordingNotFound{}
}

/*DeleteOrphanrecordingNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOrphanrecordingNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrphanrecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingRequestTimeout creates a DeleteOrphanrecordingRequestTimeout with default headers values
func NewDeleteOrphanrecordingRequestTimeout() *DeleteOrphanrecordingRequestTimeout {
	return &DeleteOrphanrecordingRequestTimeout{}
}

/*DeleteOrphanrecordingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOrphanrecordingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOrphanrecordingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingRequestEntityTooLarge creates a DeleteOrphanrecordingRequestEntityTooLarge with default headers values
func NewDeleteOrphanrecordingRequestEntityTooLarge() *DeleteOrphanrecordingRequestEntityTooLarge {
	return &DeleteOrphanrecordingRequestEntityTooLarge{}
}

/*DeleteOrphanrecordingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteOrphanrecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOrphanrecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingUnsupportedMediaType creates a DeleteOrphanrecordingUnsupportedMediaType with default headers values
func NewDeleteOrphanrecordingUnsupportedMediaType() *DeleteOrphanrecordingUnsupportedMediaType {
	return &DeleteOrphanrecordingUnsupportedMediaType{}
}

/*DeleteOrphanrecordingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOrphanrecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOrphanrecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingTooManyRequests creates a DeleteOrphanrecordingTooManyRequests with default headers values
func NewDeleteOrphanrecordingTooManyRequests() *DeleteOrphanrecordingTooManyRequests {
	return &DeleteOrphanrecordingTooManyRequests{}
}

/*DeleteOrphanrecordingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOrphanrecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOrphanrecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingInternalServerError creates a DeleteOrphanrecordingInternalServerError with default headers values
func NewDeleteOrphanrecordingInternalServerError() *DeleteOrphanrecordingInternalServerError {
	return &DeleteOrphanrecordingInternalServerError{}
}

/*DeleteOrphanrecordingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOrphanrecordingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrphanrecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingServiceUnavailable creates a DeleteOrphanrecordingServiceUnavailable with default headers values
func NewDeleteOrphanrecordingServiceUnavailable() *DeleteOrphanrecordingServiceUnavailable {
	return &DeleteOrphanrecordingServiceUnavailable{}
}

/*DeleteOrphanrecordingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOrphanrecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOrphanrecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrphanrecordingGatewayTimeout creates a DeleteOrphanrecordingGatewayTimeout with default headers values
func NewDeleteOrphanrecordingGatewayTimeout() *DeleteOrphanrecordingGatewayTimeout {
	return &DeleteOrphanrecordingGatewayTimeout{}
}

/*DeleteOrphanrecordingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOrphanrecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOrphanrecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/orphanrecordings/{orphanId}][%d] deleteOrphanrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOrphanrecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOrphanrecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
