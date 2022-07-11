package module

type DicTag struct {
	Id         uint   `db:"id"`
	Tag        string `db:"tag"`
	Status     uint   `db:"status"` // 0-添加 1-删除
	CreateTime string `db:"create_time"`
}

type File struct {
	Id         uint   `db:"id"`
	Order      uint   `db:"order"`
	Name       string `db:"name"`
	Ftype      string `db:"ftype"`
	Content    string `db:"content"`
	Author     string `db:"author"`
	Remark     string `db:"remark"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

type FileTag struct {
	Id         uint   `db:"id"`
	TagId      uint   `db:"tag_id"`
	FileId     uint   `db:"file_id"`
	Status     uint   `db:"status"` // 0-添加 1-删除
	CreateTime string `db:"create_time"`
}

type Note struct {
	Id         uint   `db:"id"`          // 笔记ID
	Content    string `db:"content"`     // 笔记内容
	Status     uint   `db:"status"`      // 0-添加 1-删除
	FileId     uint   `db:"file_id"`     // 笔记归属的文件ID
	CreateTime string `db:"create_time"` // 笔记时间
	UpdateTime string `db:"update_time"` // 笔记修改时间
}
