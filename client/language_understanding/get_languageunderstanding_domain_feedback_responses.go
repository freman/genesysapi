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

// GetLanguageunderstandingDomainFeedbackReader is a Reader for the GetLanguageunderstandingDomainFeedback structure.
type GetLanguageunderstandingDomainFeedbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainFeedbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainFeedbackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainFeedbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainFeedbackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainFeedbackForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainFeedbackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainFeedbackRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainFeedbackUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainFeedbackTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainFeedbackInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainFeedbackServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainFeedbackGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainFeedbackOK creates a GetLanguageunderstandingDomainFeedbackOK with default headers values
func NewGetLanguageunderstandingDomainFeedbackOK() *GetLanguageunderstandingDomainFeedbackOK {
	return &GetLanguageunderstandingDomainFeedbackOK{}
}

/*GetLanguageunderstandingDomainFeedbackOK handles this case with default header values.

successful operation
*/
type GetLanguageunderstandingDomainFeedbackOK struct {
	Payload *models.NluFeedbackListing
}

func (o *GetLanguageunderstandingDomainFeedbackOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackOK) GetPayload() *models.NluFeedbackListing {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluFeedbackListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackBadRequest creates a GetLanguageunderstandingDomainFeedbackBadRequest with default headers values
func NewGetLanguageunderstandingDomainFeedbackBadRequest() *GetLanguageunderstandingDomainFeedbackBadRequest {
	return &GetLanguageunderstandingDomainFeedbackBadRequest{}
}

/*GetLanguageunderstandingDomainFeedbackBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainFeedbackBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackUnauthorized creates a GetLanguageunderstandingDomainFeedbackUnauthorized with default headers values
func NewGetLanguageunderstandingDomainFeedbackUnauthorized() *GetLanguageunderstandingDomainFeedbackUnauthorized {
	return &GetLanguageunderstandingDomainFeedbackUnauthorized{}
}

/*GetLanguageunderstandingDomainFeedbackUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainFeedbackUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackForbidden creates a GetLanguageunderstandingDomainFeedbackForbidden with default headers values
func NewGetLanguageunderstandingDomainFeedbackForbidden() *GetLanguageunderstandingDomainFeedbackForbidden {
	return &GetLanguageunderstandingDomainFeedbackForbidden{}
}

/*GetLanguageunderstandingDomainFeedbackForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainFeedbackForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackNotFound creates a GetLanguageunderstandingDomainFeedbackNotFound with default headers values
func NewGetLanguageunderstandingDomainFeedbackNotFound() *GetLanguageunderstandingDomainFeedbackNotFound {
	return &GetLanguageunderstandingDomainFeedbackNotFound{}
}

/*GetLanguageunderstandingDomainFeedbackNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainFeedbackNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackRequestEntityTooLarge creates a GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainFeedbackRequestEntityTooLarge() *GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge{}
}

/*GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackUnsupportedMediaType creates a GetLanguageunderstandingDomainFeedbackUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainFeedbackUnsupportedMediaType() *GetLanguageunderstandingDomainFeedbackUnsupportedMediaType {
	return &GetLanguageunderstandingDomainFeedbackUnsupportedMediaType{}
}

/*GetLanguageunderstandingDomainFeedbackUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainFeedbackUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackTooManyRequests creates a GetLanguageunderstandingDomainFeedbackTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainFeedbackTooManyRequests() *GetLanguageunderstandingDomainFeedbackTooManyRequests {
	return &GetLanguageunderstandingDomainFeedbackTooManyRequests{}
}

/*GetLanguageunderstandingDomainFeedbackTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLanguageunderstandingDomainFeedbackTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackInternalServerError creates a GetLanguageunderstandingDomainFeedbackInternalServerError with default headers values
func NewGetLanguageunderstandingDomainFeedbackInternalServerError() *GetLanguageunderstandingDomainFeedbackInternalServerError {
	return &GetLanguageunderstandingDomainFeedbackInternalServerError{}
}

/*GetLanguageunderstandingDomainFeedbackInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainFeedbackInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackServiceUnavailable creates a GetLanguageunderstandingDomainFeedbackServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainFeedbackServiceUnavailable() *GetLanguageunderstandingDomainFeedbackServiceUnavailable {
	return &GetLanguageunderstandingDomainFeedbackServiceUnavailable{}
}

/*GetLanguageunderstandingDomainFeedbackServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainFeedbackServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackGatewayTimeout creates a GetLanguageunderstandingDomainFeedbackGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainFeedbackGatewayTimeout() *GetLanguageunderstandingDomainFeedbackGatewayTimeout {
	return &GetLanguageunderstandingDomainFeedbackGatewayTimeout{}
}

/*GetLanguageunderstandingDomainFeedbackGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainFeedbackGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback][%d] getLanguageunderstandingDomainFeedbackGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
