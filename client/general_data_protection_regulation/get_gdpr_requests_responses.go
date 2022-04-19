// Code generated by go-swagger; DO NOT EDIT.

package general_data_protection_regulation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGdprRequestsReader is a Reader for the GetGdprRequests structure.
type GetGdprRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGdprRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGdprRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGdprRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGdprRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGdprRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGdprRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGdprRequestsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGdprRequestsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGdprRequestsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGdprRequestsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGdprRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGdprRequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGdprRequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGdprRequestsOK creates a GetGdprRequestsOK with default headers values
func NewGetGdprRequestsOK() *GetGdprRequestsOK {
	return &GetGdprRequestsOK{}
}

/*GetGdprRequestsOK handles this case with default header values.

successful operation
*/
type GetGdprRequestsOK struct {
	Payload *models.GDPRRequestEntityListing
}

func (o *GetGdprRequestsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsOK  %+v", 200, o.Payload)
}

func (o *GetGdprRequestsOK) GetPayload() *models.GDPRRequestEntityListing {
	return o.Payload
}

func (o *GetGdprRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GDPRRequestEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsBadRequest creates a GetGdprRequestsBadRequest with default headers values
func NewGetGdprRequestsBadRequest() *GetGdprRequestsBadRequest {
	return &GetGdprRequestsBadRequest{}
}

/*GetGdprRequestsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGdprRequestsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGdprRequestsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsUnauthorized creates a GetGdprRequestsUnauthorized with default headers values
func NewGetGdprRequestsUnauthorized() *GetGdprRequestsUnauthorized {
	return &GetGdprRequestsUnauthorized{}
}

/*GetGdprRequestsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGdprRequestsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGdprRequestsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsForbidden creates a GetGdprRequestsForbidden with default headers values
func NewGetGdprRequestsForbidden() *GetGdprRequestsForbidden {
	return &GetGdprRequestsForbidden{}
}

/*GetGdprRequestsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGdprRequestsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsForbidden  %+v", 403, o.Payload)
}

func (o *GetGdprRequestsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsNotFound creates a GetGdprRequestsNotFound with default headers values
func NewGetGdprRequestsNotFound() *GetGdprRequestsNotFound {
	return &GetGdprRequestsNotFound{}
}

/*GetGdprRequestsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGdprRequestsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetGdprRequestsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsRequestTimeout creates a GetGdprRequestsRequestTimeout with default headers values
func NewGetGdprRequestsRequestTimeout() *GetGdprRequestsRequestTimeout {
	return &GetGdprRequestsRequestTimeout{}
}

/*GetGdprRequestsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGdprRequestsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGdprRequestsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsRequestEntityTooLarge creates a GetGdprRequestsRequestEntityTooLarge with default headers values
func NewGetGdprRequestsRequestEntityTooLarge() *GetGdprRequestsRequestEntityTooLarge {
	return &GetGdprRequestsRequestEntityTooLarge{}
}

/*GetGdprRequestsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGdprRequestsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGdprRequestsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsUnsupportedMediaType creates a GetGdprRequestsUnsupportedMediaType with default headers values
func NewGetGdprRequestsUnsupportedMediaType() *GetGdprRequestsUnsupportedMediaType {
	return &GetGdprRequestsUnsupportedMediaType{}
}

/*GetGdprRequestsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGdprRequestsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGdprRequestsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsTooManyRequests creates a GetGdprRequestsTooManyRequests with default headers values
func NewGetGdprRequestsTooManyRequests() *GetGdprRequestsTooManyRequests {
	return &GetGdprRequestsTooManyRequests{}
}

/*GetGdprRequestsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGdprRequestsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGdprRequestsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsInternalServerError creates a GetGdprRequestsInternalServerError with default headers values
func NewGetGdprRequestsInternalServerError() *GetGdprRequestsInternalServerError {
	return &GetGdprRequestsInternalServerError{}
}

/*GetGdprRequestsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGdprRequestsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGdprRequestsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsServiceUnavailable creates a GetGdprRequestsServiceUnavailable with default headers values
func NewGetGdprRequestsServiceUnavailable() *GetGdprRequestsServiceUnavailable {
	return &GetGdprRequestsServiceUnavailable{}
}

/*GetGdprRequestsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGdprRequestsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGdprRequestsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGdprRequestsGatewayTimeout creates a GetGdprRequestsGatewayTimeout with default headers values
func NewGetGdprRequestsGatewayTimeout() *GetGdprRequestsGatewayTimeout {
	return &GetGdprRequestsGatewayTimeout{}
}

/*GetGdprRequestsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGdprRequestsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGdprRequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gdpr/requests][%d] getGdprRequestsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGdprRequestsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGdprRequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
