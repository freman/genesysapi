// Code generated by go-swagger; DO NOT EDIT.

package utilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostCertificateDetailsReader is a Reader for the PostCertificateDetails structure.
type PostCertificateDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCertificateDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCertificateDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCertificateDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostCertificateDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostCertificateDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostCertificateDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostCertificateDetailsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostCertificateDetailsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostCertificateDetailsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostCertificateDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCertificateDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostCertificateDetailsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostCertificateDetailsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCertificateDetailsOK creates a PostCertificateDetailsOK with default headers values
func NewPostCertificateDetailsOK() *PostCertificateDetailsOK {
	return &PostCertificateDetailsOK{}
}

/*PostCertificateDetailsOK handles this case with default header values.

successful operation
*/
type PostCertificateDetailsOK struct {
	Payload *models.ParsedCertificate
}

func (o *PostCertificateDetailsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsOK  %+v", 200, o.Payload)
}

func (o *PostCertificateDetailsOK) GetPayload() *models.ParsedCertificate {
	return o.Payload
}

func (o *PostCertificateDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParsedCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsBadRequest creates a PostCertificateDetailsBadRequest with default headers values
func NewPostCertificateDetailsBadRequest() *PostCertificateDetailsBadRequest {
	return &PostCertificateDetailsBadRequest{}
}

/*PostCertificateDetailsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostCertificateDetailsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *PostCertificateDetailsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsUnauthorized creates a PostCertificateDetailsUnauthorized with default headers values
func NewPostCertificateDetailsUnauthorized() *PostCertificateDetailsUnauthorized {
	return &PostCertificateDetailsUnauthorized{}
}

/*PostCertificateDetailsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostCertificateDetailsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCertificateDetailsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsForbidden creates a PostCertificateDetailsForbidden with default headers values
func NewPostCertificateDetailsForbidden() *PostCertificateDetailsForbidden {
	return &PostCertificateDetailsForbidden{}
}

/*PostCertificateDetailsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostCertificateDetailsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsForbidden  %+v", 403, o.Payload)
}

func (o *PostCertificateDetailsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsNotFound creates a PostCertificateDetailsNotFound with default headers values
func NewPostCertificateDetailsNotFound() *PostCertificateDetailsNotFound {
	return &PostCertificateDetailsNotFound{}
}

/*PostCertificateDetailsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostCertificateDetailsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsNotFound  %+v", 404, o.Payload)
}

func (o *PostCertificateDetailsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsRequestTimeout creates a PostCertificateDetailsRequestTimeout with default headers values
func NewPostCertificateDetailsRequestTimeout() *PostCertificateDetailsRequestTimeout {
	return &PostCertificateDetailsRequestTimeout{}
}

/*PostCertificateDetailsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostCertificateDetailsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostCertificateDetailsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsRequestEntityTooLarge creates a PostCertificateDetailsRequestEntityTooLarge with default headers values
func NewPostCertificateDetailsRequestEntityTooLarge() *PostCertificateDetailsRequestEntityTooLarge {
	return &PostCertificateDetailsRequestEntityTooLarge{}
}

/*PostCertificateDetailsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostCertificateDetailsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostCertificateDetailsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsUnsupportedMediaType creates a PostCertificateDetailsUnsupportedMediaType with default headers values
func NewPostCertificateDetailsUnsupportedMediaType() *PostCertificateDetailsUnsupportedMediaType {
	return &PostCertificateDetailsUnsupportedMediaType{}
}

/*PostCertificateDetailsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostCertificateDetailsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostCertificateDetailsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsTooManyRequests creates a PostCertificateDetailsTooManyRequests with default headers values
func NewPostCertificateDetailsTooManyRequests() *PostCertificateDetailsTooManyRequests {
	return &PostCertificateDetailsTooManyRequests{}
}

/*PostCertificateDetailsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostCertificateDetailsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostCertificateDetailsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsInternalServerError creates a PostCertificateDetailsInternalServerError with default headers values
func NewPostCertificateDetailsInternalServerError() *PostCertificateDetailsInternalServerError {
	return &PostCertificateDetailsInternalServerError{}
}

/*PostCertificateDetailsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostCertificateDetailsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostCertificateDetailsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsServiceUnavailable creates a PostCertificateDetailsServiceUnavailable with default headers values
func NewPostCertificateDetailsServiceUnavailable() *PostCertificateDetailsServiceUnavailable {
	return &PostCertificateDetailsServiceUnavailable{}
}

/*PostCertificateDetailsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostCertificateDetailsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostCertificateDetailsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCertificateDetailsGatewayTimeout creates a PostCertificateDetailsGatewayTimeout with default headers values
func NewPostCertificateDetailsGatewayTimeout() *PostCertificateDetailsGatewayTimeout {
	return &PostCertificateDetailsGatewayTimeout{}
}

/*PostCertificateDetailsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostCertificateDetailsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostCertificateDetailsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/certificate/details][%d] postCertificateDetailsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostCertificateDetailsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostCertificateDetailsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
