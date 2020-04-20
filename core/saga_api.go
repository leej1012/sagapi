package core

import (
	"github.com/ontio/sagapi/common"
	"github.com/ontio/sagapi/core/nasa"
	"github.com/ontio/sagapi/dao"
	"github.com/ontio/sagapi/models/tables"
)

var DefSagaApi *SagaApi

type SagaApi struct {
	Nasa      *nasa.Nasa
	SagaOrder *SagaOrder
}

func NewSagaApi() *SagaApi {
	return &SagaApi{
		Nasa:      nasa.NewNasa(),
		SagaOrder: NewSagaOrder(),
	}
}

func (this *SagaApi) QueryBasicApiInfoByPage(pageNum, pageSize int) ([]*tables.ApiBasicInfo, error) {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	start := (pageNum - 1) * pageSize
	return dao.DefSagaApiDB.ApiDB.QueryApiBasicInfoByPage(start, pageSize)
}

func (this *SagaApi) QueryApiDetailInfoByApiId(apiId int) (*common.ApiDetailResponse, error) {
	apiDetail, err := dao.DefSagaApiDB.ApiDB.QueryApiDetailInfoById(apiId)
	if err != nil {
		return nil, err
	}
	if apiDetail == nil {
		return nil, nil
	}
	requestParam, err := dao.DefSagaApiDB.ApiDB.QueryRequestParamByApiDetailInfoId(apiDetail.Id)
	if err != nil {
		return nil, err
	}
	errCode, err := dao.DefSagaApiDB.ApiDB.QueryErrorCodeByApiDetailInfoId(apiDetail.Id)
	if err != nil {
		return nil, err
	}
	return &common.ApiDetailResponse{
		ApiId:               apiDetail.ApiId,
		Mark:                apiDetail.Mark,
		ResponseParam:       apiDetail.ResponseParam,
		ResponseExample:     apiDetail.ResponseExample,
		DataDesc:            apiDetail.DataDesc,
		DataSource:          apiDetail.DataSource,
		ApplicationScenario: apiDetail.ApplicationScenario,
		RequestParams:       requestParam,
		ErrorCodes:          errCode,
	}, nil
}

func (this *SagaApi) SearchApiIdByCategoryId(categoryId int) ([]*tables.ApiBasicInfo, error) {
	tagId, err := dao.DefSagaApiDB.ApiDB.QueryTagIdByCategoryId(categoryId)
	if err != nil {
		return nil, err
	}
	apiIds, err := dao.DefSagaApiDB.ApiDB.QueryApiIdByTagId(tagId)
	if err != nil {
		return nil, err
	}
	return dao.DefSagaApiDB.ApiDB.QueryApiByApiIds(apiIds)
}

//newest hot free
func (this *SagaApi) SearchApi() (map[string][]*tables.ApiBasicInfo, error) {
	res := make(map[string][]*tables.ApiBasicInfo)
	//newest
	newest, err := dao.DefSagaApiDB.ApiDB.QueryNewestApiBasicInfo()
	if err != nil {
		return nil, err
	}
	res["newest"] = newest
	hottest, err := dao.DefSagaApiDB.ApiDB.QueryHottestApiBasicInfo()
	if err != nil {
		return nil, err
	}
	res["hottest"] = hottest
	free, err := dao.DefSagaApiDB.ApiDB.QueryFreeApiBasicInfo()
	res["free"] = free
	return res, nil
}
