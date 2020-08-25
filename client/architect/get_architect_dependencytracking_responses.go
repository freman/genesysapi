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

// GetArchitectDependencytrackingReader is a Reader for the GetArchitectDependencytracking structure.
type GetArchitectDependencytrackingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectDependencytrackingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectDependencytrackingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewGetArchitectDependencytrackingPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectDependencytrackingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectDependencytrackingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectDependencytrackingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectDependencytrackingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectDependencytrackingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectDependencytrackingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectDependencytrackingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectDependencytrackingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectDependencytrackingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectDependencytrackingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectDependencytrackingOK creates a GetArchitectDependencytrackingOK with default headers values
func NewGetArchitectDependencytrackingOK() *GetArchitectDependencytrackingOK {
	return &GetArchitectDependencytrackingOK{}
}

/*GetArchitectDependencytrackingOK handles this case with default header values.

successful operation
*/
type GetArchitectDependencytrackingOK struct {
	Payload *models.DependencyObjectEntityListing
}

func (o *GetArchitectDependencytrackingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingOK) GetPayload() *models.DependencyObjectEntityListing {
	return o.Payload
}

func (o *GetArchitectDependencytrackingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DependencyObjectEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingPartialContent creates a GetArchitectDependencytrackingPartialContent with default headers values
func NewGetArchitectDependencytrackingPartialContent() *GetArchitectDependencytrackingPartialContent {
	return &GetArchitectDependencytrackingPartialContent{}
}

/*GetArchitectDependencytrackingPartialContent handles this case with default header values.

Partial Content - the organization's data is being rebuilt or needs to be rebuilt.
*/
type GetArchitectDependencytrackingPartialContent struct {
}

func (o *GetArchitectDependencytrackingPartialContent) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingPartialContent ", 206)
}

func (o *GetArchitectDependencytrackingPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetArchitectDependencytrackingBadRequest creates a GetArchitectDependencytrackingBadRequest with default headers values
func NewGetArchitectDependencytrackingBadRequest() *GetArchitectDependencytrackingBadRequest {
	return &GetArchitectDependencytrackingBadRequest{}
}

/*GetArchitectDependencytrackingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectDependencytrackingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingUnauthorized creates a GetArchitectDependencytrackingUnauthorized with default headers values
func NewGetArchitectDependencytrackingUnauthorized() *GetArchitectDependencytrackingUnauthorized {
	return &GetArchitectDependencytrackingUnauthorized{}
}

/*GetArchitectDependencytrackingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectDependencytrackingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingForbidden creates a GetArchitectDependencytrackingForbidden with default headers values
func NewGetArchitectDependencytrackingForbidden() *GetArchitectDependencytrackingForbidden {
	return &GetArchitectDependencytrackingForbidden{}
}

/*GetArchitectDependencytrackingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectDependencytrackingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingNotFound creates a GetArchitectDependencytrackingNotFound with default headers values
func NewGetArchitectDependencytrackingNotFound() *GetArchitectDependencytrackingNotFound {
	return &GetArchitectDependencytrackingNotFound{}
}

/*GetArchitectDependencytrackingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectDependencytrackingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingRequestEntityTooLarge creates a GetArchitectDependencytrackingRequestEntityTooLarge with default headers values
func NewGetArchitectDependencytrackingRequestEntityTooLarge() *GetArchitectDependencytrackingRequestEntityTooLarge {
	return &GetArchitectDependencytrackingRequestEntityTooLarge{}
}

/*GetArchitectDependencytrackingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectDependencytrackingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingUnsupportedMediaType creates a GetArchitectDependencytrackingUnsupportedMediaType with default headers values
func NewGetArchitectDependencytrackingUnsupportedMediaType() *GetArchitectDependencytrackingUnsupportedMediaType {
	return &GetArchitectDependencytrackingUnsupportedMediaType{}
}

/*GetArchitectDependencytrackingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectDependencytrackingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTooManyRequests creates a GetArchitectDependencytrackingTooManyRequests with default headers values
func NewGetArchitectDependencytrackingTooManyRequests() *GetArchitectDependencytrackingTooManyRequests {
	return &GetArchitectDependencytrackingTooManyRequests{}
}

/*GetArchitectDependencytrackingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectDependencytrackingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingInternalServerError creates a GetArchitectDependencytrackingInternalServerError with default headers values
func NewGetArchitectDependencytrackingInternalServerError() *GetArchitectDependencytrackingInternalServerError {
	return &GetArchitectDependencytrackingInternalServerError{}
}

/*GetArchitectDependencytrackingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectDependencytrackingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingServiceUnavailable creates a GetArchitectDependencytrackingServiceUnavailable with default headers values
func NewGetArchitectDependencytrackingServiceUnavailable() *GetArchitectDependencytrackingServiceUnavailable {
	return &GetArchitectDependencytrackingServiceUnavailable{}
}

/*GetArchitectDependencytrackingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectDependencytrackingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingGatewayTimeout creates a GetArchitectDependencytrackingGatewayTimeout with default headers values
func NewGetArchitectDependencytrackingGatewayTimeout() *GetArchitectDependencytrackingGatewayTimeout {
	return &GetArchitectDependencytrackingGatewayTimeout{}
}

/*GetArchitectDependencytrackingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectDependencytrackingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking][%d] getArchitectDependencytrackingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}