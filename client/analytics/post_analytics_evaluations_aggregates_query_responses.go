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

// PostAnalyticsEvaluationsAggregatesQueryReader is a Reader for the PostAnalyticsEvaluationsAggregatesQuery structure.
type PostAnalyticsEvaluationsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsEvaluationsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsEvaluationsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsEvaluationsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsEvaluationsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsEvaluationsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsEvaluationsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAnalyticsEvaluationsAggregatesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsEvaluationsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsEvaluationsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsEvaluationsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsEvaluationsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsEvaluationsAggregatesQueryOK creates a PostAnalyticsEvaluationsAggregatesQueryOK with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryOK() *PostAnalyticsEvaluationsAggregatesQueryOK {
	return &PostAnalyticsEvaluationsAggregatesQueryOK{}
}

/*PostAnalyticsEvaluationsAggregatesQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsEvaluationsAggregatesQueryOK struct {
	Payload *models.EvaluationAggregateQueryResponse
}

func (o *PostAnalyticsEvaluationsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryOK) GetPayload() *models.EvaluationAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryBadRequest creates a PostAnalyticsEvaluationsAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryBadRequest() *PostAnalyticsEvaluationsAggregatesQueryBadRequest {
	return &PostAnalyticsEvaluationsAggregatesQueryBadRequest{}
}

/*PostAnalyticsEvaluationsAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsEvaluationsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryUnauthorized creates a PostAnalyticsEvaluationsAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryUnauthorized() *PostAnalyticsEvaluationsAggregatesQueryUnauthorized {
	return &PostAnalyticsEvaluationsAggregatesQueryUnauthorized{}
}

/*PostAnalyticsEvaluationsAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsEvaluationsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryForbidden creates a PostAnalyticsEvaluationsAggregatesQueryForbidden with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryForbidden() *PostAnalyticsEvaluationsAggregatesQueryForbidden {
	return &PostAnalyticsEvaluationsAggregatesQueryForbidden{}
}

/*PostAnalyticsEvaluationsAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsEvaluationsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryNotFound creates a PostAnalyticsEvaluationsAggregatesQueryNotFound with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryNotFound() *PostAnalyticsEvaluationsAggregatesQueryNotFound {
	return &PostAnalyticsEvaluationsAggregatesQueryNotFound{}
}

/*PostAnalyticsEvaluationsAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsEvaluationsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryRequestTimeout creates a PostAnalyticsEvaluationsAggregatesQueryRequestTimeout with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryRequestTimeout() *PostAnalyticsEvaluationsAggregatesQueryRequestTimeout {
	return &PostAnalyticsEvaluationsAggregatesQueryRequestTimeout{}
}

/*PostAnalyticsEvaluationsAggregatesQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAnalyticsEvaluationsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge() *PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge{}
}

/*PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType creates a PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType() *PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType{}
}

/*PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryTooManyRequests creates a PostAnalyticsEvaluationsAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryTooManyRequests() *PostAnalyticsEvaluationsAggregatesQueryTooManyRequests {
	return &PostAnalyticsEvaluationsAggregatesQueryTooManyRequests{}
}

/*PostAnalyticsEvaluationsAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAnalyticsEvaluationsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryInternalServerError creates a PostAnalyticsEvaluationsAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryInternalServerError() *PostAnalyticsEvaluationsAggregatesQueryInternalServerError {
	return &PostAnalyticsEvaluationsAggregatesQueryInternalServerError{}
}

/*PostAnalyticsEvaluationsAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsEvaluationsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryServiceUnavailable creates a PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryServiceUnavailable() *PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable {
	return &PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable{}
}

/*PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsEvaluationsAggregatesQueryGatewayTimeout creates a PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsEvaluationsAggregatesQueryGatewayTimeout() *PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout {
	return &PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout{}
}

/*PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/evaluations/aggregates/query][%d] postAnalyticsEvaluationsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsEvaluationsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}