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

// GetAuthorizationDivisionGrantsReader is a Reader for the GetAuthorizationDivisionGrants structure.
type GetAuthorizationDivisionGrantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationDivisionGrantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationDivisionGrantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationDivisionGrantsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationDivisionGrantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationDivisionGrantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationDivisionGrantsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationDivisionGrantsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationDivisionGrantsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationDivisionGrantsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationDivisionGrantsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationDivisionGrantsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationDivisionGrantsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationDivisionGrantsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationDivisionGrantsOK creates a GetAuthorizationDivisionGrantsOK with default headers values
func NewGetAuthorizationDivisionGrantsOK() *GetAuthorizationDivisionGrantsOK {
	return &GetAuthorizationDivisionGrantsOK{}
}

/*GetAuthorizationDivisionGrantsOK handles this case with default header values.

successful operation
*/
type GetAuthorizationDivisionGrantsOK struct {
	Payload *models.AuthzDivisionGrantEntityListing
}

func (o *GetAuthorizationDivisionGrantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsOK) GetPayload() *models.AuthzDivisionGrantEntityListing {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthzDivisionGrantEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsBadRequest creates a GetAuthorizationDivisionGrantsBadRequest with default headers values
func NewGetAuthorizationDivisionGrantsBadRequest() *GetAuthorizationDivisionGrantsBadRequest {
	return &GetAuthorizationDivisionGrantsBadRequest{}
}

/*GetAuthorizationDivisionGrantsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationDivisionGrantsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsUnauthorized creates a GetAuthorizationDivisionGrantsUnauthorized with default headers values
func NewGetAuthorizationDivisionGrantsUnauthorized() *GetAuthorizationDivisionGrantsUnauthorized {
	return &GetAuthorizationDivisionGrantsUnauthorized{}
}

/*GetAuthorizationDivisionGrantsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationDivisionGrantsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsForbidden creates a GetAuthorizationDivisionGrantsForbidden with default headers values
func NewGetAuthorizationDivisionGrantsForbidden() *GetAuthorizationDivisionGrantsForbidden {
	return &GetAuthorizationDivisionGrantsForbidden{}
}

/*GetAuthorizationDivisionGrantsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationDivisionGrantsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsNotFound creates a GetAuthorizationDivisionGrantsNotFound with default headers values
func NewGetAuthorizationDivisionGrantsNotFound() *GetAuthorizationDivisionGrantsNotFound {
	return &GetAuthorizationDivisionGrantsNotFound{}
}

/*GetAuthorizationDivisionGrantsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationDivisionGrantsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsRequestTimeout creates a GetAuthorizationDivisionGrantsRequestTimeout with default headers values
func NewGetAuthorizationDivisionGrantsRequestTimeout() *GetAuthorizationDivisionGrantsRequestTimeout {
	return &GetAuthorizationDivisionGrantsRequestTimeout{}
}

/*GetAuthorizationDivisionGrantsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationDivisionGrantsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsRequestEntityTooLarge creates a GetAuthorizationDivisionGrantsRequestEntityTooLarge with default headers values
func NewGetAuthorizationDivisionGrantsRequestEntityTooLarge() *GetAuthorizationDivisionGrantsRequestEntityTooLarge {
	return &GetAuthorizationDivisionGrantsRequestEntityTooLarge{}
}

/*GetAuthorizationDivisionGrantsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAuthorizationDivisionGrantsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsUnsupportedMediaType creates a GetAuthorizationDivisionGrantsUnsupportedMediaType with default headers values
func NewGetAuthorizationDivisionGrantsUnsupportedMediaType() *GetAuthorizationDivisionGrantsUnsupportedMediaType {
	return &GetAuthorizationDivisionGrantsUnsupportedMediaType{}
}

/*GetAuthorizationDivisionGrantsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationDivisionGrantsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsTooManyRequests creates a GetAuthorizationDivisionGrantsTooManyRequests with default headers values
func NewGetAuthorizationDivisionGrantsTooManyRequests() *GetAuthorizationDivisionGrantsTooManyRequests {
	return &GetAuthorizationDivisionGrantsTooManyRequests{}
}

/*GetAuthorizationDivisionGrantsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationDivisionGrantsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsInternalServerError creates a GetAuthorizationDivisionGrantsInternalServerError with default headers values
func NewGetAuthorizationDivisionGrantsInternalServerError() *GetAuthorizationDivisionGrantsInternalServerError {
	return &GetAuthorizationDivisionGrantsInternalServerError{}
}

/*GetAuthorizationDivisionGrantsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationDivisionGrantsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsServiceUnavailable creates a GetAuthorizationDivisionGrantsServiceUnavailable with default headers values
func NewGetAuthorizationDivisionGrantsServiceUnavailable() *GetAuthorizationDivisionGrantsServiceUnavailable {
	return &GetAuthorizationDivisionGrantsServiceUnavailable{}
}

/*GetAuthorizationDivisionGrantsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationDivisionGrantsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationDivisionGrantsGatewayTimeout creates a GetAuthorizationDivisionGrantsGatewayTimeout with default headers values
func NewGetAuthorizationDivisionGrantsGatewayTimeout() *GetAuthorizationDivisionGrantsGatewayTimeout {
	return &GetAuthorizationDivisionGrantsGatewayTimeout{}
}

/*GetAuthorizationDivisionGrantsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationDivisionGrantsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationDivisionGrantsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/divisions/{divisionId}/grants][%d] getAuthorizationDivisionGrantsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationDivisionGrantsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationDivisionGrantsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
