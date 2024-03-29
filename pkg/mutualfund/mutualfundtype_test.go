package mutualfund_test

import (
	"fmt"
	"testing"

	"github.com/rakhmadbudiono/mutual-fund-aggregator/pkg/mutualfund"

	"github.com/stretchr/testify/assert"
)

var (
	mixed       mutualfund.MutualFundType = 0
	stock       mutualfund.MutualFundType = 1
	fixedincome mutualfund.MutualFundType = 2
	moneymarket mutualfund.MutualFundType = 3
)

func TestNewFixedIncomeMutualFund(t *testing.T) {
	mf := mutualfund.NewMutualFund(name, fixedincome)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)
}

func TestNewMixedMutualFund(t *testing.T) {
	mf := mutualfund.NewMutualFund(name, mixed)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)
}

func TestNewMoneyMarketMutualFund(t *testing.T) {
	mf := mutualfund.NewMutualFund(name, moneymarket)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)
}

func TestNewStockMutualFund(t *testing.T) {
	mf := mutualfund.NewMutualFund(name, stock)

	assert.Equal(
		t,
		name,
		mf.GetName(),
		fmt.Sprintf("Expect name to be %s, but got %s", name, mf.GetName()),
	)
}
