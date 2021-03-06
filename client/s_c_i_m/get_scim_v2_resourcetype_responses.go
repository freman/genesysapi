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

// GetScimV2ResourcetypeReader is a Reader for the GetScimV2Resourcetype structure.
type GetScimV2ResourcetypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimV2ResourcetypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimV2ResourcetypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScimV2ResourcetypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimV2ResourcetypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimV2ResourcetypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimV2ResourcetypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimV2ResourcetypeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimV2ResourcetypeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimV2ResourcetypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimV2ResourcetypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimV2ResourcetypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimV2ResourcetypeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimV2ResourcetypeOK creates a GetScimV2ResourcetypeOK with default headers values
func NewGetScimV2ResourcetypeOK() *GetScimV2ResourcetypeOK {
	return &GetScimV2ResourcetypeOK{}
}

/*GetScimV2ResourcetypeOK handles this case with default header values.

successful operation
*/
type GetScimV2ResourcetypeOK struct {
	Payload *models.ScimConfigResourceType
}

func (o *GetScimV2ResourcetypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeOK  %+v", 200, o.Payload)
}

func (o *GetScimV2ResourcetypeOK) GetPayload() *models.ScimConfigResourceType {
	return o.Payload
}

func (o *GetScimV2ResourcetypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimConfigResourceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeBadRequest creates a GetScimV2ResourcetypeBadRequest with default headers values
func NewGetScimV2ResourcetypeBadRequest() *GetScimV2ResourcetypeBadRequest {
	return &GetScimV2ResourcetypeBadRequest{}
}

/*GetScimV2ResourcetypeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimV2ResourcetypeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2ResourcetypeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeUnauthorized creates a GetScimV2ResourcetypeUnauthorized with default headers values
func NewGetScimV2ResourcetypeUnauthorized() *GetScimV2ResourcetypeUnauthorized {
	return &GetScimV2ResourcetypeUnauthorized{}
}

/*GetScimV2ResourcetypeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimV2ResourcetypeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2ResourcetypeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeForbidden creates a GetScimV2ResourcetypeForbidden with default headers values
func NewGetScimV2ResourcetypeForbidden() *GetScimV2ResourcetypeForbidden {
	return &GetScimV2ResourcetypeForbidden{}
}

/*GetScimV2ResourcetypeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScimV2ResourcetypeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2ResourcetypeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeNotFound creates a GetScimV2ResourcetypeNotFound with default headers values
func NewGetScimV2ResourcetypeNotFound() *GetScimV2ResourcetypeNotFound {
	return &GetScimV2ResourcetypeNotFound{}
}

/*GetScimV2ResourcetypeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScimV2ResourcetypeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2ResourcetypeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeRequestEntityTooLarge creates a GetScimV2ResourcetypeRequestEntityTooLarge with default headers values
func NewGetScimV2ResourcetypeRequestEntityTooLarge() *GetScimV2ResourcetypeRequestEntityTooLarge {
	return &GetScimV2ResourcetypeRequestEntityTooLarge{}
}

/*GetScimV2ResourcetypeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScimV2ResourcetypeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2ResourcetypeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeUnsupportedMediaType creates a GetScimV2ResourcetypeUnsupportedMediaType with default headers values
func NewGetScimV2ResourcetypeUnsupportedMediaType() *GetScimV2ResourcetypeUnsupportedMediaType {
	return &GetScimV2ResourcetypeUnsupportedMediaType{}
}

/*GetScimV2ResourcetypeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimV2ResourcetypeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2ResourcetypeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeTooManyRequests creates a GetScimV2ResourcetypeTooManyRequests with default headers values
func NewGetScimV2ResourcetypeTooManyRequests() *GetScimV2ResourcetypeTooManyRequests {
	return &GetScimV2ResourcetypeTooManyRequests{}
}

/*GetScimV2ResourcetypeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScimV2ResourcetypeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2ResourcetypeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeInternalServerError creates a GetScimV2ResourcetypeInternalServerError with default headers values
func NewGetScimV2ResourcetypeInternalServerError() *GetScimV2ResourcetypeInternalServerError {
	return &GetScimV2ResourcetypeInternalServerError{}
}

/*GetScimV2ResourcetypeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimV2ResourcetypeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2ResourcetypeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeServiceUnavailable creates a GetScimV2ResourcetypeServiceUnavailable with default headers values
func NewGetScimV2ResourcetypeServiceUnavailable() *GetScimV2ResourcetypeServiceUnavailable {
	return &GetScimV2ResourcetypeServiceUnavailable{}
}

/*GetScimV2ResourcetypeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimV2ResourcetypeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2ResourcetypeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypeGatewayTimeout creates a GetScimV2ResourcetypeGatewayTimeout with default headers values
func NewGetScimV2ResourcetypeGatewayTimeout() *GetScimV2ResourcetypeGatewayTimeout {
	return &GetScimV2ResourcetypeGatewayTimeout{}
}

/*GetScimV2ResourcetypeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScimV2ResourcetypeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScimV2ResourcetypeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes/{resourceType}][%d] getScimV2ResourcetypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2ResourcetypeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
