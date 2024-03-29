// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityFormsSurveyVersionsReader is a Reader for the GetQualityFormsSurveyVersions structure.
type GetQualityFormsSurveyVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityFormsSurveyVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityFormsSurveyVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityFormsSurveyVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityFormsSurveyVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityFormsSurveyVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityFormsSurveyVersionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetQualityFormsSurveyVersionsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityFormsSurveyVersionsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityFormsSurveyVersionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityFormsSurveyVersionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityFormsSurveyVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityFormsSurveyVersionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityFormsSurveyVersionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityFormsSurveyVersionsOK creates a GetQualityFormsSurveyVersionsOK with default headers values
func NewGetQualityFormsSurveyVersionsOK() *GetQualityFormsSurveyVersionsOK {
	return &GetQualityFormsSurveyVersionsOK{}
}

/*
GetQualityFormsSurveyVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetQualityFormsSurveyVersionsOK struct {
	Payload *models.SurveyFormEntityListing
}

// IsSuccess returns true when this get quality forms survey versions o k response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get quality forms survey versions o k response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions o k response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms survey versions o k response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions o k response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetQualityFormsSurveyVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsOK) GetPayload() *models.SurveyFormEntityListing {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SurveyFormEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsBadRequest creates a GetQualityFormsSurveyVersionsBadRequest with default headers values
func NewGetQualityFormsSurveyVersionsBadRequest() *GetQualityFormsSurveyVersionsBadRequest {
	return &GetQualityFormsSurveyVersionsBadRequest{}
}

/*
GetQualityFormsSurveyVersionsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityFormsSurveyVersionsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions bad request response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions bad request response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions bad request response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions bad request response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions bad request response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetQualityFormsSurveyVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsUnauthorized creates a GetQualityFormsSurveyVersionsUnauthorized with default headers values
func NewGetQualityFormsSurveyVersionsUnauthorized() *GetQualityFormsSurveyVersionsUnauthorized {
	return &GetQualityFormsSurveyVersionsUnauthorized{}
}

/*
GetQualityFormsSurveyVersionsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityFormsSurveyVersionsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions unauthorized response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions unauthorized response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions unauthorized response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions unauthorized response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions unauthorized response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetQualityFormsSurveyVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsForbidden creates a GetQualityFormsSurveyVersionsForbidden with default headers values
func NewGetQualityFormsSurveyVersionsForbidden() *GetQualityFormsSurveyVersionsForbidden {
	return &GetQualityFormsSurveyVersionsForbidden{}
}

/*
GetQualityFormsSurveyVersionsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityFormsSurveyVersionsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions forbidden response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions forbidden response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions forbidden response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions forbidden response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions forbidden response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetQualityFormsSurveyVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsNotFound creates a GetQualityFormsSurveyVersionsNotFound with default headers values
func NewGetQualityFormsSurveyVersionsNotFound() *GetQualityFormsSurveyVersionsNotFound {
	return &GetQualityFormsSurveyVersionsNotFound{}
}

/*
GetQualityFormsSurveyVersionsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetQualityFormsSurveyVersionsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions not found response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions not found response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions not found response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions not found response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions not found response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetQualityFormsSurveyVersionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsRequestTimeout creates a GetQualityFormsSurveyVersionsRequestTimeout with default headers values
func NewGetQualityFormsSurveyVersionsRequestTimeout() *GetQualityFormsSurveyVersionsRequestTimeout {
	return &GetQualityFormsSurveyVersionsRequestTimeout{}
}

/*
GetQualityFormsSurveyVersionsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetQualityFormsSurveyVersionsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions request timeout response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions request timeout response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions request timeout response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions request timeout response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions request timeout response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetQualityFormsSurveyVersionsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsRequestEntityTooLarge creates a GetQualityFormsSurveyVersionsRequestEntityTooLarge with default headers values
func NewGetQualityFormsSurveyVersionsRequestEntityTooLarge() *GetQualityFormsSurveyVersionsRequestEntityTooLarge {
	return &GetQualityFormsSurveyVersionsRequestEntityTooLarge{}
}

/*
GetQualityFormsSurveyVersionsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetQualityFormsSurveyVersionsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions request entity too large response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions request entity too large response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions request entity too large response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions request entity too large response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions request entity too large response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsUnsupportedMediaType creates a GetQualityFormsSurveyVersionsUnsupportedMediaType with default headers values
func NewGetQualityFormsSurveyVersionsUnsupportedMediaType() *GetQualityFormsSurveyVersionsUnsupportedMediaType {
	return &GetQualityFormsSurveyVersionsUnsupportedMediaType{}
}

/*
GetQualityFormsSurveyVersionsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityFormsSurveyVersionsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions unsupported media type response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions unsupported media type response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions unsupported media type response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions unsupported media type response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions unsupported media type response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsTooManyRequests creates a GetQualityFormsSurveyVersionsTooManyRequests with default headers values
func NewGetQualityFormsSurveyVersionsTooManyRequests() *GetQualityFormsSurveyVersionsTooManyRequests {
	return &GetQualityFormsSurveyVersionsTooManyRequests{}
}

/*
GetQualityFormsSurveyVersionsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetQualityFormsSurveyVersionsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions too many requests response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions too many requests response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions too many requests response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get quality forms survey versions too many requests response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get quality forms survey versions too many requests response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetQualityFormsSurveyVersionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsInternalServerError creates a GetQualityFormsSurveyVersionsInternalServerError with default headers values
func NewGetQualityFormsSurveyVersionsInternalServerError() *GetQualityFormsSurveyVersionsInternalServerError {
	return &GetQualityFormsSurveyVersionsInternalServerError{}
}

/*
GetQualityFormsSurveyVersionsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityFormsSurveyVersionsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions internal server error response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions internal server error response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions internal server error response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms survey versions internal server error response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality forms survey versions internal server error response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetQualityFormsSurveyVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsServiceUnavailable creates a GetQualityFormsSurveyVersionsServiceUnavailable with default headers values
func NewGetQualityFormsSurveyVersionsServiceUnavailable() *GetQualityFormsSurveyVersionsServiceUnavailable {
	return &GetQualityFormsSurveyVersionsServiceUnavailable{}
}

/*
GetQualityFormsSurveyVersionsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityFormsSurveyVersionsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions service unavailable response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions service unavailable response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions service unavailable response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms survey versions service unavailable response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality forms survey versions service unavailable response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetQualityFormsSurveyVersionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsSurveyVersionsGatewayTimeout creates a GetQualityFormsSurveyVersionsGatewayTimeout with default headers values
func NewGetQualityFormsSurveyVersionsGatewayTimeout() *GetQualityFormsSurveyVersionsGatewayTimeout {
	return &GetQualityFormsSurveyVersionsGatewayTimeout{}
}

/*
GetQualityFormsSurveyVersionsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetQualityFormsSurveyVersionsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get quality forms survey versions gateway timeout response has a 2xx status code
func (o *GetQualityFormsSurveyVersionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get quality forms survey versions gateway timeout response has a 3xx status code
func (o *GetQualityFormsSurveyVersionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get quality forms survey versions gateway timeout response has a 4xx status code
func (o *GetQualityFormsSurveyVersionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get quality forms survey versions gateway timeout response has a 5xx status code
func (o *GetQualityFormsSurveyVersionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get quality forms survey versions gateway timeout response a status code equal to that given
func (o *GetQualityFormsSurveyVersionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetQualityFormsSurveyVersionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/surveys/{formId}/versions][%d] getQualityFormsSurveyVersionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormsSurveyVersionsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsSurveyVersionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
