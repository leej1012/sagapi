package tables

import (
	"github.com/ontio/sagapi/config"
	"github.com/ontio/sagapi/models"
	"time"
)

type BuyRecord struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	OntId     string
	UserName  string
	TxHash    string
	Price     string
	ApiId     string
	ApiKey    string
	TxStatus  config.TxStatus
}

type ApiTestRecord struct {
	ID         uint `gorm:"primary_key"`
	OntId      string
	UserName   string
	ApiId      int
	TestResult int //0 test failed, 1 test success
}

type ApiBasicInfo struct {
	ApiId          uint `gorm:"primary_key"`
	ApiLogo        string
	ApiName        string
	ApiProvider    string
	ApiUrl         string
	ApiPrice       string
	ApiDesc        string
	Specifications int
	ApiExtra       models.ApiExtra
}

type ApiDetailInfo struct {
	ID                   uint `gorm:"primary_key"`
	Mark                 string
	RequestParam         string
	ResponseParam        string
	ResponseExample      string
	ParamErrorCode       string
	APIDetailInstruction models.ApiDetailInstruction
}

type APIKey struct {
	ID       uint `gorm:"primary_key"`
	ApiKey   string
	ApiId    int
	Limit    int
	UsedNum  int
	OntId    string
	UserName string
}
