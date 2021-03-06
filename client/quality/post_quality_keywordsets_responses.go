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

// PostQualityKeywordsetsReader is a Reader for the PostQualityKeywordsets structure.
type PostQualityKeywordsetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityKeywordsetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityKeywordsetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityKeywordsetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityKeywordsetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityKeywordsetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityKeywordsetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityKeywordsetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityKeywordsetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityKeywordsetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityKeywordsetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityKeywordsetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityKeywordsetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityKeywordsetsOK creates a PostQualityKeywordsetsOK with default headers values
func NewPostQualityKeywordsetsOK() *PostQualityKeywordsetsOK {
	return &PostQualityKeywordsetsOK{}
}

/*PostQualityKeywordsetsOK handles this case with default header values.

successful operation
*/
type PostQualityKeywordsetsOK struct {
	Payload *models.KeywordSet
}

func (o *PostQualityKeywordsetsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsOK  %+v", 200, o.Payload)
}

func (o *PostQualityKeywordsetsOK) GetPayload() *models.KeywordSet {
	return o.Payload
}

func (o *PostQualityKeywordsetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeywordSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsBadRequest creates a PostQualityKeywordsetsBadRequest with default headers values
func NewPostQualityKeywordsetsBadRequest() *PostQualityKeywordsetsBadRequest {
	return &PostQualityKeywordsetsBadRequest{}
}

/*PostQualityKeywordsetsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityKeywordsetsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityKeywordsetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsUnauthorized creates a PostQualityKeywordsetsUnauthorized with default headers values
func NewPostQualityKeywordsetsUnauthorized() *PostQualityKeywordsetsUnauthorized {
	return &PostQualityKeywordsetsUnauthorized{}
}

/*PostQualityKeywordsetsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityKeywordsetsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityKeywordsetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsForbidden creates a PostQualityKeywordsetsForbidden with default headers values
func NewPostQualityKeywordsetsForbidden() *PostQualityKeywordsetsForbidden {
	return &PostQualityKeywordsetsForbidden{}
}

/*PostQualityKeywordsetsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityKeywordsetsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityKeywordsetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsNotFound creates a PostQualityKeywordsetsNotFound with default headers values
func NewPostQualityKeywordsetsNotFound() *PostQualityKeywordsetsNotFound {
	return &PostQualityKeywordsetsNotFound{}
}

/*PostQualityKeywordsetsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostQualityKeywordsetsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityKeywordsetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsRequestEntityTooLarge creates a PostQualityKeywordsetsRequestEntityTooLarge with default headers values
func NewPostQualityKeywordsetsRequestEntityTooLarge() *PostQualityKeywordsetsRequestEntityTooLarge {
	return &PostQualityKeywordsetsRequestEntityTooLarge{}
}

/*PostQualityKeywordsetsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostQualityKeywordsetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityKeywordsetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsUnsupportedMediaType creates a PostQualityKeywordsetsUnsupportedMediaType with default headers values
func NewPostQualityKeywordsetsUnsupportedMediaType() *PostQualityKeywordsetsUnsupportedMediaType {
	return &PostQualityKeywordsetsUnsupportedMediaType{}
}

/*PostQualityKeywordsetsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityKeywordsetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityKeywordsetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsTooManyRequests creates a PostQualityKeywordsetsTooManyRequests with default headers values
func NewPostQualityKeywordsetsTooManyRequests() *PostQualityKeywordsetsTooManyRequests {
	return &PostQualityKeywordsetsTooManyRequests{}
}

/*PostQualityKeywordsetsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostQualityKeywordsetsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityKeywordsetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsInternalServerError creates a PostQualityKeywordsetsInternalServerError with default headers values
func NewPostQualityKeywordsetsInternalServerError() *PostQualityKeywordsetsInternalServerError {
	return &PostQualityKeywordsetsInternalServerError{}
}

/*PostQualityKeywordsetsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityKeywordsetsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityKeywordsetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsServiceUnavailable creates a PostQualityKeywordsetsServiceUnavailable with default headers values
func NewPostQualityKeywordsetsServiceUnavailable() *PostQualityKeywordsetsServiceUnavailable {
	return &PostQualityKeywordsetsServiceUnavailable{}
}

/*PostQualityKeywordsetsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityKeywordsetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityKeywordsetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityKeywordsetsGatewayTimeout creates a PostQualityKeywordsetsGatewayTimeout with default headers values
func NewPostQualityKeywordsetsGatewayTimeout() *PostQualityKeywordsetsGatewayTimeout {
	return &PostQualityKeywordsetsGatewayTimeout{}
}

/*PostQualityKeywordsetsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostQualityKeywordsetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostQualityKeywordsetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/keywordsets][%d] postQualityKeywordsetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityKeywordsetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityKeywordsetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
