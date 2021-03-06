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

// DeleteJourneyActionmapReader is a Reader for the DeleteJourneyActionmap structure.
type DeleteJourneyActionmapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJourneyActionmapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteJourneyActionmapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteJourneyActionmapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteJourneyActionmapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteJourneyActionmapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteJourneyActionmapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteJourneyActionmapRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteJourneyActionmapUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteJourneyActionmapTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteJourneyActionmapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteJourneyActionmapServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteJourneyActionmapGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteJourneyActionmapNoContent creates a DeleteJourneyActionmapNoContent with default headers values
func NewDeleteJourneyActionmapNoContent() *DeleteJourneyActionmapNoContent {
	return &DeleteJourneyActionmapNoContent{}
}

/*DeleteJourneyActionmapNoContent handles this case with default header values.

Action map deleted.
*/
type DeleteJourneyActionmapNoContent struct {
}

func (o *DeleteJourneyActionmapNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapNoContent ", 204)
}

func (o *DeleteJourneyActionmapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJourneyActionmapBadRequest creates a DeleteJourneyActionmapBadRequest with default headers values
func NewDeleteJourneyActionmapBadRequest() *DeleteJourneyActionmapBadRequest {
	return &DeleteJourneyActionmapBadRequest{}
}

/*DeleteJourneyActionmapBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteJourneyActionmapBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteJourneyActionmapBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapUnauthorized creates a DeleteJourneyActionmapUnauthorized with default headers values
func NewDeleteJourneyActionmapUnauthorized() *DeleteJourneyActionmapUnauthorized {
	return &DeleteJourneyActionmapUnauthorized{}
}

/*DeleteJourneyActionmapUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteJourneyActionmapUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteJourneyActionmapUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapForbidden creates a DeleteJourneyActionmapForbidden with default headers values
func NewDeleteJourneyActionmapForbidden() *DeleteJourneyActionmapForbidden {
	return &DeleteJourneyActionmapForbidden{}
}

/*DeleteJourneyActionmapForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteJourneyActionmapForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapForbidden  %+v", 403, o.Payload)
}

func (o *DeleteJourneyActionmapForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapNotFound creates a DeleteJourneyActionmapNotFound with default headers values
func NewDeleteJourneyActionmapNotFound() *DeleteJourneyActionmapNotFound {
	return &DeleteJourneyActionmapNotFound{}
}

/*DeleteJourneyActionmapNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteJourneyActionmapNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapNotFound  %+v", 404, o.Payload)
}

func (o *DeleteJourneyActionmapNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapRequestEntityTooLarge creates a DeleteJourneyActionmapRequestEntityTooLarge with default headers values
func NewDeleteJourneyActionmapRequestEntityTooLarge() *DeleteJourneyActionmapRequestEntityTooLarge {
	return &DeleteJourneyActionmapRequestEntityTooLarge{}
}

/*DeleteJourneyActionmapRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteJourneyActionmapRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteJourneyActionmapRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapUnsupportedMediaType creates a DeleteJourneyActionmapUnsupportedMediaType with default headers values
func NewDeleteJourneyActionmapUnsupportedMediaType() *DeleteJourneyActionmapUnsupportedMediaType {
	return &DeleteJourneyActionmapUnsupportedMediaType{}
}

/*DeleteJourneyActionmapUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteJourneyActionmapUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteJourneyActionmapUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapTooManyRequests creates a DeleteJourneyActionmapTooManyRequests with default headers values
func NewDeleteJourneyActionmapTooManyRequests() *DeleteJourneyActionmapTooManyRequests {
	return &DeleteJourneyActionmapTooManyRequests{}
}

/*DeleteJourneyActionmapTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteJourneyActionmapTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteJourneyActionmapTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapInternalServerError creates a DeleteJourneyActionmapInternalServerError with default headers values
func NewDeleteJourneyActionmapInternalServerError() *DeleteJourneyActionmapInternalServerError {
	return &DeleteJourneyActionmapInternalServerError{}
}

/*DeleteJourneyActionmapInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteJourneyActionmapInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteJourneyActionmapInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapServiceUnavailable creates a DeleteJourneyActionmapServiceUnavailable with default headers values
func NewDeleteJourneyActionmapServiceUnavailable() *DeleteJourneyActionmapServiceUnavailable {
	return &DeleteJourneyActionmapServiceUnavailable{}
}

/*DeleteJourneyActionmapServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteJourneyActionmapServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteJourneyActionmapServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJourneyActionmapGatewayTimeout creates a DeleteJourneyActionmapGatewayTimeout with default headers values
func NewDeleteJourneyActionmapGatewayTimeout() *DeleteJourneyActionmapGatewayTimeout {
	return &DeleteJourneyActionmapGatewayTimeout{}
}

/*DeleteJourneyActionmapGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteJourneyActionmapGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteJourneyActionmapGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/journey/actionmaps/{actionMapId}][%d] deleteJourneyActionmapGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteJourneyActionmapGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteJourneyActionmapGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
