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

// PostTelephonyProvidersEdgesSitesReader is a Reader for the PostTelephonyProvidersEdgesSites structure.
type PostTelephonyProvidersEdgesSitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesSitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesSitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesSitesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesSitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesSitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesSitesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesSitesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesSitesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesSitesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesSitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesSitesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesSitesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesSitesOK creates a PostTelephonyProvidersEdgesSitesOK with default headers values
func NewPostTelephonyProvidersEdgesSitesOK() *PostTelephonyProvidersEdgesSitesOK {
	return &PostTelephonyProvidersEdgesSitesOK{}
}

/*PostTelephonyProvidersEdgesSitesOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesSitesOK struct {
	Payload *models.Site
}

func (o *PostTelephonyProvidersEdgesSitesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesBadRequest creates a PostTelephonyProvidersEdgesSitesBadRequest with default headers values
func NewPostTelephonyProvidersEdgesSitesBadRequest() *PostTelephonyProvidersEdgesSitesBadRequest {
	return &PostTelephonyProvidersEdgesSitesBadRequest{}
}

/*PostTelephonyProvidersEdgesSitesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesSitesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesUnauthorized creates a PostTelephonyProvidersEdgesSitesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesSitesUnauthorized() *PostTelephonyProvidersEdgesSitesUnauthorized {
	return &PostTelephonyProvidersEdgesSitesUnauthorized{}
}

/*PostTelephonyProvidersEdgesSitesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesSitesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesForbidden creates a PostTelephonyProvidersEdgesSitesForbidden with default headers values
func NewPostTelephonyProvidersEdgesSitesForbidden() *PostTelephonyProvidersEdgesSitesForbidden {
	return &PostTelephonyProvidersEdgesSitesForbidden{}
}

/*PostTelephonyProvidersEdgesSitesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesSitesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesNotFound creates a PostTelephonyProvidersEdgesSitesNotFound with default headers values
func NewPostTelephonyProvidersEdgesSitesNotFound() *PostTelephonyProvidersEdgesSitesNotFound {
	return &PostTelephonyProvidersEdgesSitesNotFound{}
}

/*PostTelephonyProvidersEdgesSitesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesSitesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesRequestEntityTooLarge creates a PostTelephonyProvidersEdgesSitesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesSitesRequestEntityTooLarge() *PostTelephonyProvidersEdgesSitesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesSitesRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesSitesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesSitesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesUnsupportedMediaType creates a PostTelephonyProvidersEdgesSitesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesSitesUnsupportedMediaType() *PostTelephonyProvidersEdgesSitesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesSitesUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesSitesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesSitesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesTooManyRequests creates a PostTelephonyProvidersEdgesSitesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesSitesTooManyRequests() *PostTelephonyProvidersEdgesSitesTooManyRequests {
	return &PostTelephonyProvidersEdgesSitesTooManyRequests{}
}

/*PostTelephonyProvidersEdgesSitesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesSitesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesInternalServerError creates a PostTelephonyProvidersEdgesSitesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesSitesInternalServerError() *PostTelephonyProvidersEdgesSitesInternalServerError {
	return &PostTelephonyProvidersEdgesSitesInternalServerError{}
}

/*PostTelephonyProvidersEdgesSitesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesSitesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesServiceUnavailable creates a PostTelephonyProvidersEdgesSitesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesSitesServiceUnavailable() *PostTelephonyProvidersEdgesSitesServiceUnavailable {
	return &PostTelephonyProvidersEdgesSitesServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesSitesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesSitesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesSitesGatewayTimeout creates a PostTelephonyProvidersEdgesSitesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesSitesGatewayTimeout() *PostTelephonyProvidersEdgesSitesGatewayTimeout {
	return &PostTelephonyProvidersEdgesSitesGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesSitesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesSitesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesSitesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/sites][%d] postTelephonyProvidersEdgesSitesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesSitesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesSitesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
