package req

//消息体
type MqNoticeMsgRequest struct {
	Title        string            `form:"title" json:"title"`
	Content      string            `form:"content" json:"content"`
	Summary      string            `form:"summary" json:"summary"`
	FromUserId   int32             `form:"fromUserId" json:"fromUserId"`
	FromUserName string            `form:"fromUserName" json:"fromUserName"`
	ToUserId     int32             `form:"toUserId" json:"toUserId"`
	ToUserName   string            `form:"toUserName" json:"toUserName"`
	MsgType      int16             `form:"msgType" json:"msgType"`
	MsgCode      string            `form:"msgCode" json:"msgCode"`
	OfferId      int32             `form:"offerId" json:"offerId"`
	JobId        int16             `form:"jobId" json:"jobId"`
	JobOfId      int32             `form:"jobOfId" json:"jobOfId"`
	JobName      string            `form:"jobName" json:"jobName"`
	BusinessName string            `form:"businessName" json:"businessName"`
	KeyPair      map[string]string `form:"keyPair" json:"keyPair"`
}
