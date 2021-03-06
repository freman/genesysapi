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

// PostTelephonyProvidersEdgeLogicalinterfacesReader is a Reader for the PostTelephonyProvidersEdgeLogicalinterfaces structure.
type PostTelephonyProvidersEdgeLogicalinterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeLogicalinterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesOK creates a PostTelephonyProvidersEdgeLogicalinterfacesOK with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesOK() *PostTelephonyProvidersEdgeLogicalinterfacesOK {
	return &PostTelephonyProvidersEdgeLogicalinterfacesOK{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgeLogicalinterfacesOK struct {
	Payload *models.DomainLogicalInterface
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) GetPayload() *models.DomainLogicalInterface {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainLogicalInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesBadRequest creates a PostTelephonyProvidersEdgeLogicalinterfacesBadRequest with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesBadRequest() *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest {
	return &PostTelephonyProvidersEdgeLogicalinterfacesBadRequest{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesUnauthorized creates a PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesUnauthorized() *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized {
	return &PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesForbidden creates a PostTelephonyProvidersEdgeLogicalinterfacesForbidden with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesForbidden() *PostTelephonyProvidersEdgeLogicalinterfacesForbidden {
	return &PostTelephonyProvidersEdgeLogicalinterfacesForbidden{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesNotFound creates a PostTelephonyProvidersEdgeLogicalinterfacesNotFound with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesNotFound() *PostTelephonyProvidersEdgeLogicalinterfacesNotFound {
	return &PostTelephonyProvidersEdgeLogicalinterfacesNotFound{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge creates a PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge() *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType creates a PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType() *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests creates a PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests() *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests {
	return &PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesInternalServerError creates a PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesInternalServerError() *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError {
	return &PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable creates a PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable() *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable {
	return &PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout creates a PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout() *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout {
	return &PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout{}
}

/*PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/logicalinterfaces][%d] postTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeLogicalinterfacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
