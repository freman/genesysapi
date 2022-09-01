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

// GetKnowledgeKnowledgebaseLabelsReader is a Reader for the GetKnowledgeKnowledgebaseLabels structure.
type GetKnowledgeKnowledgebaseLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseLabelsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseLabelsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseLabelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseLabelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseLabelsOK creates a GetKnowledgeKnowledgebaseLabelsOK with default headers values
func NewGetKnowledgeKnowledgebaseLabelsOK() *GetKnowledgeKnowledgebaseLabelsOK {
	return &GetKnowledgeKnowledgebaseLabelsOK{}
}

/*GetKnowledgeKnowledgebaseLabelsOK handles this case with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseLabelsOK struct {
	Payload *models.LabelListing
}

func (o *GetKnowledgeKnowledgebaseLabelsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsOK) GetPayload() *models.LabelListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsBadRequest creates a GetKnowledgeKnowledgebaseLabelsBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseLabelsBadRequest() *GetKnowledgeKnowledgebaseLabelsBadRequest {
	return &GetKnowledgeKnowledgebaseLabelsBadRequest{}
}

/*GetKnowledgeKnowledgebaseLabelsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseLabelsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsUnauthorized creates a GetKnowledgeKnowledgebaseLabelsUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseLabelsUnauthorized() *GetKnowledgeKnowledgebaseLabelsUnauthorized {
	return &GetKnowledgeKnowledgebaseLabelsUnauthorized{}
}

/*GetKnowledgeKnowledgebaseLabelsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseLabelsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsForbidden creates a GetKnowledgeKnowledgebaseLabelsForbidden with default headers values
func NewGetKnowledgeKnowledgebaseLabelsForbidden() *GetKnowledgeKnowledgebaseLabelsForbidden {
	return &GetKnowledgeKnowledgebaseLabelsForbidden{}
}

/*GetKnowledgeKnowledgebaseLabelsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseLabelsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsNotFound creates a GetKnowledgeKnowledgebaseLabelsNotFound with default headers values
func NewGetKnowledgeKnowledgebaseLabelsNotFound() *GetKnowledgeKnowledgebaseLabelsNotFound {
	return &GetKnowledgeKnowledgebaseLabelsNotFound{}
}

/*GetKnowledgeKnowledgebaseLabelsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseLabelsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsRequestTimeout creates a GetKnowledgeKnowledgebaseLabelsRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseLabelsRequestTimeout() *GetKnowledgeKnowledgebaseLabelsRequestTimeout {
	return &GetKnowledgeKnowledgebaseLabelsRequestTimeout{}
}

/*GetKnowledgeKnowledgebaseLabelsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseLabelsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge() *GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge{}
}

/*GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsUnsupportedMediaType creates a GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseLabelsUnsupportedMediaType() *GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType{}
}

/*GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsTooManyRequests creates a GetKnowledgeKnowledgebaseLabelsTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseLabelsTooManyRequests() *GetKnowledgeKnowledgebaseLabelsTooManyRequests {
	return &GetKnowledgeKnowledgebaseLabelsTooManyRequests{}
}

/*GetKnowledgeKnowledgebaseLabelsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseLabelsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsInternalServerError creates a GetKnowledgeKnowledgebaseLabelsInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseLabelsInternalServerError() *GetKnowledgeKnowledgebaseLabelsInternalServerError {
	return &GetKnowledgeKnowledgebaseLabelsInternalServerError{}
}

/*GetKnowledgeKnowledgebaseLabelsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseLabelsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsServiceUnavailable creates a GetKnowledgeKnowledgebaseLabelsServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseLabelsServiceUnavailable() *GetKnowledgeKnowledgebaseLabelsServiceUnavailable {
	return &GetKnowledgeKnowledgebaseLabelsServiceUnavailable{}
}

/*GetKnowledgeKnowledgebaseLabelsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseLabelsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLabelsGatewayTimeout creates a GetKnowledgeKnowledgebaseLabelsGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseLabelsGatewayTimeout() *GetKnowledgeKnowledgebaseLabelsGatewayTimeout {
	return &GetKnowledgeKnowledgebaseLabelsGatewayTimeout{}
}

/*GetKnowledgeKnowledgebaseLabelsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseLabelsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetKnowledgeKnowledgebaseLabelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] getKnowledgeKnowledgebaseLabelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLabelsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLabelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
