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

// PostDocumentationSearchReader is a Reader for the PostDocumentationSearch structure.
type PostDocumentationSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDocumentationSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDocumentationSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDocumentationSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostDocumentationSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostDocumentationSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostDocumentationSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostDocumentationSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostDocumentationSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostDocumentationSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDocumentationSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostDocumentationSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostDocumentationSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDocumentationSearchOK creates a PostDocumentationSearchOK with default headers values
func NewPostDocumentationSearchOK() *PostDocumentationSearchOK {
	return &PostDocumentationSearchOK{}
}

/*PostDocumentationSearchOK handles this case with default header values.

successful operation
*/
type PostDocumentationSearchOK struct {
	Payload *models.DocumentationSearchResponse
}

func (o *PostDocumentationSearchOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchOK  %+v", 200, o.Payload)
}

func (o *PostDocumentationSearchOK) GetPayload() *models.DocumentationSearchResponse {
	return o.Payload
}

func (o *PostDocumentationSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentationSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchBadRequest creates a PostDocumentationSearchBadRequest with default headers values
func NewPostDocumentationSearchBadRequest() *PostDocumentationSearchBadRequest {
	return &PostDocumentationSearchBadRequest{}
}

/*PostDocumentationSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostDocumentationSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchBadRequest  %+v", 400, o.Payload)
}

func (o *PostDocumentationSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchUnauthorized creates a PostDocumentationSearchUnauthorized with default headers values
func NewPostDocumentationSearchUnauthorized() *PostDocumentationSearchUnauthorized {
	return &PostDocumentationSearchUnauthorized{}
}

/*PostDocumentationSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostDocumentationSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *PostDocumentationSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchForbidden creates a PostDocumentationSearchForbidden with default headers values
func NewPostDocumentationSearchForbidden() *PostDocumentationSearchForbidden {
	return &PostDocumentationSearchForbidden{}
}

/*PostDocumentationSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostDocumentationSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchForbidden  %+v", 403, o.Payload)
}

func (o *PostDocumentationSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchNotFound creates a PostDocumentationSearchNotFound with default headers values
func NewPostDocumentationSearchNotFound() *PostDocumentationSearchNotFound {
	return &PostDocumentationSearchNotFound{}
}

/*PostDocumentationSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostDocumentationSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchNotFound  %+v", 404, o.Payload)
}

func (o *PostDocumentationSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchRequestEntityTooLarge creates a PostDocumentationSearchRequestEntityTooLarge with default headers values
func NewPostDocumentationSearchRequestEntityTooLarge() *PostDocumentationSearchRequestEntityTooLarge {
	return &PostDocumentationSearchRequestEntityTooLarge{}
}

/*PostDocumentationSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostDocumentationSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostDocumentationSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchUnsupportedMediaType creates a PostDocumentationSearchUnsupportedMediaType with default headers values
func NewPostDocumentationSearchUnsupportedMediaType() *PostDocumentationSearchUnsupportedMediaType {
	return &PostDocumentationSearchUnsupportedMediaType{}
}

/*PostDocumentationSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostDocumentationSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostDocumentationSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchTooManyRequests creates a PostDocumentationSearchTooManyRequests with default headers values
func NewPostDocumentationSearchTooManyRequests() *PostDocumentationSearchTooManyRequests {
	return &PostDocumentationSearchTooManyRequests{}
}

/*PostDocumentationSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostDocumentationSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostDocumentationSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchInternalServerError creates a PostDocumentationSearchInternalServerError with default headers values
func NewPostDocumentationSearchInternalServerError() *PostDocumentationSearchInternalServerError {
	return &PostDocumentationSearchInternalServerError{}
}

/*PostDocumentationSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostDocumentationSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *PostDocumentationSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchServiceUnavailable creates a PostDocumentationSearchServiceUnavailable with default headers values
func NewPostDocumentationSearchServiceUnavailable() *PostDocumentationSearchServiceUnavailable {
	return &PostDocumentationSearchServiceUnavailable{}
}

/*PostDocumentationSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostDocumentationSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostDocumentationSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDocumentationSearchGatewayTimeout creates a PostDocumentationSearchGatewayTimeout with default headers values
func NewPostDocumentationSearchGatewayTimeout() *PostDocumentationSearchGatewayTimeout {
	return &PostDocumentationSearchGatewayTimeout{}
}

/*PostDocumentationSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostDocumentationSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostDocumentationSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/documentation/search][%d] postDocumentationSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostDocumentationSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostDocumentationSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
