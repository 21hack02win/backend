// Package handler provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package handler

import "github.com/gofrs/uuid"

// Defines values for WsEvent.
const (
	WsEventDRAWCANCEL WsEvent = "DRAW_CANCEL"

	WsEventDRAWFINISH WsEvent = "DRAW_FINISH"

	WsEventDRAWREADY WsEvent = "DRAW_READY"

	WsEventDRAWSEND WsEvent = "DRAW_SEND"

	WsEventDRAWSTART WsEvent = "DRAW_START"

	WsEventGAMESTART WsEvent = "GAME_START"

	WsEventODAICANCEL WsEvent = "ODAI_CANCEL"

	WsEventODAIFINISH WsEvent = "ODAI_FINISH"

	WsEventODAIREADY WsEvent = "ODAI_READY"

	WsEventODAISEND WsEvent = "ODAI_SEND"

	WsEventREQUESTGAMESTART WsEvent = "REQUEST_GAME_START"

	WsEventROOMNEWMEMBER WsEvent = "ROOM_NEW_MEMBER"

	WsEventROOMSETOPTION WsEvent = "ROOM_SET_OPTION"

	WsEventROOMUPDATEOPTION WsEvent = "ROOM_UPDATE_OPTION"
)

// 回答の入力の完了を解除する (ルームの各員 -> サーバー)
type AnswerCancelEvent struct {
	Id string `json:"id"`
}

// 回答の入力が完了していることを通知する (ルームの各員 -> サーバー)
type AnswerReadyEvent struct {
	Answer string `json:"answer"`
}

// 絵が飛んできて，回答する (サーバー -> ルーム各員)
type AnswerStartEvent struct {
	Img string `json:"img"`
}

// 絵が書き終わっている通知を解除する (ルームの各員 -> サーバー)
type DrawCancelEvent map[string]interface{}

// 全員が絵を完了したことor制限時間が来たことを通知する (サーバー -> ルーム全員)
// クライアントは絵を送信する
type DrawFinishEvent struct {
	Id string `json:"id"`
}

// 絵が書き終わっていることを通知する (ルームの各員 -> サーバー)
type DrawReadyEvent map[string]interface{}

// 絵を送信する (ルームの各員 -> サーバー)
//
// -> (DRAWフェーズが終わってなかったら) また，DRAW_START が飛んでくる
type DrawSendEvent struct {
	Img string `json:"img"`
}

// キャンバス情報とお題を送信する (サーバー -> ルーム各員)
type DrawStartEvent struct {
	AllDrawPhaseNum float32 `json:"allDrawPhaseNum"`
	DrawPhaseNum    int     `json:"drawPhaseNum"`
	Img             string  `json:"img"`
	Odai            string  `json:"odai"`
	TimeLimit       int     `json:"timeLimit"`
}

// ゲームの開始を通知する (サーバー -> ルーム全員)
type GameStartEvent struct {
	OdaiHint  string `json:"odaiHint"`
	TimeLimit int    `json:"timeLimit"`
}

// ルーム参加リクエスト
type JoinRoomRequest struct {
	Avatar int    `json:"avatar"`
	Name   string `json:"name"`
	RoomId string `json:"roomId"`
}

// 新規ルーム情報
type NewRoom struct {
	// ルームID
	RoomId string `json:"roomId"`

	// ユーザーUUID
	UserId uuid.UUID `json:"userId"`
}

// NewRoomRequest defines model for NewRoomRequest.
type NewRoomRequest struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

// お題の入力の完了を解除する (ルームの各員 -> サーバー)
type OdaiCancelEvent map[string]interface{}

// 全員がお題の入力を完了したことor制限時間が来たことを通知する (サーバー -> ルーム全員)
// クライアントはお題を送信する
type OdaiFinishEvent map[string]interface{}

// お題の入力が完了していることを通知する (ルームの各員 -> サーバー)
type OdaiReadyEvent struct {
	Id *string `json:"id,omitempty"`
}

// お題を送信する (ルームの各員 -> サーバー)
type OdaiSendEvent struct {
	Odai string `json:"odai"`
}

// ルーム情報
type Room struct {
	Capacity int       `json:"capacity"`
	HostId   uuid.UUID `json:"hostId"`
	Members  []User    `json:"members"`
	RoomId   string    `json:"roomId"`
	UserId   uuid.UUID `json:"userId"`
}

// 部屋に追加のメンバーが来たことを通知する (サーバー -> ルーム全員)
type RoomNewMemberEvent map[string]interface{}

// ゲームのオプションを設定する (ホスト -> サーバー)
type RoomSetOptionEvent map[string]interface{}

// RoomUpdateOptionEvent defines model for RoomUpdateOptionEvent.
type RoomUpdateOptionEvent map[string]interface{}

// User defines model for User.
type User struct {
	Avatar int       `json:"avatar"`
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"userId"`
}

// Websocketイベントのリスト
type WsEvent string

// RoomId defines model for roomId.
type RoomId string

// JoinRoomJSONBody defines parameters for JoinRoom.
type JoinRoomJSONBody JoinRoomRequest

// CreateRoomJSONBody defines parameters for CreateRoom.
type CreateRoomJSONBody NewRoomRequest

// WsJSONBody defines parameters for Ws.
type WsJSONBody struct {
	Body *interface{} `json:"body,omitempty"`

	// Websocketイベントのリスト
	Type *WsEvent `json:"type,omitempty"`
}

// JoinRoomJSONRequestBody defines body for JoinRoom for application/json ContentType.
type JoinRoomJSONRequestBody JoinRoomJSONBody

// CreateRoomJSONRequestBody defines body for CreateRoom for application/json ContentType.
type CreateRoomJSONRequestBody CreateRoomJSONBody

// WsJSONRequestBody defines body for Ws for application/json ContentType.
type WsJSONRequestBody WsJSONBody
