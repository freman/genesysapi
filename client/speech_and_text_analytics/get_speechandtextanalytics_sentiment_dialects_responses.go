// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetSpeechandtextanalyticsSentimentDialectsReader is a Reader for the GetSpeechandtextanalyticsSentimentDialects structure.
type GetSpeechandtextanalyticsSentimentDialectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsSentimentDialectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsSentimentDialectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsSentimentDialectsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsSentimentDialectsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsSentimentDialectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsSentimentDialectsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsSentimentDialectsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsSentimentDialectsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsSentimentDialectsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsSentimentDialectsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsSentimentDialectsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsSentimentDialectsOK creates a GetSpeechandtextanalyticsSentimentDialectsOK with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsOK() *GetSpeechandtextanalyticsSentimentDialectsOK {
	return &GetSpeechandtextanalyticsSentimentDialectsOK{}
}

/*GetSpeechandtextanalyticsSentimentDialectsOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsSentimentDialectsOK struct {
	Payload *models.EntityListing
}

func (o *GetSpeechandtextanalyticsSentimentDialectsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsOK) GetPayload() *models.EntityListing {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsBadRequest creates a GetSpeechandtextanalyticsSentimentDialectsBadRequest with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsBadRequest() *GetSpeechandtextanalyticsSentimentDialectsBadRequest {
	return &GetSpeechandtextanalyticsSentimentDialectsBadRequest{}
}

/*GetSpeechandtextanalyticsSentimentDialectsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsSentimentDialectsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsUnauthorized creates a GetSpeechandtextanalyticsSentimentDialectsUnauthorized with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsUnauthorized() *GetSpeechandtextanalyticsSentimentDialectsUnauthorized {
	return &GetSpeechandtextanalyticsSentimentDialectsUnauthorized{}
}

/*GetSpeechandtextanalyticsSentimentDialectsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsSentimentDialectsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsForbidden creates a GetSpeechandtextanalyticsSentimentDialectsForbidden with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsForbidden() *GetSpeechandtextanalyticsSentimentDialectsForbidden {
	return &GetSpeechandtextanalyticsSentimentDialectsForbidden{}
}

/*GetSpeechandtextanalyticsSentimentDialectsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsSentimentDialectsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsNotFound creates a GetSpeechandtextanalyticsSentimentDialectsNotFound with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsNotFound() *GetSpeechandtextanalyticsSentimentDialectsNotFound {
	return &GetSpeechandtextanalyticsSentimentDialectsNotFound{}
}

/*GetSpeechandtextanalyticsSentimentDialectsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsSentimentDialectsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsRequestTimeout creates a GetSpeechandtextanalyticsSentimentDialectsRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsRequestTimeout() *GetSpeechandtextanalyticsSentimentDialectsRequestTimeout {
	return &GetSpeechandtextanalyticsSentimentDialectsRequestTimeout{}
}

/*GetSpeechandtextanalyticsSentimentDialectsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsSentimentDialectsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge creates a GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge() *GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType creates a GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType() *GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType {
	return &GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsTooManyRequests creates a GetSpeechandtextanalyticsSentimentDialectsTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsTooManyRequests() *GetSpeechandtextanalyticsSentimentDialectsTooManyRequests {
	return &GetSpeechandtextanalyticsSentimentDialectsTooManyRequests{}
}

/*GetSpeechandtextanalyticsSentimentDialectsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsSentimentDialectsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsInternalServerError creates a GetSpeechandtextanalyticsSentimentDialectsInternalServerError with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsInternalServerError() *GetSpeechandtextanalyticsSentimentDialectsInternalServerError {
	return &GetSpeechandtextanalyticsSentimentDialectsInternalServerError{}
}

/*GetSpeechandtextanalyticsSentimentDialectsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsSentimentDialectsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsServiceUnavailable creates a GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsServiceUnavailable() *GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable {
	return &GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable{}
}

/*GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsSentimentDialectsGatewayTimeout creates a GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsSentimentDialectsGatewayTimeout() *GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout {
	return &GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout{}
}

/*GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/sentiment/dialects][%d] getSpeechandtextanalyticsSentimentDialectsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsSentimentDialectsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
