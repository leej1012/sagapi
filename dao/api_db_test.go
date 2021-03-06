package dao

import (
	"fmt"
	"github.com/ontio/sagapi/models/tables"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApiDB_InsertApiBasicInfo(t *testing.T) {
	info := &tables.ApiBasicInfo{
		Icon:            "",
		Title:           "",
		ApiProvider:     "",
		ApiUrl:          "",
		Price:           "",
		ApiDesc:         "",
		Specifications:  1,
		Popularity:      0,
		Delay:           0,
		SuccessRate:     0,
		InvokeFrequency: 0,
	}
	l := 11
	infos := make([]*tables.ApiBasicInfo, l)
	for i := 0; i < len(infos); i++ {
		infos[i] = info
	}
	assert.Nil(t, TestDB.ApiDB.InsertApiBasicInfo(infos))
}

func TestApiDB_UpdateApiKeyInvokeFre(t *testing.T) {
	err := TestDB.ApiDB.UpdateApiKeyInvokeFre("test_3dddcb11-041a-4184-ad69-7505542649a6", 80, 1, 89)
	assert.Nil(t, err)
}

func TestApiDB_QueryApiBasicInfoByApiId(t *testing.T) {
	info, err := TestDB.ApiDB.QueryApiBasicInfoByApiId(1)
	assert.Nil(t, err)
	assert.Equal(t, info.ApiId, 1)

	infos, err := TestDB.ApiDB.QueryApiBasicInfoByPage(1, 2)
	assert.Nil(t, err)
	assert.Equal(t, len(infos), 2)
	price, err := TestDB.ApiDB.QueryPriceByApiId(1)
	assert.Nil(t, err)
	assert.Equal(t, price, "")
}

func TestApiDB_InsertApiDetailInfo(t *testing.T) {
	info := &tables.ApiDetailInfo{
		ApiId:               1,
		Mark:                "",
		ResponseParam:       "",
		ResponseExample:     "",
		DataDesc:            "test",
		DataSource:          "",
		ApplicationScenario: "",
	}
	err := TestDB.ApiDB.InsertApiDetailInfo(info)
	assert.Nil(t, err)
}

func TestApiDB_QueryApiDetailInfoById(t *testing.T) {
	info, err := TestDB.ApiDB.QueryApiDetailInfoById(1)
	assert.Nil(t, err)
	assert.Equal(t, info.ApiId, 1)
}

func TestApiDB_InsertRequestParam(t *testing.T) {
	rp := &tables.RequestParam{
		ApiDetailInfoId: 1,
		ParamName:       "",
		Required:        true,
		ParamType:       "",
		ValueDesc:       "",
	}
	l := 10
	requestParam := make([]*tables.RequestParam, l)
	for i := 0; i < l; i++ {
		requestParam[i] = rp
	}
	err := TestDB.ApiDB.InsertRequestParam(requestParam)
	assert.Nil(t, err)

	param, err := TestDB.ApiDB.QueryRequestParamByApiDetailInfoId(1)
	assert.Nil(t, err)
	assert.Equal(t, len(param), 10)
}

func TestApiDB_InsertErrorCode(t *testing.T) {
	rp := &tables.ErrorCode{
		ApiDetailInfoId: 1,
		ErrorCode:       1,
		ErrorDesc:       "",
	}
	l := 10
	requestParam := make([]*tables.ErrorCode, l)
	for i := 0; i < l; i++ {
		requestParam[i] = rp
	}
	err := TestDB.ApiDB.InsertErrorCode(requestParam)
	assert.Nil(t, err)
}

func TestApiDB_QueryErrorCodeByApiDetailInfoId(t *testing.T) {
	param, err := TestDB.ApiDB.QueryErrorCodeByApiDetailInfoId(1)
	assert.Nil(t, err)
	assert.Equal(t, len(param), 10)
}

func TestApiDB_QueryNewestApiBasicInfo(t *testing.T) {
	infos, err := TestDB.ApiDB.QueryFreeApiBasicInfo()
	fmt.Println(err)
	assert.Nil(t, err)
	fmt.Println(infos)
}

func TestApiDB_testTag(t *testing.T) {
	a := &tables.ApiTag{
		Id:         20,
		CreateTime: time.Now(),
	}

	b := &tables.Tag{
		Id:         30,
		CreateTime: time.Now(),
	}

	c := &tables.Category{
		Id:     40,
		NameZh: "xxx",
		NameEn: "xxx",
	}

	err := TestDB.ApiDB.InsertApiTag(a)
	assert.Nil(t, err)
	err = TestDB.ApiDB.InsertTag(b)
	assert.Nil(t, err)
	err = TestDB.ApiDB.InsertCategory(c)
	assert.Nil(t, err)
}
