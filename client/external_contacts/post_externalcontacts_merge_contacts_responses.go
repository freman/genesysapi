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

// PostExternalcontactsMergeContactsReader is a Reader for the PostExternalcontactsMergeContacts structure.
type PostExternalcontactsMergeContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsMergeContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsMergeContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsMergeContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsMergeContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsMergeContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsMergeContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsMergeContactsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostExternalcontactsMergeContactsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsMergeContactsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsMergeContactsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsMergeContactsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsMergeContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsMergeContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsMergeContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsMergeContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsMergeContactsOK creates a PostExternalcontactsMergeContactsOK with default headers values
func NewPostExternalcontactsMergeContactsOK() *PostExternalcontactsMergeContactsOK {
	return &PostExternalcontactsMergeContactsOK{}
}

/*
PostExternalcontactsMergeContactsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsMergeContactsOK struct {
	Payload *models.ExternalContact
}

// IsSuccess returns true when this post externalcontacts merge contacts o k response has a 2xx status code
func (o *PostExternalcontactsMergeContactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts merge contacts o k response has a 3xx status code
func (o *PostExternalcontactsMergeContactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts o k response has a 4xx status code
func (o *PostExternalcontactsMergeContactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts merge contacts o k response has a 5xx status code
func (o *PostExternalcontactsMergeContactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts o k response a status code equal to that given
func (o *PostExternalcontactsMergeContactsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsMergeContactsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsMergeContactsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsMergeContactsOK) GetPayload() *models.ExternalContact {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsBadRequest creates a PostExternalcontactsMergeContactsBadRequest with default headers values
func NewPostExternalcontactsMergeContactsBadRequest() *PostExternalcontactsMergeContactsBadRequest {
	return &PostExternalcontactsMergeContactsBadRequest{}
}

/*
PostExternalcontactsMergeContactsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsMergeContactsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts bad request response has a 2xx status code
func (o *PostExternalcontactsMergeContactsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts bad request response has a 3xx status code
func (o *PostExternalcontactsMergeContactsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts bad request response has a 4xx status code
func (o *PostExternalcontactsMergeContactsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts bad request response has a 5xx status code
func (o *PostExternalcontactsMergeContactsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts bad request response a status code equal to that given
func (o *PostExternalcontactsMergeContactsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsMergeContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsMergeContactsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsMergeContactsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsUnauthorized creates a PostExternalcontactsMergeContactsUnauthorized with default headers values
func NewPostExternalcontactsMergeContactsUnauthorized() *PostExternalcontactsMergeContactsUnauthorized {
	return &PostExternalcontactsMergeContactsUnauthorized{}
}

/*
PostExternalcontactsMergeContactsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsMergeContactsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts unauthorized response has a 2xx status code
func (o *PostExternalcontactsMergeContactsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts unauthorized response has a 3xx status code
func (o *PostExternalcontactsMergeContactsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts unauthorized response has a 4xx status code
func (o *PostExternalcontactsMergeContactsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts unauthorized response has a 5xx status code
func (o *PostExternalcontactsMergeContactsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts unauthorized response a status code equal to that given
func (o *PostExternalcontactsMergeContactsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsMergeContactsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsMergeContactsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsMergeContactsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsForbidden creates a PostExternalcontactsMergeContactsForbidden with default headers values
func NewPostExternalcontactsMergeContactsForbidden() *PostExternalcontactsMergeContactsForbidden {
	return &PostExternalcontactsMergeContactsForbidden{}
}

/*
PostExternalcontactsMergeContactsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsMergeContactsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts forbidden response has a 2xx status code
func (o *PostExternalcontactsMergeContactsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts forbidden response has a 3xx status code
func (o *PostExternalcontactsMergeContactsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts forbidden response has a 4xx status code
func (o *PostExternalcontactsMergeContactsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts forbidden response has a 5xx status code
func (o *PostExternalcontactsMergeContactsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts forbidden response a status code equal to that given
func (o *PostExternalcontactsMergeContactsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsMergeContactsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsMergeContactsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsMergeContactsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsNotFound creates a PostExternalcontactsMergeContactsNotFound with default headers values
func NewPostExternalcontactsMergeContactsNotFound() *PostExternalcontactsMergeContactsNotFound {
	return &PostExternalcontactsMergeContactsNotFound{}
}

/*
PostExternalcontactsMergeContactsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsMergeContactsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts not found response has a 2xx status code
func (o *PostExternalcontactsMergeContactsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts not found response has a 3xx status code
func (o *PostExternalcontactsMergeContactsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts not found response has a 4xx status code
func (o *PostExternalcontactsMergeContactsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts not found response has a 5xx status code
func (o *PostExternalcontactsMergeContactsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts not found response a status code equal to that given
func (o *PostExternalcontactsMergeContactsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsMergeContactsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsMergeContactsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsMergeContactsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsRequestTimeout creates a PostExternalcontactsMergeContactsRequestTimeout with default headers values
func NewPostExternalcontactsMergeContactsRequestTimeout() *PostExternalcontactsMergeContactsRequestTimeout {
	return &PostExternalcontactsMergeContactsRequestTimeout{}
}

/*
PostExternalcontactsMergeContactsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsMergeContactsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts request timeout response has a 2xx status code
func (o *PostExternalcontactsMergeContactsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts request timeout response has a 3xx status code
func (o *PostExternalcontactsMergeContactsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts request timeout response has a 4xx status code
func (o *PostExternalcontactsMergeContactsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts request timeout response has a 5xx status code
func (o *PostExternalcontactsMergeContactsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts request timeout response a status code equal to that given
func (o *PostExternalcontactsMergeContactsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsMergeContactsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsMergeContactsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsMergeContactsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsConflict creates a PostExternalcontactsMergeContactsConflict with default headers values
func NewPostExternalcontactsMergeContactsConflict() *PostExternalcontactsMergeContactsConflict {
	return &PostExternalcontactsMergeContactsConflict{}
}

/*
PostExternalcontactsMergeContactsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostExternalcontactsMergeContactsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts conflict response has a 2xx status code
func (o *PostExternalcontactsMergeContactsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts conflict response has a 3xx status code
func (o *PostExternalcontactsMergeContactsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts conflict response has a 4xx status code
func (o *PostExternalcontactsMergeContactsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts conflict response has a 5xx status code
func (o *PostExternalcontactsMergeContactsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts conflict response a status code equal to that given
func (o *PostExternalcontactsMergeContactsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostExternalcontactsMergeContactsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsConflict  %+v", 409, o.Payload)
}

func (o *PostExternalcontactsMergeContactsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsConflict  %+v", 409, o.Payload)
}

func (o *PostExternalcontactsMergeContactsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsRequestEntityTooLarge creates a PostExternalcontactsMergeContactsRequestEntityTooLarge with default headers values
func NewPostExternalcontactsMergeContactsRequestEntityTooLarge() *PostExternalcontactsMergeContactsRequestEntityTooLarge {
	return &PostExternalcontactsMergeContactsRequestEntityTooLarge{}
}

/*
PostExternalcontactsMergeContactsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostExternalcontactsMergeContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts request entity too large response has a 2xx status code
func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts request entity too large response has a 3xx status code
func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts request entity too large response has a 4xx status code
func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts request entity too large response has a 5xx status code
func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts request entity too large response a status code equal to that given
func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsUnsupportedMediaType creates a PostExternalcontactsMergeContactsUnsupportedMediaType with default headers values
func NewPostExternalcontactsMergeContactsUnsupportedMediaType() *PostExternalcontactsMergeContactsUnsupportedMediaType {
	return &PostExternalcontactsMergeContactsUnsupportedMediaType{}
}

/*
PostExternalcontactsMergeContactsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsMergeContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts unsupported media type response has a 2xx status code
func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts unsupported media type response has a 3xx status code
func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts unsupported media type response has a 4xx status code
func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts unsupported media type response has a 5xx status code
func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts unsupported media type response a status code equal to that given
func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsUnprocessableEntity creates a PostExternalcontactsMergeContactsUnprocessableEntity with default headers values
func NewPostExternalcontactsMergeContactsUnprocessableEntity() *PostExternalcontactsMergeContactsUnprocessableEntity {
	return &PostExternalcontactsMergeContactsUnprocessableEntity{}
}

/*
PostExternalcontactsMergeContactsUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsMergeContactsUnprocessableEntity post externalcontacts merge contacts unprocessable entity
*/
type PostExternalcontactsMergeContactsUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsMergeContactsUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsMergeContactsUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsMergeContactsUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsMergeContactsUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsMergeContactsUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsMergeContactsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsMergeContactsUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsMergeContactsUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsTooManyRequests creates a PostExternalcontactsMergeContactsTooManyRequests with default headers values
func NewPostExternalcontactsMergeContactsTooManyRequests() *PostExternalcontactsMergeContactsTooManyRequests {
	return &PostExternalcontactsMergeContactsTooManyRequests{}
}

/*
PostExternalcontactsMergeContactsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsMergeContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts too many requests response has a 2xx status code
func (o *PostExternalcontactsMergeContactsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts too many requests response has a 3xx status code
func (o *PostExternalcontactsMergeContactsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts too many requests response has a 4xx status code
func (o *PostExternalcontactsMergeContactsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts merge contacts too many requests response has a 5xx status code
func (o *PostExternalcontactsMergeContactsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts merge contacts too many requests response a status code equal to that given
func (o *PostExternalcontactsMergeContactsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsMergeContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsMergeContactsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsMergeContactsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsInternalServerError creates a PostExternalcontactsMergeContactsInternalServerError with default headers values
func NewPostExternalcontactsMergeContactsInternalServerError() *PostExternalcontactsMergeContactsInternalServerError {
	return &PostExternalcontactsMergeContactsInternalServerError{}
}

/*
PostExternalcontactsMergeContactsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsMergeContactsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts internal server error response has a 2xx status code
func (o *PostExternalcontactsMergeContactsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts internal server error response has a 3xx status code
func (o *PostExternalcontactsMergeContactsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts internal server error response has a 4xx status code
func (o *PostExternalcontactsMergeContactsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts merge contacts internal server error response has a 5xx status code
func (o *PostExternalcontactsMergeContactsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts merge contacts internal server error response a status code equal to that given
func (o *PostExternalcontactsMergeContactsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsMergeContactsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsMergeContactsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsMergeContactsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsServiceUnavailable creates a PostExternalcontactsMergeContactsServiceUnavailable with default headers values
func NewPostExternalcontactsMergeContactsServiceUnavailable() *PostExternalcontactsMergeContactsServiceUnavailable {
	return &PostExternalcontactsMergeContactsServiceUnavailable{}
}

/*
PostExternalcontactsMergeContactsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsMergeContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts service unavailable response has a 2xx status code
func (o *PostExternalcontactsMergeContactsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts service unavailable response has a 3xx status code
func (o *PostExternalcontactsMergeContactsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts service unavailable response has a 4xx status code
func (o *PostExternalcontactsMergeContactsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts merge contacts service unavailable response has a 5xx status code
func (o *PostExternalcontactsMergeContactsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts merge contacts service unavailable response a status code equal to that given
func (o *PostExternalcontactsMergeContactsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsMergeContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsMergeContactsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsMergeContactsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsMergeContactsGatewayTimeout creates a PostExternalcontactsMergeContactsGatewayTimeout with default headers values
func NewPostExternalcontactsMergeContactsGatewayTimeout() *PostExternalcontactsMergeContactsGatewayTimeout {
	return &PostExternalcontactsMergeContactsGatewayTimeout{}
}

/*
PostExternalcontactsMergeContactsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsMergeContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts merge contacts gateway timeout response has a 2xx status code
func (o *PostExternalcontactsMergeContactsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts merge contacts gateway timeout response has a 3xx status code
func (o *PostExternalcontactsMergeContactsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts merge contacts gateway timeout response has a 4xx status code
func (o *PostExternalcontactsMergeContactsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts merge contacts gateway timeout response has a 5xx status code
func (o *PostExternalcontactsMergeContactsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts merge contacts gateway timeout response a status code equal to that given
func (o *PostExternalcontactsMergeContactsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsMergeContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsMergeContactsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/merge/contacts][%d] postExternalcontactsMergeContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsMergeContactsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsMergeContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
