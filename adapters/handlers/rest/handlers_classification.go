//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package rest

import (
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/weaviate/weaviate/adapters/handlers/rest/operations"
	"github.com/weaviate/weaviate/adapters/handlers/rest/operations/classifications"
	"github.com/weaviate/weaviate/entities/models"
	autherrs "github.com/weaviate/weaviate/usecases/auth/authorization/errors"
	"github.com/weaviate/weaviate/usecases/classification"
	"github.com/weaviate/weaviate/usecases/monitoring"
)

func setupClassificationHandlers(api *operations.WeaviateAPI,
	classifier *classification.Classifier, metrics *monitoring.PrometheusMetrics, logger logrus.FieldLogger,
) {
	metricRequestsTotal := newClassificationRequestsTotal(metrics, logger)
	api.ClassificationsClassificationsGetHandler = classifications.ClassificationsGetHandlerFunc(
		func(params classifications.ClassificationsGetParams, principal *models.Principal) middleware.Responder {
			res, err := classifier.Get(params.HTTPRequest.Context(), principal, strfmt.UUID(params.ID))
			if err != nil {
				metricRequestsTotal.logError("", err)
				var forbidden autherrs.Forbidden
				switch {
				case errors.As(err, &forbidden):
					return classifications.NewClassificationsGetForbidden().
						WithPayload(errPayloadFromSingleErr(err))
				default:
					return classifications.NewClassificationsPostBadRequest().WithPayload(errPayloadFromSingleErr(err))
				}
			}

			if res == nil {
				metricRequestsTotal.logUserError("")
				return classifications.NewClassificationsGetNotFound()
			}

			metricRequestsTotal.logOk("")
			return classifications.NewClassificationsGetOK().WithPayload(res)
		},
	)

	api.ClassificationsClassificationsPostHandler = classifications.ClassificationsPostHandlerFunc(
		func(params classifications.ClassificationsPostParams, principal *models.Principal) middleware.Responder {
			res, err := classifier.Schedule(params.HTTPRequest.Context(), principal, *params.Params)
			if err != nil {
				metricRequestsTotal.logUserError("")

				var forbidden autherrs.Forbidden
				switch {
				case errors.As(err, &forbidden):
					return classifications.NewClassificationsPostForbidden().
						WithPayload(errPayloadFromSingleErr(err))
				default:
					return classifications.NewClassificationsPostBadRequest().WithPayload(errPayloadFromSingleErr(err))
				}
			}

			metricRequestsTotal.logOk("")
			return classifications.NewClassificationsPostCreated().WithPayload(res)
		},
	)
}

type classificationRequestsTotal struct {
	*restApiRequestsTotalImpl
}

func newClassificationRequestsTotal(metrics *monitoring.PrometheusMetrics, logger logrus.FieldLogger) restApiRequestsTotal {
	return &classificationRequestsTotal{
		restApiRequestsTotalImpl: &restApiRequestsTotalImpl{newRequestsTotalMetric(metrics, "rest"), "rest", "classification", logger},
	}
}

func (e *classificationRequestsTotal) logError(className string, err error) {
	e.logServerError(className, err)
}
