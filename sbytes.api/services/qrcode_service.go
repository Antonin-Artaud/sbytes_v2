package services

import (
	"errors"
	"fmt"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
	"os"
)

type qrCodeImpl struct {
}

func (receiver *qrCodeImpl) Create(data string, qrcodeName string) string {
	qrc, err := qrcode.New(data)

	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return errors.New("could not generate QRCode").Error()
	}

	w, err := standard.New("./tempQrCodeService" + qrcodeName)

	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return ""
	}

	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}

	return ""
}

func (receiver *baseService) Delete(qrcodeName string) {
	err := os.Remove("./tempQrCodeService" + qrcodeName)
	if err != nil {
		return
	}
}
