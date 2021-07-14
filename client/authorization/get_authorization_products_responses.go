// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuthorizationProductsReader is a Reader for the GetAuthorizationProducts structure.
type GetAuthorizationProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationProductsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationProductsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationProductsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationProductsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationProductsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationProductsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationProductsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationProductsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationProductsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationProductsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationProductsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationProductsOK creates a GetAuthorizationProductsOK with default headers values
func NewGetAuthorizationProductsOK() *GetAuthorizationProductsOK {
	return &GetAuthorizationProductsOK{}
}

/*GetAuthorizationProductsOK handles this case with default header values.

successful operation
*/
type GetAuthorizationProductsOK struct {
	Payload *models.OrganizationProductEntityListing
}

func (o *GetAuthorizationProductsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationProductsOK) GetPayload() *models.OrganizationProductEntityListing {
	return o.Payload
}

func (o *GetAuthorizationProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationProductEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsBadRequest creates a GetAuthorizationProductsBadRequest with default headers values
func NewGetAuthorizationProductsBadRequest() *GetAuthorizationProductsBadRequest {
	return &GetAuthorizationProductsBadRequest{}
}

/*GetAuthorizationProductsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationProductsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationProductsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsUnauthorized creates a GetAuthorizationProductsUnauthorized with default headers values
func NewGetAuthorizationProductsUnauthorized() *GetAuthorizationProductsUnauthorized {
	return &GetAuthorizationProductsUnauthorized{}
}

/*GetAuthorizationProductsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationProductsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationProductsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsForbidden creates a GetAuthorizationProductsForbidden with default headers values
func NewGetAuthorizationProductsForbidden() *GetAuthorizationProductsForbidden {
	return &GetAuthorizationProductsForbidden{}
}

/*GetAuthorizationProductsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationProductsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationProductsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsNotFound creates a GetAuthorizationProductsNotFound with default headers values
func NewGetAuthorizationProductsNotFound() *GetAuthorizationProductsNotFound {
	return &GetAuthorizationProductsNotFound{}
}

/*GetAuthorizationProductsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationProductsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationProductsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsRequestTimeout creates a GetAuthorizationProductsRequestTimeout with default headers values
func NewGetAuthorizationProductsRequestTimeout() *GetAuthorizationProductsRequestTimeout {
	return &GetAuthorizationProductsRequestTimeout{}
}

/*GetAuthorizationProductsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationProductsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationProductsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsRequestEntityTooLarge creates a GetAuthorizationProductsRequestEntityTooLarge with default headers values
func NewGetAuthorizationProductsRequestEntityTooLarge() *GetAuthorizationProductsRequestEntityTooLarge {
	return &GetAuthorizationProductsRequestEntityTooLarge{}
}

/*GetAuthorizationProductsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAuthorizationProductsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsUnsupportedMediaType creates a GetAuthorizationProductsUnsupportedMediaType with default headers values
func NewGetAuthorizationProductsUnsupportedMediaType() *GetAuthorizationProductsUnsupportedMediaType {
	return &GetAuthorizationProductsUnsupportedMediaType{}
}

/*GetAuthorizationProductsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationProductsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationProductsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsTooManyRequests creates a GetAuthorizationProductsTooManyRequests with default headers values
func NewGetAuthorizationProductsTooManyRequests() *GetAuthorizationProductsTooManyRequests {
	return &GetAuthorizationProductsTooManyRequests{}
}

/*GetAuthorizationProductsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationProductsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationProductsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsInternalServerError creates a GetAuthorizationProductsInternalServerError with default headers values
func NewGetAuthorizationProductsInternalServerError() *GetAuthorizationProductsInternalServerError {
	return &GetAuthorizationProductsInternalServerError{}
}

/*GetAuthorizationProductsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationProductsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationProductsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsServiceUnavailable creates a GetAuthorizationProductsServiceUnavailable with default headers values
func NewGetAuthorizationProductsServiceUnavailable() *GetAuthorizationProductsServiceUnavailable {
	return &GetAuthorizationProductsServiceUnavailable{}
}

/*GetAuthorizationProductsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationProductsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationProductsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationProductsGatewayTimeout creates a GetAuthorizationProductsGatewayTimeout with default headers values
func NewGetAuthorizationProductsGatewayTimeout() *GetAuthorizationProductsGatewayTimeout {
	return &GetAuthorizationProductsGatewayTimeout{}
}

/*GetAuthorizationProductsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationProductsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationProductsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/products][%d] getAuthorizationProductsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationProductsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationProductsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
