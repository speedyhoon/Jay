package fuzz00

import (
	"testing"

	"github.com/speedyhoon/rando"
	"github.com/stretchr/testify/require"
)

func TestFuzz_0(t *testing.T) {
	var expected, actual B41sS07udGL5tMTa4N86u6AD5d21324W5nWkWID5m1Sk810sN4e37
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, B41sS07udGL5tMTa4N86u6AD5d21324W5nWkWID5m1Sk810sN4e37{}, expected)
	require.Equal(t, B41sS07udGL5tMTa4N86u6AD5d21324W5nWkWID5m1Sk810sN4e37{}, actual)

	actual = B41sS07udGL5tMTa4N86u6AD5d21324W5nWkWID5m1Sk810sN4e37{
		D4TFwAFiVkiF: rando.Int32(),
		B8thq1nX5wWv4PBh3d2uAL63vWFp8x0FSdwDaqOCh8DT1jKnf0m6j:                                          rando.Uint8(),
		Yjor40POR2IM5q56X36856ua764EaG1B48Suq42LMi5TrW:                                                 rando.Uint16(),
		JV6F60EfixUb7tHhr6yt2EdEsyK83UU2a4qHsS3l1q76Gx1AFiwQSf3h1y17X0cec1Mc6CQy:                       rando.Uint(),
		K5UBtgqPyvjy87FMrQMH0cV8k65R0nl7mHb4J5sn3oCR0ki76NKQ3RcY1gEUE85YwlrI3Q7T4Sf43L24H8fqqX4a6C2rwL: rando.Int64(),
		K6yL415FMC7k4aR7YA:                   rando.Bytes(),
		Q6n5vh7T6PH62h6hp8o1jThBh4v2Yiq1X70J: rando.Uint32(),
		Tqy48H5m41bwg3Ft57:                   rando.Bytes(),
		UEN78c4d6mD6eV7263KHaW6TnS6s36:       rando.Int32(),
		JTId2g154021BbYS0OCa556Qr8fW84nL771yq00bS2lcVtR7267AbTvgK1tmIf1au: rando.Int(),
		Y021g: rando.Int8(),
		Ar3U847321f64gBe6c0h7605l5e0V7he503WEdysVD2daDG8yLEQ6X7B0Mm14: rando.Int16(),
		LRAG02j481rgrBkPPRvdI8EP8vLV43a88U6n72x77gK2nmpOy2cJAHvQGVON:  rando.Byte(),
		R3b0if7HswFBRqEX6vO8CqT2504vQM2sH:                             rando.Int64(),
		QM4VjjJsw6l24mq:                                               rando.Int8(),
		Sf721e4:                                                       rando.Uint(),
		PvmFkVXkJ54Dj5Nf8je0Ips:                                       rando.Byte(),
		Fi8AFMBnRQMFlyFKbXhwMB3244CNa4gBC751e0n3H55q7512gf8OSwn3Ot70Lt33XlHGMGYETch:                         rando.Rune(),
		KSuCCpEg10b7CV520d0NvloMN0CiLrr5bNw4Om7kTYK228ngbNX6WMi0nxy1L5Vrg1yBmE1e6N1m2eCQ5M632B03eMOHGMW7m:   rando.Int64(),
		BOHOis0EaOm0q3LpGG8uA8e2u72B5662u1qHqAsVY2gTGy7A84JG84g6D1fEjBhcmsqvKR25q7f2tE4:                     rando.Time(),
		P0mS700R5yML23L41ySusJrfQ7pTKpO5FEcWAoT6JC0P5LJ2LL0CU0jYI7S8jnG8SV5vRWiTh15hAi2K3RqD:                rando.Int64(),
		J4u8xUqqES4X0rF1XsBj1GHmhq82lur7V80w50b25378PE6rN1MRSkhwrC3m52LqdD53cMSldU2jU:                       rando.Int8(),
		CQnBsPjD3JLyy3iR5R11JqjbThXJH43l6GDS3mYU8dHlLIPw723:                                                 rando.Bytes(),
		Cu0pwahPF6Y2pTNRykB8t6KS20Rd2md8u3:                                                                  rando.Byte(),
		UM4org8M2e46Y7fll30CL2F11sATLn7D5RtS64lgjxXs:                                                        rando.Int64(),
		G6om1UlH6pjGm85C6cW77RDsSys4LN1A3k5CY84vq1qjjo0lw5Hr6uWJ2rEhY2B2w8DhCkM1RR5hJaHJu8rawC:              rando.Int16(),
		BSQqSw05jR7swBq4TipJ2cB0VyJE71YHwitqTU1Co3Trve1o0Lt10Rmn3gCm8hj77IL2uWYSgU1vX3lrF4EgAmD8IyRyCtQOeW4: rando.Uint(),
		T8F3XJhk4NcTRt4fkrO:                       rando.Uint(),
		I3b2mOpeuF7P8MV2eI2b8WQc:                  rando.Float64(),
		Oc8J85I8e37PmCUS08joQLbG55sp1l74CnMI4j0MF: rando.Uint(),
		KiDvsNufk5j11OhF:                          rando.Int32(),
		U68sFarB5DB860Tc:                          rando.Uint8(),
		XpE8J0:                                    rando.Float64(),
		I6wpx4504uQ5M0vYgpIRpSyEP0mG6O63K3dJSNIMkg374kHpiY11nYdw7: rando.Uint(),
		D4WG7U2a1pfQxKS0tmOuY44reNg277:                            struct{}{},
		QKAm:                                                      rando.Bool(),
		I0mS33qcGb4PbM0cd4l0KB75W6EuhHbh5PW2SpP8n3qVOm3YUn4AQBGFC78BEf51K10GAsn220jpl7sSQ8t7YNr4p5Kl5r3: rando.Rune(),
		PK1owNFI5MEWW072sC722mmmX474A2mySL2qENwk51GoF4275f32K4U5LOA37eA05ypUiUs3Y0A:                     rando.Byte(),
		HuvbWI35u2jK1pg6:         rando.Int16(),
		FNpn56v6ju2714R5IcYFxre5: rando.Uint(),
		GJ136cqkFeq8d4Q1l6251fijHC488na1SF78OnKE8NKSTgFFaSX1u0UhPvIPa4Ju7EjJcM1RfqeJVUvxxqIEYTF6WD22wPHVRv7: rando.Bytes(),
		RMmO0myP7UpKnS3bLCge87vD1dcjM0tH13y5i4q1doXY1TFK020uFekAP3Pda2hqMfhX3Co5tTL2G:                       struct{}{},
		HP35uU2QeTI1jN1I6v5pkgrX0LT5qi4885jhT3n:                                                             rando.Int8(),
		U5C8S3pA4n:                                                                                          rando.Byte(),
		GOn5w5ndrLC4Ude1g4DTD55X1R8CQvJe2w36vu:                                                              rando.Time(),
		PKalsb1WDm8tXM24vLlR7NOha2HKo87WrEFihCy:                                                             rando.Time(),
		Re2EvS0Waa8WQ72t8QncPP6IOQc4W8E6R2G0Vyyh8pA6M0EcpG7dBv8HL07uT1:                                      rando.Int8(),
		Em2Fqx46gf5pvY38G2aqxFI8iKOS23Xv3:                                                                   rando.Bool(),
		Sa5tW7101:                                                                                           rando.Uint32(),
		EScr3UwOJQiImdxefYRVAw:                                                                              rando.Bytes(),
		UD7441UrXldhqqUV264N78mhfP86Xp:                                                                      rando.Uint16(),
		A0R8dV23112gnmsM1hp8bnUBLya05kuQu7C:                                                                 rando.Uint(),
		SyPH1e82cvq20PML2c7I5VY2JNk4L75sDm8PqD75l8nm1B7M7El1kvo7WIwRiVkgO0MnT4MkqVex6T2bX4p6xp71hl: rando.Bool(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, B41sS07udGL5tMTa4N86u6AD5d21324W5nWkWID5m1Sk810sN4e37{}, expected)
	require.NotEqual(t, B41sS07udGL5tMTa4N86u6AD5d21324W5nWkWID5m1Sk810sN4e37{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_1(t *testing.T) {
	var expected, actual H5NObK28prAo7234yN0Rl7KA
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, H5NObK28prAo7234yN0Rl7KA{}, expected)
	require.Equal(t, H5NObK28prAo7234yN0Rl7KA{}, actual)

	actual = H5NObK28prAo7234yN0Rl7KA{
		R8gy3R0aEO67d533FxGy5Bq4Bx0oK4cG4x30FMC454gpE1ix8rVNt4Q1hEuS536W6wX37X0RM: rando.Int64(),
		AYHUQ2JNrSF4vV45Leg6tIgiD44F: rando.Float32(),
		T08l2Yp6H1mq685Re2Fy14doVocQHxnh0153S6aE1AQY533wW6KGnQdAwBf0iqD5Stlsdgc3G2VUPljQ4jwlWYptPqBjR:   rando.Float64(),
		DtS74V6Wnqpa6W2TVfXilKR4P34k4kYN7rYsxDG62jpVa2t21PV1JcAD5QSScF33h8x8M8d8Ft3xv8U057ddrLwC2LcF1X2: rando.Int16(),
		FPt8A3wk: rando.Time(),
		Ofqsno01YWfh6j807Uw0720V7Le0lh74532Pv27jcEO1A1E34oGoDkke2Nkb5Y8s7t8CS: rando.Time(),
		FrR8k1m75763e5758RV11hJW: rando.Int64(),
		U1g06oJmN5710X3D3G6cKSDCRhc01v8rxVxFJEyiTEy1l035J1Ky2OX716r8G55o: rando.Uint16(),
		J23MdwLsM3cVNEuX5Uctd7s0F: rando.Uint64(),
		PG2EJJ0o63Sc22nJG2MfO274BW5APQkJkPuHiQX0351VaPIrAqTX332haGhN1gehi6XCYdQbeG3pqeF:                  rando.Uint8(),
		Bo4dwTt27l2XUimL77FkK6XsrdjMq60DtdbbsaxDAQv0L28Hk7LP0OMLJh5K0h7HV3D8grd31SAkgx3H6:                rando.Int32(),
		PeGra6k1Rf7pqRIKc4hvSkjcF6i5dO4b0O3Ixld85F7xN3t:                                                  rando.Int64(),
		KHNb4hq36I3BwqU4aQ380HaqhH2Ug1G4pUp0cNc5:                                                         rando.Uint64(),
		DV1g1mp2kfQde3J3tr5K0at1HNf0BR1SkC75Br7JGpQpIGm2nQL32L565XWWSbY2BvM2P2W5py7myy6SAL8q6eieM1NAGc07: struct{}{},
		VKttnr731X575ofosuA5m6sAhUle44fpUjA3KAY47Ourxl5Ay6vuXvL3Dh4BKIhX1NqhT:                            rando.Uint8(),
		Mr425y1Of5b5x6Xi4k8eIaqOwemFj:                                                                    rando.Uint8(),
		B46ywxj5jVSP4xWl20Nf362:                                                                          rando.Uint(),
		I7Ofg7777D4vPcgLgKLN66RVe8:                                                                       rando.Time(),
		NVA4F0vN06W24:                                                                                    rando.Uint8(),
		YhdX17r6j15DhlrSge5C8X6mGM0AgtKP6iY3tvyW:                                                         rando.Uint8(),
		AJ83485R3aTogdUB3eHG4j4827clqD1LM:                                                                rando.Rune(),
		QMAUuUk62BLuM1YQsGdV4O6I3165wV0U2:                                                                rando.Float64(),
		J6fea772aA74jhuyiX8dXoIqc184SbW21R6nNeeDoExKB3156j88HXal:                                         rando.String(),
		Y0M1J1TuDA1FhJr64527SoXQ788bk12KS04H53J2IK366PC86H7618e0HHpi73NU05u45mOnaM32E415U6u5YY3eB:        rando.Uint(),
		UE30PSlkL3IJX2fBb0jePC1jR6FlBq0gc002:                                                             rando.Rune(),
		FidwHV74TpnlQ4EuEBD3WF13l7B8RW4w0PY8n8y3Mpp6d0I111B5SbOvLQsjEay56EWMvS87227:                      rando.Bytes(),
		Cr2A2Oe27yh35nALiKr1GXQPKG81ec6q:                                                                 struct{}{},
		N4NqmlI4geBhuq67TFVjCm24r2f10A68XMmfI4KwC4rbl3Xd2Eu82d286G57oG7Onk8gV7s8h4HJhNkhvguVEsebD3TP:     rando.Float64(),
		S06N47JRYJY0uBMK8u71F0CGR6QqlA1j40CPe3hYJ5rt1BU2Rg:                                               rando.Int16(),
		R0JaBhj1LW47OYj5S2yI1eu67Pm167W372PHkxQ7by31D0:                                                   rando.Time(),
		R7KNFAi3IYJmNdFS3NnBJYw08Wg76vW43163N2x3Sv35xVwe3W35dR4mM50Cah1H33C673e3L6SfoQ4NLP5m6j2w4:        rando.Float64(),
		D2G5Xj4BoqCA1Lty3eP2WK84phYxvWq7fXus4s3gjTupQ0iiABDa:                                             rando.Int8(),
		CWWaqLVOs84Ow: rando.String(),
		Cs3801HA0:     rando.Time(),
		J54T2KdRE6xj76UpeGD168Se6167ybhllkiJEQw1CIg7gU3860L68lIiIn8rkf66wC6U027RxV:                   rando.Rune(),
		TIVV4cE50Dq3kuU5SVyM7GLjv83423KcYY8sG3miyK67rvq6P8T63h3kk6AXso3XCG1KA4yGASh73L812tVuv4G32Y0Q: rando.Float32(),
		JnknIRe642Gu5TMs0GSy1vaVUv2dfsm7o7:                                                           rando.Int8(),
		Jue6OlALkoQlP:                                                                                rando.Int16(),
		H0K3vgM6DjRFQKlx7884M8wei13j1W7rEb8t2LmH15g3effm0CVqV36Gplv1667Gabu25lbAr46b4:                rando.Float32(),
		UC3Q6IkG6o3ixjFY3K2Xnaq41FWN1Fihd00RQdS332:                                                   rando.Byte(),
		Ge6MJ88Y5NSP84hM50tl2525j3x842V7SME8qH2juX545i4otR1wk54n:                                     rando.Uint8(),
		CNR4e8Wa40Uay4j4LM7fMHkVaxO727fi28q4ST1224C130NW3V03P:                                        rando.Float64(),
		Lp2fv3OE6puBTK32T5DUAu07W8Rd2nFpmYOf4AJeSux30SOp5f83rslH3n5TXU1x7H:                           rando.Bool(),
		TaEI:                             rando.Uint32(),
		YH8:                              rando.Time(),
		F8dfCe361:                        rando.Bool(),
		KT1g8jq0MCJOL2kadt8wKXwL:         rando.Rune(),
		Mr08xmf5QfF2Ievj8EL7:             rando.Uint(),
		IjNrD2uH2E1e8v2yUV667SyY3iUlS1dJ: rando.String(),
		DsI8eybp7T5CEPL755eh1tu3bOXV3b1Cde8YfUFokR4Gdi:                                                 rando.Int32(),
		Gl55m3rwQ47W8tt7Un160n38t15JUOoyV8ngoengw1rwbAjPa4h60FcEb1R8qxI7U2C72P41JR857:                  struct{}{},
		WTCyr6c0Sp3HY6QV0xnegUH18FfPchuL41313rTFNy0picK2aUY2Fu0RB64Xgy4Q1qgs2iN2E3HYhDqen152P:          rando.Time(),
		IS1F7mbQ8e26I1f4JG8Y21XOjq1obeOP5o7XUO11G626YO3I2iPbxy05NeP4oxe8E0s8IFV7lS17yOF70m2:            rando.Int32(),
		F5Jx52Vs1Fa03qajVVoEVJFH0rx65BT416LsFt27T1S1f:                                                  rando.Uint8(),
		VR0i6OEDBTxtLnEMO35e2Kc4x17lIyPk8w4UPSqJ4UUEKpGiU6rh3v52v7tmw3ebN5O71TC6h32S4r:                 rando.Uint(),
		R8aWPY85W26dr1H256lxdwY6cgBTsAjr6cxw4YBr7D7IALSQe63om44I7GQ5ALAtIVAWiQAy1y1RG32FtGclXCI8GCEe8R: rando.Rune(),
		RCmphFNn4nMEEo08Qwfb5U5UQLTvco:                                                                 rando.Rune(),
		H48mQjWoK855N8j74isYp6Xxvl5084xSAu8:                                                            rando.Uint8(),
		R6AJ5X5NmBfIGle1H552ey3UYAV6wk58D6sbPcEjybNmTeK2846YQvhyml:                                     rando.Rune(),
		KBdlECR7HBKlqyFmmJ8NS5500h4bCXNwjlehu3TmaaPIUCW874uIcS266W0W184MF:                              rando.Uint(),
		T0MN630Alxh6d0k73PUA1RgSEm6Mi3nn7Xs6UQ7DPXGS0e8247XAa:                                          struct{}{},
		Gfk540r2rohJI71mFmvXEU7a825GX1r6fs04Qego:                                                       rando.Uint32(),
		JpEe7817qP8A8FTaikW3EbwK061l358YpGkdO1i2RCRVVm3uxB84T77vsGQhCvL0Ch4J3h0:                        rando.Int64(),
		IMEWe1OUw: rando.String(),
		Nivl3gpUNVt8514HE2EYP1UN2N2X1Tlf1F4FUv6w1Qv1OFv50HED4Q4t1hH5B70oJReHan6Kxw2JvJT4dKVHP8u: rando.Uint8(),
		Y1jLu03F0U36bx2fUN2AE41D3INXQ6474UUks22D84VU2Cj08y38YHTo4:                               rando.Int32(),
		VLY5m0m50tlAW0uXuUBQ48ynU6x44:                                                           rando.Byte(),
		A70nTDY:                                                                                 rando.Int32(),
		RSA23J6Y07aidk:                                                                          rando.Int32(),
		OIq85sAjUjkO2kbUDPP7Exm5C6H7xnett8xLi7T6UfQYjbaLE3vV:                                    rando.Uint(),
		FStUI1evJ8IMpF1FNkFUAxKcNp5mnPBKk5wve3813xr6Ur:                                          rando.Bytes(),
		UPSKmy5fY3qf875AnwWUt037H10If12s57kj703I7EGDRTJPn07CU2O76r5rvdRrHL:                      rando.Int64(),
		OY2B6MR13h5Xyl8dJL3o68LTD5r04Masc54445LWfBI308w8WAK1305u:                                rando.Float64(),
		Q06AwRq6E740I8OV66Sj4HCcI8tLI6vq1b5qtEj46T45O65Ms8yME4Fu2M3RS666176O:                    rando.Int32(),
		A56AAxu0U5MqKUQS5QlM710Y7b253iy70VNjsAiX850hTXax2ByYWwiYDMj:                             rando.Int32(),
		I1xT1MyLu5: rando.Uint64(),
		MN8TEqDPN45TFWk5XRMq8spGw7m4610X4WvJuM624mb3LqG788F0h766Q053db6B0Y04ldxhyB4MWXRs: rando.Int64(),
		WKVXS3uQKnVfe0TK1lSu3j2LPfrYkogEuc1nWJDqNdSQ5C3Dao4ojE15AdA:                      rando.Int32(),
		O16LkkMpeTm261XA: rando.Int64(),
		S3GXR1aHM31le8Ov6vxgIrR08Kp3d0Af62oYIipU4N3Sxp313bMkrvUkfysxrxau547L708:                          rando.Rune(),
		Sgd1mA7JWraTo47RyktQjQ3IW15LU5Y2mvg6:                                                             rando.Byte(),
		Idg7p60Fvef8kTw73Vr4W36M8QA546Ja:                                                                 rando.Byte(),
		Vq3S2PAE4EoprVn636LlpMN71tnafRqGLV20y4v71j155l0f34rfCYI4Q3JeXu4071quKU2khKfp6NuPEXWd:             rando.Int16(),
		FdWU2mK76Wq5F8IBEvHONBE0m3lLC3lr36HXe2JfBGBe7xJXrDAWY3a1X4cJq6dP1evN8DdyNXnnxFt:                  rando.Int64(),
		OT0vCT88Mj6yI86RO5aCwML7GK:                                                                       rando.Float64(),
		BL4AVV1U1pUT57u4BE32hmKk6P5P7qun5W8f13aqY5ury3W407m4g3AL8BoMRYW04O:                               rando.Int16(),
		S05Fg4J6Bwd6kq612M6pT6CEE3OhtKxM8c3vNgmMb6Wx4tj1uTIvOd82LxvC1LUg51jA738wd68qc85LTj5E0p85:         rando.Int8(),
		Sy8bkk8El0R7576s67g6NFpEF4hfMTw67aW0UYIbBQxQ432PW7:                                               rando.Int32(),
		Oc7JcP2jTBwyfv54m2G1sBc4CMR3MnPu5Pc587qPn2VGo8dsmi54m63TR7gRM0d7D2BF78Wno73XyrW0njo:              rando.Uint(),
		GJU1n3befn4LHR7Sn5Ak2f4n:                                                                         rando.Int(),
		Kmnlvv78XWs45dP548GcGfTqt78o0uQ0XF0507XKwrpO5V106OJ5IVFkk6Jh6F177juG7X64jSI5:                     rando.Bytes(),
		GIJP6Fnb4kr0opMyg8vBb0Vo4vkgE06CF0qj8g66Alf0Ucf1JuPquDS3x1uWQ57IHb7wo4Wyhq0thhxLffA51JH2aQP0ojP0: rando.Float32(),
		F012485rRxHmoANDYX5S4m6Ci1M0Go5a2dgsSetcUG7Gev3644F65GnxXA1QFLS87QfYnh:                           rando.Int16(),
		G6JrkpJPqP708KOjsW247a0WfSfuMADe1X5dby02c52MBcpN3FoS262J4lg:                                      rando.Float64(),
		Q6lheL80vV4Vuej88Xy8d6x6P0AKar2cM508V8tj7ll2Xlf1H0y5l1pRfKtJi:                                    rando.Time(),
		Q76B6y10XEIAf2ryo23VyKRvd0uUUiUJf0W5t7073t0P7Xe478lfoP5JQPHhB4QV5ycNORc524tEslp:                  rando.Float64(),
		TqCYVSEpcGO7C3Oil4H3tqM63HatbC67aY2viL8p6egJ:                                                     rando.Uint(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, H5NObK28prAo7234yN0Rl7KA{}, expected)
	require.NotEqual(t, H5NObK28prAo7234yN0Rl7KA{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_2(t *testing.T) {
	var expected, actual U0liUQXFx3QJh62G31JY5vGvF2Y52cmOK7g8jvV5NXYh7xfmjQ000MmpUF81g8
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, U0liUQXFx3QJh62G31JY5vGvF2Y52cmOK7g8jvV5NXYh7xfmjQ000MmpUF81g8{}, expected)
	require.Equal(t, U0liUQXFx3QJh62G31JY5vGvF2Y52cmOK7g8jvV5NXYh7xfmjQ000MmpUF81g8{}, actual)

	actual = U0liUQXFx3QJh62G31JY5vGvF2Y52cmOK7g8jvV5NXYh7xfmjQ000MmpUF81g8{
		LT5HgUf805OSxc073t1VxxIcTQBOk72r4Dn0C7P1:                                                             rando.Int64(),
		KxtkQjDKvervUbC43Q0b4i53r7vJ5t2rQBHjVp81ov7m4lyX0Txp5N1qhJPjChy3gVGQdaYR4LTT6Q51Kd:                   rando.Uint32(),
		QL7DYx1WFpL8Ha481VB4251Br6eFs75dx46xyrkwMTm0eacbt5YgrFs0Ue7UfQdDSTm7X74:                              rando.Time(),
		Wj340qgcd6J6sgG6RbbF1Cj6HTOJa8BhneI0PVeR2wpiHBft5Mpu3Ln26oAs24lBN38WABRE27XL7C72DsrSaI6Y6768:         rando.Int32(),
		CnB3YV8wQwVgjMv3dK8DGqL5um1SH2sOgPYwteO8EPR5qD3xcMv:                                                  rando.Int16(),
		HCryS23q8R15C4OpeNj1Wsf3EQaSR1kw8Ct473gph5IK48UtTsiAEe0EnEQdaH3O263122pn4:                            rando.Time(),
		BRm180iIjsx1Mu3iV15O70k56iI7360Q7wU83I2v7y3eK4b5A:                                                    struct{}{},
		IqHV4ax78257pE1qiUH17HwJp06a3fJNs6Insd146iB8iQ708OE64I82KXVk0yMvr042qnUvMu4Rdck7t70Exw7347aKJAA46w13: rando.Uint64(),
		HbOnWi7N5c7H0f5m1N6xKEK5s:                    rando.Time(),
		Vdwgh2Ov5lY4E1S006bRaOs2L301:                 rando.Rune(),
		Hmn0fC7ScPm7Bh4XI4uVipsrm0xDkYs6N1:           rando.Rune(),
		Ic74Lfa1dNVUyeA5w3UGEUB7DkDQ048PU3SD110PC28x: rando.Float64(),
		QTSNHqE3Gij4J8lpFs12c1fXB1fbBw41Wky3Egegr5t3Wc8qtVHskapc3xt67H1Ld40oWASXQsK3SK36d06e: rando.Float64(),
		RW0I1dX57du71w36bCE4N53oKpJP6: rando.Time(),
		QeVv81r0OrGi04235lkNT87Ba3c07Ja8fCJtKw84T0405V3Vl0O1Tld8a1Jk6vg8g4X4rn0WB5usc8wGij6qMrw1:  struct{}{},
		M782fH6C2GOtQ41wsXk74bxrbabP5RgV1NO2MxHVC72J57:                                            rando.Int32(),
		SKGIO2pQcB6E6yQT2e7531GrP47F81C1BHN6i2b4x1OmlWf4g8AB6mP50uKMlTe5Ok0Qv40ejb7M1YvWJ2bVRF6qL: rando.String(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, U0liUQXFx3QJh62G31JY5vGvF2Y52cmOK7g8jvV5NXYh7xfmjQ000MmpUF81g8{}, expected)
	require.NotEqual(t, U0liUQXFx3QJh62G31JY5vGvF2Y52cmOK7g8jvV5NXYh7xfmjQ000MmpUF81g8{}, actual)
	require.Equal(t, expected, actual)
}

func TestFuzz_3(t *testing.T) {
	var expected, actual ORjBp8Pg4Hi73lc04
	require.NoError(t, expected.UnmarshalJ(actual.MarshalJ()))
	require.Equal(t, expected, actual)
	require.Equal(t, ORjBp8Pg4Hi73lc04{}, expected)
	require.Equal(t, ORjBp8Pg4Hi73lc04{}, actual)

	actual = ORjBp8Pg4Hi73lc04{
		B01DI1Or7HDdV4o0RI4iMLEo15trJ635Xw6s6nGb3oxwGYY6t5bbmIQcM047y54Uu3ta6HM083pAUMvxPHRJDF36i0S7wVL: rando.Int8(),
		PI54RW52X1Mni52Skt30A64Mj4smx1T7FxVOP4pLS5T3U010ce4EeUOPvNDXRD:                                  rando.Int64(),
		CqV5G15Ka13U2S1V1hS4yoA223G383W2WRMdgkITIRb76LC3lCbyAXFlQFAHrdLW48:                              rando.Uint8(),
		G70MKO58m37LWGxNc5Aiv4icSCtH7MNkko1y3WMl63LXq5Ee2bV6eANi5u:                                      rando.Int32(),
		WhQy8yx1pWrKGanyRahHlTprCuMtTbm1fR52n8hRYAK64rgohI:                                              rando.Bool(),
		Rg: rando.Uint64(),
		Cr8KPrjn6rNyydH7yyYVoF7bWnkV2o46fyUTnvgBER:                   rando.Bytes(),
		G0UglxCHl4756dv54V628C15XfNs4E007c5oj10R36Y40NqTjD8MlwBXB6Oy: rando.Bytes(),
		Dr5EXJd0OX01GPvUSf6rcSjqATa4GR7e2OmF13kV4r1HnAiU8imOQABte:    rando.Int(),
		ArC1: rando.Bytes(),
		Bk8IoUosEYvigkRFut1i22S54x67ve4rmx1tjJ3EUGnUU4Vbo04DlI4IOtdR875xB48BI3tL13o:                     rando.Bytes(),
		L16Ib2RBupPaSv8vooE23VW7e38Mn6b3Xc2aKS8fIA3xlA03klPh:                                            rando.Int8(),
		SHWgP2Bb8WvC7v1Td0UjFVq2D0c3Ae2Ah1Kre:                                                           rando.Int16(),
		B7iXM2tp3OYhjfs3662vEOl36FOM0a5x3:                                                               rando.Int(),
		P51ouY6e4HmT8KWeby76wWKB8WtGApo4068DFkMXbHJy4YL50onHWs4J871r25tRdmk173QEoORK0no7y3Tl48t6l74u6W4: rando.Time(),
		Ke3xV0T5P0O3ar4siNC8hotWoK:                                                                      rando.Byte(),
		D28L2K2y13xf6shLn481SwDxoEJt5308COplG0pESm3GE41f375PT61uv:                                       rando.Int8(),
		VDcgINoeXx08lPEbd51w7qh4I04Fad7Mku26PbsEAdMKlv6AXT4RCKDgcPDSil7WJJwFr25dkBSn11VrdjtnHq1:         rando.Int64(),
		F04Dwxt8g8pIF3L0bJ4j1Wt15YAmpdFsvYSJstPEna3x6o8N07NL4B28u4BADcd8561kYg:                          rando.Byte(),
		GT6YdCXQTOuR53RoYK: rando.Float32(),
		Ni08MLj4X:          rando.Int8(),
		PA6XuQ2YyyCJ4wuI67: rando.Int32(),
		Dk61B3x6Lrc2XI6X0Ro7bY0LY0FL30JMp7rV82QN3E3VLiWERV6h2U7L8CUmFjJX0r5x4l0b2KF4P77xKkS: rando.String(),
		QL8JI5I2p5y3FMJrYG8k0wN6L57P7Ki5V62315v470SpVra88QW4sGa6m653y51Gj4n6S:               struct{}{},
		UR6aRyE8Vg6akN3XpbOPggvSC: rando.Bytes(),
		Usujl8l2jk5DN55AnBNmoq15JFGrJoCfcIb8OSodT4UNJv33bM4EjyEsuC5rGtt7WBtSHan8I34eGCm4vVFY2wN4N0EtH0qXW5: rando.Int(),
		CGT7u8H06xNcW8IS2y0rr65VweJmkX1IDf38Ngf66NP6IUADYde2g3tqPr0kdY417nH8AKeRBuan1GnePPW5w6QOS1u0ORY:    rando.Rune(),
		E4OpbO1pa5cp2364EL1k25NQ8WsoCM88CrXf3jIPeMcDXwoUJ5J4dCiPJy3CJ1Mmy0UDaG38IlpK8:                      rando.Bool(),
		I3I8P: rando.Uint32(),
		RFf5rotTB1a603RP25l7Hy42400KUEnGXoO223iGadtFhqsr0tN8AM8e7cO2w4NPmXGEN5:                        rando.Float64(),
		KHTH4EF8084GjIOOD3wW5Jc2EFYQ7vr5bST3ylHK54DWQr3m5el128d:                                       rando.Int8(),
		De5P35VH7Ne6rSTyw7MaT71r2osY7wTeGE68JbEOj016Gr:                                                rando.Int16(),
		SbDi51X1K17uRY503U8G011kc7DGJxu5H8mnElS8sB6:                                                   rando.Int64(),
		AYLQH8FJaglD5VnHi1hxwiiB6M6Vt1LslX6g7raYu8VdNqgqM43r8352TA320f:                                rando.Uint16(),
		IRy6wV74w76D28j3Hnqo54NWjHKaR755g:                                                             rando.Int64(),
		Ivv7tXCXbmplV77DBr0HVSF187dc5xwdTbY4eAb08E540TYDpQDa5Q7:                                       rando.Float32(),
		QMcmemd6xwqdS8fGsv1syTfDKmf2kn37y8IjpwDxv1q7P8IPaxq7aPW531H2076cy0a5h51432N76hjKxbbQJ6eL331lU: rando.Uint(),
		Yy8R2Mo5H18n1SX3aWyJcT8T5JToI1vPRysv3LM8:                                                      rando.Int8(),
		T1sk6JgtQU0w02oUS16sAv3C6SnjkO05P7XjfNPYB8t4DNJbDCiBu18HwSlQN42Ut7dpODYDW61nYsA54g2sqU5u:      rando.String(),
		TN22X8Ks1pkGXdC6vEhQxh4D4dWVVYdU0UVJPF4qDXy1VSo:                                               struct{}{},
		LlC4MJuD8eYi24mWt14waEPv8S7bK6VK7q23sX6y66XM33O72UvyxRBY3wW238T5:                              rando.Int16(),
		Ro0q: rando.Uint32(),
		Wk2PGYri6dhsgT6SAXw76PK3sYCeks8mYa20WYKWO:                                                         rando.String(),
		HSY23ooE5X44wcsH28k7Q4XDU8v3lu5o556rDVn:                                                           rando.Int64(),
		P3p7W6t0ua1Tm45FtHC46175B0Jt52BOaF7lATGM16ca2qV2dS00O1IH14ftXsN55NI3NP3CHDY0bqW7d47y6Q5S6UVeWv5W6: rando.Rune(),
		V38V5X0egH5xx4d2QRNG04C6657yI0rcOf5J:                                                              rando.Rune(),
		PA8IAgYW7mmqL65g8C74xU5UkcscyocHVA3Qv457hI3qO61RP61Hft24yDEJ12eLAhckiA:                            rando.Int(),
		FM5v2A3Uw3QCi6SHfBs254JN7dqia5aJVJl26S550pyabo6yf73tK1f:                                           rando.Int16(),
		RO2kIB860oSCPB4wF2o85Iho72Vc33uDRLUTw8lrd5gJ6R0W6G:                                                rando.Bool(),
		If103qdbE5X18ev0uEx1e83S0f4:                                                                       rando.Int(),
		Ae7UF5d4BIgmEUwvw08HtPwXqj1h7Ynj67o3A8uC6OJ0s8Ajm3ekDHqJI35aubj14U55Y0G4731Kj3U486:                rando.Int32(),
		Ujw0oMmuR6: rando.Rune(),
		XAQ4e7sKpwBYyoStgnSQJdc0e756FvnYuBj4L6mg2t4W1e:                                                      rando.Int8(),
		OS4lwKcv3oE7D6Hd4XPc3rjrJXkGCTp0RSMXtRdI3jr84w23qky1iE2XkTKFKgU3oc380CKo51uo1tAh7G8VKq3:             rando.Float32(),
		VCc83QxcM2L2OnOy8606YwTyb3e8S0i8b5SEX7m1isqI1IXe2Ur:                                                 rando.Bool(),
		Mh7yv7Wn56euvYNIbrvjBIiGflP05CT0K1u42Kp6edu6kbahGaPo1UytF0JU8XB4v7SdB8tRj8g72n1m7:                   struct{}{},
		R5X8f4iND838T2Sx0CPkod16i6L06421:                                                                    rando.Float64(),
		I47GNaFf4bAlT1sgg00rV3SiFLE4:                                                                        rando.Float64(),
		ImOSOC5WnfvFfNr3F55e7jCq1wAAbL6nO7qL57i3iegwce233KOscVRXe8Yxphp23d620v84YW2GSYA37fY7MPa4C4Nkdo31xXY: rando.Int64(),
		Fuq8d3t7PM5qxlNFK5FWwm735ly045Xn4UMYHHah8k0:                                                         rando.Float64(),
		BmOI4rXvqKiniOHUNMO34NuMb277Mwu7LTf0ful06Ytcd7iCJ2xK1Q4D2JcX7TI650N65783YUB0eTxc67a4Vy:              rando.Uint8(),
		PkhS2lkkE3B8in3G73q1aHo7Yu1Upflfb4nt6sygpY8e85Nw4BFwp62ajI1Ey60R1NnVTua8:                            rando.Uint16(),
		Hkk3FTs8DpVC24KoQOLi1822IkGW25HK5rgXv6FMMmwwKfLWemGc8s81103bAl1U215c2n3ulLA7C3s:                     rando.Bool(),
		R4hWoy42cWR7chFhbt6R0lIe7lbmPp5H7NB5no74l7Mj1cAX5VTT2NNkG8YuA864633lxrG8KK:                          rando.Rune(),
		XL8iHpE0umy1lsRbDg4L7PeWme0V3WI172pp40mhLNv4tab4:                                                    rando.Rune(),
		M5RmbR161X54Lk0sj5325dpStJET7Yx1hGnlC8at8v61V5G7cfVf0rts04vaMbgSD7E6236a7piTU22Jdjm7xhBqMKE5Tm:      rando.Bytes(),
		FS5UX7RRb36tD: rando.Bool(),
		LOhiG4sLjkmj045O6mEKC720nDackFgJ2Dr84W325AXRkr6LdRRAt020ktgxi52yutvH54A2s110JMmV66702Qfv3: rando.Uint8(),
		GH:                               rando.Int(),
		C7cLp60mQ:                        rando.Time(),
		Yg40LSo7V7u8Tw74RrTl1oaAaConw4K0: rando.Int(),
		Xu6EvoX1ajn:                      rando.Byte(),
		AOiq1n20yIX7Xa4IKF6a8G8jf3auHyOFA7f5fYL46axxlhK8sBnA43HcPG3jOE0fsNG7JNAL6aV0v10TwotFUX4myM: rando.Bytes(),
		HDDo06684Ac1Rn51441rJWEDyinL3w0:              rando.Int16(),
		Sxaa8sAr6103gKq:                              rando.Float64(),
		PqoM7FupFp11y1T34dCqi5J66tN6d6:               rando.Uint8(),
		GVpe8X46GqlU04KdvujDSG74p0T5i2eW1SSnvM58N104: rando.Int(),
		CHTM28Ga3O4x:                                 rando.Float64(),
		LyM3o5O50JxpRH2NfWUF5VX7j6w4Y2dvVvDdgsYaSXjWYEH56JCglDpad7Utl8o:        rando.Int(),
		D752s636j4C7jv74iHeNxU8l1dEp156JdgwwnGQc217td17IFtI4iqS0RiBXjBX2ciuV76: rando.Float64(),
		U1274wXb7a5IUUxO: rando.Uint16(),
		FaDQ4RN8cAd21K7pDhwc1OSuS5UEVAR180wt0b8LM27UA8v6r1LLP08rr7f24Xe:          rando.Int(),
		V3Nj7ylYCk32ima28f2l2DNONtbgj26185T6j600SV8X50Uc1tp85ue8y1RT:             rando.Int(),
		SbT5m67T832SU0mT37a0PF3iva7nNO7304L7ESwVck07J3SL3DpgQ1MKP:                rando.Int(),
		XkJ2t61d1r6iLr7VwCBj3JDua2eJ1N6M6bJ1mu31xe1Wl1b6QeL4R0FC075TR2y7o5sjBCPK: rando.Int64(),
		Ts7dy844yygPLam5v3xspq42UA4148pAINXxWpV5yMWF778d252giy41EQ65:             rando.Uint8(),
		Fi5tL0f7Ay6G0k0e7uthRe254MtYYgu545dRj62Q8qE7ait1b68m2o3dBp8E5qX4b:        rando.Int16(),
		CD1khpvtLa3w4q1d07:                              rando.Bool(),
		CFEHv0S2Yv8XG45f61Q5j01WANN4UHuh0810iC10:        rando.Uint(),
		Bw688I32cR6Lhwf57w68v14xUS0Ov2fI06tegH68uMs6E:   rando.Uint32(),
		LX7rAsf44c4dt25Bdn0otfLWKbID:                    rando.String(),
		Q7UjJt8s27YJtR0p6poLv67ygwY5JY4kp06oWdVW2eef2G2: rando.Uint64(),
		KFC: rando.Uint32(),
		C878YkDfnrPaPIgUp7qvnF5e77KJId812PJk6Lk3s234p4U8bxJM07iQ4liCR024R3g1dSt6p4gHV4Hh5jdFT8v4JI5: rando.Time(),
		Vwj4lDeoF7j1BS2Ppq3hyuMoJ15q5: rando.Bool(),
		Wx8f2E0lM:                     rando.Int8(),
		It5m7jSqSD44qbR3s6c4M0Vy8bN6i7kWdA75B1wNyk6tXiGA0T3q1gMbwHu3vti3tvU162J228wtcUxM0rpsayE6:            rando.Uint32(),
		I15q83lh2753E3hh5SaDR114cmvG276N07YPlpQ5sWG81Xc6LjJD53a7rTOIY1jkYy2L0YR6f:                           rando.Bool(),
		JJVVDQCp8seeHo27b41cS76yC1whY4JT50sE82Hx6h7vpjL76:                                                   rando.Float64(),
		Y345kTV8ELq4dSA4J0iF2stgO7U3k5hop733tPCuY3v6H2031i85A70j06rd3er8F0IfHMUeaOUW2xes712t3N1e0fVQ5q8:     struct{}{},
		S6FuvHuJx7Edy68k6I2xs6I8fNF6R707gYFTK8SaHL1pUMukJca7K0n7QbJLpF6602GM7151FF372ACN48hd:                rando.Uint8(),
		E873yuduPYLipIYS46OI8Bn713qC37X13Q6yRvhUfjQF7ogE874mE5ij74V08b3H3u0TBrcytBukJ7R0swi6S7H04iiQBBWi8l0: rando.Byte(),
		Efi6qL34A03n54g480mH5T67Yw7HAt403fjnH:                                                               rando.Int16(),
		BJo0TyddGaju7Rcd812QQq54LK7PAPBxgij30Mosg7435dnxL3V5JqGc8TUJjG38pD1V8325pl7MwS8f5Wod7ABMU1:          rando.Float32(),
		B0Uw: rando.Int32(),
		RkIRJlk123cwymOjqM5AD6VuIVwOL036XvrXf5LIicFsBckFbN81BFOo26oo4Kf407F4H714KL8X: rando.Int32(),
		W8p1705lnNd3abIrJJ: rando.Int32(),
		M7AJ23C2sPR11P3we8pTHXBUWXtvM1PxPBcbGn2t01Y87P: rando.Float32(),
	}
	src := actual.MarshalJ()
	require.NoError(t, expected.UnmarshalJ(src))
	require.NotEqual(t, ORjBp8Pg4Hi73lc04{}, expected)
	require.NotEqual(t, ORjBp8Pg4Hi73lc04{}, actual)
	require.Equal(t, expected, actual)
}
