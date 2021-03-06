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

// PostTelephonyProvidersEdgesDidpoolsReader is a Reader for the PostTelephonyProvidersEdgesDidpools structure.
type PostTelephonyProvidersEdgesDidpoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesDidpoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesDidpoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesDidpoolsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesDidpoolsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesDidpoolsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesDidpoolsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostTelephonyProvidersEdgesDidpoolsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesDidpoolsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesDidpoolsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesDidpoolsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesDidpoolsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesDidpoolsOK creates a PostTelephonyProvidersEdgesDidpoolsOK with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsOK() *PostTelephonyProvidersEdgesDidpoolsOK {
	return &PostTelephonyProvidersEdgesDidpoolsOK{}
}

/*PostTelephonyProvidersEdgesDidpoolsOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesDidpoolsOK struct {
	Payload *models.DIDPool
}

func (o *PostTelephonyProvidersEdgesDidpoolsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsOK) GetPayload() *models.DIDPool {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DIDPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsBadRequest creates a PostTelephonyProvidersEdgesDidpoolsBadRequest with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsBadRequest() *PostTelephonyProvidersEdgesDidpoolsBadRequest {
	return &PostTelephonyProvidersEdgesDidpoolsBadRequest{}
}

/*PostTelephonyProvidersEdgesDidpoolsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesDidpoolsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsUnauthorized creates a PostTelephonyProvidersEdgesDidpoolsUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsUnauthorized() *PostTelephonyProvidersEdgesDidpoolsUnauthorized {
	return &PostTelephonyProvidersEdgesDidpoolsUnauthorized{}
}

/*PostTelephonyProvidersEdgesDidpoolsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesDidpoolsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsForbidden creates a PostTelephonyProvidersEdgesDidpoolsForbidden with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsForbidden() *PostTelephonyProvidersEdgesDidpoolsForbidden {
	return &PostTelephonyProvidersEdgesDidpoolsForbidden{}
}

/*PostTelephonyProvidersEdgesDidpoolsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesDidpoolsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsNotFound creates a PostTelephonyProvidersEdgesDidpoolsNotFound with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsNotFound() *PostTelephonyProvidersEdgesDidpoolsNotFound {
	return &PostTelephonyProvidersEdgesDidpoolsNotFound{}
}

/*PostTelephonyProvidersEdgesDidpoolsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesDidpoolsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsConflict creates a PostTelephonyProvidersEdgesDidpoolsConflict with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsConflict() *PostTelephonyProvidersEdgesDidpoolsConflict {
	return &PostTelephonyProvidersEdgesDidpoolsConflict{}
}

/*PostTelephonyProvidersEdgesDidpoolsConflict handles this case with default header values.

Conflict
*/
type PostTelephonyProvidersEdgesDidpoolsConflict struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsConflict  %+v", 409, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge creates a PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge() *PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType creates a PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType() *PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsTooManyRequests creates a PostTelephonyProvidersEdgesDidpoolsTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsTooManyRequests() *PostTelephonyProvidersEdgesDidpoolsTooManyRequests {
	return &PostTelephonyProvidersEdgesDidpoolsTooManyRequests{}
}

/*PostTelephonyProvidersEdgesDidpoolsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesDidpoolsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsInternalServerError creates a PostTelephonyProvidersEdgesDidpoolsInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsInternalServerError() *PostTelephonyProvidersEdgesDidpoolsInternalServerError {
	return &PostTelephonyProvidersEdgesDidpoolsInternalServerError{}
}

/*PostTelephonyProvidersEdgesDidpoolsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesDidpoolsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsServiceUnavailable creates a PostTelephonyProvidersEdgesDidpoolsServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsServiceUnavailable() *PostTelephonyProvidersEdgesDidpoolsServiceUnavailable {
	return &PostTelephonyProvidersEdgesDidpoolsServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesDidpoolsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesDidpoolsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesDidpoolsGatewayTimeout creates a PostTelephonyProvidersEdgesDidpoolsGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesDidpoolsGatewayTimeout() *PostTelephonyProvidersEdgesDidpoolsGatewayTimeout {
	return &PostTelephonyProvidersEdgesDidpoolsGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesDidpoolsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesDidpoolsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesDidpoolsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/didpools][%d] postTelephonyProvidersEdgesDidpoolsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesDidpoolsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesDidpoolsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
