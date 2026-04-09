package id_storage

import (
	"context"
	"time"

	"github.com/team-ide/framework/db"
	"github.com/team-ide/modules/module_id"
)

type Id struct {
	Type    string `json:"type" column:"id_type"`
	Value   int64  `json:"value" column:"id_value"`
	Comment string `json:"comment" column:"account"`

	CreateAt int64 `json:"createAt" column:"create_at"`
	UpdateAt int64 `json:"updateAt" column:"update_at"`
}

func (this_ *Id) GetTableName() string {
	return TableNameTbId
}
func (this_ *Id) GetPrimaryKey() []string {
	return []string{"id_type"}
}

type IdStorage struct {
	DbService db.IService
}

func (this_ *IdStorage) StorageQueryId(ctx context.Context, idType module_id.IdType) (int64, error) {
	m := this_.DbService.ModelSelect(&Id{Type: string(idType)})
	find, err := db.DoQueryOneWithModel[*Id](ctx, m)
	if err != nil {
		return 0, err
	}
	if find != nil {
		return find.Value, nil
	}
	return 0, err
}
func (this_ *IdStorage) StorageSave(ctx context.Context, idType module_id.IdType, id int64) error {
	m := this_.DbService.ModelSelect(&Id{Type: string(idType)})
	find, err := db.DoQueryOneWithModel[*Id](ctx, m)
	if err != nil {
		return err
	}
	if find == nil {
		in := &Id{Type: string(idType), Value: id}
		in.CreateAt = time.Now().UnixMilli()
		_, err = this_.DbService.Insert(ctx, in)
	} else {
		in := &Id{Type: string(idType), Value: id}
		in.UpdateAt = time.Now().UnixMilli()
		_, err = this_.DbService.Update(ctx, in)
	}
	if err != nil {
		return err
	}
	return nil
}

func (this_ *IdStorage) StorageInsert(ctx context.Context, idType module_id.IdType, id int64) error {
	in := &Id{Type: string(idType), Value: id}
	in.CreateAt = time.Now().UnixMilli()
	_, err := this_.DbService.Insert(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
func (this_ *IdStorage) StorageUpdate(ctx context.Context, idType module_id.IdType, id int64) error {
	in := &Id{Type: string(idType), Value: id}
	in.UpdateAt = time.Now().UnixMilli()
	_, err := this_.DbService.Update(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
