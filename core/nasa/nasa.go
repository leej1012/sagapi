package nasa

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ontio/sagapi/core/http"
	"github.com/ontio/sagapi/dao"
	"github.com/ontio/sagapi/models"
	"github.com/ontio/sagapi/models/tables"
	"github.com/ontio/sagapi/sagaconfig"
	"sync"
	"sync/atomic"
)

var (
	apod = "https://api.nasa.gov/planetary/apod?api_key=%s"
	feed = "https://api.nasa.gov/neo/rest/v1/feed?start_date=%s&end_date=%s&api_key=%s"
)

type Nasa struct {
	apiKeyInvokeFreCache *sync.Map //apikey -> ApiKeyInvokeFre
	freqlock             *sync.Mutex
}

func NewNasa() *Nasa {
	return &Nasa{
		apiKeyInvokeFreCache: new(sync.Map),
		freqlock:             new(sync.Mutex),
	}
}

func (this *Nasa) beforeCheckApiKey(apiKey string, tx *sql.Tx) (*models.ApiKeyInvokeFre, error) {
	this.freqlock.Lock()
	defer this.freqlock.Unlock()
	key, err := this.getApiKeyInvokeFre(apiKey)
	if err != nil {
		return nil, err
	}

	if key.UsedNum >= int32(key.RequestLimit) {
		return nil, fmt.Errorf("apikey: %s, useNum: %d, limit:%d", apiKey, key.UsedNum, key.RequestLimit)
	}

	err = dao.DefSagaApiDB.ApiDB.UpdateApiKeyInvokeFre(key.ApiKey, int(key.UsedNum), key.ApiId, int(key.InvokeFre), tx)
	if err != nil {
		return nil, err
	}

	key.UsedNum += 1
	key.InvokeFre += 1

	return key, nil
}

func (this *Nasa) Apod(apiKey string) (res []byte, e error) {
	tx, err := dao.DefSagaApiDB.ApiDB.DB().Begin()
	if err != nil {
		return nil, err
	}

	key, err := this.beforeCheckApiKey(apiKey, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	url := fmt.Sprintf(apod, sagaconfig.DefSagaConfig.NASAAPIKey)
	res, e = http.DefClient.Get(url)
	if e != nil {
		atomic.AddInt32(&key.UsedNum, -1)
		atomic.AddInt32(&key.InvokeFre, -1)
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (this *Nasa) Feed(startDate, endDate string, apiKey string) (res []byte, e error) {
	tx, err := dao.DefSagaApiDB.ApiDB.DB().Begin()
	if err != nil {
		return nil, err
	}

	key, err := this.beforeCheckApiKey(apiKey, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	url := fmt.Sprintf(feed, startDate, endDate, sagaconfig.DefSagaConfig.NASAAPIKey)
	res, e = http.DefClient.Get(url)
	if e != nil {
		atomic.AddInt32(&key.UsedNum, -1)
		atomic.AddInt32(&key.InvokeFre, -1)
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return
}

func (this *Nasa) ApodParams(params []tables.RequestParam) ([]byte, error) {
	if len(params) == 1 && params[0].ParamName == "apiKey" {
		return this.Apod(params[0].ValueDesc)
	}
	return nil, errors.New("Apod params error")
}

func (this *Nasa) FeedParams(params []tables.RequestParam) ([]byte, error) {
	if len(params) == 3 && params[0].ParamName == "startDate" && params[1].ParamName == "endDate" && params[2].ParamName == "apiKey" {
		return this.Feed(params[0].ValueDesc, params[1].ValueDesc, params[2].ValueDesc)
	}
	return nil, errors.New("Apod params error")
}

func (this *Nasa) getApiKeyInvokeFre(apiKey string) (*models.ApiKeyInvokeFre, error) {
	keyIn, ok := this.apiKeyInvokeFreCache.Load(apiKey)
	if !ok || keyIn == nil {
		key, err := dao.DefSagaApiDB.ApiDB.QueryApiKeyAndInvokeFreByApiKey(apiKey)
		if err != nil {
			return nil, err
		}
		this.apiKeyInvokeFreCache.Store(apiKey, key)
		return key, nil
	} else {
		return keyIn.(*models.ApiKeyInvokeFre), nil
	}
}

func (this *Nasa) updateApiKeyInvokeFre(key *models.ApiKeyInvokeFre, tx *sql.Tx) error {
	return dao.DefSagaApiDB.ApiDB.UpdateApiKeyInvokeFre(key.ApiKey, int(key.UsedNum), key.ApiId, int(key.InvokeFre), tx)
}
