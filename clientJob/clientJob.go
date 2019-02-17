package clientJob

import (
	"bytes"
	"encoding/json"
	"github.com/Deansquirrel/goClientManager/global"
	"github.com/Deansquirrel/goClientManager/object"
	"github.com/Deansquirrel/goToolEnvironment"
	log "github.com/Deansquirrel/goToolLog"
	"io/ioutil"
	"net/http"
	"time"
)

type clientJob struct {
	interval int
}

func NewClientJob(interval int) *clientJob {
	return &clientJob{
		interval: interval,
	}
}

//启动Client事件循环
func (cj *clientJob) StartClientJob() {
	cj.worker()
	time.AfterFunc(time.Duration(cj.interval*1000*1000*1000), func() {
		cj.StartClientJob()
	})
}

//事件循环体
func (cj *clientJob) worker() {
	log.Debug("ClientJob start working")
	defer log.Debug("ClientJob work complete")
	clientInfo := cj.getClientInfo()
	data, err := json.Marshal(clientInfo)
	if err != nil {
		log.Warn("数据准备失败：" + err.Error())
		return
	}
	cj.sendData(data, cj.getUrl())
}

//获取客户端信息
func (cj *clientJob) getClientInfo() *object.ClientInfo {
	clientInfo := object.ClientInfo{}
	clientInfo.OsInfo.Type = goToolEnvironment.GetOsType()
	clientInfo.OsInfo.Name = goToolEnvironment.GetOsName()
	osVer, err := goToolEnvironment.GetOsVer()
	if err != nil {
		log.Warn(err.Error())
	} else {
		clientInfo.OsInfo.Ver = osVer
	}
	hostName, err := goToolEnvironment.GetHostName()
	if err != nil {
		log.Warn(err.Error())
	} else {
		clientInfo.OsInfo.HostName = hostName
	}
	internetAddr, err := goToolEnvironment.GetInternetAddr()
	if err != nil {
		log.Warn(err.Error())
	} else {
		clientInfo.NetInfo.InternetIp = internetAddr
	}
	intranetAddr, err := goToolEnvironment.GetIntranetAddr()
	if err != nil {
		log.Warn(err.Error())
	} else {
		clientInfo.NetInfo.IntranetIp = intranetAddr
	}
	return &clientInfo
}

//POST发送数据
func (cj *clientJob) sendData(data []byte, url string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Warn("构造http请求数据时发生错误：" + err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Warn("发送http请求时错误：" + err.Error())
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	rData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Warn("读取http返回数据时发生错误：" + err.Error())
		return
	}
	log.Debug(string(rData))
	//============================================================================
}

func (cj *clientJob) getUrl() string {
	return global.ClientConfig.Config.ServerUrl + "/ClientInfo" + "/Info"
}
