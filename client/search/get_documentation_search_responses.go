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

// GetDocumentationSearchReader is a Reader for the GetDocumentationSearch structure.
type GetDocumentationSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDocumentationSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDocumentationSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDocumentationSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetDocumentationSearchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetDocumentationSearchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDocumentationSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetDocumentationSearchRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetDocumentationSearchUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetDocumentationSearchTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDocumentationSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetDocumentationSearchServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetDocumentationSearchGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDocumentationSearchOK creates a GetDocumentationSearchOK with default headers values
func NewGetDocumentationSearchOK() *GetDocumentationSearchOK {
	return &GetDocumentationSearchOK{}
}

/*GetDocumentationSearchOK handles this case with default header values.

successful operation
*/
type GetDocumentationSearchOK struct {
	Payload *models.DocumentationSearchResponse
}

func (o *GetDocumentationSearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchOK  %+v", 200, o.Payload)
}

func (o *GetDocumentationSearchOK) GetPayload() *models.DocumentationSearchResponse {
	return o.Payload
}

func (o *GetDocumentationSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentationSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchBadRequest creates a GetDocumentationSearchBadRequest with default headers values
func NewGetDocumentationSearchBadRequest() *GetDocumentationSearchBadRequest {
	return &GetDocumentationSearchBadRequest{}
}

/*GetDocumentationSearchBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetDocumentationSearchBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchBadRequest  %+v", 400, o.Payload)
}

func (o *GetDocumentationSearchBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchUnauthorized creates a GetDocumentationSearchUnauthorized with default headers values
func NewGetDocumentationSearchUnauthorized() *GetDocumentationSearchUnauthorized {
	return &GetDocumentationSearchUnauthorized{}
}

/*GetDocumentationSearchUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetDocumentationSearchUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchUnauthorized  %+v", 401, o.Payload)
}

func (o *GetDocumentationSearchUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchForbidden creates a GetDocumentationSearchForbidden with default headers values
func NewGetDocumentationSearchForbidden() *GetDocumentationSearchForbidden {
	return &GetDocumentationSearchForbidden{}
}

/*GetDocumentationSearchForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetDocumentationSearchForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchForbidden  %+v", 403, o.Payload)
}

func (o *GetDocumentationSearchForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchNotFound creates a GetDocumentationSearchNotFound with default headers values
func NewGetDocumentationSearchNotFound() *GetDocumentationSearchNotFound {
	return &GetDocumentationSearchNotFound{}
}

/*GetDocumentationSearchNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetDocumentationSearchNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchNotFound  %+v", 404, o.Payload)
}

func (o *GetDocumentationSearchNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchRequestEntityTooLarge creates a GetDocumentationSearchRequestEntityTooLarge with default headers values
func NewGetDocumentationSearchRequestEntityTooLarge() *GetDocumentationSearchRequestEntityTooLarge {
	return &GetDocumentationSearchRequestEntityTooLarge{}
}

/*GetDocumentationSearchRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetDocumentationSearchRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetDocumentationSearchRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchUnsupportedMediaType creates a GetDocumentationSearchUnsupportedMediaType with default headers values
func NewGetDocumentationSearchUnsupportedMediaType() *GetDocumentationSearchUnsupportedMediaType {
	return &GetDocumentationSearchUnsupportedMediaType{}
}

/*GetDocumentationSearchUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetDocumentationSearchUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetDocumentationSearchUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchTooManyRequests creates a GetDocumentationSearchTooManyRequests with default headers values
func NewGetDocumentationSearchTooManyRequests() *GetDocumentationSearchTooManyRequests {
	return &GetDocumentationSearchTooManyRequests{}
}

/*GetDocumentationSearchTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetDocumentationSearchTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetDocumentationSearchTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchInternalServerError creates a GetDocumentationSearchInternalServerError with default headers values
func NewGetDocumentationSearchInternalServerError() *GetDocumentationSearchInternalServerError {
	return &GetDocumentationSearchInternalServerError{}
}

/*GetDocumentationSearchInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetDocumentationSearchInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDocumentationSearchInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchServiceUnavailable creates a GetDocumentationSearchServiceUnavailable with default headers values
func NewGetDocumentationSearchServiceUnavailable() *GetDocumentationSearchServiceUnavailable {
	return &GetDocumentationSearchServiceUnavailable{}
}

/*GetDocumentationSearchServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetDocumentationSearchServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDocumentationSearchServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDocumentationSearchGatewayTimeout creates a GetDocumentationSearchGatewayTimeout with default headers values
func NewGetDocumentationSearchGatewayTimeout() *GetDocumentationSearchGatewayTimeout {
	return &GetDocumentationSearchGatewayTimeout{}
}

/*GetDocumentationSearchGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetDocumentationSearchGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetDocumentationSearchGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/documentation/search][%d] getDocumentationSearchGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDocumentationSearchGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetDocumentationSearchGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
