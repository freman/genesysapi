// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectIvrsReader is a Reader for the PostArchitectIvrs structure.
type PostArchitectIvrsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectIvrsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostArchitectIvrsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectIvrsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectIvrsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectIvrsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectIvrsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectIvrsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectIvrsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectIvrsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectIvrsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectIvrsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectIvrsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectIvrsOK creates a PostArchitectIvrsOK with default headers values
func NewPostArchitectIvrsOK() *PostArchitectIvrsOK {
	return &PostArchitectIvrsOK{}
}

/*PostArchitectIvrsOK handles this case with default header values.

successful operation
*/
type PostArchitectIvrsOK struct {
	Payload *models.IVR
}

func (o *PostArchitectIvrsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsOK  %+v", 200, o.Payload)
}

func (o *PostArchitectIvrsOK) GetPayload() *models.IVR {
	return o.Payload
}

func (o *PostArchitectIvrsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IVR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsBadRequest creates a PostArchitectIvrsBadRequest with default headers values
func NewPostArchitectIvrsBadRequest() *PostArchitectIvrsBadRequest {
	return &PostArchitectIvrsBadRequest{}
}

/*PostArchitectIvrsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectIvrsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectIvrsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsUnauthorized creates a PostArchitectIvrsUnauthorized with default headers values
func NewPostArchitectIvrsUnauthorized() *PostArchitectIvrsUnauthorized {
	return &PostArchitectIvrsUnauthorized{}
}

/*PostArchitectIvrsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectIvrsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectIvrsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsForbidden creates a PostArchitectIvrsForbidden with default headers values
func NewPostArchitectIvrsForbidden() *PostArchitectIvrsForbidden {
	return &PostArchitectIvrsForbidden{}
}

/*PostArchitectIvrsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectIvrsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectIvrsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsNotFound creates a PostArchitectIvrsNotFound with default headers values
func NewPostArchitectIvrsNotFound() *PostArchitectIvrsNotFound {
	return &PostArchitectIvrsNotFound{}
}

/*PostArchitectIvrsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostArchitectIvrsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectIvrsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsRequestEntityTooLarge creates a PostArchitectIvrsRequestEntityTooLarge with default headers values
func NewPostArchitectIvrsRequestEntityTooLarge() *PostArchitectIvrsRequestEntityTooLarge {
	return &PostArchitectIvrsRequestEntityTooLarge{}
}

/*PostArchitectIvrsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostArchitectIvrsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectIvrsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsUnsupportedMediaType creates a PostArchitectIvrsUnsupportedMediaType with default headers values
func NewPostArchitectIvrsUnsupportedMediaType() *PostArchitectIvrsUnsupportedMediaType {
	return &PostArchitectIvrsUnsupportedMediaType{}
}

/*PostArchitectIvrsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectIvrsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectIvrsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsTooManyRequests creates a PostArchitectIvrsTooManyRequests with default headers values
func NewPostArchitectIvrsTooManyRequests() *PostArchitectIvrsTooManyRequests {
	return &PostArchitectIvrsTooManyRequests{}
}

/*PostArchitectIvrsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostArchitectIvrsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectIvrsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsInternalServerError creates a PostArchitectIvrsInternalServerError with default headers values
func NewPostArchitectIvrsInternalServerError() *PostArchitectIvrsInternalServerError {
	return &PostArchitectIvrsInternalServerError{}
}

/*PostArchitectIvrsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectIvrsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectIvrsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsServiceUnavailable creates a PostArchitectIvrsServiceUnavailable with default headers values
func NewPostArchitectIvrsServiceUnavailable() *PostArchitectIvrsServiceUnavailable {
	return &PostArchitectIvrsServiceUnavailable{}
}

/*PostArchitectIvrsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectIvrsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectIvrsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectIvrsGatewayTimeout creates a PostArchitectIvrsGatewayTimeout with default headers values
func NewPostArchitectIvrsGatewayTimeout() *PostArchitectIvrsGatewayTimeout {
	return &PostArchitectIvrsGatewayTimeout{}
}

/*PostArchitectIvrsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostArchitectIvrsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectIvrsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/ivrs][%d] postArchitectIvrsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectIvrsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectIvrsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
