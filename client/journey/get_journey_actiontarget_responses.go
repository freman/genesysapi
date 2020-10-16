// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyActiontargetReader is a Reader for the GetJourneyActiontarget structure.
type GetJourneyActiontargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyActiontargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyActiontargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyActiontargetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyActiontargetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyActiontargetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyActiontargetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyActiontargetRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyActiontargetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyActiontargetTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyActiontargetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyActiontargetServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyActiontargetGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyActiontargetOK creates a GetJourneyActiontargetOK with default headers values
func NewGetJourneyActiontargetOK() *GetJourneyActiontargetOK {
	return &GetJourneyActiontargetOK{}
}

/*GetJourneyActiontargetOK handles this case with default header values.

successful operation
*/
type GetJourneyActiontargetOK struct {
	Payload *models.ActionTarget
}

func (o *GetJourneyActiontargetOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActiontargetOK) GetPayload() *models.ActionTarget {
	return o.Payload
}

func (o *GetJourneyActiontargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetBadRequest creates a GetJourneyActiontargetBadRequest with default headers values
func NewGetJourneyActiontargetBadRequest() *GetJourneyActiontargetBadRequest {
	return &GetJourneyActiontargetBadRequest{}
}

/*GetJourneyActiontargetBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyActiontargetBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActiontargetBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetUnauthorized creates a GetJourneyActiontargetUnauthorized with default headers values
func NewGetJourneyActiontargetUnauthorized() *GetJourneyActiontargetUnauthorized {
	return &GetJourneyActiontargetUnauthorized{}
}

/*GetJourneyActiontargetUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyActiontargetUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActiontargetUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetForbidden creates a GetJourneyActiontargetForbidden with default headers values
func NewGetJourneyActiontargetForbidden() *GetJourneyActiontargetForbidden {
	return &GetJourneyActiontargetForbidden{}
}

/*GetJourneyActiontargetForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyActiontargetForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActiontargetForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetNotFound creates a GetJourneyActiontargetNotFound with default headers values
func NewGetJourneyActiontargetNotFound() *GetJourneyActiontargetNotFound {
	return &GetJourneyActiontargetNotFound{}
}

/*GetJourneyActiontargetNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetJourneyActiontargetNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActiontargetNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetRequestEntityTooLarge creates a GetJourneyActiontargetRequestEntityTooLarge with default headers values
func NewGetJourneyActiontargetRequestEntityTooLarge() *GetJourneyActiontargetRequestEntityTooLarge {
	return &GetJourneyActiontargetRequestEntityTooLarge{}
}

/*GetJourneyActiontargetRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetJourneyActiontargetRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetUnsupportedMediaType creates a GetJourneyActiontargetUnsupportedMediaType with default headers values
func NewGetJourneyActiontargetUnsupportedMediaType() *GetJourneyActiontargetUnsupportedMediaType {
	return &GetJourneyActiontargetUnsupportedMediaType{}
}

/*GetJourneyActiontargetUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyActiontargetUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActiontargetUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetTooManyRequests creates a GetJourneyActiontargetTooManyRequests with default headers values
func NewGetJourneyActiontargetTooManyRequests() *GetJourneyActiontargetTooManyRequests {
	return &GetJourneyActiontargetTooManyRequests{}
}

/*GetJourneyActiontargetTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetJourneyActiontargetTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActiontargetTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetInternalServerError creates a GetJourneyActiontargetInternalServerError with default headers values
func NewGetJourneyActiontargetInternalServerError() *GetJourneyActiontargetInternalServerError {
	return &GetJourneyActiontargetInternalServerError{}
}

/*GetJourneyActiontargetInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyActiontargetInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActiontargetInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetServiceUnavailable creates a GetJourneyActiontargetServiceUnavailable with default headers values
func NewGetJourneyActiontargetServiceUnavailable() *GetJourneyActiontargetServiceUnavailable {
	return &GetJourneyActiontargetServiceUnavailable{}
}

/*GetJourneyActiontargetServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyActiontargetServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActiontargetServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActiontargetGatewayTimeout creates a GetJourneyActiontargetGatewayTimeout with default headers values
func NewGetJourneyActiontargetGatewayTimeout() *GetJourneyActiontargetGatewayTimeout {
	return &GetJourneyActiontargetGatewayTimeout{}
}

/*GetJourneyActiontargetGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetJourneyActiontargetGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActiontargetGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actiontargets/{actionTargetId}][%d] getJourneyActiontargetGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActiontargetGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActiontargetGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
