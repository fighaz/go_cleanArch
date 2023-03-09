package repository

import (
	"blog/model"

	"gorm.io/gorm"
)

type MysqlPostRepository struct {
	Conn *gorm.DB
}

func NewPostRepository(Conn *gorm.DB) model.PostRepository {
	return &MysqlPostRepository{Conn}
}
func (m *MysqlPostRepository) GetAll() (res []model.Post, err error) {
	result := m.Conn.Find(&res)
	return res, result.Error
}
func (m *MysqlPostRepository) GetByID(id int) (res model.Post, err error) {

	result := m.Conn.First(&res, id)

	return res, result.Error
}
func (m *MysqlPostRepository) Insert(p *model.Post) (res model.Post, err error) {
	result := m.Conn.Create(&p)

	return res, result.Error

}
func (m *MysqlPostRepository) Update(p *model.Post) (res model.Post, err error) {
	// result := m.Conn.First(&p, p.Id)
	result := m.Conn.Save(&p)

	return res, result.Error

}
func (m *MysqlPostRepository) Delete(id int) (err error) {

	result := m.Conn.Delete(&model.Post{Id: id})

	return result.Error
}
