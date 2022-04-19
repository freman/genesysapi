// Code generated by go-swagger; DO NOT EDIT.

package uploads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostLanguageunderstandingMinerUploadsReader is a Reader for the PostLanguageunderstandingMinerUploads structure.
type PostLanguageunderstandingMinerUploadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLanguageunderstandingMinerUploadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLanguageunderstandingMinerUploadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLanguageunderstandingMinerUploadsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLanguageunderstandingMinerUploadsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLanguageunderstandingMinerUploadsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLanguageunderstandingMinerUploadsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLanguageunderstandingMinerUploadsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLanguageunderstandingMinerUploadsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLanguageunderstandingMinerUploadsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLanguageunderstandingMinerUploadsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLanguageunderstandingMinerUploadsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLanguageunderstandingMinerUploadsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLanguageunderstandingMinerUploadsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLanguageunderstandingMinerUploadsOK creates a PostLanguageunderstandingMinerUploadsOK with default headers values
func NewPostLanguageunderstandingMinerUploadsOK() *PostLanguageunderstandingMinerUploadsOK {
	return &PostLanguageunderstandingMinerUploadsOK{}
}

/*PostLanguageunderstandingMinerUploadsOK handles this case with default header values.

Presigned URL successfully created.
*/
type PostLanguageunderstandingMinerUploadsOK struct {
	Payload *models.UploadURLResponse
}

func (o *PostLanguageunderstandingMinerUploadsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsOK) GetPayload() *models.UploadURLResponse {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsBadRequest creates a PostLanguageunderstandingMinerUploadsBadRequest with default headers values
func NewPostLanguageunderstandingMinerUploadsBadRequest() *PostLanguageunderstandingMinerUploadsBadRequest {
	return &PostLanguageunderstandingMinerUploadsBadRequest{}
}

/*PostLanguageunderstandingMinerUploadsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLanguageunderstandingMinerUploadsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsUnauthorized creates a PostLanguageunderstandingMinerUploadsUnauthorized with default headers values
func NewPostLanguageunderstandingMinerUploadsUnauthorized() *PostLanguageunderstandingMinerUploadsUnauthorized {
	return &PostLanguageunderstandingMinerUploadsUnauthorized{}
}

/*PostLanguageunderstandingMinerUploadsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLanguageunderstandingMinerUploadsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsForbidden creates a PostLanguageunderstandingMinerUploadsForbidden with default headers values
func NewPostLanguageunderstandingMinerUploadsForbidden() *PostLanguageunderstandingMinerUploadsForbidden {
	return &PostLanguageunderstandingMinerUploadsForbidden{}
}

/*PostLanguageunderstandingMinerUploadsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLanguageunderstandingMinerUploadsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsNotFound creates a PostLanguageunderstandingMinerUploadsNotFound with default headers values
func NewPostLanguageunderstandingMinerUploadsNotFound() *PostLanguageunderstandingMinerUploadsNotFound {
	return &PostLanguageunderstandingMinerUploadsNotFound{}
}

/*PostLanguageunderstandingMinerUploadsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLanguageunderstandingMinerUploadsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsRequestTimeout creates a PostLanguageunderstandingMinerUploadsRequestTimeout with default headers values
func NewPostLanguageunderstandingMinerUploadsRequestTimeout() *PostLanguageunderstandingMinerUploadsRequestTimeout {
	return &PostLanguageunderstandingMinerUploadsRequestTimeout{}
}

/*PostLanguageunderstandingMinerUploadsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLanguageunderstandingMinerUploadsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsRequestEntityTooLarge creates a PostLanguageunderstandingMinerUploadsRequestEntityTooLarge with default headers values
func NewPostLanguageunderstandingMinerUploadsRequestEntityTooLarge() *PostLanguageunderstandingMinerUploadsRequestEntityTooLarge {
	return &PostLanguageunderstandingMinerUploadsRequestEntityTooLarge{}
}

/*PostLanguageunderstandingMinerUploadsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostLanguageunderstandingMinerUploadsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsUnsupportedMediaType creates a PostLanguageunderstandingMinerUploadsUnsupportedMediaType with default headers values
func NewPostLanguageunderstandingMinerUploadsUnsupportedMediaType() *PostLanguageunderstandingMinerUploadsUnsupportedMediaType {
	return &PostLanguageunderstandingMinerUploadsUnsupportedMediaType{}
}

/*PostLanguageunderstandingMinerUploadsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLanguageunderstandingMinerUploadsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsTooManyRequests creates a PostLanguageunderstandingMinerUploadsTooManyRequests with default headers values
func NewPostLanguageunderstandingMinerUploadsTooManyRequests() *PostLanguageunderstandingMinerUploadsTooManyRequests {
	return &PostLanguageunderstandingMinerUploadsTooManyRequests{}
}

/*PostLanguageunderstandingMinerUploadsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLanguageunderstandingMinerUploadsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsInternalServerError creates a PostLanguageunderstandingMinerUploadsInternalServerError with default headers values
func NewPostLanguageunderstandingMinerUploadsInternalServerError() *PostLanguageunderstandingMinerUploadsInternalServerError {
	return &PostLanguageunderstandingMinerUploadsInternalServerError{}
}

/*PostLanguageunderstandingMinerUploadsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLanguageunderstandingMinerUploadsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsServiceUnavailable creates a PostLanguageunderstandingMinerUploadsServiceUnavailable with default headers values
func NewPostLanguageunderstandingMinerUploadsServiceUnavailable() *PostLanguageunderstandingMinerUploadsServiceUnavailable {
	return &PostLanguageunderstandingMinerUploadsServiceUnavailable{}
}

/*PostLanguageunderstandingMinerUploadsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLanguageunderstandingMinerUploadsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingMinerUploadsGatewayTimeout creates a PostLanguageunderstandingMinerUploadsGatewayTimeout with default headers values
func NewPostLanguageunderstandingMinerUploadsGatewayTimeout() *PostLanguageunderstandingMinerUploadsGatewayTimeout {
	return &PostLanguageunderstandingMinerUploadsGatewayTimeout{}
}

/*PostLanguageunderstandingMinerUploadsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLanguageunderstandingMinerUploadsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingMinerUploadsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/miners/{minerId}/uploads][%d] postLanguageunderstandingMinerUploadsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingMinerUploadsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingMinerUploadsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
