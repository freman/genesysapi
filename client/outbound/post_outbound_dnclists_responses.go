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

// PostOutboundDnclistsReader is a Reader for the PostOutboundDnclists structure.
type PostOutboundDnclistsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundDnclistsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundDnclistsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundDnclistsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundDnclistsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundDnclistsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundDnclistsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundDnclistsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundDnclistsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundDnclistsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundDnclistsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundDnclistsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundDnclistsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundDnclistsOK creates a PostOutboundDnclistsOK with default headers values
func NewPostOutboundDnclistsOK() *PostOutboundDnclistsOK {
	return &PostOutboundDnclistsOK{}
}

/*PostOutboundDnclistsOK handles this case with default header values.

successful operation
*/
type PostOutboundDnclistsOK struct {
	Payload *models.DncList
}

func (o *PostOutboundDnclistsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundDnclistsOK) GetPayload() *models.DncList {
	return o.Payload
}

func (o *PostOutboundDnclistsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DncList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsBadRequest creates a PostOutboundDnclistsBadRequest with default headers values
func NewPostOutboundDnclistsBadRequest() *PostOutboundDnclistsBadRequest {
	return &PostOutboundDnclistsBadRequest{}
}

/*PostOutboundDnclistsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundDnclistsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundDnclistsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsUnauthorized creates a PostOutboundDnclistsUnauthorized with default headers values
func NewPostOutboundDnclistsUnauthorized() *PostOutboundDnclistsUnauthorized {
	return &PostOutboundDnclistsUnauthorized{}
}

/*PostOutboundDnclistsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundDnclistsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundDnclistsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsForbidden creates a PostOutboundDnclistsForbidden with default headers values
func NewPostOutboundDnclistsForbidden() *PostOutboundDnclistsForbidden {
	return &PostOutboundDnclistsForbidden{}
}

/*PostOutboundDnclistsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundDnclistsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundDnclistsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsNotFound creates a PostOutboundDnclistsNotFound with default headers values
func NewPostOutboundDnclistsNotFound() *PostOutboundDnclistsNotFound {
	return &PostOutboundDnclistsNotFound{}
}

/*PostOutboundDnclistsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundDnclistsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundDnclistsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsRequestEntityTooLarge creates a PostOutboundDnclistsRequestEntityTooLarge with default headers values
func NewPostOutboundDnclistsRequestEntityTooLarge() *PostOutboundDnclistsRequestEntityTooLarge {
	return &PostOutboundDnclistsRequestEntityTooLarge{}
}

/*PostOutboundDnclistsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundDnclistsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsUnsupportedMediaType creates a PostOutboundDnclistsUnsupportedMediaType with default headers values
func NewPostOutboundDnclistsUnsupportedMediaType() *PostOutboundDnclistsUnsupportedMediaType {
	return &PostOutboundDnclistsUnsupportedMediaType{}
}

/*PostOutboundDnclistsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundDnclistsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundDnclistsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsTooManyRequests creates a PostOutboundDnclistsTooManyRequests with default headers values
func NewPostOutboundDnclistsTooManyRequests() *PostOutboundDnclistsTooManyRequests {
	return &PostOutboundDnclistsTooManyRequests{}
}

/*PostOutboundDnclistsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOutboundDnclistsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundDnclistsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsInternalServerError creates a PostOutboundDnclistsInternalServerError with default headers values
func NewPostOutboundDnclistsInternalServerError() *PostOutboundDnclistsInternalServerError {
	return &PostOutboundDnclistsInternalServerError{}
}

/*PostOutboundDnclistsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundDnclistsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundDnclistsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsServiceUnavailable creates a PostOutboundDnclistsServiceUnavailable with default headers values
func NewPostOutboundDnclistsServiceUnavailable() *PostOutboundDnclistsServiceUnavailable {
	return &PostOutboundDnclistsServiceUnavailable{}
}

/*PostOutboundDnclistsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundDnclistsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundDnclistsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundDnclistsGatewayTimeout creates a PostOutboundDnclistsGatewayTimeout with default headers values
func NewPostOutboundDnclistsGatewayTimeout() *PostOutboundDnclistsGatewayTimeout {
	return &PostOutboundDnclistsGatewayTimeout{}
}

/*PostOutboundDnclistsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundDnclistsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundDnclistsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/dnclists][%d] postOutboundDnclistsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundDnclistsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundDnclistsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}