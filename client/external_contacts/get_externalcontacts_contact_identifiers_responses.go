// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetExternalcontactsContactIdentifiersReader is a Reader for the GetExternalcontactsContactIdentifiers structure.
type GetExternalcontactsContactIdentifiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsContactIdentifiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsContactIdentifiersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsContactIdentifiersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsContactIdentifiersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsContactIdentifiersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsContactIdentifiersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsContactIdentifiersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsContactIdentifiersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsContactIdentifiersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsContactIdentifiersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsContactIdentifiersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsContactIdentifiersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsContactIdentifiersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsContactIdentifiersOK creates a GetExternalcontactsContactIdentifiersOK with default headers values
func NewGetExternalcontactsContactIdentifiersOK() *GetExternalcontactsContactIdentifiersOK {
	return &GetExternalcontactsContactIdentifiersOK{}
}

/*
GetExternalcontactsContactIdentifiersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsContactIdentifiersOK struct {
	Payload *models.EntityListing
}

// IsSuccess returns true when this get externalcontacts contact identifiers o k response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts contact identifiers o k response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers o k response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact identifiers o k response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers o k response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsContactIdentifiersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersOK) GetPayload() *models.EntityListing {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersBadRequest creates a GetExternalcontactsContactIdentifiersBadRequest with default headers values
func NewGetExternalcontactsContactIdentifiersBadRequest() *GetExternalcontactsContactIdentifiersBadRequest {
	return &GetExternalcontactsContactIdentifiersBadRequest{}
}

/*
GetExternalcontactsContactIdentifiersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsContactIdentifiersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers bad request response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers bad request response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers bad request response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers bad request response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers bad request response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsContactIdentifiersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersUnauthorized creates a GetExternalcontactsContactIdentifiersUnauthorized with default headers values
func NewGetExternalcontactsContactIdentifiersUnauthorized() *GetExternalcontactsContactIdentifiersUnauthorized {
	return &GetExternalcontactsContactIdentifiersUnauthorized{}
}

/*
GetExternalcontactsContactIdentifiersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsContactIdentifiersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers unauthorized response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers unauthorized response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers unauthorized response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers unauthorized response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers unauthorized response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsContactIdentifiersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersForbidden creates a GetExternalcontactsContactIdentifiersForbidden with default headers values
func NewGetExternalcontactsContactIdentifiersForbidden() *GetExternalcontactsContactIdentifiersForbidden {
	return &GetExternalcontactsContactIdentifiersForbidden{}
}

/*
GetExternalcontactsContactIdentifiersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsContactIdentifiersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers forbidden response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers forbidden response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers forbidden response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers forbidden response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers forbidden response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsContactIdentifiersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersNotFound creates a GetExternalcontactsContactIdentifiersNotFound with default headers values
func NewGetExternalcontactsContactIdentifiersNotFound() *GetExternalcontactsContactIdentifiersNotFound {
	return &GetExternalcontactsContactIdentifiersNotFound{}
}

/*
GetExternalcontactsContactIdentifiersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsContactIdentifiersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers not found response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers not found response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers not found response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers not found response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers not found response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsContactIdentifiersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersRequestTimeout creates a GetExternalcontactsContactIdentifiersRequestTimeout with default headers values
func NewGetExternalcontactsContactIdentifiersRequestTimeout() *GetExternalcontactsContactIdentifiersRequestTimeout {
	return &GetExternalcontactsContactIdentifiersRequestTimeout{}
}

/*
GetExternalcontactsContactIdentifiersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsContactIdentifiersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers request timeout response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers request timeout response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers request timeout response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers request timeout response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers request timeout response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsContactIdentifiersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersRequestEntityTooLarge creates a GetExternalcontactsContactIdentifiersRequestEntityTooLarge with default headers values
func NewGetExternalcontactsContactIdentifiersRequestEntityTooLarge() *GetExternalcontactsContactIdentifiersRequestEntityTooLarge {
	return &GetExternalcontactsContactIdentifiersRequestEntityTooLarge{}
}

/*
GetExternalcontactsContactIdentifiersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetExternalcontactsContactIdentifiersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers request entity too large response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers request entity too large response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers request entity too large response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers request entity too large response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers request entity too large response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersUnsupportedMediaType creates a GetExternalcontactsContactIdentifiersUnsupportedMediaType with default headers values
func NewGetExternalcontactsContactIdentifiersUnsupportedMediaType() *GetExternalcontactsContactIdentifiersUnsupportedMediaType {
	return &GetExternalcontactsContactIdentifiersUnsupportedMediaType{}
}

/*
GetExternalcontactsContactIdentifiersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsContactIdentifiersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers unsupported media type response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers unsupported media type response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers unsupported media type response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers unsupported media type response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers unsupported media type response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersTooManyRequests creates a GetExternalcontactsContactIdentifiersTooManyRequests with default headers values
func NewGetExternalcontactsContactIdentifiersTooManyRequests() *GetExternalcontactsContactIdentifiersTooManyRequests {
	return &GetExternalcontactsContactIdentifiersTooManyRequests{}
}

/*
GetExternalcontactsContactIdentifiersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsContactIdentifiersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers too many requests response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers too many requests response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers too many requests response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact identifiers too many requests response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact identifiers too many requests response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsContactIdentifiersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersInternalServerError creates a GetExternalcontactsContactIdentifiersInternalServerError with default headers values
func NewGetExternalcontactsContactIdentifiersInternalServerError() *GetExternalcontactsContactIdentifiersInternalServerError {
	return &GetExternalcontactsContactIdentifiersInternalServerError{}
}

/*
GetExternalcontactsContactIdentifiersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsContactIdentifiersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers internal server error response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers internal server error response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers internal server error response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact identifiers internal server error response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts contact identifiers internal server error response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsContactIdentifiersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersServiceUnavailable creates a GetExternalcontactsContactIdentifiersServiceUnavailable with default headers values
func NewGetExternalcontactsContactIdentifiersServiceUnavailable() *GetExternalcontactsContactIdentifiersServiceUnavailable {
	return &GetExternalcontactsContactIdentifiersServiceUnavailable{}
}

/*
GetExternalcontactsContactIdentifiersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsContactIdentifiersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers service unavailable response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers service unavailable response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers service unavailable response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact identifiers service unavailable response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts contact identifiers service unavailable response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactIdentifiersGatewayTimeout creates a GetExternalcontactsContactIdentifiersGatewayTimeout with default headers values
func NewGetExternalcontactsContactIdentifiersGatewayTimeout() *GetExternalcontactsContactIdentifiersGatewayTimeout {
	return &GetExternalcontactsContactIdentifiersGatewayTimeout{}
}

/*
GetExternalcontactsContactIdentifiersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsContactIdentifiersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact identifiers gateway timeout response has a 2xx status code
func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact identifiers gateway timeout response has a 3xx status code
func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact identifiers gateway timeout response has a 4xx status code
func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact identifiers gateway timeout response has a 5xx status code
func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts contact identifiers gateway timeout response a status code equal to that given
func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/identifiers][%d] getExternalcontactsContactIdentifiersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactIdentifiersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
