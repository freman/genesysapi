// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostContentmanagementDocumentsReader is a Reader for the PostContentmanagementDocuments structure.
type PostContentmanagementDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContentmanagementDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContentmanagementDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContentmanagementDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostContentmanagementDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostContentmanagementDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostContentmanagementDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostContentmanagementDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostContentmanagementDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 423:
		result := NewPostContentmanagementDocumentsLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostContentmanagementDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostContentmanagementDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostContentmanagementDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostContentmanagementDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostContentmanagementDocumentsOK creates a PostContentmanagementDocumentsOK with default headers values
func NewPostContentmanagementDocumentsOK() *PostContentmanagementDocumentsOK {
	return &PostContentmanagementDocumentsOK{}
}

/*PostContentmanagementDocumentsOK handles this case with default header values.

successful operation
*/
type PostContentmanagementDocumentsOK struct {
	Payload *models.Document
}

func (o *PostContentmanagementDocumentsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsOK  %+v", 200, o.Payload)
}

func (o *PostContentmanagementDocumentsOK) GetPayload() *models.Document {
	return o.Payload
}

func (o *PostContentmanagementDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsBadRequest creates a PostContentmanagementDocumentsBadRequest with default headers values
func NewPostContentmanagementDocumentsBadRequest() *PostContentmanagementDocumentsBadRequest {
	return &PostContentmanagementDocumentsBadRequest{}
}

/*PostContentmanagementDocumentsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostContentmanagementDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *PostContentmanagementDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsUnauthorized creates a PostContentmanagementDocumentsUnauthorized with default headers values
func NewPostContentmanagementDocumentsUnauthorized() *PostContentmanagementDocumentsUnauthorized {
	return &PostContentmanagementDocumentsUnauthorized{}
}

/*PostContentmanagementDocumentsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostContentmanagementDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostContentmanagementDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsForbidden creates a PostContentmanagementDocumentsForbidden with default headers values
func NewPostContentmanagementDocumentsForbidden() *PostContentmanagementDocumentsForbidden {
	return &PostContentmanagementDocumentsForbidden{}
}

/*PostContentmanagementDocumentsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostContentmanagementDocumentsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *PostContentmanagementDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsNotFound creates a PostContentmanagementDocumentsNotFound with default headers values
func NewPostContentmanagementDocumentsNotFound() *PostContentmanagementDocumentsNotFound {
	return &PostContentmanagementDocumentsNotFound{}
}

/*PostContentmanagementDocumentsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostContentmanagementDocumentsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *PostContentmanagementDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsRequestEntityTooLarge creates a PostContentmanagementDocumentsRequestEntityTooLarge with default headers values
func NewPostContentmanagementDocumentsRequestEntityTooLarge() *PostContentmanagementDocumentsRequestEntityTooLarge {
	return &PostContentmanagementDocumentsRequestEntityTooLarge{}
}

/*PostContentmanagementDocumentsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostContentmanagementDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostContentmanagementDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsUnsupportedMediaType creates a PostContentmanagementDocumentsUnsupportedMediaType with default headers values
func NewPostContentmanagementDocumentsUnsupportedMediaType() *PostContentmanagementDocumentsUnsupportedMediaType {
	return &PostContentmanagementDocumentsUnsupportedMediaType{}
}

/*PostContentmanagementDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostContentmanagementDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostContentmanagementDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsLocked creates a PostContentmanagementDocumentsLocked with default headers values
func NewPostContentmanagementDocumentsLocked() *PostContentmanagementDocumentsLocked {
	return &PostContentmanagementDocumentsLocked{}
}

/*PostContentmanagementDocumentsLocked handles this case with default header values.

Locked - The source document is locked by another operation
*/
type PostContentmanagementDocumentsLocked struct {
}

func (o *PostContentmanagementDocumentsLocked) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsLocked ", 423)
}

func (o *PostContentmanagementDocumentsLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostContentmanagementDocumentsTooManyRequests creates a PostContentmanagementDocumentsTooManyRequests with default headers values
func NewPostContentmanagementDocumentsTooManyRequests() *PostContentmanagementDocumentsTooManyRequests {
	return &PostContentmanagementDocumentsTooManyRequests{}
}

/*PostContentmanagementDocumentsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostContentmanagementDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostContentmanagementDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsInternalServerError creates a PostContentmanagementDocumentsInternalServerError with default headers values
func NewPostContentmanagementDocumentsInternalServerError() *PostContentmanagementDocumentsInternalServerError {
	return &PostContentmanagementDocumentsInternalServerError{}
}

/*PostContentmanagementDocumentsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostContentmanagementDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostContentmanagementDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsServiceUnavailable creates a PostContentmanagementDocumentsServiceUnavailable with default headers values
func NewPostContentmanagementDocumentsServiceUnavailable() *PostContentmanagementDocumentsServiceUnavailable {
	return &PostContentmanagementDocumentsServiceUnavailable{}
}

/*PostContentmanagementDocumentsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostContentmanagementDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostContentmanagementDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementDocumentsGatewayTimeout creates a PostContentmanagementDocumentsGatewayTimeout with default headers values
func NewPostContentmanagementDocumentsGatewayTimeout() *PostContentmanagementDocumentsGatewayTimeout {
	return &PostContentmanagementDocumentsGatewayTimeout{}
}

/*PostContentmanagementDocumentsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostContentmanagementDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/documents][%d] postContentmanagementDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostContentmanagementDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
