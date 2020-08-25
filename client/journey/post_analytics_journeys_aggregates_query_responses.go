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

// PostAnalyticsJourneysAggregatesQueryReader is a Reader for the PostAnalyticsJourneysAggregatesQuery structure.
type PostAnalyticsJourneysAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsJourneysAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsJourneysAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsJourneysAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsJourneysAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsJourneysAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsJourneysAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsJourneysAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsJourneysAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsJourneysAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsJourneysAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsJourneysAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsJourneysAggregatesQueryOK creates a PostAnalyticsJourneysAggregatesQueryOK with default headers values
func NewPostAnalyticsJourneysAggregatesQueryOK() *PostAnalyticsJourneysAggregatesQueryOK {
	return &PostAnalyticsJourneysAggregatesQueryOK{}
}

/*PostAnalyticsJourneysAggregatesQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsJourneysAggregatesQueryOK struct {
	Payload *models.JourneyAggregateQueryResponse
}

func (o *PostAnalyticsJourneysAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryOK) GetPayload() *models.JourneyAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JourneyAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryBadRequest creates a PostAnalyticsJourneysAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsJourneysAggregatesQueryBadRequest() *PostAnalyticsJourneysAggregatesQueryBadRequest {
	return &PostAnalyticsJourneysAggregatesQueryBadRequest{}
}

/*PostAnalyticsJourneysAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsJourneysAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryUnauthorized creates a PostAnalyticsJourneysAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsJourneysAggregatesQueryUnauthorized() *PostAnalyticsJourneysAggregatesQueryUnauthorized {
	return &PostAnalyticsJourneysAggregatesQueryUnauthorized{}
}

/*PostAnalyticsJourneysAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsJourneysAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryForbidden creates a PostAnalyticsJourneysAggregatesQueryForbidden with default headers values
func NewPostAnalyticsJourneysAggregatesQueryForbidden() *PostAnalyticsJourneysAggregatesQueryForbidden {
	return &PostAnalyticsJourneysAggregatesQueryForbidden{}
}

/*PostAnalyticsJourneysAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsJourneysAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryNotFound creates a PostAnalyticsJourneysAggregatesQueryNotFound with default headers values
func NewPostAnalyticsJourneysAggregatesQueryNotFound() *PostAnalyticsJourneysAggregatesQueryNotFound {
	return &PostAnalyticsJourneysAggregatesQueryNotFound{}
}

/*PostAnalyticsJourneysAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsJourneysAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge() *PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge{}
}

/*PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryUnsupportedMediaType creates a PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsJourneysAggregatesQueryUnsupportedMediaType() *PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType{}
}

/*PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryTooManyRequests creates a PostAnalyticsJourneysAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsJourneysAggregatesQueryTooManyRequests() *PostAnalyticsJourneysAggregatesQueryTooManyRequests {
	return &PostAnalyticsJourneysAggregatesQueryTooManyRequests{}
}

/*PostAnalyticsJourneysAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsJourneysAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryInternalServerError creates a PostAnalyticsJourneysAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsJourneysAggregatesQueryInternalServerError() *PostAnalyticsJourneysAggregatesQueryInternalServerError {
	return &PostAnalyticsJourneysAggregatesQueryInternalServerError{}
}

/*PostAnalyticsJourneysAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsJourneysAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryServiceUnavailable creates a PostAnalyticsJourneysAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsJourneysAggregatesQueryServiceUnavailable() *PostAnalyticsJourneysAggregatesQueryServiceUnavailable {
	return &PostAnalyticsJourneysAggregatesQueryServiceUnavailable{}
}

/*PostAnalyticsJourneysAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsJourneysAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsJourneysAggregatesQueryGatewayTimeout creates a PostAnalyticsJourneysAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsJourneysAggregatesQueryGatewayTimeout() *PostAnalyticsJourneysAggregatesQueryGatewayTimeout {
	return &PostAnalyticsJourneysAggregatesQueryGatewayTimeout{}
}

/*PostAnalyticsJourneysAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsJourneysAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsJourneysAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/journeys/aggregates/query][%d] postAnalyticsJourneysAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsJourneysAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsJourneysAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
