// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetExternalcontactsReversewhitepageslookupReader is a Reader for the GetExternalcontactsReversewhitepageslookup structure.
type GetExternalcontactsReversewhitepageslookupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsReversewhitepageslookupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsReversewhitepageslookupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsReversewhitepageslookupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsReversewhitepageslookupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsReversewhitepageslookupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsReversewhitepageslookupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsReversewhitepageslookupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsReversewhitepageslookupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsReversewhitepageslookupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsReversewhitepageslookupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsReversewhitepageslookupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsReversewhitepageslookupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsReversewhitepageslookupOK creates a GetExternalcontactsReversewhitepageslookupOK with default headers values
func NewGetExternalcontactsReversewhitepageslookupOK() *GetExternalcontactsReversewhitepageslookupOK {
	return &GetExternalcontactsReversewhitepageslookupOK{}
}

/*GetExternalcontactsReversewhitepageslookupOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsReversewhitepageslookupOK struct {
	Payload *models.ReverseWhitepagesLookupResult
}

func (o *GetExternalcontactsReversewhitepageslookupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupOK) GetPayload() *models.ReverseWhitepagesLookupResult {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReverseWhitepagesLookupResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupBadRequest creates a GetExternalcontactsReversewhitepageslookupBadRequest with default headers values
func NewGetExternalcontactsReversewhitepageslookupBadRequest() *GetExternalcontactsReversewhitepageslookupBadRequest {
	return &GetExternalcontactsReversewhitepageslookupBadRequest{}
}

/*GetExternalcontactsReversewhitepageslookupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsReversewhitepageslookupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupUnauthorized creates a GetExternalcontactsReversewhitepageslookupUnauthorized with default headers values
func NewGetExternalcontactsReversewhitepageslookupUnauthorized() *GetExternalcontactsReversewhitepageslookupUnauthorized {
	return &GetExternalcontactsReversewhitepageslookupUnauthorized{}
}

/*GetExternalcontactsReversewhitepageslookupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsReversewhitepageslookupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupForbidden creates a GetExternalcontactsReversewhitepageslookupForbidden with default headers values
func NewGetExternalcontactsReversewhitepageslookupForbidden() *GetExternalcontactsReversewhitepageslookupForbidden {
	return &GetExternalcontactsReversewhitepageslookupForbidden{}
}

/*GetExternalcontactsReversewhitepageslookupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsReversewhitepageslookupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupNotFound creates a GetExternalcontactsReversewhitepageslookupNotFound with default headers values
func NewGetExternalcontactsReversewhitepageslookupNotFound() *GetExternalcontactsReversewhitepageslookupNotFound {
	return &GetExternalcontactsReversewhitepageslookupNotFound{}
}

/*GetExternalcontactsReversewhitepageslookupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsReversewhitepageslookupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupRequestEntityTooLarge creates a GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge with default headers values
func NewGetExternalcontactsReversewhitepageslookupRequestEntityTooLarge() *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge {
	return &GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge{}
}

/*GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupUnsupportedMediaType creates a GetExternalcontactsReversewhitepageslookupUnsupportedMediaType with default headers values
func NewGetExternalcontactsReversewhitepageslookupUnsupportedMediaType() *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType {
	return &GetExternalcontactsReversewhitepageslookupUnsupportedMediaType{}
}

/*GetExternalcontactsReversewhitepageslookupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsReversewhitepageslookupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupTooManyRequests creates a GetExternalcontactsReversewhitepageslookupTooManyRequests with default headers values
func NewGetExternalcontactsReversewhitepageslookupTooManyRequests() *GetExternalcontactsReversewhitepageslookupTooManyRequests {
	return &GetExternalcontactsReversewhitepageslookupTooManyRequests{}
}

/*GetExternalcontactsReversewhitepageslookupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsReversewhitepageslookupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupInternalServerError creates a GetExternalcontactsReversewhitepageslookupInternalServerError with default headers values
func NewGetExternalcontactsReversewhitepageslookupInternalServerError() *GetExternalcontactsReversewhitepageslookupInternalServerError {
	return &GetExternalcontactsReversewhitepageslookupInternalServerError{}
}

/*GetExternalcontactsReversewhitepageslookupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsReversewhitepageslookupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupServiceUnavailable creates a GetExternalcontactsReversewhitepageslookupServiceUnavailable with default headers values
func NewGetExternalcontactsReversewhitepageslookupServiceUnavailable() *GetExternalcontactsReversewhitepageslookupServiceUnavailable {
	return &GetExternalcontactsReversewhitepageslookupServiceUnavailable{}
}

/*GetExternalcontactsReversewhitepageslookupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsReversewhitepageslookupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsReversewhitepageslookupGatewayTimeout creates a GetExternalcontactsReversewhitepageslookupGatewayTimeout with default headers values
func NewGetExternalcontactsReversewhitepageslookupGatewayTimeout() *GetExternalcontactsReversewhitepageslookupGatewayTimeout {
	return &GetExternalcontactsReversewhitepageslookupGatewayTimeout{}
}

/*GetExternalcontactsReversewhitepageslookupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsReversewhitepageslookupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/reversewhitepageslookup][%d] getExternalcontactsReversewhitepageslookupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsReversewhitepageslookupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
