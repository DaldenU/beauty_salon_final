package main

import (
	"net/http"

	"github.com/bmizerany/pat" // New import
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	// Use the nosurf middleware on all our 'dynamic' routes.
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	mux := pat.New()

	mux.Get("/services", dynamicMiddleware.ThenFunc(app.servicess))
	mux.Get("/products", dynamicMiddleware.ThenFunc(app.products))
	mux.Get("/reviews", dynamicMiddleware.ThenFunc(app.reviews))
	mux.Get("/about", dynamicMiddleware.ThenFunc(app.about))

	//	CREATE TABLE appointments (
	//	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	// 	user_id int NOT NULL,
	// 	service_id int NOT NULL,
	// 	time DATETIME NOT NULL
	// 	);

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	mux.Get("/services/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createServiceForm))
	mux.Post("/services/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createService))
	mux.Get("/services/update", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.updateServiceForm))
	mux.Post("/services/update", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.updateService))
	mux.Get("/services/delete", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.deleteServiceForm))
	mux.Post("/services/delete", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.deleteService))
	mux.Get("/services/:id", dynamicMiddleware.ThenFunc(app.showService))

	mux.Get("/appointments", dynamicMiddleware.ThenFunc(app.appointmentss))
	mux.Get("/appointments/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createAppointmentForm))
	mux.Post("/appointments/create", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.createAppointment))
	mux.Get("/appointments/update", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.updateAppointmentForm))
	mux.Post("/appointments/update", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.updateAppointment))
	mux.Get("/appointments/delete", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.deleteAppointmentForm))
	mux.Post("/appointments/delete", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.deleteAppointment))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthentication).ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return standardMiddleware.Then(mux)

}
