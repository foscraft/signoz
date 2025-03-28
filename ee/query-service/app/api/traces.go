package api

import (
	"net/http"

	"github.com/SigNoz/signoz/ee/query-service/app/db"
	"github.com/SigNoz/signoz/ee/query-service/model"
	baseapp "github.com/SigNoz/signoz/pkg/query-service/app"
	basemodel "github.com/SigNoz/signoz/pkg/query-service/model"
	"go.uber.org/zap"
)

func (ah *APIHandler) searchTraces(w http.ResponseWriter, r *http.Request) {

	if !ah.CheckFeature(basemodel.SmartTraceDetail) {
		zap.L().Info("SmartTraceDetail feature is not enabled in this plan")
		ah.APIHandler.SearchTraces(w, r)
		return
	}
	searchTracesParams, err := baseapp.ParseSearchTracesParams(r)
	if err != nil {
		RespondError(w, &model.ApiError{Typ: model.ErrorBadData, Err: err}, "Error reading params")
		return
	}

	result, err := ah.opts.DataConnector.SearchTraces(r.Context(), searchTracesParams, db.SmartTraceAlgorithm)
	if ah.HandleError(w, err, http.StatusBadRequest) {
		return
	}

	ah.WriteJSON(w, r, result)

}
