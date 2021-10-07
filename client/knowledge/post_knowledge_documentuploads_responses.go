// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostKnowledgeDocumentuploadsReader is a Reader for the PostKnowledgeDocumentuploads structure.
type PostKnowledgeDocumentuploadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKnowledgeDocumentuploadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKnowledgeDocumentuploadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKnowledgeDocumentuploadsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKnowledgeDocumentuploadsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKnowledgeDocumentuploadsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKnowledgeDocumentuploadsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostKnowledgeDocumentuploadsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostKnowledgeDocumentuploadsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostKnowledgeDocumentuploadsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKnowledgeDocumentuploadsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKnowledgeDocumentuploadsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKnowledgeDocumentuploadsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostKnowledgeDocumentuploadsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKnowledgeDocumentuploadsOK creates a PostKnowledgeDocumentuploadsOK with default headers values
func NewPostKnowledgeDocumentuploadsOK() *PostKnowledgeDocumentuploadsOK {
	return &PostKnowledgeDocumentuploadsOK{}
}

/*PostKnowledgeDocumentuploadsOK handles this case with default header values.

Presigned URL successfully created.
*/
type PostKnowledgeDocumentuploadsOK struct {
	Payload *models.UploadURLResponse
}

func (o *PostKnowledgeDocumentuploadsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsOK) GetPayload() *models.UploadURLResponse {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UploadURLResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsBadRequest creates a PostKnowledgeDocumentuploadsBadRequest with default headers values
func NewPostKnowledgeDocumentuploadsBadRequest() *PostKnowledgeDocumentuploadsBadRequest {
	return &PostKnowledgeDocumentuploadsBadRequest{}
}

/*PostKnowledgeDocumentuploadsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostKnowledgeDocumentuploadsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsUnauthorized creates a PostKnowledgeDocumentuploadsUnauthorized with default headers values
func NewPostKnowledgeDocumentuploadsUnauthorized() *PostKnowledgeDocumentuploadsUnauthorized {
	return &PostKnowledgeDocumentuploadsUnauthorized{}
}

/*PostKnowledgeDocumentuploadsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostKnowledgeDocumentuploadsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsForbidden creates a PostKnowledgeDocumentuploadsForbidden with default headers values
func NewPostKnowledgeDocumentuploadsForbidden() *PostKnowledgeDocumentuploadsForbidden {
	return &PostKnowledgeDocumentuploadsForbidden{}
}

/*PostKnowledgeDocumentuploadsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostKnowledgeDocumentuploadsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsNotFound creates a PostKnowledgeDocumentuploadsNotFound with default headers values
func NewPostKnowledgeDocumentuploadsNotFound() *PostKnowledgeDocumentuploadsNotFound {
	return &PostKnowledgeDocumentuploadsNotFound{}
}

/*PostKnowledgeDocumentuploadsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostKnowledgeDocumentuploadsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsRequestTimeout creates a PostKnowledgeDocumentuploadsRequestTimeout with default headers values
func NewPostKnowledgeDocumentuploadsRequestTimeout() *PostKnowledgeDocumentuploadsRequestTimeout {
	return &PostKnowledgeDocumentuploadsRequestTimeout{}
}

/*PostKnowledgeDocumentuploadsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostKnowledgeDocumentuploadsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsRequestEntityTooLarge creates a PostKnowledgeDocumentuploadsRequestEntityTooLarge with default headers values
func NewPostKnowledgeDocumentuploadsRequestEntityTooLarge() *PostKnowledgeDocumentuploadsRequestEntityTooLarge {
	return &PostKnowledgeDocumentuploadsRequestEntityTooLarge{}
}

/*PostKnowledgeDocumentuploadsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostKnowledgeDocumentuploadsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsUnsupportedMediaType creates a PostKnowledgeDocumentuploadsUnsupportedMediaType with default headers values
func NewPostKnowledgeDocumentuploadsUnsupportedMediaType() *PostKnowledgeDocumentuploadsUnsupportedMediaType {
	return &PostKnowledgeDocumentuploadsUnsupportedMediaType{}
}

/*PostKnowledgeDocumentuploadsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostKnowledgeDocumentuploadsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsTooManyRequests creates a PostKnowledgeDocumentuploadsTooManyRequests with default headers values
func NewPostKnowledgeDocumentuploadsTooManyRequests() *PostKnowledgeDocumentuploadsTooManyRequests {
	return &PostKnowledgeDocumentuploadsTooManyRequests{}
}

/*PostKnowledgeDocumentuploadsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostKnowledgeDocumentuploadsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsInternalServerError creates a PostKnowledgeDocumentuploadsInternalServerError with default headers values
func NewPostKnowledgeDocumentuploadsInternalServerError() *PostKnowledgeDocumentuploadsInternalServerError {
	return &PostKnowledgeDocumentuploadsInternalServerError{}
}

/*PostKnowledgeDocumentuploadsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostKnowledgeDocumentuploadsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsServiceUnavailable creates a PostKnowledgeDocumentuploadsServiceUnavailable with default headers values
func NewPostKnowledgeDocumentuploadsServiceUnavailable() *PostKnowledgeDocumentuploadsServiceUnavailable {
	return &PostKnowledgeDocumentuploadsServiceUnavailable{}
}

/*PostKnowledgeDocumentuploadsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostKnowledgeDocumentuploadsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeDocumentuploadsGatewayTimeout creates a PostKnowledgeDocumentuploadsGatewayTimeout with default headers values
func NewPostKnowledgeDocumentuploadsGatewayTimeout() *PostKnowledgeDocumentuploadsGatewayTimeout {
	return &PostKnowledgeDocumentuploadsGatewayTimeout{}
}

/*PostKnowledgeDocumentuploadsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostKnowledgeDocumentuploadsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeDocumentuploadsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/documentuploads][%d] postKnowledgeDocumentuploadsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeDocumentuploadsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeDocumentuploadsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
