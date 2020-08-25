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

// GetArchitectDependencytrackingDeletedresourceconsumersReader is a Reader for the GetArchitectDependencytrackingDeletedresourceconsumers structure.
type GetArchitectDependencytrackingDeletedresourceconsumersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectDependencytrackingDeletedresourceconsumersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersOK creates a GetArchitectDependencytrackingDeletedresourceconsumersOK with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersOK() *GetArchitectDependencytrackingDeletedresourceconsumersOK {
	return &GetArchitectDependencytrackingDeletedresourceconsumersOK{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersOK handles this case with default header values.

successful operation
*/
type GetArchitectDependencytrackingDeletedresourceconsumersOK struct {
	Payload *models.DependencyObjectEntityListing
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersOK) GetPayload() *models.DependencyObjectEntityListing {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DependencyObjectEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersPartialContent creates a GetArchitectDependencytrackingDeletedresourceconsumersPartialContent with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersPartialContent() *GetArchitectDependencytrackingDeletedresourceconsumersPartialContent {
	return &GetArchitectDependencytrackingDeletedresourceconsumersPartialContent{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersPartialContent handles this case with default header values.

Partial Content - the org data is being rebuilt or needs to be rebuilt.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersPartialContent struct {
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersPartialContent) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersPartialContent ", 206)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersBadRequest creates a GetArchitectDependencytrackingDeletedresourceconsumersBadRequest with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersBadRequest() *GetArchitectDependencytrackingDeletedresourceconsumersBadRequest {
	return &GetArchitectDependencytrackingDeletedresourceconsumersBadRequest{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersUnauthorized creates a GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersUnauthorized() *GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized {
	return &GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersForbidden creates a GetArchitectDependencytrackingDeletedresourceconsumersForbidden with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersForbidden() *GetArchitectDependencytrackingDeletedresourceconsumersForbidden {
	return &GetArchitectDependencytrackingDeletedresourceconsumersForbidden{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersNotFound creates a GetArchitectDependencytrackingDeletedresourceconsumersNotFound with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersNotFound() *GetArchitectDependencytrackingDeletedresourceconsumersNotFound {
	return &GetArchitectDependencytrackingDeletedresourceconsumersNotFound{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge creates a GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge() *GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge {
	return &GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType creates a GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType() *GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType {
	return &GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests creates a GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests() *GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests {
	return &GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersInternalServerError creates a GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersInternalServerError() *GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError {
	return &GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable creates a GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable() *GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable {
	return &GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout creates a GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout with default headers values
func NewGetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout() *GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout {
	return &GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout{}
}

/*GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/deletedresourceconsumers][%d] getArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingDeletedresourceconsumersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}