package crt

import (
	"testing"

	"github.com/sourcekris/goRsaTool/keys"
	"github.com/sourcekris/goRsaTool/ln"

	fmp "github.com/sourcekris/goflint"
)

func TestAttack(t *testing.T) {
	for _, tc := range []struct {
		name   string
		n      *fmp.Fmpz
		primes []*fmp.Fmpz
		dp     *fmp.Fmpz
		dq     *fmp.Fmpz
		c      *fmp.Fmpz
		want   string
	}{
		{
			name: "large int crt with both primes",
			primes: []*fmp.Fmpz{
				ln.FmpString("8637633767257008567099653486541091171320491509433615447539162437911244175885667806398411790524083553445158113502227745206205327690939504032994699902053229"),
				ln.FmpString("12640674973996472769176047937170883420927050821480010581593137135372473880595613737337630629752577346147039284030082593490776630572584959954205336880228469"),
			},
			dp:   ln.FmpString("6500795702216834621109042351193261530650043841056252930930949663358625016881832840728066026150264693076109354874099841380454881716097778307268116910582929 "),
			dq:   ln.FmpString("783472263673553449019532580386470672380574033551303889137911760438881683674556098098256795673512201963002175438762767516968043599582527539160811120550041 "),
			c:    ln.FmpString("24722305403887382073567316467649080662631552905960229399079107995602154418176056335800638887527614164073530437657085079676157350205351945222989351316076486573599576041978339872265925062764318536089007310270278526159678937431903862892400747915525118983959970607934142974736675784325993445942031372107342103852"),
			want: "noxCTF{W31c0m3_70_Ch1n470wn}",
		},
		{
			name:   "large int crt with 1 prime and N",
			n:      ln.FmpString("109185520996312541892170311492569840427989404929338472706839065360161994637255409112468354512184559845894072909465678390142560223819189118515217770244432400366111866477825962880380741745193047922889710523658653189145814237895884199725959755553890822782508549045051612849654194388769655622729647728707719176401"),
			primes: []*fmp.Fmpz{ln.FmpString("8637633767257008567099653486541091171320491509433615447539162437911244175885667806398411790524083553445158113502227745206205327690939504032994699902053229")},
			dp:     ln.FmpString("6500795702216834621109042351193261530650043841056252930930949663358625016881832840728066026150264693076109354874099841380454881716097778307268116910582929 "),
			dq:     ln.FmpString("783472263673553449019532580386470672380574033551303889137911760438881683674556098098256795673512201963002175438762767516968043599582527539160811120550041 "),
			c:      ln.FmpString("24722305403887382073567316467649080662631552905960229399079107995602154418176056335800638887527614164073530437657085079676157350205351945222989351316076486573599576041978339872265925062764318536089007310270278526159678937431903862892400747915525118983959970607934142974736675784325993445942031372107342103852"),
			want:   "noxCTF{W31c0m3_70_Ch1n470wn}",
		},
	} {

		// Construct the bare minimum RSA precomputed values key.
		k := &keys.RSA{
			Key: keys.FMPPrivateKey{
				N:      tc.n,
				Primes: tc.primes,
				Precomputed: &keys.PrecomputedValues{
					Dp: tc.dp,
					Dq: tc.dq,
				},
			},
			CipherText:  ln.NumberToBytes(tc.c),
			KeyFilename: tc.name,
		}

		if err := Attack([]*keys.RSA{k}); err != nil {
			t.Errorf("%s failed - got unexpected error: %v", tc.name, err)
		}

		got := string(k.PlainText)
		if got != tc.want {
			t.Errorf("%s failed - got / want mismatched: %s / %s", tc.name, got, tc.want)
		}
	}
}
