package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-irankish"
)

func TestIranKish_MakeToken(t *testing.T) {
	ik := irankish.IranKish{MerchantId: "BB00"}
	payment := irankish.Payment{}
	payment.Amount = "20000"
	payment.InvoiceId = "1"
	payment.CallbackUrl = "http://localhost/test/"
	ik.Payment = &payment
	token, err := ik.MakeToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Token:", token.Token)
}
