// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"go-skeleton/internal/app/model"
)

func newYfoAdmin(db *gorm.DB, opts ...gen.DOOption) yfoAdmin {
	_yfoAdmin := yfoAdmin{}

	_yfoAdmin.yfoAdminDo.UseDB(db, opts...)
	_yfoAdmin.yfoAdminDo.UseModel(&model.YfoAdmin{})

	tableName := _yfoAdmin.yfoAdminDo.TableName()
	_yfoAdmin.ALL = field.NewAsterisk(tableName)
	_yfoAdmin.ID = field.NewInt64(tableName, "id")
	_yfoAdmin.UUID = field.NewString(tableName, "uuid")
	_yfoAdmin.Account = field.NewString(tableName, "account")
	_yfoAdmin.Password = field.NewString(tableName, "password")
	_yfoAdmin.Phone = field.NewString(tableName, "phone")
	_yfoAdmin.Avatar = field.NewString(tableName, "avatar")
	_yfoAdmin.Salt = field.NewString(tableName, "salt")
	_yfoAdmin.RealName = field.NewString(tableName, "real_name")
	_yfoAdmin.RegisterTime = field.NewInt64(tableName, "register_time")
	_yfoAdmin.RegisterIP = field.NewString(tableName, "register_ip")
	_yfoAdmin.LoginTime = field.NewInt64(tableName, "login_time")
	_yfoAdmin.LoginIP = field.NewString(tableName, "login_ip")
	_yfoAdmin.RoleIds = field.NewString(tableName, "role_ids")
	_yfoAdmin.Status = field.NewInt8(tableName, "status")
	_yfoAdmin.CreatedAt = field.NewInt64(tableName, "created_at")
	_yfoAdmin.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_yfoAdmin.fillFieldMap()

	return _yfoAdmin
}

// yfoAdmin 后台管理员表
type yfoAdmin struct {
	yfoAdminDo

	ALL          field.Asterisk
	ID           field.Int64
	UUID         field.String // 唯一id号
	Account      field.String // 账号
	Password     field.String // 密码
	Phone        field.String // 手机号
	Avatar       field.String // 头像
	Salt         field.String // 密码
	RealName     field.String // 真实姓名
	RegisterTime field.Int64  // 注册时间
	RegisterIP   field.String // 注册ip
	LoginTime    field.Int64  // 登录时间
	LoginIP      field.String // 登录ip
	RoleIds      field.String // 角色IDs
	Status       field.Int8   // 状态 1：正常 2：禁用
	CreatedAt    field.Int64
	UpdatedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (y yfoAdmin) Table(newTableName string) *yfoAdmin {
	y.yfoAdminDo.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y yfoAdmin) As(alias string) *yfoAdmin {
	y.yfoAdminDo.DO = *(y.yfoAdminDo.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *yfoAdmin) updateTableName(table string) *yfoAdmin {
	y.ALL = field.NewAsterisk(table)
	y.ID = field.NewInt64(table, "id")
	y.UUID = field.NewString(table, "uuid")
	y.Account = field.NewString(table, "account")
	y.Password = field.NewString(table, "password")
	y.Phone = field.NewString(table, "phone")
	y.Avatar = field.NewString(table, "avatar")
	y.Salt = field.NewString(table, "salt")
	y.RealName = field.NewString(table, "real_name")
	y.RegisterTime = field.NewInt64(table, "register_time")
	y.RegisterIP = field.NewString(table, "register_ip")
	y.LoginTime = field.NewInt64(table, "login_time")
	y.LoginIP = field.NewString(table, "login_ip")
	y.RoleIds = field.NewString(table, "role_ids")
	y.Status = field.NewInt8(table, "status")
	y.CreatedAt = field.NewInt64(table, "created_at")
	y.UpdatedAt = field.NewInt64(table, "updated_at")

	y.fillFieldMap()

	return y
}

func (y *yfoAdmin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *yfoAdmin) fillFieldMap() {
	y.fieldMap = make(map[string]field.Expr, 16)
	y.fieldMap["id"] = y.ID
	y.fieldMap["uuid"] = y.UUID
	y.fieldMap["account"] = y.Account
	y.fieldMap["password"] = y.Password
	y.fieldMap["phone"] = y.Phone
	y.fieldMap["avatar"] = y.Avatar
	y.fieldMap["salt"] = y.Salt
	y.fieldMap["real_name"] = y.RealName
	y.fieldMap["register_time"] = y.RegisterTime
	y.fieldMap["register_ip"] = y.RegisterIP
	y.fieldMap["login_time"] = y.LoginTime
	y.fieldMap["login_ip"] = y.LoginIP
	y.fieldMap["role_ids"] = y.RoleIds
	y.fieldMap["status"] = y.Status
	y.fieldMap["created_at"] = y.CreatedAt
	y.fieldMap["updated_at"] = y.UpdatedAt
}

func (y yfoAdmin) clone(db *gorm.DB) yfoAdmin {
	y.yfoAdminDo.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y yfoAdmin) replaceDB(db *gorm.DB) yfoAdmin {
	y.yfoAdminDo.ReplaceDB(db)
	return y
}

type yfoAdminDo struct{ gen.DO }

type IYfoAdminDo interface {
	gen.SubQuery
	Debug() IYfoAdminDo
	WithContext(ctx context.Context) IYfoAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IYfoAdminDo
	WriteDB() IYfoAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IYfoAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IYfoAdminDo
	Not(conds ...gen.Condition) IYfoAdminDo
	Or(conds ...gen.Condition) IYfoAdminDo
	Select(conds ...field.Expr) IYfoAdminDo
	Where(conds ...gen.Condition) IYfoAdminDo
	Order(conds ...field.Expr) IYfoAdminDo
	Distinct(cols ...field.Expr) IYfoAdminDo
	Omit(cols ...field.Expr) IYfoAdminDo
	Join(table schema.Tabler, on ...field.Expr) IYfoAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IYfoAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) IYfoAdminDo
	Group(cols ...field.Expr) IYfoAdminDo
	Having(conds ...gen.Condition) IYfoAdminDo
	Limit(limit int) IYfoAdminDo
	Offset(offset int) IYfoAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IYfoAdminDo
	Unscoped() IYfoAdminDo
	Create(values ...*model.YfoAdmin) error
	CreateInBatches(values []*model.YfoAdmin, batchSize int) error
	Save(values ...*model.YfoAdmin) error
	First() (*model.YfoAdmin, error)
	Take() (*model.YfoAdmin, error)
	Last() (*model.YfoAdmin, error)
	Find() ([]*model.YfoAdmin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YfoAdmin, err error)
	FindInBatches(result *[]*model.YfoAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.YfoAdmin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IYfoAdminDo
	Assign(attrs ...field.AssignExpr) IYfoAdminDo
	Joins(fields ...field.RelationField) IYfoAdminDo
	Preload(fields ...field.RelationField) IYfoAdminDo
	FirstOrInit() (*model.YfoAdmin, error)
	FirstOrCreate() (*model.YfoAdmin, error)
	FindByPage(offset int, limit int) (result []*model.YfoAdmin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IYfoAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	GetByID(id int) (result model.YfoAdmin, err error)
	FindAll() (result []model.YfoAdmin, err error)
	FindOne() (result model.YfoAdmin)
	GetByAccount(account string) (result model.YfoAdmin, err error)
}

// SELECT * FROM @@table WHERE id = @id
func (y yfoAdminDo) GetByID(id int) (result model.YfoAdmin, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("SELECT * FROM yfo_admin WHERE id = ? ")

	var executeSQL *gorm.DB
	executeSQL = y.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table
func (y yfoAdminDo) FindAll() (result []model.YfoAdmin, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM yfo_admin ")

	var executeSQL *gorm.DB
	executeSQL = y.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// SELECT * FROM @@table LIMIT 1
func (y yfoAdminDo) FindOne() (result model.YfoAdmin) {
	var generateSQL strings.Builder
	generateSQL.WriteString("SELECT * FROM yfo_admin LIMIT 1 ")

	var executeSQL *gorm.DB
	executeSQL = y.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	_ = executeSQL

	return
}

// SELECT * FROM @@table WHERE account = @account
func (y yfoAdminDo) GetByAccount(account string) (result model.YfoAdmin, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, account)
	generateSQL.WriteString("SELECT * FROM yfo_admin WHERE account = ? ")

	var executeSQL *gorm.DB
	executeSQL = y.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (y yfoAdminDo) Debug() IYfoAdminDo {
	return y.withDO(y.DO.Debug())
}

func (y yfoAdminDo) WithContext(ctx context.Context) IYfoAdminDo {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y yfoAdminDo) ReadDB() IYfoAdminDo {
	return y.Clauses(dbresolver.Read)
}

func (y yfoAdminDo) WriteDB() IYfoAdminDo {
	return y.Clauses(dbresolver.Write)
}

func (y yfoAdminDo) Session(config *gorm.Session) IYfoAdminDo {
	return y.withDO(y.DO.Session(config))
}

func (y yfoAdminDo) Clauses(conds ...clause.Expression) IYfoAdminDo {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y yfoAdminDo) Returning(value interface{}, columns ...string) IYfoAdminDo {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y yfoAdminDo) Not(conds ...gen.Condition) IYfoAdminDo {
	return y.withDO(y.DO.Not(conds...))
}

func (y yfoAdminDo) Or(conds ...gen.Condition) IYfoAdminDo {
	return y.withDO(y.DO.Or(conds...))
}

func (y yfoAdminDo) Select(conds ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.Select(conds...))
}

func (y yfoAdminDo) Where(conds ...gen.Condition) IYfoAdminDo {
	return y.withDO(y.DO.Where(conds...))
}

func (y yfoAdminDo) Order(conds ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.Order(conds...))
}

func (y yfoAdminDo) Distinct(cols ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y yfoAdminDo) Omit(cols ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.Omit(cols...))
}

func (y yfoAdminDo) Join(table schema.Tabler, on ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.Join(table, on...))
}

func (y yfoAdminDo) LeftJoin(table schema.Tabler, on ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y yfoAdminDo) RightJoin(table schema.Tabler, on ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y yfoAdminDo) Group(cols ...field.Expr) IYfoAdminDo {
	return y.withDO(y.DO.Group(cols...))
}

func (y yfoAdminDo) Having(conds ...gen.Condition) IYfoAdminDo {
	return y.withDO(y.DO.Having(conds...))
}

func (y yfoAdminDo) Limit(limit int) IYfoAdminDo {
	return y.withDO(y.DO.Limit(limit))
}

func (y yfoAdminDo) Offset(offset int) IYfoAdminDo {
	return y.withDO(y.DO.Offset(offset))
}

func (y yfoAdminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IYfoAdminDo {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y yfoAdminDo) Unscoped() IYfoAdminDo {
	return y.withDO(y.DO.Unscoped())
}

func (y yfoAdminDo) Create(values ...*model.YfoAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y yfoAdminDo) CreateInBatches(values []*model.YfoAdmin, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y yfoAdminDo) Save(values ...*model.YfoAdmin) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y yfoAdminDo) First() (*model.YfoAdmin, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.YfoAdmin), nil
	}
}

func (y yfoAdminDo) Take() (*model.YfoAdmin, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.YfoAdmin), nil
	}
}

func (y yfoAdminDo) Last() (*model.YfoAdmin, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.YfoAdmin), nil
	}
}

func (y yfoAdminDo) Find() ([]*model.YfoAdmin, error) {
	result, err := y.DO.Find()
	return result.([]*model.YfoAdmin), err
}

func (y yfoAdminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YfoAdmin, err error) {
	buf := make([]*model.YfoAdmin, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y yfoAdminDo) FindInBatches(result *[]*model.YfoAdmin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y yfoAdminDo) Attrs(attrs ...field.AssignExpr) IYfoAdminDo {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y yfoAdminDo) Assign(attrs ...field.AssignExpr) IYfoAdminDo {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y yfoAdminDo) Joins(fields ...field.RelationField) IYfoAdminDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y yfoAdminDo) Preload(fields ...field.RelationField) IYfoAdminDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y yfoAdminDo) FirstOrInit() (*model.YfoAdmin, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.YfoAdmin), nil
	}
}

func (y yfoAdminDo) FirstOrCreate() (*model.YfoAdmin, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.YfoAdmin), nil
	}
}

func (y yfoAdminDo) FindByPage(offset int, limit int) (result []*model.YfoAdmin, count int64, err error) {
	result, err = y.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = y.Offset(-1).Limit(-1).Count()
	return
}

func (y yfoAdminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y yfoAdminDo) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y yfoAdminDo) Delete(models ...*model.YfoAdmin) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *yfoAdminDo) withDO(do gen.Dao) *yfoAdminDo {
	y.DO = *do.(*gen.DO)
	return y
}
