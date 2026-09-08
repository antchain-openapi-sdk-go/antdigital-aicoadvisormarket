// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

// Description:
//
// Model for initing client
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	//
	// example:
	//
	// http
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	//
	// example:
	//
	// 10
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	//
	// example:
	//
	// 10
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	//
	// example:
	//
	// http://localhost
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	//
	// example:
	//
	// https://localhost
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	//
	// example:
	//
	// cs.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	//
	// example:
	//
	// http://localhost
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	//
	// example:
	//
	// 3
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	//
	// example:
	//
	// Alibabacloud/1
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	//
	// example:
	//
	// TCP
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

// 二维码详情
type QRCodeDetail struct {
	// 位置，Top、Left、Width、Height
	Location []*int64 `json:"location,omitempty" xml:"location,omitempty" require:"true" type:"Repeated"`
	// 二维码类型
	// example:
	//
	// 二维码类型
	Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	// 二维码内容
	// example:
	//
	// 二维码内容
	Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true"`
	// 置信度
	// example:
	//
	// 置信度
	Probability *string `json:"probability,omitempty" xml:"probability,omitempty" require:"true"`
}

func (s QRCodeDetail) String() string {
	return tea.Prettify(s)
}

func (s QRCodeDetail) GoString() string {
	return s.String()
}

func (s *QRCodeDetail) SetLocation(v []*int64) *QRCodeDetail {
	s.Location = v
	return s
}

func (s *QRCodeDetail) SetType(v string) *QRCodeDetail {
	s.Type = &v
	return s
}

func (s *QRCodeDetail) SetContent(v string) *QRCodeDetail {
	s.Content = &v
	return s
}

func (s *QRCodeDetail) SetProbability(v string) *QRCodeDetail {
	s.Probability = &v
	return s
}

// Logo详情
type LogoDetail struct {
	// 位置，Top、Left、Width、Height
	Location []*int64 `json:"location,omitempty" xml:"location,omitempty" require:"true" type:"Repeated"`
	// 类型
	// example:
	//
	// LOGO类型
	Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	// 置信度
	// example:
	//
	// 置信度
	Probability *string `json:"probability,omitempty" xml:"probability,omitempty" require:"true"`
}

func (s LogoDetail) String() string {
	return tea.Prettify(s)
}

func (s LogoDetail) GoString() string {
	return s.String()
}

func (s *LogoDetail) SetLocation(v []*int64) *LogoDetail {
	s.Location = v
	return s
}

func (s *LogoDetail) SetType(v string) *LogoDetail {
	s.Type = &v
	return s
}

func (s *LogoDetail) SetProbability(v string) *LogoDetail {
	s.Probability = &v
	return s
}

// maya响应体
type MayaRedGptResponseDTO struct {
	// 消息的ID
	// example:
	//
	// 123123
	MessageId *string `json:"message_id,omitempty" xml:"message_id,omitempty" require:"true"`
	// 请求ID
	// example:
	//
	// 4564546
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
	// 会话ID
	// example:
	//
	// 312414124
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty" require:"true"`
	// 应答内容
	// example:
	//
	// 你可以问我信息安全的问题
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty" require:"true"`
	// 应答内容格式
	// example:
	//
	// PLAINTEXT
	AnswerFormat *string `json:"answer_format,omitempty" xml:"answer_format,omitempty" require:"true"`
	// 是否回答结束
	// example:
	//
	// true, false
	AnswerEnd *bool `json:"answer_end,omitempty" xml:"answer_end,omitempty" require:"true"`
	// 是否问题有风险
	// example:
	//
	// true, false
	Safe *bool `json:"safe,omitempty" xml:"safe,omitempty" require:"true"`
}

func (s MayaRedGptResponseDTO) String() string {
	return tea.Prettify(s)
}

func (s MayaRedGptResponseDTO) GoString() string {
	return s.String()
}

func (s *MayaRedGptResponseDTO) SetMessageId(v string) *MayaRedGptResponseDTO {
	s.MessageId = &v
	return s
}

func (s *MayaRedGptResponseDTO) SetRequestId(v string) *MayaRedGptResponseDTO {
	s.RequestId = &v
	return s
}

func (s *MayaRedGptResponseDTO) SetSessionId(v string) *MayaRedGptResponseDTO {
	s.SessionId = &v
	return s
}

func (s *MayaRedGptResponseDTO) SetAnswer(v string) *MayaRedGptResponseDTO {
	s.Answer = &v
	return s
}

func (s *MayaRedGptResponseDTO) SetAnswerFormat(v string) *MayaRedGptResponseDTO {
	s.AnswerFormat = &v
	return s
}

func (s *MayaRedGptResponseDTO) SetAnswerEnd(v bool) *MayaRedGptResponseDTO {
	s.AnswerEnd = &v
	return s
}

func (s *MayaRedGptResponseDTO) SetSafe(v bool) *MayaRedGptResponseDTO {
	s.Safe = &v
	return s
}

// 审核同步标签列表
type AuditSyncLabel struct {
	// 标签名：sex-色情
	// example:
	//
	// sex
	Label *string `json:"label,omitempty" xml:"label,omitempty" require:"true"`
	// 检测到单个风险标签的置信度：66.25
	// example:
	//
	// 66.25
	Probability *int64 `json:"probability,omitempty" xml:"probability,omitempty" require:"true"`
	// 风险标签说明
	// example:
	//
	// 疑似色情内容
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 检测到的敏感词，多个词用逗号分隔，部分标签不会返回敏感词：AA,BB,CC
	// example:
	//
	// AA,BB,CC
	RiskWords *string `json:"risk_words,omitempty" xml:"risk_words,omitempty" require:"true"`
}

func (s AuditSyncLabel) String() string {
	return tea.Prettify(s)
}

func (s AuditSyncLabel) GoString() string {
	return s.String()
}

func (s *AuditSyncLabel) SetLabel(v string) *AuditSyncLabel {
	s.Label = &v
	return s
}

func (s *AuditSyncLabel) SetProbability(v int64) *AuditSyncLabel {
	s.Probability = &v
	return s
}

func (s *AuditSyncLabel) SetDescription(v string) *AuditSyncLabel {
	s.Description = &v
	return s
}

func (s *AuditSyncLabel) SetRiskWords(v string) *AuditSyncLabel {
	s.RiskWords = &v
	return s
}

// QRCode审核结果
type QRCodeAuditResult struct {
	// 检测到二维码个数
	DetectNum *int64 `json:"detect_num,omitempty" xml:"detect_num,omitempty" require:"true"`
	// 二维码详情
	Details []*QRCodeDetail `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s QRCodeAuditResult) String() string {
	return tea.Prettify(s)
}

func (s QRCodeAuditResult) GoString() string {
	return s.String()
}

func (s *QRCodeAuditResult) SetDetectNum(v int64) *QRCodeAuditResult {
	s.DetectNum = &v
	return s
}

func (s *QRCodeAuditResult) SetDetails(v []*QRCodeDetail) *QRCodeAuditResult {
	s.Details = v
	return s
}

// 攻击手法二级标签
type AttackSubLabel struct {
	// 提示词攻击手法二级标签
	// example:
	//
	// role_play：角色扮演
	AttackSubLabel *string `json:"attack_sub_label,omitempty" xml:"attack_sub_label,omitempty"`
}

func (s AttackSubLabel) String() string {
	return tea.Prettify(s)
}

func (s AttackSubLabel) GoString() string {
	return s.String()
}

func (s *AttackSubLabel) SetAttackSubLabel(v string) *AttackSubLabel {
	s.AttackSubLabel = &v
	return s
}

// 主题信息
type MeiyouTopicWebInfo struct {
	// 内容文本
	// example:
	//
	// 内容文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 当前楼层
	// example:
	//
	// 1
	CurrentFloor *int64 `json:"current_floor,omitempty" xml:"current_floor,omitempty"`
	// 回复楼层
	// example:
	//
	// 1
	CallBackFloor *int64 `json:"call_back_floor,omitempty" xml:"call_back_floor,omitempty"`
	// 发布时间戳(毫秒)
	//
	PublishTime *int64 `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 图片URL数组(JSON字符串)
	// example:
	//
	// {}
	Images *string `json:"images,omitempty" xml:"images,omitempty"`
	// 用户昵称
	// example:
	//
	// a
	UserNickname *string `json:"user_nickname,omitempty" xml:"user_nickname,omitempty"`
	// 用户ID
	// example:
	//
	// 1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户类型
	//
	// example:
	//
	// q
	UserType *string `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 用户头像URL
	// example:
	//
	// http://test
	UserAvatar *string `json:"user_avatar,omitempty" xml:"user_avatar,omitempty"`
	// 主题ID
	// example:
	//
	// 1
	TopicId *int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 回复楼信息
	// example:
	//
	// test
	CallBackFloorContent *string `json:"call_back_floor_content,omitempty" xml:"call_back_floor_content,omitempty"`
	// oss存储地址
	// example:
	//
	// {}
	OssImages *string `json:"oss_images,omitempty" xml:"oss_images,omitempty"`
}

func (s MeiyouTopicWebInfo) String() string {
	return tea.Prettify(s)
}

func (s MeiyouTopicWebInfo) GoString() string {
	return s.String()
}

func (s *MeiyouTopicWebInfo) SetContent(v string) *MeiyouTopicWebInfo {
	s.Content = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetCurrentFloor(v int64) *MeiyouTopicWebInfo {
	s.CurrentFloor = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetCallBackFloor(v int64) *MeiyouTopicWebInfo {
	s.CallBackFloor = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetPublishTime(v int64) *MeiyouTopicWebInfo {
	s.PublishTime = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetImages(v string) *MeiyouTopicWebInfo {
	s.Images = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetUserNickname(v string) *MeiyouTopicWebInfo {
	s.UserNickname = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetUserId(v string) *MeiyouTopicWebInfo {
	s.UserId = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetUserType(v string) *MeiyouTopicWebInfo {
	s.UserType = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetUserAvatar(v string) *MeiyouTopicWebInfo {
	s.UserAvatar = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetTopicId(v int64) *MeiyouTopicWebInfo {
	s.TopicId = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetCallBackFloorContent(v string) *MeiyouTopicWebInfo {
	s.CallBackFloorContent = &v
	return s
}

func (s *MeiyouTopicWebInfo) SetOssImages(v string) *MeiyouTopicWebInfo {
	s.OssImages = &v
	return s
}

// maya流式响应结果
type MayaStreamResult struct {
	// maya响应数据
	Data *MayaRedGptResponseDTO `json:"data,omitempty" xml:"data,omitempty" require:"true"`
	// 是否成功
	// example:
	//
	// true, false
	Success *bool `json:"success,omitempty" xml:"success,omitempty" require:"true"`
	// 错误码
	// example:
	//
	// 10002
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty" require:"true"`
	// 错误信息
	// example:
	//
	// maya平台调用失败
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty" require:"true"`
}

func (s MayaStreamResult) String() string {
	return tea.Prettify(s)
}

func (s MayaStreamResult) GoString() string {
	return s.String()
}

func (s *MayaStreamResult) SetData(v *MayaRedGptResponseDTO) *MayaStreamResult {
	s.Data = v
	return s
}

func (s *MayaStreamResult) SetSuccess(v bool) *MayaStreamResult {
	s.Success = &v
	return s
}

func (s *MayaStreamResult) SetErrorCode(v string) *MayaStreamResult {
	s.ErrorCode = &v
	return s
}

func (s *MayaStreamResult) SetErrorMsg(v string) *MayaStreamResult {
	s.ErrorMsg = &v
	return s
}

// 二级标签结构
type SubLabelModel struct {
	// 二级标签
	// example:
	//
	// 正常
	SubLabel *string `json:"sub_label,omitempty" xml:"sub_label,omitempty" require:"true"`
	// 风险关键词列表
	RiskWords []*string `json:"risk_words,omitempty" xml:"risk_words,omitempty" type:"Repeated"`
	// 风险关键词索引列表
	RiskWordsIndex []*string `json:"risk_words_index,omitempty" xml:"risk_words_index,omitempty" type:"Repeated"`
	// 三级标签列表
	ThirdLabels []*string `json:"third_labels,omitempty" xml:"third_labels,omitempty" type:"Repeated"`
}

func (s SubLabelModel) String() string {
	return tea.Prettify(s)
}

func (s SubLabelModel) GoString() string {
	return s.String()
}

func (s *SubLabelModel) SetSubLabel(v string) *SubLabelModel {
	s.SubLabel = &v
	return s
}

func (s *SubLabelModel) SetRiskWords(v []*string) *SubLabelModel {
	s.RiskWords = v
	return s
}

func (s *SubLabelModel) SetRiskWordsIndex(v []*string) *SubLabelModel {
	s.RiskWordsIndex = v
	return s
}

func (s *SubLabelModel) SetThirdLabels(v []*string) *SubLabelModel {
	s.ThirdLabels = v
	return s
}

// logo审核结果
type LogoAuditResult struct {
	// 检测到LOGO个数
	DetectNum *int64 `json:"detect_num,omitempty" xml:"detect_num,omitempty" require:"true"`
	// LOGO详情
	Details []*LogoDetail `json:"details,omitempty" xml:"details,omitempty" require:"true" type:"Repeated"`
}

func (s LogoAuditResult) String() string {
	return tea.Prettify(s)
}

func (s LogoAuditResult) GoString() string {
	return s.String()
}

func (s *LogoAuditResult) SetDetectNum(v int64) *LogoAuditResult {
	s.DetectNum = &v
	return s
}

func (s *LogoAuditResult) SetDetails(v []*LogoDetail) *LogoAuditResult {
	s.Details = v
	return s
}

// 领域信息
type FieldModel struct {
	// 领域一级标签
	// example:
	//
	// 金融
	FieldCategory *string `json:"field_category,omitempty" xml:"field_category,omitempty"`
	// 领域二级标签
	// example:
	//
	// 金融
	FieldLabel *string `json:"field_label,omitempty" xml:"field_label,omitempty"`
	// 领域一级标签的等级分数
	FieldScore *int64 `json:"field_score,omitempty" xml:"field_score,omitempty"`
}

func (s FieldModel) String() string {
	return tea.Prettify(s)
}

func (s FieldModel) GoString() string {
	return s.String()
}

func (s *FieldModel) SetFieldCategory(v string) *FieldModel {
	s.FieldCategory = &v
	return s
}

func (s *FieldModel) SetFieldLabel(v string) *FieldModel {
	s.FieldLabel = &v
	return s
}

func (s *FieldModel) SetFieldScore(v int64) *FieldModel {
	s.FieldScore = &v
	return s
}

// 美柚审核信息存储请求
type MeiyouAuditSaveWebRequest struct {
	// 审核记录ID
	// example:
	//
	// 1
	AuditId *int64 `json:"audit_id,omitempty" xml:"audit_id,omitempty"`
	// 主题ID
	// example:
	//
	// 1
	TopicId *int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 内容文本
	// example:
	//
	// test
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 发布时间戳(毫秒)
	// example:
	//
	// 1
	PublishTime *int64 `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 图片URL数组(JSON字符串)
	// example:
	//
	// {}
	Images *string `json:"images,omitempty" xml:"images,omitempty"`
	// 一级业务
	// example:
	//
	// test
	PrimaryBusiness *string `json:"primary_business,omitempty" xml:"primary_business,omitempty"`
	// 二级业务
	// example:
	//
	// 1
	SecondaryBusiness *string `json:"secondary_business,omitempty" xml:"secondary_business,omitempty"`
	// 用户昵称
	// example:
	//
	// 小蜜
	UserNickname *string `json:"user_nickname,omitempty" xml:"user_nickname,omitempty"`
	// 用户ID
	// example:
	//
	// 1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户类型
	// example:
	//
	// INIT
	UserType *string `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 用户头像URL
	// example:
	//
	// http://test
	UserAvatar *string `json:"user_avatar,omitempty" xml:"user_avatar,omitempty"`
	// 操作人
	// example:
	//
	// 小蜜
	AuditOperator *string `json:"audit_operator,omitempty" xml:"audit_operator,omitempty"`
	// 审核楼
	// example:
	//
	// 1
	AuditFloor *int64 `json:"audit_floor,omitempty" xml:"audit_floor,omitempty"`
	// 版本
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 主题信息
	// example:
	//
	// {}
	TopicInfos []*MeiyouTopicWebInfo `json:"topic_infos,omitempty" xml:"topic_infos,omitempty" type:"Repeated"`
}

func (s MeiyouAuditSaveWebRequest) String() string {
	return tea.Prettify(s)
}

func (s MeiyouAuditSaveWebRequest) GoString() string {
	return s.String()
}

func (s *MeiyouAuditSaveWebRequest) SetAuditId(v int64) *MeiyouAuditSaveWebRequest {
	s.AuditId = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetTopicId(v int64) *MeiyouAuditSaveWebRequest {
	s.TopicId = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetContent(v string) *MeiyouAuditSaveWebRequest {
	s.Content = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetPublishTime(v int64) *MeiyouAuditSaveWebRequest {
	s.PublishTime = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetImages(v string) *MeiyouAuditSaveWebRequest {
	s.Images = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetPrimaryBusiness(v string) *MeiyouAuditSaveWebRequest {
	s.PrimaryBusiness = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetSecondaryBusiness(v string) *MeiyouAuditSaveWebRequest {
	s.SecondaryBusiness = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetUserNickname(v string) *MeiyouAuditSaveWebRequest {
	s.UserNickname = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetUserId(v string) *MeiyouAuditSaveWebRequest {
	s.UserId = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetUserType(v string) *MeiyouAuditSaveWebRequest {
	s.UserType = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetUserAvatar(v string) *MeiyouAuditSaveWebRequest {
	s.UserAvatar = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetAuditOperator(v string) *MeiyouAuditSaveWebRequest {
	s.AuditOperator = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetAuditFloor(v int64) *MeiyouAuditSaveWebRequest {
	s.AuditFloor = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetVersion(v string) *MeiyouAuditSaveWebRequest {
	s.Version = &v
	return s
}

func (s *MeiyouAuditSaveWebRequest) SetTopicInfos(v []*MeiyouTopicWebInfo) *MeiyouAuditSaveWebRequest {
	s.TopicInfos = v
	return s
}

// 网关响应模型
type AntCloudProdProviderHttpResponse struct {
	// maya返回结果
	Response *MayaStreamResult `json:"response,omitempty" xml:"response,omitempty" require:"true"`
	// 签名
	// example:
	//
	// 5Okl4F2SNc9L2zrWCF8xZ+QPUyA=
	Sign *string `json:"sign,omitempty" xml:"sign,omitempty" require:"true"`
}

func (s AntCloudProdProviderHttpResponse) String() string {
	return tea.Prettify(s)
}

func (s AntCloudProdProviderHttpResponse) GoString() string {
	return s.String()
}

func (s *AntCloudProdProviderHttpResponse) SetResponse(v *MayaStreamResult) *AntCloudProdProviderHttpResponse {
	s.Response = v
	return s
}

func (s *AntCloudProdProviderHttpResponse) SetSign(v string) *AntCloudProdProviderHttpResponse {
	s.Sign = &v
	return s
}

// 一级标签信息
type LabelModel struct {
	// 一级标签
	// example:
	//
	// 正常
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// 子标签
	SubLabels []*SubLabelModel `json:"sub_labels,omitempty" xml:"sub_labels,omitempty" type:"Repeated"`
}

func (s LabelModel) String() string {
	return tea.Prettify(s)
}

func (s LabelModel) GoString() string {
	return s.String()
}

func (s *LabelModel) SetLabel(v string) *LabelModel {
	s.Label = &v
	return s
}

func (s *LabelModel) SetSubLabels(v []*SubLabelModel) *LabelModel {
	s.SubLabels = v
	return s
}

// 文本同步审核结果
type TextSyncAuditResult struct {
	// 任务id
	// example:
	//
	// 123
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	// 数据Id
	// example:
	//
	// 123
	DataId *string `json:"data_id,omitempty" xml:"data_id,omitempty" require:"true"`
	// 业务Id - 调用方透传
	// example:
	//
	// 123
	BusinessId *string `json:"business_id,omitempty" xml:"business_id,omitempty" require:"true"`
	// 风险等级，根据设置的高低风险分返回，返回值包括： - high：高风险（若命中自定义词库，风险等级默认为高风险） - medium：中风险 - low：低风险 - none：未检测到风险
	// example:
	//
	// high
	RiskLevel *string `json:"risk_level,omitempty" xml:"risk_level,omitempty" require:"true"`
	// 审核标签列表
	// example:
	//
	// [{"label":"sex", "probability":66.5, "description":"疑似色情内容", "riskWords":"AA,BB,CC"}]
	Labels []*AuditSyncLabel `json:"labels,omitempty" xml:"labels,omitempty" require:"true" type:"Repeated"`
}

func (s TextSyncAuditResult) String() string {
	return tea.Prettify(s)
}

func (s TextSyncAuditResult) GoString() string {
	return s.String()
}

func (s *TextSyncAuditResult) SetTaskId(v string) *TextSyncAuditResult {
	s.TaskId = &v
	return s
}

func (s *TextSyncAuditResult) SetDataId(v string) *TextSyncAuditResult {
	s.DataId = &v
	return s
}

func (s *TextSyncAuditResult) SetBusinessId(v string) *TextSyncAuditResult {
	s.BusinessId = &v
	return s
}

func (s *TextSyncAuditResult) SetRiskLevel(v string) *TextSyncAuditResult {
	s.RiskLevel = &v
	return s
}

func (s *TextSyncAuditResult) SetLabels(v []*AuditSyncLabel) *TextSyncAuditResult {
	s.Labels = v
	return s
}

// 美柚审核信息存储
type MeiyouAuditSaveWebInfo struct {
	// 主题ID
	// example:
	//
	// 1
	TopicId *int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 审核记录ID
	// example:
	//
	// 1
	AuditId *int64 `json:"audit_id,omitempty" xml:"audit_id,omitempty" require:"true"`
	// 内容文本
	// example:
	//
	// 小蜜
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 发布时间戳(毫秒)
	PublishTime *int64 `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 图片URL数组(JSON字符串)
	// example:
	//
	// {}
	Images *string `json:"images,omitempty" xml:"images,omitempty"`
	// 一级业务
	// example:
	//
	// 一级业务
	PrimaryBusiness *string `json:"primary_business,omitempty" xml:"primary_business,omitempty"`
	// 二级业务
	// example:
	//
	// 二级业务
	SecondaryBusiness *string `json:"secondary_business,omitempty" xml:"secondary_business,omitempty"`
	// 用户昵称
	// example:
	//
	// 用户昵称
	UserNickname *string `json:"user_nickname,omitempty" xml:"user_nickname,omitempty"`
	// 用户ID
	// example:
	//
	// 用户ID
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户类型
	//
	// example:
	//
	// 用户类型
	UserType *string `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 用户头像URL
	// example:
	//
	// http://
	UserAvatar *string `json:"user_avatar,omitempty" xml:"user_avatar,omitempty"`
	//  主题信息
	// example:
	//
	//  主题信息
	TopicInfos *MeiyouTopicWebInfo `json:"topic_infos,omitempty" xml:"topic_infos,omitempty"`
}

func (s MeiyouAuditSaveWebInfo) String() string {
	return tea.Prettify(s)
}

func (s MeiyouAuditSaveWebInfo) GoString() string {
	return s.String()
}

func (s *MeiyouAuditSaveWebInfo) SetTopicId(v int64) *MeiyouAuditSaveWebInfo {
	s.TopicId = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetAuditId(v int64) *MeiyouAuditSaveWebInfo {
	s.AuditId = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetContent(v string) *MeiyouAuditSaveWebInfo {
	s.Content = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetPublishTime(v int64) *MeiyouAuditSaveWebInfo {
	s.PublishTime = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetImages(v string) *MeiyouAuditSaveWebInfo {
	s.Images = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetPrimaryBusiness(v string) *MeiyouAuditSaveWebInfo {
	s.PrimaryBusiness = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetSecondaryBusiness(v string) *MeiyouAuditSaveWebInfo {
	s.SecondaryBusiness = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetUserNickname(v string) *MeiyouAuditSaveWebInfo {
	s.UserNickname = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetUserId(v string) *MeiyouAuditSaveWebInfo {
	s.UserId = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetUserType(v string) *MeiyouAuditSaveWebInfo {
	s.UserType = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetUserAvatar(v string) *MeiyouAuditSaveWebInfo {
	s.UserAvatar = &v
	return s
}

func (s *MeiyouAuditSaveWebInfo) SetTopicInfos(v *MeiyouTopicWebInfo) *MeiyouAuditSaveWebInfo {
	s.TopicInfos = v
	return s
}

// 图片审核结果
type ImageAuditResult struct {
	// 任务ID
	// example:
	//
	// 任务ID
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty" require:"true"`
	// 数据ID
	// example:
	//
	// 数据ID
	DataId *string `json:"data_id,omitempty" xml:"data_id,omitempty" require:"true"`
	// LOGO审核结果
	LogoAuditResult *LogoAuditResult `json:"logo_audit_result,omitempty" xml:"logo_audit_result,omitempty" require:"true"`
	// QRCode审核结果
	QrCodeAuditResult *QRCodeAuditResult `json:"qr_code_audit_result,omitempty" xml:"qr_code_audit_result,omitempty" require:"true"`
}

func (s ImageAuditResult) String() string {
	return tea.Prettify(s)
}

func (s ImageAuditResult) GoString() string {
	return s.String()
}

func (s *ImageAuditResult) SetTaskId(v string) *ImageAuditResult {
	s.TaskId = &v
	return s
}

func (s *ImageAuditResult) SetDataId(v string) *ImageAuditResult {
	s.DataId = &v
	return s
}

func (s *ImageAuditResult) SetLogoAuditResult(v *LogoAuditResult) *ImageAuditResult {
	s.LogoAuditResult = v
	return s
}

func (s *ImageAuditResult) SetQrCodeAuditResult(v *QRCodeAuditResult) *ImageAuditResult {
	s.QrCodeAuditResult = v
	return s
}

// 提示词攻击手法一级标签
type AttackLabel struct {
	// 提示词攻击手法一级标签
	// example:
	//
	// jailbreak：越狱攻击
	AttackLabel *string `json:"attack_label,omitempty" xml:"attack_label,omitempty"`
	// 提示词攻击手法二级标签列表
	AttackSubLabels *AttackSubLabel `json:"attack_sub_labels,omitempty" xml:"attack_sub_labels,omitempty"`
}

func (s AttackLabel) String() string {
	return tea.Prettify(s)
}

func (s AttackLabel) GoString() string {
	return s.String()
}

func (s *AttackLabel) SetAttackLabel(v string) *AttackLabel {
	s.AttackLabel = &v
	return s
}

func (s *AttackLabel) SetAttackSubLabels(v *AttackSubLabel) *AttackLabel {
	s.AttackSubLabels = v
	return s
}

// 美柚审核信息
type MeiyouAuditInfo struct {
	// 主键id
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// 审核记录ID
	// example:
	//
	// 1
	AuditId *int64 `json:"audit_id,omitempty" xml:"audit_id,omitempty" require:"true"`
	// 主题ID
	// example:
	//
	// 1
	TopicId *int64 `json:"topic_id,omitempty" xml:"topic_id,omitempty"`
	// 内容文本
	// example:
	//
	// 内容文本
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 发布时间戳
	// example:
	//
	// 2018-10-10T10:10:00Z
	PublishTime *string `json:"publish_time,omitempty" xml:"publish_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 图片URL数组(JSON字符串)
	// example:
	//
	// {}
	Images *string `json:"images,omitempty" xml:"images,omitempty"`
	// 一级业务
	// example:
	//
	// 1
	PrimaryBusiness *string `json:"primary_business,omitempty" xml:"primary_business,omitempty"`
	// 二级业务
	// example:
	//
	// 2
	SecondaryBusiness *string `json:"secondary_business,omitempty" xml:"secondary_business,omitempty"`
	// 用户昵称
	// example:
	//
	// 小蜜
	UserNickname *string `json:"user_nickname,omitempty" xml:"user_nickname,omitempty"`
	// 用户ID
	// example:
	//
	// 1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户类型
	// example:
	//
	// test
	UserType *string `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 用户头像URL
	// example:
	//
	// http://test
	UserAvatar *string `json:"user_avatar,omitempty" xml:"user_avatar,omitempty"`
	// 审核结果
	// example:
	//
	// OK
	AuditResult *string `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	// 审核原因
	// example:
	//
	// 30
	AuditReason *string `json:"audit_reason,omitempty" xml:"audit_reason,omitempty"`
	// 审核原因
	// example:
	//
	// test
	AuditReasonMsg *string `json:"audit_reason_msg,omitempty" xml:"audit_reason_msg,omitempty"`
	// 操作人
	// example:
	//
	// 小蜜
	AuditOperator *string `json:"audit_operator,omitempty" xml:"audit_operator,omitempty"`
	// 操作时间戳
	// example:
	//
	// 2018-10-10T10:10:00Z
	AuditTime *string `json:"audit_time,omitempty" xml:"audit_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// oss 转存后的图像地址
	// example:
	//
	// {}
	OssImages *string `json:"oss_images,omitempty" xml:"oss_images,omitempty"`
	// 创建时间
	// example:
	//
	// 2018-10-10T10:10:00Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 更新时间
	// example:
	//
	// 2018-10-10T10:10:00Z
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
}

func (s MeiyouAuditInfo) String() string {
	return tea.Prettify(s)
}

func (s MeiyouAuditInfo) GoString() string {
	return s.String()
}

func (s *MeiyouAuditInfo) SetId(v int64) *MeiyouAuditInfo {
	s.Id = &v
	return s
}

func (s *MeiyouAuditInfo) SetAuditId(v int64) *MeiyouAuditInfo {
	s.AuditId = &v
	return s
}

func (s *MeiyouAuditInfo) SetTopicId(v int64) *MeiyouAuditInfo {
	s.TopicId = &v
	return s
}

func (s *MeiyouAuditInfo) SetContent(v string) *MeiyouAuditInfo {
	s.Content = &v
	return s
}

func (s *MeiyouAuditInfo) SetPublishTime(v string) *MeiyouAuditInfo {
	s.PublishTime = &v
	return s
}

func (s *MeiyouAuditInfo) SetImages(v string) *MeiyouAuditInfo {
	s.Images = &v
	return s
}

func (s *MeiyouAuditInfo) SetPrimaryBusiness(v string) *MeiyouAuditInfo {
	s.PrimaryBusiness = &v
	return s
}

func (s *MeiyouAuditInfo) SetSecondaryBusiness(v string) *MeiyouAuditInfo {
	s.SecondaryBusiness = &v
	return s
}

func (s *MeiyouAuditInfo) SetUserNickname(v string) *MeiyouAuditInfo {
	s.UserNickname = &v
	return s
}

func (s *MeiyouAuditInfo) SetUserId(v string) *MeiyouAuditInfo {
	s.UserId = &v
	return s
}

func (s *MeiyouAuditInfo) SetUserType(v string) *MeiyouAuditInfo {
	s.UserType = &v
	return s
}

func (s *MeiyouAuditInfo) SetUserAvatar(v string) *MeiyouAuditInfo {
	s.UserAvatar = &v
	return s
}

func (s *MeiyouAuditInfo) SetAuditResult(v string) *MeiyouAuditInfo {
	s.AuditResult = &v
	return s
}

func (s *MeiyouAuditInfo) SetAuditReason(v string) *MeiyouAuditInfo {
	s.AuditReason = &v
	return s
}

func (s *MeiyouAuditInfo) SetAuditReasonMsg(v string) *MeiyouAuditInfo {
	s.AuditReasonMsg = &v
	return s
}

func (s *MeiyouAuditInfo) SetAuditOperator(v string) *MeiyouAuditInfo {
	s.AuditOperator = &v
	return s
}

func (s *MeiyouAuditInfo) SetAuditTime(v string) *MeiyouAuditInfo {
	s.AuditTime = &v
	return s
}

func (s *MeiyouAuditInfo) SetOssImages(v string) *MeiyouAuditInfo {
	s.OssImages = &v
	return s
}

func (s *MeiyouAuditInfo) SetGmtCreate(v string) *MeiyouAuditInfo {
	s.GmtCreate = &v
	return s
}

func (s *MeiyouAuditInfo) SetGmtModified(v string) *MeiyouAuditInfo {
	s.GmtModified = &v
	return s
}

// 更新美柚itag关系信息
type UpdateMeiyouItagRelationWebInfo struct {
	// 主键id
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty" require:"true"`
	// 审核记录ID
	// example:
	//
	// 1
	AuditId *string `json:"audit_id,omitempty" xml:"audit_id,omitempty"`
	// itag数据ID
	// example:
	//
	// 1
	ItagDataId *int64 `json:"itag_data_id,omitempty" xml:"itag_data_id,omitempty"`
	// 美柚任务审核结果推送状态
	// example:
	//
	// INIT
	MeiyouAuditState *string `json:"meiyou_audit_state,omitempty" xml:"meiyou_audit_state,omitempty"`
	// 审核不通过原因
	// example:
	//
	// 34
	RefuseReson *string `json:"refuse_reson,omitempty" xml:"refuse_reson,omitempty"`
	// 审核不通过图片序号
	// example:
	//
	// ["1","2"]
	RefuseImages *string `json:"refuse_images,omitempty" xml:"refuse_images,omitempty"`
	// 美柚itag关联状态
	// example:
	//
	// 1
	AuditState *string `json:"audit_state,omitempty" xml:"audit_state,omitempty"`
	// 美柚itag关联状态
	// example:
	//
	// 1
	TopicState *string `json:"topic_state,omitempty" xml:"topic_state,omitempty"`
	// 审核结果
	// example:
	//
	// 1
	AuditResult *string `json:"audit_result,omitempty" xml:"audit_result,omitempty"`
	// 审核人员
	// example:
	//
	// 1
	AuditUser *string `json:"audit_user,omitempty" xml:"audit_user,omitempty"`
	// 审核时间
	// example:
	//
	// 1
	AuditTime *int64 `json:"audit_time,omitempty" xml:"audit_time,omitempty"`
}

func (s UpdateMeiyouItagRelationWebInfo) String() string {
	return tea.Prettify(s)
}

func (s UpdateMeiyouItagRelationWebInfo) GoString() string {
	return s.String()
}

func (s *UpdateMeiyouItagRelationWebInfo) SetId(v int64) *UpdateMeiyouItagRelationWebInfo {
	s.Id = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetAuditId(v string) *UpdateMeiyouItagRelationWebInfo {
	s.AuditId = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetItagDataId(v int64) *UpdateMeiyouItagRelationWebInfo {
	s.ItagDataId = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetMeiyouAuditState(v string) *UpdateMeiyouItagRelationWebInfo {
	s.MeiyouAuditState = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetRefuseReson(v string) *UpdateMeiyouItagRelationWebInfo {
	s.RefuseReson = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetRefuseImages(v string) *UpdateMeiyouItagRelationWebInfo {
	s.RefuseImages = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetAuditState(v string) *UpdateMeiyouItagRelationWebInfo {
	s.AuditState = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetTopicState(v string) *UpdateMeiyouItagRelationWebInfo {
	s.TopicState = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetAuditResult(v string) *UpdateMeiyouItagRelationWebInfo {
	s.AuditResult = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetAuditUser(v string) *UpdateMeiyouItagRelationWebInfo {
	s.AuditUser = &v
	return s
}

func (s *UpdateMeiyouItagRelationWebInfo) SetAuditTime(v int64) *UpdateMeiyouItagRelationWebInfo {
	s.AuditTime = &v
	return s
}

type QueryAitechCommAdvisormarketDataRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 查询类型STATION_LIST、STATION_DETAIL、REACHABLE_ROADS、VEHICLE_TYPE_DISTRIBUTION、ENTRY_FLOW、FRONT_FLOW、SURROUNDING_FLOW、NEW_ENTRY_VEHICLES、FREQUENT_ENTRY、INFREQUENT_ENTRY
	Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	// 行政区名称，如“浙江省”、“杭州市”；type=STATION_LIST时必填
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// 品牌清单筛选，如 ["中石油", "中石化"]；不传则返回所有品牌；仅 type=STATION_LIST 时生效
	Brands []*string `json:"brands,omitempty" xml:"brands,omitempty" type:"Repeated"`
	// 站点名称列表，支持批量查询；type 非 STATION_LIST 时必填
	Names []*string `json:"names,omitempty" xml:"names,omitempty" type:"Repeated"`
	// 起始日期，yyyy-MM-dd；时间类查询时必填
	Startdate *string `json:"startdate,omitempty" xml:"startdate,omitempty"`
	// 结束日期，yyyy-MM-dd；时间类查询时必填
	Enddate *string `json:"enddate,omitempty" xml:"enddate,omitempty"`
	// 最低进站次数阈值，默认 3；仅 type=FREQUENT_ENTRY 时生效
	Mintimes *string `json:"mintimes,omitempty" xml:"mintimes,omitempty"`
	// 最高进站次数阈值（不含），默认 3；仅 type=INFREQUENT_ENTRY 时生效
	Maxtimes *string `json:"maxtimes,omitempty" xml:"maxtimes,omitempty"`
}

func (s QueryAitechCommAdvisormarketDataRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAitechCommAdvisormarketDataRequest) GoString() string {
	return s.String()
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetAuthToken(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetProductInstanceId(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetType(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.Type = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetRegion(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.Region = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetBrands(v []*string) *QueryAitechCommAdvisormarketDataRequest {
	s.Brands = v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetNames(v []*string) *QueryAitechCommAdvisormarketDataRequest {
	s.Names = v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetStartdate(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.Startdate = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetEnddate(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.Enddate = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetMintimes(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.Mintimes = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataRequest) SetMaxtimes(v string) *QueryAitechCommAdvisormarketDataRequest {
	s.Maxtimes = &v
	return s
}

type QueryAitechCommAdvisormarketDataResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 状态码，"0" 表示成功，非 "0" 为失败
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 状态描述
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID，用于问题排查
	Requestid *string `json:"requestid,omitempty" xml:"requestid,omitempty"`
	// 服务端响应时间戳（毫秒）
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// 业务数据 JSON 字符串，调用方需反序列化后使用，具体结构见各 type 定义
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s QueryAitechCommAdvisormarketDataResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAitechCommAdvisormarketDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetReqMsgId(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetResultCode(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetResultMsg(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetCode(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.Code = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetMessage(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.Message = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetRequestid(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.Requestid = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetTimestamp(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.Timestamp = &v
	return s
}

func (s *QueryAitechCommAdvisormarketDataResponse) SetData(v string) *QueryAitechCommAdvisormarketDataResponse {
	s.Data = &v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

// Description:
//
// # Init client with Config
//
// @param config - config contains the necessary information to create a client
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(config)) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

// Description:
//
// # Encapsulate the request and invoke the network
//
// @param action - api name
//
// @param protocol - http or https
//
// @param method - e.g. GET
//
// @param pathname - pathname of every api
//
// @param request - which contains request params
//
// @param runtime - which controls some details of call api, such as retry times
//
// @return the response
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":          "retry",
		"readTimeout":        tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":     tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":          tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":         tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":            tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":       tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":  tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDuration":  tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":        tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost": tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("1.0.1"),
				"_prod_code":       tea.String("AICOADVISORMARKET"),
				"_prod_channel":    tea.String("default"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res, _err := util.AssertAsMap(obj)
			if _err != nil {
				return _result, _err
			}

			resp, _err := util.AssertAsMap(res["response"])
			if _err != nil {
				return _result, _err
			}

			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

// Description:
//
// Description: 客户统计结果查询
//
// Summary: 客户统计结果查询
func (client *Client) QueryAitechCommAdvisormarketData(request *QueryAitechCommAdvisormarketDataRequest) (_result *QueryAitechCommAdvisormarketDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAitechCommAdvisormarketDataResponse{}
	_body, _err := client.QueryAitechCommAdvisormarketDataEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Description:
//
// Description: 客户统计结果查询
//
// Summary: 客户统计结果查询
func (client *Client) QueryAitechCommAdvisormarketDataEx(request *QueryAitechCommAdvisormarketDataRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryAitechCommAdvisormarketDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryAitechCommAdvisormarketDataResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("aitech.comm.advisormarket.data.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
