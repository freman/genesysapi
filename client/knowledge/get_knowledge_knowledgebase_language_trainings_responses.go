// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetKnowledgeKnowledgebaseLanguageTrainingsReader is a Reader for the GetKnowledgeKnowledgebaseLanguageTrainings structure.
type GetKnowledgeKnowledgebaseLanguageTrainingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsOK creates a GetKnowledgeKnowledgebaseLanguageTrainingsOK with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsOK() *GetKnowledgeKnowledgebaseLanguageTrainingsOK {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsOK{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsOK struct {
	Payload *models.TrainingListing
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings o k response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings o k response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings o k response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language trainings o k response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings o k response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsOK  %+v", 200, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) GetPayload() *models.TrainingListing {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrainingListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsBadRequest creates a GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsBadRequest() *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings bad request response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings bad request response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings bad request response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings bad request response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings bad request response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized creates a GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized() *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings unauthorized response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings unauthorized response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings unauthorized response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings unauthorized response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings unauthorized response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsForbidden creates a GetKnowledgeKnowledgebaseLanguageTrainingsForbidden with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsForbidden() *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsForbidden{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings forbidden response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings forbidden response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings forbidden response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings forbidden response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings forbidden response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsForbidden  %+v", 403, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsNotFound creates a GetKnowledgeKnowledgebaseLanguageTrainingsNotFound with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsNotFound() *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsNotFound{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings not found response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings not found response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings not found response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings not found response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings not found response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsNotFound  %+v", 404, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout creates a GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout() *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings request timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings request timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings request timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings request timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings request timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge creates a GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge() *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings request entity too large response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings request entity too large response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings request entity too large response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings request entity too large response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings request entity too large response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType creates a GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType() *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings unsupported media type response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings unsupported media type response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings unsupported media type response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings unsupported media type response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings unsupported media type response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests creates a GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests() *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings too many requests response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings too many requests response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings too many requests response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get knowledge knowledgebase language trainings too many requests response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get knowledge knowledgebase language trainings too many requests response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError creates a GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError() *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings internal server error response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings internal server error response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings internal server error response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language trainings internal server error response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase language trainings internal server error response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable creates a GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable() *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings service unavailable response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings service unavailable response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings service unavailable response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language trainings service unavailable response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase language trainings service unavailable response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout creates a GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout with default headers values
func NewGetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout() *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout {
	return &GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout{}
}

/*
GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get knowledge knowledgebase language trainings gateway timeout response has a 2xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get knowledge knowledgebase language trainings gateway timeout response has a 3xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get knowledge knowledgebase language trainings gateway timeout response has a 4xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get knowledge knowledgebase language trainings gateway timeout response has a 5xx status code
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get knowledge knowledgebase language trainings gateway timeout response a status code equal to that given
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/trainings][%d] getKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetKnowledgeKnowledgebaseLanguageTrainingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
