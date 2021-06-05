package lpostgres

import (
	"fmt"
	"strings"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/model"
	"github.com/kyleu/admini/app/result"
)

func (l *Loader) Get(m *model.Model, ids []interface{}) (*result.Result, error) {
	q, err := modelGetByPKQuery(m, l.logger)
	if err != nil {
		return nil, err
	}
	rows, err := l.db.Query(q, nil, ids...)
	if err != nil {
		return nil, errors.Wrapf(err, "error listing models for [%s]", m.Key)
	}

	var timing *result.Timing
	ret, err := ParseResultFields(m.Name(), 0, q, timing, m.Fields, rows)
	if err != nil {
		return nil, errors.Wrapf(err, "error constructing result for [%s]", m.Key)
	}

	return ret, nil
}

func modelGetByPKQuery(m *model.Model, logger *zap.SugaredLogger) (string, error) {
	cols, tbl := forTable(m)
	pk := m.GetPK(logger)
	if len(pk) == 0 {
		return "", errors.Errorf("no PK for model [%s]", m.Key)
	}
	where := make([]string, 0, len(pk))
	for idx, pkf := range pk {
		where = append(where, fmt.Sprintf(`"%s" = $%d`, pkf, idx+1))
	}
	return database.SQLSelectSimple(cols, tbl, strings.Join(where, " and ")), nil
}
