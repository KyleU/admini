package ldb

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
)

func Get(ctx context.Context, db *database.Service, m *model.Model, ids []interface{}, logger *zap.SugaredLogger) (*result.Result, error) {
	q, err := modelGetByPKQuery(db.Type, m, logger)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(ctx, q, nil, ids...)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%s]", m.String())
	}
	defer func() { _ = rows.Close() }()

	var timing *result.Timing
	ret, err := ParseResultFields(m.Name(), 0, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.String())
	}

	return ret, nil
}

func modelGetByPKQuery(typ *database.DBType, m *model.Model, logger *zap.SugaredLogger) (string, error) {
	cols, tbl := forTable(typ, m)
	pk := m.GetPK(logger)
	if len(pk) == 0 {
		return "", errors.Errorf("no PK for model [%s]", m.String())
	}
	where := make([]string, 0, len(pk))
	for idx, pkf := range pk {
		if typ.Placeholder == "?" {
			where = append(where, fmt.Sprintf(`%s%s%s = ?`, typ.Quote, pkf, typ.Quote))
		} else {
			where = append(where, fmt.Sprintf(`%s%s%s = $%d`, typ.Quote, pkf, typ.Quote, idx+1))
		}
	}
	return database.SQLSelectSimple(cols, tbl, strings.Join(where, " and ")), nil
}
