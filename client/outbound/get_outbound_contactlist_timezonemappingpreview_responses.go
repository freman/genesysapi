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

// GetOutboundContactlistTimezonemappingpreviewReader is a Reader for the GetOutboundContactlistTimezonemappingpreview structure.
type GetOutboundContactlistTimezonemappingpreviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundContactlistTimezonemappingpreviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundContactlistTimezonemappingpreviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundContactlistTimezonemappingpreviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundContactlistTimezonemappingpreviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundContactlistTimezonemappingpreviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundContactlistTimezonemappingpreviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundContactlistTimezonemappingpreviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundContactlistTimezonemappingpreviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundContactlistTimezonemappingpreviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundContactlistTimezonemappingpreviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundContactlistTimezonemappingpreviewOK creates a GetOutboundContactlistTimezonemappingpreviewOK with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewOK() *GetOutboundContactlistTimezonemappingpreviewOK {
	return &GetOutboundContactlistTimezonemappingpreviewOK{}
}

/*GetOutboundContactlistTimezonemappingpreviewOK handles this case with default header values.

successful operation
*/
type GetOutboundContactlistTimezonemappingpreviewOK struct {
	Payload *models.TimeZoneMappingPreview
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) GetPayload() *models.TimeZoneMappingPreview {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TimeZoneMappingPreview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewBadRequest creates a GetOutboundContactlistTimezonemappingpreviewBadRequest with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewBadRequest() *GetOutboundContactlistTimezonemappingpreviewBadRequest {
	return &GetOutboundContactlistTimezonemappingpreviewBadRequest{}
}

/*GetOutboundContactlistTimezonemappingpreviewBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundContactlistTimezonemappingpreviewBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewUnauthorized creates a GetOutboundContactlistTimezonemappingpreviewUnauthorized with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewUnauthorized() *GetOutboundContactlistTimezonemappingpreviewUnauthorized {
	return &GetOutboundContactlistTimezonemappingpreviewUnauthorized{}
}

/*GetOutboundContactlistTimezonemappingpreviewUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundContactlistTimezonemappingpreviewUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewForbidden creates a GetOutboundContactlistTimezonemappingpreviewForbidden with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewForbidden() *GetOutboundContactlistTimezonemappingpreviewForbidden {
	return &GetOutboundContactlistTimezonemappingpreviewForbidden{}
}

/*GetOutboundContactlistTimezonemappingpreviewForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundContactlistTimezonemappingpreviewForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewNotFound creates a GetOutboundContactlistTimezonemappingpreviewNotFound with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewNotFound() *GetOutboundContactlistTimezonemappingpreviewNotFound {
	return &GetOutboundContactlistTimezonemappingpreviewNotFound{}
}

/*GetOutboundContactlistTimezonemappingpreviewNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundContactlistTimezonemappingpreviewNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge creates a GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge() *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge {
	return &GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge{}
}

/*GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType creates a GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType() *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType {
	return &GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType{}
}

/*GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewTooManyRequests creates a GetOutboundContactlistTimezonemappingpreviewTooManyRequests with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewTooManyRequests() *GetOutboundContactlistTimezonemappingpreviewTooManyRequests {
	return &GetOutboundContactlistTimezonemappingpreviewTooManyRequests{}
}

/*GetOutboundContactlistTimezonemappingpreviewTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundContactlistTimezonemappingpreviewTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewInternalServerError creates a GetOutboundContactlistTimezonemappingpreviewInternalServerError with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewInternalServerError() *GetOutboundContactlistTimezonemappingpreviewInternalServerError {
	return &GetOutboundContactlistTimezonemappingpreviewInternalServerError{}
}

/*GetOutboundContactlistTimezonemappingpreviewInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundContactlistTimezonemappingpreviewInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewServiceUnavailable creates a GetOutboundContactlistTimezonemappingpreviewServiceUnavailable with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewServiceUnavailable() *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable {
	return &GetOutboundContactlistTimezonemappingpreviewServiceUnavailable{}
}

/*GetOutboundContactlistTimezonemappingpreviewServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundContactlistTimezonemappingpreviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistTimezonemappingpreviewGatewayTimeout creates a GetOutboundContactlistTimezonemappingpreviewGatewayTimeout with default headers values
func NewGetOutboundContactlistTimezonemappingpreviewGatewayTimeout() *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout {
	return &GetOutboundContactlistTimezonemappingpreviewGatewayTimeout{}
}

/*GetOutboundContactlistTimezonemappingpreviewGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundContactlistTimezonemappingpreviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/{contactListId}/timezonemappingpreview][%d] getOutboundContactlistTimezonemappingpreviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistTimezonemappingpreviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
