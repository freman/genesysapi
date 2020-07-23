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

// PostKnowledgeKnowledgebasesReader is a Reader for the PostKnowledgeKnowledgebases structure.
type PostKnowledgeKnowledgebasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKnowledgeKnowledgebasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKnowledgeKnowledgebasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKnowledgeKnowledgebasesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKnowledgeKnowledgebasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKnowledgeKnowledgebasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKnowledgeKnowledgebasesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostKnowledgeKnowledgebasesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostKnowledgeKnowledgebasesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKnowledgeKnowledgebasesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKnowledgeKnowledgebasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKnowledgeKnowledgebasesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostKnowledgeKnowledgebasesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKnowledgeKnowledgebasesOK creates a PostKnowledgeKnowledgebasesOK with default headers values
func NewPostKnowledgeKnowledgebasesOK() *PostKnowledgeKnowledgebasesOK {
	return &PostKnowledgeKnowledgebasesOK{}
}

/*PostKnowledgeKnowledgebasesOK handles this case with default header values.

successful operation
*/
type PostKnowledgeKnowledgebasesOK struct {
	Payload *models.KnowledgeBase
}

func (o *PostKnowledgeKnowledgebasesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesOK) GetPayload() *models.KnowledgeBase {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KnowledgeBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesBadRequest creates a PostKnowledgeKnowledgebasesBadRequest with default headers values
func NewPostKnowledgeKnowledgebasesBadRequest() *PostKnowledgeKnowledgebasesBadRequest {
	return &PostKnowledgeKnowledgebasesBadRequest{}
}

/*PostKnowledgeKnowledgebasesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostKnowledgeKnowledgebasesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesUnauthorized creates a PostKnowledgeKnowledgebasesUnauthorized with default headers values
func NewPostKnowledgeKnowledgebasesUnauthorized() *PostKnowledgeKnowledgebasesUnauthorized {
	return &PostKnowledgeKnowledgebasesUnauthorized{}
}

/*PostKnowledgeKnowledgebasesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostKnowledgeKnowledgebasesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesForbidden creates a PostKnowledgeKnowledgebasesForbidden with default headers values
func NewPostKnowledgeKnowledgebasesForbidden() *PostKnowledgeKnowledgebasesForbidden {
	return &PostKnowledgeKnowledgebasesForbidden{}
}

/*PostKnowledgeKnowledgebasesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostKnowledgeKnowledgebasesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesNotFound creates a PostKnowledgeKnowledgebasesNotFound with default headers values
func NewPostKnowledgeKnowledgebasesNotFound() *PostKnowledgeKnowledgebasesNotFound {
	return &PostKnowledgeKnowledgebasesNotFound{}
}

/*PostKnowledgeKnowledgebasesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostKnowledgeKnowledgebasesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesRequestEntityTooLarge creates a PostKnowledgeKnowledgebasesRequestEntityTooLarge with default headers values
func NewPostKnowledgeKnowledgebasesRequestEntityTooLarge() *PostKnowledgeKnowledgebasesRequestEntityTooLarge {
	return &PostKnowledgeKnowledgebasesRequestEntityTooLarge{}
}

/*PostKnowledgeKnowledgebasesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostKnowledgeKnowledgebasesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesUnsupportedMediaType creates a PostKnowledgeKnowledgebasesUnsupportedMediaType with default headers values
func NewPostKnowledgeKnowledgebasesUnsupportedMediaType() *PostKnowledgeKnowledgebasesUnsupportedMediaType {
	return &PostKnowledgeKnowledgebasesUnsupportedMediaType{}
}

/*PostKnowledgeKnowledgebasesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostKnowledgeKnowledgebasesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesTooManyRequests creates a PostKnowledgeKnowledgebasesTooManyRequests with default headers values
func NewPostKnowledgeKnowledgebasesTooManyRequests() *PostKnowledgeKnowledgebasesTooManyRequests {
	return &PostKnowledgeKnowledgebasesTooManyRequests{}
}

/*PostKnowledgeKnowledgebasesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostKnowledgeKnowledgebasesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesInternalServerError creates a PostKnowledgeKnowledgebasesInternalServerError with default headers values
func NewPostKnowledgeKnowledgebasesInternalServerError() *PostKnowledgeKnowledgebasesInternalServerError {
	return &PostKnowledgeKnowledgebasesInternalServerError{}
}

/*PostKnowledgeKnowledgebasesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostKnowledgeKnowledgebasesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesServiceUnavailable creates a PostKnowledgeKnowledgebasesServiceUnavailable with default headers values
func NewPostKnowledgeKnowledgebasesServiceUnavailable() *PostKnowledgeKnowledgebasesServiceUnavailable {
	return &PostKnowledgeKnowledgebasesServiceUnavailable{}
}

/*PostKnowledgeKnowledgebasesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostKnowledgeKnowledgebasesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebasesGatewayTimeout creates a PostKnowledgeKnowledgebasesGatewayTimeout with default headers values
func NewPostKnowledgeKnowledgebasesGatewayTimeout() *PostKnowledgeKnowledgebasesGatewayTimeout {
	return &PostKnowledgeKnowledgebasesGatewayTimeout{}
}

/*PostKnowledgeKnowledgebasesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostKnowledgeKnowledgebasesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebasesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases][%d] postKnowledgeKnowledgebasesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeKnowledgebasesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebasesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
