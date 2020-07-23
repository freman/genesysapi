// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetFlowsOutcomesReader is a Reader for the GetFlowsOutcomes structure.
type GetFlowsOutcomesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFlowsOutcomesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFlowsOutcomesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFlowsOutcomesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFlowsOutcomesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFlowsOutcomesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFlowsOutcomesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFlowsOutcomesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetFlowsOutcomesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetFlowsOutcomesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFlowsOutcomesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFlowsOutcomesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetFlowsOutcomesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetFlowsOutcomesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFlowsOutcomesOK creates a GetFlowsOutcomesOK with default headers values
func NewGetFlowsOutcomesOK() *GetFlowsOutcomesOK {
	return &GetFlowsOutcomesOK{}
}

/*GetFlowsOutcomesOK handles this case with default header values.

successful operation
*/
type GetFlowsOutcomesOK struct {
	Payload *models.FlowOutcomeListing
}

func (o *GetFlowsOutcomesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesOK  %+v", 200, o.Payload)
}

func (o *GetFlowsOutcomesOK) GetPayload() *models.FlowOutcomeListing {
	return o.Payload
}

func (o *GetFlowsOutcomesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowOutcomeListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesBadRequest creates a GetFlowsOutcomesBadRequest with default headers values
func NewGetFlowsOutcomesBadRequest() *GetFlowsOutcomesBadRequest {
	return &GetFlowsOutcomesBadRequest{}
}

/*GetFlowsOutcomesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetFlowsOutcomesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesBadRequest  %+v", 400, o.Payload)
}

func (o *GetFlowsOutcomesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesUnauthorized creates a GetFlowsOutcomesUnauthorized with default headers values
func NewGetFlowsOutcomesUnauthorized() *GetFlowsOutcomesUnauthorized {
	return &GetFlowsOutcomesUnauthorized{}
}

/*GetFlowsOutcomesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetFlowsOutcomesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFlowsOutcomesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesForbidden creates a GetFlowsOutcomesForbidden with default headers values
func NewGetFlowsOutcomesForbidden() *GetFlowsOutcomesForbidden {
	return &GetFlowsOutcomesForbidden{}
}

/*GetFlowsOutcomesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetFlowsOutcomesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesForbidden  %+v", 403, o.Payload)
}

func (o *GetFlowsOutcomesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesNotFound creates a GetFlowsOutcomesNotFound with default headers values
func NewGetFlowsOutcomesNotFound() *GetFlowsOutcomesNotFound {
	return &GetFlowsOutcomesNotFound{}
}

/*GetFlowsOutcomesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetFlowsOutcomesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesNotFound  %+v", 404, o.Payload)
}

func (o *GetFlowsOutcomesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesMethodNotAllowed creates a GetFlowsOutcomesMethodNotAllowed with default headers values
func NewGetFlowsOutcomesMethodNotAllowed() *GetFlowsOutcomesMethodNotAllowed {
	return &GetFlowsOutcomesMethodNotAllowed{}
}

/*GetFlowsOutcomesMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetFlowsOutcomesMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetFlowsOutcomesMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesRequestEntityTooLarge creates a GetFlowsOutcomesRequestEntityTooLarge with default headers values
func NewGetFlowsOutcomesRequestEntityTooLarge() *GetFlowsOutcomesRequestEntityTooLarge {
	return &GetFlowsOutcomesRequestEntityTooLarge{}
}

/*GetFlowsOutcomesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetFlowsOutcomesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetFlowsOutcomesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesUnsupportedMediaType creates a GetFlowsOutcomesUnsupportedMediaType with default headers values
func NewGetFlowsOutcomesUnsupportedMediaType() *GetFlowsOutcomesUnsupportedMediaType {
	return &GetFlowsOutcomesUnsupportedMediaType{}
}

/*GetFlowsOutcomesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetFlowsOutcomesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetFlowsOutcomesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesTooManyRequests creates a GetFlowsOutcomesTooManyRequests with default headers values
func NewGetFlowsOutcomesTooManyRequests() *GetFlowsOutcomesTooManyRequests {
	return &GetFlowsOutcomesTooManyRequests{}
}

/*GetFlowsOutcomesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetFlowsOutcomesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetFlowsOutcomesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesInternalServerError creates a GetFlowsOutcomesInternalServerError with default headers values
func NewGetFlowsOutcomesInternalServerError() *GetFlowsOutcomesInternalServerError {
	return &GetFlowsOutcomesInternalServerError{}
}

/*GetFlowsOutcomesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetFlowsOutcomesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFlowsOutcomesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesServiceUnavailable creates a GetFlowsOutcomesServiceUnavailable with default headers values
func NewGetFlowsOutcomesServiceUnavailable() *GetFlowsOutcomesServiceUnavailable {
	return &GetFlowsOutcomesServiceUnavailable{}
}

/*GetFlowsOutcomesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetFlowsOutcomesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetFlowsOutcomesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFlowsOutcomesGatewayTimeout creates a GetFlowsOutcomesGatewayTimeout with default headers values
func NewGetFlowsOutcomesGatewayTimeout() *GetFlowsOutcomesGatewayTimeout {
	return &GetFlowsOutcomesGatewayTimeout{}
}

/*GetFlowsOutcomesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetFlowsOutcomesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetFlowsOutcomesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/flows/outcomes][%d] getFlowsOutcomesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetFlowsOutcomesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetFlowsOutcomesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
