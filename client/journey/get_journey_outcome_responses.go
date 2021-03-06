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

// GetJourneyOutcomeReader is a Reader for the GetJourneyOutcome structure.
type GetJourneyOutcomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyOutcomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyOutcomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyOutcomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyOutcomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyOutcomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyOutcomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyOutcomeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyOutcomeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyOutcomeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyOutcomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyOutcomeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyOutcomeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyOutcomeOK creates a GetJourneyOutcomeOK with default headers values
func NewGetJourneyOutcomeOK() *GetJourneyOutcomeOK {
	return &GetJourneyOutcomeOK{}
}

/*GetJourneyOutcomeOK handles this case with default header values.

successful operation
*/
type GetJourneyOutcomeOK struct {
	Payload *models.Outcome
}

func (o *GetJourneyOutcomeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeOK  %+v", 200, o.Payload)
}

func (o *GetJourneyOutcomeOK) GetPayload() *models.Outcome {
	return o.Payload
}

func (o *GetJourneyOutcomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Outcome)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeBadRequest creates a GetJourneyOutcomeBadRequest with default headers values
func NewGetJourneyOutcomeBadRequest() *GetJourneyOutcomeBadRequest {
	return &GetJourneyOutcomeBadRequest{}
}

/*GetJourneyOutcomeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyOutcomeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyOutcomeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeUnauthorized creates a GetJourneyOutcomeUnauthorized with default headers values
func NewGetJourneyOutcomeUnauthorized() *GetJourneyOutcomeUnauthorized {
	return &GetJourneyOutcomeUnauthorized{}
}

/*GetJourneyOutcomeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyOutcomeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyOutcomeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeForbidden creates a GetJourneyOutcomeForbidden with default headers values
func NewGetJourneyOutcomeForbidden() *GetJourneyOutcomeForbidden {
	return &GetJourneyOutcomeForbidden{}
}

/*GetJourneyOutcomeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyOutcomeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyOutcomeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeNotFound creates a GetJourneyOutcomeNotFound with default headers values
func NewGetJourneyOutcomeNotFound() *GetJourneyOutcomeNotFound {
	return &GetJourneyOutcomeNotFound{}
}

/*GetJourneyOutcomeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetJourneyOutcomeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyOutcomeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeRequestEntityTooLarge creates a GetJourneyOutcomeRequestEntityTooLarge with default headers values
func NewGetJourneyOutcomeRequestEntityTooLarge() *GetJourneyOutcomeRequestEntityTooLarge {
	return &GetJourneyOutcomeRequestEntityTooLarge{}
}

/*GetJourneyOutcomeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetJourneyOutcomeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeUnsupportedMediaType creates a GetJourneyOutcomeUnsupportedMediaType with default headers values
func NewGetJourneyOutcomeUnsupportedMediaType() *GetJourneyOutcomeUnsupportedMediaType {
	return &GetJourneyOutcomeUnsupportedMediaType{}
}

/*GetJourneyOutcomeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyOutcomeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyOutcomeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeTooManyRequests creates a GetJourneyOutcomeTooManyRequests with default headers values
func NewGetJourneyOutcomeTooManyRequests() *GetJourneyOutcomeTooManyRequests {
	return &GetJourneyOutcomeTooManyRequests{}
}

/*GetJourneyOutcomeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetJourneyOutcomeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyOutcomeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeInternalServerError creates a GetJourneyOutcomeInternalServerError with default headers values
func NewGetJourneyOutcomeInternalServerError() *GetJourneyOutcomeInternalServerError {
	return &GetJourneyOutcomeInternalServerError{}
}

/*GetJourneyOutcomeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyOutcomeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyOutcomeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeServiceUnavailable creates a GetJourneyOutcomeServiceUnavailable with default headers values
func NewGetJourneyOutcomeServiceUnavailable() *GetJourneyOutcomeServiceUnavailable {
	return &GetJourneyOutcomeServiceUnavailable{}
}

/*GetJourneyOutcomeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyOutcomeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyOutcomeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyOutcomeGatewayTimeout creates a GetJourneyOutcomeGatewayTimeout with default headers values
func NewGetJourneyOutcomeGatewayTimeout() *GetJourneyOutcomeGatewayTimeout {
	return &GetJourneyOutcomeGatewayTimeout{}
}

/*GetJourneyOutcomeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetJourneyOutcomeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyOutcomeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/outcomes/{outcomeId}][%d] getJourneyOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyOutcomeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyOutcomeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
