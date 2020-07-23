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

// GetLanguageunderstandingDomainsReader is a Reader for the GetLanguageunderstandingDomains structure.
type GetLanguageunderstandingDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainsOK creates a GetLanguageunderstandingDomainsOK with default headers values
func NewGetLanguageunderstandingDomainsOK() *GetLanguageunderstandingDomainsOK {
	return &GetLanguageunderstandingDomainsOK{}
}

/*GetLanguageunderstandingDomainsOK handles this case with default header values.

successful operation
*/
type GetLanguageunderstandingDomainsOK struct {
	Payload *models.NluDomainListing
}

func (o *GetLanguageunderstandingDomainsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainsOK) GetPayload() *models.NluDomainListing {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomainListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsBadRequest creates a GetLanguageunderstandingDomainsBadRequest with default headers values
func NewGetLanguageunderstandingDomainsBadRequest() *GetLanguageunderstandingDomainsBadRequest {
	return &GetLanguageunderstandingDomainsBadRequest{}
}

/*GetLanguageunderstandingDomainsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsUnauthorized creates a GetLanguageunderstandingDomainsUnauthorized with default headers values
func NewGetLanguageunderstandingDomainsUnauthorized() *GetLanguageunderstandingDomainsUnauthorized {
	return &GetLanguageunderstandingDomainsUnauthorized{}
}

/*GetLanguageunderstandingDomainsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsForbidden creates a GetLanguageunderstandingDomainsForbidden with default headers values
func NewGetLanguageunderstandingDomainsForbidden() *GetLanguageunderstandingDomainsForbidden {
	return &GetLanguageunderstandingDomainsForbidden{}
}

/*GetLanguageunderstandingDomainsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsNotFound creates a GetLanguageunderstandingDomainsNotFound with default headers values
func NewGetLanguageunderstandingDomainsNotFound() *GetLanguageunderstandingDomainsNotFound {
	return &GetLanguageunderstandingDomainsNotFound{}
}

/*GetLanguageunderstandingDomainsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsRequestEntityTooLarge creates a GetLanguageunderstandingDomainsRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainsRequestEntityTooLarge() *GetLanguageunderstandingDomainsRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainsRequestEntityTooLarge{}
}

/*GetLanguageunderstandingDomainsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingDomainsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsUnsupportedMediaType creates a GetLanguageunderstandingDomainsUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainsUnsupportedMediaType() *GetLanguageunderstandingDomainsUnsupportedMediaType {
	return &GetLanguageunderstandingDomainsUnsupportedMediaType{}
}

/*GetLanguageunderstandingDomainsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsTooManyRequests creates a GetLanguageunderstandingDomainsTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainsTooManyRequests() *GetLanguageunderstandingDomainsTooManyRequests {
	return &GetLanguageunderstandingDomainsTooManyRequests{}
}

/*GetLanguageunderstandingDomainsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLanguageunderstandingDomainsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsInternalServerError creates a GetLanguageunderstandingDomainsInternalServerError with default headers values
func NewGetLanguageunderstandingDomainsInternalServerError() *GetLanguageunderstandingDomainsInternalServerError {
	return &GetLanguageunderstandingDomainsInternalServerError{}
}

/*GetLanguageunderstandingDomainsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsServiceUnavailable creates a GetLanguageunderstandingDomainsServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainsServiceUnavailable() *GetLanguageunderstandingDomainsServiceUnavailable {
	return &GetLanguageunderstandingDomainsServiceUnavailable{}
}

/*GetLanguageunderstandingDomainsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainsGatewayTimeout creates a GetLanguageunderstandingDomainsGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainsGatewayTimeout() *GetLanguageunderstandingDomainsGatewayTimeout {
	return &GetLanguageunderstandingDomainsGatewayTimeout{}
}

/*GetLanguageunderstandingDomainsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains][%d] getLanguageunderstandingDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
