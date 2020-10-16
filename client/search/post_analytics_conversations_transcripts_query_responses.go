// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAnalyticsConversationsTranscriptsQueryReader is a Reader for the PostAnalyticsConversationsTranscriptsQuery structure.
type PostAnalyticsConversationsTranscriptsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsConversationsTranscriptsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsConversationsTranscriptsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsConversationsTranscriptsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsConversationsTranscriptsQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsConversationsTranscriptsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsConversationsTranscriptsQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsConversationsTranscriptsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsConversationsTranscriptsQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsConversationsTranscriptsQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsConversationsTranscriptsQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsConversationsTranscriptsQueryOK creates a PostAnalyticsConversationsTranscriptsQueryOK with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryOK() *PostAnalyticsConversationsTranscriptsQueryOK {
	return &PostAnalyticsConversationsTranscriptsQueryOK{}
}

/*PostAnalyticsConversationsTranscriptsQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsConversationsTranscriptsQueryOK struct {
	Payload *models.AnalyticsConversationWithoutAttributesMultiGetResponse
}

func (o *PostAnalyticsConversationsTranscriptsQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryOK) GetPayload() *models.AnalyticsConversationWithoutAttributesMultiGetResponse {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsConversationWithoutAttributesMultiGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryBadRequest creates a PostAnalyticsConversationsTranscriptsQueryBadRequest with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryBadRequest() *PostAnalyticsConversationsTranscriptsQueryBadRequest {
	return &PostAnalyticsConversationsTranscriptsQueryBadRequest{}
}

/*PostAnalyticsConversationsTranscriptsQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsConversationsTranscriptsQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryUnauthorized creates a PostAnalyticsConversationsTranscriptsQueryUnauthorized with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryUnauthorized() *PostAnalyticsConversationsTranscriptsQueryUnauthorized {
	return &PostAnalyticsConversationsTranscriptsQueryUnauthorized{}
}

/*PostAnalyticsConversationsTranscriptsQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsConversationsTranscriptsQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryForbidden creates a PostAnalyticsConversationsTranscriptsQueryForbidden with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryForbidden() *PostAnalyticsConversationsTranscriptsQueryForbidden {
	return &PostAnalyticsConversationsTranscriptsQueryForbidden{}
}

/*PostAnalyticsConversationsTranscriptsQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsConversationsTranscriptsQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryNotFound creates a PostAnalyticsConversationsTranscriptsQueryNotFound with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryNotFound() *PostAnalyticsConversationsTranscriptsQueryNotFound {
	return &PostAnalyticsConversationsTranscriptsQueryNotFound{}
}

/*PostAnalyticsConversationsTranscriptsQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsConversationsTranscriptsQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge creates a PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge() *PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge {
	return &PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge{}
}

/*PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType creates a PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType() *PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType {
	return &PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType{}
}

/*PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryTooManyRequests creates a PostAnalyticsConversationsTranscriptsQueryTooManyRequests with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryTooManyRequests() *PostAnalyticsConversationsTranscriptsQueryTooManyRequests {
	return &PostAnalyticsConversationsTranscriptsQueryTooManyRequests{}
}

/*PostAnalyticsConversationsTranscriptsQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsConversationsTranscriptsQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryInternalServerError creates a PostAnalyticsConversationsTranscriptsQueryInternalServerError with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryInternalServerError() *PostAnalyticsConversationsTranscriptsQueryInternalServerError {
	return &PostAnalyticsConversationsTranscriptsQueryInternalServerError{}
}

/*PostAnalyticsConversationsTranscriptsQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsConversationsTranscriptsQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryServiceUnavailable creates a PostAnalyticsConversationsTranscriptsQueryServiceUnavailable with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryServiceUnavailable() *PostAnalyticsConversationsTranscriptsQueryServiceUnavailable {
	return &PostAnalyticsConversationsTranscriptsQueryServiceUnavailable{}
}

/*PostAnalyticsConversationsTranscriptsQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsConversationsTranscriptsQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsTranscriptsQueryGatewayTimeout creates a PostAnalyticsConversationsTranscriptsQueryGatewayTimeout with default headers values
func NewPostAnalyticsConversationsTranscriptsQueryGatewayTimeout() *PostAnalyticsConversationsTranscriptsQueryGatewayTimeout {
	return &PostAnalyticsConversationsTranscriptsQueryGatewayTimeout{}
}

/*PostAnalyticsConversationsTranscriptsQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsConversationsTranscriptsQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsTranscriptsQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/transcripts/query][%d] postAnalyticsConversationsTranscriptsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsConversationsTranscriptsQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsTranscriptsQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}