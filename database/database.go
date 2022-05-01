package database

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Database struct {
	impl *gorm.DB
}

func (db *Database) Debug() *Database {
	return &Database{impl: db.impl.Debug()}
}

func (db *Database) Unsafe() *gorm.DB {
	return db.Debug().impl
}

func (db *Database) Error() error {
	return db.impl.Error
}

func (db *Database) RowsAffected() int64 {
	return db.impl.RowsAffected
}

func (db *Database) Updates(values interface{}) *Database {
	tx := db.impl.Updates(values)
	return &Database{impl: tx}
}

func (db *Database) Where(query interface{}, args ...interface{}) *Database {
	return &Database{impl: db.impl.Where(query, args...)}
}

func (db *Database) Select(query interface{}, args ...interface{}) *Database {
	return &Database{impl: db.impl.Select(query, args...)}
}

func (db *Database) Scan(dest interface{}) *Database {
	return &Database{impl: db.impl.Scan(dest)}
}

func (db *Database) Find(dest interface{}) *Database {
	return &Database{impl: db.impl.Find(dest)}
}

func (db *Database) Count(i *int64) error {
	return db.impl.Count(i).Error
}

func (db *Database) First(dest interface{}, conds ...interface{}) *Database {
	return &Database{impl: db.impl.First(dest, conds...)}
}

func (db *Database) Create(value interface{}) *Database {
	return &Database{impl: db.impl.Create(value)}
}

func (db *Database) Delete(value interface{}, conds ...interface{}) *Database {
	return &Database{impl: db.impl.Delete(value, conds...)}
}

func (db *Database) Exec(sql string, values ...interface{}) *Database {
	return &Database{impl: db.impl.Exec(sql, values...)}
}

func (db *Database) Unscoped() *Database {
	return &Database{impl: db.impl.Unscoped()}
}

func (db *Database) Raw(sql string, values ...interface{}) *Database {
	return &Database{impl: db.impl.Raw(sql, values...)}
}

func (db *Database) Packsafe(g *gorm.DB) *Database {

	return &Database{impl: g}
}

func (db *Database) Table(i interface{ TableName() string }) *Database {
	return &Database{impl: db.impl.Table(i.TableName())}
}

func (db *Database) Save(value interface{}) *Database {

	return &Database{impl: db.impl.Save(value)}
}

func (db *Database) Parse(data map[string]interface{}) *Database {
	query := db.impl

	// where 语句处理
	if where, ok := data["where"]; ok {
		for k, v := range where.(map[string]interface{}) {
			if !strings.Contains(k, "?") {
				k = fmt.Sprintf("%s = ?", k)
			}
			query = query.Where(k, v)
		}
	}

	// 排序
	if order, ok := data["order"]; ok {
		for _, v := range order.([]interface{}) {
			query = query.Order(v.(string))
		}
	}

	// 偏移
	if offset, ok := data["offset"]; ok {
		query = query.Offset(int(offset.(float64)))
	}

	// 条数限制
	if limit, ok := data["limit"]; ok {
		query = query.Limit(int(limit.(float64)))
	}

	return &Database{impl: query}
}

func (db *Database) Transaction(transaction func(tx *Database) error) error {
	return db.impl.Transaction(func(tx *gorm.DB) error {
		return transaction(&Database{impl: tx})
	})
}
