// Package router
package router

import (
	"context"
	"encoding/json"

	"net/http"
	"runtime/debug"

	"gitlab.com/willysihombing/task-c3/internal/appctx"
	"gitlab.com/willysihombing/task-c3/internal/bootstrap"
	"gitlab.com/willysihombing/task-c3/internal/consts"
	"gitlab.com/willysihombing/task-c3/internal/handler"
	"gitlab.com/willysihombing/task-c3/internal/middleware"
	"gitlab.com/willysihombing/task-c3/internal/ucase"
	"gitlab.com/willysihombing/task-c3/internal/ucase/activity"
	"gitlab.com/willysihombing/task-c3/internal/ucase/auth"
	user "gitlab.com/willysihombing/task-c3/internal/ucase/users"
	"gitlab.com/willysihombing/task-c3/pkg/logger"
	"gitlab.com/willysihombing/task-c3/pkg/routerkit"

	//"gitlab.com/willysihombing/task-c3/pkg/mariadb"
	//"gitlab.com/willysihombing/task-c3/internal/repositories"
	//"gitlab.com/willysihombing/task-c3/internal/ucase/example"

	activitySupabase "gitlab.com/willysihombing/task-c3/internal/connector/supabase/activity"
	userSupabase "gitlab.com/willysihombing/task-c3/internal/connector/supabase/user"
	ucaseContract "gitlab.com/willysihombing/task-c3/internal/ucase/contract"
)

type router struct {
	config *appctx.Config
	router *routerkit.Router
}

// NewRouter initialize new router wil return Router Interface
func NewRouter(cfg *appctx.Config) Router {
	bootstrap.RegistryMessage()
	bootstrap.RegistryLogger(cfg)

	return &router{
		config: cfg,
		router: routerkit.NewRouter(routerkit.WithServiceName(cfg.App.AppName)),
	}
}

func (rtr *router) handle(hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)

				w.WriteHeader(http.StatusInternalServerError)
				res := appctx.Response{
					Code: consts.CodeInternalServerError,
				}

				res.GenerateMessage()
				logger.Error(logger.MessageFormat("error %v", string(debug.Stack())))
				json.NewEncoder(w).Encode(res)
				return
			}
		}()

		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
			"path":      r.URL.Path,
			"remote_ip": r.RemoteAddr,
			"method":    r.Method,
		})

		req := r.WithContext(ctx)

		if status := middleware.FilterFunc(rtr.config, req, mdws); status != 200 {
			rtr.response(w, appctx.Response{
				Code: status,
			})

			return
		}

		resp := hfn(req, svc, rtr.config)
		resp.Lang = rtr.defaultLang(req.Header.Get(consts.HeaderLanguageKey))
		rtr.response(w, resp)
	}
}

// response prints as a json and formatted string for DGP legacy
func (rtr *router) response(w http.ResponseWriter, resp appctx.Response) {

	w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)

	defer func() {
		resp.GenerateMessage()
		w.WriteHeader(resp.GetCode())
		json.NewEncoder(w).Encode(resp)
	}()

	return

}

// Route preparing http router and will return mux router object
func (rtr *router) Route() *routerkit.Router {

	root := rtr.router.PathPrefix("/").Subrouter()
	api := root.PathPrefix("/api").Subrouter()
	v1 := api.PathPrefix("/v1").Subrouter()
	// ex := root.PathPrefix("/external").Subrouter()
	liveness := root.PathPrefix("/").Subrouter()
	//ad := root.PathPrefix("/users").Subrouter()

	// open tracer setup
	bootstrap.RegistryOpenTracing(rtr.config)

	//db := bootstrap.RegistryMariaMasterSlave(rtr.config.WriteDB, rtr.config.ReadDB, rtr.config.App.Timezone)

	// use case
	healthy := ucase.NewHealthCheck()

	//DB
	userDb := userSupabase.NewClient(rtr.config)
	activityDb := activitySupabase.NewClient(rtr.config)

	//middlewares
	middlewares := middleware.NewValidateBearer(userDb)

	// Users
	createNewUser := user.CreateUser(userDb)
	getAllUser := user.GetAllUser(userDb)
	getDetailUser := user.GetDetailUser(userDb)
	deleteUser := user.DeleteUser(userDb)
	updateUser := user.UpdateUser(userDb)

	// Intern
	createNewActivity := activity.CreateActivity(activityDb)
	getAllActivity := activity.GetAllActivity(activityDb)
	GetAllActivityByMonth := activity.GetAllActivityMonth(activityDb)
	updateActivity := activity.UpdateActivity(activityDb)
	deleteActivity := activity.DeleteActivity(activityDb)
	verifyActivity := activity.NewVerifyActivity(activityDb)

	// Login
	loginUcase := auth.NewLoginUser(userDb)

	//generateJWT := middleware.CreateJWT()

	// healthy
	liveness.HandleFunc("/liveness", rtr.handle(
		handler.HttpRequest,
		healthy,
	)).Methods(http.MethodGet)

	// Users
	// create users
	v1.HandleFunc("/user/new", rtr.handle(
		handler.HttpRequest,
		createNewUser,
		middlewares.ValidateToken,
	)).Methods(http.MethodPost)

	// read all users
	v1.HandleFunc("/user", rtr.handle(
		handler.HttpRequest,
		getAllUser,
		middlewares.ValidateToken,
	)).Methods(http.MethodGet)

	// read detail users
	v1.HandleFunc("/user/{id}", rtr.handle(
		handler.HttpRequest,
		getDetailUser,
		middlewares.ValidateToken,
	)).Methods(http.MethodGet)

	// Update Users
	v1.HandleFunc("/user/update/{id}", rtr.handle(
		handler.HttpRequest,
		updateUser,
		middlewares.ValidateToken,
	)).Methods(http.MethodPut)

	// Delete Users
	v1.HandleFunc("/user/{id}", rtr.handle(
		handler.HttpRequest,
		deleteUser,
		middlewares.ValidateToken,
	)).Methods(http.MethodDelete)

	// Activity
	// Create Activity
	v1.HandleFunc("/activity/new", rtr.handle(
		handler.HttpRequest,
		createNewActivity,
		middlewares.ValidateToken,
	)).Methods(http.MethodPost)
	// Get All Activity
	v1.HandleFunc("/activity", rtr.handle(
		handler.HttpRequest,
		getAllActivity,
		middlewares.ValidateToken,
	)).Methods(http.MethodGet)
	v1.HandleFunc("/activity/allByMonth", rtr.handle(
		handler.HttpRequest,
		GetAllActivityByMonth,
		middlewares.ValidateToken,
	)).Methods(http.MethodGet)
	// Update Activity
	v1.HandleFunc("/activity/update/{id}", rtr.handle(
		handler.HttpRequest,
		updateActivity,
		middlewares.ValidateToken,
	)).Methods(http.MethodPut)
	// Delete Activity
	v1.HandleFunc("/activity/delete/{id}", rtr.handle(
		handler.HttpRequest,
		deleteActivity,
		middlewares.ValidateToken,
	)).Methods(http.MethodDelete)

	v1.HandleFunc("/activity/verify", rtr.handle(
		handler.HttpRequest,
		verifyActivity,
		middlewares.ValidateToken,
	)).Methods(http.MethodPost)

	// Login
	v1.HandleFunc("/login", rtr.handle(
		handler.HttpRequest,
		loginUcase,
	)).Methods(http.MethodPost)

	return rtr.router
}

func (rtr *router) defaultLang(l string) string {

	if len(l) == 0 {
		return rtr.config.App.DefaultLang
	}

	return l
}
