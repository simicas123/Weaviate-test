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

package objects

import (
	"context"
	"fmt"

	"github.com/weaviate/weaviate/entities/additional"
	"github.com/weaviate/weaviate/entities/filters"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/usecases/auth/authorization"
	"github.com/weaviate/weaviate/usecases/auth/authorization/filter"
)

type QueryInput struct {
	Class      string
	Offset     int
	Limit      int
	Cursor     *filters.Cursor
	Filters    *filters.LocalFilter
	Sort       []filters.Sort
	Tenant     string
	Additional additional.Properties
}

type QueryParams struct {
	Class      string
	Offset     *int64
	Limit      *int64
	After      *string
	Sort       *string
	Order      *string
	Tenant     *string
	Additional additional.Properties
}

func (q *QueryParams) inputs(m *Manager) (*QueryInput, error) {
	smartOffset, smartLimit, err := m.localOffsetLimit(q.Offset, q.Limit)
	if err != nil {
		return nil, err
	}
	sort := m.getSort(q.Sort, q.Order)
	cursor := m.getCursor(q.After, q.Limit)
	tenant := ""
	if q.Tenant != nil {
		tenant = *q.Tenant
	}
	return &QueryInput{
		Class:      q.Class,
		Offset:     smartOffset,
		Limit:      smartLimit,
		Sort:       sort,
		Cursor:     cursor,
		Tenant:     tenant,
		Additional: q.Additional,
	}, nil
}

func (m *Manager) Query(ctx context.Context, principal *models.Principal, params *QueryParams,
) ([]*models.Object, *Error) {
	class := "*"
	aliasName := ""

	if params != nil && params.Class != "" {
		params.Class, aliasName = m.resolveAlias(params.Class)
		class = params.Class
	}

	if err := m.authorizer.Authorize(ctx, principal, authorization.READ, authorization.CollectionsData(class)...); err != nil {
		return nil, &Error{err.Error(), StatusForbidden, err}
	}

	m.metrics.GetObjectInc()
	defer m.metrics.GetObjectDec()

	q, err := params.inputs(m)
	if err != nil {
		return nil, &Error{"offset or limit", StatusBadRequest, err}
	}

	filteredQuery := filter.New[*QueryInput](m.authorizer, m.config.Config.Authorization.Rbac).Filter(
		ctx,
		m.logger,
		principal,
		[]*QueryInput{q},
		authorization.READ,
		func(qi *QueryInput) string {
			return authorization.CollectionsData(qi.Class)[0]
		},
	)
	if len(filteredQuery) == 0 {
		err = fmt.Errorf("unauthorized to access collection %s", q.Class)
		return nil, &Error{err.Error(), StatusForbidden, err}
	}

	res, rerr := m.vectorRepo.Query(ctx, filteredQuery[0])
	if rerr != nil {
		return nil, rerr
	}

	if m.modulesProvider != nil {
		res, err = m.modulesProvider.ListObjectsAdditionalExtend(ctx, res, q.Additional.ModuleParams)
		if err != nil {
			return nil, &Error{"extend results", StatusInternalServerError, err}
		}
	}

	if q.Additional.Vector {
		m.trackUsageList(res)
	}

	if aliasName != "" {
		return m.classNamesToAliases(res.ObjectsWithVector(q.Additional.Vector), aliasName), nil
	}
	return res.ObjectsWithVector(q.Additional.Vector), nil
}
