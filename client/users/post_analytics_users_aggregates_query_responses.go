// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAnalyticsUsersAggregatesQueryReader is a Reader for the PostAnalyticsUsersAggregatesQuery structure.
type PostAnalyticsUsersAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsUsersAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsUsersAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsUsersAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsUsersAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsUsersAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsUsersAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsUsersAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsUsersAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsUsersAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsUsersAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsUsersAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsUsersAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsUsersAggregatesQueryOK creates a PostAnalyticsUsersAggregatesQueryOK with default headers values
func NewPostAnalyticsUsersAggregatesQueryOK() *PostAnalyticsUsersAggregatesQueryOK {
	return &PostAnalyticsUsersAggregatesQueryOK{}
}

/*PostAnalyticsUsersAggregatesQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsUsersAggregatesQueryOK struct {
	Payload *models.UserAggregateQueryResponse
}

func (o *PostAnalyticsUsersAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryOK) GetPayload() *models.UserAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryBadRequest creates a PostAnalyticsUsersAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsUsersAggregatesQueryBadRequest() *PostAnalyticsUsersAggregatesQueryBadRequest {
	return &PostAnalyticsUsersAggregatesQueryBadRequest{}
}

/*PostAnalyticsUsersAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsUsersAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryUnauthorized creates a PostAnalyticsUsersAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsUsersAggregatesQueryUnauthorized() *PostAnalyticsUsersAggregatesQueryUnauthorized {
	return &PostAnalyticsUsersAggregatesQueryUnauthorized{}
}

/*PostAnalyticsUsersAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsUsersAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryForbidden creates a PostAnalyticsUsersAggregatesQueryForbidden with default headers values
func NewPostAnalyticsUsersAggregatesQueryForbidden() *PostAnalyticsUsersAggregatesQueryForbidden {
	return &PostAnalyticsUsersAggregatesQueryForbidden{}
}

/*PostAnalyticsUsersAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsUsersAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryNotFound creates a PostAnalyticsUsersAggregatesQueryNotFound with default headers values
func NewPostAnalyticsUsersAggregatesQueryNotFound() *PostAnalyticsUsersAggregatesQueryNotFound {
	return &PostAnalyticsUsersAggregatesQueryNotFound{}
}

/*PostAnalyticsUsersAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsUsersAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsUsersAggregatesQueryRequestEntityTooLarge() *PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge{}
}

/*PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryUnsupportedMediaType creates a PostAnalyticsUsersAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsUsersAggregatesQueryUnsupportedMediaType() *PostAnalyticsUsersAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsUsersAggregatesQueryUnsupportedMediaType{}
}

/*PostAnalyticsUsersAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsUsersAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryTooManyRequests creates a PostAnalyticsUsersAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsUsersAggregatesQueryTooManyRequests() *PostAnalyticsUsersAggregatesQueryTooManyRequests {
	return &PostAnalyticsUsersAggregatesQueryTooManyRequests{}
}

/*PostAnalyticsUsersAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsUsersAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryInternalServerError creates a PostAnalyticsUsersAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsUsersAggregatesQueryInternalServerError() *PostAnalyticsUsersAggregatesQueryInternalServerError {
	return &PostAnalyticsUsersAggregatesQueryInternalServerError{}
}

/*PostAnalyticsUsersAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsUsersAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryServiceUnavailable creates a PostAnalyticsUsersAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsUsersAggregatesQueryServiceUnavailable() *PostAnalyticsUsersAggregatesQueryServiceUnavailable {
	return &PostAnalyticsUsersAggregatesQueryServiceUnavailable{}
}

/*PostAnalyticsUsersAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsUsersAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersAggregatesQueryGatewayTimeout creates a PostAnalyticsUsersAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsUsersAggregatesQueryGatewayTimeout() *PostAnalyticsUsersAggregatesQueryGatewayTimeout {
	return &PostAnalyticsUsersAggregatesQueryGatewayTimeout{}
}

/*PostAnalyticsUsersAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsUsersAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/aggregates/query][%d] postAnalyticsUsersAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsUsersAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
