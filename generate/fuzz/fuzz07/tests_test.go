package fuzz07

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_0(t *testing.T) {
	var expected, actual Q7b1o0
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Q7b1o0{}, expected)
	require.Equal(t, Q7b1o0{}, actual)

	actual = Q7b1o0{
		EW0PvUj7hVq5Xo75fX0h5QId8f85KNx4tsvnXD8eiV38By5OnjN0kiAdYbJf61jsK1lueDp8r1UAr47n: rando.Rune(),
		Il0F4j82ACubpef2j1IrAhx5m6nNoNTH20j5D5OU1TN80aWeTXhMF3hj77Q1DxepHtAX7:            rando.Uint(),
		DrkvC35tAlw6QHJ1DfP0BBdev8xMtqH2wYkLY0b1orDAIf7Nw30N81XSOC38:                     rando.Uint64(),
		NP: rando.Int64(),
		KXpDV6xS5158NjPEDL3S4oSuEH73Txwa1TEt3Ou4d1m02COptqNSH84t5XWlVev86H64t4naaj0Iv0J8ffFSbCxM3gnisKHJq: rando.Int8(),
		OYOqXAAFPW2hseYkXaj7lNrjEgoJGA0qxG1U31ePrlpN:                                                      rando.Int16(),
		WlHUDXlaKk0OQct001wR37867FnaBAUwX1JJ:                                                              rando.String(),
		E0nDN4FS2VM0q0tS0Fr58WAA10sEp5Xd5ef5ghlFHg6J7146UqD18v0l6Iu4eIKCq16P4LDH0eJ2CgG47:                 rando.Uint64(),
		UVhTv7V2r4aHqVnaanL0QQ2Dm7GUALFaJOlslXnFWEX4c321g16dTJt0RYJbswYdfSIbFil:                           rando.Uint(),
		HyU225f2MG3wP7SK3Mw8Md8UvO55ix3g57O704Y6thTM1H2yiB7oVYnVncEBt61AtO13:                              rando.Uint64(),
		W0Qo4qs0pl0nIEa3t1qYB380Yuq3EbJA628kIoyykh2457c83k0I56qeny:                                        rando.Bytes(),
		Ma4R73UrwGrrnN6kj28W64GU3heVROyH5ysDOs4LGcImK1ps5KJ7fT3552OlW7b77u8BHB:                            rando.Int64(),
		LNn0fKv20C7mKGTvXsJtipxs5ahqQbkmeGR0HEm2JNem07Jxn3A2350cqNoV0vFlGcBku43:                           rando.Int(),
		C8MhA507qAty0326Egc4Crh77siKW38Bt6G4aqu0cIaS0007b87fnX48iktc2r:                                    rando.Byte(),
		DT2PQh3xhe3ExtH4d1JL1qwo17BicGM4Gs3nRoIS1o:                                                        rando.Rune(),
		EpHxkXF816e7NLMYo1d4o2136eH7KS3Gt8u7e5w2ui8XDaD0S8B6MmKdkn678E2:                                   rando.Uint16(),
		MOHXe1JWyC56LQUyMxQsgXniool7I0575FFEmv63xetn6FA44:                                                 rando.Float32(),
		Ts28LFwKRu0: rando.Int(),
		FQoHu6H85U87keY6raMoO7up375IqV41H0d00UyBPC6Knta7Whb8FxIaq81cxWQl:                            rando.Bool(),
		RNc1XeHafPvw5oiHr151pb3c02r1vuhf5ee0gf422SOvOe6vVB4gkBAXulummA177Ydy67Y52sHNKC16Q:           rando.Uint64(),
		Ox80anNh654Mp0g7sUIFG1rj7VcG35jgsf300FckXf6o6tfMedGRJVe:                                     rando.Uint16(),
		A0D7Ie6psS026jGA7tMa788Yr4FB61XiebHk63Qoc2WcyvJ5WhwLOm6mCB5YOT25Pb7Ulh48mtDJGgq1Hf3535Bpj0h: rando.Int64(),
		Pp0fxddUe365MTEKafpv6CfhMaW3DJxE3I0J6dyfxM6j0F3B37:                                          rando.Uint32(),
		Vqgnp6Y78Gr50k1viP: struct{}{},
		GaPl6SjT34DnPFK3Nc1D0MBq67aaXSm4lDAqTMw150eK2mhwKp08q6u:                                          rando.Int32(),
		ArAWX023i0R1HBuFesm53mmekEoClY3347lgDNi1ulbJM81YWV3sw8R:                                          rando.Uint8(),
		Jp7N48646mFl1gx2Cd21317t8wuQr37lRG0gN8YENw103:                                                    rando.Int8(),
		G3N0283bBPV1ISvjTRpe41ch1b8v1N414enc875KuC:                                                       rando.String(),
		YJ8fhV5Nj00es42MlMQx4LJG8KHCV18R4l5771w5j7A1YaKbwa14huykNa8s2BpvKY:                               rando.Time(),
		Q41C0reF01ewLG5087AG0A6fq1vyC:                                                                    rando.Uint(),
		B3quDVQ15HcHAj20250i0lM85rlhtKxpXwe7205msabhi5L5TfHmS7y36ggfC67cLQawx7mF8KSUj73Lw65AK1MO8x8:      rando.Int32(),
		DxL2t83shoY3Qe1jTdnOJ681V3VigL181aTp41cl1ndgOiWg4pE4V4UNdX32QnbdbxC1f25Y6v6xrkhNS51U626lOd856rE7: rando.Int(),
		Yf2pt074xrY50aJ8RA5Iv2xB18178121kyUU8f011dl7H:                                                    rando.Uint32(),
		XAdGPnm842r8682Ky3Ax6t3ptW7:                                                                      rando.Uint8(),
		Nh72JeXWq06I0ouJNrwQy4bXIBs1:                                                                     rando.Int(),
		W2L1ef0qs3DEipbDrM27gD2h4x:                                                                       rando.Int8(),
		T6w7EY5cIAu:                                                                                      rando.Uint(),
		HHp2ij0X31QA1K8u6p1Q0p5LxUi1L6j3Dl4DHP1H606Il2I0Y0UYifD7a5FwVjQQ3J573pnyr6EEUNxHyg:               rando.Uint(),
		F4PhpRLJx3xll8W01MANi4jF68jD8pdwg8Rl5gq267K0Pm3P8bDh3T7nlRH6g8X:                                  rando.Uint64(),
		TF401xPq0K5sxd3fKbFHp8KTaY13tupV6XHnQ3p4:                                                         rando.Int8(),
		IXjFqlS1ptE6KQ80102YYM5i325r2jP5n:                                                                rando.Int32(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, Q7b1o0{}, expected)
	require.NotEqual(t, Q7b1o0{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_1(t *testing.T) {
	var expected, actual HtMw8EU
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, HtMw8EU{}, expected)
	require.Equal(t, HtMw8EU{}, actual)

	actual = HtMw8EU{
		Iah8PQ3KH33L60XTw5YgHQDi8q3g0p7B7d7rh0AqH465O1garOrVmUUB2TpB3fg661Eixs8k3: rando.String(),
		U0aG4ETsX1r3v713dY2dK212xW:                                 rando.Int32(),
		XGgJ2gp56ME0b77n8u6rCx2pGB2N4ruoK50LLORh:                   rando.Uint32(),
		HLkbKv45s2GJPo5iv58CJmQ4fBN6Tk22rMs4gGmk1j3aTSHBydtVD6J2Tr: rando.Uint32(),
		AV4iU5qq6xvQR3G3KopjNj:                                     rando.Byte(),
		Uxr1BlQn0C4M10oaH8Jq234ixS5Bk1BgWpkHNalMoR3NbogGhauw4U154L4ljN3LW313py6XpbVJcBp2QClt7DedV:  rando.Time(),
		XtT0Lr326RYWj6d2AHW1yb0tRvpex24Ysj5JO7Irsb5Dm1JJ2KgmQQf44VVV2vkI7ODclqfIeOtL281G1FjF6av3Cy: rando.Rune(),
		UJNW0B6df4EeD58mx: rando.Bool(),
		TcsFu5o5fb37uE0bfeW2WM7InhV0CW2xs3JshKl8Y8He85d85Pp:                                          struct{}{},
		Drk4lQ7gny5L55882JJc0CDmfwP41y2yFQr0dHL453PTFNL17lI0lcxMUO5g3Lxhl7TsNF6OQ87EbTu0qT5e10TcjXP1: rando.Rune(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, HtMw8EU{}, expected)
	require.NotEqual(t, HtMw8EU{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual Rm1wILWkrY5u7usw1WdA8H3UPGUkY4p2wFrYuyIOLTO
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, Rm1wILWkrY5u7usw1WdA8H3UPGUkY4p2wFrYuyIOLTO{}, expected)
	require.Equal(t, Rm1wILWkrY5u7usw1WdA8H3UPGUkY4p2wFrYuyIOLTO{}, actual)

	actual = Rm1wILWkrY5u7usw1WdA8H3UPGUkY4p2wFrYuyIOLTO{
		AP1Cjl21cim4rM:                         rando.Float64(),
		ON40YelI368B6GSh0W6II5Pk145cTP5na3BD06: rando.Uint8(),
		HK0WoBjo2ecwC2rqee16IBg22gCs8c66RXlbNK4c0KpP2062I6aPfFK0kFXw10NIc388q0vS3WRVqgg7R: rando.Uint64(),
		Akj0vIi8wu6TLt1TaIMW2RuDJ2Y6vPL03:                                                 rando.Int32(),
		H088o:                                                                             rando.Bytes(),
		Uls88Tv4n5oc8R7bPfKFM7ViF6r55v2O373enO5j:                                          rando.Int64(),
		JHFFqS27RVa45:                                                                     rando.Int32(),
		H8gBN0Y1sT724J7IHs0oU15Kw0hH363sc7xv0862Mc6814jb3:                                 rando.Uint64(),
		WKa010865l33EgwPd538a6Tcq:                                                         rando.Time(),
		U6J58xqXU13TEUKIw0B13ReNcHGIWg05cy0kibI3x37FdS4rreC3h3yjc5Y6Gg:                    rando.String(),
		Ui3gIFuS0vUa1644E78OS72770oe3iUKN4xQ7868xKew8Y6Hls2W67A3y6yG2h387iiP7hn:           rando.Byte(),
		Wdd7Dbxvue0:                           rando.Uint(),
		Ox7nhK1r154jxc6MkotPqdV8n1y4FKDJX62Bh: rando.Int(),
		BwqifsBsMH0II2iME7pr25AaD6MdU3J5:      rando.Uint8(),
		GHxvrVJkIt4f46f21PR3GaH88LjnUsM41Q24K7b0nu4YlTiSqoIXhn5iA3FIgY4Oyi: rando.Int16(),
		I6w1:                                 rando.Int64(),
		FT8Phy3vyHVwv02u258iS06on1viobSWH2W2: rando.Uint(),
		O5YSc735jpmwyj2uC3sJYlO31G2yxQ86cLRQ0tr1sbdL216X260r8g7bP5K5:                                  struct{}{},
		Pb5dSod2nAKEBCL5rw2vCpn76IdO77IJ7xs8:                                                          rando.Float32(),
		PJ3kaE5W1UC7G5ui48YE2:                                                                         rando.Int32(),
		P40a8Jpb2qQ4WoK132h6ACL5vA21fLtAHXOwTw0Bco5egm7gNPP7161kXJQyHH5G77I1GYHBtW4BucrpC3JB7:         rando.Uint16(),
		LjrVjod2BD2S12oc3A8fe6blQX7VujNMKA7216T50HX21kb1a647s7831VTf32XYC0Re4M5cMjr3Jbh4X43G353:       struct{}{},
		Q5j0BL1WDuiod3rj1vixmWAh0h5eSt8PlEhVOJTEo0330sr2x2WM3oVI4e1UiG5o0aQD7IEJa4wxs2DHj7u7C:         rando.Uint(),
		P4T6HkYxg06BJN7AQQk14gjWWN8uLvURx38i:                                                          rando.Uint32(),
		P1PuAGuKYXFG6k78S4578qgcFdSf5766s6i0LrA4wT:                                                    rando.Bytes(),
		VS2j1qBh8B747k68aB3AP54Ryki43VrXwk4h:                                                          rando.Uint16(),
		Iu4j6K3810xRRmrGU7u6diX1h3AqrcEDH1RBAAHXt62ljgQ250:                                            rando.Bytes(),
		QL2200VX52l550ub86Kk4gj4n78abj78kx0Ja6mjRV5:                                                   rando.Int32(),
		Ov7XY3d4xuOxrK1Co8h8gX6u5gtdfKkxyYcP04H736v3LB6BK:                                             rando.Uint16(),
		E78IQAe317yyxteSJMW4KyHA361k8a2W2P3mXHGCSrxX3810nvAdepRpmw4542u5mNJgl0atdG50R74Hy1Gj:          rando.String(),
		UimS2UR5BDyFEd3n8a27125omGsGYhX10uo5hA1A7MC5udg325TI418K4Qh83JqF01S1VqSr3D0g5LCg:              struct{}{},
		A2F8F2HN8AdOCNs1h0HM2622:                                                                      rando.Uint32(),
		N2MPq23pxaG2vjwYgvhXgLlN1yyCY6mcipT0j7B2SM40yQyF84G72L8f7qCwc3Xh1Y8p:                          struct{}{},
		JLr618s373m8qjy4r2hh5xET4T7s6RiyXV113Iw5Pw37p7E4lbal76RUlA:                                    rando.Uint32(),
		XfCUPGOn5x4g43vh3vc7NQ02S6b6a3nw7QGiI8Ngn106C84KCLT32NK477j8q8SelP26b3Ec8vy0yL6LgW1JpiU5tP5a6: rando.Bytes(),
		I5fjdA7qMg6Yf72Xe1xK4ut2Aq6Jh26S6562cCUvL:                                                     struct{}{},
		FSd6gXeDT31lYrH0h1EV5AHSTMFlt7RhX234ESm138u4YA8pHSO728Pxsvudr4x6BLy8F5T5ETSdosD22q1KmBBBE:     rando.Rune(),
		IS22g7yLO152iXRYb5hLeI17wq0x5qDC85r4r82c0O0Ght0h2EhgT5xM34GMqa0lJ3q0A32glXMeX488Sc54FGTWX4gF:  rando.Int(),
		C1xsNLn5G664G1xV2Y4w4I04oQ373o0Q4a0aV1Kyr8n0B2VLl2A4arnVud75aCBQpC4Ofxs4g0:                    rando.Int64(),
		Hpc:           rando.Rune(),
		MxYM685tWHchY: rando.Int32(),
		E61Q7pvXFSe52531VhK20Fx3WiPF4MrHU136tiN4y0R0yB5l4cmPFju1tWF33d25ve08D8f77Tp62E0M53:                rando.Bytes(),
		Py3388ar6BT81KO2Vb3Wt2NciH4bD3r6I0861ggPPwcb5yT2ti7mEC6i7jfxMaCkcdgbJW5hX7N:                       rando.Bool(),
		WTOGcsVku7UQe8LN24A8p2qQ48syBAd4giJfXMNK72r6JL8Xwkpa6:                                             rando.Rune(),
		P51e1ufV1gtgq2A51Cyd37248DS5gke0AOiKc7keJjwa:                                                      rando.Int32(),
		K8N36g1JUioJt6h1Y7XpQ7qQONNx10dh37YcFYk5tgV3wp3wYEw3xPfpU311V13y2Fe7j5XT0Qc:                       rando.Bool(),
		Jq24bgp53q4SdSh8PJIm8aayLKn5GC6V4B1UXXrAC:                                                         rando.Int16(),
		K4ABPP5YfUc066VR08F1pIyCH7CKbhLmp36woyELFEqg2RQi7MlI4t708237T2p7OEQoci56axlGpe1838boSF06c5S3dbnEg: rando.String(),
		ToBe8: rando.Int32(),
		DHsl17Dwox51Wu5I08P0i764qA4Yyg27727GSVeMvp:                                                        rando.Uint8(),
		WN06VhK61I2E04Fj83xIfC8p6wwDGaRVJOir1rG2OMO0:                                                      rando.Bool(),
		Awd51c0y7E1OFh4p6qMPu54sE7rp7rm6WL15I:                                                             rando.Bool(),
		X2wyn24E1uc5qgji0U4H0H5x5jbDOlncdlfYD73ApkexUc27JY71RY5fJdN50t3Tf:                                 rando.Byte(),
		JF8v31pLP7JhbH227LF57no6sMi1p7cWipA2XV6V78N5ilK0b4q8k306D1e6vw45jd:                                struct{}{},
		WdKjX7b312Ua1FePtSUrhEtNF:                                                                         rando.Uint8(),
		NW6PIl51i717xauHKxVc5xBrVpYc7xpx3PkodXv8D:                                                         rando.Float32(),
		KM7cMXSE8oDuh80sgRPkxW4eotpRh6Ep6hY74VAdw86x:                                                      rando.Int64(),
		EbQJGdEWN7LfKy0jJsEPdeoPh5oxq3m4NL413yWHxI36SjgKn2V5cjwX8gHY06Y05INBu7YpPsm8unX1Jt87w8s8q2mLecoy:  rando.Bytes(),
		Q3i8fUO6Hm4aASiAgG4dM8B3KMMLf4Lt5dnL6sQB714byHh7by5Ry2e36wV6BDr1344O0Y874M4im23lnPWa:              rando.Byte(),
		DPKa2WXB8JVys10vrbk7572KhKII77RB4oIjia4KSrTaWkCY58SEncvKq6nxkxA6O60eBeSpXDXYk8VATa61R58w82frpI:    rando.Int32(),
		E1j6533Ht348v36PEQ1NPaF3PPdDN7YKe4P382XrF7O284NaSK261jpTQG3w5tWl1K46:                              rando.Uint(),
		XdC3vj1UMxnuGIaho44g7W3AT2C1spPxMkUOln8P2L6q0ArEjbMGn:                                             rando.Int32(),
		Wc36135215YEadthrW35vYnjNyusm185Bx0j87P2GQ145e4bLnT3FpX1mFOgYt5b6TkSMr25mQ08jK4G5WQ0637:           rando.Rune(),
		R4g8jt48gMRmt425i2vtM1tF70ax7Q1euokP7khgLE1RlX3k7NuoE5Sd7SjA61j28eR5te2V:                          rando.Int64(),
		JCQXXkyISiU8UjRXCRU58wKLx4sHRgn28407BoEgq6BKBkdGvKe17gcn1qYo5jV4PPH7FLB0b6YQx8kgn0a4SX4013IPX2qyR: rando.Int8(),
		Hg0ybpdF4B34p2tnSJ3Cr7HqA6e6B6PTC82n6J0t320318roD1:                                                rando.Uint8(),
		K4XLeysRgk6ELwnrO8IgLx8iKW4DRA85hHqLS88UkhE8a1nv3Y111xIraYyuAd:                                    rando.Int(),
		SbJQHCmFdHRvh1580oINFOFY5leJ0O:                                                                    rando.Int(),
		ATeDkAg47ba80O007frW4c5wHUb36Itm0rny273gvSM0AE36vFwSAdoM3AyFYEccl43o551N10Ijt:                     rando.Uint32(),
		T03ih5eW5B2YvhjL4PgWPFh7ihLtUiYD4YpHQQywAuVkU3e72nrmebqruU:                                        rando.String(),
		N3561d00J4BH7V6Q12v73cd7y0pS5XT4ycmQkLV2q6jU0KwPIY760wV83aN38UC62Jj0b35HsfKM7bl7Iix37YbMDtj5AW4xo: rando.Bytes(),
		F6S4SAlab4td0i0kFm2ubep0g4TKardByPBf:                                                              rando.Bytes(),
		Ls12i4Iw030667WR1ha13v7d0G5tY30337S1enROyXSXxeMPO5jK3mE663y37:                                     rando.Uint(),
		NnogpQA712G4dRq03q5w77U:                                                                       rando.Uint32(),
		Dw7mS67eu7W3wtH3LET8jQKR1e6vH7F3:                                                              rando.Time(),
		Vi4cwh272xtHq62470Dmpk7sG4fJPjnbyA0Iq8Q2WM6T44jD2wt7126oC713:                                  rando.Int32(),
		FOOx2O40c748yvwYe6lvsIMbtY5:                                                                   struct{}{},
		RB5445d3726Qy87561L40u6:                                                                       rando.Int16(),
		EKjD20dFe3ySe2GE3lN0PlyJQXWjLEge1G4VYcW3sBWPqW575Btp0X1IJF33:                                  rando.Int64(),
		BFgMubJHTvQS2QKiY585rFisL1QsF7WePOnEeP0NU5xh5KFeycWwEcgH:                                      rando.Bytes(),
		PTEwijP4437gvK710u0AEgGi242d84pwJvMJ0q6lK888n6mlx8D:                                           rando.Int16(),
		Wd4B3J83BLXVree3C76Blu3fMjK1b0cLR5VG3w1013ne3SDbU26P35K372p2HkeEeRQH1mrmK6dGGDC6S26y6BnS6Q7vD: rando.Int8(),
		Eyc: rando.Int16(),
		Tl1mRtUuVDmGrG4A3lFHC26EmUv6C3Ss651pTDf4pt6jlyqVp1F2NVIONIw47WjOtN43C76J2M2uQo37: rando.Bool(),
		G4W363dKT8Rb0Rr0ryool0: rando.Int64(),
		PK1rta1g67YT4V7:        rando.Float64(),
		LhL3xQ5Ik6WIUE2siLV5eyhD26fGU8ndO73xWGLHp3l40i57r536vc3b887ovT7PVsgebES7850: rando.Byte(),
		DMhDE73H347X407scUKlLRN031l4:                                                                        rando.Bytes(),
		A8w2u8XMNf86Mtul8v8k3HLNG53T5sHk7d07MqpnfB:                                                          rando.Uint(),
		S6sErE1e3B1tp3Y0LPl0Bo5uBb8u6XjapqUldMkoybhH66M8K4aw:                                                rando.Bytes(),
		L315oQ3U0g0tWTyS8N88qwm76wimLuaQN6lRXX5mpr4JH8h5gM3b10ICes1qHU1ChsDyj3DK86boQ6Q4m:                   rando.Time(),
		Fo4yTdD34VLVSI4k7S5RJ1xbyHS4LqUv62Jo8sV:                                                             rando.Int16(),
		WMS0o6G4EQWw2:                                                                                       rando.Byte(),
		Md86nA3uQ2sS3oPPNuRFGAD55RqFi1YrNytQiC5Lva113vU4:                                                    rando.Float64(),
		Qn4ekC7AYJNpD43f2h68XDHYGUXFdvU0XH:                                                                  rando.Int8(),
		AO5I25UBPaDp43s05y5l35FAPK87QIHm5tTN7q05Hq0E6HO45Ug31VWn12KHvH:                                      rando.Uint32(),
		LfTbFiN5hGnS06Y4450Ae5SQxNy3i1D3fhXsFyhXO87K1THJnHXtMsR01k0q3rj2uxH3jrXV7l2LmNK5psd:                 rando.Int64(),
		T84wGxXke7dhQaM26yMetu1E846iLAHBGG4CPur3IOx6m6ESvUh77moqc2nJDvUPa7d1PyvJwJ860n5ODm02IAbjF357:        rando.Float64(),
		H5G6sl3JOb3sw83044XU7Y54p7xn52I88Yl51mx4A4wD:                                                        rando.Int64(),
		CJ5mr08tRC4vkI0VL4bF45NVnHG0w01VKsctC:                                                               rando.Bool(),
		A3Ii1tPp3JVcgI46qfa0XY:                                                                              rando.Uint32(),
		UW22ENnY2iVlt2HG6FSlodEx43uQE1p4Sfjwpp5m6Cf5tmH4xc24h547prjmmmQFd1YiqSc6Kn5avuIGH4XJPJGQu4U0A48WsS4: rando.Uint(),
		HEUM0Ea6iwXT50uiAsTpBL15bB4gIs0454NkvCE2x1n21HqO6T1McHcQm:                                           rando.Int16(),
		RjMlTxg5Ryq8ivVnle1K31G458r4qOWjY182cBNS123M1Ku0DamhbURndWMKBdT18fC173x2oqIm:                        rando.Int(),
		BKU3EGwaO1MfjPF4vbs1p1B177bQVbuTlVLuULsOvY7P58yBtDt8aaIFan1Lm6ldh7YaM2vo2hG18N6J8vW2LROH2U5bTd:      rando.Int64(),
		R5baTyjWrOk31P8018ja1:                                                                               rando.String(),
		I0sKNcDo3m7dIyyGHfu1YsuK5:                                                                           rando.Uint32(),
		HFv8o1Vwrx4RM6c7gU2O782TfX6426Fr1n33VB4G030IJq6KkCn873FJOiuI8Hh4wE5KLLjxn35F2sa1Ya3ICx8J43J8yKG7HR:  rando.Bytes(),
		Ct670bw6xcG5waVAA71ldT47ouBT63Xdwa04eVtX6RYJ3FnfK1l8Yl6cii8501O:                                     rando.Int(),
		X0G0aky3: rando.Int(),
		Px14hfhkn3pL62sw76605g6bv55w2n45OLSD04C1H6O7n87TL158rnJpqg0yVAq42Ml4VuqWFfC31Q5yn: rando.String(),
		TK1qNk6115BmBs1QRbXA6: rando.Uint16(),
		SGwFQJMq7Wt0f23S2aWVmDaypaj0t0oa7477ik0e25dGRgH2he3N13hs6RNM5CR8dK7A6I347mBl3BWueb26RWnKuHN1pW2ukQ: rando.Int(),
		HWcH81AOl0o1H4VyRHi8ejcm2FxsDCT476Ea6r380NK:                                                        rando.Byte(),
		As6sr8JVsP: rando.Uint(),
		I14u0JrecufS160LJVYH2Qiulh4p0jTf7k7tK008Jln23373YoOc: rando.Uint64(),
		UsGb33b4R04i2f5rR3n4Br5AKJOTdByHwe5SDv87mm72oSS:      rando.Int64(),
		X1cuwGDwyfCM48R0em5g3pRPvS8q667YJc68WR47XsUwrI:       rando.Rune(),
		Vl8S2hI7TPlhPb4wD1:                            rando.Int64(),
		WD10vg0QfI3CM42NF2r8u1ouPKYqEvW53r:            rando.Uint32(),
		BcwRGB01fCDo408obDC8bR7XlQ0WN34xxR722Dd0gjQ82: rando.Uint16(),
		Awku5176qE6n5fvJf18Kw403a:                     rando.Int64(),
		N858R0R8g3Ys58G0hpukEhTs544:                   rando.Uint16(),
		D5w2VbiwDTnP3QvMap40RU2gCMw12FGMM3j4fOa8TBjp0S86DE0m6ixU60BfIt5y55gd45xuKcm:           rando.Int(),
		POSQRb2y7GWt354dGDXKhM1oDKoI2w4w5vf5OGVsxGqr68GrUcVSfY8K0oYedc0h4jc3LR7XqISMsNM4otUi0: rando.Int8(),
		SO0P00VdEUln83jWt4gbX2qUbOP6SV3pT4TG3d2p4s8c8gE5exEi5ELaLbcC3aLiI:                     rando.Uint8(),
		W32W7pXYtjc4RoV7a7yLycCHtm3MrKx8cX73gX637W2Sf1paEY22gm1U53QblQtQkD:                    rando.Int16(),
		ABHW68nS37f11YTjhij37J17RsBO8081jU:                                                    rando.Bool(),
		W:                                                                                     rando.Bytes(),
		FMrJ5rFr3ns5VXeD3NH3H78170N3m1PxhwmSa31QPA7ikInr4g288PfUcW6t4hJH2Nyr3DvXG0HB575ffJH12Sus41HS3qW781: rando.Uint32(),
		FYmi8QXQ4E5bU0hMT743qE2dQ1uXwR7700sryMXLu8O0e84C6h43WpgRyHiT:                                       rando.Uint64(),
		MXyH842hTbj1kD554uwP556XIBma1Y5yybx0LGF6kXpYxJOrjVKs6:                                              rando.Int16(),
		G5W4m1ABK3Jfc3tJo624QOHm5b2JdrDNWM8ckS:                                                             rando.Uint8(),
		U473Ad405FmUiYqC4QRp:                                                                               rando.Int16(),
		R3xoLQ6gNl7V71038wuv45018CHuR8ltuQeG17C0k61LkTq2m:                                                  rando.Int8(),
		MU4XABMA45j44f80O843hC8wveoyd14x:                                                                   rando.Uint64(),
		WDTrpJsP64dJ7LnRDfkYc0mGiGA6u6rEQW3I3k08l8nI1WBjr4Sj21GJfukB4K71an8F5cd75:                          rando.Bytes(),
		Uk67re2Le553L75b1E4aM2GJVSI4l8gkqikBW2IU:                                                           rando.Int16(),
		E384S1GW15u384p0CqTxfP8NUUMaQmxN116hBfGJk7w3cG17H24KOTBf8131TeH4BkG4mTIx7U5Hcrr78B4QrWh4FE2l:       rando.String(),
		IlMksEP: rando.Uint(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, Rm1wILWkrY5u7usw1WdA8H3UPGUkY4p2wFrYuyIOLTO{}, expected)
	require.NotEqual(t, Rm1wILWkrY5u7usw1WdA8H3UPGUkY4p2wFrYuyIOLTO{}, actual)
	require.Equal(t, expected, actual)
}
