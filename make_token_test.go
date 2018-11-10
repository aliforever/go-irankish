package irankish

import (
	"testing"
	"fmt"
	"github.com/kr/pretty"
)

func TestIranKish_MakeToken(t *testing.T) {
	ik := IranKish{MerchantId:"BB00"}
	payment := Payment{}
	payment.Amount = "20000"
	payment.InvoiceId = "1"
	payment.CallbackUrl = "http://localhost/test/"
	ik.Payment = &payment
	token, err := ik.MakeToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	pretty.Println(token)
}
