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

// GetExternalcontactsContactJourneySessionsReader is a Reader for the GetExternalcontactsContactJourneySessions structure.
type GetExternalcontactsContactJourneySessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsContactJourneySessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsContactJourneySessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsContactJourneySessionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsContactJourneySessionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsContactJourneySessionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsContactJourneySessionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsContactJourneySessionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsContactJourneySessionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsContactJourneySessionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsContactJourneySessionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsContactJourneySessionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsContactJourneySessionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsContactJourneySessionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsContactJourneySessionsOK creates a GetExternalcontactsContactJourneySessionsOK with default headers values
func NewGetExternalcontactsContactJourneySessionsOK() *GetExternalcontactsContactJourneySessionsOK {
	return &GetExternalcontactsContactJourneySessionsOK{}
}

/*
GetExternalcontactsContactJourneySessionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetExternalcontactsContactJourneySessionsOK struct {
	Payload *models.SessionListing
}

// IsSuccess returns true when this get externalcontacts contact journey sessions o k response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get externalcontacts contact journey sessions o k response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions o k response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact journey sessions o k response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions o k response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetExternalcontactsContactJourneySessionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsOK) GetPayload() *models.SessionListing {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsBadRequest creates a GetExternalcontactsContactJourneySessionsBadRequest with default headers values
func NewGetExternalcontactsContactJourneySessionsBadRequest() *GetExternalcontactsContactJourneySessionsBadRequest {
	return &GetExternalcontactsContactJourneySessionsBadRequest{}
}

/*
GetExternalcontactsContactJourneySessionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsContactJourneySessionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions bad request response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions bad request response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions bad request response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions bad request response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions bad request response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetExternalcontactsContactJourneySessionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsUnauthorized creates a GetExternalcontactsContactJourneySessionsUnauthorized with default headers values
func NewGetExternalcontactsContactJourneySessionsUnauthorized() *GetExternalcontactsContactJourneySessionsUnauthorized {
	return &GetExternalcontactsContactJourneySessionsUnauthorized{}
}

/*
GetExternalcontactsContactJourneySessionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsContactJourneySessionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions unauthorized response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions unauthorized response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions unauthorized response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions unauthorized response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions unauthorized response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetExternalcontactsContactJourneySessionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsForbidden creates a GetExternalcontactsContactJourneySessionsForbidden with default headers values
func NewGetExternalcontactsContactJourneySessionsForbidden() *GetExternalcontactsContactJourneySessionsForbidden {
	return &GetExternalcontactsContactJourneySessionsForbidden{}
}

/*
GetExternalcontactsContactJourneySessionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsContactJourneySessionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions forbidden response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions forbidden response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions forbidden response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions forbidden response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions forbidden response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetExternalcontactsContactJourneySessionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsNotFound creates a GetExternalcontactsContactJourneySessionsNotFound with default headers values
func NewGetExternalcontactsContactJourneySessionsNotFound() *GetExternalcontactsContactJourneySessionsNotFound {
	return &GetExternalcontactsContactJourneySessionsNotFound{}
}

/*
GetExternalcontactsContactJourneySessionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetExternalcontactsContactJourneySessionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions not found response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions not found response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions not found response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions not found response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions not found response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetExternalcontactsContactJourneySessionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsRequestTimeout creates a GetExternalcontactsContactJourneySessionsRequestTimeout with default headers values
func NewGetExternalcontactsContactJourneySessionsRequestTimeout() *GetExternalcontactsContactJourneySessionsRequestTimeout {
	return &GetExternalcontactsContactJourneySessionsRequestTimeout{}
}

/*
GetExternalcontactsContactJourneySessionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsContactJourneySessionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions request timeout response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions request timeout response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions request timeout response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions request timeout response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions request timeout response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsRequestEntityTooLarge creates a GetExternalcontactsContactJourneySessionsRequestEntityTooLarge with default headers values
func NewGetExternalcontactsContactJourneySessionsRequestEntityTooLarge() *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge {
	return &GetExternalcontactsContactJourneySessionsRequestEntityTooLarge{}
}

/*
GetExternalcontactsContactJourneySessionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetExternalcontactsContactJourneySessionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions request entity too large response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions request entity too large response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions request entity too large response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions request entity too large response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions request entity too large response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsUnsupportedMediaType creates a GetExternalcontactsContactJourneySessionsUnsupportedMediaType with default headers values
func NewGetExternalcontactsContactJourneySessionsUnsupportedMediaType() *GetExternalcontactsContactJourneySessionsUnsupportedMediaType {
	return &GetExternalcontactsContactJourneySessionsUnsupportedMediaType{}
}

/*
GetExternalcontactsContactJourneySessionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsContactJourneySessionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions unsupported media type response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions unsupported media type response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions unsupported media type response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions unsupported media type response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions unsupported media type response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsTooManyRequests creates a GetExternalcontactsContactJourneySessionsTooManyRequests with default headers values
func NewGetExternalcontactsContactJourneySessionsTooManyRequests() *GetExternalcontactsContactJourneySessionsTooManyRequests {
	return &GetExternalcontactsContactJourneySessionsTooManyRequests{}
}

/*
GetExternalcontactsContactJourneySessionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsContactJourneySessionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions too many requests response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions too many requests response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions too many requests response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get externalcontacts contact journey sessions too many requests response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get externalcontacts contact journey sessions too many requests response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsInternalServerError creates a GetExternalcontactsContactJourneySessionsInternalServerError with default headers values
func NewGetExternalcontactsContactJourneySessionsInternalServerError() *GetExternalcontactsContactJourneySessionsInternalServerError {
	return &GetExternalcontactsContactJourneySessionsInternalServerError{}
}

/*
GetExternalcontactsContactJourneySessionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsContactJourneySessionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions internal server error response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions internal server error response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions internal server error response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact journey sessions internal server error response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts contact journey sessions internal server error response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetExternalcontactsContactJourneySessionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsServiceUnavailable creates a GetExternalcontactsContactJourneySessionsServiceUnavailable with default headers values
func NewGetExternalcontactsContactJourneySessionsServiceUnavailable() *GetExternalcontactsContactJourneySessionsServiceUnavailable {
	return &GetExternalcontactsContactJourneySessionsServiceUnavailable{}
}

/*
GetExternalcontactsContactJourneySessionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsContactJourneySessionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions service unavailable response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions service unavailable response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions service unavailable response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact journey sessions service unavailable response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts contact journey sessions service unavailable response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactJourneySessionsGatewayTimeout creates a GetExternalcontactsContactJourneySessionsGatewayTimeout with default headers values
func NewGetExternalcontactsContactJourneySessionsGatewayTimeout() *GetExternalcontactsContactJourneySessionsGatewayTimeout {
	return &GetExternalcontactsContactJourneySessionsGatewayTimeout{}
}

/*
GetExternalcontactsContactJourneySessionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetExternalcontactsContactJourneySessionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get externalcontacts contact journey sessions gateway timeout response has a 2xx status code
func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get externalcontacts contact journey sessions gateway timeout response has a 3xx status code
func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get externalcontacts contact journey sessions gateway timeout response has a 4xx status code
func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get externalcontacts contact journey sessions gateway timeout response has a 5xx status code
func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get externalcontacts contact journey sessions gateway timeout response a status code equal to that given
func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/{contactId}/journey/sessions][%d] getExternalcontactsContactJourneySessionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactJourneySessionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
