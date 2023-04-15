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

// PostKnowledgeKnowledgebaseLabelsReader is a Reader for the PostKnowledgeKnowledgebaseLabels structure.
type PostKnowledgeKnowledgebaseLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostKnowledgeKnowledgebaseLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostKnowledgeKnowledgebaseLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostKnowledgeKnowledgebaseLabelsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostKnowledgeKnowledgebaseLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostKnowledgeKnowledgebaseLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostKnowledgeKnowledgebaseLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostKnowledgeKnowledgebaseLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostKnowledgeKnowledgebaseLabelsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostKnowledgeKnowledgebaseLabelsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostKnowledgeKnowledgebaseLabelsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostKnowledgeKnowledgebaseLabelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostKnowledgeKnowledgebaseLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostKnowledgeKnowledgebaseLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostKnowledgeKnowledgebaseLabelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostKnowledgeKnowledgebaseLabelsOK creates a PostKnowledgeKnowledgebaseLabelsOK with default headers values
func NewPostKnowledgeKnowledgebaseLabelsOK() *PostKnowledgeKnowledgebaseLabelsOK {
	return &PostKnowledgeKnowledgebaseLabelsOK{}
}

/*
PostKnowledgeKnowledgebaseLabelsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostKnowledgeKnowledgebaseLabelsOK struct {
	Payload *models.LabelResponse
}

// IsSuccess returns true when this post knowledge knowledgebase labels o k response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post knowledge knowledgebase labels o k response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels o k response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post knowledge knowledgebase labels o k response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels o k response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsOK  %+v", 200, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) GetPayload() *models.LabelResponse {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsCreated creates a PostKnowledgeKnowledgebaseLabelsCreated with default headers values
func NewPostKnowledgeKnowledgebaseLabelsCreated() *PostKnowledgeKnowledgebaseLabelsCreated {
	return &PostKnowledgeKnowledgebaseLabelsCreated{}
}

/*
PostKnowledgeKnowledgebaseLabelsCreated describes a response with status code 201, with default header values.

Label created
*/
type PostKnowledgeKnowledgebaseLabelsCreated struct {
	Payload *models.LabelResponse
}

// IsSuccess returns true when this post knowledge knowledgebase labels created response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post knowledge knowledgebase labels created response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels created response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post knowledge knowledgebase labels created response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels created response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsCreated  %+v", 201, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsCreated  %+v", 201, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) GetPayload() *models.LabelResponse {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LabelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsBadRequest creates a PostKnowledgeKnowledgebaseLabelsBadRequest with default headers values
func NewPostKnowledgeKnowledgebaseLabelsBadRequest() *PostKnowledgeKnowledgebaseLabelsBadRequest {
	return &PostKnowledgeKnowledgebaseLabelsBadRequest{}
}

/*
PostKnowledgeKnowledgebaseLabelsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostKnowledgeKnowledgebaseLabelsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels bad request response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels bad request response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels bad request response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels bad request response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels bad request response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsUnauthorized creates a PostKnowledgeKnowledgebaseLabelsUnauthorized with default headers values
func NewPostKnowledgeKnowledgebaseLabelsUnauthorized() *PostKnowledgeKnowledgebaseLabelsUnauthorized {
	return &PostKnowledgeKnowledgebaseLabelsUnauthorized{}
}

/*
PostKnowledgeKnowledgebaseLabelsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostKnowledgeKnowledgebaseLabelsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels unauthorized response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels unauthorized response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels unauthorized response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels unauthorized response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels unauthorized response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsForbidden creates a PostKnowledgeKnowledgebaseLabelsForbidden with default headers values
func NewPostKnowledgeKnowledgebaseLabelsForbidden() *PostKnowledgeKnowledgebaseLabelsForbidden {
	return &PostKnowledgeKnowledgebaseLabelsForbidden{}
}

/*
PostKnowledgeKnowledgebaseLabelsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostKnowledgeKnowledgebaseLabelsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels forbidden response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels forbidden response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels forbidden response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels forbidden response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels forbidden response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsForbidden  %+v", 403, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsNotFound creates a PostKnowledgeKnowledgebaseLabelsNotFound with default headers values
func NewPostKnowledgeKnowledgebaseLabelsNotFound() *PostKnowledgeKnowledgebaseLabelsNotFound {
	return &PostKnowledgeKnowledgebaseLabelsNotFound{}
}

/*
PostKnowledgeKnowledgebaseLabelsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostKnowledgeKnowledgebaseLabelsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels not found response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels not found response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels not found response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels not found response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels not found response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsNotFound  %+v", 404, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsRequestTimeout creates a PostKnowledgeKnowledgebaseLabelsRequestTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLabelsRequestTimeout() *PostKnowledgeKnowledgebaseLabelsRequestTimeout {
	return &PostKnowledgeKnowledgebaseLabelsRequestTimeout{}
}

/*
PostKnowledgeKnowledgebaseLabelsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostKnowledgeKnowledgebaseLabelsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels request timeout response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels request timeout response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels request timeout response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels request timeout response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels request timeout response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsConflict creates a PostKnowledgeKnowledgebaseLabelsConflict with default headers values
func NewPostKnowledgeKnowledgebaseLabelsConflict() *PostKnowledgeKnowledgebaseLabelsConflict {
	return &PostKnowledgeKnowledgebaseLabelsConflict{}
}

/*
PostKnowledgeKnowledgebaseLabelsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PostKnowledgeKnowledgebaseLabelsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels conflict response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels conflict response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels conflict response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels conflict response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels conflict response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsConflict  %+v", 409, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsConflict  %+v", 409, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge creates a PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge with default headers values
func NewPostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge() *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge {
	return &PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge{}
}

/*
PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels request entity too large response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels request entity too large response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels request entity too large response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels request entity too large response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels request entity too large response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsUnsupportedMediaType creates a PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType with default headers values
func NewPostKnowledgeKnowledgebaseLabelsUnsupportedMediaType() *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType {
	return &PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType{}
}

/*
PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels unsupported media type response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels unsupported media type response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels unsupported media type response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels unsupported media type response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels unsupported media type response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsTooManyRequests creates a PostKnowledgeKnowledgebaseLabelsTooManyRequests with default headers values
func NewPostKnowledgeKnowledgebaseLabelsTooManyRequests() *PostKnowledgeKnowledgebaseLabelsTooManyRequests {
	return &PostKnowledgeKnowledgebaseLabelsTooManyRequests{}
}

/*
PostKnowledgeKnowledgebaseLabelsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostKnowledgeKnowledgebaseLabelsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels too many requests response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels too many requests response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels too many requests response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post knowledge knowledgebase labels too many requests response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post knowledge knowledgebase labels too many requests response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsInternalServerError creates a PostKnowledgeKnowledgebaseLabelsInternalServerError with default headers values
func NewPostKnowledgeKnowledgebaseLabelsInternalServerError() *PostKnowledgeKnowledgebaseLabelsInternalServerError {
	return &PostKnowledgeKnowledgebaseLabelsInternalServerError{}
}

/*
PostKnowledgeKnowledgebaseLabelsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostKnowledgeKnowledgebaseLabelsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels internal server error response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels internal server error response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels internal server error response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post knowledge knowledgebase labels internal server error response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post knowledge knowledgebase labels internal server error response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsServiceUnavailable creates a PostKnowledgeKnowledgebaseLabelsServiceUnavailable with default headers values
func NewPostKnowledgeKnowledgebaseLabelsServiceUnavailable() *PostKnowledgeKnowledgebaseLabelsServiceUnavailable {
	return &PostKnowledgeKnowledgebaseLabelsServiceUnavailable{}
}

/*
PostKnowledgeKnowledgebaseLabelsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostKnowledgeKnowledgebaseLabelsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels service unavailable response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels service unavailable response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels service unavailable response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post knowledge knowledgebase labels service unavailable response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post knowledge knowledgebase labels service unavailable response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostKnowledgeKnowledgebaseLabelsGatewayTimeout creates a PostKnowledgeKnowledgebaseLabelsGatewayTimeout with default headers values
func NewPostKnowledgeKnowledgebaseLabelsGatewayTimeout() *PostKnowledgeKnowledgebaseLabelsGatewayTimeout {
	return &PostKnowledgeKnowledgebaseLabelsGatewayTimeout{}
}

/*
PostKnowledgeKnowledgebaseLabelsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostKnowledgeKnowledgebaseLabelsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post knowledge knowledgebase labels gateway timeout response has a 2xx status code
func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post knowledge knowledgebase labels gateway timeout response has a 3xx status code
func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post knowledge knowledgebase labels gateway timeout response has a 4xx status code
func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post knowledge knowledgebase labels gateway timeout response has a 5xx status code
func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post knowledge knowledgebase labels gateway timeout response a status code equal to that given
func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/labels][%d] postKnowledgeKnowledgebaseLabelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostKnowledgeKnowledgebaseLabelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
