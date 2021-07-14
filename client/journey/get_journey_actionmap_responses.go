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

// GetJourneyActionmapReader is a Reader for the GetJourneyActionmap structure.
type GetJourneyActionmapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyActionmapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyActionmapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyActionmapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyActionmapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyActionmapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyActionmapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyActionmapRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyActionmapRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyActionmapUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyActionmapTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyActionmapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyActionmapServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyActionmapGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyActionmapOK creates a GetJourneyActionmapOK with default headers values
func NewGetJourneyActionmapOK() *GetJourneyActionmapOK {
	return &GetJourneyActionmapOK{}
}

/*GetJourneyActionmapOK handles this case with default header values.

successful operation
*/
type GetJourneyActionmapOK struct {
	Payload *models.ActionMap
}

func (o *GetJourneyActionmapOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActionmapOK) GetPayload() *models.ActionMap {
	return o.Payload
}

func (o *GetJourneyActionmapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapBadRequest creates a GetJourneyActionmapBadRequest with default headers values
func NewGetJourneyActionmapBadRequest() *GetJourneyActionmapBadRequest {
	return &GetJourneyActionmapBadRequest{}
}

/*GetJourneyActionmapBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyActionmapBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActionmapBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapUnauthorized creates a GetJourneyActionmapUnauthorized with default headers values
func NewGetJourneyActionmapUnauthorized() *GetJourneyActionmapUnauthorized {
	return &GetJourneyActionmapUnauthorized{}
}

/*GetJourneyActionmapUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyActionmapUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActionmapUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapForbidden creates a GetJourneyActionmapForbidden with default headers values
func NewGetJourneyActionmapForbidden() *GetJourneyActionmapForbidden {
	return &GetJourneyActionmapForbidden{}
}

/*GetJourneyActionmapForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyActionmapForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActionmapForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapNotFound creates a GetJourneyActionmapNotFound with default headers values
func NewGetJourneyActionmapNotFound() *GetJourneyActionmapNotFound {
	return &GetJourneyActionmapNotFound{}
}

/*GetJourneyActionmapNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetJourneyActionmapNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActionmapNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapRequestTimeout creates a GetJourneyActionmapRequestTimeout with default headers values
func NewGetJourneyActionmapRequestTimeout() *GetJourneyActionmapRequestTimeout {
	return &GetJourneyActionmapRequestTimeout{}
}

/*GetJourneyActionmapRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyActionmapRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActionmapRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapRequestEntityTooLarge creates a GetJourneyActionmapRequestEntityTooLarge with default headers values
func NewGetJourneyActionmapRequestEntityTooLarge() *GetJourneyActionmapRequestEntityTooLarge {
	return &GetJourneyActionmapRequestEntityTooLarge{}
}

/*GetJourneyActionmapRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetJourneyActionmapRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActionmapRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapUnsupportedMediaType creates a GetJourneyActionmapUnsupportedMediaType with default headers values
func NewGetJourneyActionmapUnsupportedMediaType() *GetJourneyActionmapUnsupportedMediaType {
	return &GetJourneyActionmapUnsupportedMediaType{}
}

/*GetJourneyActionmapUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyActionmapUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActionmapUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapTooManyRequests creates a GetJourneyActionmapTooManyRequests with default headers values
func NewGetJourneyActionmapTooManyRequests() *GetJourneyActionmapTooManyRequests {
	return &GetJourneyActionmapTooManyRequests{}
}

/*GetJourneyActionmapTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyActionmapTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActionmapTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapInternalServerError creates a GetJourneyActionmapInternalServerError with default headers values
func NewGetJourneyActionmapInternalServerError() *GetJourneyActionmapInternalServerError {
	return &GetJourneyActionmapInternalServerError{}
}

/*GetJourneyActionmapInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyActionmapInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActionmapInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapServiceUnavailable creates a GetJourneyActionmapServiceUnavailable with default headers values
func NewGetJourneyActionmapServiceUnavailable() *GetJourneyActionmapServiceUnavailable {
	return &GetJourneyActionmapServiceUnavailable{}
}

/*GetJourneyActionmapServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyActionmapServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActionmapServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapGatewayTimeout creates a GetJourneyActionmapGatewayTimeout with default headers values
func NewGetJourneyActionmapGatewayTimeout() *GetJourneyActionmapGatewayTimeout {
	return &GetJourneyActionmapGatewayTimeout{}
}

/*GetJourneyActionmapGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetJourneyActionmapGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetJourneyActionmapGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps/{actionMapId}][%d] getJourneyActionmapGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActionmapGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
