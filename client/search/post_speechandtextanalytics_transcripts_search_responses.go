// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostSpeechandtextanalyticsTranscriptsSearchReader is a Reader for the PostSpeechandtextanalyticsTranscriptsSearch structure.
type PostSpeechandtextanalyticsTranscriptsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSpeechandtextanalyticsTranscriptsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSpeechandtextanalyticsTranscriptsSearchOK creates a PostSpeechandtextanalyticsTranscriptsSearchOK with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchOK() *PostSpeechandtextanalyticsTranscriptsSearchOK {
	return &PostSpeechandtextanalyticsTranscriptsSearchOK{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchOK handles this case with default header values.

successful operation
*/
type PostSpeechandtextanalyticsTranscriptsSearchOK struct {
	Payload *models.JSONSearchResponse
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchOK  %+v", 200, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchOK) GetPayload() *models.JSONSearchResponse {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchBadRequest creates a PostSpeechandtextanalyticsTranscriptsSearchBadRequest with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchBadRequest() *PostSpeechandtextanalyticsTranscriptsSearchBadRequest {
	return &PostSpeechandtextanalyticsTranscriptsSearchBadRequest{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostSpeechandtextanalyticsTranscriptsSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchUnauthorized creates a PostSpeechandtextanalyticsTranscriptsSearchUnauthorized with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchUnauthorized() *PostSpeechandtextanalyticsTranscriptsSearchUnauthorized {
	return &PostSpeechandtextanalyticsTranscriptsSearchUnauthorized{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostSpeechandtextanalyticsTranscriptsSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchForbidden creates a PostSpeechandtextanalyticsTranscriptsSearchForbidden with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchForbidden() *PostSpeechandtextanalyticsTranscriptsSearchForbidden {
	return &PostSpeechandtextanalyticsTranscriptsSearchForbidden{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostSpeechandtextanalyticsTranscriptsSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchNotFound creates a PostSpeechandtextanalyticsTranscriptsSearchNotFound with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchNotFound() *PostSpeechandtextanalyticsTranscriptsSearchNotFound {
	return &PostSpeechandtextanalyticsTranscriptsSearchNotFound{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostSpeechandtextanalyticsTranscriptsSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchRequestTimeout creates a PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchRequestTimeout() *PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout {
	return &PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge creates a PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge() *PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge {
	return &PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType creates a PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType() *PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType {
	return &PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchTooManyRequests creates a PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchTooManyRequests() *PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests {
	return &PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchInternalServerError creates a PostSpeechandtextanalyticsTranscriptsSearchInternalServerError with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchInternalServerError() *PostSpeechandtextanalyticsTranscriptsSearchInternalServerError {
	return &PostSpeechandtextanalyticsTranscriptsSearchInternalServerError{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostSpeechandtextanalyticsTranscriptsSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable creates a PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable() *PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable {
	return &PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout creates a PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout with default headers values
func NewPostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout() *PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout {
	return &PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout{}
}

/*PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/speechandtextanalytics/transcripts/search][%d] postSpeechandtextanalyticsTranscriptsSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostSpeechandtextanalyticsTranscriptsSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
