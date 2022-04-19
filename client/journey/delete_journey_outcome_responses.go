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

// DeleteJourneyOutcomeReader is a Reader for the DeleteJourneyOutcome structure.
type DeleteJourneyOutcomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJourneyOutcomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteJourneyOutcomeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteJourneyOutcomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteJourneyOutcomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteJourneyOutcomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteJourneyOutcomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteJourneyOutcomeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteJourneyOutcomeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteJourneyOutcomeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteJourneyOutcomeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJourneyOutcomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteJourneyOutcomeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteJourneyOutcomeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteJourneyOutcomeNoContent creates a DeleteJourneyOutcomeNoContent with default headers values
func NewDeleteJourneyOutcomeNoContent() *DeleteJourneyOutcomeNoContent {
	return &DeleteJourneyOutcomeNoContent{}
}

/*DeleteJourneyOutcomeNoContent handles this case with default header values.

Outcome deleted.
*/
type DeleteJourneyOutcomeNoContent struct {
}

func (o *DeleteJourneyOutcomeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeNoContent ", 204)
}

func (o *DeleteJourneyOutcomeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJourneyOutcomeBadRequest creates a DeleteJourneyOutcomeBadRequest with default headers values
func NewDeleteJourneyOutcomeBadRequest() *DeleteJourneyOutcomeBadRequest {
	return &DeleteJourneyOutcomeBadRequest{}
}

/*DeleteJourneyOutcomeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteJourneyOutcomeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteJourneyOutcomeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeUnauthorized creates a DeleteJourneyOutcomeUnauthorized with default headers values
func NewDeleteJourneyOutcomeUnauthorized() *DeleteJourneyOutcomeUnauthorized {
	return &DeleteJourneyOutcomeUnauthorized{}
}

/*DeleteJourneyOutcomeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteJourneyOutcomeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteJourneyOutcomeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeForbidden creates a DeleteJourneyOutcomeForbidden with default headers values
func NewDeleteJourneyOutcomeForbidden() *DeleteJourneyOutcomeForbidden {
	return &DeleteJourneyOutcomeForbidden{}
}

/*DeleteJourneyOutcomeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteJourneyOutcomeForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteJourneyOutcomeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeNotFound creates a DeleteJourneyOutcomeNotFound with default headers values
func NewDeleteJourneyOutcomeNotFound() *DeleteJourneyOutcomeNotFound {
	return &DeleteJourneyOutcomeNotFound{}
}

/*DeleteJourneyOutcomeNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteJourneyOutcomeNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteJourneyOutcomeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeRequestTimeout creates a DeleteJourneyOutcomeRequestTimeout with default headers values
func NewDeleteJourneyOutcomeRequestTimeout() *DeleteJourneyOutcomeRequestTimeout {
	return &DeleteJourneyOutcomeRequestTimeout{}
}

/*DeleteJourneyOutcomeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteJourneyOutcomeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteJourneyOutcomeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeRequestEntityTooLarge creates a DeleteJourneyOutcomeRequestEntityTooLarge with default headers values
func NewDeleteJourneyOutcomeRequestEntityTooLarge() *DeleteJourneyOutcomeRequestEntityTooLarge {
	return &DeleteJourneyOutcomeRequestEntityTooLarge{}
}

/*DeleteJourneyOutcomeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteJourneyOutcomeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteJourneyOutcomeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeUnsupportedMediaType creates a DeleteJourneyOutcomeUnsupportedMediaType with default headers values
func NewDeleteJourneyOutcomeUnsupportedMediaType() *DeleteJourneyOutcomeUnsupportedMediaType {
	return &DeleteJourneyOutcomeUnsupportedMediaType{}
}

/*DeleteJourneyOutcomeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteJourneyOutcomeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteJourneyOutcomeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeTooManyRequests creates a DeleteJourneyOutcomeTooManyRequests with default headers values
func NewDeleteJourneyOutcomeTooManyRequests() *DeleteJourneyOutcomeTooManyRequests {
	return &DeleteJourneyOutcomeTooManyRequests{}
}

/*DeleteJourneyOutcomeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteJourneyOutcomeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteJourneyOutcomeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeInternalServerError creates a DeleteJourneyOutcomeInternalServerError with default headers values
func NewDeleteJourneyOutcomeInternalServerError() *DeleteJourneyOutcomeInternalServerError {
	return &DeleteJourneyOutcomeInternalServerError{}
}

/*DeleteJourneyOutcomeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteJourneyOutcomeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteJourneyOutcomeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeServiceUnavailable creates a DeleteJourneyOutcomeServiceUnavailable with default headers values
func NewDeleteJourneyOutcomeServiceUnavailable() *DeleteJourneyOutcomeServiceUnavailable {
	return &DeleteJourneyOutcomeServiceUnavailable{}
}

/*DeleteJourneyOutcomeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteJourneyOutcomeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteJourneyOutcomeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyOutcomeGatewayTimeout creates a DeleteJourneyOutcomeGatewayTimeout with default headers values
func NewDeleteJourneyOutcomeGatewayTimeout() *DeleteJourneyOutcomeGatewayTimeout {
	return &DeleteJourneyOutcomeGatewayTimeout{}
}

/*DeleteJourneyOutcomeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteJourneyOutcomeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyOutcomeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/outcomes/{outcomeId}][%d] deleteJourneyOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteJourneyOutcomeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyOutcomeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
