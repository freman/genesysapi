// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScimResourcetypeReader is a Reader for the GetScimResourcetype structure.
type GetScimResourcetypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimResourcetypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimResourcetypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScimResourcetypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimResourcetypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimResourcetypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimResourcetypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimResourcetypeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimResourcetypeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimResourcetypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimResourcetypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimResourcetypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimResourcetypeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimResourcetypeOK creates a GetScimResourcetypeOK with default headers values
func NewGetScimResourcetypeOK() *GetScimResourcetypeOK {
	return &GetScimResourcetypeOK{}
}

/*GetScimResourcetypeOK handles this case with default header values.

successful operation
*/
type GetScimResourcetypeOK struct {
	Payload *models.ScimConfigResourceType
}

func (o *GetScimResourcetypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeOK  %+v", 200, o.Payload)
}

func (o *GetScimResourcetypeOK) GetPayload() *models.ScimConfigResourceType {
	return o.Payload
}

func (o *GetScimResourcetypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimConfigResourceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeBadRequest creates a GetScimResourcetypeBadRequest with default headers values
func NewGetScimResourcetypeBadRequest() *GetScimResourcetypeBadRequest {
	return &GetScimResourcetypeBadRequest{}
}

/*GetScimResourcetypeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimResourcetypeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimResourcetypeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeUnauthorized creates a GetScimResourcetypeUnauthorized with default headers values
func NewGetScimResourcetypeUnauthorized() *GetScimResourcetypeUnauthorized {
	return &GetScimResourcetypeUnauthorized{}
}

/*GetScimResourcetypeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimResourcetypeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimResourcetypeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeForbidden creates a GetScimResourcetypeForbidden with default headers values
func NewGetScimResourcetypeForbidden() *GetScimResourcetypeForbidden {
	return &GetScimResourcetypeForbidden{}
}

/*GetScimResourcetypeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScimResourcetypeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeForbidden  %+v", 403, o.Payload)
}

func (o *GetScimResourcetypeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeNotFound creates a GetScimResourcetypeNotFound with default headers values
func NewGetScimResourcetypeNotFound() *GetScimResourcetypeNotFound {
	return &GetScimResourcetypeNotFound{}
}

/*GetScimResourcetypeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScimResourcetypeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeNotFound  %+v", 404, o.Payload)
}

func (o *GetScimResourcetypeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeRequestEntityTooLarge creates a GetScimResourcetypeRequestEntityTooLarge with default headers values
func NewGetScimResourcetypeRequestEntityTooLarge() *GetScimResourcetypeRequestEntityTooLarge {
	return &GetScimResourcetypeRequestEntityTooLarge{}
}

/*GetScimResourcetypeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScimResourcetypeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimResourcetypeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeUnsupportedMediaType creates a GetScimResourcetypeUnsupportedMediaType with default headers values
func NewGetScimResourcetypeUnsupportedMediaType() *GetScimResourcetypeUnsupportedMediaType {
	return &GetScimResourcetypeUnsupportedMediaType{}
}

/*GetScimResourcetypeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimResourcetypeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimResourcetypeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeTooManyRequests creates a GetScimResourcetypeTooManyRequests with default headers values
func NewGetScimResourcetypeTooManyRequests() *GetScimResourcetypeTooManyRequests {
	return &GetScimResourcetypeTooManyRequests{}
}

/*GetScimResourcetypeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScimResourcetypeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimResourcetypeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeInternalServerError creates a GetScimResourcetypeInternalServerError with default headers values
func NewGetScimResourcetypeInternalServerError() *GetScimResourcetypeInternalServerError {
	return &GetScimResourcetypeInternalServerError{}
}

/*GetScimResourcetypeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimResourcetypeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimResourcetypeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeServiceUnavailable creates a GetScimResourcetypeServiceUnavailable with default headers values
func NewGetScimResourcetypeServiceUnavailable() *GetScimResourcetypeServiceUnavailable {
	return &GetScimResourcetypeServiceUnavailable{}
}

/*GetScimResourcetypeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimResourcetypeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimResourcetypeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimResourcetypeGatewayTimeout creates a GetScimResourcetypeGatewayTimeout with default headers values
func NewGetScimResourcetypeGatewayTimeout() *GetScimResourcetypeGatewayTimeout {
	return &GetScimResourcetypeGatewayTimeout{}
}

/*GetScimResourcetypeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScimResourcetypeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScimResourcetypeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/resourcetypes/{resourceType}][%d] getScimResourcetypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimResourcetypeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimResourcetypeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
