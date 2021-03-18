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

// PostJourneyOutcomesReader is a Reader for the PostJourneyOutcomes structure.
type PostJourneyOutcomesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostJourneyOutcomesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostJourneyOutcomesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostJourneyOutcomesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostJourneyOutcomesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostJourneyOutcomesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostJourneyOutcomesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostJourneyOutcomesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostJourneyOutcomesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostJourneyOutcomesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostJourneyOutcomesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostJourneyOutcomesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostJourneyOutcomesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostJourneyOutcomesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostJourneyOutcomesOK creates a PostJourneyOutcomesOK with default headers values
func NewPostJourneyOutcomesOK() *PostJourneyOutcomesOK {
	return &PostJourneyOutcomesOK{}
}

/*PostJourneyOutcomesOK handles this case with default header values.

successful operation
*/
type PostJourneyOutcomesOK struct {
	Payload *models.Outcome
}

func (o *PostJourneyOutcomesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesOK  %+v", 200, o.Payload)
}

func (o *PostJourneyOutcomesOK) GetPayload() *models.Outcome {
	return o.Payload
}

func (o *PostJourneyOutcomesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Outcome)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesCreated creates a PostJourneyOutcomesCreated with default headers values
func NewPostJourneyOutcomesCreated() *PostJourneyOutcomesCreated {
	return &PostJourneyOutcomesCreated{}
}

/*PostJourneyOutcomesCreated handles this case with default header values.

Outcome created.
*/
type PostJourneyOutcomesCreated struct {
	Payload *models.Outcome
}

func (o *PostJourneyOutcomesCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesCreated  %+v", 201, o.Payload)
}

func (o *PostJourneyOutcomesCreated) GetPayload() *models.Outcome {
	return o.Payload
}

func (o *PostJourneyOutcomesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Outcome)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesBadRequest creates a PostJourneyOutcomesBadRequest with default headers values
func NewPostJourneyOutcomesBadRequest() *PostJourneyOutcomesBadRequest {
	return &PostJourneyOutcomesBadRequest{}
}

/*PostJourneyOutcomesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostJourneyOutcomesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesBadRequest  %+v", 400, o.Payload)
}

func (o *PostJourneyOutcomesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesUnauthorized creates a PostJourneyOutcomesUnauthorized with default headers values
func NewPostJourneyOutcomesUnauthorized() *PostJourneyOutcomesUnauthorized {
	return &PostJourneyOutcomesUnauthorized{}
}

/*PostJourneyOutcomesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostJourneyOutcomesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostJourneyOutcomesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesForbidden creates a PostJourneyOutcomesForbidden with default headers values
func NewPostJourneyOutcomesForbidden() *PostJourneyOutcomesForbidden {
	return &PostJourneyOutcomesForbidden{}
}

/*PostJourneyOutcomesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostJourneyOutcomesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesForbidden  %+v", 403, o.Payload)
}

func (o *PostJourneyOutcomesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesNotFound creates a PostJourneyOutcomesNotFound with default headers values
func NewPostJourneyOutcomesNotFound() *PostJourneyOutcomesNotFound {
	return &PostJourneyOutcomesNotFound{}
}

/*PostJourneyOutcomesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostJourneyOutcomesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesNotFound  %+v", 404, o.Payload)
}

func (o *PostJourneyOutcomesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesRequestEntityTooLarge creates a PostJourneyOutcomesRequestEntityTooLarge with default headers values
func NewPostJourneyOutcomesRequestEntityTooLarge() *PostJourneyOutcomesRequestEntityTooLarge {
	return &PostJourneyOutcomesRequestEntityTooLarge{}
}

/*PostJourneyOutcomesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostJourneyOutcomesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostJourneyOutcomesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesUnsupportedMediaType creates a PostJourneyOutcomesUnsupportedMediaType with default headers values
func NewPostJourneyOutcomesUnsupportedMediaType() *PostJourneyOutcomesUnsupportedMediaType {
	return &PostJourneyOutcomesUnsupportedMediaType{}
}

/*PostJourneyOutcomesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostJourneyOutcomesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostJourneyOutcomesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesTooManyRequests creates a PostJourneyOutcomesTooManyRequests with default headers values
func NewPostJourneyOutcomesTooManyRequests() *PostJourneyOutcomesTooManyRequests {
	return &PostJourneyOutcomesTooManyRequests{}
}

/*PostJourneyOutcomesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostJourneyOutcomesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostJourneyOutcomesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesInternalServerError creates a PostJourneyOutcomesInternalServerError with default headers values
func NewPostJourneyOutcomesInternalServerError() *PostJourneyOutcomesInternalServerError {
	return &PostJourneyOutcomesInternalServerError{}
}

/*PostJourneyOutcomesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostJourneyOutcomesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostJourneyOutcomesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesServiceUnavailable creates a PostJourneyOutcomesServiceUnavailable with default headers values
func NewPostJourneyOutcomesServiceUnavailable() *PostJourneyOutcomesServiceUnavailable {
	return &PostJourneyOutcomesServiceUnavailable{}
}

/*PostJourneyOutcomesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostJourneyOutcomesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostJourneyOutcomesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostJourneyOutcomesGatewayTimeout creates a PostJourneyOutcomesGatewayTimeout with default headers values
func NewPostJourneyOutcomesGatewayTimeout() *PostJourneyOutcomesGatewayTimeout {
	return &PostJourneyOutcomesGatewayTimeout{}
}

/*PostJourneyOutcomesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostJourneyOutcomesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostJourneyOutcomesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/journey/outcomes][%d] postJourneyOutcomesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostJourneyOutcomesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostJourneyOutcomesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}