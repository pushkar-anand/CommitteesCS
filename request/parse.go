package request

import (
	"committees/config"
	"committees/helpers"
	"committees/validation"
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

var decoder = schema.NewDecoder()

func ReadFormDataAndValidate(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	err := ReadFormData(r, v)
	if err != nil {
		config.GetLogger().WithError(err).Error("error reading form")
		helpers.InternalError(w)
		return false
	}

	// TODO improve error handling
	ve := validate(v)
	if ve != nil && len(ve) > 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "%v is not valid", ve)
		return false
	}

	return true
}

func ReadFormData(r *http.Request, v interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	return decoder.Decode(v, r.PostForm)
}

func validate(v interface{}) []string {
	ok, fields, err := validation.DoValidation(v)
	if err != nil {
		panic(err)
	}

	if ok && len(fields) == 0 {
		return nil
	}

	return fields
}
