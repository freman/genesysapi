// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetContentmanagementShareReader is a Reader for the GetContentmanagementShare structure.
type GetContentmanagementShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementShareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementShareBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementShareUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementShareForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementShareNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementShareRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementShareUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementShareTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementShareInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementShareServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementShareGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementShareOK creates a GetContentmanagementShareOK with default headers values
func NewGetContentmanagementShareOK() *GetContentmanagementShareOK {
	return &GetContentmanagementShareOK{}
}

/*GetContentmanagementShareOK handles this case with default header values.

successful operation
*/
type GetContentmanagementShareOK struct {
	Payload *models.Share
}

func (o *GetContentmanagementShareOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementShareOK) GetPayload() *models.Share {
	return o.Payload
}

func (o *GetContentmanagementShareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Share)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareBadRequest creates a GetContentmanagementShareBadRequest with default headers values
func NewGetContentmanagementShareBadRequest() *GetContentmanagementShareBadRequest {
	return &GetContentmanagementShareBadRequest{}
}

/*GetContentmanagementShareBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementShareBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementShareBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareUnauthorized creates a GetContentmanagementShareUnauthorized with default headers values
func NewGetContentmanagementShareUnauthorized() *GetContentmanagementShareUnauthorized {
	return &GetContentmanagementShareUnauthorized{}
}

/*GetContentmanagementShareUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementShareUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementShareUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareForbidden creates a GetContentmanagementShareForbidden with default headers values
func NewGetContentmanagementShareForbidden() *GetContentmanagementShareForbidden {
	return &GetContentmanagementShareForbidden{}
}

/*GetContentmanagementShareForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementShareForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementShareForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareNotFound creates a GetContentmanagementShareNotFound with default headers values
func NewGetContentmanagementShareNotFound() *GetContentmanagementShareNotFound {
	return &GetContentmanagementShareNotFound{}
}

/*GetContentmanagementShareNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetContentmanagementShareNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementShareNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareRequestEntityTooLarge creates a GetContentmanagementShareRequestEntityTooLarge with default headers values
func NewGetContentmanagementShareRequestEntityTooLarge() *GetContentmanagementShareRequestEntityTooLarge {
	return &GetContentmanagementShareRequestEntityTooLarge{}
}

/*GetContentmanagementShareRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetContentmanagementShareRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementShareRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareUnsupportedMediaType creates a GetContentmanagementShareUnsupportedMediaType with default headers values
func NewGetContentmanagementShareUnsupportedMediaType() *GetContentmanagementShareUnsupportedMediaType {
	return &GetContentmanagementShareUnsupportedMediaType{}
}

/*GetContentmanagementShareUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementShareUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementShareUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareTooManyRequests creates a GetContentmanagementShareTooManyRequests with default headers values
func NewGetContentmanagementShareTooManyRequests() *GetContentmanagementShareTooManyRequests {
	return &GetContentmanagementShareTooManyRequests{}
}

/*GetContentmanagementShareTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetContentmanagementShareTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementShareTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareInternalServerError creates a GetContentmanagementShareInternalServerError with default headers values
func NewGetContentmanagementShareInternalServerError() *GetContentmanagementShareInternalServerError {
	return &GetContentmanagementShareInternalServerError{}
}

/*GetContentmanagementShareInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementShareInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementShareInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareServiceUnavailable creates a GetContentmanagementShareServiceUnavailable with default headers values
func NewGetContentmanagementShareServiceUnavailable() *GetContentmanagementShareServiceUnavailable {
	return &GetContentmanagementShareServiceUnavailable{}
}

/*GetContentmanagementShareServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementShareServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementShareServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementShareGatewayTimeout creates a GetContentmanagementShareGatewayTimeout with default headers values
func NewGetContentmanagementShareGatewayTimeout() *GetContentmanagementShareGatewayTimeout {
	return &GetContentmanagementShareGatewayTimeout{}
}

/*GetContentmanagementShareGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetContentmanagementShareGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetContentmanagementShareGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/shares/{shareId}][%d] getContentmanagementShareGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementShareGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementShareGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
