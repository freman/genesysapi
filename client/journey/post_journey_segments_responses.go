// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostJourneySegmentsReader is a Reader for the PostJourneySegments structure.
type PostJourneySegmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostJourneySegmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostJourneySegmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostJourneySegmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostJourneySegmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostJourneySegmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostJourneySegmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostJourneySegmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostJourneySegmentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostJourneySegmentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostJourneySegmentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostJourneySegmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostJourneySegmentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostJourneySegmentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostJourneySegmentsOK creates a PostJourneySegmentsOK with default headers values
func NewPostJourneySegmentsOK() *PostJourneySegmentsOK {
	return &PostJourneySegmentsOK{}
}

/*PostJourneySegmentsOK handles this case with default header values.

successful operation
*/
type PostJourneySegmentsOK struct {
	Payload *models.JourneySegment
}

func (o *PostJourneySegmentsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsOK  %+v", 200, o.Payload)
}

func (o *PostJourneySegmentsOK) GetPayload() *models.JourneySegment {
	return o.Payload
}

func (o *PostJourneySegmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JourneySegment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsCreated creates a PostJourneySegmentsCreated with default headers values
func NewPostJourneySegmentsCreated() *PostJourneySegmentsCreated {
	return &PostJourneySegmentsCreated{}
}

/*PostJourneySegmentsCreated handles this case with default header values.

Segment created.
*/
type PostJourneySegmentsCreated struct {
	Payload *models.JourneySegment
}

func (o *PostJourneySegmentsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsCreated  %+v", 201, o.Payload)
}

func (o *PostJourneySegmentsCreated) GetPayload() *models.JourneySegment {
	return o.Payload
}

func (o *PostJourneySegmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JourneySegment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsBadRequest creates a PostJourneySegmentsBadRequest with default headers values
func NewPostJourneySegmentsBadRequest() *PostJourneySegmentsBadRequest {
	return &PostJourneySegmentsBadRequest{}
}

/*PostJourneySegmentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostJourneySegmentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostJourneySegmentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsUnauthorized creates a PostJourneySegmentsUnauthorized with default headers values
func NewPostJourneySegmentsUnauthorized() *PostJourneySegmentsUnauthorized {
	return &PostJourneySegmentsUnauthorized{}
}

/*PostJourneySegmentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostJourneySegmentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostJourneySegmentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsForbidden creates a PostJourneySegmentsForbidden with default headers values
func NewPostJourneySegmentsForbidden() *PostJourneySegmentsForbidden {
	return &PostJourneySegmentsForbidden{}
}

/*PostJourneySegmentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostJourneySegmentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsForbidden  %+v", 403, o.Payload)
}

func (o *PostJourneySegmentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsNotFound creates a PostJourneySegmentsNotFound with default headers values
func NewPostJourneySegmentsNotFound() *PostJourneySegmentsNotFound {
	return &PostJourneySegmentsNotFound{}
}

/*PostJourneySegmentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostJourneySegmentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsNotFound  %+v", 404, o.Payload)
}

func (o *PostJourneySegmentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsRequestEntityTooLarge creates a PostJourneySegmentsRequestEntityTooLarge with default headers values
func NewPostJourneySegmentsRequestEntityTooLarge() *PostJourneySegmentsRequestEntityTooLarge {
	return &PostJourneySegmentsRequestEntityTooLarge{}
}

/*PostJourneySegmentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostJourneySegmentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostJourneySegmentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsUnsupportedMediaType creates a PostJourneySegmentsUnsupportedMediaType with default headers values
func NewPostJourneySegmentsUnsupportedMediaType() *PostJourneySegmentsUnsupportedMediaType {
	return &PostJourneySegmentsUnsupportedMediaType{}
}

/*PostJourneySegmentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostJourneySegmentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostJourneySegmentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsTooManyRequests creates a PostJourneySegmentsTooManyRequests with default headers values
func NewPostJourneySegmentsTooManyRequests() *PostJourneySegmentsTooManyRequests {
	return &PostJourneySegmentsTooManyRequests{}
}

/*PostJourneySegmentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostJourneySegmentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostJourneySegmentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsInternalServerError creates a PostJourneySegmentsInternalServerError with default headers values
func NewPostJourneySegmentsInternalServerError() *PostJourneySegmentsInternalServerError {
	return &PostJourneySegmentsInternalServerError{}
}

/*PostJourneySegmentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostJourneySegmentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostJourneySegmentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsServiceUnavailable creates a PostJourneySegmentsServiceUnavailable with default headers values
func NewPostJourneySegmentsServiceUnavailable() *PostJourneySegmentsServiceUnavailable {
	return &PostJourneySegmentsServiceUnavailable{}
}

/*PostJourneySegmentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostJourneySegmentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostJourneySegmentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneySegmentsGatewayTimeout creates a PostJourneySegmentsGatewayTimeout with default headers values
func NewPostJourneySegmentsGatewayTimeout() *PostJourneySegmentsGatewayTimeout {
	return &PostJourneySegmentsGatewayTimeout{}
}

/*PostJourneySegmentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostJourneySegmentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostJourneySegmentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/segments][%d] postJourneySegmentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostJourneySegmentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneySegmentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
