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

// PostTelephonyProvidersEdgesCertificateauthoritiesReader is a Reader for the PostTelephonyProvidersEdgesCertificateauthorities structure.
type PostTelephonyProvidersEdgesCertificateauthoritiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesCertificateauthoritiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesOK creates a PostTelephonyProvidersEdgesCertificateauthoritiesOK with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesOK() *PostTelephonyProvidersEdgesCertificateauthoritiesOK {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesOK{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesOK struct {
	Payload *models.DomainCertificateAuthority
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesOK) GetPayload() *models.DomainCertificateAuthority {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCertificateAuthority)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesBadRequest creates a PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesBadRequest() *PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized creates a PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized() *PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesForbidden creates a PostTelephonyProvidersEdgesCertificateauthoritiesForbidden with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesForbidden() *PostTelephonyProvidersEdgesCertificateauthoritiesForbidden {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesForbidden{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesNotFound creates a PostTelephonyProvidersEdgesCertificateauthoritiesNotFound with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesNotFound() *PostTelephonyProvidersEdgesCertificateauthoritiesNotFound {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesNotFound{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge creates a PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge() *PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType creates a PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType() *PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests creates a PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests() *PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError creates a PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError() *PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable creates a PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable() *PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout creates a PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout() *PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout {
	return &PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/certificateauthorities][%d] postTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesCertificateauthoritiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}