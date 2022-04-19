// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostQualityEvaluationsAggregatesQueryMeReader is a Reader for the PostQualityEvaluationsAggregatesQueryMe structure.
type PostQualityEvaluationsAggregatesQueryMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityEvaluationsAggregatesQueryMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityEvaluationsAggregatesQueryMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityEvaluationsAggregatesQueryMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityEvaluationsAggregatesQueryMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityEvaluationsAggregatesQueryMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityEvaluationsAggregatesQueryMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityEvaluationsAggregatesQueryMeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityEvaluationsAggregatesQueryMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityEvaluationsAggregatesQueryMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityEvaluationsAggregatesQueryMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityEvaluationsAggregatesQueryMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityEvaluationsAggregatesQueryMeOK creates a PostQualityEvaluationsAggregatesQueryMeOK with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeOK() *PostQualityEvaluationsAggregatesQueryMeOK {
	return &PostQualityEvaluationsAggregatesQueryMeOK{}
}

/*PostQualityEvaluationsAggregatesQueryMeOK handles this case with default header values.

successful operation
*/
type PostQualityEvaluationsAggregatesQueryMeOK struct {
	Payload *models.EvaluationAggregateQueryResponse
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeOK  %+v", 200, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) GetPayload() *models.EvaluationAggregateQueryResponse {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeBadRequest creates a PostQualityEvaluationsAggregatesQueryMeBadRequest with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeBadRequest() *PostQualityEvaluationsAggregatesQueryMeBadRequest {
	return &PostQualityEvaluationsAggregatesQueryMeBadRequest{}
}

/*PostQualityEvaluationsAggregatesQueryMeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityEvaluationsAggregatesQueryMeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeUnauthorized creates a PostQualityEvaluationsAggregatesQueryMeUnauthorized with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeUnauthorized() *PostQualityEvaluationsAggregatesQueryMeUnauthorized {
	return &PostQualityEvaluationsAggregatesQueryMeUnauthorized{}
}

/*PostQualityEvaluationsAggregatesQueryMeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityEvaluationsAggregatesQueryMeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeForbidden creates a PostQualityEvaluationsAggregatesQueryMeForbidden with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeForbidden() *PostQualityEvaluationsAggregatesQueryMeForbidden {
	return &PostQualityEvaluationsAggregatesQueryMeForbidden{}
}

/*PostQualityEvaluationsAggregatesQueryMeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityEvaluationsAggregatesQueryMeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeNotFound creates a PostQualityEvaluationsAggregatesQueryMeNotFound with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeNotFound() *PostQualityEvaluationsAggregatesQueryMeNotFound {
	return &PostQualityEvaluationsAggregatesQueryMeNotFound{}
}

/*PostQualityEvaluationsAggregatesQueryMeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualityEvaluationsAggregatesQueryMeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeRequestTimeout creates a PostQualityEvaluationsAggregatesQueryMeRequestTimeout with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeRequestTimeout() *PostQualityEvaluationsAggregatesQueryMeRequestTimeout {
	return &PostQualityEvaluationsAggregatesQueryMeRequestTimeout{}
}

/*PostQualityEvaluationsAggregatesQueryMeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityEvaluationsAggregatesQueryMeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge creates a PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge() *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge {
	return &PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge{}
}

/*PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType creates a PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType() *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType {
	return &PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType{}
}

/*PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeTooManyRequests creates a PostQualityEvaluationsAggregatesQueryMeTooManyRequests with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeTooManyRequests() *PostQualityEvaluationsAggregatesQueryMeTooManyRequests {
	return &PostQualityEvaluationsAggregatesQueryMeTooManyRequests{}
}

/*PostQualityEvaluationsAggregatesQueryMeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityEvaluationsAggregatesQueryMeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeInternalServerError creates a PostQualityEvaluationsAggregatesQueryMeInternalServerError with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeInternalServerError() *PostQualityEvaluationsAggregatesQueryMeInternalServerError {
	return &PostQualityEvaluationsAggregatesQueryMeInternalServerError{}
}

/*PostQualityEvaluationsAggregatesQueryMeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityEvaluationsAggregatesQueryMeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeServiceUnavailable creates a PostQualityEvaluationsAggregatesQueryMeServiceUnavailable with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeServiceUnavailable() *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable {
	return &PostQualityEvaluationsAggregatesQueryMeServiceUnavailable{}
}

/*PostQualityEvaluationsAggregatesQueryMeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityEvaluationsAggregatesQueryMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeGatewayTimeout creates a PostQualityEvaluationsAggregatesQueryMeGatewayTimeout with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeGatewayTimeout() *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout {
	return &PostQualityEvaluationsAggregatesQueryMeGatewayTimeout{}
}

/*PostQualityEvaluationsAggregatesQueryMeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualityEvaluationsAggregatesQueryMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
