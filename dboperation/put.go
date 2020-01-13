package dboperation

import (
	"github.com/prometheus/common/log"
	"github.com/syndtr/goleveldb/leveldb"
)

func Put(key, value, path string) error {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Error("Get the db failed!", err)
	}
	defer db.Close()

	err = db.Put([]byte(key), []byte(value), nil)

	if err != nil {
		log.Errorf("Got the error, when put [%s, %s]", key, value)
		return err
	}

	log.Infof("Put the [%s, %s] in db success!", key, value)
	return nil
}
