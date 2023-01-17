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
	case 408:
		result := NewPostLanguageunderstandingDomainsRequestTimeout()
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

/*
PostLanguageunderstandingDomainsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostLanguageunderstandingDomainsOK struct {
	Payload *models.NluDomain
}

// IsSuccess returns true when this post languageunderstanding domains o k response has a 2xx status code
func (o *PostLanguageunderstandingDomainsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post languageunderstanding domains o k response has a 3xx status code
func (o *PostLanguageunderstandingDomainsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains o k response has a 4xx status code
func (o *PostLanguageunderstandingDomainsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding domains o k response has a 5xx status code
func (o *PostLanguageunderstandingDomainsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains o k response a status code equal to that given
func (o *PostLanguageunderstandingDomainsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostLanguageunderstandingDomainsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsOK  %+v", 200, o.Payload)
}

func (o *PostLanguageunderstandingDomainsOK) String() string {
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

/*
PostLanguageunderstandingDomainsCreated describes a response with status code 201, with default header values.

PostLanguageunderstandingDomainsCreated post languageunderstanding domains created
*/
type PostLanguageunderstandingDomainsCreated struct {
	Payload *models.NluDomain
}

// IsSuccess returns true when this post languageunderstanding domains created response has a 2xx status code
func (o *PostLanguageunderstandingDomainsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post languageunderstanding domains created response has a 3xx status code
func (o *PostLanguageunderstandingDomainsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains created response has a 4xx status code
func (o *PostLanguageunderstandingDomainsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding domains created response has a 5xx status code
func (o *PostLanguageunderstandingDomainsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains created response a status code equal to that given
func (o *PostLanguageunderstandingDomainsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostLanguageunderstandingDomainsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsCreated  %+v", 201, o.Payload)
}

func (o *PostLanguageunderstandingDomainsCreated) String() string {
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

/*
PostLanguageunderstandingDomainsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostLanguageunderstandingDomainsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains bad request response has a 2xx status code
func (o *PostLanguageunderstandingDomainsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains bad request response has a 3xx status code
func (o *PostLanguageunderstandingDomainsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains bad request response has a 4xx status code
func (o *PostLanguageunderstandingDomainsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains bad request response has a 5xx status code
func (o *PostLanguageunderstandingDomainsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains bad request response a status code equal to that given
func (o *PostLanguageunderstandingDomainsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostLanguageunderstandingDomainsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsBadRequest  %+v", 400, o.Payload)
}

func (o *PostLanguageunderstandingDomainsBadRequest) String() string {
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

/*
PostLanguageunderstandingDomainsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostLanguageunderstandingDomainsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains unauthorized response has a 2xx status code
func (o *PostLanguageunderstandingDomainsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains unauthorized response has a 3xx status code
func (o *PostLanguageunderstandingDomainsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains unauthorized response has a 4xx status code
func (o *PostLanguageunderstandingDomainsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains unauthorized response has a 5xx status code
func (o *PostLanguageunderstandingDomainsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains unauthorized response a status code equal to that given
func (o *PostLanguageunderstandingDomainsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostLanguageunderstandingDomainsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostLanguageunderstandingDomainsUnauthorized) String() string {
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

/*
PostLanguageunderstandingDomainsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostLanguageunderstandingDomainsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains forbidden response has a 2xx status code
func (o *PostLanguageunderstandingDomainsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains forbidden response has a 3xx status code
func (o *PostLanguageunderstandingDomainsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains forbidden response has a 4xx status code
func (o *PostLanguageunderstandingDomainsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains forbidden response has a 5xx status code
func (o *PostLanguageunderstandingDomainsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains forbidden response a status code equal to that given
func (o *PostLanguageunderstandingDomainsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostLanguageunderstandingDomainsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsForbidden  %+v", 403, o.Payload)
}

func (o *PostLanguageunderstandingDomainsForbidden) String() string {
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

/*
PostLanguageunderstandingDomainsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostLanguageunderstandingDomainsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains not found response has a 2xx status code
func (o *PostLanguageunderstandingDomainsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains not found response has a 3xx status code
func (o *PostLanguageunderstandingDomainsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains not found response has a 4xx status code
func (o *PostLanguageunderstandingDomainsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains not found response has a 5xx status code
func (o *PostLanguageunderstandingDomainsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains not found response a status code equal to that given
func (o *PostLanguageunderstandingDomainsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostLanguageunderstandingDomainsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsNotFound  %+v", 404, o.Payload)
}

func (o *PostLanguageunderstandingDomainsNotFound) String() string {
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

// NewPostLanguageunderstandingDomainsRequestTimeout creates a PostLanguageunderstandingDomainsRequestTimeout with default headers values
func NewPostLanguageunderstandingDomainsRequestTimeout() *PostLanguageunderstandingDomainsRequestTimeout {
	return &PostLanguageunderstandingDomainsRequestTimeout{}
}

/*
PostLanguageunderstandingDomainsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostLanguageunderstandingDomainsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains request timeout response has a 2xx status code
func (o *PostLanguageunderstandingDomainsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains request timeout response has a 3xx status code
func (o *PostLanguageunderstandingDomainsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains request timeout response has a 4xx status code
func (o *PostLanguageunderstandingDomainsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains request timeout response has a 5xx status code
func (o *PostLanguageunderstandingDomainsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains request timeout response a status code equal to that given
func (o *PostLanguageunderstandingDomainsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostLanguageunderstandingDomainsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingDomainsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostLanguageunderstandingDomainsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostLanguageunderstandingDomainsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
PostLanguageunderstandingDomainsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostLanguageunderstandingDomainsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains request entity too large response has a 2xx status code
func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains request entity too large response has a 3xx status code
func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains request entity too large response has a 4xx status code
func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains request entity too large response has a 5xx status code
func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains request entity too large response a status code equal to that given
func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostLanguageunderstandingDomainsRequestEntityTooLarge) String() string {
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

/*
PostLanguageunderstandingDomainsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostLanguageunderstandingDomainsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains unsupported media type response has a 2xx status code
func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains unsupported media type response has a 3xx status code
func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains unsupported media type response has a 4xx status code
func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains unsupported media type response has a 5xx status code
func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains unsupported media type response a status code equal to that given
func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostLanguageunderstandingDomainsUnsupportedMediaType) String() string {
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

/*
PostLanguageunderstandingDomainsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostLanguageunderstandingDomainsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains too many requests response has a 2xx status code
func (o *PostLanguageunderstandingDomainsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains too many requests response has a 3xx status code
func (o *PostLanguageunderstandingDomainsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains too many requests response has a 4xx status code
func (o *PostLanguageunderstandingDomainsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post languageunderstanding domains too many requests response has a 5xx status code
func (o *PostLanguageunderstandingDomainsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post languageunderstanding domains too many requests response a status code equal to that given
func (o *PostLanguageunderstandingDomainsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostLanguageunderstandingDomainsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostLanguageunderstandingDomainsTooManyRequests) String() string {
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

/*
PostLanguageunderstandingDomainsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostLanguageunderstandingDomainsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains internal server error response has a 2xx status code
func (o *PostLanguageunderstandingDomainsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains internal server error response has a 3xx status code
func (o *PostLanguageunderstandingDomainsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains internal server error response has a 4xx status code
func (o *PostLanguageunderstandingDomainsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding domains internal server error response has a 5xx status code
func (o *PostLanguageunderstandingDomainsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding domains internal server error response a status code equal to that given
func (o *PostLanguageunderstandingDomainsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostLanguageunderstandingDomainsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostLanguageunderstandingDomainsInternalServerError) String() string {
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

/*
PostLanguageunderstandingDomainsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostLanguageunderstandingDomainsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains service unavailable response has a 2xx status code
func (o *PostLanguageunderstandingDomainsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains service unavailable response has a 3xx status code
func (o *PostLanguageunderstandingDomainsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains service unavailable response has a 4xx status code
func (o *PostLanguageunderstandingDomainsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding domains service unavailable response has a 5xx status code
func (o *PostLanguageunderstandingDomainsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding domains service unavailable response a status code equal to that given
func (o *PostLanguageunderstandingDomainsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostLanguageunderstandingDomainsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostLanguageunderstandingDomainsServiceUnavailable) String() string {
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

/*
PostLanguageunderstandingDomainsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostLanguageunderstandingDomainsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post languageunderstanding domains gateway timeout response has a 2xx status code
func (o *PostLanguageunderstandingDomainsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post languageunderstanding domains gateway timeout response has a 3xx status code
func (o *PostLanguageunderstandingDomainsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post languageunderstanding domains gateway timeout response has a 4xx status code
func (o *PostLanguageunderstandingDomainsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post languageunderstanding domains gateway timeout response has a 5xx status code
func (o *PostLanguageunderstandingDomainsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post languageunderstanding domains gateway timeout response a status code equal to that given
func (o *PostLanguageunderstandingDomainsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostLanguageunderstandingDomainsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/languageunderstanding/domains][%d] postLanguageunderstandingDomainsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostLanguageunderstandingDomainsGatewayTimeout) String() string {
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
