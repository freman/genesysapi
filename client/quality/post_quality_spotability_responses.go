// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostQualitySpotabilityReader is a Reader for the PostQualitySpotability structure.
type PostQualitySpotabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualitySpotabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualitySpotabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualitySpotabilityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualitySpotabilityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualitySpotabilityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualitySpotabilityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualitySpotabilityRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualitySpotabilityUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualitySpotabilityTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualitySpotabilityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualitySpotabilityServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualitySpotabilityGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualitySpotabilityOK creates a PostQualitySpotabilityOK with default headers values
func NewPostQualitySpotabilityOK() *PostQualitySpotabilityOK {
	return &PostQualitySpotabilityOK{}
}

/*PostQualitySpotabilityOK handles this case with default header values.

successful operation
*/
type PostQualitySpotabilityOK struct {
	Payload *models.KeywordSet
}

func (o *PostQualitySpotabilityOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityOK  %+v", 200, o.Payload)
}

func (o *PostQualitySpotabilityOK) GetPayload() *models.KeywordSet {
	return o.Payload
}

func (o *PostQualitySpotabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeywordSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityBadRequest creates a PostQualitySpotabilityBadRequest with default headers values
func NewPostQualitySpotabilityBadRequest() *PostQualitySpotabilityBadRequest {
	return &PostQualitySpotabilityBadRequest{}
}

/*PostQualitySpotabilityBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualitySpotabilityBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualitySpotabilityBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityUnauthorized creates a PostQualitySpotabilityUnauthorized with default headers values
func NewPostQualitySpotabilityUnauthorized() *PostQualitySpotabilityUnauthorized {
	return &PostQualitySpotabilityUnauthorized{}
}

/*PostQualitySpotabilityUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualitySpotabilityUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualitySpotabilityUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityForbidden creates a PostQualitySpotabilityForbidden with default headers values
func NewPostQualitySpotabilityForbidden() *PostQualitySpotabilityForbidden {
	return &PostQualitySpotabilityForbidden{}
}

/*PostQualitySpotabilityForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualitySpotabilityForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityForbidden  %+v", 403, o.Payload)
}

func (o *PostQualitySpotabilityForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityNotFound creates a PostQualitySpotabilityNotFound with default headers values
func NewPostQualitySpotabilityNotFound() *PostQualitySpotabilityNotFound {
	return &PostQualitySpotabilityNotFound{}
}

/*PostQualitySpotabilityNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualitySpotabilityNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityNotFound  %+v", 404, o.Payload)
}

func (o *PostQualitySpotabilityNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityRequestEntityTooLarge creates a PostQualitySpotabilityRequestEntityTooLarge with default headers values
func NewPostQualitySpotabilityRequestEntityTooLarge() *PostQualitySpotabilityRequestEntityTooLarge {
	return &PostQualitySpotabilityRequestEntityTooLarge{}
}

/*PostQualitySpotabilityRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostQualitySpotabilityRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualitySpotabilityRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityUnsupportedMediaType creates a PostQualitySpotabilityUnsupportedMediaType with default headers values
func NewPostQualitySpotabilityUnsupportedMediaType() *PostQualitySpotabilityUnsupportedMediaType {
	return &PostQualitySpotabilityUnsupportedMediaType{}
}

/*PostQualitySpotabilityUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualitySpotabilityUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualitySpotabilityUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityTooManyRequests creates a PostQualitySpotabilityTooManyRequests with default headers values
func NewPostQualitySpotabilityTooManyRequests() *PostQualitySpotabilityTooManyRequests {
	return &PostQualitySpotabilityTooManyRequests{}
}

/*PostQualitySpotabilityTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostQualitySpotabilityTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualitySpotabilityTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityInternalServerError creates a PostQualitySpotabilityInternalServerError with default headers values
func NewPostQualitySpotabilityInternalServerError() *PostQualitySpotabilityInternalServerError {
	return &PostQualitySpotabilityInternalServerError{}
}

/*PostQualitySpotabilityInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualitySpotabilityInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualitySpotabilityInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityServiceUnavailable creates a PostQualitySpotabilityServiceUnavailable with default headers values
func NewPostQualitySpotabilityServiceUnavailable() *PostQualitySpotabilityServiceUnavailable {
	return &PostQualitySpotabilityServiceUnavailable{}
}

/*PostQualitySpotabilityServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualitySpotabilityServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualitySpotabilityServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualitySpotabilityGatewayTimeout creates a PostQualitySpotabilityGatewayTimeout with default headers values
func NewPostQualitySpotabilityGatewayTimeout() *PostQualitySpotabilityGatewayTimeout {
	return &PostQualitySpotabilityGatewayTimeout{}
}

/*PostQualitySpotabilityGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualitySpotabilityGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualitySpotabilityGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/spotability][%d] postQualitySpotabilityGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualitySpotabilityGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualitySpotabilityGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
