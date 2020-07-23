// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOutboundAttemptlimitsReader is a Reader for the PostOutboundAttemptlimits structure.
type PostOutboundAttemptlimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundAttemptlimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOutboundAttemptlimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOutboundAttemptlimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundAttemptlimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundAttemptlimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundAttemptlimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundAttemptlimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundAttemptlimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundAttemptlimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundAttemptlimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundAttemptlimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundAttemptlimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOutboundAttemptlimitsOK creates a PostOutboundAttemptlimitsOK with default headers values
func NewPostOutboundAttemptlimitsOK() *PostOutboundAttemptlimitsOK {
	return &PostOutboundAttemptlimitsOK{}
}

/*PostOutboundAttemptlimitsOK handles this case with default header values.

successful operation
*/
type PostOutboundAttemptlimitsOK struct {
	Payload *models.AttemptLimits
}

func (o *PostOutboundAttemptlimitsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsOK  %+v", 200, o.Payload)
}

func (o *PostOutboundAttemptlimitsOK) GetPayload() *models.AttemptLimits {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttemptLimits)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsBadRequest creates a PostOutboundAttemptlimitsBadRequest with default headers values
func NewPostOutboundAttemptlimitsBadRequest() *PostOutboundAttemptlimitsBadRequest {
	return &PostOutboundAttemptlimitsBadRequest{}
}

/*PostOutboundAttemptlimitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundAttemptlimitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundAttemptlimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsUnauthorized creates a PostOutboundAttemptlimitsUnauthorized with default headers values
func NewPostOutboundAttemptlimitsUnauthorized() *PostOutboundAttemptlimitsUnauthorized {
	return &PostOutboundAttemptlimitsUnauthorized{}
}

/*PostOutboundAttemptlimitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundAttemptlimitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundAttemptlimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsForbidden creates a PostOutboundAttemptlimitsForbidden with default headers values
func NewPostOutboundAttemptlimitsForbidden() *PostOutboundAttemptlimitsForbidden {
	return &PostOutboundAttemptlimitsForbidden{}
}

/*PostOutboundAttemptlimitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundAttemptlimitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundAttemptlimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsNotFound creates a PostOutboundAttemptlimitsNotFound with default headers values
func NewPostOutboundAttemptlimitsNotFound() *PostOutboundAttemptlimitsNotFound {
	return &PostOutboundAttemptlimitsNotFound{}
}

/*PostOutboundAttemptlimitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundAttemptlimitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundAttemptlimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsRequestEntityTooLarge creates a PostOutboundAttemptlimitsRequestEntityTooLarge with default headers values
func NewPostOutboundAttemptlimitsRequestEntityTooLarge() *PostOutboundAttemptlimitsRequestEntityTooLarge {
	return &PostOutboundAttemptlimitsRequestEntityTooLarge{}
}

/*PostOutboundAttemptlimitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostOutboundAttemptlimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundAttemptlimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsUnsupportedMediaType creates a PostOutboundAttemptlimitsUnsupportedMediaType with default headers values
func NewPostOutboundAttemptlimitsUnsupportedMediaType() *PostOutboundAttemptlimitsUnsupportedMediaType {
	return &PostOutboundAttemptlimitsUnsupportedMediaType{}
}

/*PostOutboundAttemptlimitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundAttemptlimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundAttemptlimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsTooManyRequests creates a PostOutboundAttemptlimitsTooManyRequests with default headers values
func NewPostOutboundAttemptlimitsTooManyRequests() *PostOutboundAttemptlimitsTooManyRequests {
	return &PostOutboundAttemptlimitsTooManyRequests{}
}

/*PostOutboundAttemptlimitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostOutboundAttemptlimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundAttemptlimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsInternalServerError creates a PostOutboundAttemptlimitsInternalServerError with default headers values
func NewPostOutboundAttemptlimitsInternalServerError() *PostOutboundAttemptlimitsInternalServerError {
	return &PostOutboundAttemptlimitsInternalServerError{}
}

/*PostOutboundAttemptlimitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundAttemptlimitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundAttemptlimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsServiceUnavailable creates a PostOutboundAttemptlimitsServiceUnavailable with default headers values
func NewPostOutboundAttemptlimitsServiceUnavailable() *PostOutboundAttemptlimitsServiceUnavailable {
	return &PostOutboundAttemptlimitsServiceUnavailable{}
}

/*PostOutboundAttemptlimitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundAttemptlimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundAttemptlimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundAttemptlimitsGatewayTimeout creates a PostOutboundAttemptlimitsGatewayTimeout with default headers values
func NewPostOutboundAttemptlimitsGatewayTimeout() *PostOutboundAttemptlimitsGatewayTimeout {
	return &PostOutboundAttemptlimitsGatewayTimeout{}
}

/*PostOutboundAttemptlimitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundAttemptlimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundAttemptlimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/attemptlimits][%d] postOutboundAttemptlimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundAttemptlimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundAttemptlimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
