// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundDnclistsDivisionviewsReader is a Reader for the GetOutboundDnclistsDivisionviews structure.
type GetOutboundDnclistsDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundDnclistsDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundDnclistsDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundDnclistsDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundDnclistsDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundDnclistsDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundDnclistsDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundDnclistsDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundDnclistsDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundDnclistsDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundDnclistsDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundDnclistsDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundDnclistsDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundDnclistsDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundDnclistsDivisionviewsOK creates a GetOutboundDnclistsDivisionviewsOK with default headers values
func NewGetOutboundDnclistsDivisionviewsOK() *GetOutboundDnclistsDivisionviewsOK {
	return &GetOutboundDnclistsDivisionviewsOK{}
}

/*GetOutboundDnclistsDivisionviewsOK handles this case with default header values.

successful operation
*/
type GetOutboundDnclistsDivisionviewsOK struct {
	Payload *models.DncListDivisionViewListing
}

func (o *GetOutboundDnclistsDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsOK) GetPayload() *models.DncListDivisionViewListing {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DncListDivisionViewListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsBadRequest creates a GetOutboundDnclistsDivisionviewsBadRequest with default headers values
func NewGetOutboundDnclistsDivisionviewsBadRequest() *GetOutboundDnclistsDivisionviewsBadRequest {
	return &GetOutboundDnclistsDivisionviewsBadRequest{}
}

/*GetOutboundDnclistsDivisionviewsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundDnclistsDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsUnauthorized creates a GetOutboundDnclistsDivisionviewsUnauthorized with default headers values
func NewGetOutboundDnclistsDivisionviewsUnauthorized() *GetOutboundDnclistsDivisionviewsUnauthorized {
	return &GetOutboundDnclistsDivisionviewsUnauthorized{}
}

/*GetOutboundDnclistsDivisionviewsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundDnclistsDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsForbidden creates a GetOutboundDnclistsDivisionviewsForbidden with default headers values
func NewGetOutboundDnclistsDivisionviewsForbidden() *GetOutboundDnclistsDivisionviewsForbidden {
	return &GetOutboundDnclistsDivisionviewsForbidden{}
}

/*GetOutboundDnclistsDivisionviewsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundDnclistsDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsNotFound creates a GetOutboundDnclistsDivisionviewsNotFound with default headers values
func NewGetOutboundDnclistsDivisionviewsNotFound() *GetOutboundDnclistsDivisionviewsNotFound {
	return &GetOutboundDnclistsDivisionviewsNotFound{}
}

/*GetOutboundDnclistsDivisionviewsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundDnclistsDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsRequestTimeout creates a GetOutboundDnclistsDivisionviewsRequestTimeout with default headers values
func NewGetOutboundDnclistsDivisionviewsRequestTimeout() *GetOutboundDnclistsDivisionviewsRequestTimeout {
	return &GetOutboundDnclistsDivisionviewsRequestTimeout{}
}

/*GetOutboundDnclistsDivisionviewsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundDnclistsDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsRequestEntityTooLarge creates a GetOutboundDnclistsDivisionviewsRequestEntityTooLarge with default headers values
func NewGetOutboundDnclistsDivisionviewsRequestEntityTooLarge() *GetOutboundDnclistsDivisionviewsRequestEntityTooLarge {
	return &GetOutboundDnclistsDivisionviewsRequestEntityTooLarge{}
}

/*GetOutboundDnclistsDivisionviewsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOutboundDnclistsDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsUnsupportedMediaType creates a GetOutboundDnclistsDivisionviewsUnsupportedMediaType with default headers values
func NewGetOutboundDnclistsDivisionviewsUnsupportedMediaType() *GetOutboundDnclistsDivisionviewsUnsupportedMediaType {
	return &GetOutboundDnclistsDivisionviewsUnsupportedMediaType{}
}

/*GetOutboundDnclistsDivisionviewsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundDnclistsDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsTooManyRequests creates a GetOutboundDnclistsDivisionviewsTooManyRequests with default headers values
func NewGetOutboundDnclistsDivisionviewsTooManyRequests() *GetOutboundDnclistsDivisionviewsTooManyRequests {
	return &GetOutboundDnclistsDivisionviewsTooManyRequests{}
}

/*GetOutboundDnclistsDivisionviewsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundDnclistsDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsInternalServerError creates a GetOutboundDnclistsDivisionviewsInternalServerError with default headers values
func NewGetOutboundDnclistsDivisionviewsInternalServerError() *GetOutboundDnclistsDivisionviewsInternalServerError {
	return &GetOutboundDnclistsDivisionviewsInternalServerError{}
}

/*GetOutboundDnclistsDivisionviewsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundDnclistsDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsServiceUnavailable creates a GetOutboundDnclistsDivisionviewsServiceUnavailable with default headers values
func NewGetOutboundDnclistsDivisionviewsServiceUnavailable() *GetOutboundDnclistsDivisionviewsServiceUnavailable {
	return &GetOutboundDnclistsDivisionviewsServiceUnavailable{}
}

/*GetOutboundDnclistsDivisionviewsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundDnclistsDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewsGatewayTimeout creates a GetOutboundDnclistsDivisionviewsGatewayTimeout with default headers values
func NewGetOutboundDnclistsDivisionviewsGatewayTimeout() *GetOutboundDnclistsDivisionviewsGatewayTimeout {
	return &GetOutboundDnclistsDivisionviewsGatewayTimeout{}
}

/*GetOutboundDnclistsDivisionviewsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundDnclistsDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews][%d] getOutboundDnclistsDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
