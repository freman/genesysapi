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

// PostKnowledgeKnowledgebaseLabelsReader is a Reader for the PostKnowledgeKnowledgebaseLabels structure.
type PostKnowledgeKnowledgebaseLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKnowledgeKnowledgebaseLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKnowledgeKnowledgebaseLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostKnowledgeKnowledgebaseLabelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKnowledgeKnowledgebaseLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKnowledgeKnowledgebaseLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKnowledgeKnowledgebaseLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKnowledgeKnowledgebaseLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostKnowledgeKnowledgebaseLabelsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostKnowledgeKnowledgebaseLabelsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostKnowledgeKnowledgebaseLabelsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKnowledgeKnowledgebaseLabelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKnowledgeKnowledgebaseLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKnowledgeKnowledgebaseLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostKnowledgeKnowledgebaseLabelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKnowledgeKnowledgebaseLabelsOK creates a PostKnowledgeKnowledgebaseLabelsOK with default headers values
func NewPostKnowledgeKnowledgebaseLabelsOK() *PostKnowledgeKnowledgebaseLabelsOK {
	return &PostKnowledgeKnowledgebaseLabelsOK{}
}

/*PostKnowledgeKnowledgebaseLabelsOK handles this case with default header values.

successful operation
*/
type PostKnowledgeKnowledgebaseLabelsOK struct {
	Payload *models.LabelResponse
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) GetPayload() *models.LabelResponse {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsCreated creates a PostKnowledgeKnowledgebaseLabelsCreated with default headers values
func NewPostKnowledgeKnowledgebaseLabelsCreated() *PostKnowledgeKnowledgebaseLabelsCreated {
	return &PostKnowledgeKnowledgebaseLabelsCreated{}
}

/*PostKnowledgeKnowledgebaseLabelsCreated handles this case with default header values.

Label created
*/
type PostKnowledgeKnowledgebaseLabelsCreated struct {
	Payload *models.LabelResponse
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsCreated  %+v", 201, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) GetPayload() *models.LabelResponse {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsBadRequest creates a PostKnowledgeKnowledgebaseLabelsBadRequest with default headers values
func NewPostKnowledgeKnowledgebaseLabelsBadRequest() *PostKnowledgeKnowledgebaseLabelsBadRequest {
	return &PostKnowledgeKnowledgebaseLabelsBadRequest{}
}

/*PostKnowledgeKnowledgebaseLabelsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostKnowledgeKnowledgebaseLabelsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsUnauthorized creates a PostKnowledgeKnowledgebaseLabelsUnauthorized with default headers values
func NewPostKnowledgeKnowledgebaseLabelsUnauthorized() *PostKnowledgeKnowledgebaseLabelsUnauthorized {
	return &PostKnowledgeKnowledgebaseLabelsUnauthorized{}
}

/*PostKnowledgeKnowledgebaseLabelsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostKnowledgeKnowledgebaseLabelsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsForbidden creates a PostKnowledgeKnowledgebaseLabelsForbidden with default headers values
func NewPostKnowledgeKnowledgebaseLabelsForbidden() *PostKnowledgeKnowledgebaseLabelsForbidden {
	return &PostKnowledgeKnowledgebaseLabelsForbidden{}
}

/*PostKnowledgeKnowledgebaseLabelsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostKnowledgeKnowledgebaseLabelsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsNotFound creates a PostKnowledgeKnowledgebaseLabelsNotFound with default headers values
func NewPostKnowledgeKnowledgebaseLabelsNotFound() *PostKnowledgeKnowledgebaseLabelsNotFound {
	return &PostKnowledgeKnowledgebaseLabelsNotFound{}
}

/*PostKnowledgeKnowledgebaseLabelsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostKnowledgeKnowledgebaseLabelsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsRequestTimeout creates a PostKnowledgeKnowledgebaseLabelsRequestTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLabelsRequestTimeout() *PostKnowledgeKnowledgebaseLabelsRequestTimeout {
	return &PostKnowledgeKnowledgebaseLabelsRequestTimeout{}
}

/*PostKnowledgeKnowledgebaseLabelsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostKnowledgeKnowledgebaseLabelsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsConflict creates a PostKnowledgeKnowledgebaseLabelsConflict with default headers values
func NewPostKnowledgeKnowledgebaseLabelsConflict() *PostKnowledgeKnowledgebaseLabelsConflict {
	return &PostKnowledgeKnowledgebaseLabelsConflict{}
}

/*PostKnowledgeKnowledgebaseLabelsConflict handles this case with default header values.

Conflict
*/
type PostKnowledgeKnowledgebaseLabelsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsConflict  %+v", 409, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge creates a PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge with default headers values
func NewPostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge() *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge {
	return &PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge{}
}

/*PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsUnsupportedMediaType creates a PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType with default headers values
func NewPostKnowledgeKnowledgebaseLabelsUnsupportedMediaType() *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType {
	return &PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType{}
}

/*PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsTooManyRequests creates a PostKnowledgeKnowledgebaseLabelsTooManyRequests with default headers values
func NewPostKnowledgeKnowledgebaseLabelsTooManyRequests() *PostKnowledgeKnowledgebaseLabelsTooManyRequests {
	return &PostKnowledgeKnowledgebaseLabelsTooManyRequests{}
}

/*PostKnowledgeKnowledgebaseLabelsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostKnowledgeKnowledgebaseLabelsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsInternalServerError creates a PostKnowledgeKnowledgebaseLabelsInternalServerError with default headers values
func NewPostKnowledgeKnowledgebaseLabelsInternalServerError() *PostKnowledgeKnowledgebaseLabelsInternalServerError {
	return &PostKnowledgeKnowledgebaseLabelsInternalServerError{}
}

/*PostKnowledgeKnowledgebaseLabelsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostKnowledgeKnowledgebaseLabelsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsServiceUnavailable creates a PostKnowledgeKnowledgebaseLabelsServiceUnavailable with default headers values
func NewPostKnowledgeKnowledgebaseLabelsServiceUnavailable() *PostKnowledgeKnowledgebaseLabelsServiceUnavailable {
	return &PostKnowledgeKnowledgebaseLabelsServiceUnavailable{}
}

/*PostKnowledgeKnowledgebaseLabelsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostKnowledgeKnowledgebaseLabelsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsGatewayTimeout creates a PostKnowledgeKnowledgebaseLabelsGatewayTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLabelsGatewayTimeout() *PostKnowledgeKnowledgebaseLabelsGatewayTimeout {
	return &PostKnowledgeKnowledgebaseLabelsGatewayTimeout{}
}

/*PostKnowledgeKnowledgebaseLabelsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostKnowledgeKnowledgebaseLabelsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}