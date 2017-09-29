package transaction

import (
	"testing"
	"time"

	"github.com/icrowley/fake"
)

var amount1, amount2 = 25.11, 76.46
var collection = Collection{
	{
		Date:      time.Now(),
		Name:      fake.FirstName(),
		Account:   fake.CreditCardNum("visa"),
		ToAccount: "",
		Code:      "",
		Direction: Incoming,
		Amount:    amount1,
		Type:      "",
		Notes:     "",
	},
	{
		Date:      time.Now(),
		Name:      fake.FirstName(),
		Account:   fake.CreditCardNum("visa"),
		ToAccount: "",
		Code:      "",
		Direction: Incoming,
		Amount:    amount2,
		Type:      "",
		Notes:     "",
	},
}

func TestGetTotalAmount(t *testing.T) {
	expectedResult := amount1 + amount2

	testResult := collection.GetTotalAmount(Incoming)
	if testResult != expectedResult {
		t.Errorf("GetTotalAmount: expected %.2f got %.2f", expectedResult, testResult)
	}
}

func TestReportByName(t *testing.T) {

}
