package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/tapi/models"
)

/*CreateTeamCreated Newly created team

swagger:response createTeamCreated
*/
type CreateTeamCreated struct {

	// In: body
	Payload *models.Team `json:"body,omitempty"`
}

// NewCreateTeamCreated creates CreateTeamCreated with default headers values
func NewCreateTeamCreated() *CreateTeamCreated {
	return &CreateTeamCreated{}
}

// WithPayload adds the payload to the create team created response
func (o *CreateTeamCreated) WithPayload(payload *models.Team) *CreateTeamCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team created response
func (o *CreateTeamCreated) SetPayload(payload *models.Team) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTeamBadRequest User sent a bad request

swagger:response createTeamBadRequest
*/
type CreateTeamBadRequest struct {

	// In: body
	Payload *models.BadRequest `json:"body,omitempty"`
}

// NewCreateTeamBadRequest creates CreateTeamBadRequest with default headers values
func NewCreateTeamBadRequest() *CreateTeamBadRequest {
	return &CreateTeamBadRequest{}
}

// WithPayload adds the payload to the create team bad request response
func (o *CreateTeamBadRequest) WithPayload(payload *models.BadRequest) *CreateTeamBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team bad request response
func (o *CreateTeamBadRequest) SetPayload(payload *models.BadRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTeamUnauthorized User not authorized

swagger:response createTeamUnauthorized
*/
type CreateTeamUnauthorized struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewCreateTeamUnauthorized creates CreateTeamUnauthorized with default headers values
func NewCreateTeamUnauthorized() *CreateTeamUnauthorized {
	return &CreateTeamUnauthorized{}
}

// WithPayload adds the payload to the create team unauthorized response
func (o *CreateTeamUnauthorized) WithPayload(payload *models.Unauthorized) *CreateTeamUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team unauthorized response
func (o *CreateTeamUnauthorized) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTeamForbidden User does not have the credentials to access this resource


swagger:response createTeamForbidden
*/
type CreateTeamForbidden struct {

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewCreateTeamForbidden creates CreateTeamForbidden with default headers values
func NewCreateTeamForbidden() *CreateTeamForbidden {
	return &CreateTeamForbidden{}
}

// WithPayload adds the payload to the create team forbidden response
func (o *CreateTeamForbidden) WithPayload(payload *models.Unauthorized) *CreateTeamForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team forbidden response
func (o *CreateTeamForbidden) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTeamDefault User not authorized

swagger:response createTeamDefault
*/
type CreateTeamDefault struct {
	_statusCode int

	// In: body
	Payload *models.Unauthorized `json:"body,omitempty"`
}

// NewCreateTeamDefault creates CreateTeamDefault with default headers values
func NewCreateTeamDefault(code int) *CreateTeamDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTeamDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create team default response
func (o *CreateTeamDefault) WithStatusCode(code int) *CreateTeamDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create team default response
func (o *CreateTeamDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create team default response
func (o *CreateTeamDefault) WithPayload(payload *models.Unauthorized) *CreateTeamDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create team default response
func (o *CreateTeamDefault) SetPayload(payload *models.Unauthorized) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTeamDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}