package attacks

import (
	"fmt"

	"github.com/kavehmz/prime"

	"github.com/sourcekris/goRsaTool/keys"
	"github.com/sourcekris/goRsaTool/ln"

	fmp "github.com/sourcekris/goflint"
)

// go seems so fast making small primes we can probably make this much larger
const maxprimeint = 1000000

// SmallQ iterate small primes < maxprimeint and test them as factors of N at a memory cost.
func SmallQ(t *keys.RSA) error {
	if t.Key.D != nil {
		return nil
	}

	primes := prime.Primes(maxprimeint)
	modp := new(fmp.Fmpz)

	for _, p := range primes {
		modp = modp.Mod(t.Key.N, fmp.NewFmpz(int64(p)))
		if modp.Cmp(ln.BigZero) == 0 {
			t.PackGivenP(fmp.NewFmpz(int64(p)))
			fmt.Printf("[+] Small q Factor found\n")
			return nil
		}
	}

	return nil
}
