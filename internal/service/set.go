package service

import (
	"github.com/yemingfeng/sdb/internal/collection"
	"github.com/yemingfeng/sdb/internal/engine"
	"github.com/yemingfeng/sdb/internal/pb"
	"math"
)

var setCollection = collection.NewCollection(pb.DataType_SET)

func SPush(key []byte, values [][]byte, batch engine.Batch) error {
	for _, value := range values {
		if err := setCollection.UpsertRow(&collection.Row{
			Key:   key,
			Id:    value,
			Value: value,
		}, batch); err != nil {
			return err
		}
	}
	return nil
}

func SPop(key []byte, values [][]byte, batch engine.Batch) error {
	for _, value := range values {
		if err := setCollection.DelRowById(key, value, batch); err != nil {
			return err
		}
	}
	return nil
}

func SExist(key []byte, values [][]byte) ([]bool, error) {
	res := make([]bool, len(values))
	for i, value := range values {
		exist, err := setCollection.ExistRowById(key, value)
		if err != nil {
			return nil, err
		}
		res[i] = exist
	}
	return res, nil
}

func SDel(key []byte, batch engine.Batch) error {
	return setCollection.Del(key, batch)
}

func SCount(key []byte) (uint32, error) {
	return setCollection.Count(key)
}

func SMembers(key []byte) ([][]byte, error) {
	rows, err := setCollection.Page(key, 0, math.MaxUint32)
	if err != nil {
		return nil, err
	}
	res := make([][]byte, len(rows))
	for i := range rows {
		res[i] = rows[i].Value
	}
	return res, nil
}
