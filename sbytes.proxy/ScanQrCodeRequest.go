package main

type ScanQrCodeRequest struct {
	FireBaseId string     `json:"fireBaseId" binding:"required"`
	Url        string     `json:"url" binding:"required"`
	Credential Credential `json:"credential" binding:"required"`
}

type Credential struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
