package bank

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertBalance := func(t testing.TB, want Bitcoint, wallet Wallet) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoint(10)
		assertBalance(t, want, wallet)

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(20)}
		wallet.Withdraw(Bitcoint(10))

		want := Bitcoint(10)
		assertBalance(t, want, wallet)
	})

	t.Run("withdraw insufficient founds", func(t *testing.T) {
		startingBalance := Bitcoint(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoint(100))

		assertError(t, err, "insufficient funds")

		assertBalance(t, startingBalance, wallet)
	})

}
