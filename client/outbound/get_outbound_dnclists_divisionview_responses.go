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

// GetOutboundDnclistsDivisionviewReader is a Reader for the GetOutboundDnclistsDivisionview structure.
type GetOutboundDnclistsDivisionviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundDnclistsDivisionviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundDnclistsDivisionviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundDnclistsDivisionviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundDnclistsDivisionviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundDnclistsDivisionviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundDnclistsDivisionviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundDnclistsDivisionviewRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundDnclistsDivisionviewUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundDnclistsDivisionviewTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundDnclistsDivisionviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundDnclistsDivisionviewServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundDnclistsDivisionviewGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundDnclistsDivisionviewOK creates a GetOutboundDnclistsDivisionviewOK with default headers values
func NewGetOutboundDnclistsDivisionviewOK() *GetOutboundDnclistsDivisionviewOK {
	return &GetOutboundDnclistsDivisionviewOK{}
}

/*GetOutboundDnclistsDivisionviewOK handles this case with default header values.

successful operation
*/
type GetOutboundDnclistsDivisionviewOK struct {
	Payload *models.DncListDivisionView
}

func (o *GetOutboundDnclistsDivisionviewOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewOK  %+v", 200, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewOK) GetPayload() *models.DncListDivisionView {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DncListDivisionView)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewBadRequest creates a GetOutboundDnclistsDivisionviewBadRequest with default headers values
func NewGetOutboundDnclistsDivisionviewBadRequest() *GetOutboundDnclistsDivisionviewBadRequest {
	return &GetOutboundDnclistsDivisionviewBadRequest{}
}

/*GetOutboundDnclistsDivisionviewBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundDnclistsDivisionviewBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewUnauthorized creates a GetOutboundDnclistsDivisionviewUnauthorized with default headers values
func NewGetOutboundDnclistsDivisionviewUnauthorized() *GetOutboundDnclistsDivisionviewUnauthorized {
	return &GetOutboundDnclistsDivisionviewUnauthorized{}
}

/*GetOutboundDnclistsDivisionviewUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundDnclistsDivisionviewUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewForbidden creates a GetOutboundDnclistsDivisionviewForbidden with default headers values
func NewGetOutboundDnclistsDivisionviewForbidden() *GetOutboundDnclistsDivisionviewForbidden {
	return &GetOutboundDnclistsDivisionviewForbidden{}
}

/*GetOutboundDnclistsDivisionviewForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundDnclistsDivisionviewForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewNotFound creates a GetOutboundDnclistsDivisionviewNotFound with default headers values
func NewGetOutboundDnclistsDivisionviewNotFound() *GetOutboundDnclistsDivisionviewNotFound {
	return &GetOutboundDnclistsDivisionviewNotFound{}
}

/*GetOutboundDnclistsDivisionviewNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundDnclistsDivisionviewNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewRequestEntityTooLarge creates a GetOutboundDnclistsDivisionviewRequestEntityTooLarge with default headers values
func NewGetOutboundDnclistsDivisionviewRequestEntityTooLarge() *GetOutboundDnclistsDivisionviewRequestEntityTooLarge {
	return &GetOutboundDnclistsDivisionviewRequestEntityTooLarge{}
}

/*GetOutboundDnclistsDivisionviewRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundDnclistsDivisionviewRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewUnsupportedMediaType creates a GetOutboundDnclistsDivisionviewUnsupportedMediaType with default headers values
func NewGetOutboundDnclistsDivisionviewUnsupportedMediaType() *GetOutboundDnclistsDivisionviewUnsupportedMediaType {
	return &GetOutboundDnclistsDivisionviewUnsupportedMediaType{}
}

/*GetOutboundDnclistsDivisionviewUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundDnclistsDivisionviewUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewTooManyRequests creates a GetOutboundDnclistsDivisionviewTooManyRequests with default headers values
func NewGetOutboundDnclistsDivisionviewTooManyRequests() *GetOutboundDnclistsDivisionviewTooManyRequests {
	return &GetOutboundDnclistsDivisionviewTooManyRequests{}
}

/*GetOutboundDnclistsDivisionviewTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundDnclistsDivisionviewTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewInternalServerError creates a GetOutboundDnclistsDivisionviewInternalServerError with default headers values
func NewGetOutboundDnclistsDivisionviewInternalServerError() *GetOutboundDnclistsDivisionviewInternalServerError {
	return &GetOutboundDnclistsDivisionviewInternalServerError{}
}

/*GetOutboundDnclistsDivisionviewInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundDnclistsDivisionviewInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewServiceUnavailable creates a GetOutboundDnclistsDivisionviewServiceUnavailable with default headers values
func NewGetOutboundDnclistsDivisionviewServiceUnavailable() *GetOutboundDnclistsDivisionviewServiceUnavailable {
	return &GetOutboundDnclistsDivisionviewServiceUnavailable{}
}

/*GetOutboundDnclistsDivisionviewServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundDnclistsDivisionviewServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundDnclistsDivisionviewGatewayTimeout creates a GetOutboundDnclistsDivisionviewGatewayTimeout with default headers values
func NewGetOutboundDnclistsDivisionviewGatewayTimeout() *GetOutboundDnclistsDivisionviewGatewayTimeout {
	return &GetOutboundDnclistsDivisionviewGatewayTimeout{}
}

/*GetOutboundDnclistsDivisionviewGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundDnclistsDivisionviewGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundDnclistsDivisionviewGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/dnclists/divisionviews/{dncListId}][%d] getOutboundDnclistsDivisionviewGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundDnclistsDivisionviewGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundDnclistsDivisionviewGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}