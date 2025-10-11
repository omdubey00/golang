package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds from the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(50))
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}

func assertNoError(t *testing.T, err error) { // this asserts that there should be no error if then tell us in test.
	if err != nil {
		t.Error("did not expected an error but got one")
	}
}
func assertError(t *testing.T, err error, want string) { // this terminology asserts there shouold be an error
	if err == nil {
		t.Fatal("Did not get an error but expected one") // and if there is not an error then there is something wrong with the test.
	}

	if err.Error() != want {
		t.Errorf("got the error : %s , wanted the error : %s,  didnot get the exepected error description", err.Error(), want)
	}
}
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf(" got : %s , want : %s ", got, want)
	}
}

// Here I have to understand that the test is written to throw an error its not the real one like when it doesn't then
// tell me. Because the entire concept of writing of the insufficient funds test is that it should throw an error.
//
//
