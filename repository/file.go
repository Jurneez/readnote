package repository

import (
	"errors"
	"fmt"
	"tool/readnote/common"
	"tool/readnote/log"
	"tool/readnote/module"
)

// search file infomation by file-no
func SearchFileByNo(fileNo string) (*module.File, error) {
	if fileNo == "" {
		return nil, errors.New(fmt.Sprintf(common.ErrParamEmpty, "fileNo"))
	}

	sql := fmt.Sprintf(`
		select id,no from file where no = '%s'
	`, fileNo)
	log.Write("log.txt", fmt.Sprintf("[%s]  %s", "SearchFileByNo", sql))

	rows, err := common.Test_DB.Query(sql)
	if err != nil {
		return nil, err
	}
	var fileData module.File
	for rows.Next() {
		if err := rows.Scan(&fileData.Id, &fileData.No); err != nil {
			return nil, err
		}
	}
	return &fileData, nil
}
