package model

import "time"

type Word struct {
	// ID 自增id
	ID int32 `json:"id"`
	// SrcContent 需要翻译单词
	SrcContent string `json:"src_content"`
	// DstContent 翻译结果
	DstContent string `json:"dst_content"`
	// DstAttr 词性
	DstAttr string `json:"dst_attr"`
	// DstExplain 解释
	DstExplain string `json:"dst_explain"`
	// DstExample 例子
	DstExample string `json:"dst_example"`
	// MediaID 素材id
	MediaID   string    `json:"media_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//type Word struct {
//	Id           int       `json:"id" xorm:"pk autoincr"`
//	SrcContent   string    `json:"src_content"`
//	DstContent   string    `json:"dst_content"`
//	DstExplain   string    `json:"dst_explain"`
//	DstExample   string    `json:"dst_example"`
//	MediaId      string    `json:"media_id"`
//	CreatedAt    time.Time `json:"created_at"`
//	UpdatedAt    time.Time `json:"updated_at"`
//}
