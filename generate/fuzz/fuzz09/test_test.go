package fuzz42

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_0(t *testing.T) {
	var expected, actual QQXGKCa4tQi1lgGl81IDkGRMc3627IiA3Fh4aC6y3815PWv684S
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, QQXGKCa4tQi1lgGl81IDkGRMc3627IiA3Fh4aC6y3815PWv684S{}, expected)
	require.Equal(t, QQXGKCa4tQi1lgGl81IDkGRMc3627IiA3Fh4aC6y3815PWv684S{}, actual)

	actual = QQXGKCa4tQi1lgGl81IDkGRMc3627IiA3Fh4aC6y3815PWv684S{
		MuX3b80ex64d8fcJLqs23gdqXpho2N621oVjtHvsv234l7xtQMmemQ321MaOKaDu:        rando.Uint64(),
		Kl6hsSyQgQLvH3K4rTu28T00sPm0nTQ6876X6FlY785TpT1ePW50Yyb2EXgXV1m:         rando.Byte(),
		X0mNu7E0Bq4g2uEcwgD3WOtt7TNFOjsR18sdjq3R0Ts3xiPI84s6xcpM567863oq068BBSi: rando.Uint16(),
		FgfHynKH3b011DN7P17Odg8gHq758qo8am7O71Qvjlu6y77S11x8S18uphU2hbg8S4F7A:   rando.Uint(),
		EF6628P:          rando.Int8(),
		MEKnFAP8R:        rando.Rune(),
		HmlcnTJ3R1B37BHX: rando.Int16(),
		D:                rando.Int64(),
		AS3Vt6OV5JTb13t21eH626kShBV70J57Bbx5B56Vwb7dR3dbm28iYEIAPjm3f3MG1Ypu78a1cuu8v4P1Oh8:            rando.Int(),
		D7GbaJha7uStx35L35hkYi32NSDwo8RHR5lpQ1U8443h8FG2Gms8Vww8OSHxiw2mKPFJdc7084S7q8KW6q68fgO037L5es: struct{}{},
		B: rando.Uint16(),
		Ffo5YW84WxKrXfVHStGAXb0EFp7khH7VTi1l4FnM405:                                                     rando.Float64(),
		Q0t5CeJr6BywBK6s8y87q4Nr38X37jNm168V28M7CP5l1181A6j410Kcir0Eg6SJDNfxpB5u6Kh0xoR26XMET1WQrf0RTY6: rando.Byte(),
		Jp0g3ikY715Kh5va2ox8irgDq826c4GY1g6k2F8Si3Tu2EdAB6846HbhEsPWC:                                   rando.Int(),
		Ol7ErVxW4ms64gxMeK5kN68vhApn8oH4GtJOe6Ys08UY2Mj46qf4P6e8fGeXJp30q854U:                           rando.Float64(),
		R0gNl5W1Y7S01O4N1m7tvE0G1hX4E5vRshPxRg38Xvjs1JEXBYT2oPLrnLLdhr3vFW1:                             rando.Rune(),
		U6YfshV7THy2wx4544U0joJW37DetTgAl2eJdSPVy5A1qkq440IK7m361t1w31PqI2D70ah2S8x30fWCcv4fiTT80AGI6:   rando.Duration(),
		X61syw: rando.Int32(),
		TxA17U6gQy7N1hRYgfPuO21D3AOa3X25LlhoYhmnWPEx6728Cp3VmOlRl2ih55axOhr0TIR: rando.Int32(),
		Ss4gkq73jxq2S5ksxdY4e0Iw631pK67E145m42jr85ok3O10t76yIebP0n3H6nJE8:       rando.Duration(),
		OSQwSib7810mPBo3Q2f2CC8kN: rando.Int16(),
		Ry0RA6Bvfdicx4jSqHTPAR16Y5cPy3U41Vc6Do1k6fWha6k3mCSUwAb74dd556N:                                 rando.Uint8(),
		Pq5g5BrTr3884lPMp5ODw38A3ElJ56ulQco8HcUC5iqTG4E6YDB7x24w1VHPf52sc7P264o10Cdh28EL7F84Bq8F88871Qr: rando.Rune(),
		PbOlLcl2ETwy20l3lknJxJNnFWlmkmRSWVe5HESpc4vaet0E:                                                rando.Uint8s(),
		F0Jy4OD5Xlq7xWObrf4taIeUjChm750mkeCcBFV38Ag61s6y0I1224847RDRig:                                  rando.Int(),
		BTYUl67hcFwI8FQj2yCHjKcLjiBe28qcUDu78QkJd62242ru:                                                rando.Int16(),
		JrtcLc53677I4b1p5JDtAD6sN8r27XB87uMpbL0gR1o7qSjnJG8Dr2XtG0RJAT53KmTphHl2ruv02yj7xMF87rrQ:        rando.Byte(),
		DT1cHS5U8uMF:                          rando.String(),
		GE4wd70NE1k823236W73GST0w13M8fusliwOv: rando.Uint64(),
		Kf7J2vx8aQ03I1Q8Gj4k57Sq50tT4snwxOCtmGBSHc0K47nLWl51XL5lO6uxRtK:                                    rando.Uint8(),
		QQveS1npmR62N6vg8mRNCtpS14H37vmnxkm6a3QSY7R85CDeuxKWOiH:                                            rando.Int8(),
		S31XMuxgS730Ihh51A1Lvq8eBWQhmL63FlYNR6I132252KISk058547:                                            rando.Int8(),
		TpX1vex53wY5Bc45T4GD1q5RT6aUb5L5qQS:                                                                rando.Float64(),
		VjxS7cluGqTaiWYtbSvw41tolOE722uJ57M6ak7Bx7KPCI:                                                     rando.Int16(),
		LICQO0485uME8y2LpEM6trLj5n32D5kB3sBTK48o8425FTVUVpf12DV438G2IOJS0YFi8A8A23OX4B0ucw642k00XxiNpo6sWH: rando.Int16(),
		AAeM0kX56qD5hD6LWYY0isYcnW6rno:                                                                     rando.Uint64(),
		E6tA6O74SrjSJCUU7oA6l7lu5wVeRQ60Vs0OHLhL0Ou4O76Vt8pMTvUW4A2gQ7sptUU0tegMpd5xxIGbr:                  rando.Int8s(),
		M21Q0IfB7eb1lo7i3rJL0YblVPiumn55LMLS33084k8nFusY61V2HI842y3CiK5283LwesJmVF2v6kuP:                   rando.Rune(),
		E3mW21T67LjqSfsK6hQYugQu67DNuHai31AbPwpLL3KlW6M04IJ1a0Q7Dc7qou356A2Ewla1mx:                         rando.Uint64(),
		SfHbj72THgHfi5s02LdHujj7I54FG6x3ki0754WQ5r773JYrdWG:                                                rando.Uint8(),
		Q3I2p8sC4s1G71kd4l64O3xY3m826eRbjBaVhGA7epdeGUuNM8UVBdXb30Yq7OUh4Q433L1xQ0U47805Dsh2fgW413Yk:       rando.Int32(),
		J38Lm7EKYs10M3:                   rando.Uint32(),
		V1Wdc2grgqnNH5u3TpV56YA32J7gBda4: rando.Int8s(),
		GAUj4KTiPfKy8b06sp7p6D52WrHN7y72xi5ri16y4cn7XK2q578u1N261HulwAFUTs7V2MwAW66Qyampk50u3w4OlFw1O: rando.Uint16(),
		Ub3c5eFlmM7x12OC0qJFQ24cpS4Rt624VXANPD65eQlTm4653fwg60eMmsF6oYHFD5c66yxKXbL1:                  rando.Int64(),
		GTl1OP2Ji7jHUYWqSXnT56:      rando.Uint(),
		UtEX65ni33Bt3rw8RuW8FmYblO6: rando.Float64s(),
		TI5CU0AW2WQdkTBh6r38C8ArX:   rando.Uint16(),
		Wq2F5oc4LQ67Ll58QwB05O10f1cVCojjg8BG3kxF1kp4LIRrswS41JWl1mr22lT6843xgOT8XR: rando.Rune(),
		E26f8OdVYlG1B6aq4Pv5: rando.String(),
		Tr1KYq0T4Uy1K24Lcq2y3BDyy2V01c8Ef7htMVVs5QGuR5J7kHfF4q5utq2Bk8HsSS30W38fp44edxrF4N7vXFKnBvGQPE:     rando.Bools(),
		ErAxhPs1e065b055oav6fqmdJyEV5r81Fx41qvw65m5EUr7Qyr7JiQK101oUIUxVdlE54dFoK72d27GD2ER5EoIvu62PF2aOAS: rando.Byte(),
		T1Wf8Mq1:                             rando.Uint(),
		O33Ew7kG81gwc3pwp57O7MN2jGSrxX5d2Fo2: rando.Float32(),
		OWfea50Ak84g43yFJXqoj7jxcah7F7s2e:    struct{}{},
		CYJ4gd7D71Lf7yX8k0MND2GNhY3tQrTsI5N8xBlW60u4aNakJ06V40o4ibIBl20:                                     rando.Byte(),
		HR0b38iSXTi427UBMHqd41a7Uw6d4h4Yku3nyp4USL:                                                          rando.Uint8(),
		OTC2L4uIAG1PPR6358V0BE637REP8u2QWwDQkCo66Kolju2TdsvPOON3Cv32Ip1F7qgr4oy7LK5v3aeaKHs8YEGyFHyygi:      rando.String(),
		CJ8h7YjJ7FNVh80yaLoq877s3Tk08657bSoJgMtcE7850421jnxG8Wxaev7swX6U4Ujw8TKMO7gVsVY2xGlQsP074EE1rHOoW65: rando.Float64s(),
		PPL4MNNL5O4XIUcAmcN6QH5USM11Ww1ExnOX264F26:                                                          rando.String(),
		UN08gmtiLIJb0hHN5wo6G4HXjo823Kn7E1v77786Xr20y3lEqm6mqC7U3B0Bb6In4T:                                  rando.Int32(),
		AY3f6ey75rMm081JexU2f1y00205t3Ex6fIRyO582B2V843b51q0S:                                               rando.Int8s(),
		Oob2TRO0Jo3CajFW3m5UH2e0Y033klWSMn26h6B2Q5UBCfb64613H731br8ufSNY:                                    rando.Int8(),
		AH3: rando.Float64(),
		NhcKkk8qm300W80Kn3G35mPi5awFuuW81UFPg5NE04hLJ3ceqPv:                                                 rando.String(),
		R8UDKJSU8167fPkTu17CxH2FdHVXBh5VPq3R4yVpDsP18DDEQl7cPpfhrFkM38G5w4QVUO1:                             rando.Float32(),
		Kfp3iA2rR8b62nh8n46IwAV4aIs0dTkP6R6E6RJqST0jd2fij45DBKBu5JiuvCnbbkpcNbn8MPfDX2H:                     rando.Float32(),
		P75tcNmsot8mqS4rwxYjkV7g5BHG8qdNF1qly8GfT3VV3Ib2p3B68RJH8o7O2m3yA08AH5eqhHcu2KaUn:                   rando.Uint64(),
		YV0sM8YJ4ntqRHl2dNuVn31WdTtcXV020:                                                                   rando.Bools(),
		Uiix48QL6HC6c5QYI653p7TLjD1160xwURC6ues631tskuS11xa52y7A5pBi7rA20YnYQuXy28:                          rando.Float32s(),
		Tl6Trh5iif2oulE84HsoRom2S5JR021IgfNgO6K88oVCx4:                                                      rando.Duration(),
		JtfT1dnDQ7cAI3O50K1jN8aG31xF4efCoh:                                                                  rando.Float32s(),
		J86L1V4ua1ehpyDjP60kdsfA5S12aIc10tfqJUBf5B2I7S7PX24o1T7p62S5OJ:                                      rando.Uint16(),
		Tf0e1MRG6Col1h2HiXmkuxJF3Y2S361umPe6N1HYKMKq65Ey01iwu620Yjrpj7m6I5Sg2tF7nn57qyIXBL:                  rando.Uint(),
		GCV31P3M8580QmvW1Xv332cuI7D7S8GpnJMOh05pX4q:                                                         rando.Uint64(),
		GWyS13GXAp4Oh52eOn3wGoL5c7I27BF3VPs1kR76b3fil:                                                       rando.Bools(),
		GnGV51md55qe71RRUa2RVj8QpS68C2b8UqjmAWyOl6Ymwk6a8oqQe8hcI8d4hKFk0uwdp:                               rando.Float64(),
		FEyA0EkyXlx3yI5JORj5o8IF7V0Q726pja7080D4qW486p772AB07daMIGnpu0ik2kjXnh4Sl244d:                       struct{}{},
		QvNQQWYI15FKN1y8Vuv1G235hrXd1JVSQFP66p5PJQ6H0vNapIcXvUaKY2hHC10XAKc7Efbi2N52y7lWV3:                  rando.Int64(),
		H4wbF0H83jCypUWqy6b0YjL6cku3M2Oe3c5S7PT55YHge80XV634Dgqa5a05uc3PIdc1o5s1C8tPSpK37dhRKx5jHL5Kv32BJL:  rando.Int32(),
		Ho588cd4bslL53IdVPpMmECJHoaJg58C4Y7:                                                                 rando.Int32(),
		T2405J00Dsn6Qjjr6yr5k:                                                                               rando.Rune(),
		TN15K1Fehm5JIgTfKN7GO2lQlvSII0Pa2nMhYo2oP:                                                           rando.Uint16(),
		Rmm5Jo6pf4hM4v4Hf27h14xFWabhN251xCVMr1ffsV2RBHoL6ine477PDc6Dl0:                                      rando.Uint16(),
		QH6ywWNEOjsi1ShF87adR70nPS3Rb32Ij1bJJ7q867LPCJTfKS3h:                                                rando.Int8(),
		Ev7r6CQjD110elh3GRUNaUNj40l36Y81d4m45FV3MaoVl0u2sF48e7nJ8S4rkAHJMvlU20sFP5oLmM655BOc3NJ:             rando.String(),
		S11T12a5UNo32q8NcvjtJ7o5C1SuERyXTh4c1gY3kWn78HmSJNh5xBe5W5j:                                         rando.Float32(),
		D0IE544n4aN450Dy46hpeT4LxWihLs1yDhbh78gEDO136K6x8YN6JVfgiOa4fjKs72k1mTj1j2eL52uD54L0h06Q76xGBeCgyc0: struct{}{},
		XaTRl2j757754Nsm4HIu6nsveTiwL6wwY4F84p76n4MPuw45h:                                                   rando.Rune(),
		TW20ixrlWY6E254cp5eCV:                                               rando.Uint32(),
		YgMouab4ATb0Mr2aAEb32Pum15M8853k02483SBSMEKj:                        rando.Int(),
		WIdcImvm6OV7OVR30omY2gbNQfmJduQ81I451it4238j6q8hlU0B14Kx66DR5opEJ52: rando.Int8s(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, QQXGKCa4tQi1lgGl81IDkGRMc3627IiA3Fh4aC6y3815PWv684S{}, expected)
	require.NotEqual(t, QQXGKCa4tQi1lgGl81IDkGRMc3627IiA3Fh4aC6y3815PWv684S{}, actual)
	require.Equal(t, expected, actual)
}
