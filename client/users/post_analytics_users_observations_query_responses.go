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

// PostAnalyticsUsersObservationsQueryReader is a Reader for the PostAnalyticsUsersObservationsQuery structure.
type PostAnalyticsUsersObservationsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsUsersObservationsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsUsersObservationsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsUsersObservationsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsUsersObservationsQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsUsersObservationsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsUsersObservationsQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsUsersObservationsQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsUsersObservationsQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsUsersObservationsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsUsersObservationsQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsUsersObservationsQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsUsersObservationsQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsUsersObservationsQueryOK creates a PostAnalyticsUsersObservationsQueryOK with default headers values
func NewPostAnalyticsUsersObservationsQueryOK() *PostAnalyticsUsersObservationsQueryOK {
	return &PostAnalyticsUsersObservationsQueryOK{}
}

/*PostAnalyticsUsersObservationsQueryOK handles this case with default header values.

successful operation
*/
type PostAnalyticsUsersObservationsQueryOK struct {
	Payload *models.UserObservationQueryResponse
}

func (o *PostAnalyticsUsersObservationsQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryOK) GetPayload() *models.UserObservationQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserObservationQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryBadRequest creates a PostAnalyticsUsersObservationsQueryBadRequest with default headers values
func NewPostAnalyticsUsersObservationsQueryBadRequest() *PostAnalyticsUsersObservationsQueryBadRequest {
	return &PostAnalyticsUsersObservationsQueryBadRequest{}
}

/*PostAnalyticsUsersObservationsQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsUsersObservationsQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryUnauthorized creates a PostAnalyticsUsersObservationsQueryUnauthorized with default headers values
func NewPostAnalyticsUsersObservationsQueryUnauthorized() *PostAnalyticsUsersObservationsQueryUnauthorized {
	return &PostAnalyticsUsersObservationsQueryUnauthorized{}
}

/*PostAnalyticsUsersObservationsQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsUsersObservationsQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryForbidden creates a PostAnalyticsUsersObservationsQueryForbidden with default headers values
func NewPostAnalyticsUsersObservationsQueryForbidden() *PostAnalyticsUsersObservationsQueryForbidden {
	return &PostAnalyticsUsersObservationsQueryForbidden{}
}

/*PostAnalyticsUsersObservationsQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsUsersObservationsQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryNotFound creates a PostAnalyticsUsersObservationsQueryNotFound with default headers values
func NewPostAnalyticsUsersObservationsQueryNotFound() *PostAnalyticsUsersObservationsQueryNotFound {
	return &PostAnalyticsUsersObservationsQueryNotFound{}
}

/*PostAnalyticsUsersObservationsQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAnalyticsUsersObservationsQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryRequestEntityTooLarge creates a PostAnalyticsUsersObservationsQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsUsersObservationsQueryRequestEntityTooLarge() *PostAnalyticsUsersObservationsQueryRequestEntityTooLarge {
	return &PostAnalyticsUsersObservationsQueryRequestEntityTooLarge{}
}

/*PostAnalyticsUsersObservationsQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostAnalyticsUsersObservationsQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryUnsupportedMediaType creates a PostAnalyticsUsersObservationsQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsUsersObservationsQueryUnsupportedMediaType() *PostAnalyticsUsersObservationsQueryUnsupportedMediaType {
	return &PostAnalyticsUsersObservationsQueryUnsupportedMediaType{}
}

/*PostAnalyticsUsersObservationsQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsUsersObservationsQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryTooManyRequests creates a PostAnalyticsUsersObservationsQueryTooManyRequests with default headers values
func NewPostAnalyticsUsersObservationsQueryTooManyRequests() *PostAnalyticsUsersObservationsQueryTooManyRequests {
	return &PostAnalyticsUsersObservationsQueryTooManyRequests{}
}

/*PostAnalyticsUsersObservationsQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostAnalyticsUsersObservationsQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryInternalServerError creates a PostAnalyticsUsersObservationsQueryInternalServerError with default headers values
func NewPostAnalyticsUsersObservationsQueryInternalServerError() *PostAnalyticsUsersObservationsQueryInternalServerError {
	return &PostAnalyticsUsersObservationsQueryInternalServerError{}
}

/*PostAnalyticsUsersObservationsQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsUsersObservationsQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryServiceUnavailable creates a PostAnalyticsUsersObservationsQueryServiceUnavailable with default headers values
func NewPostAnalyticsUsersObservationsQueryServiceUnavailable() *PostAnalyticsUsersObservationsQueryServiceUnavailable {
	return &PostAnalyticsUsersObservationsQueryServiceUnavailable{}
}

/*PostAnalyticsUsersObservationsQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsUsersObservationsQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersObservationsQueryGatewayTimeout creates a PostAnalyticsUsersObservationsQueryGatewayTimeout with default headers values
func NewPostAnalyticsUsersObservationsQueryGatewayTimeout() *PostAnalyticsUsersObservationsQueryGatewayTimeout {
	return &PostAnalyticsUsersObservationsQueryGatewayTimeout{}
}

/*PostAnalyticsUsersObservationsQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAnalyticsUsersObservationsQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAnalyticsUsersObservationsQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/observations/query][%d] postAnalyticsUsersObservationsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsUsersObservationsQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersObservationsQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
