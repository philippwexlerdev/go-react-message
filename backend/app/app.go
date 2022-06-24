package app

import (
	"backend-test-template/app/handler"
	"backend-test-template/app/repo"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

type Application struct {
	Router   *mux.Router
	Database *gorm.DB
	// TODO: Add storage for comments
}

func NewApplication(r *mux.Router, db *gorm.DB) *Application {
	myApp := &Application{Router: r, Database: db}
	myApp.setupRoutes()
	return myApp
}

func (a *Application) setupRoutes() {
	// TODO: Setup routes and link to functions

	//REPO
	userRepo := repo.NewUserRepo(a.Database)
	userTokenRepo := repo.NewUserTokenRepo(a.Database)
	postRepo := repo.NewPostRepo(a.Database)
	postResponseRepo := repo.NewPostResponseRepo(a.Database)

	//HANDLER
	userHandler := handler.NewUserHandlers(userRepo, userTokenRepo)
	postHandler := handler.NewPostHandlers(postRepo, userRepo)
	postResponseHandler := handler.NewPostResponseHandlers(postResponseRepo, userRepo, postRepo)

	middleware := handler.NewMiddleware(userTokenRepo)
	databaseMigrationHandler := handler.NewMigrationHandler(a.Database)

	a.Router.HandleFunc("/db/migrate", databaseMigrationHandler.MigrateDatabase).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/user/create", userHandler.Create).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/user/login", userHandler.Login).Methods(http.MethodPost)

	a.Router.HandleFunc("/api/user/get/{id}", middleware.Authen(userHandler.GetUserById)).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/user/update/{id}", middleware.Authen(userHandler.UpdateUserByID)).Methods(http.MethodPut)
	a.Router.HandleFunc("/api/user/delete/{id}", middleware.Authen(userHandler.DeleteUserByID)).Methods(http.MethodDelete)

	a.Router.HandleFunc("/api/post/create", middleware.Authen(postHandler.Create)).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/post/get/{id}", middleware.Authen(postHandler.GetPostById)).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/post/update/{id}", middleware.Authen(postHandler.UpdatePostByID)).Methods(http.MethodPut)
	a.Router.HandleFunc("/api/post/get-all", middleware.Authen(postHandler.GetAll)).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/post/delete/{id}", middleware.Authen(postHandler.DeletePostByID)).Methods(http.MethodDelete)

	a.Router.HandleFunc("/api/post-response/create", middleware.Authen(postResponseHandler.Create)).Methods(http.MethodPost)
	a.Router.HandleFunc("/api/post-response/get/{id}", middleware.Authen(postResponseHandler.GetPostResponseById)).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/post-response/update/{id}", middleware.Authen(postResponseHandler.UpdatePostResponseByID)).Methods(http.MethodPut)
	a.Router.HandleFunc("/api/post-response/get-all/{id}", middleware.Authen(postResponseHandler.GetAllResponseOfAPost)).Methods(http.MethodGet)
	a.Router.HandleFunc("/api/post-response/delete/{id}", middleware.Authen(postResponseHandler.DeletePostResponseByID)).Methods(http.MethodDelete)
}

// Template function, needs to be linked to a route
func (a *Application) getComments(w http.ResponseWriter, r *http.Request) {
	// TODO: Get comments from database...
}
