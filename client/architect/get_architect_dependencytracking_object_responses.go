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

// GetArchitectDependencytrackingObjectReader is a Reader for the GetArchitectDependencytrackingObject structure.
type GetArchitectDependencytrackingObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectDependencytrackingObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectDependencytrackingObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewGetArchitectDependencytrackingObjectPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectDependencytrackingObjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectDependencytrackingObjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectDependencytrackingObjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectDependencytrackingObjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetArchitectDependencytrackingObjectGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectDependencytrackingObjectRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectDependencytrackingObjectUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectDependencytrackingObjectTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectDependencytrackingObjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectDependencytrackingObjectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectDependencytrackingObjectGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectDependencytrackingObjectOK creates a GetArchitectDependencytrackingObjectOK with default headers values
func NewGetArchitectDependencytrackingObjectOK() *GetArchitectDependencytrackingObjectOK {
	return &GetArchitectDependencytrackingObjectOK{}
}

/*GetArchitectDependencytrackingObjectOK handles this case with default header values.

successful operation
*/
type GetArchitectDependencytrackingObjectOK struct {
	Payload *models.DependencyObject
}

func (o *GetArchitectDependencytrackingObjectOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectOK) GetPayload() *models.DependencyObject {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DependencyObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectPartialContent creates a GetArchitectDependencytrackingObjectPartialContent with default headers values
func NewGetArchitectDependencytrackingObjectPartialContent() *GetArchitectDependencytrackingObjectPartialContent {
	return &GetArchitectDependencytrackingObjectPartialContent{}
}

/*GetArchitectDependencytrackingObjectPartialContent handles this case with default header values.

Partial Content - the org data is being rebuilt or needs to be rebuilt.
*/
type GetArchitectDependencytrackingObjectPartialContent struct {
}

func (o *GetArchitectDependencytrackingObjectPartialContent) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectPartialContent ", 206)
}

func (o *GetArchitectDependencytrackingObjectPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetArchitectDependencytrackingObjectBadRequest creates a GetArchitectDependencytrackingObjectBadRequest with default headers values
func NewGetArchitectDependencytrackingObjectBadRequest() *GetArchitectDependencytrackingObjectBadRequest {
	return &GetArchitectDependencytrackingObjectBadRequest{}
}

/*GetArchitectDependencytrackingObjectBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectDependencytrackingObjectBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectUnauthorized creates a GetArchitectDependencytrackingObjectUnauthorized with default headers values
func NewGetArchitectDependencytrackingObjectUnauthorized() *GetArchitectDependencytrackingObjectUnauthorized {
	return &GetArchitectDependencytrackingObjectUnauthorized{}
}

/*GetArchitectDependencytrackingObjectUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectDependencytrackingObjectUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectForbidden creates a GetArchitectDependencytrackingObjectForbidden with default headers values
func NewGetArchitectDependencytrackingObjectForbidden() *GetArchitectDependencytrackingObjectForbidden {
	return &GetArchitectDependencytrackingObjectForbidden{}
}

/*GetArchitectDependencytrackingObjectForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectDependencytrackingObjectForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectNotFound creates a GetArchitectDependencytrackingObjectNotFound with default headers values
func NewGetArchitectDependencytrackingObjectNotFound() *GetArchitectDependencytrackingObjectNotFound {
	return &GetArchitectDependencytrackingObjectNotFound{}
}

/*GetArchitectDependencytrackingObjectNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectDependencytrackingObjectNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectGone creates a GetArchitectDependencytrackingObjectGone with default headers values
func NewGetArchitectDependencytrackingObjectGone() *GetArchitectDependencytrackingObjectGone {
	return &GetArchitectDependencytrackingObjectGone{}
}

/*GetArchitectDependencytrackingObjectGone handles this case with default header values.

Gone
*/
type GetArchitectDependencytrackingObjectGone struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectGone) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectGone  %+v", 410, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectRequestEntityTooLarge creates a GetArchitectDependencytrackingObjectRequestEntityTooLarge with default headers values
func NewGetArchitectDependencytrackingObjectRequestEntityTooLarge() *GetArchitectDependencytrackingObjectRequestEntityTooLarge {
	return &GetArchitectDependencytrackingObjectRequestEntityTooLarge{}
}

/*GetArchitectDependencytrackingObjectRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectDependencytrackingObjectRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectUnsupportedMediaType creates a GetArchitectDependencytrackingObjectUnsupportedMediaType with default headers values
func NewGetArchitectDependencytrackingObjectUnsupportedMediaType() *GetArchitectDependencytrackingObjectUnsupportedMediaType {
	return &GetArchitectDependencytrackingObjectUnsupportedMediaType{}
}

/*GetArchitectDependencytrackingObjectUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectDependencytrackingObjectUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectTooManyRequests creates a GetArchitectDependencytrackingObjectTooManyRequests with default headers values
func NewGetArchitectDependencytrackingObjectTooManyRequests() *GetArchitectDependencytrackingObjectTooManyRequests {
	return &GetArchitectDependencytrackingObjectTooManyRequests{}
}

/*GetArchitectDependencytrackingObjectTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectDependencytrackingObjectTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectInternalServerError creates a GetArchitectDependencytrackingObjectInternalServerError with default headers values
func NewGetArchitectDependencytrackingObjectInternalServerError() *GetArchitectDependencytrackingObjectInternalServerError {
	return &GetArchitectDependencytrackingObjectInternalServerError{}
}

/*GetArchitectDependencytrackingObjectInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectDependencytrackingObjectInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectServiceUnavailable creates a GetArchitectDependencytrackingObjectServiceUnavailable with default headers values
func NewGetArchitectDependencytrackingObjectServiceUnavailable() *GetArchitectDependencytrackingObjectServiceUnavailable {
	return &GetArchitectDependencytrackingObjectServiceUnavailable{}
}

/*GetArchitectDependencytrackingObjectServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectDependencytrackingObjectServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingObjectGatewayTimeout creates a GetArchitectDependencytrackingObjectGatewayTimeout with default headers values
func NewGetArchitectDependencytrackingObjectGatewayTimeout() *GetArchitectDependencytrackingObjectGatewayTimeout {
	return &GetArchitectDependencytrackingObjectGatewayTimeout{}
}

/*GetArchitectDependencytrackingObjectGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectDependencytrackingObjectGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectDependencytrackingObjectGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/object][%d] getArchitectDependencytrackingObjectGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingObjectGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingObjectGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}