package valkeystore

import (
	"testing"

	"github.com/EktaMind/Thank_ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/kalibroida/ThankyouCoin_Node/inter/validatorpk"
)

var (
	pubkey1, _ = validatorpk.FromString("0xc0045ea4ce3ab0748574f0290dadcb45545aff82d8baa72e5b4c84a19d2e1f16fb3dc487430b4189ded650a94148e57a60ca8cbf4da414dbfd3b072f0a5b9a746235")
	key1       = common.FromHex("e77b3e0e1bfb52a1e22b73dd7941336443363c4942c5c70869302f66940eefc2")
	name1      = "c0045ea4ce3ab0748574f0290dadcb45545aff82d8baa72e5b4c84a19d2e1f16fb3dc487430b4189ded650a94148e57a60ca8cbf4da414dbfd3b072f0a5b9a746235"
	file1      = common.FromHex("7b2274797065223a3139322c227075626b6579223a2230343565613463653361623037343835373466303239306461646362343535343561666638326438626161373265356234633834613139643265316631366662336463343837343330623431383964656436353061393431343865353761363063613863626634646134313464626664336230373266306135623961373436323335222c2263727970746f223a7b22636970686572223a226165732d3132382d637472222c2263697068657274657874223a2262623662363638336636316633363231636131313530366137633666366661616130313761663833613861656163373139303666336332643664613265353132222c22636970686572706172616d73223a7b226976223a223963643333343332373230386164616666373162653936323434643339666263227d2c226b6466223a22736372797074222c226b6466706172616d73223a7b22646b6c656e223a33322c226e223a343039362c2270223a362c2272223a382c2273616c74223a2232383232656134316338366462366435353065333733326565333661343639393765656438326661366530646536383234343530373562356261396461633934227d2c226d6163223a2261623366363934396234306130366664326264396663386237316664643566353933386164333866616236366236396636663931393363373362336439613939227d7d")
	pubkey2, _ = validatorpk.FromString("0xc00459b25a40ac4af6d114deb2f899bb371869b467955dd3106302309263c6c7786209306dae5564cbeb75805ff517bb49dce467f785c138837782a0c0becf4b122c")
	key2       = common.FromHex("72c7c0305f3bb74720683aad5342b44bec96efed8256ed76bb3ba6421947f0a5")
	name2      = "c00459b25a40ac4af6d114deb2f899bb371869b467955dd3106302309263c6c7786209306dae5564cbeb75805ff517bb49dce467f785c138837782a0c0becf4b122c"
	file2      = common.FromHex("7b2274797065223a3139322c227075626b6579223a2230343539623235613430616334616636643131346465623266383939626233373138363962343637393535646433313036333032333039323633633663373738363230393330366461653535363463626562373538303566663531376262343964636534363766373835633133383833373738326130633062656366346231323263222c2263727970746f223a7b22636970686572223a226165732d3132382d637472222c2263697068657274657874223a2237373833373830643537633835373530366234646139636461643632316638653161346132386130376335636264343564653332663536313566323630396532222c22636970686572706172616d73223a7b226976223a226338346563613438333231346364393461353933663539336362633032616437227d2c226b6466223a22736372797074222c226b6466706172616d73223a7b22646b6c656e223a33322c226e223a343039362c2270223a362c2272223a382c2273616c74223a2265353061623135366138636430633537363431336331346563373162336637666465373466363362633161323631376233343465363933616136633733626237227d2c226d6163223a2231383535343832663266363837313236393931613233313665613061656535386636363932306637653366633330653038656133663832666430636162323232227d7d")
)

func testGet(t *testing.T, keystore RawKeystoreI, expPubkey validatorpk.PubKey, expKey []byte, auth string) {
	require := require.New(t)

	wrongPubkey := expPubkey
	wrongPubkey.Type++
	key, err := keystore.Get(wrongPubkey, auth)
	require.EqualError(err, ErrNotFound.Error())
	require.Nil(key)

	wrongPubkey = expPubkey
	wrongPubkey.Raw = []byte{0}
	key, err = keystore.Get(wrongPubkey, auth)
	require.EqualError(err, ErrNotFound.Error())
	require.Nil(key)

	key, err = keystore.Get(expPubkey, auth)
	require.NoError(err)
	require.Equal(expPubkey.Type, key.Type)
	require.Equal(expKey, key.Bytes)

	key, err = keystore.Get(expPubkey, auth+"1")
	require.EqualError(err, "could not decrypt key with given password")
	require.Nil(key)
}
