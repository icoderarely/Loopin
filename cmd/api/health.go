package main

import (
	"log"
	"net/http"
)

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"status":  "ok",
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		log.Print(err.Error())
		// writeJSONError(w, http.StatusBadRequest, "Errorrrr")
	}
}
