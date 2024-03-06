package main

import (
	"errors"
	"net/http"
	"strconv"

	// New import
	// New import
	"aitunews.kz/snippetbox/pkg/forms"
	"aitunews.kz/snippetbox/pkg/models"
)

func (app *application) showService(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.services.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Service: s,
	})

}

func (app *application) createServiceForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
}

func (app *application) createService(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "master", "price")
	form.MaxLength("title", 100)
	form.MaxLength("master", 100)
	//form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	i, err1 := strconv.Atoi(form.Get("price"))
	if err1 != nil {
		return
	}

	id, err := app.services.Insert(form.Get("title"), form.Get("content"), form.Get("master"), i)
	if err != nil && id == 0 {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Service successfully created!")

	http.Redirect(w, r, "/services", http.StatusSeeOther)

}

func (app *application) updateServiceForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "update.page.tmpl", &templateData{
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
}

func (app *application) updateService(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "price")

	if !form.Valid() {
		app.render(w, r, "update.page.tmpl", &templateData{Form: form})
		return
	}

	i, err1 := strconv.Atoi(form.Get("price"))
	if err1 != nil {
		return
	}

	err = app.services.Update(form.Get("title"), i)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Service successfully updated!")

	http.Redirect(w, r, "/services", http.StatusSeeOther)

}

func (app *application) deleteServiceForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "delete.page.tmpl", &templateData{
		// Pass a new empty forms.Form object to the template.
		Form: forms.New(nil),
	})
}

func (app *application) deleteService(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title")

	if !form.Valid() {
		app.render(w, r, "delete.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.services.Delete(form.Get("title"))
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.session.Put(r, "flash", "Service successfully deleted!")

	http.Redirect(w, r, "/services", http.StatusSeeOther)

}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	s, err := app.snippets.Latest("")
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) servicess(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/services" {
		app.notFound(w)
		return
	}
	s, err := app.services.Latest("services")
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "services.page.tmpl", &templateData{
		Services: s,
	})
}

func (app *application) applicants(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/applicants" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest("applicants")
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "applicants.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) products(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/products" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest("products")
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "products.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) reviews(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reviews" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest("reviews")
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "reviews.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Latest("about")
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "about.page.tmpl", &templateData{
		Snippets: s,
	})
}

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form := forms.New(r.PostForm)
	form.Required("full_name", "email", "password", "phone")
	form.MaxLength("full_name", 255)
	form.MaxLength("email", 255)
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)
	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}
	err = app.users.Insert(form.Get("full_name"), form.Get("email"), form.Get("phone"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.Errors.Add("email", "Address is already in use")
			app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.session.Put(r, "flash", "Your signup was successful. Please log in.")
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	// Check whether the credentials are valid. If they're not, add a generic error
	// message to the form failures map and re-display the login page.
	form := forms.New(r.PostForm)
	id, err := app.users.Authenticate(form.Get("email"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Email or Password is incorrect")
			app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Add the ID of the current user to the session, so that they are now 'logged
	// in'.
	app.session.Put(r, "authenticatedUserID", id)
	// Redirect the user to the create snippet page.
	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	// Remove the authenticatedUserID from the session data so that the user is
	// 'logged out'.
	app.session.Remove(r, "authenticatedUserID")
	app.session.Remove(r, "authenticatedUserRole")
	// Add a flash message to the session to confirm to the user that they've been
	// logged out.
	app.session.Put(r, "flash", "You've been logged out successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
