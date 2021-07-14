// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostResponsemanagementLibrariesReader is a Reader for the PostResponsemanagementLibraries structure.
type PostResponsemanagementLibrariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostResponsemanagementLibrariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostResponsemanagementLibrariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostResponsemanagementLibrariesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostResponsemanagementLibrariesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostResponsemanagementLibrariesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostResponsemanagementLibrariesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostResponsemanagementLibrariesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostResponsemanagementLibrariesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostResponsemanagementLibrariesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostResponsemanagementLibrariesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostResponsemanagementLibrariesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostResponsemanagementLibrariesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostResponsemanagementLibrariesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostResponsemanagementLibrariesOK creates a PostResponsemanagementLibrariesOK with default headers values
func NewPostResponsemanagementLibrariesOK() *PostResponsemanagementLibrariesOK {
	return &PostResponsemanagementLibrariesOK{}
}

/*PostResponsemanagementLibrariesOK handles this case with default header values.

successful operation
*/
type PostResponsemanagementLibrariesOK struct {
	Payload *models.Library
}

func (o *PostResponsemanagementLibrariesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesOK  %+v", 200, o.Payload)
}

func (o *PostResponsemanagementLibrariesOK) GetPayload() *models.Library {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Library)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesBadRequest creates a PostResponsemanagementLibrariesBadRequest with default headers values
func NewPostResponsemanagementLibrariesBadRequest() *PostResponsemanagementLibrariesBadRequest {
	return &PostResponsemanagementLibrariesBadRequest{}
}

/*PostResponsemanagementLibrariesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostResponsemanagementLibrariesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesBadRequest  %+v", 400, o.Payload)
}

func (o *PostResponsemanagementLibrariesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesUnauthorized creates a PostResponsemanagementLibrariesUnauthorized with default headers values
func NewPostResponsemanagementLibrariesUnauthorized() *PostResponsemanagementLibrariesUnauthorized {
	return &PostResponsemanagementLibrariesUnauthorized{}
}

/*PostResponsemanagementLibrariesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostResponsemanagementLibrariesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostResponsemanagementLibrariesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesForbidden creates a PostResponsemanagementLibrariesForbidden with default headers values
func NewPostResponsemanagementLibrariesForbidden() *PostResponsemanagementLibrariesForbidden {
	return &PostResponsemanagementLibrariesForbidden{}
}

/*PostResponsemanagementLibrariesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostResponsemanagementLibrariesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesForbidden  %+v", 403, o.Payload)
}

func (o *PostResponsemanagementLibrariesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesNotFound creates a PostResponsemanagementLibrariesNotFound with default headers values
func NewPostResponsemanagementLibrariesNotFound() *PostResponsemanagementLibrariesNotFound {
	return &PostResponsemanagementLibrariesNotFound{}
}

/*PostResponsemanagementLibrariesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostResponsemanagementLibrariesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesNotFound  %+v", 404, o.Payload)
}

func (o *PostResponsemanagementLibrariesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesRequestTimeout creates a PostResponsemanagementLibrariesRequestTimeout with default headers values
func NewPostResponsemanagementLibrariesRequestTimeout() *PostResponsemanagementLibrariesRequestTimeout {
	return &PostResponsemanagementLibrariesRequestTimeout{}
}

/*PostResponsemanagementLibrariesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostResponsemanagementLibrariesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostResponsemanagementLibrariesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesRequestEntityTooLarge creates a PostResponsemanagementLibrariesRequestEntityTooLarge with default headers values
func NewPostResponsemanagementLibrariesRequestEntityTooLarge() *PostResponsemanagementLibrariesRequestEntityTooLarge {
	return &PostResponsemanagementLibrariesRequestEntityTooLarge{}
}

/*PostResponsemanagementLibrariesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostResponsemanagementLibrariesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostResponsemanagementLibrariesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesUnsupportedMediaType creates a PostResponsemanagementLibrariesUnsupportedMediaType with default headers values
func NewPostResponsemanagementLibrariesUnsupportedMediaType() *PostResponsemanagementLibrariesUnsupportedMediaType {
	return &PostResponsemanagementLibrariesUnsupportedMediaType{}
}

/*PostResponsemanagementLibrariesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostResponsemanagementLibrariesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostResponsemanagementLibrariesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesTooManyRequests creates a PostResponsemanagementLibrariesTooManyRequests with default headers values
func NewPostResponsemanagementLibrariesTooManyRequests() *PostResponsemanagementLibrariesTooManyRequests {
	return &PostResponsemanagementLibrariesTooManyRequests{}
}

/*PostResponsemanagementLibrariesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostResponsemanagementLibrariesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostResponsemanagementLibrariesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesInternalServerError creates a PostResponsemanagementLibrariesInternalServerError with default headers values
func NewPostResponsemanagementLibrariesInternalServerError() *PostResponsemanagementLibrariesInternalServerError {
	return &PostResponsemanagementLibrariesInternalServerError{}
}

/*PostResponsemanagementLibrariesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostResponsemanagementLibrariesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostResponsemanagementLibrariesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesServiceUnavailable creates a PostResponsemanagementLibrariesServiceUnavailable with default headers values
func NewPostResponsemanagementLibrariesServiceUnavailable() *PostResponsemanagementLibrariesServiceUnavailable {
	return &PostResponsemanagementLibrariesServiceUnavailable{}
}

/*PostResponsemanagementLibrariesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostResponsemanagementLibrariesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostResponsemanagementLibrariesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementLibrariesGatewayTimeout creates a PostResponsemanagementLibrariesGatewayTimeout with default headers values
func NewPostResponsemanagementLibrariesGatewayTimeout() *PostResponsemanagementLibrariesGatewayTimeout {
	return &PostResponsemanagementLibrariesGatewayTimeout{}
}

/*PostResponsemanagementLibrariesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostResponsemanagementLibrariesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostResponsemanagementLibrariesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/libraries][%d] postResponsemanagementLibrariesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostResponsemanagementLibrariesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementLibrariesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
