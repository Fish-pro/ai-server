package files

import (
	"database/sql"
	"gorm.io/gorm"

	"github.com/Fish-pro/ai-server/pkg/models"
)

type Files struct {
	models.CommonModel
	Name sql.NullString `json:"name" gorm:"column:name;type:varchar(32);comment:文件名称"`
	Url  sql.NullString `json:"URL" gorm:"column:URL;type:varchar(32);comment:文件链接"`
	Size sql.NullInt32  `json:"size" gorm:"column:del_flag;type:int;size:32;comment:文件大小"`
}

func (*Files) TableName() string {
	return "tb_ai_files"
}

type FileDao interface {
	Create(*Files) error
	Get(string) (*Files, error)
	Update(*Files) error
	Delete(string) error
	List() ([]*Files, error)
}

type FileDaoImpl struct {
	db *gorm.DB
}

func NewFileDao() FileDao {
	return &FileDaoImpl{}
}

func (f *FileDaoImpl) Create(*Files) error {
	return nil
}

func (f *FileDaoImpl) Get(name string) (*Files, error) {
	return nil, nil
}

func (f *FileDaoImpl) Update(*Files) error {
	return nil
}

func (f *FileDaoImpl) Delete(string) error {
	return nil
}

func (f *FileDaoImpl) List() ([]*Files, error) {
	return nil, nil
}
