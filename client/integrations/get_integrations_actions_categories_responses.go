// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIntegrationsActionsCategoriesReader is a Reader for the GetIntegrationsActionsCategories structure.
type GetIntegrationsActionsCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionsCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionsCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionsCategoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionsCategoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionsCategoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionsCategoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsActionsCategoriesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionsCategoriesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionsCategoriesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionsCategoriesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionsCategoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionsCategoriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionsCategoriesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionsCategoriesOK creates a GetIntegrationsActionsCategoriesOK with default headers values
func NewGetIntegrationsActionsCategoriesOK() *GetIntegrationsActionsCategoriesOK {
	return &GetIntegrationsActionsCategoriesOK{}
}

/*
GetIntegrationsActionsCategoriesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsActionsCategoriesOK struct {
	Payload *models.CategoryEntityListing
}

// IsSuccess returns true when this get integrations actions categories o k response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations actions categories o k response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories o k response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions categories o k response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories o k response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsActionsCategoriesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesOK) GetPayload() *models.CategoryEntityListing {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesBadRequest creates a GetIntegrationsActionsCategoriesBadRequest with default headers values
func NewGetIntegrationsActionsCategoriesBadRequest() *GetIntegrationsActionsCategoriesBadRequest {
	return &GetIntegrationsActionsCategoriesBadRequest{}
}

/*
GetIntegrationsActionsCategoriesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionsCategoriesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories bad request response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories bad request response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories bad request response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories bad request response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories bad request response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsActionsCategoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesUnauthorized creates a GetIntegrationsActionsCategoriesUnauthorized with default headers values
func NewGetIntegrationsActionsCategoriesUnauthorized() *GetIntegrationsActionsCategoriesUnauthorized {
	return &GetIntegrationsActionsCategoriesUnauthorized{}
}

/*
GetIntegrationsActionsCategoriesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionsCategoriesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories unauthorized response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories unauthorized response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories unauthorized response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories unauthorized response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories unauthorized response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsActionsCategoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesForbidden creates a GetIntegrationsActionsCategoriesForbidden with default headers values
func NewGetIntegrationsActionsCategoriesForbidden() *GetIntegrationsActionsCategoriesForbidden {
	return &GetIntegrationsActionsCategoriesForbidden{}
}

/*
GetIntegrationsActionsCategoriesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionsCategoriesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories forbidden response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories forbidden response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories forbidden response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories forbidden response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories forbidden response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsActionsCategoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesNotFound creates a GetIntegrationsActionsCategoriesNotFound with default headers values
func NewGetIntegrationsActionsCategoriesNotFound() *GetIntegrationsActionsCategoriesNotFound {
	return &GetIntegrationsActionsCategoriesNotFound{}
}

/*
GetIntegrationsActionsCategoriesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionsCategoriesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories not found response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories not found response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories not found response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories not found response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories not found response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsActionsCategoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesRequestTimeout creates a GetIntegrationsActionsCategoriesRequestTimeout with default headers values
func NewGetIntegrationsActionsCategoriesRequestTimeout() *GetIntegrationsActionsCategoriesRequestTimeout {
	return &GetIntegrationsActionsCategoriesRequestTimeout{}
}

/*
GetIntegrationsActionsCategoriesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsActionsCategoriesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories request timeout response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories request timeout response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories request timeout response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories request timeout response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories request timeout response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsActionsCategoriesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesRequestEntityTooLarge creates a GetIntegrationsActionsCategoriesRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionsCategoriesRequestEntityTooLarge() *GetIntegrationsActionsCategoriesRequestEntityTooLarge {
	return &GetIntegrationsActionsCategoriesRequestEntityTooLarge{}
}

/*
GetIntegrationsActionsCategoriesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsActionsCategoriesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories request entity too large response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories request entity too large response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories request entity too large response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories request entity too large response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories request entity too large response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesUnsupportedMediaType creates a GetIntegrationsActionsCategoriesUnsupportedMediaType with default headers values
func NewGetIntegrationsActionsCategoriesUnsupportedMediaType() *GetIntegrationsActionsCategoriesUnsupportedMediaType {
	return &GetIntegrationsActionsCategoriesUnsupportedMediaType{}
}

/*
GetIntegrationsActionsCategoriesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionsCategoriesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories unsupported media type response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories unsupported media type response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories unsupported media type response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories unsupported media type response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories unsupported media type response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesTooManyRequests creates a GetIntegrationsActionsCategoriesTooManyRequests with default headers values
func NewGetIntegrationsActionsCategoriesTooManyRequests() *GetIntegrationsActionsCategoriesTooManyRequests {
	return &GetIntegrationsActionsCategoriesTooManyRequests{}
}

/*
GetIntegrationsActionsCategoriesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsActionsCategoriesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories too many requests response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories too many requests response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories too many requests response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations actions categories too many requests response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations actions categories too many requests response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsActionsCategoriesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesInternalServerError creates a GetIntegrationsActionsCategoriesInternalServerError with default headers values
func NewGetIntegrationsActionsCategoriesInternalServerError() *GetIntegrationsActionsCategoriesInternalServerError {
	return &GetIntegrationsActionsCategoriesInternalServerError{}
}

/*
GetIntegrationsActionsCategoriesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionsCategoriesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories internal server error response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories internal server error response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories internal server error response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions categories internal server error response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations actions categories internal server error response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsActionsCategoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesServiceUnavailable creates a GetIntegrationsActionsCategoriesServiceUnavailable with default headers values
func NewGetIntegrationsActionsCategoriesServiceUnavailable() *GetIntegrationsActionsCategoriesServiceUnavailable {
	return &GetIntegrationsActionsCategoriesServiceUnavailable{}
}

/*
GetIntegrationsActionsCategoriesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionsCategoriesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories service unavailable response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories service unavailable response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories service unavailable response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions categories service unavailable response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations actions categories service unavailable response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsActionsCategoriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionsCategoriesGatewayTimeout creates a GetIntegrationsActionsCategoriesGatewayTimeout with default headers values
func NewGetIntegrationsActionsCategoriesGatewayTimeout() *GetIntegrationsActionsCategoriesGatewayTimeout {
	return &GetIntegrationsActionsCategoriesGatewayTimeout{}
}

/*
GetIntegrationsActionsCategoriesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsActionsCategoriesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations actions categories gateway timeout response has a 2xx status code
func (o *GetIntegrationsActionsCategoriesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations actions categories gateway timeout response has a 3xx status code
func (o *GetIntegrationsActionsCategoriesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations actions categories gateway timeout response has a 4xx status code
func (o *GetIntegrationsActionsCategoriesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations actions categories gateway timeout response has a 5xx status code
func (o *GetIntegrationsActionsCategoriesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations actions categories gateway timeout response a status code equal to that given
func (o *GetIntegrationsActionsCategoriesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsActionsCategoriesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/categories][%d] getIntegrationsActionsCategoriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionsCategoriesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionsCategoriesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
