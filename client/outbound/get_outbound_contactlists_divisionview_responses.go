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

// GetOutboundContactlistsDivisionviewReader is a Reader for the GetOutboundContactlistsDivisionview structure.
type GetOutboundContactlistsDivisionviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundContactlistsDivisionviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundContactlistsDivisionviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundContactlistsDivisionviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundContactlistsDivisionviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundContactlistsDivisionviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundContactlistsDivisionviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundContactlistsDivisionviewRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundContactlistsDivisionviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundContactlistsDivisionviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundContactlistsDivisionviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundContactlistsDivisionviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundContactlistsDivisionviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundContactlistsDivisionviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundContactlistsDivisionviewOK creates a GetOutboundContactlistsDivisionviewOK with default headers values
func NewGetOutboundContactlistsDivisionviewOK() *GetOutboundContactlistsDivisionviewOK {
	return &GetOutboundContactlistsDivisionviewOK{}
}

/*GetOutboundContactlistsDivisionviewOK handles this case with default header values.

successful operation
*/
type GetOutboundContactlistsDivisionviewOK struct {
	Payload *models.ContactListDivisionView
}

func (o *GetOutboundContactlistsDivisionviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewOK) GetPayload() *models.ContactListDivisionView {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactListDivisionView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewBadRequest creates a GetOutboundContactlistsDivisionviewBadRequest with default headers values
func NewGetOutboundContactlistsDivisionviewBadRequest() *GetOutboundContactlistsDivisionviewBadRequest {
	return &GetOutboundContactlistsDivisionviewBadRequest{}
}

/*GetOutboundContactlistsDivisionviewBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundContactlistsDivisionviewBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewUnauthorized creates a GetOutboundContactlistsDivisionviewUnauthorized with default headers values
func NewGetOutboundContactlistsDivisionviewUnauthorized() *GetOutboundContactlistsDivisionviewUnauthorized {
	return &GetOutboundContactlistsDivisionviewUnauthorized{}
}

/*GetOutboundContactlistsDivisionviewUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundContactlistsDivisionviewUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewForbidden creates a GetOutboundContactlistsDivisionviewForbidden with default headers values
func NewGetOutboundContactlistsDivisionviewForbidden() *GetOutboundContactlistsDivisionviewForbidden {
	return &GetOutboundContactlistsDivisionviewForbidden{}
}

/*GetOutboundContactlistsDivisionviewForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundContactlistsDivisionviewForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewNotFound creates a GetOutboundContactlistsDivisionviewNotFound with default headers values
func NewGetOutboundContactlistsDivisionviewNotFound() *GetOutboundContactlistsDivisionviewNotFound {
	return &GetOutboundContactlistsDivisionviewNotFound{}
}

/*GetOutboundContactlistsDivisionviewNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundContactlistsDivisionviewNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewRequestTimeout creates a GetOutboundContactlistsDivisionviewRequestTimeout with default headers values
func NewGetOutboundContactlistsDivisionviewRequestTimeout() *GetOutboundContactlistsDivisionviewRequestTimeout {
	return &GetOutboundContactlistsDivisionviewRequestTimeout{}
}

/*GetOutboundContactlistsDivisionviewRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundContactlistsDivisionviewRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewRequestEntityTooLarge creates a GetOutboundContactlistsDivisionviewRequestEntityTooLarge with default headers values
func NewGetOutboundContactlistsDivisionviewRequestEntityTooLarge() *GetOutboundContactlistsDivisionviewRequestEntityTooLarge {
	return &GetOutboundContactlistsDivisionviewRequestEntityTooLarge{}
}

/*GetOutboundContactlistsDivisionviewRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundContactlistsDivisionviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewUnsupportedMediaType creates a GetOutboundContactlistsDivisionviewUnsupportedMediaType with default headers values
func NewGetOutboundContactlistsDivisionviewUnsupportedMediaType() *GetOutboundContactlistsDivisionviewUnsupportedMediaType {
	return &GetOutboundContactlistsDivisionviewUnsupportedMediaType{}
}

/*GetOutboundContactlistsDivisionviewUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundContactlistsDivisionviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewTooManyRequests creates a GetOutboundContactlistsDivisionviewTooManyRequests with default headers values
func NewGetOutboundContactlistsDivisionviewTooManyRequests() *GetOutboundContactlistsDivisionviewTooManyRequests {
	return &GetOutboundContactlistsDivisionviewTooManyRequests{}
}

/*GetOutboundContactlistsDivisionviewTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundContactlistsDivisionviewTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewInternalServerError creates a GetOutboundContactlistsDivisionviewInternalServerError with default headers values
func NewGetOutboundContactlistsDivisionviewInternalServerError() *GetOutboundContactlistsDivisionviewInternalServerError {
	return &GetOutboundContactlistsDivisionviewInternalServerError{}
}

/*GetOutboundContactlistsDivisionviewInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundContactlistsDivisionviewInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewServiceUnavailable creates a GetOutboundContactlistsDivisionviewServiceUnavailable with default headers values
func NewGetOutboundContactlistsDivisionviewServiceUnavailable() *GetOutboundContactlistsDivisionviewServiceUnavailable {
	return &GetOutboundContactlistsDivisionviewServiceUnavailable{}
}

/*GetOutboundContactlistsDivisionviewServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundContactlistsDivisionviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundContactlistsDivisionviewGatewayTimeout creates a GetOutboundContactlistsDivisionviewGatewayTimeout with default headers values
func NewGetOutboundContactlistsDivisionviewGatewayTimeout() *GetOutboundContactlistsDivisionviewGatewayTimeout {
	return &GetOutboundContactlistsDivisionviewGatewayTimeout{}
}

/*GetOutboundContactlistsDivisionviewGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundContactlistsDivisionviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundContactlistsDivisionviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/contactlists/divisionviews/{contactListId}][%d] getOutboundContactlistsDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundContactlistsDivisionviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundContactlistsDivisionviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
