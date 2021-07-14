// Code generated by go-swagger; DO NOT EDIT.

package learning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostLearningAssignmentsAggregatesQueryReader is a Reader for the PostLearningAssignmentsAggregatesQuery structure.
type PostLearningAssignmentsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLearningAssignmentsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLearningAssignmentsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLearningAssignmentsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLearningAssignmentsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLearningAssignmentsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLearningAssignmentsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostLearningAssignmentsAggregatesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLearningAssignmentsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLearningAssignmentsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLearningAssignmentsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLearningAssignmentsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLearningAssignmentsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLearningAssignmentsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLearningAssignmentsAggregatesQueryOK creates a PostLearningAssignmentsAggregatesQueryOK with default headers values
func NewPostLearningAssignmentsAggregatesQueryOK() *PostLearningAssignmentsAggregatesQueryOK {
	return &PostLearningAssignmentsAggregatesQueryOK{}
}

/*PostLearningAssignmentsAggregatesQueryOK handles this case with default header values.

Query completed successfully
*/
type PostLearningAssignmentsAggregatesQueryOK struct {
	Payload *models.LearningAssignmentAggregateResponse
}

func (o *PostLearningAssignmentsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryOK) GetPayload() *models.LearningAssignmentAggregateResponse {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LearningAssignmentAggregateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryBadRequest creates a PostLearningAssignmentsAggregatesQueryBadRequest with default headers values
func NewPostLearningAssignmentsAggregatesQueryBadRequest() *PostLearningAssignmentsAggregatesQueryBadRequest {
	return &PostLearningAssignmentsAggregatesQueryBadRequest{}
}

/*PostLearningAssignmentsAggregatesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLearningAssignmentsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryUnauthorized creates a PostLearningAssignmentsAggregatesQueryUnauthorized with default headers values
func NewPostLearningAssignmentsAggregatesQueryUnauthorized() *PostLearningAssignmentsAggregatesQueryUnauthorized {
	return &PostLearningAssignmentsAggregatesQueryUnauthorized{}
}

/*PostLearningAssignmentsAggregatesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLearningAssignmentsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryForbidden creates a PostLearningAssignmentsAggregatesQueryForbidden with default headers values
func NewPostLearningAssignmentsAggregatesQueryForbidden() *PostLearningAssignmentsAggregatesQueryForbidden {
	return &PostLearningAssignmentsAggregatesQueryForbidden{}
}

/*PostLearningAssignmentsAggregatesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLearningAssignmentsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryNotFound creates a PostLearningAssignmentsAggregatesQueryNotFound with default headers values
func NewPostLearningAssignmentsAggregatesQueryNotFound() *PostLearningAssignmentsAggregatesQueryNotFound {
	return &PostLearningAssignmentsAggregatesQueryNotFound{}
}

/*PostLearningAssignmentsAggregatesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLearningAssignmentsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryRequestTimeout creates a PostLearningAssignmentsAggregatesQueryRequestTimeout with default headers values
func NewPostLearningAssignmentsAggregatesQueryRequestTimeout() *PostLearningAssignmentsAggregatesQueryRequestTimeout {
	return &PostLearningAssignmentsAggregatesQueryRequestTimeout{}
}

/*PostLearningAssignmentsAggregatesQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLearningAssignmentsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryRequestEntityTooLarge creates a PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostLearningAssignmentsAggregatesQueryRequestEntityTooLarge() *PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge {
	return &PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge{}
}

/*PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryUnsupportedMediaType creates a PostLearningAssignmentsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostLearningAssignmentsAggregatesQueryUnsupportedMediaType() *PostLearningAssignmentsAggregatesQueryUnsupportedMediaType {
	return &PostLearningAssignmentsAggregatesQueryUnsupportedMediaType{}
}

/*PostLearningAssignmentsAggregatesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLearningAssignmentsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryTooManyRequests creates a PostLearningAssignmentsAggregatesQueryTooManyRequests with default headers values
func NewPostLearningAssignmentsAggregatesQueryTooManyRequests() *PostLearningAssignmentsAggregatesQueryTooManyRequests {
	return &PostLearningAssignmentsAggregatesQueryTooManyRequests{}
}

/*PostLearningAssignmentsAggregatesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLearningAssignmentsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryInternalServerError creates a PostLearningAssignmentsAggregatesQueryInternalServerError with default headers values
func NewPostLearningAssignmentsAggregatesQueryInternalServerError() *PostLearningAssignmentsAggregatesQueryInternalServerError {
	return &PostLearningAssignmentsAggregatesQueryInternalServerError{}
}

/*PostLearningAssignmentsAggregatesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLearningAssignmentsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryServiceUnavailable creates a PostLearningAssignmentsAggregatesQueryServiceUnavailable with default headers values
func NewPostLearningAssignmentsAggregatesQueryServiceUnavailable() *PostLearningAssignmentsAggregatesQueryServiceUnavailable {
	return &PostLearningAssignmentsAggregatesQueryServiceUnavailable{}
}

/*PostLearningAssignmentsAggregatesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLearningAssignmentsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLearningAssignmentsAggregatesQueryGatewayTimeout creates a PostLearningAssignmentsAggregatesQueryGatewayTimeout with default headers values
func NewPostLearningAssignmentsAggregatesQueryGatewayTimeout() *PostLearningAssignmentsAggregatesQueryGatewayTimeout {
	return &PostLearningAssignmentsAggregatesQueryGatewayTimeout{}
}

/*PostLearningAssignmentsAggregatesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLearningAssignmentsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLearningAssignmentsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/learning/assignments/aggregates/query][%d] postLearningAssignmentsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLearningAssignmentsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLearningAssignmentsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
