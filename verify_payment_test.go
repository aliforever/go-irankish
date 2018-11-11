package irankish

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestIranKish_VerifyPayment(t *testing.T) {
	ik := IranKish{MerchantId: "BB00", Sha1Key: sha1Key}
	payment := VerifyPayment{}
	payment.Token = "82243619433299825345"
	payment.ReferenceNumber = "1"
	ik.Verify = &payment
	verify, err := ik.VerifyPayment()
	if err != nil {
		fmt.Println(err)
		return
	}
	pretty.Println(verify)
}
