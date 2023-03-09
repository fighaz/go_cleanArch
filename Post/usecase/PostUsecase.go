package usecase

import (
	"blog/model"
)

type PostUsecase struct {
	PostRepo model.PostRepository
}

func NewPostUsecase(PostRepo model.PostRepository) model.PostUsecase {
	return &PostUsecase{PostRepo: PostRepo}
}
func (pu *PostUsecase) GetAll() (res []model.Post, err error) {
	res, err = pu.PostRepo.GetAll()
	return res, nil
}
func (pu *PostUsecase) GetByID(id int) (res model.Post, err error) {

	res, err = pu.PostRepo.GetByID(id)

	return res, nil
}
func (pu *PostUsecase) Insert(p *model.Post) (res model.Post, err error) {
	res, err = pu.PostRepo.Insert(p)

	return res, err

}
func (pu *PostUsecase) Update(p *model.Post) (res model.Post, err error) {
	res, err = pu.PostRepo.Update(p)

	return res, nil

}
func (pu *PostUsecase) Delete(id int) (err error) {

	err = pu.PostRepo.Delete(id)

	return err
}
