package merkle

import "testing"

func TestSimple(t *testing.T) {
	input := []string{"abcd", "efgh", "ijkl"}
	correct := "A8uyUMnjjSXD0Tb0Tm-oyrJ3tVnif0ry4l0YbTSxvos="
	if h := Hash(input); h != correct {
		t.Fatalf("Hash incorrect, recevied: %s, expected: %s", h, correct)
	}
}

func TestBitCoin(t *testing.T) {
	input := []string{
		"4e10436ca8206a2dd760dd351210a5120a3824d4eb53011be0a7b9a33b368208",
		"76dc5788be4a8cf6925aff15fd8c8fbf6417b4ad6c30a1ac12cd117e95c5820b",
		"7452bfa629b104985f7c937e0f7836206935d83872882c88ae183234fe9bcf97",
		"e325a8a968368aeb6d89bcdb362d311833f5a9fe5a80f3f0730b684922439a68",
		"8510c531b585e77a66a986cb0dfdd0ca280ff0747d2dca0d6fa87b0f8af4810a",
		"f9a5d31e7894c3983d38215060c55665db0024ad7fb373fa58db7316dd223ea9",
		"83e74406b0876fed2db187444dd0a4f3eedad42e9adb32ce82b3ff729fe77b58",
		"0f3e5f5b833dfad8ff19115ecd29fa40566aace67f2880da0ab4fa1acac00bcc",
		"a1743f0803926ef4343e217c78324ec3e3ebc4cfc7c96739f3696b3c510cc7d1",
		"5015ed6455e7d3fb50ce5ef1f63888fe7d4c37042f17f5f89bea176dce4ca0bd",
		"7b14b05552f053a862df22824f92e94f155722f9e5d91341b934a6de010b6560",
		"2564b40ef226c73eb63409f9dbdcc64aeacca6e3a7136e86156a64573ba3f6d1",
	}
	correct := "bb988af992654871e8cefe8bbe05e9f9679611eadcfa53980ee515978eebcd52"
	if h := Hash(input); h != correct {
		t.Fatalf("Hash incorrect, recevied: %s, expected: %s", h, correct)
	}
}
