package repo

import "github.com/372572571/xq/database"

type IRepo interface {
	TableName() string                   // model name
	Empty() interface{}                  // model{}
	EmptyList() interface{}            // []model
	Primary(key interface{}) interface{} // {ID:id}
}

type BaseRepo struct {
	db   *database.Database
	repo IRepo
}

func NewBaseRepo(ins IRepo, db *database.Database) *BaseRepo {
	return &BaseRepo{db: db, repo: ins}
}

// create model.
func (r *BaseRepo) Create(order interface{}) error {
	if err := r.db.Create(order).Error(); err != nil {
		return err
	}
	return nil
}

// primary key get model
func (r *BaseRepo) Get(id int64) interface{} {
	order := r.repo.Primary(id)

	if err := r.db.First(&order).Error(); err != nil {
		return r.repo.Empty()
	}

	return order
}

func (r *BaseRepo) First(data interface{}, w ...func(*database.Database) *database.Database) error {
	imp := r.db.Table(r.repo)

	for _, v := range w {
		imp = v(imp)
	}

	if err := imp.First(&data).Error(); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepo) Update(data interface{}, where ...func(*database.Database) *database.Database) int64 {
	imp := r.db.Table(r.repo)

	for _, v := range where {
		imp = v(imp)
	}

	db := imp.Updates(&data)
	err := db.Error()

	if err != nil {
		return 0
	}
	return db.RowsAffected()
}

func (r *BaseRepo) Find(m map[string]interface{}) interface{} {
	list := r.repo.EmptyList()
	r.db.Table(r.repo).Parse(m).Scan(&list)
	return list
}

func (r *BaseRepo) Count(m map[string]interface{}, count *int64) error {
	return r.db.Table(r.repo).Parse(m).Count(count)
}

func (r *BaseRepo) Delete(data interface{}, where ...func(*database.Database) *database.Database) int64 {
	imp := r.db.Table(r.repo)

	for _, v := range where {
		imp = v(imp)
	}

	db := imp.Delete(data)
	err := db.Error()

	if err != nil {
		return 0
	}
	return db.RowsAffected()
}
