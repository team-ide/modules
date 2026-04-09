package storage

import (
	"context"
	"github.com/team-ide/framework/db"
	"time"
)

type TbTest1 struct {
	UserId  int64  `json:"userId" column:"user_id"`
	Name    string `json:"name" column:"name"`
	Account string `json:"account" column:"account"`

	CreateAt int64 `json:"createAt" column:"create_at"`
	UpdateAt int64 `json:"updateAt" column:"update_at"`
}

func (this_ *TbTest1) GetTableName() string {
	return "tb_test1"
}
func (this_ *TbTest1) GetPrimaryKey() []string {
	return []string{"user_id"}
}

func NewTbTest1Storage(dbService db.IService) (res *TbTest1Storage) {
	res = &TbTest1Storage{}
	res.dbService = dbService
	return
}

type TbTest1Storage struct {
	dbService db.IService
}

func (this_ *TbTest1Storage) GetById(ctx context.Context, userId int64) (res *TbTest1, err error) {
	m := this_.dbService.ModelSelect(&TbTest1{UserId: userId})
	res, err = db.DoQueryOneWithModel[*TbTest1](ctx, m)
	return
}

func (this_ *TbTest1Storage) GetByIds(ctx context.Context, userIds ...int64) (res []*TbTest1, err error) {
	m := this_.dbService.ModelSelect(&TbTest1{})
	m.Where().In("user_id", userIds)
	res, err = db.DoQueryListWithModel[*TbTest1](ctx, m)
	return
}

func (this_ *TbTest1Storage) Insert(ctx context.Context, in *TbTest1) (err error) {
	in.CreateAt = time.Now().UnixMilli()

	r, err := this_.dbService.Insert(ctx, in)
	if err != nil {
		return
	}
	if in.UserId == 0 {
		in.UserId, err = r.LastInsertId()
		if err != nil {
			return
		}
	}
	return
}

func (this_ *TbTest1Storage) Update(ctx context.Context, in *TbTest1) (err error) {
	in.UpdateAt = time.Now().UnixMilli()

	_, err = this_.dbService.Update(ctx, in)
	return
}

func (this_ *TbTest1Storage) Delete(ctx context.Context, in *TbTest1) (err error) {

	_, err = this_.dbService.Delete(ctx, in)
	return
}

func (this_ *TbTest1Storage) Query(ctx context.Context, in *TbTest1) (res []*TbTest1, err error) {

	m := this_.dbService.ModelSelect(in)
	m.Where().Ne("status", 9)

	res, err = db.DoQueryListWithModel[*TbTest1](ctx, m)
	return
}
