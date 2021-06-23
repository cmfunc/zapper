// zaper
// context 从context中获取有效的请求信息
// TODO 考虑context中保留的done、err相关上下文信息，不做处理时，是否会影响到日志的执行goroutine
package zaper

import (
	"context"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

const gfwLogHeaderKey = "_logHeader"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type logHeader struct {
	LogId      string `json:"logid"`
	CallerIp   string `json:"caller_ip"`
	HostIp     string `json:"host_ip"`
	Port       int
	Product    string        `json:"product"`
	Module     string        `json:"module"`
	ServiceId  string        `json:"service_id"`
	InstanceId string        `json:"instance_id"`
	UriPath    string        `json:"uri_path"`
	Tag        string        `json:"tag"`
	Env        string        `json:"env"`
	SVersion   string        `json:"sversion"`
	Stag       stag          `json:"stag"`
	Request    *http.Request `json:"-"`
}

type stag struct {
	StId        int    `json:"st_id"`
	LastLogid   string `json:"lastlogid"`
	StType      string `json:"st_type"`
	StChannel   string `json:"st_channel"`
	StChannelL2 string `json:"st_channel_l2"`
	StChannelL3 string `json:"st_channel_l3"`
	StPos       int    `json:"st_pos"`
	StRelatedId int64  `json:"st_related_id"`
}

func getLogHeaderFromCtx(ctx context.Context) (h *logHeader) {
	h, _ = ctx.Value(gfwLogHeaderKey).(*logHeader)
	return
}

func (h *logHeader) fieldsMarshal() (fields []zap.Field) {
	fields = make([]zap.Field, 0)
	fields = append(fields, zap.String("logid", h.LogId))
	fields = append(fields, zap.String("caller_ip", h.CallerIp))
	fields = append(fields, zap.String("host_ip", h.HostIp))
	fields = append(fields, zap.String("product", h.Product))
	fields = append(fields, zap.String("module", h.Module))
	fields = append(fields, zap.String("service_id", h.ServiceId))
	fields = append(fields, zap.String("instance_id", h.InstanceId))
	fields = append(fields, zap.String("uri_path", h.UriPath))
	fields = append(fields, zap.String("tag", h.Tag))
	fields = append(fields, zap.String("env", h.Env))
	fields = append(fields, zap.String("sversion", h.SVersion))
	stag, _ := json.MarshalToString(h.Stag)// TODO 优化，需实现zapcore.Object接口
	fields = append(fields, zap.String("stag", stag)) 
	return
}
