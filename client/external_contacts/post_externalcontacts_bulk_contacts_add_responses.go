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

// PostExternalcontactsBulkContactsAddReader is a Reader for the PostExternalcontactsBulkContactsAdd structure.
type PostExternalcontactsBulkContactsAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkContactsAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkContactsAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkContactsAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkContactsAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkContactsAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkContactsAddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkContactsAddRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkContactsAddRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkContactsAddUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkContactsAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkContactsAddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkContactsAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkContactsAddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkContactsAddGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkContactsAddOK creates a PostExternalcontactsBulkContactsAddOK with default headers values
func NewPostExternalcontactsBulkContactsAddOK() *PostExternalcontactsBulkContactsAddOK {
	return &PostExternalcontactsBulkContactsAddOK{}
}

/*
PostExternalcontactsBulkContactsAddOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkContactsAddOK struct {
	Payload *models.BulkContactsResponse
}

// IsSuccess returns true when this post externalcontacts bulk contacts add o k response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk contacts add o k response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add o k response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts add o k response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add o k response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkContactsAddOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddOK) GetPayload() *models.BulkContactsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkContactsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddBadRequest creates a PostExternalcontactsBulkContactsAddBadRequest with default headers values
func NewPostExternalcontactsBulkContactsAddBadRequest() *PostExternalcontactsBulkContactsAddBadRequest {
	return &PostExternalcontactsBulkContactsAddBadRequest{}
}

/*
PostExternalcontactsBulkContactsAddBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkContactsAddBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add bad request response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add bad request response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add bad request response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add bad request response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add bad request response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkContactsAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddUnauthorized creates a PostExternalcontactsBulkContactsAddUnauthorized with default headers values
func NewPostExternalcontactsBulkContactsAddUnauthorized() *PostExternalcontactsBulkContactsAddUnauthorized {
	return &PostExternalcontactsBulkContactsAddUnauthorized{}
}

/*
PostExternalcontactsBulkContactsAddUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkContactsAddUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkContactsAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddForbidden creates a PostExternalcontactsBulkContactsAddForbidden with default headers values
func NewPostExternalcontactsBulkContactsAddForbidden() *PostExternalcontactsBulkContactsAddForbidden {
	return &PostExternalcontactsBulkContactsAddForbidden{}
}

/*
PostExternalcontactsBulkContactsAddForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkContactsAddForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkContactsAddForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddNotFound creates a PostExternalcontactsBulkContactsAddNotFound with default headers values
func NewPostExternalcontactsBulkContactsAddNotFound() *PostExternalcontactsBulkContactsAddNotFound {
	return &PostExternalcontactsBulkContactsAddNotFound{}
}

/*
PostExternalcontactsBulkContactsAddNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkContactsAddNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add not found response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add not found response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add not found response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add not found response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add not found response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkContactsAddNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddRequestTimeout creates a PostExternalcontactsBulkContactsAddRequestTimeout with default headers values
func NewPostExternalcontactsBulkContactsAddRequestTimeout() *PostExternalcontactsBulkContactsAddRequestTimeout {
	return &PostExternalcontactsBulkContactsAddRequestTimeout{}
}

/*
PostExternalcontactsBulkContactsAddRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkContactsAddRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkContactsAddRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddRequestEntityTooLarge creates a PostExternalcontactsBulkContactsAddRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkContactsAddRequestEntityTooLarge() *PostExternalcontactsBulkContactsAddRequestEntityTooLarge {
	return &PostExternalcontactsBulkContactsAddRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkContactsAddRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostExternalcontactsBulkContactsAddRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddUnsupportedMediaType creates a PostExternalcontactsBulkContactsAddUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkContactsAddUnsupportedMediaType() *PostExternalcontactsBulkContactsAddUnsupportedMediaType {
	return &PostExternalcontactsBulkContactsAddUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkContactsAddUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkContactsAddUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddUnprocessableEntity creates a PostExternalcontactsBulkContactsAddUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkContactsAddUnprocessableEntity() *PostExternalcontactsBulkContactsAddUnprocessableEntity {
	return &PostExternalcontactsBulkContactsAddUnprocessableEntity{}
}

/*
PostExternalcontactsBulkContactsAddUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsBulkContactsAddUnprocessableEntity post externalcontacts bulk contacts add unprocessable entity
*/
type PostExternalcontactsBulkContactsAddUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddTooManyRequests creates a PostExternalcontactsBulkContactsAddTooManyRequests with default headers values
func NewPostExternalcontactsBulkContactsAddTooManyRequests() *PostExternalcontactsBulkContactsAddTooManyRequests {
	return &PostExternalcontactsBulkContactsAddTooManyRequests{}
}

/*
PostExternalcontactsBulkContactsAddTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkContactsAddTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts add too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts add too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkContactsAddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddInternalServerError creates a PostExternalcontactsBulkContactsAddInternalServerError with default headers values
func NewPostExternalcontactsBulkContactsAddInternalServerError() *PostExternalcontactsBulkContactsAddInternalServerError {
	return &PostExternalcontactsBulkContactsAddInternalServerError{}
}

/*
PostExternalcontactsBulkContactsAddInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkContactsAddInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts add internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts add internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkContactsAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddServiceUnavailable creates a PostExternalcontactsBulkContactsAddServiceUnavailable with default headers values
func NewPostExternalcontactsBulkContactsAddServiceUnavailable() *PostExternalcontactsBulkContactsAddServiceUnavailable {
	return &PostExternalcontactsBulkContactsAddServiceUnavailable{}
}

/*
PostExternalcontactsBulkContactsAddServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkContactsAddServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts add service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts add service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsAddGatewayTimeout creates a PostExternalcontactsBulkContactsAddGatewayTimeout with default headers values
func NewPostExternalcontactsBulkContactsAddGatewayTimeout() *PostExternalcontactsBulkContactsAddGatewayTimeout {
	return &PostExternalcontactsBulkContactsAddGatewayTimeout{}
}

/*
PostExternalcontactsBulkContactsAddGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkContactsAddGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts add gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts add gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts add gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts add gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts add gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/add][%d] postExternalcontactsBulkContactsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsAddGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
