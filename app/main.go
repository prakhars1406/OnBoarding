package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	adminHandler "gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/admin"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/config"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/factory"
	"gitlab.com/coffeebeansdev/wru/wru-backend-new-merchant-on-boarding/logger"
	"log"
	"net/http"
)

var AppConfig *config.Configuration = nil

func init() {
	AppConfig = config.GetAppConfiguration()
}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	dataStoreClient := factory.MongoConnector()
	configuration := config.GetAppConfiguration()
	allowedHeaders := handlers.AllowedHeaders([]string{config.ORIGIN, config.ACCEPT, config.CONTENT_TYPE_HEADER, config.AUTHORIZATION, config.DATE_USED, config.X_REQUESTED_WITH})
	allowedOrigins := handlers.AllowedOrigins([]string{config.ALL})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})
	fmt.Println(AppConfig.MONGO_SERVER)
	authoriseHandler := adminHandler.AuthoriseHandler(dataStoreClient)
	registerHandler := adminHandler.RegisterHandler(dataStoreClient)
	//onBoardHandler := adminHandler.OnBoardHandler(dataStoreClient)
	getAllMerchants:=adminHandler.GetAllMerchants(dataStoreClient)
	actionHandle:=adminHandler.ActionHandle(dataStoreClient)
	addNewMerchantConfig:=adminHandler.AddNewMerchantConfig(dataStoreClient)
	addNewPageType:=adminHandler.AddNewPageType(dataStoreClient)
	addMerchantPage:=adminHandler.AddMerchantPage(dataStoreClient)
	getPageRecommendations:=adminHandler.GetPageRecommendations(dataStoreClient)
	addPageRecommendations:=adminHandler.AddPageRecommendations(dataStoreClient)
	getRouteInfo:=adminHandler.GetRouteInfo(dataStoreClient)

	logger.Init(config.DIR_NAME, config.MAX_FILES, config.FILES_TO_DELETE, config.MAX_SIZE, false, configuration.APP_MODE, config.LOGS_BUCKET,configuration.LOG_STRATEGY)
	fileServer := http.FileServer(http.Dir("./swaggerui/"))
	router.PathPrefix("/api/").Handler(http.StripPrefix("/api/", fileServer))
	router.HandleFunc("/authorise", authoriseHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", registerHandler).Methods(http.MethodPost)
	router.HandleFunc("/merchants", getAllMerchants).Methods(http.MethodGet)
	router.HandleFunc("/{merchantId}/action", actionHandle).Methods(http.MethodPost)
	router.HandleFunc("/{merchantId}/merchantConfig", addNewMerchantConfig).Methods(http.MethodPost)
	router.HandleFunc("/{merchantId}/pageType", addNewPageType).Methods(http.MethodPost)
	router.HandleFunc("/{merchantId}/merchantPages", addMerchantPage).Methods(http.MethodPost)
	router.HandleFunc("/pageRecommendations", getPageRecommendations).Methods(http.MethodGet)
	router.HandleFunc("/pageRecommendations", addPageRecommendations).Methods(http.MethodPost)
	router.HandleFunc("/{merchantId}/routeInfo", getRouteInfo).Methods(http.MethodGet)

	fmt.Println("Register Service loaded successfully ")
	logger.Info("Register Service loaded successfully ")
	log.Fatal(http.ListenAndServe(config.KEY_SEPARATOR+configuration.PORT, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
func AddContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
