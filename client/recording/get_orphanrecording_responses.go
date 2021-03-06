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

// GetOrphanrecordingReader is a Reader for the GetOrphanrecording structure.
type GetOrphanrecordingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrphanrecordingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrphanrecordingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrphanrecordingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrphanrecordingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrphanrecordingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrphanrecordingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrphanrecordingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrphanrecordingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrphanrecordingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrphanrecordingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrphanrecordingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrphanrecordingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrphanrecordingOK creates a GetOrphanrecordingOK with default headers values
func NewGetOrphanrecordingOK() *GetOrphanrecordingOK {
	return &GetOrphanrecordingOK{}
}

/*GetOrphanrecordingOK handles this case with default header values.

successful operation
*/
type GetOrphanrecordingOK struct {
	Payload *models.OrphanRecording
}

func (o *GetOrphanrecordingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingOK  %+v", 200, o.Payload)
}

func (o *GetOrphanrecordingOK) GetPayload() *models.OrphanRecording {
	return o.Payload
}

func (o *GetOrphanrecordingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrphanRecording)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingBadRequest creates a GetOrphanrecordingBadRequest with default headers values
func NewGetOrphanrecordingBadRequest() *GetOrphanrecordingBadRequest {
	return &GetOrphanrecordingBadRequest{}
}

/*GetOrphanrecordingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrphanrecordingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrphanrecordingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingUnauthorized creates a GetOrphanrecordingUnauthorized with default headers values
func NewGetOrphanrecordingUnauthorized() *GetOrphanrecordingUnauthorized {
	return &GetOrphanrecordingUnauthorized{}
}

/*GetOrphanrecordingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrphanrecordingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrphanrecordingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingForbidden creates a GetOrphanrecordingForbidden with default headers values
func NewGetOrphanrecordingForbidden() *GetOrphanrecordingForbidden {
	return &GetOrphanrecordingForbidden{}
}

/*GetOrphanrecordingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOrphanrecordingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingForbidden  %+v", 403, o.Payload)
}

func (o *GetOrphanrecordingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingNotFound creates a GetOrphanrecordingNotFound with default headers values
func NewGetOrphanrecordingNotFound() *GetOrphanrecordingNotFound {
	return &GetOrphanrecordingNotFound{}
}

/*GetOrphanrecordingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOrphanrecordingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingNotFound  %+v", 404, o.Payload)
}

func (o *GetOrphanrecordingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingRequestEntityTooLarge creates a GetOrphanrecordingRequestEntityTooLarge with default headers values
func NewGetOrphanrecordingRequestEntityTooLarge() *GetOrphanrecordingRequestEntityTooLarge {
	return &GetOrphanrecordingRequestEntityTooLarge{}
}

/*GetOrphanrecordingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOrphanrecordingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrphanrecordingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingUnsupportedMediaType creates a GetOrphanrecordingUnsupportedMediaType with default headers values
func NewGetOrphanrecordingUnsupportedMediaType() *GetOrphanrecordingUnsupportedMediaType {
	return &GetOrphanrecordingUnsupportedMediaType{}
}

/*GetOrphanrecordingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrphanrecordingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrphanrecordingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingTooManyRequests creates a GetOrphanrecordingTooManyRequests with default headers values
func NewGetOrphanrecordingTooManyRequests() *GetOrphanrecordingTooManyRequests {
	return &GetOrphanrecordingTooManyRequests{}
}

/*GetOrphanrecordingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOrphanrecordingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrphanrecordingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingInternalServerError creates a GetOrphanrecordingInternalServerError with default headers values
func NewGetOrphanrecordingInternalServerError() *GetOrphanrecordingInternalServerError {
	return &GetOrphanrecordingInternalServerError{}
}

/*GetOrphanrecordingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrphanrecordingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrphanrecordingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingServiceUnavailable creates a GetOrphanrecordingServiceUnavailable with default headers values
func NewGetOrphanrecordingServiceUnavailable() *GetOrphanrecordingServiceUnavailable {
	return &GetOrphanrecordingServiceUnavailable{}
}

/*GetOrphanrecordingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrphanrecordingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrphanrecordingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrphanrecordingGatewayTimeout creates a GetOrphanrecordingGatewayTimeout with default headers values
func NewGetOrphanrecordingGatewayTimeout() *GetOrphanrecordingGatewayTimeout {
	return &GetOrphanrecordingGatewayTimeout{}
}

/*GetOrphanrecordingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOrphanrecordingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOrphanrecordingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orphanrecordings/{orphanId}][%d] getOrphanrecordingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrphanrecordingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrphanrecordingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
