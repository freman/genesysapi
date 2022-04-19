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

// PostRecordingsScreensessionsMetadataReader is a Reader for the PostRecordingsScreensessionsMetadata structure.
type PostRecordingsScreensessionsMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecordingsScreensessionsMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostRecordingsScreensessionsMetadataNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRecordingsScreensessionsMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRecordingsScreensessionsMetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRecordingsScreensessionsMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRecordingsScreensessionsMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRecordingsScreensessionsMetadataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRecordingsScreensessionsMetadataRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRecordingsScreensessionsMetadataUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRecordingsScreensessionsMetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRecordingsScreensessionsMetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRecordingsScreensessionsMetadataServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRecordingsScreensessionsMetadataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRecordingsScreensessionsMetadataNoContent creates a PostRecordingsScreensessionsMetadataNoContent with default headers values
func NewPostRecordingsScreensessionsMetadataNoContent() *PostRecordingsScreensessionsMetadataNoContent {
	return &PostRecordingsScreensessionsMetadataNoContent{}
}

/*PostRecordingsScreensessionsMetadataNoContent handles this case with default header values.

Meta-data supplied to screen recording.
*/
type PostRecordingsScreensessionsMetadataNoContent struct {
}

func (o *PostRecordingsScreensessionsMetadataNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataNoContent ", 204)
}

func (o *PostRecordingsScreensessionsMetadataNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRecordingsScreensessionsMetadataBadRequest creates a PostRecordingsScreensessionsMetadataBadRequest with default headers values
func NewPostRecordingsScreensessionsMetadataBadRequest() *PostRecordingsScreensessionsMetadataBadRequest {
	return &PostRecordingsScreensessionsMetadataBadRequest{}
}

/*PostRecordingsScreensessionsMetadataBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRecordingsScreensessionsMetadataBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataBadRequest  %+v", 400, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataUnauthorized creates a PostRecordingsScreensessionsMetadataUnauthorized with default headers values
func NewPostRecordingsScreensessionsMetadataUnauthorized() *PostRecordingsScreensessionsMetadataUnauthorized {
	return &PostRecordingsScreensessionsMetadataUnauthorized{}
}

/*PostRecordingsScreensessionsMetadataUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRecordingsScreensessionsMetadataUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataForbidden creates a PostRecordingsScreensessionsMetadataForbidden with default headers values
func NewPostRecordingsScreensessionsMetadataForbidden() *PostRecordingsScreensessionsMetadataForbidden {
	return &PostRecordingsScreensessionsMetadataForbidden{}
}

/*PostRecordingsScreensessionsMetadataForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostRecordingsScreensessionsMetadataForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataForbidden  %+v", 403, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataNotFound creates a PostRecordingsScreensessionsMetadataNotFound with default headers values
func NewPostRecordingsScreensessionsMetadataNotFound() *PostRecordingsScreensessionsMetadataNotFound {
	return &PostRecordingsScreensessionsMetadataNotFound{}
}

/*PostRecordingsScreensessionsMetadataNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostRecordingsScreensessionsMetadataNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataNotFound  %+v", 404, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataRequestTimeout creates a PostRecordingsScreensessionsMetadataRequestTimeout with default headers values
func NewPostRecordingsScreensessionsMetadataRequestTimeout() *PostRecordingsScreensessionsMetadataRequestTimeout {
	return &PostRecordingsScreensessionsMetadataRequestTimeout{}
}

/*PostRecordingsScreensessionsMetadataRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRecordingsScreensessionsMetadataRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataRequestEntityTooLarge creates a PostRecordingsScreensessionsMetadataRequestEntityTooLarge with default headers values
func NewPostRecordingsScreensessionsMetadataRequestEntityTooLarge() *PostRecordingsScreensessionsMetadataRequestEntityTooLarge {
	return &PostRecordingsScreensessionsMetadataRequestEntityTooLarge{}
}

/*PostRecordingsScreensessionsMetadataRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRecordingsScreensessionsMetadataRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataUnsupportedMediaType creates a PostRecordingsScreensessionsMetadataUnsupportedMediaType with default headers values
func NewPostRecordingsScreensessionsMetadataUnsupportedMediaType() *PostRecordingsScreensessionsMetadataUnsupportedMediaType {
	return &PostRecordingsScreensessionsMetadataUnsupportedMediaType{}
}

/*PostRecordingsScreensessionsMetadataUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRecordingsScreensessionsMetadataUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataTooManyRequests creates a PostRecordingsScreensessionsMetadataTooManyRequests with default headers values
func NewPostRecordingsScreensessionsMetadataTooManyRequests() *PostRecordingsScreensessionsMetadataTooManyRequests {
	return &PostRecordingsScreensessionsMetadataTooManyRequests{}
}

/*PostRecordingsScreensessionsMetadataTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRecordingsScreensessionsMetadataTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataInternalServerError creates a PostRecordingsScreensessionsMetadataInternalServerError with default headers values
func NewPostRecordingsScreensessionsMetadataInternalServerError() *PostRecordingsScreensessionsMetadataInternalServerError {
	return &PostRecordingsScreensessionsMetadataInternalServerError{}
}

/*PostRecordingsScreensessionsMetadataInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRecordingsScreensessionsMetadataInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataServiceUnavailable creates a PostRecordingsScreensessionsMetadataServiceUnavailable with default headers values
func NewPostRecordingsScreensessionsMetadataServiceUnavailable() *PostRecordingsScreensessionsMetadataServiceUnavailable {
	return &PostRecordingsScreensessionsMetadataServiceUnavailable{}
}

/*PostRecordingsScreensessionsMetadataServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRecordingsScreensessionsMetadataServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRecordingsScreensessionsMetadataGatewayTimeout creates a PostRecordingsScreensessionsMetadataGatewayTimeout with default headers values
func NewPostRecordingsScreensessionsMetadataGatewayTimeout() *PostRecordingsScreensessionsMetadataGatewayTimeout {
	return &PostRecordingsScreensessionsMetadataGatewayTimeout{}
}

/*PostRecordingsScreensessionsMetadataGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostRecordingsScreensessionsMetadataGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRecordingsScreensessionsMetadataGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/recordings/screensessions/metadata][%d] postRecordingsScreensessionsMetadataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRecordingsScreensessionsMetadataGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRecordingsScreensessionsMetadataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
