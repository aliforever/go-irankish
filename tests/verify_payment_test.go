package tests

import (
	"fmt"
	"testing"

	"github.com/aliforever/go-irankish"
)

func TestIranKish_VerifyPayment(t *testing.T) {
	ik := irankish.IranKish{MerchantId: "BB00", Sha1Key: irankish.Sha1Key}
	payment := irankish.VerifyPayment{}
	payment.Token = "82243619433299825345"
	payment.ReferenceNumber = "1"
	ik.Verify = &payment
	verify, err := ik.VerifyPayment()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(verify)
}
