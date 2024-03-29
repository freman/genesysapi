// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetSpeechandtextanalyticsProgramMappingsReader is a Reader for the GetSpeechandtextanalyticsProgramMappings structure.
type GetSpeechandtextanalyticsProgramMappingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsProgramMappingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsProgramMappingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsProgramMappingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsProgramMappingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsProgramMappingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsProgramMappingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsProgramMappingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsProgramMappingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsProgramMappingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsProgramMappingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsProgramMappingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsProgramMappingsOK creates a GetSpeechandtextanalyticsProgramMappingsOK with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsOK() *GetSpeechandtextanalyticsProgramMappingsOK {
	return &GetSpeechandtextanalyticsProgramMappingsOK{}
}

/*
GetSpeechandtextanalyticsProgramMappingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSpeechandtextanalyticsProgramMappingsOK struct {
	Payload *models.ProgramMappings
}

// IsSuccess returns true when this get speechandtextanalytics program mappings o k response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get speechandtextanalytics program mappings o k response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings o k response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics program mappings o k response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings o k response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSpeechandtextanalyticsProgramMappingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsOK) GetPayload() *models.ProgramMappings {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProgramMappings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsBadRequest creates a GetSpeechandtextanalyticsProgramMappingsBadRequest with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsBadRequest() *GetSpeechandtextanalyticsProgramMappingsBadRequest {
	return &GetSpeechandtextanalyticsProgramMappingsBadRequest{}
}

/*
GetSpeechandtextanalyticsProgramMappingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsProgramMappingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings bad request response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings bad request response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings bad request response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings bad request response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings bad request response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsUnauthorized creates a GetSpeechandtextanalyticsProgramMappingsUnauthorized with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsUnauthorized() *GetSpeechandtextanalyticsProgramMappingsUnauthorized {
	return &GetSpeechandtextanalyticsProgramMappingsUnauthorized{}
}

/*
GetSpeechandtextanalyticsProgramMappingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsProgramMappingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings unauthorized response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings unauthorized response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings unauthorized response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings unauthorized response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings unauthorized response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsForbidden creates a GetSpeechandtextanalyticsProgramMappingsForbidden with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsForbidden() *GetSpeechandtextanalyticsProgramMappingsForbidden {
	return &GetSpeechandtextanalyticsProgramMappingsForbidden{}
}

/*
GetSpeechandtextanalyticsProgramMappingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsProgramMappingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings forbidden response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings forbidden response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings forbidden response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings forbidden response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings forbidden response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsNotFound creates a GetSpeechandtextanalyticsProgramMappingsNotFound with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsNotFound() *GetSpeechandtextanalyticsProgramMappingsNotFound {
	return &GetSpeechandtextanalyticsProgramMappingsNotFound{}
}

/*
GetSpeechandtextanalyticsProgramMappingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsProgramMappingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings not found response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings not found response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings not found response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings not found response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings not found response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsRequestTimeout creates a GetSpeechandtextanalyticsProgramMappingsRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsRequestTimeout() *GetSpeechandtextanalyticsProgramMappingsRequestTimeout {
	return &GetSpeechandtextanalyticsProgramMappingsRequestTimeout{}
}

/*
GetSpeechandtextanalyticsProgramMappingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsProgramMappingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings request timeout response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings request timeout response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings request timeout response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings request timeout response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings request timeout response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge creates a GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge() *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge{}
}

/*
GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings request entity too large response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings request entity too large response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings request entity too large response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings request entity too large response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings request entity too large response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType creates a GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType() *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType {
	return &GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType{}
}

/*
GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings unsupported media type response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings unsupported media type response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings unsupported media type response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings unsupported media type response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings unsupported media type response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsTooManyRequests creates a GetSpeechandtextanalyticsProgramMappingsTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsTooManyRequests() *GetSpeechandtextanalyticsProgramMappingsTooManyRequests {
	return &GetSpeechandtextanalyticsProgramMappingsTooManyRequests{}
}

/*
GetSpeechandtextanalyticsProgramMappingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsProgramMappingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings too many requests response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings too many requests response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings too many requests response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get speechandtextanalytics program mappings too many requests response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get speechandtextanalytics program mappings too many requests response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsInternalServerError creates a GetSpeechandtextanalyticsProgramMappingsInternalServerError with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsInternalServerError() *GetSpeechandtextanalyticsProgramMappingsInternalServerError {
	return &GetSpeechandtextanalyticsProgramMappingsInternalServerError{}
}

/*
GetSpeechandtextanalyticsProgramMappingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsProgramMappingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings internal server error response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings internal server error response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings internal server error response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics program mappings internal server error response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics program mappings internal server error response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsServiceUnavailable creates a GetSpeechandtextanalyticsProgramMappingsServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsServiceUnavailable() *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable {
	return &GetSpeechandtextanalyticsProgramMappingsServiceUnavailable{}
}

/*
GetSpeechandtextanalyticsProgramMappingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsProgramMappingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings service unavailable response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings service unavailable response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings service unavailable response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics program mappings service unavailable response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics program mappings service unavailable response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsProgramMappingsGatewayTimeout creates a GetSpeechandtextanalyticsProgramMappingsGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsProgramMappingsGatewayTimeout() *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout {
	return &GetSpeechandtextanalyticsProgramMappingsGatewayTimeout{}
}

/*
GetSpeechandtextanalyticsProgramMappingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsProgramMappingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get speechandtextanalytics program mappings gateway timeout response has a 2xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get speechandtextanalytics program mappings gateway timeout response has a 3xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get speechandtextanalytics program mappings gateway timeout response has a 4xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get speechandtextanalytics program mappings gateway timeout response has a 5xx status code
func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get speechandtextanalytics program mappings gateway timeout response a status code equal to that given
func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/programs/{programId}/mappings][%d] getSpeechandtextanalyticsProgramMappingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsProgramMappingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
