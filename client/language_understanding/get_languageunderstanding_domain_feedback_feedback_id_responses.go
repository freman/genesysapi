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

// GetLanguageunderstandingDomainFeedbackFeedbackIDReader is a Reader for the GetLanguageunderstandingDomainFeedbackFeedbackID structure.
type GetLanguageunderstandingDomainFeedbackFeedbackIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDOK creates a GetLanguageunderstandingDomainFeedbackFeedbackIDOK with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDOK() *GetLanguageunderstandingDomainFeedbackFeedbackIDOK {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDOK{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDOK handles this case with default header values.

successful operation
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDOK struct {
	Payload *models.NluFeedbackResponse
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdOK  %+v", 200, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDOK) GetPayload() *models.NluFeedbackResponse {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NluFeedbackResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest creates a GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest() *GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized creates a GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized() *GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDForbidden creates a GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDForbidden() *GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDNotFound creates a GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDNotFound() *GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge creates a GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge() *GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType creates a GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType() *GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests creates a GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests() *GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError creates a GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError() *GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable creates a GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable() *GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout creates a GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout with default headers values
func NewGetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout() *GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout {
	return &GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout{}
}

/*GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languageunderstanding/domains/{domainId}/feedback/{feedbackId}][%d] getLanguageunderstandingDomainFeedbackFeedbackIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguageunderstandingDomainFeedbackFeedbackIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}