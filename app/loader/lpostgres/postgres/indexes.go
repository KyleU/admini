package postgres

import (
	"context"

	model2 "github.com/kyleu/admini/app/schema/model"
	"github.com/pkg/errors"

	"github.com/kyleu/admini/app/database"
	"github.com/kyleu/admini/app/util"
	"github.com/kyleu/admini/queries/qpostgres"
)

type indexResult struct {
	Schema      string `db:"schema_name"`
	Table       string `db:"table_name"`
	Index       string `db:"index_name"`
	PrimaryKey  bool   `db:"pk"`
	Unique      bool   `db:"u"`
	ColumnNames string `db:"column_names"`
}

func (r indexResult) AsIndex() *model2.Index {
	return &model2.Index{
		Key:      r.Index,
		Fields:   util.StringSplitAndTrim(r.ColumnNames, ","),
		Unique:   r.Unique,
		Primary:  r.PrimaryKey,
		Metadata: nil,
	}
}

func loadIndexes(ctx context.Context, models model2.Models, db *database.Service) error {
	var idxs []*indexResult
	err := db.Select(ctx, &idxs, qpostgres.ListIndexes(db.SchemaName), nil)
	if err != nil {
		return errors.Wrap(err, "can't list indexes")
	}

	for _, idx := range idxs {
		mod := models.Get(util.Pkg{idx.Schema}, idx.Table)
		if mod == nil {
			return errors.Errorf("no table [%s] found among [%d] candidates", idx.Table, len(models))
		}
		err = mod.AddIndex(idx.AsIndex())
		if err != nil {
			return errors.Wrap(err, "can't add index")
		}
	}
	return nil
}
