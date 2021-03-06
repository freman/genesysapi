// Code generated by go-swagger; DO NOT EDIT.

package flows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAnalyticsFlowsAggregatesQueryReader is a Reader for the PostAnalyticsFlowsAggregatesQuery structure.
type PostAnalyticsFlowsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsFlowsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsFlowsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsFlowsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsFlowsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsFlowsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsFlowsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsFlowsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsFlowsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsFlowsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsFlowsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsFlowsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsFlowsAggregatesQueryOK creates a PostAnalyticsFlowsAggregatesQueryOK with default headers values
func NewPostAnalyticsFlowsAggregatesQueryOK() *PostAnalyticsFlowsAggregatesQueryOK {
	return &PostAnalyticsFlowsAggregatesQueryOK{}
}

/*PostAnalyticsFlowsAggregatesQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsFlowsAggregatesQueryOK struct {
	Payload *models.FlowAggregateQueryResponse
}

func (o *PostAnalyticsFlowsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryOK) GetPayload() *models.FlowAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryBadRequest creates a PostAnalyticsFlowsAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsFlowsAggregatesQueryBadRequest() *PostAnalyticsFlowsAggregatesQueryBadRequest {
	return &PostAnalyticsFlowsAggregatesQueryBadRequest{}
}

/*PostAnalyticsFlowsAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsFlowsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryUnauthorized creates a PostAnalyticsFlowsAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsFlowsAggregatesQueryUnauthorized() *PostAnalyticsFlowsAggregatesQueryUnauthorized {
	return &PostAnalyticsFlowsAggregatesQueryUnauthorized{}
}

/*PostAnalyticsFlowsAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsFlowsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryForbidden creates a PostAnalyticsFlowsAggregatesQueryForbidden with default headers values
func NewPostAnalyticsFlowsAggregatesQueryForbidden() *PostAnalyticsFlowsAggregatesQueryForbidden {
	return &PostAnalyticsFlowsAggregatesQueryForbidden{}
}

/*PostAnalyticsFlowsAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsFlowsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryNotFound creates a PostAnalyticsFlowsAggregatesQueryNotFound with default headers values
func NewPostAnalyticsFlowsAggregatesQueryNotFound() *PostAnalyticsFlowsAggregatesQueryNotFound {
	return &PostAnalyticsFlowsAggregatesQueryNotFound{}
}

/*PostAnalyticsFlowsAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsFlowsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge() *PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge{}
}

/*PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryUnsupportedMediaType creates a PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsFlowsAggregatesQueryUnsupportedMediaType() *PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType{}
}

/*PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryTooManyRequests creates a PostAnalyticsFlowsAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsFlowsAggregatesQueryTooManyRequests() *PostAnalyticsFlowsAggregatesQueryTooManyRequests {
	return &PostAnalyticsFlowsAggregatesQueryTooManyRequests{}
}

/*PostAnalyticsFlowsAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsFlowsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryInternalServerError creates a PostAnalyticsFlowsAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsFlowsAggregatesQueryInternalServerError() *PostAnalyticsFlowsAggregatesQueryInternalServerError {
	return &PostAnalyticsFlowsAggregatesQueryInternalServerError{}
}

/*PostAnalyticsFlowsAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsFlowsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryServiceUnavailable creates a PostAnalyticsFlowsAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsFlowsAggregatesQueryServiceUnavailable() *PostAnalyticsFlowsAggregatesQueryServiceUnavailable {
	return &PostAnalyticsFlowsAggregatesQueryServiceUnavailable{}
}

/*PostAnalyticsFlowsAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsFlowsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsFlowsAggregatesQueryGatewayTimeout creates a PostAnalyticsFlowsAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsFlowsAggregatesQueryGatewayTimeout() *PostAnalyticsFlowsAggregatesQueryGatewayTimeout {
	return &PostAnalyticsFlowsAggregatesQueryGatewayTimeout{}
}

/*PostAnalyticsFlowsAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsFlowsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsFlowsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/flows/aggregates/query][%d] postAnalyticsFlowsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsFlowsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsFlowsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
