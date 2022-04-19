// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAnalyticsConversationsAggregatesQueryReader is a Reader for the PostAnalyticsConversationsAggregatesQuery structure.
type PostAnalyticsConversationsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsConversationsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsConversationsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsConversationsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsConversationsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsConversationsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsConversationsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAnalyticsConversationsAggregatesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsConversationsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsConversationsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsConversationsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsConversationsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsConversationsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsConversationsAggregatesQueryOK creates a PostAnalyticsConversationsAggregatesQueryOK with default headers values
func NewPostAnalyticsConversationsAggregatesQueryOK() *PostAnalyticsConversationsAggregatesQueryOK {
	return &PostAnalyticsConversationsAggregatesQueryOK{}
}

/*PostAnalyticsConversationsAggregatesQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsConversationsAggregatesQueryOK struct {
	Payload *models.ConversationAggregateQueryResponse
}

func (o *PostAnalyticsConversationsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryOK) GetPayload() *models.ConversationAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConversationAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryBadRequest creates a PostAnalyticsConversationsAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsConversationsAggregatesQueryBadRequest() *PostAnalyticsConversationsAggregatesQueryBadRequest {
	return &PostAnalyticsConversationsAggregatesQueryBadRequest{}
}

/*PostAnalyticsConversationsAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsConversationsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryUnauthorized creates a PostAnalyticsConversationsAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsConversationsAggregatesQueryUnauthorized() *PostAnalyticsConversationsAggregatesQueryUnauthorized {
	return &PostAnalyticsConversationsAggregatesQueryUnauthorized{}
}

/*PostAnalyticsConversationsAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsConversationsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryForbidden creates a PostAnalyticsConversationsAggregatesQueryForbidden with default headers values
func NewPostAnalyticsConversationsAggregatesQueryForbidden() *PostAnalyticsConversationsAggregatesQueryForbidden {
	return &PostAnalyticsConversationsAggregatesQueryForbidden{}
}

/*PostAnalyticsConversationsAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsConversationsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryNotFound creates a PostAnalyticsConversationsAggregatesQueryNotFound with default headers values
func NewPostAnalyticsConversationsAggregatesQueryNotFound() *PostAnalyticsConversationsAggregatesQueryNotFound {
	return &PostAnalyticsConversationsAggregatesQueryNotFound{}
}

/*PostAnalyticsConversationsAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsConversationsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryRequestTimeout creates a PostAnalyticsConversationsAggregatesQueryRequestTimeout with default headers values
func NewPostAnalyticsConversationsAggregatesQueryRequestTimeout() *PostAnalyticsConversationsAggregatesQueryRequestTimeout {
	return &PostAnalyticsConversationsAggregatesQueryRequestTimeout{}
}

/*PostAnalyticsConversationsAggregatesQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAnalyticsConversationsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge() *PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge{}
}

/*PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryUnsupportedMediaType creates a PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsConversationsAggregatesQueryUnsupportedMediaType() *PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType{}
}

/*PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryTooManyRequests creates a PostAnalyticsConversationsAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsConversationsAggregatesQueryTooManyRequests() *PostAnalyticsConversationsAggregatesQueryTooManyRequests {
	return &PostAnalyticsConversationsAggregatesQueryTooManyRequests{}
}

/*PostAnalyticsConversationsAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAnalyticsConversationsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryInternalServerError creates a PostAnalyticsConversationsAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsConversationsAggregatesQueryInternalServerError() *PostAnalyticsConversationsAggregatesQueryInternalServerError {
	return &PostAnalyticsConversationsAggregatesQueryInternalServerError{}
}

/*PostAnalyticsConversationsAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsConversationsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryServiceUnavailable creates a PostAnalyticsConversationsAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsConversationsAggregatesQueryServiceUnavailable() *PostAnalyticsConversationsAggregatesQueryServiceUnavailable {
	return &PostAnalyticsConversationsAggregatesQueryServiceUnavailable{}
}

/*PostAnalyticsConversationsAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsConversationsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsConversationsAggregatesQueryGatewayTimeout creates a PostAnalyticsConversationsAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsConversationsAggregatesQueryGatewayTimeout() *PostAnalyticsConversationsAggregatesQueryGatewayTimeout {
	return &PostAnalyticsConversationsAggregatesQueryGatewayTimeout{}
}

/*PostAnalyticsConversationsAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsConversationsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsConversationsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/conversations/aggregates/query][%d] postAnalyticsConversationsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsConversationsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsConversationsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
