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

// DeleteQualityKeywordsetsReader is a Reader for the DeleteQualityKeywordsets structure.
type DeleteQualityKeywordsetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteQualityKeywordsetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteQualityKeywordsetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteQualityKeywordsetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteQualityKeywordsetsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteQualityKeywordsetsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteQualityKeywordsetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteQualityKeywordsetsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteQualityKeywordsetsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteQualityKeywordsetsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteQualityKeywordsetsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteQualityKeywordsetsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteQualityKeywordsetsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteQualityKeywordsetsOK creates a DeleteQualityKeywordsetsOK with default headers values
func NewDeleteQualityKeywordsetsOK() *DeleteQualityKeywordsetsOK {
	return &DeleteQualityKeywordsetsOK{}
}

/*DeleteQualityKeywordsetsOK handles this case with default header values.

Operation was successful.
*/
type DeleteQualityKeywordsetsOK struct {
}

func (o *DeleteQualityKeywordsetsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsOK ", 200)
}

func (o *DeleteQualityKeywordsetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteQualityKeywordsetsBadRequest creates a DeleteQualityKeywordsetsBadRequest with default headers values
func NewDeleteQualityKeywordsetsBadRequest() *DeleteQualityKeywordsetsBadRequest {
	return &DeleteQualityKeywordsetsBadRequest{}
}

/*DeleteQualityKeywordsetsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteQualityKeywordsetsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteQualityKeywordsetsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsUnauthorized creates a DeleteQualityKeywordsetsUnauthorized with default headers values
func NewDeleteQualityKeywordsetsUnauthorized() *DeleteQualityKeywordsetsUnauthorized {
	return &DeleteQualityKeywordsetsUnauthorized{}
}

/*DeleteQualityKeywordsetsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteQualityKeywordsetsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteQualityKeywordsetsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsForbidden creates a DeleteQualityKeywordsetsForbidden with default headers values
func NewDeleteQualityKeywordsetsForbidden() *DeleteQualityKeywordsetsForbidden {
	return &DeleteQualityKeywordsetsForbidden{}
}

/*DeleteQualityKeywordsetsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteQualityKeywordsetsForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteQualityKeywordsetsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsNotFound creates a DeleteQualityKeywordsetsNotFound with default headers values
func NewDeleteQualityKeywordsetsNotFound() *DeleteQualityKeywordsetsNotFound {
	return &DeleteQualityKeywordsetsNotFound{}
}

/*DeleteQualityKeywordsetsNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteQualityKeywordsetsNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteQualityKeywordsetsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsRequestEntityTooLarge creates a DeleteQualityKeywordsetsRequestEntityTooLarge with default headers values
func NewDeleteQualityKeywordsetsRequestEntityTooLarge() *DeleteQualityKeywordsetsRequestEntityTooLarge {
	return &DeleteQualityKeywordsetsRequestEntityTooLarge{}
}

/*DeleteQualityKeywordsetsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteQualityKeywordsetsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteQualityKeywordsetsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsUnsupportedMediaType creates a DeleteQualityKeywordsetsUnsupportedMediaType with default headers values
func NewDeleteQualityKeywordsetsUnsupportedMediaType() *DeleteQualityKeywordsetsUnsupportedMediaType {
	return &DeleteQualityKeywordsetsUnsupportedMediaType{}
}

/*DeleteQualityKeywordsetsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteQualityKeywordsetsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteQualityKeywordsetsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsTooManyRequests creates a DeleteQualityKeywordsetsTooManyRequests with default headers values
func NewDeleteQualityKeywordsetsTooManyRequests() *DeleteQualityKeywordsetsTooManyRequests {
	return &DeleteQualityKeywordsetsTooManyRequests{}
}

/*DeleteQualityKeywordsetsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteQualityKeywordsetsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteQualityKeywordsetsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsInternalServerError creates a DeleteQualityKeywordsetsInternalServerError with default headers values
func NewDeleteQualityKeywordsetsInternalServerError() *DeleteQualityKeywordsetsInternalServerError {
	return &DeleteQualityKeywordsetsInternalServerError{}
}

/*DeleteQualityKeywordsetsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteQualityKeywordsetsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteQualityKeywordsetsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsServiceUnavailable creates a DeleteQualityKeywordsetsServiceUnavailable with default headers values
func NewDeleteQualityKeywordsetsServiceUnavailable() *DeleteQualityKeywordsetsServiceUnavailable {
	return &DeleteQualityKeywordsetsServiceUnavailable{}
}

/*DeleteQualityKeywordsetsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteQualityKeywordsetsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteQualityKeywordsetsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteQualityKeywordsetsGatewayTimeout creates a DeleteQualityKeywordsetsGatewayTimeout with default headers values
func NewDeleteQualityKeywordsetsGatewayTimeout() *DeleteQualityKeywordsetsGatewayTimeout {
	return &DeleteQualityKeywordsetsGatewayTimeout{}
}

/*DeleteQualityKeywordsetsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteQualityKeywordsetsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteQualityKeywordsetsGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/quality/keywordsets][%d] deleteQualityKeywordsetsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteQualityKeywordsetsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteQualityKeywordsetsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
