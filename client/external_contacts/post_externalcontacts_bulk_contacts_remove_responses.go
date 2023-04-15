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

// PostExternalcontactsBulkContactsRemoveReader is a Reader for the PostExternalcontactsBulkContactsRemove structure.
type PostExternalcontactsBulkContactsRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkContactsRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkContactsRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkContactsRemoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkContactsRemoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkContactsRemoveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkContactsRemoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkContactsRemoveRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkContactsRemoveRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkContactsRemoveUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkContactsRemoveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkContactsRemoveTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkContactsRemoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkContactsRemoveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkContactsRemoveGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkContactsRemoveOK creates a PostExternalcontactsBulkContactsRemoveOK with default headers values
func NewPostExternalcontactsBulkContactsRemoveOK() *PostExternalcontactsBulkContactsRemoveOK {
	return &PostExternalcontactsBulkContactsRemoveOK{}
}

/*
PostExternalcontactsBulkContactsRemoveOK describes a response with status code 200, with default header values.

successful operation
*/
type PostExternalcontactsBulkContactsRemoveOK struct {
	Payload *models.BulkDeleteResponse
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove o k response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove o k response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove o k response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts remove o k response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove o k response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostExternalcontactsBulkContactsRemoveOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveOK) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveOK) GetPayload() *models.BulkDeleteResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveBadRequest creates a PostExternalcontactsBulkContactsRemoveBadRequest with default headers values
func NewPostExternalcontactsBulkContactsRemoveBadRequest() *PostExternalcontactsBulkContactsRemoveBadRequest {
	return &PostExternalcontactsBulkContactsRemoveBadRequest{}
}

/*
PostExternalcontactsBulkContactsRemoveBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkContactsRemoveBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove bad request response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove bad request response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove bad request response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove bad request response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove bad request response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostExternalcontactsBulkContactsRemoveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveUnauthorized creates a PostExternalcontactsBulkContactsRemoveUnauthorized with default headers values
func NewPostExternalcontactsBulkContactsRemoveUnauthorized() *PostExternalcontactsBulkContactsRemoveUnauthorized {
	return &PostExternalcontactsBulkContactsRemoveUnauthorized{}
}

/*
PostExternalcontactsBulkContactsRemoveUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkContactsRemoveUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove unauthorized response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove unauthorized response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove unauthorized response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove unauthorized response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove unauthorized response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveForbidden creates a PostExternalcontactsBulkContactsRemoveForbidden with default headers values
func NewPostExternalcontactsBulkContactsRemoveForbidden() *PostExternalcontactsBulkContactsRemoveForbidden {
	return &PostExternalcontactsBulkContactsRemoveForbidden{}
}

/*
PostExternalcontactsBulkContactsRemoveForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkContactsRemoveForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove forbidden response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove forbidden response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove forbidden response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove forbidden response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove forbidden response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostExternalcontactsBulkContactsRemoveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveNotFound creates a PostExternalcontactsBulkContactsRemoveNotFound with default headers values
func NewPostExternalcontactsBulkContactsRemoveNotFound() *PostExternalcontactsBulkContactsRemoveNotFound {
	return &PostExternalcontactsBulkContactsRemoveNotFound{}
}

/*
PostExternalcontactsBulkContactsRemoveNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkContactsRemoveNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove not found response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove not found response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove not found response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove not found response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove not found response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostExternalcontactsBulkContactsRemoveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveRequestTimeout creates a PostExternalcontactsBulkContactsRemoveRequestTimeout with default headers values
func NewPostExternalcontactsBulkContactsRemoveRequestTimeout() *PostExternalcontactsBulkContactsRemoveRequestTimeout {
	return &PostExternalcontactsBulkContactsRemoveRequestTimeout{}
}

/*
PostExternalcontactsBulkContactsRemoveRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkContactsRemoveRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove request timeout response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove request timeout response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove request timeout response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove request timeout response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove request timeout response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveRequestEntityTooLarge creates a PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkContactsRemoveRequestEntityTooLarge() *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge {
	return &PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge{}
}

/*
PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove request entity too large response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove request entity too large response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove request entity too large response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove request entity too large response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove request entity too large response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveUnsupportedMediaType creates a PostExternalcontactsBulkContactsRemoveUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkContactsRemoveUnsupportedMediaType() *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType {
	return &PostExternalcontactsBulkContactsRemoveUnsupportedMediaType{}
}

/*
PostExternalcontactsBulkContactsRemoveUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkContactsRemoveUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove unsupported media type response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove unsupported media type response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove unsupported media type response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove unsupported media type response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove unsupported media type response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveUnprocessableEntity creates a PostExternalcontactsBulkContactsRemoveUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkContactsRemoveUnprocessableEntity() *PostExternalcontactsBulkContactsRemoveUnprocessableEntity {
	return &PostExternalcontactsBulkContactsRemoveUnprocessableEntity{}
}

/*
PostExternalcontactsBulkContactsRemoveUnprocessableEntity describes a response with status code 422, with default header values.

PostExternalcontactsBulkContactsRemoveUnprocessableEntity post externalcontacts bulk contacts remove unprocessable entity
*/
type PostExternalcontactsBulkContactsRemoveUnprocessableEntity struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove unprocessable entity response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove unprocessable entity response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove unprocessable entity response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove unprocessable entity response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove unprocessable entity response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveTooManyRequests creates a PostExternalcontactsBulkContactsRemoveTooManyRequests with default headers values
func NewPostExternalcontactsBulkContactsRemoveTooManyRequests() *PostExternalcontactsBulkContactsRemoveTooManyRequests {
	return &PostExternalcontactsBulkContactsRemoveTooManyRequests{}
}

/*
PostExternalcontactsBulkContactsRemoveTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkContactsRemoveTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove too many requests response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove too many requests response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove too many requests response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post externalcontacts bulk contacts remove too many requests response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post externalcontacts bulk contacts remove too many requests response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveInternalServerError creates a PostExternalcontactsBulkContactsRemoveInternalServerError with default headers values
func NewPostExternalcontactsBulkContactsRemoveInternalServerError() *PostExternalcontactsBulkContactsRemoveInternalServerError {
	return &PostExternalcontactsBulkContactsRemoveInternalServerError{}
}

/*
PostExternalcontactsBulkContactsRemoveInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkContactsRemoveInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove internal server error response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove internal server error response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove internal server error response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts remove internal server error response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts remove internal server error response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveServiceUnavailable creates a PostExternalcontactsBulkContactsRemoveServiceUnavailable with default headers values
func NewPostExternalcontactsBulkContactsRemoveServiceUnavailable() *PostExternalcontactsBulkContactsRemoveServiceUnavailable {
	return &PostExternalcontactsBulkContactsRemoveServiceUnavailable{}
}

/*
PostExternalcontactsBulkContactsRemoveServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkContactsRemoveServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove service unavailable response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove service unavailable response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove service unavailable response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts remove service unavailable response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts remove service unavailable response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRemoveGatewayTimeout creates a PostExternalcontactsBulkContactsRemoveGatewayTimeout with default headers values
func NewPostExternalcontactsBulkContactsRemoveGatewayTimeout() *PostExternalcontactsBulkContactsRemoveGatewayTimeout {
	return &PostExternalcontactsBulkContactsRemoveGatewayTimeout{}
}

/*
PostExternalcontactsBulkContactsRemoveGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostExternalcontactsBulkContactsRemoveGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post externalcontacts bulk contacts remove gateway timeout response has a 2xx status code
func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post externalcontacts bulk contacts remove gateway timeout response has a 3xx status code
func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post externalcontacts bulk contacts remove gateway timeout response has a 4xx status code
func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post externalcontacts bulk contacts remove gateway timeout response has a 5xx status code
func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post externalcontacts bulk contacts remove gateway timeout response a status code equal to that given
func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/remove][%d] postExternalcontactsBulkContactsRemoveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRemoveGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
