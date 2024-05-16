package repository

import (
	"service/internal/domain/info/do"
	"service/internal/infrastructure/database"
)

type InfoRepository interface {
	Insert(do do.InfoDo) int64
	BatchInsert(dos []do.InfoDo) int64
	DeleteByPrimaryKey(id int64) int64
	DeleteByParam(param do.InfoDoParam) int64
	UpdateByPrimaryKey(do do.InfoDo) int64
	UpdateByParam(do do.InfoDo, param do.InfoDoParam) int64
	CountByParam(param do.InfoDoParam) int64
	SelectByPrimaryKey(id int64) do.InfoDo
	SelectByParam(param do.InfoDoParam) []do.InfoDo
	//PageableSelectByParam(param *do.InfoDoParam, pageIndex, pageSize int) []do.InfoDo
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type InfoRepositoryImpl struct {
}

func (r *InfoRepositoryImpl) Insert(do do.InfoDo) int64 {
	return database.Db.Create(&do).RowsAffected
}

func (r *InfoRepositoryImpl) BatchInsert(dos []do.InfoDo) int64 {
	return database.Db.Create(&dos).RowsAffected
}

func (r *InfoRepositoryImpl) DeleteByPrimaryKey(id int64) int64 {
	return database.Db.Delete(&do.InfoDo{}, id).RowsAffected
}

func (r *InfoRepositoryImpl) DeleteByParam(param do.InfoDoParam) int64 {
	return 0
}

func (r *InfoRepositoryImpl) UpdateByPrimaryKey(do do.InfoDo) int64 {
	return database.Db.Save(&do).RowsAffected
}

func (r *InfoRepositoryImpl) UpdateByParam(do do.InfoDo, param do.InfoDoParam) int64 {
	return 0
}

func (r *InfoRepositoryImpl) CountByParam(param do.InfoDoParam) int64 {
	var count int64
	db := database.Db
	for k, v := range param.Criteria.Condition {
		db = db.Where(k, v)
	}
	db.Count(&count)
	return count
}

func (r *InfoRepositoryImpl) SelectByPrimaryKey(id int64) do.InfoDo {
	do_ := do.InfoDo{}
	database.Db.Where("id = ?", id).First(&do_)
	return do_
}

func (r *InfoRepositoryImpl) SelectByParam(param do.InfoDoParam) []do.InfoDo {
	var dos []do.InfoDo
	db := database.Db
	for k, v := range param.Criteria.Condition {
		db = db.Where(k, v)
	}
	db.Find(&dos)
	return dos
}
