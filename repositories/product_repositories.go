package repositories

import (
	"ProductPro/models"
	"database/sql"
)

// 第一步，先开发对应接口
// 第二步，实现定义接口

type IProduct interface {
	// 连接数据
	Conn() error
	Insert(*models.Product)(int64, error)
	Delete(int64) bool
	Update(*models.Product) error
	SelectByKey(int64) (*models.Product, error)
	SelectAll()([]*models.Product, error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string, db *sql.DB) IProduct {
	return &ProductManager{table: table, mysqlConn: db}
}

func (p *ProductManager) Conn() error {
	if p.mysqlConn == nil{
		mysql, err :=
	}
}