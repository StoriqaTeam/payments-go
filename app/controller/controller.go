package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/storiqateam/payments/app/controller/requests"
	"github.com/storiqateam/payments/app/services"
)

type Controller struct {
	AuthService *services.AuthService
}

func (c *Controller) HandleSessionCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err := errors.Wrap(err, "error reading body")
		log.Errorf("%+v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	var input requests.SessionCreateInput
	fmt.Printf("Body: %s\n", body)
	if err = json.Unmarshal(body, &input); err != nil {
		err := errors.Wrap(err, "error parsing json")
		log.Errorf("%+v", err)
		http.Error(w, "Can't parse json", http.StatusUnprocessableEntity)
		return
	}
	// s, _ := json.Marshal(input)
	fmt.Printf("Input: %+v\n", input)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Success")
}
