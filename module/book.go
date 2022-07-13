package module

type DicTag struct {
	Id         uint   `db:"id"`
	Tag        string `db:"tag"`
	Status     uint   `db:"status"` // 0-添加 1-删除
	CreateTime string `db:"create_time"`
}

func (d DicTag) TableName() string {
	return "dic_tag"
}

type File struct {
	Id         uint   `db:"id"`
	Order      uint   `db:"order"`
	No         string `db:"No"` // 文件编号
	Name       string `db:"name"`
	Ftype      string `db:"ftype"`
	Content    string `db:"content"`
	Author     string `db:"author"`
	Remark     string `db:"remark"`
	CustomerId uint   `db:"customer_id"` // 上传文件的用户
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
}

func (f File) TableName() string {
	return "file"
}

type MapFileTag struct {
	Id         uint   `db:"id"`
	TagId      uint   `db:"tag_id"`
	FileNo     uint   `db:"file_no"`
	Status     uint   `db:"status"` // 0-添加 1-删除
	CreateTime string `db:"create_time"`
}

func (m MapFileTag) TableName() string {
	return "map_file_tag"
}

type Note struct {
	Id         uint   `db:"id"`          // 笔记ID
	OriginC    string `db:"origin_c"`    // 源文件的内容
	Content    string `db:"content"`     // 笔记内容
	Status     uint   `db:"status"`      // 0-添加 1-删除
	FileNo     string `db:"file_no"`     // 笔记归属的文件的编号
	OrderId    uint   `db:"order_id"`    // 笔记的顺序，返回数据的时候按照orderId排序，会变
	CreateTime string `db:"create_time"` // 笔记时间
	UpdateTime string `db:"update_time"` // 笔记修改时间
}

func (n Note) TableName() string {
	return "note"
}
