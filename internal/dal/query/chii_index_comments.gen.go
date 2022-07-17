// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newIndexComment(db *gorm.DB) indexComment {
	_indexComment := indexComment{}

	_indexComment.indexCommentDo.UseDB(db)
	_indexComment.indexCommentDo.UseModel(&dao.IndexComment{})

	tableName := _indexComment.indexCommentDo.TableName()
	_indexComment.ALL = field.NewField(tableName, "*")
	_indexComment.ID = field.NewUint32(tableName, "idx_pst_id")
	_indexComment.TopicID = field.NewUint32(tableName, "idx_pst_mid")
	_indexComment.UID = field.NewUint32(tableName, "idx_pst_uid")
	_indexComment.Related = field.NewUint32(tableName, "idx_pst_related")
	_indexComment.CreatedTime = field.NewUint32(tableName, "idx_pst_dateline")
	_indexComment.Content = field.NewString(tableName, "idx_pst_content")

	_indexComment.fillFieldMap()

	return _indexComment
}

type indexComment struct {
	indexCommentDo indexCommentDo

	ALL         field.Field
	ID          field.Uint32
	TopicID     field.Uint32
	UID         field.Uint32
	Related     field.Uint32
	CreatedTime field.Uint32
	Content     field.String

	fieldMap map[string]field.Expr
}

func (i indexComment) Table(newTableName string) *indexComment {
	i.indexCommentDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i indexComment) As(alias string) *indexComment {
	i.indexCommentDo.DO = *(i.indexCommentDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *indexComment) updateTableName(table string) *indexComment {
	i.ALL = field.NewField(table, "*")
	i.ID = field.NewUint32(table, "idx_pst_id")
	i.TopicID = field.NewUint32(table, "idx_pst_mid")
	i.UID = field.NewUint32(table, "idx_pst_uid")
	i.Related = field.NewUint32(table, "idx_pst_related")
	i.CreatedTime = field.NewUint32(table, "idx_pst_dateline")
	i.Content = field.NewString(table, "idx_pst_content")

	i.fillFieldMap()

	return i
}

func (i *indexComment) WithContext(ctx context.Context) *indexCommentDo {
	return i.indexCommentDo.WithContext(ctx)
}

func (i indexComment) TableName() string { return i.indexCommentDo.TableName() }

func (i indexComment) Alias() string { return i.indexCommentDo.Alias() }

func (i *indexComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *indexComment) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 6)
	i.fieldMap["idx_pst_id"] = i.ID
	i.fieldMap["idx_pst_mid"] = i.TopicID
	i.fieldMap["idx_pst_uid"] = i.UID
	i.fieldMap["idx_pst_related"] = i.Related
	i.fieldMap["idx_pst_dateline"] = i.CreatedTime
	i.fieldMap["idx_pst_content"] = i.Content
}

func (i indexComment) clone(db *gorm.DB) indexComment {
	i.indexCommentDo.ReplaceDB(db)
	return i
}

type indexCommentDo struct{ gen.DO }

func (i indexCommentDo) Debug() *indexCommentDo {
	return i.withDO(i.DO.Debug())
}

func (i indexCommentDo) WithContext(ctx context.Context) *indexCommentDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i indexCommentDo) ReadDB() *indexCommentDo {
	return i.Clauses(dbresolver.Read)
}

func (i indexCommentDo) WriteDB() *indexCommentDo {
	return i.Clauses(dbresolver.Write)
}

func (i indexCommentDo) Clauses(conds ...clause.Expression) *indexCommentDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i indexCommentDo) Returning(value interface{}, columns ...string) *indexCommentDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i indexCommentDo) Not(conds ...gen.Condition) *indexCommentDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i indexCommentDo) Or(conds ...gen.Condition) *indexCommentDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i indexCommentDo) Select(conds ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i indexCommentDo) Where(conds ...gen.Condition) *indexCommentDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i indexCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *indexCommentDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i indexCommentDo) Order(conds ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i indexCommentDo) Distinct(cols ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i indexCommentDo) Omit(cols ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i indexCommentDo) Join(table schema.Tabler, on ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i indexCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i indexCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i indexCommentDo) Group(cols ...field.Expr) *indexCommentDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i indexCommentDo) Having(conds ...gen.Condition) *indexCommentDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i indexCommentDo) Limit(limit int) *indexCommentDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i indexCommentDo) Offset(offset int) *indexCommentDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i indexCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *indexCommentDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i indexCommentDo) Unscoped() *indexCommentDo {
	return i.withDO(i.DO.Unscoped())
}

func (i indexCommentDo) Create(values ...*dao.IndexComment) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i indexCommentDo) CreateInBatches(values []*dao.IndexComment, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i indexCommentDo) Save(values ...*dao.IndexComment) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i indexCommentDo) First() (*dao.IndexComment, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexComment), nil
	}
}

func (i indexCommentDo) Take() (*dao.IndexComment, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexComment), nil
	}
}

func (i indexCommentDo) Last() (*dao.IndexComment, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexComment), nil
	}
}

func (i indexCommentDo) Find() ([]*dao.IndexComment, error) {
	result, err := i.DO.Find()
	return result.([]*dao.IndexComment), err
}

func (i indexCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.IndexComment, err error) {
	buf := make([]*dao.IndexComment, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i indexCommentDo) FindInBatches(result *[]*dao.IndexComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i indexCommentDo) Attrs(attrs ...field.AssignExpr) *indexCommentDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i indexCommentDo) Assign(attrs ...field.AssignExpr) *indexCommentDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i indexCommentDo) Joins(fields ...field.RelationField) *indexCommentDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i indexCommentDo) Preload(fields ...field.RelationField) *indexCommentDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i indexCommentDo) FirstOrInit() (*dao.IndexComment, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexComment), nil
	}
}

func (i indexCommentDo) FirstOrCreate() (*dao.IndexComment, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.IndexComment), nil
	}
}

func (i indexCommentDo) FindByPage(offset int, limit int) (result []*dao.IndexComment, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i indexCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i indexCommentDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i *indexCommentDo) withDO(do gen.Dao) *indexCommentDo {
	i.DO = *do.(*gen.DO)
	return i
}
