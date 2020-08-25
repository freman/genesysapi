// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteTelephonyProvidersEdgesSiteReader is a Reader for the DeleteTelephonyProvidersEdgesSite structure.
type DeleteTelephonyProvidersEdgesSiteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTelephonyProvidersEdgesSiteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTelephonyProvidersEdgesSiteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTelephonyProvidersEdgesSiteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTelephonyProvidersEdgesSiteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteTelephonyProvidersEdgesSiteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTelephonyProvidersEdgesSiteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteTelephonyProvidersEdgesSiteConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteTelephonyProvidersEdgesSiteUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteTelephonyProvidersEdgesSiteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteTelephonyProvidersEdgesSiteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteTelephonyProvidersEdgesSiteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteTelephonyProvidersEdgesSiteGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTelephonyProvidersEdgesSiteOK creates a DeleteTelephonyProvidersEdgesSiteOK with default headers values
func NewDeleteTelephonyProvidersEdgesSiteOK() *DeleteTelephonyProvidersEdgesSiteOK {
	return &DeleteTelephonyProvidersEdgesSiteOK{}
}

/*DeleteTelephonyProvidersEdgesSiteOK handles this case with default header values.

Operation was successful.
*/
type DeleteTelephonyProvidersEdgesSiteOK struct {
}

func (o *DeleteTelephonyProvidersEdgesSiteOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteOK ", 200)
}

func (o *DeleteTelephonyProvidersEdgesSiteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteBadRequest creates a DeleteTelephonyProvidersEdgesSiteBadRequest with default headers values
func NewDeleteTelephonyProvidersEdgesSiteBadRequest() *DeleteTelephonyProvidersEdgesSiteBadRequest {
	return &DeleteTelephonyProvidersEdgesSiteBadRequest{}
}

/*DeleteTelephonyProvidersEdgesSiteBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteTelephonyProvidersEdgesSiteBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteUnauthorized creates a DeleteTelephonyProvidersEdgesSiteUnauthorized with default headers values
func NewDeleteTelephonyProvidersEdgesSiteUnauthorized() *DeleteTelephonyProvidersEdgesSiteUnauthorized {
	return &DeleteTelephonyProvidersEdgesSiteUnauthorized{}
}

/*DeleteTelephonyProvidersEdgesSiteUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteTelephonyProvidersEdgesSiteUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteForbidden creates a DeleteTelephonyProvidersEdgesSiteForbidden with default headers values
func NewDeleteTelephonyProvidersEdgesSiteForbidden() *DeleteTelephonyProvidersEdgesSiteForbidden {
	return &DeleteTelephonyProvidersEdgesSiteForbidden{}
}

/*DeleteTelephonyProvidersEdgesSiteForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteTelephonyProvidersEdgesSiteForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteNotFound creates a DeleteTelephonyProvidersEdgesSiteNotFound with default headers values
func NewDeleteTelephonyProvidersEdgesSiteNotFound() *DeleteTelephonyProvidersEdgesSiteNotFound {
	return &DeleteTelephonyProvidersEdgesSiteNotFound{}
}

/*DeleteTelephonyProvidersEdgesSiteNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteTelephonyProvidersEdgesSiteNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteConflict creates a DeleteTelephonyProvidersEdgesSiteConflict with default headers values
func NewDeleteTelephonyProvidersEdgesSiteConflict() *DeleteTelephonyProvidersEdgesSiteConflict {
	return &DeleteTelephonyProvidersEdgesSiteConflict{}
}

/*DeleteTelephonyProvidersEdgesSiteConflict handles this case with default header values.

Conflict
*/
type DeleteTelephonyProvidersEdgesSiteConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteConflict  %+v", 409, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge creates a DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge with default headers values
func NewDeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge() *DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge {
	return &DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge{}
}

/*DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteUnsupportedMediaType creates a DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType with default headers values
func NewDeleteTelephonyProvidersEdgesSiteUnsupportedMediaType() *DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType {
	return &DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType{}
}

/*DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteTooManyRequests creates a DeleteTelephonyProvidersEdgesSiteTooManyRequests with default headers values
func NewDeleteTelephonyProvidersEdgesSiteTooManyRequests() *DeleteTelephonyProvidersEdgesSiteTooManyRequests {
	return &DeleteTelephonyProvidersEdgesSiteTooManyRequests{}
}

/*DeleteTelephonyProvidersEdgesSiteTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteTelephonyProvidersEdgesSiteTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteInternalServerError creates a DeleteTelephonyProvidersEdgesSiteInternalServerError with default headers values
func NewDeleteTelephonyProvidersEdgesSiteInternalServerError() *DeleteTelephonyProvidersEdgesSiteInternalServerError {
	return &DeleteTelephonyProvidersEdgesSiteInternalServerError{}
}

/*DeleteTelephonyProvidersEdgesSiteInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteTelephonyProvidersEdgesSiteInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteServiceUnavailable creates a DeleteTelephonyProvidersEdgesSiteServiceUnavailable with default headers values
func NewDeleteTelephonyProvidersEdgesSiteServiceUnavailable() *DeleteTelephonyProvidersEdgesSiteServiceUnavailable {
	return &DeleteTelephonyProvidersEdgesSiteServiceUnavailable{}
}

/*DeleteTelephonyProvidersEdgesSiteServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteTelephonyProvidersEdgesSiteServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTelephonyProvidersEdgesSiteGatewayTimeout creates a DeleteTelephonyProvidersEdgesSiteGatewayTimeout with default headers values
func NewDeleteTelephonyProvidersEdgesSiteGatewayTimeout() *DeleteTelephonyProvidersEdgesSiteGatewayTimeout {
	return &DeleteTelephonyProvidersEdgesSiteGatewayTimeout{}
}

/*DeleteTelephonyProvidersEdgesSiteGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteTelephonyProvidersEdgesSiteGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteTelephonyProvidersEdgesSiteGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/telephony/providers/edges/sites/{siteId}][%d] deleteTelephonyProvidersEdgesSiteGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteTelephonyProvidersEdgesSiteGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteTelephonyProvidersEdgesSiteGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}