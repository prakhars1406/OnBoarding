package admin

import (
	"encoding/json"
	"github.com/prakhars1406/OnBoarding/config"
	"github.com/prakhars1406/OnBoarding/factory"
	"github.com/prakhars1406/OnBoarding/model"
	"github.com/prakhars1406/OnBoarding/utility"
	"net/http"
	"sort"
)



func GetPageRecommendations(datastoreClient factory.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		writer.Header().Set(config.CONTENT_TYPE_HEADER, config.APPLICATION_JSON)

		allPageRecommendations, err := datastoreClient.GetAllPageRecommendations()
		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		}  else {
			sort.Slice(allPageRecommendations, func(i, j int) bool { return allPageRecommendations[i].PageTypeId < allPageRecommendations[j].PageTypeId })
			err := json.NewEncoder(writer).Encode(allPageRecommendations)
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			} else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}

func AddPageRecommendations(datastoreClient factory.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		writer.Header().Set(config.CONTENT_TYPE_HEADER, config.APPLICATION_JSON)
		var pageRecommendation []model.AllPageRecommendations
		err := json.NewDecoder(request.Body).Decode(&pageRecommendation)
		if err != nil {
			err := json.NewEncoder(writer).Encode(struct {
				Result  bool   `json:"result"`
				Error string `json:"err,omitempty"`
			}{Result: false, Error: err.Error()})
			if err != nil {
				utility.HandleError(writer, http.StatusBadRequest, err)
			} else {
				writer.WriteHeader(http.StatusBadRequest)
			}
			return
		}
		response, err := datastoreClient.AddNewPageRecommendations(pageRecommendation)
		if err != nil {
			utility.HandleError(writer, http.StatusBadRequest, err)
		}  else {
			err := json.NewEncoder(writer).Encode(response)
			if err != nil {
				utility.HandleError(writer, http.StatusInternalServerError, err)
			} else {
				writer.WriteHeader(http.StatusOK)
			}
		}
	}
}