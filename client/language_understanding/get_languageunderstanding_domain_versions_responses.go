// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguageunderstandingDomainVersionsReader is a Reader for the GetLanguageunderstandingDomainVersions structure.
type GetLanguageunderstandingDomainVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainVersionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainVersionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainVersionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainVersionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainVersionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainVersionsOK creates a GetLanguageunderstandingDomainVersionsOK with default headers values
func NewGetLanguageunderstandingDomainVersionsOK() *GetLanguageunderstandingDomainVersionsOK {
	return &GetLanguageunderstandingDomainVersionsOK{}
}

/*GetLanguageunderstandingDomainVersionsOK handles this case with default header values.

successful operation
*/
type GetLanguageunderstandingDomainVersionsOK struct {
	Payload *models.NluDomainVersionListing
}

func (o *GetLanguageunderstandingDomainVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsOK) GetPayload() *models.NluDomainVersionListing {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomainVersionListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsBadRequest creates a GetLanguageunderstandingDomainVersionsBadRequest with default headers values
func NewGetLanguageunderstandingDomainVersionsBadRequest() *GetLanguageunderstandingDomainVersionsBadRequest {
	return &GetLanguageunderstandingDomainVersionsBadRequest{}
}

/*GetLanguageunderstandingDomainVersionsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainVersionsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsUnauthorized creates a GetLanguageunderstandingDomainVersionsUnauthorized with default headers values
func NewGetLanguageunderstandingDomainVersionsUnauthorized() *GetLanguageunderstandingDomainVersionsUnauthorized {
	return &GetLanguageunderstandingDomainVersionsUnauthorized{}
}

/*GetLanguageunderstandingDomainVersionsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainVersionsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsForbidden creates a GetLanguageunderstandingDomainVersionsForbidden with default headers values
func NewGetLanguageunderstandingDomainVersionsForbidden() *GetLanguageunderstandingDomainVersionsForbidden {
	return &GetLanguageunderstandingDomainVersionsForbidden{}
}

/*GetLanguageunderstandingDomainVersionsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainVersionsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsNotFound creates a GetLanguageunderstandingDomainVersionsNotFound with default headers values
func NewGetLanguageunderstandingDomainVersionsNotFound() *GetLanguageunderstandingDomainVersionsNotFound {
	return &GetLanguageunderstandingDomainVersionsNotFound{}
}

/*GetLanguageunderstandingDomainVersionsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainVersionsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsRequestEntityTooLarge creates a GetLanguageunderstandingDomainVersionsRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainVersionsRequestEntityTooLarge() *GetLanguageunderstandingDomainVersionsRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainVersionsRequestEntityTooLarge{}
}

/*GetLanguageunderstandingDomainVersionsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingDomainVersionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsUnsupportedMediaType creates a GetLanguageunderstandingDomainVersionsUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainVersionsUnsupportedMediaType() *GetLanguageunderstandingDomainVersionsUnsupportedMediaType {
	return &GetLanguageunderstandingDomainVersionsUnsupportedMediaType{}
}

/*GetLanguageunderstandingDomainVersionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainVersionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsTooManyRequests creates a GetLanguageunderstandingDomainVersionsTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainVersionsTooManyRequests() *GetLanguageunderstandingDomainVersionsTooManyRequests {
	return &GetLanguageunderstandingDomainVersionsTooManyRequests{}
}

/*GetLanguageunderstandingDomainVersionsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLanguageunderstandingDomainVersionsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsInternalServerError creates a GetLanguageunderstandingDomainVersionsInternalServerError with default headers values
func NewGetLanguageunderstandingDomainVersionsInternalServerError() *GetLanguageunderstandingDomainVersionsInternalServerError {
	return &GetLanguageunderstandingDomainVersionsInternalServerError{}
}

/*GetLanguageunderstandingDomainVersionsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainVersionsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsServiceUnavailable creates a GetLanguageunderstandingDomainVersionsServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainVersionsServiceUnavailable() *GetLanguageunderstandingDomainVersionsServiceUnavailable {
	return &GetLanguageunderstandingDomainVersionsServiceUnavailable{}
}

/*GetLanguageunderstandingDomainVersionsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainVersionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainVersionsGatewayTimeout creates a GetLanguageunderstandingDomainVersionsGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainVersionsGatewayTimeout() *GetLanguageunderstandingDomainVersionsGatewayTimeout {
	return &GetLanguageunderstandingDomainVersionsGatewayTimeout{}
}

/*GetLanguageunderstandingDomainVersionsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainVersionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainVersionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/versions][%d] getLanguageunderstandingDomainVersionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainVersionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainVersionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
