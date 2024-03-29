// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguagesTranslationsOrganizationReader is a Reader for the GetLanguagesTranslationsOrganization structure.
type GetLanguagesTranslationsOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguagesTranslationsOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguagesTranslationsOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguagesTranslationsOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguagesTranslationsOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguagesTranslationsOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguagesTranslationsOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguagesTranslationsOrganizationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguagesTranslationsOrganizationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguagesTranslationsOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguagesTranslationsOrganizationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguagesTranslationsOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguagesTranslationsOrganizationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguagesTranslationsOrganizationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguagesTranslationsOrganizationOK creates a GetLanguagesTranslationsOrganizationOK with default headers values
func NewGetLanguagesTranslationsOrganizationOK() *GetLanguagesTranslationsOrganizationOK {
	return &GetLanguagesTranslationsOrganizationOK{}
}

/*
GetLanguagesTranslationsOrganizationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLanguagesTranslationsOrganizationOK struct {
	Payload map[string]interface{}
}

// IsSuccess returns true when this get languages translations organization o k response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get languages translations organization o k response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization o k response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages translations organization o k response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization o k response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLanguagesTranslationsOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationBadRequest creates a GetLanguagesTranslationsOrganizationBadRequest with default headers values
func NewGetLanguagesTranslationsOrganizationBadRequest() *GetLanguagesTranslationsOrganizationBadRequest {
	return &GetLanguagesTranslationsOrganizationBadRequest{}
}

/*
GetLanguagesTranslationsOrganizationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguagesTranslationsOrganizationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization bad request response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization bad request response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization bad request response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization bad request response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization bad request response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLanguagesTranslationsOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationUnauthorized creates a GetLanguagesTranslationsOrganizationUnauthorized with default headers values
func NewGetLanguagesTranslationsOrganizationUnauthorized() *GetLanguagesTranslationsOrganizationUnauthorized {
	return &GetLanguagesTranslationsOrganizationUnauthorized{}
}

/*
GetLanguagesTranslationsOrganizationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguagesTranslationsOrganizationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization unauthorized response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization unauthorized response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization unauthorized response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization unauthorized response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization unauthorized response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLanguagesTranslationsOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationForbidden creates a GetLanguagesTranslationsOrganizationForbidden with default headers values
func NewGetLanguagesTranslationsOrganizationForbidden() *GetLanguagesTranslationsOrganizationForbidden {
	return &GetLanguagesTranslationsOrganizationForbidden{}
}

/*
GetLanguagesTranslationsOrganizationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguagesTranslationsOrganizationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization forbidden response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization forbidden response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization forbidden response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization forbidden response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization forbidden response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLanguagesTranslationsOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationNotFound creates a GetLanguagesTranslationsOrganizationNotFound with default headers values
func NewGetLanguagesTranslationsOrganizationNotFound() *GetLanguagesTranslationsOrganizationNotFound {
	return &GetLanguagesTranslationsOrganizationNotFound{}
}

/*
GetLanguagesTranslationsOrganizationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLanguagesTranslationsOrganizationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization not found response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization not found response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization not found response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization not found response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization not found response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLanguagesTranslationsOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationRequestTimeout creates a GetLanguagesTranslationsOrganizationRequestTimeout with default headers values
func NewGetLanguagesTranslationsOrganizationRequestTimeout() *GetLanguagesTranslationsOrganizationRequestTimeout {
	return &GetLanguagesTranslationsOrganizationRequestTimeout{}
}

/*
GetLanguagesTranslationsOrganizationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguagesTranslationsOrganizationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization request timeout response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization request timeout response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization request timeout response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization request timeout response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization request timeout response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLanguagesTranslationsOrganizationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationRequestEntityTooLarge creates a GetLanguagesTranslationsOrganizationRequestEntityTooLarge with default headers values
func NewGetLanguagesTranslationsOrganizationRequestEntityTooLarge() *GetLanguagesTranslationsOrganizationRequestEntityTooLarge {
	return &GetLanguagesTranslationsOrganizationRequestEntityTooLarge{}
}

/*
GetLanguagesTranslationsOrganizationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLanguagesTranslationsOrganizationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization request entity too large response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization request entity too large response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization request entity too large response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization request entity too large response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization request entity too large response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationUnsupportedMediaType creates a GetLanguagesTranslationsOrganizationUnsupportedMediaType with default headers values
func NewGetLanguagesTranslationsOrganizationUnsupportedMediaType() *GetLanguagesTranslationsOrganizationUnsupportedMediaType {
	return &GetLanguagesTranslationsOrganizationUnsupportedMediaType{}
}

/*
GetLanguagesTranslationsOrganizationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguagesTranslationsOrganizationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization unsupported media type response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization unsupported media type response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization unsupported media type response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization unsupported media type response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization unsupported media type response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationTooManyRequests creates a GetLanguagesTranslationsOrganizationTooManyRequests with default headers values
func NewGetLanguagesTranslationsOrganizationTooManyRequests() *GetLanguagesTranslationsOrganizationTooManyRequests {
	return &GetLanguagesTranslationsOrganizationTooManyRequests{}
}

/*
GetLanguagesTranslationsOrganizationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguagesTranslationsOrganizationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization too many requests response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization too many requests response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization too many requests response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages translations organization too many requests response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages translations organization too many requests response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLanguagesTranslationsOrganizationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationInternalServerError creates a GetLanguagesTranslationsOrganizationInternalServerError with default headers values
func NewGetLanguagesTranslationsOrganizationInternalServerError() *GetLanguagesTranslationsOrganizationInternalServerError {
	return &GetLanguagesTranslationsOrganizationInternalServerError{}
}

/*
GetLanguagesTranslationsOrganizationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguagesTranslationsOrganizationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization internal server error response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization internal server error response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization internal server error response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages translations organization internal server error response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get languages translations organization internal server error response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLanguagesTranslationsOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationServiceUnavailable creates a GetLanguagesTranslationsOrganizationServiceUnavailable with default headers values
func NewGetLanguagesTranslationsOrganizationServiceUnavailable() *GetLanguagesTranslationsOrganizationServiceUnavailable {
	return &GetLanguagesTranslationsOrganizationServiceUnavailable{}
}

/*
GetLanguagesTranslationsOrganizationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguagesTranslationsOrganizationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization service unavailable response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization service unavailable response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization service unavailable response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages translations organization service unavailable response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get languages translations organization service unavailable response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTranslationsOrganizationGatewayTimeout creates a GetLanguagesTranslationsOrganizationGatewayTimeout with default headers values
func NewGetLanguagesTranslationsOrganizationGatewayTimeout() *GetLanguagesTranslationsOrganizationGatewayTimeout {
	return &GetLanguagesTranslationsOrganizationGatewayTimeout{}
}

/*
GetLanguagesTranslationsOrganizationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLanguagesTranslationsOrganizationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages translations organization gateway timeout response has a 2xx status code
func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages translations organization gateway timeout response has a 3xx status code
func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages translations organization gateway timeout response has a 4xx status code
func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages translations organization gateway timeout response has a 5xx status code
func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get languages translations organization gateway timeout response a status code equal to that given
func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languages/translations/organization][%d] getLanguagesTranslationsOrganizationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTranslationsOrganizationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
