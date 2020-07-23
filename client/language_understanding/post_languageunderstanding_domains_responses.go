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

// PostLanguageunderstandingDomainsReader is a Reader for the PostLanguageunderstandingDomains structure.
type PostLanguageunderstandingDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLanguageunderstandingDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLanguageunderstandingDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostLanguageunderstandingDomainsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLanguageunderstandingDomainsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLanguageunderstandingDomainsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostLanguageunderstandingDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLanguageunderstandingDomainsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostLanguageunderstandingDomainsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostLanguageunderstandingDomainsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostLanguageunderstandingDomainsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostLanguageunderstandingDomainsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostLanguageunderstandingDomainsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostLanguageunderstandingDomainsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostLanguageunderstandingDomainsOK creates a PostLanguageunderstandingDomainsOK with default headers values
func NewPostLanguageunderstandingDomainsOK() *PostLanguageunderstandingDomainsOK {
	return &PostLanguageunderstandingDomainsOK{}
}

/*PostLanguageunderstandingDomainsOK handles this case with default header values.

successful operation
*/
type PostLanguageunderstandingDomainsOK struct {
	Payload *models.NluDomain
}

func (o *PostLanguageunderstandingDomainsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingDomainsOK) GetPayload() *models.NluDomain {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsCreated creates a PostLanguageunderstandingDomainsCreated with default headers values
func NewPostLanguageunderstandingDomainsCreated() *PostLanguageunderstandingDomainsCreated {
	return &PostLanguageunderstandingDomainsCreated{}
}

/*PostLanguageunderstandingDomainsCreated handles this case with default header values.

PostLanguageunderstandingDomainsCreated post languageunderstanding domains created
*/
type PostLanguageunderstandingDomainsCreated struct {
	Payload *models.NluDomain
}

func (o *PostLanguageunderstandingDomainsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsCreated  %+v", 201, o.Payload)
}

func (o *PostLanguageunderstandingDomainsCreated) GetPayload() *models.NluDomain {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluDomain)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsBadRequest creates a PostLanguageunderstandingDomainsBadRequest with default headers values
func NewPostLanguageunderstandingDomainsBadRequest() *PostLanguageunderstandingDomainsBadRequest {
	return &PostLanguageunderstandingDomainsBadRequest{}
}

/*PostLanguageunderstandingDomainsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLanguageunderstandingDomainsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingDomainsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsUnauthorized creates a PostLanguageunderstandingDomainsUnauthorized with default headers values
func NewPostLanguageunderstandingDomainsUnauthorized() *PostLanguageunderstandingDomainsUnauthorized {
	return &PostLanguageunderstandingDomainsUnauthorized{}
}

/*PostLanguageunderstandingDomainsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLanguageunderstandingDomainsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingDomainsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsForbidden creates a PostLanguageunderstandingDomainsForbidden with default headers values
func NewPostLanguageunderstandingDomainsForbidden() *PostLanguageunderstandingDomainsForbidden {
	return &PostLanguageunderstandingDomainsForbidden{}
}

/*PostLanguageunderstandingDomainsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostLanguageunderstandingDomainsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingDomainsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsNotFound creates a PostLanguageunderstandingDomainsNotFound with default headers values
func NewPostLanguageunderstandingDomainsNotFound() *PostLanguageunderstandingDomainsNotFound {
	return &PostLanguageunderstandingDomainsNotFound{}
}

/*PostLanguageunderstandingDomainsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostLanguageunderstandingDomainsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingDomainsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsRequestEntityTooLarge creates a PostLanguageunderstandingDomainsRequestEntityTooLarge with default headers values
func NewPostLanguageunderstandingDomainsRequestEntityTooLarge() *PostLanguageunderstandingDomainsRequestEntityTooLarge {
	return &PostLanguageunderstandingDomainsRequestEntityTooLarge{}
}

/*PostLanguageunderstandingDomainsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostLanguageunderstandingDomainsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsUnsupportedMediaType creates a PostLanguageunderstandingDomainsUnsupportedMediaType with default headers values
func NewPostLanguageunderstandingDomainsUnsupportedMediaType() *PostLanguageunderstandingDomainsUnsupportedMediaType {
	return &PostLanguageunderstandingDomainsUnsupportedMediaType{}
}

/*PostLanguageunderstandingDomainsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLanguageunderstandingDomainsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsTooManyRequests creates a PostLanguageunderstandingDomainsTooManyRequests with default headers values
func NewPostLanguageunderstandingDomainsTooManyRequests() *PostLanguageunderstandingDomainsTooManyRequests {
	return &PostLanguageunderstandingDomainsTooManyRequests{}
}

/*PostLanguageunderstandingDomainsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostLanguageunderstandingDomainsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingDomainsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsInternalServerError creates a PostLanguageunderstandingDomainsInternalServerError with default headers values
func NewPostLanguageunderstandingDomainsInternalServerError() *PostLanguageunderstandingDomainsInternalServerError {
	return &PostLanguageunderstandingDomainsInternalServerError{}
}

/*PostLanguageunderstandingDomainsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLanguageunderstandingDomainsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingDomainsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsServiceUnavailable creates a PostLanguageunderstandingDomainsServiceUnavailable with default headers values
func NewPostLanguageunderstandingDomainsServiceUnavailable() *PostLanguageunderstandingDomainsServiceUnavailable {
	return &PostLanguageunderstandingDomainsServiceUnavailable{}
}

/*PostLanguageunderstandingDomainsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLanguageunderstandingDomainsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingDomainsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLanguageunderstandingDomainsGatewayTimeout creates a PostLanguageunderstandingDomainsGatewayTimeout with default headers values
func NewPostLanguageunderstandingDomainsGatewayTimeout() *PostLanguageunderstandingDomainsGatewayTimeout {
	return &PostLanguageunderstandingDomainsGatewayTimeout{}
}

/*PostLanguageunderstandingDomainsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostLanguageunderstandingDomainsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostLanguageunderstandingDomainsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingDomainsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
