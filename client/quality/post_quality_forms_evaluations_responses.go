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

// PostQualityFormsEvaluationsReader is a Reader for the PostQualityFormsEvaluations structure.
type PostQualityFormsEvaluationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityFormsEvaluationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityFormsEvaluationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityFormsEvaluationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityFormsEvaluationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityFormsEvaluationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityFormsEvaluationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityFormsEvaluationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostQualityFormsEvaluationsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityFormsEvaluationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityFormsEvaluationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityFormsEvaluationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityFormsEvaluationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityFormsEvaluationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityFormsEvaluationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityFormsEvaluationsOK creates a PostQualityFormsEvaluationsOK with default headers values
func NewPostQualityFormsEvaluationsOK() *PostQualityFormsEvaluationsOK {
	return &PostQualityFormsEvaluationsOK{}
}

/*PostQualityFormsEvaluationsOK handles this case with default header values.

successful operation
*/
type PostQualityFormsEvaluationsOK struct {
	Payload *models.EvaluationForm
}

func (o *PostQualityFormsEvaluationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsOK  %+v", 200, o.Payload)
}

func (o *PostQualityFormsEvaluationsOK) GetPayload() *models.EvaluationForm {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationForm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsBadRequest creates a PostQualityFormsEvaluationsBadRequest with default headers values
func NewPostQualityFormsEvaluationsBadRequest() *PostQualityFormsEvaluationsBadRequest {
	return &PostQualityFormsEvaluationsBadRequest{}
}

/*PostQualityFormsEvaluationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityFormsEvaluationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityFormsEvaluationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsUnauthorized creates a PostQualityFormsEvaluationsUnauthorized with default headers values
func NewPostQualityFormsEvaluationsUnauthorized() *PostQualityFormsEvaluationsUnauthorized {
	return &PostQualityFormsEvaluationsUnauthorized{}
}

/*PostQualityFormsEvaluationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityFormsEvaluationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityFormsEvaluationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsForbidden creates a PostQualityFormsEvaluationsForbidden with default headers values
func NewPostQualityFormsEvaluationsForbidden() *PostQualityFormsEvaluationsForbidden {
	return &PostQualityFormsEvaluationsForbidden{}
}

/*PostQualityFormsEvaluationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityFormsEvaluationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityFormsEvaluationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsNotFound creates a PostQualityFormsEvaluationsNotFound with default headers values
func NewPostQualityFormsEvaluationsNotFound() *PostQualityFormsEvaluationsNotFound {
	return &PostQualityFormsEvaluationsNotFound{}
}

/*PostQualityFormsEvaluationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualityFormsEvaluationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityFormsEvaluationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsRequestTimeout creates a PostQualityFormsEvaluationsRequestTimeout with default headers values
func NewPostQualityFormsEvaluationsRequestTimeout() *PostQualityFormsEvaluationsRequestTimeout {
	return &PostQualityFormsEvaluationsRequestTimeout{}
}

/*PostQualityFormsEvaluationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityFormsEvaluationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityFormsEvaluationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsConflict creates a PostQualityFormsEvaluationsConflict with default headers values
func NewPostQualityFormsEvaluationsConflict() *PostQualityFormsEvaluationsConflict {
	return &PostQualityFormsEvaluationsConflict{}
}

/*PostQualityFormsEvaluationsConflict handles this case with default header values.

Conflict
*/
type PostQualityFormsEvaluationsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsConflict  %+v", 409, o.Payload)
}

func (o *PostQualityFormsEvaluationsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsRequestEntityTooLarge creates a PostQualityFormsEvaluationsRequestEntityTooLarge with default headers values
func NewPostQualityFormsEvaluationsRequestEntityTooLarge() *PostQualityFormsEvaluationsRequestEntityTooLarge {
	return &PostQualityFormsEvaluationsRequestEntityTooLarge{}
}

/*PostQualityFormsEvaluationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostQualityFormsEvaluationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsUnsupportedMediaType creates a PostQualityFormsEvaluationsUnsupportedMediaType with default headers values
func NewPostQualityFormsEvaluationsUnsupportedMediaType() *PostQualityFormsEvaluationsUnsupportedMediaType {
	return &PostQualityFormsEvaluationsUnsupportedMediaType{}
}

/*PostQualityFormsEvaluationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityFormsEvaluationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsTooManyRequests creates a PostQualityFormsEvaluationsTooManyRequests with default headers values
func NewPostQualityFormsEvaluationsTooManyRequests() *PostQualityFormsEvaluationsTooManyRequests {
	return &PostQualityFormsEvaluationsTooManyRequests{}
}

/*PostQualityFormsEvaluationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityFormsEvaluationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityFormsEvaluationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsInternalServerError creates a PostQualityFormsEvaluationsInternalServerError with default headers values
func NewPostQualityFormsEvaluationsInternalServerError() *PostQualityFormsEvaluationsInternalServerError {
	return &PostQualityFormsEvaluationsInternalServerError{}
}

/*PostQualityFormsEvaluationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityFormsEvaluationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityFormsEvaluationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsServiceUnavailable creates a PostQualityFormsEvaluationsServiceUnavailable with default headers values
func NewPostQualityFormsEvaluationsServiceUnavailable() *PostQualityFormsEvaluationsServiceUnavailable {
	return &PostQualityFormsEvaluationsServiceUnavailable{}
}

/*PostQualityFormsEvaluationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityFormsEvaluationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityFormsEvaluationsGatewayTimeout creates a PostQualityFormsEvaluationsGatewayTimeout with default headers values
func NewPostQualityFormsEvaluationsGatewayTimeout() *PostQualityFormsEvaluationsGatewayTimeout {
	return &PostQualityFormsEvaluationsGatewayTimeout{}
}

/*PostQualityFormsEvaluationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualityFormsEvaluationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/forms/evaluations][%d] postQualityFormsEvaluationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityFormsEvaluationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
