package fuzz03

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_0(t *testing.T) {
	var expected, actual U6Ba28tAovfMN2Gps74G0I5ErVW18M0rpFA7o7v78k1155Oa50taNaINQ14Mfv0S8CQHHo4272N03el3uMYVC370201VR3mKRYud
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, U6Ba28tAovfMN2Gps74G0I5ErVW18M0rpFA7o7v78k1155Oa50taNaINQ14Mfv0S8CQHHo4272N03el3uMYVC370201VR3mKRYud{}, expected)
	require.Equal(t, U6Ba28tAovfMN2Gps74G0I5ErVW18M0rpFA7o7v78k1155Oa50taNaINQ14Mfv0S8CQHHo4272N03el3uMYVC370201VR3mKRYud{}, actual)

	actual = U6Ba28tAovfMN2Gps74G0I5ErVW18M0rpFA7o7v78k1155Oa50taNaINQ14Mfv0S8CQHHo4272N03el3uMYVC370201VR3mKRYud{
		YmswLkr8Dgw2: rando.String(),
		FB5:          rando.Int8(),
		DfOtR3hc4BO6M8hCedOUE6O3Js87au8IP07UuhlTVsP807V85434OelJ08bO:     rando.Byte(),
		Y4IxMY5J5q5ou8358GsF6vQ784orH8Cj4gg8Lk0uwsTHwb0T:                 rando.Uint(),
		Xj6WsmID64IN1ycdcOkpf0O23c68WL5UkJ7hLI3y:                         rando.Uint32(),
		OQ7THTAIhV2VpaiKSD4xuTRiuw1LTtM7qyxTw3:                           rando.Uint8(),
		ToTj35ccb85A6ekp46r6ROc5SgEX4sVl1Md63rWII8O8ELv10EE354s8M0a8Gy:   rando.Int16(),
		DL0nDLKeo325W5P0008RBNUuCE3s262lSrXQ6p30A27XmwrUcIkKWT15A52Omk1G: rando.Rune(),
		AaIfn2E85rmH1x251p8:      rando.Int32(),
		Mt184oX0Ro4x4T82V5rsTOY:  rando.Int8(),
		C1vXkuYrKkd32lVEgQv1LJGD: rando.Uint32(),
		BKr2:                     rando.Bool(),
		B2kN848uQn1CXPpXP3d5vh6rAkQ33X7He846T3074L245E3uKEQ7FnkVI1v4: rando.Uint64(),
		Qe23F71nf138W0: rando.Rune(),
		EoD5sv8XET4eJE1M7DVE5WvdViH3Gj7C1wRg8g2ODOBxMosDoCv0ASPY36qh28AL76O2cup172g561HLS3XS7m38L4hJYx: rando.Uint64(),
		Lgwds7s7GpomEUqtienXd00nLIpbmwmU60TuqYx:                                                        rando.Byte(),
		M8fr2E7t6Im5aRsb7Q4gmwh4s3ssjki276VSh2bt1Djj6L:                                                 rando.Bool(),
		LsgWd470C4GU8461MayF3L24q5bt2Stb5:                                                              rando.Float64(),
		WYDNq14085mbH6h6lRpff08ik8U4IV5PAB5Nxri661YS1i7Jh5NW2822xTpQxrX3cWheQ5Hc1PxI33y:                rando.Time(),
		C7U8wtKq67Hj0x2rSoQ5cyEV:                                                                       rando.Uint16(),
		R5jb2V36AAs32Um6o4d7DP86rATW4REH82C355ohp0BmytuSXDPAq5D3384vlA2U1MMk87Uf783CyWd6PG2:            rando.Bool(),
		WPxPg7t8eEKe8BX005wbFcu381f2PPh2vfauWdUEJ3ticDPY0cur7366aXq:                                    rando.Uint64(),
		Ob3e7p2C8pH8ejwBUy61nFbHFjRueurwmi72j2tfr35cVjS5O3y:                                            rando.Bytes(),
		U17xJNRrn56JXiwOY2I2tE:                                                                         rando.Int32(),
		FyQw246H0g8WSB5Gg27H12U02qaxGAp3UNCg2504U7d78G3M5DwJBu2cEw0NxVlhi15w2:                          rando.Float32(),
		MyGN5ejNrXI6q7Q4M3R640NU67WKY3Q3M6:                                                             rando.Uint8(),
		D57w4L0XKdNOKXHU8Nldg3g4rvcjd44rUj:                                                             rando.Uint64(),
		Gi6uWWcsAyW62180u1QvTxgm1JAa5Ur6sW85StLuDm0EbUJg1qh:                                            rando.Uint32(),
		X0S04XSMSFsD4yimwcVptaVTC107bX0iaGhvPFRuKP:                                                     rando.Uint32(),
		RFuyK3bxptLK6l0a4w6204lIPmufL3LDcc1kTWutolY7JNWdCYxbF3bTp:                                      rando.Int64(),
		EnmLK1177yoa4Clg7n4rr873goMnL7klUAcJ13J6w2CnTHhgOj:                                             rando.Int32(),
		OMgOr088rM8RK0C7nCJ3IAbI6rl512g64nrx1RtW18RcA6V2fo7738R1VB0GO6p51s6klyWYM3KSqAtr:               rando.Uint32(),
		W7R21Qu80xlJSD04IyU5G:                                                                          rando.Uint64(),
		Ww:                                                                                             rando.Uint32(),
		FbH70qR073dEIWDCXxpJ4a0qJhC4W72g2pS7D6gEQoYU2y7AW:                                              rando.Bytes(),
		JC15jt63PL0QrJ6Y6k406A8gcE5acWPLOU14ey155t5q278VSuWl8qtkigws:                                   rando.Int(),
		K81iHA2S2U58WP3mGn5NRwokr08dE:                                                                  rando.Int8(),
		KWNupuB872iLb7o77bpv5X46o7XiA52rVllRy3735I:                                                     rando.Int8(),
		KQG3at6K83cs0Uu57lpa2c8cXd6Sf1c4LQyE612chsV67XhV2HdyK3ymhd:                                     rando.Int(),
		VlvgXa2AhcAFR203Ak2IhvbG6yBH76EkUTFC5uvlV4hYUdBna2cvb6Sh0j2jTXDHhf6Rjb0jF7uBR3KLo3:             rando.Int8(),
		Xu7tb4fiyv3443:     rando.Float32(),
		H6pK0obynl2xCovLTQ: rando.Uint(),
		FJnOx87UYQuf8tBDQx2p13u8d51womM0lIiGDJ7336L562l5m5c4qg6t2tal0p67hcV:                           rando.Rune(),
		IF17s8dr4564vW0ySknh4lmxdNKkKS5Bs7MqrIbe0cDR7Q3R8S71G6N4BSmvC:                                 struct{}{},
		UiM7DeG6vVFWf6x28h246YGLVyUGY15g7ppvkGmGE14y11aEbq1Q72TytReLMF003p3S7dvMkhdlhpR5r7OP5P1GYY4vL: rando.Int32(),
		P0c2w4LcfBBO5oQx2L8541hY5Gq4WGMoQsa2wE72BE6h43oUG17pUmc80XYMt87b6:                             rando.Uint8(),
		V6mx4W0SrM7h38j35nDoD1yCJg5u:                 rando.Int32(),
		JauN4lWtqUoPjPjtu5i1tsvlxHtnn228KVHgb:        rando.Int(),
		I30p6eGXN0E26oGIL6y2P5PILei4113fPpTG3dlE3Si2: rando.Uint32(),
		H: struct{}{},
		JApJX5w85QrBN8Gx6Vd7HnQpb7UXJhrx7Eh6vcfIPW:                                              rando.Byte(),
		W2Jwe5l34SF14305wWd2w33fWCo6XjqEuvi7D0m3wWNda837iQ16tJ71o0ry44PuWSdi1h470i8sgt57l5rCMs7: rando.Float64(),
		L25053SD1a3A43KvK30ac066Y1SL6SvV0ttWOxTufhjpavt37PW4f1o5m7:                              rando.Uint64(),
		X526723l8y8NK2nCX6gFyLx:      rando.Int64(),
		Y2R0HOwwlx1K8PA657S0UOxfXgPX: rando.Uint8(),
		NYPBko0r11Ie6F4BtT1qN6Wv0ovL0pXyS72G02XV6IYLgBgfal6RJ3wU4lHVKDya3m1L0gjRefGd7e8G2O8475bnOx3: rando.Byte(),
		KjOXx4VqJuY01327itq0cMJ: rando.String(),
		I41i0:                   rando.Float64(),
		K43U6foxWwro6M4J1P:      rando.Bytes(),
		QeRYJnnTLBpVcVYd5BoP8gu0x2vBpaMQ6SGg68vcISubbTE676AXrc7Uw61: rando.Float64(),
		K5qIxy7r4i61Dn2I: rando.Time(),
		N73Y8sR1qnP37C2umdayTB0f7p7kDRehElRpdLOte8AW4jm1YrlCP2kbEjM0qb1EkyFBsNkMHdC: rando.Float32(),
		O1: rando.Time(),
		PqN5YPsr14qM5724j1A1psJiGvViBiUQ7IuB33Aowoxy7plRoF2AQkmH8581MqpQE1hFxW6T0N:  rando.Int64(),
		BUw556Nqj3ytxnel3bm05caCa5uTgDK0sDv730qrt62srrfS1LoOM821:                    rando.Bytes(),
		B5iXRo45537P8Np1PLu1pqeRD6S5dCR8skLecU65k5h7UgC1iO82vSc0iWlnC352I7sK20FH4YJ: rando.String(),
		Liq4r3:       rando.Float32(),
		NkPL31kYNMlQ: rando.Int8(),
		M0cTVE3ugBkQ1LoCYpg4M2be01GYHh3u0kkXCoRf5aCAc81G8Q8Yru5yxGp03:                        rando.Uint(),
		I255rW7f35Wux633KX46sE76134s0q44MYJfS5r6M6r2F53Mjo4IHo3S6Qjk4an3231co2hALMsVLEw:      rando.Float64(),
		Qj6Nu36K3a2TLm4Dn3136s8jhtoBbJ7I585400sJ62fpfiQKg5QD2g2eEJclsr0r24HK7803PME8eBIf6cJG: rando.Float32(),
		UajCPLMc1j135Y74dk3yS6Gj6DC618UDf6262A4pUB0GH8JY7aJjp7kAQVjO3Q40xAQCe7t:              rando.Bytes(),
		YR3m3m55o8lk153g56syAoG: rando.Uint16(),
		V1rbMQNgMpNbBI1nC1cOxo55BaG002gqbSAN1aJTjpMTq6iRfrY5aWv6jt27nTV43G26827jY1L8mBApEyvXI4gsklSWSH: rando.Int8(),
		YCUi7gefL3bN8m8AVaXv2i2Or43Jruqf6b7I5MaG:                                                       rando.Uint8(),
		I1x65d4VlQOPjiW5an0NXbls4cos3d3g2Nk3T7715wM236UGDaSsr77OYOtDm4YC:                               rando.String(),
		H0xmq10YD74U08b5uGh2l0eOpvbw2qL55c23hF13vi15Tp62p8OPBfjwaHMEqKx05ilHlN3Cd47eAKIKHLP1x06V8MA7v4: rando.String(),
		UBoO3Ydf2: rando.Int64(),
		Q2:        rando.Float32(),
		P2AV5ydE43o7R3gICkA1pfvl5C11B48uYo8h6OxirTq735d65G0OB502:                                      rando.Uint16(),
		O8GNbDOR0VDLa7gepyLT3WaaU51SJbLyI844j0X0IG5wYo2vBy1S6KqUi6h8XD5UA4U:                           rando.Rune(),
		VCXjCc5wi3RC7p8553xG1pqEDV4bwh:                                                                rando.Int64(),
		G7O2WdU858fiFlj2Eh00Qi0o0UJ1Eo5yGA26J08LEx6S1ld6A8D864pXDKC1371xkU8L3L0fuLw77ge0v5gBNr74NxFrA: rando.Int64(),
		M058HpLSba: rando.Uint(),
		Gc5Dmq4XRoHK5b5M4B3R2y0X8MCK6qvVACW5Eo1y6In7342Vuq2JKO2yJLdi0qBAvKEIV2ETDm8SHTg8J: rando.Uint16(),
		Habi:                                 rando.Uint64(),
		Fn0jLrHddFY4BjR7lY5CtfK6HL7G123meNYM: rando.Int(),
		Hwx1kPR7nk0055PlmkkI6SDQSj:           rando.Float32(),
		JY43kklK4fbWWgsO3OCa617s868KhT2hu08SdN8E60dEEtNX134qDXy0d0RsRr54cCDcV5KF8F8SEMPmF3D:      struct{}{},
		Nch8snS4JBFtq0tlc66bbj0XgO2156CXao6P30q45qyXdP5oBkXYo:                                    rando.Uint8(),
		S2ftw6CjCpV0l0ie1d2X50wT6Kmqphq1EaC4FBATe60T0j3AAFbCo3yJsYQ5754vANBej4mGog832baX78oTf7ek: rando.Bytes(),
		Y2ej7y: rando.Int(),
		T8s83xw2VGQXbx75f4YXtjxaD4yJ2nW3v2xr5GH4m572a63vD3AQqviL1p7Vjw2rXVbC2NOhrFPj6OmGd7wu5S8L8ykNFtb:     rando.Byte(),
		TqK1kI25Ra35fU21Nd5HUQV6P6U6CBD0i2b5KKq0pyl4DO0nfXv533H582k402btyETaMjFTXly487oY51Pa8S:              rando.Uint64(),
		U6K0Gile4TqLh77VLmO0wXJntY8A2K4FU2Mj86155fqkNx:                                                      rando.Uint(),
		SPGKsb4U17NLrw6f3D53Dcvt0BC4422J1br44biN641WS0d5V8yv6rEAIohp4S0k7aoU0CO0Y:                           rando.Int64(),
		X7Al578fv1pVN2BJxJ60p24AuBWiDAGL83Fu30Tv1w7vWUp8Hf23NAJ5Cgll1U3j0Ql2FlCP2WkNfQuMVj0NoU38E323kduhPGk: rando.Byte(),
		XYtyA45acpr: rando.Uint8(),
		WGfpP5D83jy240J0o033ExaQR1G6i1tOD5R8koKQBifcY17P:                            rando.Bytes(),
		CCJq6l32gjXkm6PLqGggTX4Ds1H821bbA3r3HP6KiW5Tys83TJ6SB2damIHGKTs874Bm8J4eIP7: rando.Uint64(),
		D6J37KOh31dq15wuWeM7k3ik0ys55Q2V5nJmRSH5h0Y1CB:                              rando.Byte(),
		VUl8nexnfQen: rando.Int16(),
		TLN:          rando.Int16(),
		MWPHR37T2r875tMc6vmUJg2Rx3SGmIb6JB5j52FfqY0VXu33MJxN:                         rando.Uint16(),
		Vsf08730jSEcoNxCNcCk26f7Fd8KVN7dBK8J0bqfE6q7WQ2n3OQ6b7135M6B6YA80x1:          struct{}{},
		TH1gcwG3ra86ejMrVX1bq6Bx4FRxhatW1G36QGrA:                                     rando.Uint8(),
		El0lD5gXvu601EENN338dX7c8O7qD6a2u64X06QI13g77G4qse7Vg7yn727ui716JRfSsb1:      rando.Int(),
		TqI6Ry2JgQ6rS845B8mhe31S38VPt35cq44VxSsE77Wnpfmlsp4IvVAQ4733lNkd84VS00t578J:  rando.Uint16(),
		DS7SNB237Ya76NsxIK1d5W1peC5r5i1U8FDt:                                         rando.Float64(),
		G4ocC2hiJpL28D:                                                               rando.Rune(),
		PNiwx3K52q41YCuLqu3lgpxoA4P653v5H5jFX11u3vbS3VtaOmB06y0yf6NyAspOe0UR1O4qE0n8: rando.Uint(),
		KqJfR4UE2Y2X3iiD7RJICs65lrAPyia2rNFFr35263R8vdgg2w6sGRS1HNAD5kxm:             rando.Uint32(),
		L16sOdk2hr7VA:                        struct{}{},
		P2u0U4a6NR3G65S5I7y67j5uU2MW3pMypJmd: rando.Float64(),
		Ro8R3AFWg0kD120KTI0Wv2OQU7uGbxugRF3GmqOf4qbqF533wKCefi7hG17e1Aft1digEJWG8op26n5: struct{}{},
		O2wC1CIKAFCKR68MykRpvTWy6Q5g: rando.Int16(),
		L2j7b531X1ihaJn6pu2ek2v3WdmbN867Gvw50iH6ot2Sit0Nb6UN33Xxn651I8dfv408Q81oUm2EurWd5kJvEwB: rando.Int64(),
		UlFA1IDU83uru855o84HYBK64W4teyL18l7282HIMrqWpvGiKH86Kb2AedAsrjm:                         rando.Uint8(),
		C68vT35W147w0t4j0wo0yM0bVJMTE64Wnj4pGJX35Q11H2ocrb01Av3UM0:                              rando.String(),
		LkP8r3g55A33cG8o1X0dQ6vDmsxkM8xr7MijI3764ro6HbORKr7m063XSRd3GPTMhE5Fa0v24hIe2qsO602mhX:  rando.Int16(),
		E53776TgHEH2xvl0HlI8u85Xj451GFK1E3v3X6CPv5A3PFCbVl4hGm:                                  rando.Uint8(),
		Y4eg67RO1b2CLCL7H7SsPX2bsw65eEb01Cu5C8vC2Ner40ytEep0mLuBV2vMe0S88NE0E8KUh72E8:           rando.Uint8(),
		FHquL5rYEk543pm324LHg08tS7D02133dcFmJx8Ilw3vi4TJ5a6q3PHrgDqsmUMwyeWu48nQn0FbwN2uiEh2x:   rando.Uint64(),
		Wv2qRIkOgt1D4b45pN0t40X04nV3f:                                                           rando.Int(),
		D54GF0atqJiJouFiV0uk:                                                                    rando.Bool(),
		J0C074ChJ74p2f42eX5v33jGHkqY45y0m8Oe7Eli0vwoD12yfebDKHRgp67n3gdqYinCC:                   rando.Rune(),
		W40BE7b61Jvc0c5syJ3I62sesO410oR2Bg1W8mRn04kEjwCY12xt5AbUTmcskSNBk7dG5QxM3tlVy8g:         rando.String(),
		U3gt4UOFBfn7T22Msi4HXH3mO2YULF:                                                          rando.Byte(),
		EsmkoN4Ra8kH88:                                                                          rando.Int64(),
		Kfxx2L3Rh6a5Tf50p4c5:                                                                    rando.Int16(),
		O6CVCbxx22qifwL:                                                                         rando.Uint8(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, U6Ba28tAovfMN2Gps74G0I5ErVW18M0rpFA7o7v78k1155Oa50taNaINQ14Mfv0S8CQHHo4272N03el3uMYVC370201VR3mKRYud{}, expected)
	require.NotEqual(t, U6Ba28tAovfMN2Gps74G0I5ErVW18M0rpFA7o7v78k1155Oa50taNaINQ14Mfv0S8CQHHo4272N03el3uMYVC370201VR3mKRYud{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_1(t *testing.T) {
	var expected, actual MYh4g5vnmfR0Cl0fDWMn13r7KlRB8J6c3la4bN4nD3m8wcT0eOFSUk3t
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, MYh4g5vnmfR0Cl0fDWMn13r7KlRB8J6c3la4bN4nD3m8wcT0eOFSUk3t{}, expected)
	require.Equal(t, MYh4g5vnmfR0Cl0fDWMn13r7KlRB8J6c3la4bN4nD3m8wcT0eOFSUk3t{}, actual)

	actual = MYh4g5vnmfR0Cl0fDWMn13r7KlRB8J6c3la4bN4nD3m8wcT0eOFSUk3t{
		N71Sb0VaK57Or5s5:                                      rando.Uint16(),
		HJn5oC6o67h07Jq7n5KFn5v3k1dkj560gQOF:                  rando.Int8(),
		XHvtfmW60wYlW52gQ0D2L2Pn0uH8f0lHA3PCNmgrm6pRNDXij68O2: rando.Uint16(),
		WNjSRx06bG2rRkmKJ2r0Hu5j3y0Q88110uSsPOha61Uxlwc5sO24uLR2rDxwMduP56O6EXPqAMf7SYwmuOwInYEahG7: rando.Uint64(),
		TffGEPr0FL:          rando.Float64(),
		U4fmGkQtfA5AaBl0bh7: rando.Bool(),
		RUc5e5dH7ia2L2awyRKKWyr486YPgNy2r1YTvP7X8f24xR3WJ6iuw3tllrk688F77peKkEqswNy34ad4Dvb7lB414M: rando.Uint(),
		QIAfLs1MJ6AmV1Otl47: rando.Uint16(),
		E3270gSfDLppevk0pdh62T6l0fk42u65v70rPR0sfo72H7GBLNsc1LnHi52OOvo5y0: rando.Int16(),
		NiR2LdRgk7qYcJ67qkC6od6:         struct{}{},
		E2e3VnjcT11t3mnb04gpWEeFb4KvYN3: rando.Uint16(),
		TLt5o8BUd6X12iDsfG40587PNoU7G0uds7wIg5RVcw72Nyedhi8X0Mj11WFmoBSDgN4QmdwHWdamS:                        rando.Int(),
		JU675641NIwp1A51UHL86gxU0cmp2FYkwCSsi2oGF3LUQ06ku42H:                                                 rando.Uint32(),
		N0i6Okd130ck5cBYxtEW121JNOS15SgLr2HRcLPLWjO1A8x31x2il56BfK7AvnQCtGQeo7p4pLj:                          rando.Uint64(),
		RHhV0W185xrV162y0sOfCdood46WJ6GlR7rUa4kQwt017P5q168l43aSF5UP:                                         rando.Uint32(),
		I5N7120u6b5t6ICoqLBYSW3mv3dWY36x3T8CDjSo7frY4uvPylgt632YTx3sg80O25eJ32XV2:                            rando.Int16(),
		IMS2Ow8dMkEOJXlNF1cwCEY4cnR86BHby4lNwjJV:                                                             struct{}{},
		SEp0kW8oostEGaqO5nBG73RuM06aB3wnI7lJu5M2Ir2NaEP4aqrwS:                                                rando.Float32(),
		Ma1e5W6I6NS2820Yqh0uyIE7S28lpbKCyU17JjRnN3ess67X4VuT3KvOsntaTrHj8YImG8OgXmj0I6T:                      rando.Rune(),
		Dt7wD5Tl8KnmwL574PKg3O41cufsDLa6feTYh6k6GdDg6VAQkAmA82p2Vy836xQkuGKs3S7aF7mSlipI0G1q76iQosEpPee84BdJ: rando.Rune(),
		Kd2CfE274tOMYecW72j28D51oIg183e5a8dmSYKlP0OE7uyKr48B:                                                 rando.Bytes(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, MYh4g5vnmfR0Cl0fDWMn13r7KlRB8J6c3la4bN4nD3m8wcT0eOFSUk3t{}, expected)
	require.NotEqual(t, MYh4g5vnmfR0Cl0fDWMn13r7KlRB8J6c3la4bN4nD3m8wcT0eOFSUk3t{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual HUEnJpdeWBcCf7465T267y4A4Inn0l2CMNAkbC
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, HUEnJpdeWBcCf7465T267y4A4Inn0l2CMNAkbC{}, expected)
	require.Equal(t, HUEnJpdeWBcCf7465T267y4A4Inn0l2CMNAkbC{}, actual)

	actual = HUEnJpdeWBcCf7465T267y4A4Inn0l2CMNAkbC{
		G126mt0ekf63prSv4WK7x4lpS3wcTD22xt2AsV0E0c6h5Ehs2t63G3wgnLANj2W: rando.Float32(),
		GWSY14a557y23WjBRO6W2av6nsB11Tf50F0BjJADFDn1GjyCqAoL:            rando.Uint64(),
		NtAcJ62mU488cD2Uu2MQtrPvxq6A45N:                                 struct{}{},
		SY:                                                              rando.Float64(),
		W5Poi:                                                           rando.Byte(),
		PdM25DDXmn:                                                      rando.Uint16(),
		PG2M7Op8KbwXxKi82L6AQ4u3e57DR4HVNG47jVH85bOrD80ry5: rando.Uint32(),
		X8thJMy:              rando.Uint8(),
		VBQIhLG4TrDXNJH0ES0V: rando.Int16(),
		K5rb6RG6P547XV8e31O7Ka1P276d5xWScWr4DG3gAs08s:                                       rando.Int32(),
		RdaIa6ATpho1YK46M80i0UYVIX76U252AG3HP35x6843Y8dHdWkb3q7mEY2u:                        rando.Uint16(),
		NK84h07ce5wc367Lc2h8d6Xd1w4nCfwD5om85itKj5ce5AV2IW3YQxWfA0jmklnH7x811L25wvpuq61J:    rando.Int16(),
		YUDpsW3A5154c3nh04Oh215bGgbI02xPyGJWBtb7448M50pau8Sr7T1erv5m8wqxD:                   rando.Uint16(),
		Gob54P7Is0iaFplWW6g0YW1Fpo1cBfXS08kreXK1v:                                           rando.Int64(),
		I0gftd3vV60cUpA15qMk7BbBqr5ovc58aRobvUg54r4VxS31x5sGgDxOTW8r7BS17nWXpeaR6tr21pgh:    rando.Int8(),
		DhtO30miJ1CigNaB50PbhY1H0IkW3LyBTw7C2PgIK4mJhJ83L0:                                  rando.Int16(),
		ITRHJ1J6rBi41ib22HX6R475wU61Ywo77MeVseKE6P6Y2a5co8of8efqdyEGp36HsW177nup6So5y4:      rando.Uint16(),
		HPX7jtH0P451o85jcRxVhcOsg6Pr:                                                        rando.Int8(),
		PCYS4t3t7Lnw2NUjQ1DN2j064q56AL3kpiy1A1d2s021mba64jQ24n0cB65al3:                      rando.Int(),
		ISTc2Af1eov7i7w7234iD2riYPgR8M7uU5Dwm67jMY87Ws503pm11LCCHOjB5xHNCg1xQuOb2:           rando.Float64(),
		Aa6rIWR10cRqYL82ptVrKiSg465M6xGd05Mgq5xO5f1lqq00W722PM48gJNBa8F16xwFr22rb4YsqTF7sne: rando.Float64(),
		AY6uYkOfhQa152y4g3tSLcPPPcJNUtToJ620h75n6g3CKd04213Poja0Sq4MG3d58y2K4WU0E15rd8AT:    rando.Int64(),
		WYw73A2hJ5lBr866STr3537548NmFSji7:                                                   rando.Int64(),
		BdXBSF0B83P81hM2178EvJvd08dC7R55RDvf27UdJWO:                                         rando.Int16(),
		QWHL6u3170p64R8: struct{}{},
		XDQjt8Ybgr20ph0cox57C23vtc2nBDcMe0E5fEx25kxpCS15ulFw4xIYkOe4qDi4V3Ct3QGBLrwlFknnC: rando.Float32(),
		Wvr8Yw0O4V2T65l10D7r4K7i4wr5h: struct{}{},
		J6F7V6:                        rando.String(),
		NVXiKhqQts3Ligocr4nDW:         rando.Uint16(),
		LaYm65Hej78FT7wUWrb42WOq2gH4Klc01e1WQqRv1Bc1LkHugeKwn8erXnEV7tvOOIOb72vfvDD25EA1dRovpu02K57R7VQ6: rando.Byte(),
		F3PQ8gIcSB58T84ky682inWR58fBxs54L1whWuE5AS65Ao4s1t5Hl657XwTq1BoN:                                 rando.Time(),
		Oy0MPln83S3lVB53TLKqXr4l4PsYN7Ej15qGQqd537ipGrJT3p34545mWS:                                       rando.Bytes(),
		YQb2PAA5uDdutQu7yrLj3ASTL3M7pEQcscmMGS2Doy6l2jOafs1t1X4e82MSU:                                    rando.Int(),
		K1qm1iY7mqd30O721bmawsKWI7FPA0J17cYa0f1X627Ly585han2P:                                            rando.String(),
		T53226NG4617UO0C615rbop44j83MIENQEs62FiD4p4xEi5U1K4m8c642W47u2U701X2QV2XruOOUSI:                  rando.Float32(),
		FeJ42rl3jI2DU8NXOTv4qe7e7aT2Y8o22oPoi3cbfyO8qK8TD73:                                              rando.Bytes(),
		KgxUaTK66N6kwj7Y4q51p807gUuIU8dlmRxeb3pbm1811xtUVf:                                               rando.Uint(),
		V60B261UedL0686e1XmYG8h: rando.Float32(),
		UaLA18b8hT6NN7Q67Iwj38KjhdRCSHERfL8l7sVgN3dcmhRtu68K6M66Y4Gs77rt: rando.Time(),
		U07ppaq2x0c74vB66b1: rando.String(),
		TvEVgWmkKd77vK4jci:  rando.Bool(),
		OXr0vL3T0w0Ddt47J4hbkpwb81cI60L23nwyaTHy74iM41F68MYnk8EIpQxhJ07VpwAoyV00nrra0kUDVo: rando.Int8(),
		Vp1kxF3g1s4e6W2: rando.Rune(),
		Ke7vfch7:        rando.Uint16(),
		Wa5K2552bSnB6123K46C8jpp22wIl73EOT03wMO2m0eP7513XFoVuNDjxLrcf471A2Id50Y07I1H43Di6eo8oC03: rando.Bytes(),
		NeG7sAlaNC3w0DCUUwl0: rando.Uint16(),
		Yd2sqKoo:             rando.Uint32(),
		C4EvAraDW88Ae541f5KqR5l762C4dA7OVHWCy7h6w2KXWfqllip:                                 rando.Bytes(),
		N7RUXC3hq01nnXwYxB41W0FjsfNyF1LviN:                                                  rando.Int32(),
		SwF5F37vHVb1BXiQOEdP3a1e:                                                            rando.Uint(),
		Ju7Ag14p01uTrByUMK6Dj2tp6if13r0qLJEsIKKXR11Ih1Le61TSlMNW8JAT1FEKM0A1wTpojsWDY4vO24T: rando.String(),
		CWG06hS:                                  rando.Int16(),
		U4DUa3V8W4KwARaj6F1RXwBP5KWexdgK37LK:     rando.Int32(),
		Imo4oy0qRIx1nl4adgS:                      rando.Uint64(),
		P7Yh80tP7q3OWpVDSB60d6sP57V5O2DR7pcilaxp: rando.Float32(),
		E66R73Hee5cJ3b4in2fO:                     rando.Byte(),
		RTnDjpF5rOHX34DVH0Kx3L4gqwi3x00ia6LGl41M2rsD7YiwLwOH2R2FUspMs6CuB82sdg:                       rando.Int8(),
		V2Ch07hrD8XBsRvlRBWyye7E114871Ia010sNK507mCfXer1a0rf5KdEM4272XLVY4KvurWbOa3v4WCS8na4w:        rando.Uint16(),
		GXw212r1EhWvUl82IrJu15ruYn8gnJ276Vv43w577e006aWObCs5na1C7Xd5DN0uj2WnIA5FtxY256Pf2C0uAe82WTW8: rando.Uint(),
		PM7QNFhfXO54UFH: rando.Int(),
		FJwFu36jVAXsS4EhQquqv4CfdPk6S753Fx0ooV6jy4Ir2HKHKKT4Kwj5pns37O786T0: rando.Uint16(),
		O6aX4Qjed3:        rando.Int32(),
		A2CCQila4eHdQt38g: rando.Int(),
		K6HxP04pvg3j8vDH8aDWT76rPisiAftWkJ54U11XNmVE5fXxUni4bHXBNg4NLlr57PdEV51J5H08b1n6ij37it:          rando.Byte(),
		IXV6p7JEwYYxR16GP55X27NQjEa486Fe3r48t6U08U6bmx61jgHCY2IJQU518pvlB4x6BYMkVcWDkOYmya5W587dP0TE3tw: rando.Time(),
		W2EO3802d72Q0EIiiaCnOeq6Ij16I3R58KTwh:                                                           rando.Int32(),
		FO2b6o8V6ub10hN1q1DtbQwiYwJrHqThTDniM5ym6AR1Xat877l6RyJ:                                         rando.Rune(),
		Yhkke0s356R5U61Hn2o0Hq2g80qQllO44bQsD156Ob5oBgY4LmEXchD0On1va5ItVq41UMX5El5M63lR4kJR30x0p:       rando.Uint16(),
		SRjyu46667ria45qeQm62ClW85vMgpD2UP43C544y675hD386Mocgow4l8jsHu2M771Y4R3V:                        rando.Uint8(),
		P2OC82hiO70N3I7BEW0ek5VyB142Pq7Psu4eI70xF434tf:                                                  rando.String(),
		EepjkP87kiwkf208w0oyVm58eBxV03PGP7cR8Qkql4s8fx5MAsAnc2EX576R1IR0eEoY40UX0kktMr0yoOhWJt2uMEP:     rando.Bool(),
		S5660: rando.Uint64(),
		L8N318VeGSiwKWo775Cy5jO7m8K0DN06t8l4UgAmTsrQ128DQ302gw: struct{}{},
		ET840iKNawa1iy6SWxNT701aHrKtqsVG7mLnlEf76I6Q70q:        rando.Uint8(),
		EcnM3nwbhA8G61kCLW:                 rando.Uint32(),
		H8CTn1lYwc8YH5qj5Tj8L1s:            rando.Int64(),
		BMT7QYrph2yT2ipQ0vl73y3rq4c7C2dUkr: rando.Bytes(),
		PV23jGwBinbxHPkO6W5TW65AEpF83nl1lcJ66713YBBhpWifyxG6PWyJa7qH1dQP7PGYFu82k860D7EVu: rando.Uint8(),
		HBepPbN0rew4Afo4oEBm2vkLP8bUSP2OKywoa3K5N8uUk7POF5Wm8rgbw5TKO3o3My1UXQyL03Bj4R23:  struct{}{},
		Fq2t0: rando.Bytes(),
		E6gJB12oBTVM1aTmFScW3eW01gH3sJ2QOeoWSL6Yp: rando.Int8(),
		DF2FjhNx:          rando.Uint64(),
		UJWux2TARIVgN:     struct{}{},
		U3F5Dx26kc7c6UyLf: rando.Byte(),
		WKKf584CH0LxFWtePGT2bGo20Whru7pTSK36p60272nME5j5v5: rando.Int64(),
		RMFGbE5VfCIGU3D:                      rando.Uint32(),
		Di43v27R16Q5ls4NGaGt4v2xV3pY0q3WAeTJ: rando.Float32(),
		Hj7HaJQ1fJ380euiJ7:                   struct{}{},
		UB6BQ08jJ2aECLM4cqdB35Sw27aQ3H7u6n0G741L03qb3HFT55Nm1Q6JyN70CJ4W: rando.String(),
		FmE0GYdUFc45YI3pOVQ4o2F4C1Oh6S8jYWfxJ7b1qG4vs0KOuiU5:             rando.Bytes(),
		TnHXimlqg8m1LwB24UwbbMCNC5D2Yv520Wbr35CRSr7:                      rando.Uint32(),
		YJIWtKj4OB6SEPd44: rando.Int(),
		Fc8R313AU32dncvO3X2UHYYdfWS452yUPImo4b6dn5ka2mI0qGh252JG7D57VwruefLYuwl41GeRwI8Oxt71jgXckL2og0nny5: rando.Uint32(),
		BSnQ177Y46It5hRI0Ls77VPtPftQ3R12tmH2:                                                     rando.String(),
		HA4pILnTlUG0E0o568agIVNLLoJqpAe0:                                                         rando.Uint(),
		FU5NvAjBf27ry3bmE1Qd84doE4sI652Ix8g5kgAj7evQ1FJkl8uQNSKYQmN3Pbtl67EU8u82Vs7URBR8Yy8:      rando.Uint16(),
		SGS2G23djBG8t4m2b0h10dhuFp8257Q6ptfNsifohyjN2vc6h4B:                                      rando.Float32(),
		D5ga03e2P8sk70j2X0qyq3SiH417GQ2Cx5p:                                                      rando.Uint16(),
		XcYcXj3c3X5H3XeFC5hPSdw4x21854H31nP54P0x6LiaG5BFy2CT0RWo0oQpV34lcHdoXOIBQQjN60sl1IN3DQ30: struct{}{},
		NP1CJbHo01E5vw6xn51phlNNVn1MHW33aDlkAgaH6q738v:                                           rando.Float64(),
		MoJ5VH76Gqq7X05A4w7pjGE2LeJ80x0guGKNdWi4hv:                                               rando.Uint32(),
		BR3: rando.Time(),
		G3dMX37XYLwhrGy6kVs8r54uKl52766R8waSS83C3sEyF6n4UqCR7Dx0p: rando.String(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, HUEnJpdeWBcCf7465T267y4A4Inn0l2CMNAkbC{}, expected)
	require.NotEqual(t, HUEnJpdeWBcCf7465T267y4A4Inn0l2CMNAkbC{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual AMly01oki2JC2n77UJT22AVlo2VAiUh06W43tGk43P7th6t6MSj3ejR5PrL103BBIN1ji
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, AMly01oki2JC2n77UJT22AVlo2VAiUh06W43tGk43P7th6t6MSj3ejR5PrL103BBIN1ji{}, expected)
	require.Equal(t, AMly01oki2JC2n77UJT22AVlo2VAiUh06W43tGk43P7th6t6MSj3ejR5PrL103BBIN1ji{}, actual)

	actual = AMly01oki2JC2n77UJT22AVlo2VAiUh06W43tGk43P7th6t6MSj3ejR5PrL103BBIN1ji{
		Kc3xL2bowR7TwOLpKkDMP44WlLgB5Rh5Ro82kAob3knH34:                                         rando.Bytes(),
		S3KUyIH2uGr1YN27tQLcv2eAkEG1Dt4Y8hs2Ia:                                                 rando.Float64(),
		F3ft3g2uB7YU526LV0c0Oeq0Gf4YRaUB:                                                       rando.Int32(),
		C33lxaDuW2A6ybtWKj2TF5y0iDadC802S7RU5QD20WJ28iUW3P77s60eEb81Xi:                         rando.Int(),
		DSuEeV4220rRXcqTr5w3bf48:                                                               rando.Uint(),
		RqwPs577d5lTDAQO171A0NwWB1L4cev:                                                        rando.Byte(),
		YCiS0s6Jgw2bWcHF685JKDC6Pi8N45aliRgAgcdtePj23p4w3OeAXvO78W6fb:                          rando.Uint8(),
		U1CY87y2437vhKnBn6I3w527L05:                                                            rando.Uint8(),
		YU0B53F03KOQ3xJtAubkm4rsmJHHG2o1MD1L5JamF3oV8y3rJ727He2gukuutQMgqmOqWwFAx4gWv14k5OojCt: struct{}{},
		Hn1Qexo2s4C4Td5e7XJFxpi2AApQqe5mcSBaYOlrs3L5l8MEB8418EwUB3HSn:                          rando.Int16(),
		Ue0J8g2JhAeC0I:              rando.Uint8(),
		Rk10bi2iS60caL5S5H33LCp7Y1C: rando.Uint64(),
		N7U00Jq7M4rb:                rando.Float32(),
		Mu7L1fA5uQCm6fpEr38j7sL7C6i0Kh144G4mdhPTT77s62c7q3dDY2OKVwkeSM84D6tP6cECnc6BDKCtCb6DdX3N27D: rando.Uint(),
		UUNCBYqg0t5dAmgj2yv1f82:                     rando.Int64(),
		O4rMNxf3aO5IJWc6y2kkp:                       rando.Int64(),
		Ru61x3A3uvCu2762uCvK8d0V5rSR67s56Q3gp46xxfp: rando.Uint32(),
		YBEW261AWb7VYuPT31nJ30LbCPg7y8YdXtj1sUN5q8rfxsNn1s024e6E17cfN21vng81mh1Y5jrD1QgHgRs7uTSw5h: rando.Uint8(),
		SV7844lBSVUnHGIkv15V8LW8s6u4:                                    rando.Uint(),
		T05AXG1G6h42vGYicYHni7hSo:                                       rando.Uint8(),
		GWk0JT2GFoNPBwHS3WN151E6tcOfe3aNNxH0jxNQm36lVbQ8n8Iam400t4dJurQ: rando.Uint32(),
		J77kvKtw5KjuKE6iXmk45R8eLXIEdYPKWM4yj34hKN3H8L3:                 rando.Float32(),
		NbF3sI6ScBuembX1mg3Tm0iGh2mC26R1M1FabP3rOc31M3Nmgpf7:            rando.Byte(),
		UfKf3EwHDcf65g4Fs6aAdRg12lTrCi4efXaQ0M80KO0x17HUOTxGJm14IBf5:    rando.Int32(),
		N: rando.Float32(),
		PPC7AtPx6SNQJK0WbEMPSau50v5Ll7VhlWhSKpnS2Y0U5dc1s1: struct{}{},
		A5lQ: rando.Uint8(),
		FBK2Kv7Ig0qcgR8iQi0OY6upi1yBwxv6XisrWCUWTPa05mMI7bh2qg3WBSLN5mr1A:                                rando.Float32(),
		H5GV5iLAgebIu7g4o2SKCc4myHxL3FeKdQ7JUopgV86fl3n8h6bbhw5KAiA6KgO2Jojxm7:                           rando.Time(),
		Bwf7Y3J7ot17wch11036jO53W2XtoJ50QYy11s8wWQkuk2JnOb31wd1WHLH3VNSywJ7J7i2Lim0cGJKRJ6I6l5M:          rando.Uint(),
		VcK0X26NlJuL85au66Ov66Kk2aO225tv0I1cAY3m2WVt087vVMm86MtTCSgj802OV1Fyxy1R2Ed4GhGgt1lgks:           rando.Uint64(),
		LVUcb8km4Wvq7d0Mi7ugkkvE2dS13sJL0Vchw24g6jIK5ETr4WrgT1DPb8e7xp:                                   rando.Uint64(),
		E6g5vtTKU5c03cO4EwArxLgqh8pnuRnthQiT6UW6187i35bj6u1K6851sk2q8sN3Q5c8B3GORHGXSXi8MUYT8:            rando.Int32(),
		SM45EUMsP7c0BqSu05OwiHlKUx4wj56222mrA53gP2M5NHB80:                                                rando.Int64(),
		IW0B1ASu1777OmeqCcPQ45lC5YX2iqq:                                                                  rando.String(),
		IGYxbTLvrFlab41veBQj15qTX6uy7S:                                                                   rando.Int64(),
		O7Ud4GgACsfhnl7Q7uQ2K24yw4yCIuqrT36qge32RFSRjmmj38RFRah5gcQpF7wcx8717WKGwLQ04RMO70L1j67e58C8uO4:  struct{}{},
		OfX04QBY2XQI8npCSaEYGTYtVb1mdP6tlg05AO3mq422RT8jDi2FRqN5FF3Y5cF1OK1y2yW1Y8JK05Hk1CX87xxOL2F2iJT1: rando.Time(),
		F20en5: rando.Rune(),
		SmD5u3Dp3u1YymhIc58r13ho80s2q56q1tR335181AYgRLt0TRM5Uxd5UM0Bs4xUA0KB4L28Q0HRs0wJgr1lYE8P6U00ca57Rdp: rando.Int(),
		BAh65bIl5I44gOY47v4BrAH0Obql6fjCEd1r7S5xXumU1RWm6MTN8D6oBESvnG4IM1q3GieGLB30fO:                      rando.Int16(),
		I5oWnf6tyS2mk4v3al1inMxpN7Wj08x7M15usu3xt63bM0epMffOCixf2Ogi3R351gO0:                                rando.Byte(),
		TAp: rando.Rune(),
		U82Uao8X38P7YLkt4d8qmyl8dNBtG5HxD0DWt161SeQWT6mE:     rando.String(),
		IT2FPQXXca0Fu4BtXa1GFXV837EPBSI83e7la84TTd4g0e867m7u: rando.Int16(),
		Uq23yRI4W6:                        rando.Uint8(),
		B46TTl2OIN0mgMdUcD8p21rsFa04Cm4sv: rando.Bytes(),
		Vb8BTfmBH262aGW8cBclN2VVWBHE45eEfAghn5kEd6lF61KoMwjHKd8hmiYl8gCWuLt6ExEs7VF680sv4:          rando.Uint32(),
		G55T53ie6mm1qR4S8j1k2np8LCc841sW2b7pHI3B64F4ei0kdJae67FRU8Arx6Ra:                           rando.Int32(),
		XUvg4evc1auQ0rt1XPL8xLVt38qqiKjN0Tn5iy0u411e1h176PW1pIt3gU28jChT7Ub1VN2LJS606Ev3q78KO44wB8: rando.Time(),
		XhvCr720NYe5IxPtWWWW4va:                                 rando.Float64(),
		C81nWkpqm630B8F0hS107UXT5T8:                             rando.Bytes(),
		D2nGeA50e8B7tNxbiG227RV864Qmt71T62mXGMT75cotRqCugu87swn: rando.Int16(),
		UTg7TawC4QO80:                                           rando.Int(),
		Ma2Q0O2Wf7sTR678Pj525jr7h85134nlo4p14d31B7fqFknjskj82fm2GcjL4JjqACIt3nIB:    struct{}{},
		HaTJ5wctBlj0TS1aiiDS4w1vJGIMR51v4T47o5kxwQblt8rArOuuKd08X15x5444b6d2647J8f3: rando.Uint32(),
		MO5n60vWB312SWHISn2wd4qQoer1y337f55axYs3Rc74h4k63yYxN42XtbV1h6mt208u64:      rando.Int(),
		Fb173eCxU17I: rando.Uint64(),
		E5wqcATKLnLHQ60rdrWDJ1yQ4oI16rxrN68FVMrN5RV1bjaLL6p7kNi6C4xLuyTPxb0cr7d7I287: rando.Int8(),
		Ta5536bV3TH3nb7tGbBAm418ey4sBBVe2RdT6K2X8nq2eah2C3G55Y8:                      rando.Int32(),
		Wl12l4: rando.Uint16(),
		K3EJ1UrKYkbWAauAtqwSN8AuP1rPdIEf2WJ886V006o:                                                       rando.Bool(),
		V5JvbOv0107N8qM2m5EPuM3rTIEPGSWFpkGLt4tvRWAk66181NKn7u3C1jak6A4YRc0S6hGQ8l7:                       rando.Bool(),
		J7FNI4821OH2VMBmjG3qNFy2SSI851nx17Q41RuWW8C2d:                                                     rando.Rune(),
		XAoxjdPDO8N28i1Kmwa547xJQS37al1Oy5OIYB37PmMVt6PXUTxvGQTBB38G4XpL4rS07gdhlTPm6mtC0m5JXmt8:          rando.Uint(),
		IHHNryTqdNyOEXQD1P3O6vC12oeE6j2mml8dK882ok21pWNyW328DRbw1xCb5aFjfEwrh2M1WU5H23Ol45Aq0:             rando.Uint64(),
		Pe7BR1E2mr7BWA5y18HgeVEIsxKl0PPSjs1RTYEr03x7pBRbM64470Al41V76HMOc4P7runL:                          rando.Uint32(),
		Wp7a6qr86735t1tTv338Snvr45jdCBJ45FpN0FFQf4tokXJw:                                                  rando.Uint16(),
		W2xc78gY3V3xUESRUk44pxX16Ti002wW308Gk24HfgCckt70P3wmjmVwheea4w5b42Ju2m7S2AvO8thp507341yC7UHaIT0NB: rando.Byte(),
		UNW8I: rando.Bool(),
		EYW6cXT3PdjjbV5k52YMo7e7yY5FMDPuoY4GuUF11:                                                          rando.Uint8(),
		T0NDiW7asNiR5A655ssSe4box32Roc7NL8H420pYjPS8a23jhIINjo70bI7Sew14TD84F8M2IM6t0IaNoMBx:               rando.Uint32(),
		Nhjj038NJ04vqJDKgqkJ4wq6vlJAsD27hoPA7P:                                                             rando.Time(),
		Lv6uxvmRy6k50UkufS0i3LAC622511QkNF3pDmUCGpH6P6teD3oMJS0HkL5OLBVmjelU4O56f45ypWqy6UTLXD3x7GY8Q8MPRx: rando.Bool(),
		E161pw2FDkee46cYa85y2Y4X8:                                                                          rando.String(),
		X6H47D6k02LwQQ08:                                                                                   rando.Uint8(),
		Gn01Up7J08XvEv6G5ure51GV0u8c7DGj0n7phy5U0clGd4dC2Cd3P5FNq1:                                         rando.Byte(),
		EFUQx0FBGN345SBPrwf42L82MGov7DkUBQk:                                                                rando.Byte(),
		KLCcq5F86SI53yi6454X1Q5X0OIDkl1aEjiIB1nkn45P10j0r5qb5y:                                             rando.Time(),
		Ln7qxIc1xGosRK8M62I5j80IJN7H1oB3KpkmN0LqqDlT87h:                                                    rando.Uint(),
		F2Gfck8ev41X08qQE3wVKJ0weRiYTjgB5eYD0Ho4sx7jbMfs1Ftb42c52lYqtN22e4k6h617Ai:                         rando.Uint8(),
		P0BfqN1q: rando.Int16(),
		I3SWg4CV77750OdfDI4Kt4An5N83Al66gd6yigS0q16X8w7hC: rando.Int16(),
		UvSK:                          rando.String(),
		XDUv6BF2hW1u588d76C6qS0F64Q5c: rando.Bool(),
		F6w1voK3gWJ7wphds0Li8dJQ52W503DLR2lfRup0XxJ0Jf5VBkLN4F56Q: rando.String(),
		FE0118sftca22Uv:            rando.Int8(),
		F2iof7ha2MhsYvU3MQhB6ivy6X: rando.Bool(),
		X2QRx455Wp2:                rando.Uint(),
		YE348JHYKM4IewHYsW223u6l8qUOerISlv61J67T0khhR0dhXm68l24QTs018Hb863Fe0HRo4o0i1eRLfWQ: rando.Int16(),
		C56q26r:                  rando.Bytes(),
		I4n0klX453ugOp6teq651uW4: rando.Rune(),
		C7gx8vVuco76psH10yhRE6nL45rr1b4fvWpbIN33411t02iP5DP4J3huU7Xl4B8w2t0h7471H5iW2jyxcD36UN58rsgN07N0770a: rando.Uint(),
		Y3Ox80LRyRuLCO3P2yJ6I:                            rando.Float64(),
		DjE0g7jxOBuhNhf51pXvj5LhU6b6qsP6MX5Nv81EBy:       rando.Rune(),
		Brnni3PeEi2XU10cyr6j0ppdG2a05BdIGlBpM687A1P66c2t: rando.Uint32(),
		URL7gIQ16gfSyqcw51FQaE3wfRgTe5CE2uU0STpV7D:       rando.Uint32(),
		P2J035F8Nk1: rando.Bool(),
		PX06Ty52WFLl3HwggY387niD34HBEp4QTvD37SXFBog7vh02RUMc4OI1OKJ641Ib3r67P: rando.Bool(),
		We4o7Xt5TTm7Q0Xxb8XqTc1xj5LQ1daGEaqlQvJ6dF7d6TT42xIpDh:                rando.Int8(),
		CHaMB7Cw311V418f2W24J0kVYkVe548wfhv08Bk0275QnXbI:                      rando.Int8(),
		PuNeJcTcveW3p5: rando.Int8(),
		OW6r70Djqj:     rando.Uint(),
		J86wCAlDW41w0cl4FjCU08471hp2Q5Eo2Dw4Iwj8E2ySsy26N5yF30GUw0:                                   rando.Time(),
		Tn0RerY8Y81KFf1w7o82KukpV68736TLVKJP66s3Y1V54qT86MXtoWNUa2fudidETwHgV:                        rando.Uint(),
		EqXn3kdLI15O5GQyr6253kb7VdNv2uC4mwNLfOmR5SdMl5rdtum6vypL3:                                    rando.Uint16(),
		Q7vqKj51Qb84d48JEIvYNU75aWo0qAYSVjXY44j8AC7UD41pl51AR1d264nv2:                                rando.Uint32(),
		JT0Mq1CSXk2rO4t1PMDW0YOiB14ih1pm6FgdvCbG3332i2Kl5c:                                           rando.Byte(),
		FmUJmlMOtNqi4uUR0N1K4o72E064hbj8jX5N1sEx6U67l680FymAOtc7StlcwoAdyD2TeJ12arXQg:                rando.Float32(),
		BA0rA3ud8m5KCYi78NeWo7c7NSHQ3PdEMIPky2QQ3v6S8Ur6uqy1OK:                                       rando.Int32(),
		BXBVF4Q4iNdqfgAeaqK64s6t6Bav8ojL6KOwyY1052sL6bT1u04OSE:                                       rando.Time(),
		O58y041UjLx2ew73u70WJcqFKrC654XTLsmb4EcM0aitN05VUyjp1kW15i0736h0pjco10OpgY:                   rando.Uint8(),
		LO82V5X1RHx52i61e51fp6x2bbBgSBo70x34fuW5EtcR5GUlu4TP1B3QD472K8y8QFoa5YGg7425VTCka7A0rGNUyHU6: rando.Int16(),
		A6SPqF3gTP3b7kAHKcaQCjV1tT7pCTJOlP7Jtd5YWGP78Diw7N5F26kNPS:                                   rando.Bool(),
		VT4F711YMnK5mS3RNF531Qpj8Kpo68q7nTp560YRLMsrAdD87o6U:                                         rando.Int64(),
		Ex3gdR0qi8fJJ0U86e0V1312cPMIldHI0cOsUrv4Y2NkAwtJ4TnHBS36nt1Hytn8il748is3xgE1i011j30A0sq2:     rando.Float32(),
		U6MGiW6FA2eKQwok4h4ED45D4x8PX3X8BUeBy2GYHGvGvgm67Us8g03:                                      rando.Time(),
		TLb58mN:                         rando.Int(),
		VE00cVoA77jj656Vy7oj643l2pA:     rando.Uint(),
		Uar0CX3o5u3WfqE7o6T3QSq2fMeM4QF: rando.Uint32(),
		LYLaR8VdnVR:                     rando.Time(),
		Rnuds8B23000PDI6wyYK1nl12VYXV1a4mtyRRcRcsO24C5yB:                       rando.Int64(),
		OAG28SLLnNl77d27AKMI58db7M3LR65jX07870MAfIi24C2I3Go8cG2751b1H271H42Qp2: rando.Uint32(),
		TGa26aGI2D3tjK4f72LR007N35qKNY4XvFpw5:                                  rando.Uint16(),
		HoB7nFxQHUY652NfN3D7g826sU8PN4K6xev42oO78OxmgFb4D:                      rando.Byte(),
		U2ii03A226243h1p18Wg7oUPc4sI:                                           rando.Bool(),
		G13VEkg4508ej7kiKJ3XKjDT8t6P55T711miJKSR0H27705r77BX5242YatijJe6X:      rando.Uint64(),
		H82PSTYkUKHQ: rando.Bool(),
		Q1yYE381754DtTlMHg2UBKRd0hWe6Xp1nnrVuVc38WB2WDlTx:          rando.Int16(),
		N22e881N0OcAnCFP7gD3o1lu66pnU07f8U8fhGjAE5fsm617Wknd20aBJC: rando.Uint8(),
		JjyJi8iM45edK2A: struct{}{},
		I535RC1JeWO4pDRE81cHuoubcU6aUc1Hogq04NdIx8dTYo5vjq8UtR50F2J02v63v0qD6l8634PmdPW4yY:                  rando.Int16(),
		X56gJ8sLWa2wbeWqSI7qK17563i6yh3O6fLjiQMh:                                                            rando.Bool(),
		Om40t0dTFg4XxBk247p8v26OYHbQKDogx04o8qEmNMdbOP:                                                      rando.Int(),
		V4oAFTJfic0g5r7Aj864wCtu7lMblg0L7i42xUHfB3M4L4EejhqpxusAAASj7Y0b87xlUB0l6485C0n5iL832aknd28pG7L:     rando.Time(),
		NEFmR3GYcsX67Uai11Ybyo4KyPhTa25E17EUEX36agNChN2j2k5064G830g08ta5Oje4rskK0xToMvYXWQYjCP21wFQF82xEbRw: rando.Bytes(),
		GsXO60OQk2J1D1Fo5wDU86as6fmGC3yR21UgU83YOE2ROf3Sd0XpN2BCjr6R6WLtW3h2hUunD2e7gOURL3ieJoq:             rando.Float32(),
		BUL7SnjImLQ: rando.Uint8(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, AMly01oki2JC2n77UJT22AVlo2VAiUh06W43tGk43P7th6t6MSj3ejR5PrL103BBIN1ji{}, expected)
	require.NotEqual(t, AMly01oki2JC2n77UJT22AVlo2VAiUh06W43tGk43P7th6t6MSj3ejR5PrL103BBIN1ji{}, actual)

	assert.Equal(t, expected.N, actual.N)

	assert.Equal(t, expected.Kc3xL2bowR7TwOLpKkDMP44WlLgB5Rh5Ro82kAob3knH34, actual.Kc3xL2bowR7TwOLpKkDMP44WlLgB5Rh5Ro82kAob3knH34)
	assert.Equal(t, expected.S3KUyIH2uGr1YN27tQLcv2eAkEG1Dt4Y8hs2Ia, actual.S3KUyIH2uGr1YN27tQLcv2eAkEG1Dt4Y8hs2Ia)
	assert.Equal(t, expected.F3ft3g2uB7YU526LV0c0Oeq0Gf4YRaUB, actual.F3ft3g2uB7YU526LV0c0Oeq0Gf4YRaUB)
	assert.Equal(t, expected.C33lxaDuW2A6ybtWKj2TF5y0iDadC802S7RU5QD20WJ28iUW3P77s60eEb81Xi, actual.C33lxaDuW2A6ybtWKj2TF5y0iDadC802S7RU5QD20WJ28iUW3P77s60eEb81Xi)
	assert.Equal(t, expected.DSuEeV4220rRXcqTr5w3bf48, actual.DSuEeV4220rRXcqTr5w3bf48)
	assert.Equal(t, expected.RqwPs577d5lTDAQO171A0NwWB1L4cev, actual.RqwPs577d5lTDAQO171A0NwWB1L4cev)
	assert.Equal(t, expected.YCiS0s6Jgw2bWcHF685JKDC6Pi8N45aliRgAgcdtePj23p4w3OeAXvO78W6fb, actual.YCiS0s6Jgw2bWcHF685JKDC6Pi8N45aliRgAgcdtePj23p4w3OeAXvO78W6fb)
	assert.Equal(t, expected.U1CY87y2437vhKnBn6I3w527L05, actual.U1CY87y2437vhKnBn6I3w527L05)
	assert.Equal(t, expected.YU0B53F03KOQ3xJtAubkm4rsmJHHG2o1MD1L5JamF3oV8y3rJ727He2gukuutQMgqmOqWwFAx4gWv14k5OojCt, actual.YU0B53F03KOQ3xJtAubkm4rsmJHHG2o1MD1L5JamF3oV8y3rJ727He2gukuutQMgqmOqWwFAx4gWv14k5OojCt)
	assert.Equal(t, expected.Hn1Qexo2s4C4Td5e7XJFxpi2AApQqe5mcSBaYOlrs3L5l8MEB8418EwUB3HSn, actual.Hn1Qexo2s4C4Td5e7XJFxpi2AApQqe5mcSBaYOlrs3L5l8MEB8418EwUB3HSn)
	assert.Equal(t, expected.Ue0J8g2JhAeC0I, actual.Ue0J8g2JhAeC0I)
	assert.Equal(t, expected.Rk10bi2iS60caL5S5H33LCp7Y1C, actual.Rk10bi2iS60caL5S5H33LCp7Y1C)
	assert.Equal(t, expected.N7U00Jq7M4rb, actual.N7U00Jq7M4rb)
	assert.Equal(t, expected.Mu7L1fA5uQCm6fpEr38j7sL7C6i0Kh144G4mdhPTT77s62c7q3dDY2OKVwkeSM84D6tP6cECnc6BDKCtCb6DdX3N27D, actual.Mu7L1fA5uQCm6fpEr38j7sL7C6i0Kh144G4mdhPTT77s62c7q3dDY2OKVwkeSM84D6tP6cECnc6BDKCtCb6DdX3N27D)
	assert.Equal(t, expected.UUNCBYqg0t5dAmgj2yv1f82, actual.UUNCBYqg0t5dAmgj2yv1f82)
	assert.Equal(t, expected.O4rMNxf3aO5IJWc6y2kkp, actual.O4rMNxf3aO5IJWc6y2kkp)
	assert.Equal(t, expected.Ru61x3A3uvCu2762uCvK8d0V5rSR67s56Q3gp46xxfp, actual.Ru61x3A3uvCu2762uCvK8d0V5rSR67s56Q3gp46xxfp)
	assert.Equal(t, expected.YBEW261AWb7VYuPT31nJ30LbCPg7y8YdXtj1sUN5q8rfxsNn1s024e6E17cfN21vng81mh1Y5jrD1QgHgRs7uTSw5h, actual.YBEW261AWb7VYuPT31nJ30LbCPg7y8YdXtj1sUN5q8rfxsNn1s024e6E17cfN21vng81mh1Y5jrD1QgHgRs7uTSw5h)
	assert.Equal(t, expected.SV7844lBSVUnHGIkv15V8LW8s6u4, actual.SV7844lBSVUnHGIkv15V8LW8s6u4)
	assert.Equal(t, expected.T05AXG1G6h42vGYicYHni7hSo, actual.T05AXG1G6h42vGYicYHni7hSo)
	assert.Equal(t, expected.GWk0JT2GFoNPBwHS3WN151E6tcOfe3aNNxH0jxNQm36lVbQ8n8Iam400t4dJurQ, actual.GWk0JT2GFoNPBwHS3WN151E6tcOfe3aNNxH0jxNQm36lVbQ8n8Iam400t4dJurQ)
	assert.Equal(t, expected.J77kvKtw5KjuKE6iXmk45R8eLXIEdYPKWM4yj34hKN3H8L3, actual.J77kvKtw5KjuKE6iXmk45R8eLXIEdYPKWM4yj34hKN3H8L3)
	assert.Equal(t, expected.NbF3sI6ScBuembX1mg3Tm0iGh2mC26R1M1FabP3rOc31M3Nmgpf7, actual.NbF3sI6ScBuembX1mg3Tm0iGh2mC26R1M1FabP3rOc31M3Nmgpf7)
	assert.Equal(t, expected.UfKf3EwHDcf65g4Fs6aAdRg12lTrCi4efXaQ0M80KO0x17HUOTxGJm14IBf5, actual.UfKf3EwHDcf65g4Fs6aAdRg12lTrCi4efXaQ0M80KO0x17HUOTxGJm14IBf5)
	assert.Equal(t, expected.N, actual.N)
	assert.Equal(t, expected.PPC7AtPx6SNQJK0WbEMPSau50v5Ll7VhlWhSKpnS2Y0U5dc1s1, actual.PPC7AtPx6SNQJK0WbEMPSau50v5Ll7VhlWhSKpnS2Y0U5dc1s1)
	assert.Equal(t, expected.A5lQ, actual.A5lQ)
	assert.Equal(t, expected.FBK2Kv7Ig0qcgR8iQi0OY6upi1yBwxv6XisrWCUWTPa05mMI7bh2qg3WBSLN5mr1A, actual.FBK2Kv7Ig0qcgR8iQi0OY6upi1yBwxv6XisrWCUWTPa05mMI7bh2qg3WBSLN5mr1A)
	assert.Equal(t, expected.H5GV5iLAgebIu7g4o2SKCc4myHxL3FeKdQ7JUopgV86fl3n8h6bbhw5KAiA6KgO2Jojxm7, actual.H5GV5iLAgebIu7g4o2SKCc4myHxL3FeKdQ7JUopgV86fl3n8h6bbhw5KAiA6KgO2Jojxm7)
	assert.Equal(t, expected.Bwf7Y3J7ot17wch11036jO53W2XtoJ50QYy11s8wWQkuk2JnOb31wd1WHLH3VNSywJ7J7i2Lim0cGJKRJ6I6l5M, actual.Bwf7Y3J7ot17wch11036jO53W2XtoJ50QYy11s8wWQkuk2JnOb31wd1WHLH3VNSywJ7J7i2Lim0cGJKRJ6I6l5M)
	assert.Equal(t, expected.VcK0X26NlJuL85au66Ov66Kk2aO225tv0I1cAY3m2WVt087vVMm86MtTCSgj802OV1Fyxy1R2Ed4GhGgt1lgks, actual.VcK0X26NlJuL85au66Ov66Kk2aO225tv0I1cAY3m2WVt087vVMm86MtTCSgj802OV1Fyxy1R2Ed4GhGgt1lgks)
	assert.Equal(t, expected.LVUcb8km4Wvq7d0Mi7ugkkvE2dS13sJL0Vchw24g6jIK5ETr4WrgT1DPb8e7xp, actual.LVUcb8km4Wvq7d0Mi7ugkkvE2dS13sJL0Vchw24g6jIK5ETr4WrgT1DPb8e7xp)
	assert.Equal(t, expected.E6g5vtTKU5c03cO4EwArxLgqh8pnuRnthQiT6UW6187i35bj6u1K6851sk2q8sN3Q5c8B3GORHGXSXi8MUYT8, actual.E6g5vtTKU5c03cO4EwArxLgqh8pnuRnthQiT6UW6187i35bj6u1K6851sk2q8sN3Q5c8B3GORHGXSXi8MUYT8)
	assert.Equal(t, expected.SM45EUMsP7c0BqSu05OwiHlKUx4wj56222mrA53gP2M5NHB80, actual.SM45EUMsP7c0BqSu05OwiHlKUx4wj56222mrA53gP2M5NHB80)
	assert.Equal(t, expected.IW0B1ASu1777OmeqCcPQ45lC5YX2iqq, actual.IW0B1ASu1777OmeqCcPQ45lC5YX2iqq)
	assert.Equal(t, expected.IGYxbTLvrFlab41veBQj15qTX6uy7S, actual.IGYxbTLvrFlab41veBQj15qTX6uy7S)
	assert.Equal(t, expected.O7Ud4GgACsfhnl7Q7uQ2K24yw4yCIuqrT36qge32RFSRjmmj38RFRah5gcQpF7wcx8717WKGwLQ04RMO70L1j67e58C8uO4, actual.O7Ud4GgACsfhnl7Q7uQ2K24yw4yCIuqrT36qge32RFSRjmmj38RFRah5gcQpF7wcx8717WKGwLQ04RMO70L1j67e58C8uO4)
	assert.Equal(t, expected.OfX04QBY2XQI8npCSaEYGTYtVb1mdP6tlg05AO3mq422RT8jDi2FRqN5FF3Y5cF1OK1y2yW1Y8JK05Hk1CX87xxOL2F2iJT1, actual.OfX04QBY2XQI8npCSaEYGTYtVb1mdP6tlg05AO3mq422RT8jDi2FRqN5FF3Y5cF1OK1y2yW1Y8JK05Hk1CX87xxOL2F2iJT1)
	assert.Equal(t, expected.F20en5, actual.F20en5)
	assert.Equal(t, expected.SmD5u3Dp3u1YymhIc58r13ho80s2q56q1tR335181AYgRLt0TRM5Uxd5UM0Bs4xUA0KB4L28Q0HRs0wJgr1lYE8P6U00ca57Rdp, actual.SmD5u3Dp3u1YymhIc58r13ho80s2q56q1tR335181AYgRLt0TRM5Uxd5UM0Bs4xUA0KB4L28Q0HRs0wJgr1lYE8P6U00ca57Rdp)
	assert.Equal(t, expected.BAh65bIl5I44gOY47v4BrAH0Obql6fjCEd1r7S5xXumU1RWm6MTN8D6oBESvnG4IM1q3GieGLB30fO, actual.BAh65bIl5I44gOY47v4BrAH0Obql6fjCEd1r7S5xXumU1RWm6MTN8D6oBESvnG4IM1q3GieGLB30fO)
	assert.Equal(t, expected.I5oWnf6tyS2mk4v3al1inMxpN7Wj08x7M15usu3xt63bM0epMffOCixf2Ogi3R351gO0, actual.I5oWnf6tyS2mk4v3al1inMxpN7Wj08x7M15usu3xt63bM0epMffOCixf2Ogi3R351gO0)
	assert.Equal(t, expected.TAp, actual.TAp)
	assert.Equal(t, expected.U82Uao8X38P7YLkt4d8qmyl8dNBtG5HxD0DWt161SeQWT6mE, actual.U82Uao8X38P7YLkt4d8qmyl8dNBtG5HxD0DWt161SeQWT6mE)
	assert.Equal(t, expected.IT2FPQXXca0Fu4BtXa1GFXV837EPBSI83e7la84TTd4g0e867m7u, actual.IT2FPQXXca0Fu4BtXa1GFXV837EPBSI83e7la84TTd4g0e867m7u)
	assert.Equal(t, expected.Uq23yRI4W6, actual.Uq23yRI4W6)
	assert.Equal(t, expected.B46TTl2OIN0mgMdUcD8p21rsFa04Cm4sv, actual.B46TTl2OIN0mgMdUcD8p21rsFa04Cm4sv)
	assert.Equal(t, expected.Vb8BTfmBH262aGW8cBclN2VVWBHE45eEfAghn5kEd6lF61KoMwjHKd8hmiYl8gCWuLt6ExEs7VF680sv4, actual.Vb8BTfmBH262aGW8cBclN2VVWBHE45eEfAghn5kEd6lF61KoMwjHKd8hmiYl8gCWuLt6ExEs7VF680sv4)
	assert.Equal(t, expected.G55T53ie6mm1qR4S8j1k2np8LCc841sW2b7pHI3B64F4ei0kdJae67FRU8Arx6Ra, actual.G55T53ie6mm1qR4S8j1k2np8LCc841sW2b7pHI3B64F4ei0kdJae67FRU8Arx6Ra)
	assert.Equal(t, expected.XUvg4evc1auQ0rt1XPL8xLVt38qqiKjN0Tn5iy0u411e1h176PW1pIt3gU28jChT7Ub1VN2LJS606Ev3q78KO44wB8, actual.XUvg4evc1auQ0rt1XPL8xLVt38qqiKjN0Tn5iy0u411e1h176PW1pIt3gU28jChT7Ub1VN2LJS606Ev3q78KO44wB8)
	assert.Equal(t, expected.XhvCr720NYe5IxPtWWWW4va, actual.XhvCr720NYe5IxPtWWWW4va)
	assert.Equal(t, expected.C81nWkpqm630B8F0hS107UXT5T8, actual.C81nWkpqm630B8F0hS107UXT5T8)
	assert.Equal(t, expected.D2nGeA50e8B7tNxbiG227RV864Qmt71T62mXGMT75cotRqCugu87swn, actual.D2nGeA50e8B7tNxbiG227RV864Qmt71T62mXGMT75cotRqCugu87swn)
	assert.Equal(t, expected.UTg7TawC4QO80, actual.UTg7TawC4QO80)
	assert.Equal(t, expected.Ma2Q0O2Wf7sTR678Pj525jr7h85134nlo4p14d31B7fqFknjskj82fm2GcjL4JjqACIt3nIB, actual.Ma2Q0O2Wf7sTR678Pj525jr7h85134nlo4p14d31B7fqFknjskj82fm2GcjL4JjqACIt3nIB)
	assert.Equal(t, expected.HaTJ5wctBlj0TS1aiiDS4w1vJGIMR51v4T47o5kxwQblt8rArOuuKd08X15x5444b6d2647J8f3, actual.HaTJ5wctBlj0TS1aiiDS4w1vJGIMR51v4T47o5kxwQblt8rArOuuKd08X15x5444b6d2647J8f3)
	assert.Equal(t, expected.MO5n60vWB312SWHISn2wd4qQoer1y337f55axYs3Rc74h4k63yYxN42XtbV1h6mt208u64, actual.MO5n60vWB312SWHISn2wd4qQoer1y337f55axYs3Rc74h4k63yYxN42XtbV1h6mt208u64)
	assert.Equal(t, expected.Fb173eCxU17I, actual.Fb173eCxU17I)
	assert.Equal(t, expected.E5wqcATKLnLHQ60rdrWDJ1yQ4oI16rxrN68FVMrN5RV1bjaLL6p7kNi6C4xLuyTPxb0cr7d7I287, actual.E5wqcATKLnLHQ60rdrWDJ1yQ4oI16rxrN68FVMrN5RV1bjaLL6p7kNi6C4xLuyTPxb0cr7d7I287)
	assert.Equal(t, expected.Ta5536bV3TH3nb7tGbBAm418ey4sBBVe2RdT6K2X8nq2eah2C3G55Y8, actual.Ta5536bV3TH3nb7tGbBAm418ey4sBBVe2RdT6K2X8nq2eah2C3G55Y8)
	assert.Equal(t, expected.Wl12l4, actual.Wl12l4)
	assert.Equal(t, expected.K3EJ1UrKYkbWAauAtqwSN8AuP1rPdIEf2WJ886V006o, actual.K3EJ1UrKYkbWAauAtqwSN8AuP1rPdIEf2WJ886V006o)
	assert.Equal(t, expected.V5JvbOv0107N8qM2m5EPuM3rTIEPGSWFpkGLt4tvRWAk66181NKn7u3C1jak6A4YRc0S6hGQ8l7, actual.V5JvbOv0107N8qM2m5EPuM3rTIEPGSWFpkGLt4tvRWAk66181NKn7u3C1jak6A4YRc0S6hGQ8l7)
	assert.Equal(t, expected.J7FNI4821OH2VMBmjG3qNFy2SSI851nx17Q41RuWW8C2d, actual.J7FNI4821OH2VMBmjG3qNFy2SSI851nx17Q41RuWW8C2d)
	assert.Equal(t, expected.XAoxjdPDO8N28i1Kmwa547xJQS37al1Oy5OIYB37PmMVt6PXUTxvGQTBB38G4XpL4rS07gdhlTPm6mtC0m5JXmt8, actual.XAoxjdPDO8N28i1Kmwa547xJQS37al1Oy5OIYB37PmMVt6PXUTxvGQTBB38G4XpL4rS07gdhlTPm6mtC0m5JXmt8)
	assert.Equal(t, expected.IHHNryTqdNyOEXQD1P3O6vC12oeE6j2mml8dK882ok21pWNyW328DRbw1xCb5aFjfEwrh2M1WU5H23Ol45Aq0, actual.IHHNryTqdNyOEXQD1P3O6vC12oeE6j2mml8dK882ok21pWNyW328DRbw1xCb5aFjfEwrh2M1WU5H23Ol45Aq0)
	assert.Equal(t, expected.Pe7BR1E2mr7BWA5y18HgeVEIsxKl0PPSjs1RTYEr03x7pBRbM64470Al41V76HMOc4P7runL, actual.Pe7BR1E2mr7BWA5y18HgeVEIsxKl0PPSjs1RTYEr03x7pBRbM64470Al41V76HMOc4P7runL)
	assert.Equal(t, expected.Wp7a6qr86735t1tTv338Snvr45jdCBJ45FpN0FFQf4tokXJw, actual.Wp7a6qr86735t1tTv338Snvr45jdCBJ45FpN0FFQf4tokXJw)
	assert.Equal(t, expected.W2xc78gY3V3xUESRUk44pxX16Ti002wW308Gk24HfgCckt70P3wmjmVwheea4w5b42Ju2m7S2AvO8thp507341yC7UHaIT0NB, actual.W2xc78gY3V3xUESRUk44pxX16Ti002wW308Gk24HfgCckt70P3wmjmVwheea4w5b42Ju2m7S2AvO8thp507341yC7UHaIT0NB)
	assert.Equal(t, expected.UNW8I, actual.UNW8I)
	assert.Equal(t, expected.EYW6cXT3PdjjbV5k52YMo7e7yY5FMDPuoY4GuUF11, actual.EYW6cXT3PdjjbV5k52YMo7e7yY5FMDPuoY4GuUF11)
	assert.Equal(t, expected.T0NDiW7asNiR5A655ssSe4box32Roc7NL8H420pYjPS8a23jhIINjo70bI7Sew14TD84F8M2IM6t0IaNoMBx, actual.T0NDiW7asNiR5A655ssSe4box32Roc7NL8H420pYjPS8a23jhIINjo70bI7Sew14TD84F8M2IM6t0IaNoMBx)
	assert.Equal(t, expected.Nhjj038NJ04vqJDKgqkJ4wq6vlJAsD27hoPA7P, actual.Nhjj038NJ04vqJDKgqkJ4wq6vlJAsD27hoPA7P)
	assert.Equal(t, expected.Lv6uxvmRy6k50UkufS0i3LAC622511QkNF3pDmUCGpH6P6teD3oMJS0HkL5OLBVmjelU4O56f45ypWqy6UTLXD3x7GY8Q8MPRx, actual.Lv6uxvmRy6k50UkufS0i3LAC622511QkNF3pDmUCGpH6P6teD3oMJS0HkL5OLBVmjelU4O56f45ypWqy6UTLXD3x7GY8Q8MPRx)
	assert.Equal(t, expected.E161pw2FDkee46cYa85y2Y4X8, actual.E161pw2FDkee46cYa85y2Y4X8)
	assert.Equal(t, expected.X6H47D6k02LwQQ08, actual.X6H47D6k02LwQQ08)
	assert.Equal(t, expected.Gn01Up7J08XvEv6G5ure51GV0u8c7DGj0n7phy5U0clGd4dC2Cd3P5FNq1, actual.Gn01Up7J08XvEv6G5ure51GV0u8c7DGj0n7phy5U0clGd4dC2Cd3P5FNq1)
	assert.Equal(t, expected.EFUQx0FBGN345SBPrwf42L82MGov7DkUBQk, actual.EFUQx0FBGN345SBPrwf42L82MGov7DkUBQk)
	assert.Equal(t, expected.KLCcq5F86SI53yi6454X1Q5X0OIDkl1aEjiIB1nkn45P10j0r5qb5y, actual.KLCcq5F86SI53yi6454X1Q5X0OIDkl1aEjiIB1nkn45P10j0r5qb5y)
	assert.Equal(t, expected.Ln7qxIc1xGosRK8M62I5j80IJN7H1oB3KpkmN0LqqDlT87h, actual.Ln7qxIc1xGosRK8M62I5j80IJN7H1oB3KpkmN0LqqDlT87h)
	assert.Equal(t, expected.F2Gfck8ev41X08qQE3wVKJ0weRiYTjgB5eYD0Ho4sx7jbMfs1Ftb42c52lYqtN22e4k6h617Ai, actual.F2Gfck8ev41X08qQE3wVKJ0weRiYTjgB5eYD0Ho4sx7jbMfs1Ftb42c52lYqtN22e4k6h617Ai)
	assert.Equal(t, expected.P0BfqN1q, actual.P0BfqN1q)
	assert.Equal(t, expected.I3SWg4CV77750OdfDI4Kt4An5N83Al66gd6yigS0q16X8w7hC, actual.I3SWg4CV77750OdfDI4Kt4An5N83Al66gd6yigS0q16X8w7hC)
	assert.Equal(t, expected.UvSK, actual.UvSK)
	assert.Equal(t, expected.XDUv6BF2hW1u588d76C6qS0F64Q5c, actual.XDUv6BF2hW1u588d76C6qS0F64Q5c)
	assert.Equal(t, expected.F6w1voK3gWJ7wphds0Li8dJQ52W503DLR2lfRup0XxJ0Jf5VBkLN4F56Q, actual.F6w1voK3gWJ7wphds0Li8dJQ52W503DLR2lfRup0XxJ0Jf5VBkLN4F56Q)
	assert.Equal(t, expected.FE0118sftca22Uv, actual.FE0118sftca22Uv)
	assert.Equal(t, expected.F2iof7ha2MhsYvU3MQhB6ivy6X, actual.F2iof7ha2MhsYvU3MQhB6ivy6X)
	assert.Equal(t, expected.X2QRx455Wp2, actual.X2QRx455Wp2)
	assert.Equal(t, expected.YE348JHYKM4IewHYsW223u6l8qUOerISlv61J67T0khhR0dhXm68l24QTs018Hb863Fe0HRo4o0i1eRLfWQ, actual.YE348JHYKM4IewHYsW223u6l8qUOerISlv61J67T0khhR0dhXm68l24QTs018Hb863Fe0HRo4o0i1eRLfWQ)
	assert.Equal(t, expected.C56q26r, actual.C56q26r)
	assert.Equal(t, expected.I4n0klX453ugOp6teq651uW4, actual.I4n0klX453ugOp6teq651uW4)
	assert.Equal(t, expected.C7gx8vVuco76psH10yhRE6nL45rr1b4fvWpbIN33411t02iP5DP4J3huU7Xl4B8w2t0h7471H5iW2jyxcD36UN58rsgN07N0770a, actual.C7gx8vVuco76psH10yhRE6nL45rr1b4fvWpbIN33411t02iP5DP4J3huU7Xl4B8w2t0h7471H5iW2jyxcD36UN58rsgN07N0770a)
	assert.Equal(t, expected.Y3Ox80LRyRuLCO3P2yJ6I, actual.Y3Ox80LRyRuLCO3P2yJ6I)
	assert.Equal(t, expected.DjE0g7jxOBuhNhf51pXvj5LhU6b6qsP6MX5Nv81EBy, actual.DjE0g7jxOBuhNhf51pXvj5LhU6b6qsP6MX5Nv81EBy)
	assert.Equal(t, expected.Brnni3PeEi2XU10cyr6j0ppdG2a05BdIGlBpM687A1P66c2t, actual.Brnni3PeEi2XU10cyr6j0ppdG2a05BdIGlBpM687A1P66c2t)
	assert.Equal(t, expected.URL7gIQ16gfSyqcw51FQaE3wfRgTe5CE2uU0STpV7D, actual.URL7gIQ16gfSyqcw51FQaE3wfRgTe5CE2uU0STpV7D)
	assert.Equal(t, expected.P2J035F8Nk1, actual.P2J035F8Nk1)
	assert.Equal(t, expected.PX06Ty52WFLl3HwggY387niD34HBEp4QTvD37SXFBog7vh02RUMc4OI1OKJ641Ib3r67P, actual.PX06Ty52WFLl3HwggY387niD34HBEp4QTvD37SXFBog7vh02RUMc4OI1OKJ641Ib3r67P)
	assert.Equal(t, expected.We4o7Xt5TTm7Q0Xxb8XqTc1xj5LQ1daGEaqlQvJ6dF7d6TT42xIpDh, actual.We4o7Xt5TTm7Q0Xxb8XqTc1xj5LQ1daGEaqlQvJ6dF7d6TT42xIpDh)
	assert.Equal(t, expected.CHaMB7Cw311V418f2W24J0kVYkVe548wfhv08Bk0275QnXbI, actual.CHaMB7Cw311V418f2W24J0kVYkVe548wfhv08Bk0275QnXbI)
	assert.Equal(t, expected.PuNeJcTcveW3p5, actual.PuNeJcTcveW3p5)
	assert.Equal(t, expected.OW6r70Djqj, actual.OW6r70Djqj)
	assert.Equal(t, expected.J86wCAlDW41w0cl4FjCU08471hp2Q5Eo2Dw4Iwj8E2ySsy26N5yF30GUw0, actual.J86wCAlDW41w0cl4FjCU08471hp2Q5Eo2Dw4Iwj8E2ySsy26N5yF30GUw0)
	assert.Equal(t, expected.Tn0RerY8Y81KFf1w7o82KukpV68736TLVKJP66s3Y1V54qT86MXtoWNUa2fudidETwHgV, actual.Tn0RerY8Y81KFf1w7o82KukpV68736TLVKJP66s3Y1V54qT86MXtoWNUa2fudidETwHgV)
	assert.Equal(t, expected.EqXn3kdLI15O5GQyr6253kb7VdNv2uC4mwNLfOmR5SdMl5rdtum6vypL3, actual.EqXn3kdLI15O5GQyr6253kb7VdNv2uC4mwNLfOmR5SdMl5rdtum6vypL3)
	assert.Equal(t, expected.Q7vqKj51Qb84d48JEIvYNU75aWo0qAYSVjXY44j8AC7UD41pl51AR1d264nv2, actual.Q7vqKj51Qb84d48JEIvYNU75aWo0qAYSVjXY44j8AC7UD41pl51AR1d264nv2)
	assert.Equal(t, expected.JT0Mq1CSXk2rO4t1PMDW0YOiB14ih1pm6FgdvCbG3332i2Kl5c, actual.JT0Mq1CSXk2rO4t1PMDW0YOiB14ih1pm6FgdvCbG3332i2Kl5c)
	assert.Equal(t, expected.FmUJmlMOtNqi4uUR0N1K4o72E064hbj8jX5N1sEx6U67l680FymAOtc7StlcwoAdyD2TeJ12arXQg, actual.FmUJmlMOtNqi4uUR0N1K4o72E064hbj8jX5N1sEx6U67l680FymAOtc7StlcwoAdyD2TeJ12arXQg)
	assert.Equal(t, expected.BA0rA3ud8m5KCYi78NeWo7c7NSHQ3PdEMIPky2QQ3v6S8Ur6uqy1OK, actual.BA0rA3ud8m5KCYi78NeWo7c7NSHQ3PdEMIPky2QQ3v6S8Ur6uqy1OK)
	assert.Equal(t, expected.BXBVF4Q4iNdqfgAeaqK64s6t6Bav8ojL6KOwyY1052sL6bT1u04OSE, actual.BXBVF4Q4iNdqfgAeaqK64s6t6Bav8ojL6KOwyY1052sL6bT1u04OSE)
	assert.Equal(t, expected.O58y041UjLx2ew73u70WJcqFKrC654XTLsmb4EcM0aitN05VUyjp1kW15i0736h0pjco10OpgY, actual.O58y041UjLx2ew73u70WJcqFKrC654XTLsmb4EcM0aitN05VUyjp1kW15i0736h0pjco10OpgY)
	assert.Equal(t, expected.LO82V5X1RHx52i61e51fp6x2bbBgSBo70x34fuW5EtcR5GUlu4TP1B3QD472K8y8QFoa5YGg7425VTCka7A0rGNUyHU6, actual.LO82V5X1RHx52i61e51fp6x2bbBgSBo70x34fuW5EtcR5GUlu4TP1B3QD472K8y8QFoa5YGg7425VTCka7A0rGNUyHU6)
	assert.Equal(t, expected.A6SPqF3gTP3b7kAHKcaQCjV1tT7pCTJOlP7Jtd5YWGP78Diw7N5F26kNPS, actual.A6SPqF3gTP3b7kAHKcaQCjV1tT7pCTJOlP7Jtd5YWGP78Diw7N5F26kNPS)
	assert.Equal(t, expected.VT4F711YMnK5mS3RNF531Qpj8Kpo68q7nTp560YRLMsrAdD87o6U, actual.VT4F711YMnK5mS3RNF531Qpj8Kpo68q7nTp560YRLMsrAdD87o6U)
	assert.Equal(t, expected.Ex3gdR0qi8fJJ0U86e0V1312cPMIldHI0cOsUrv4Y2NkAwtJ4TnHBS36nt1Hytn8il748is3xgE1i011j30A0sq2, actual.Ex3gdR0qi8fJJ0U86e0V1312cPMIldHI0cOsUrv4Y2NkAwtJ4TnHBS36nt1Hytn8il748is3xgE1i011j30A0sq2)
	assert.Equal(t, expected.U6MGiW6FA2eKQwok4h4ED45D4x8PX3X8BUeBy2GYHGvGvgm67Us8g03, actual.U6MGiW6FA2eKQwok4h4ED45D4x8PX3X8BUeBy2GYHGvGvgm67Us8g03)
	assert.Equal(t, expected.TLb58mN, actual.TLb58mN)
	assert.Equal(t, expected.VE00cVoA77jj656Vy7oj643l2pA, actual.VE00cVoA77jj656Vy7oj643l2pA)
	assert.Equal(t, expected.Uar0CX3o5u3WfqE7o6T3QSq2fMeM4QF, actual.Uar0CX3o5u3WfqE7o6T3QSq2fMeM4QF)
	assert.Equal(t, expected.LYLaR8VdnVR, actual.LYLaR8VdnVR)
	assert.Equal(t, expected.Rnuds8B23000PDI6wyYK1nl12VYXV1a4mtyRRcRcsO24C5yB, actual.Rnuds8B23000PDI6wyYK1nl12VYXV1a4mtyRRcRcsO24C5yB)
	assert.Equal(t, expected.OAG28SLLnNl77d27AKMI58db7M3LR65jX07870MAfIi24C2I3Go8cG2751b1H271H42Qp2, actual.OAG28SLLnNl77d27AKMI58db7M3LR65jX07870MAfIi24C2I3Go8cG2751b1H271H42Qp2)
	assert.Equal(t, expected.TGa26aGI2D3tjK4f72LR007N35qKNY4XvFpw5, actual.TGa26aGI2D3tjK4f72LR007N35qKNY4XvFpw5)
	assert.Equal(t, expected.HoB7nFxQHUY652NfN3D7g826sU8PN4K6xev42oO78OxmgFb4D, actual.HoB7nFxQHUY652NfN3D7g826sU8PN4K6xev42oO78OxmgFb4D)
	assert.Equal(t, expected.U2ii03A226243h1p18Wg7oUPc4sI, actual.U2ii03A226243h1p18Wg7oUPc4sI)
	assert.Equal(t, expected.G13VEkg4508ej7kiKJ3XKjDT8t6P55T711miJKSR0H27705r77BX5242YatijJe6X, actual.G13VEkg4508ej7kiKJ3XKjDT8t6P55T711miJKSR0H27705r77BX5242YatijJe6X)
	assert.Equal(t, expected.H82PSTYkUKHQ, actual.H82PSTYkUKHQ)
	assert.Equal(t, expected.Q1yYE381754DtTlMHg2UBKRd0hWe6Xp1nnrVuVc38WB2WDlTx, actual.Q1yYE381754DtTlMHg2UBKRd0hWe6Xp1nnrVuVc38WB2WDlTx)
	assert.Equal(t, expected.N22e881N0OcAnCFP7gD3o1lu66pnU07f8U8fhGjAE5fsm617Wknd20aBJC, actual.N22e881N0OcAnCFP7gD3o1lu66pnU07f8U8fhGjAE5fsm617Wknd20aBJC)
	assert.Equal(t, expected.JjyJi8iM45edK2A, actual.JjyJi8iM45edK2A)
	assert.Equal(t, expected.I535RC1JeWO4pDRE81cHuoubcU6aUc1Hogq04NdIx8dTYo5vjq8UtR50F2J02v63v0qD6l8634PmdPW4yY, actual.I535RC1JeWO4pDRE81cHuoubcU6aUc1Hogq04NdIx8dTYo5vjq8UtR50F2J02v63v0qD6l8634PmdPW4yY)
	assert.Equal(t, expected.X56gJ8sLWa2wbeWqSI7qK17563i6yh3O6fLjiQMh, actual.X56gJ8sLWa2wbeWqSI7qK17563i6yh3O6fLjiQMh)
	assert.Equal(t, expected.Om40t0dTFg4XxBk247p8v26OYHbQKDogx04o8qEmNMdbOP, actual.Om40t0dTFg4XxBk247p8v26OYHbQKDogx04o8qEmNMdbOP)
	assert.Equal(t, expected.V4oAFTJfic0g5r7Aj864wCtu7lMblg0L7i42xUHfB3M4L4EejhqpxusAAASj7Y0b87xlUB0l6485C0n5iL832aknd28pG7L, actual.V4oAFTJfic0g5r7Aj864wCtu7lMblg0L7i42xUHfB3M4L4EejhqpxusAAASj7Y0b87xlUB0l6485C0n5iL832aknd28pG7L)
	assert.Equal(t, expected.NEFmR3GYcsX67Uai11Ybyo4KyPhTa25E17EUEX36agNChN2j2k5064G830g08ta5Oje4rskK0xToMvYXWQYjCP21wFQF82xEbRw, actual.NEFmR3GYcsX67Uai11Ybyo4KyPhTa25E17EUEX36agNChN2j2k5064G830g08ta5Oje4rskK0xToMvYXWQYjCP21wFQF82xEbRw)
	log.Println(expected.NEFmR3GYcsX67Uai11Ybyo4KyPhTa25E17EUEX36agNChN2j2k5064G830g08ta5Oje4rskK0xToMvYXWQYjCP21wFQF82xEbRw)
	log.Println(actual.NEFmR3GYcsX67Uai11Ybyo4KyPhTa25E17EUEX36agNChN2j2k5064G830g08ta5Oje4rskK0xToMvYXWQYjCP21wFQF82xEbRw)
	assert.Equal(t, expected.GsXO60OQk2J1D1Fo5wDU86as6fmGC3yR21UgU83YOE2ROf3Sd0XpN2BCjr6R6WLtW3h2hUunD2e7gOURL3ieJoq, actual.GsXO60OQk2J1D1Fo5wDU86as6fmGC3yR21UgU83YOE2ROf3Sd0XpN2BCjr6R6WLtW3h2hUunD2e7gOURL3ieJoq)
	assert.Equal(t, expected.BUL7SnjImLQ, actual.BUL7SnjImLQ)

	require.Equal(t, expected, actual)
}

func TestFuzz_4(t *testing.T) {
	var expected, actual O0x5DXIdW5D267tvCvQ35CkEn7oYIU30oGOTYy5EfC42i8PPpwin8n55cNSWllXyod
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, O0x5DXIdW5D267tvCvQ35CkEn7oYIU30oGOTYy5EfC42i8PPpwin8n55cNSWllXyod{}, expected)
	require.Equal(t, O0x5DXIdW5D267tvCvQ35CkEn7oYIU30oGOTYy5EfC42i8PPpwin8n55cNSWllXyod{}, actual)

	actual = O0x5DXIdW5D267tvCvQ35CkEn7oYIU30oGOTYy5EfC42i8PPpwin8n55cNSWllXyod{
		CwnoiTT: rando.Bool(),
		K3tR0F62ayb40x4Gtn4e7kOol87p7sJo55jH548y8Y0fYeK6:                          rando.Bytes(),
		CxPGQOKcIY1gS7HDiPs8gMNQkvwx4oDq511AI4HnKf07oJhum6M6vA2C47158JY0kAcHD:     rando.Byte(),
		DY72mjEfkQof3eUSDFCYo3qbdnxGsyl4oU4XkL0H2yLEdRx00844sYDmo716jQJIGI08hME0G: rando.Float64(),
		Sf44mn: rando.Uint64(),
		Y4202JJ58lJFxC5M3IGSCM1arBRxjUG156YFm1UEu21rXb71xiMks3c4uRMH01JnvG4AkW8u1O5WbQfKTd0: rando.Bool(),
		A7sako5sUl618Cw55rd8E40X3W58wgwg5278C5mqSssANwC3h8W8xVAsX:                           rando.Int64(),
		WPy74ranJyUTBhO6j1O14608e1j14AIJa3Ld0k4Vfx6O4R7INLT0:                                rando.Byte(),
		KF8:                                     rando.Int8(),
		KdAU406a8jWCmU0M6Nx87XOSUVBKgNQp2cPC4Wh: rando.Uint8(),
		K2IE5QD3kD7qelMe72n60d2U71njtdKf:        rando.Time(),
		BL3yC8ql88yU6CjT8Cvo7jST711xldA4Nyb47503C6jULeLNIE4VYJrRrQmKIQghgWMOI06mOj:            rando.Int8(),
		GA4pJHF6k0JfUbMJ0k5Eo8Wg4c5S0Nd1bf30M63860jtkk7s2i8ROJ6rWHOmXBfI655VLOOoc1nim0kYJ7nE0: rando.Int32(),
		JvSTSs42KDBJ5600KBTw: rando.Int(),
		SBf:                  rando.Int16(),
		Uat82y8083s8XTd5mM5kFEHXgFU3uvYl8Vw24dQ1M8PU0J14J7j4DoGj0Yv1w2:                   rando.Uint16(),
		J5m5WgMC1kDwGX83AQKKQOB3jmayDw85ly5331O:                                          rando.Uint16(),
		M1T1vkIWcmAo6b7D8RJVVjN1ML7iFC7RQ7MhuLDlcu273UgUkKCOcdhJcXA535MVKAU7bK1tA0b11qP4: rando.Float32(),
		YhI6wLQx8642Qxs62hnLFj3f38j5830dADpQCC1TAMSDoCx:                                  rando.Int64(),
		P1MTN2K53f4vFu47Nq77ChYFHbo3cugNLFC7IjmPB6O2Yh6l45ByLYsm84xk06:                   rando.Time(),
		NQ25SXoR2W6EdcJ2qh885: rando.Uint(),
		UfyuExdG5MyYsm1:       rando.Time(),
		UMwf0Xi6876K5PBppfrRuT1FLVWh88NuPy6d3mj7r3uaSCWe5cPUsSHqa7RB82bHcSOhchkv2Ko2yMD4mU60Y11eG: rando.Float64(),
		EH4n5b55CYsMuP2CP8eDvg2ndeUI6lv1x8yEtW510IdxG5:                                            rando.Time(),
		E861rf02gHlam6RkarIbdYt86ODCLlnS3323yuXdTXdhXE67GiIhtSHvTWxiqaf1Am7P:                      rando.Uint16(),
		WgN4MRAKSU6fsOsIU4627: rando.Float64(),
		H3w6mA3vPBQIix80fXx8cwdWlpAW72Iajj37sl3kBhm5ilFp3B50SLEt: rando.Int8(),
		E3E26VLqNXWkVK5YFc2YjL1tC05g0Qx5Mpx:                      rando.Time(),
		Yxk635XJgHai6GTOiIuBi1:                                   rando.Float32(),
		Xm528xGY5VN724Qw6TI64WQP62MFjixAVsP4CLB3h255rsngUMsj5GK00Rk6p3kygvG5R42438ESf8skHXte11wrX24n1: rando.Int64(),
		E12ro:                           rando.Uint32(),
		J3IVanIdCgocc8ckBu42h6NIN5oVITm: rando.Byte(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, O0x5DXIdW5D267tvCvQ35CkEn7oYIU30oGOTYy5EfC42i8PPpwin8n55cNSWllXyod{}, expected)
	require.NotEqual(t, O0x5DXIdW5D267tvCvQ35CkEn7oYIU30oGOTYy5EfC42i8PPpwin8n55cNSWllXyod{}, actual)
	require.Equal(t, expected, actual)
}
