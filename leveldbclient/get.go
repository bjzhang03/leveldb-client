package leveldbclient

import (
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/syndtr/goleveldb/leveldb"
)

func Get(key string, path string) string {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Error("Get the db failed! ", err)
	}
	defer db.Close()

	result, err := db.Get([]byte(key), nil)

	if err != nil {
		log.Errorf("Got the error, when get the [%s] ", key)
	}
	return string(result)
}

func GetAll(path string) []string {
	db, err := leveldb.OpenFile(path, nil)
	//数据库遍历
	iter := db.NewIterator(nil, nil)
	result := []string{}
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		result = append(result, fmt.Sprintf("%s,%s", string(key), string(value)))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		log.Errorf("Release the iterator failed!")
	}
	return result
}
