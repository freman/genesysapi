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

// GetArchitectDependencytrackingConsumedresourcesReader is a Reader for the GetArchitectDependencytrackingConsumedresources structure.
type GetArchitectDependencytrackingConsumedresourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectDependencytrackingConsumedresourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectDependencytrackingConsumedresourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewGetArchitectDependencytrackingConsumedresourcesPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectDependencytrackingConsumedresourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectDependencytrackingConsumedresourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectDependencytrackingConsumedresourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectDependencytrackingConsumedresourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetArchitectDependencytrackingConsumedresourcesGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectDependencytrackingConsumedresourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectDependencytrackingConsumedresourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectDependencytrackingConsumedresourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectDependencytrackingConsumedresourcesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectDependencytrackingConsumedresourcesOK creates a GetArchitectDependencytrackingConsumedresourcesOK with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesOK() *GetArchitectDependencytrackingConsumedresourcesOK {
	return &GetArchitectDependencytrackingConsumedresourcesOK{}
}

/*GetArchitectDependencytrackingConsumedresourcesOK handles this case with default header values.

successful operation
*/
type GetArchitectDependencytrackingConsumedresourcesOK struct {
	Payload *models.ConsumedResourcesEntityListing
}

func (o *GetArchitectDependencytrackingConsumedresourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesOK) GetPayload() *models.ConsumedResourcesEntityListing {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsumedResourcesEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesPartialContent creates a GetArchitectDependencytrackingConsumedresourcesPartialContent with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesPartialContent() *GetArchitectDependencytrackingConsumedresourcesPartialContent {
	return &GetArchitectDependencytrackingConsumedresourcesPartialContent{}
}

/*GetArchitectDependencytrackingConsumedresourcesPartialContent handles this case with default header values.

Partial Content - the org data is being rebuilt or needs to be rebuilt.
*/
type GetArchitectDependencytrackingConsumedresourcesPartialContent struct {
}

func (o *GetArchitectDependencytrackingConsumedresourcesPartialContent) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesPartialContent ", 206)
}

func (o *GetArchitectDependencytrackingConsumedresourcesPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesBadRequest creates a GetArchitectDependencytrackingConsumedresourcesBadRequest with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesBadRequest() *GetArchitectDependencytrackingConsumedresourcesBadRequest {
	return &GetArchitectDependencytrackingConsumedresourcesBadRequest{}
}

/*GetArchitectDependencytrackingConsumedresourcesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectDependencytrackingConsumedresourcesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesUnauthorized creates a GetArchitectDependencytrackingConsumedresourcesUnauthorized with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesUnauthorized() *GetArchitectDependencytrackingConsumedresourcesUnauthorized {
	return &GetArchitectDependencytrackingConsumedresourcesUnauthorized{}
}

/*GetArchitectDependencytrackingConsumedresourcesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectDependencytrackingConsumedresourcesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesForbidden creates a GetArchitectDependencytrackingConsumedresourcesForbidden with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesForbidden() *GetArchitectDependencytrackingConsumedresourcesForbidden {
	return &GetArchitectDependencytrackingConsumedresourcesForbidden{}
}

/*GetArchitectDependencytrackingConsumedresourcesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectDependencytrackingConsumedresourcesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesNotFound creates a GetArchitectDependencytrackingConsumedresourcesNotFound with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesNotFound() *GetArchitectDependencytrackingConsumedresourcesNotFound {
	return &GetArchitectDependencytrackingConsumedresourcesNotFound{}
}

/*GetArchitectDependencytrackingConsumedresourcesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectDependencytrackingConsumedresourcesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesGone creates a GetArchitectDependencytrackingConsumedresourcesGone with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesGone() *GetArchitectDependencytrackingConsumedresourcesGone {
	return &GetArchitectDependencytrackingConsumedresourcesGone{}
}

/*GetArchitectDependencytrackingConsumedresourcesGone handles this case with default header values.

Gone
*/
type GetArchitectDependencytrackingConsumedresourcesGone struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesGone) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesGone  %+v", 410, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge creates a GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge() *GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge {
	return &GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge{}
}

/*GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType creates a GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType() *GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType {
	return &GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType{}
}

/*GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesTooManyRequests creates a GetArchitectDependencytrackingConsumedresourcesTooManyRequests with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesTooManyRequests() *GetArchitectDependencytrackingConsumedresourcesTooManyRequests {
	return &GetArchitectDependencytrackingConsumedresourcesTooManyRequests{}
}

/*GetArchitectDependencytrackingConsumedresourcesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectDependencytrackingConsumedresourcesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesInternalServerError creates a GetArchitectDependencytrackingConsumedresourcesInternalServerError with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesInternalServerError() *GetArchitectDependencytrackingConsumedresourcesInternalServerError {
	return &GetArchitectDependencytrackingConsumedresourcesInternalServerError{}
}

/*GetArchitectDependencytrackingConsumedresourcesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectDependencytrackingConsumedresourcesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesServiceUnavailable creates a GetArchitectDependencytrackingConsumedresourcesServiceUnavailable with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesServiceUnavailable() *GetArchitectDependencytrackingConsumedresourcesServiceUnavailable {
	return &GetArchitectDependencytrackingConsumedresourcesServiceUnavailable{}
}

/*GetArchitectDependencytrackingConsumedresourcesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectDependencytrackingConsumedresourcesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingConsumedresourcesGatewayTimeout creates a GetArchitectDependencytrackingConsumedresourcesGatewayTimeout with default headers values
func NewGetArchitectDependencytrackingConsumedresourcesGatewayTimeout() *GetArchitectDependencytrackingConsumedresourcesGatewayTimeout {
	return &GetArchitectDependencytrackingConsumedresourcesGatewayTimeout{}
}

/*GetArchitectDependencytrackingConsumedresourcesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectDependencytrackingConsumedresourcesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingConsumedresourcesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/consumedresources][%d] getArchitectDependencytrackingConsumedresourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingConsumedresourcesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingConsumedresourcesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
